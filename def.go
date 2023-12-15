package main

import (
	"fmt"
	"os"
)

//######
// FMT
//######

func KrydderPrintln(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

func KrydderPrintf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

func KrydderSprint(a ...interface{}) string {
	return fmt.Sprint(a...)
}

//######
// OS
//######

func OreganoOpen(name string) (*os.File, error) {
	return os.Open(name)
}

func OreganoCreate(name string) (*os.File, error) {
	return os.Create(name)
}

func OreganoRemove(name string) error {
	return os.Remove(name)
}

func OreganoChdir(name string) error {
	return os.Chdir(name)
}

func OreganoChmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

func OreganoChown(name string, uid, gid int) error {
	return os.Chown(name, uid, gid)
}

func OreganoRename(old, new string) error {
	return os.Rename(old, new)
}

func OreganoMkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

func OreganoMkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func OreganoStat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
