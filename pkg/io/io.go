package io

import (
	"os"
	"os/exec"
)

// FileExists verifies if a directory or a file exists
func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

// FileSize get file size if a directory or a file exists
func FileSize(fileName string) (int64, error) {
	f, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return 0, err
	}
	return f.Size(), nil
}

// Copy Copy a directory or a file from origen to a specific destiny
// (for ex: insidy the rootFS of a container)
func Copy(source, destination string, done chan bool) error {
	if !FileExists(destination) {
		err := os.MkdirAll(destination, 0777)
		if err != nil {
			return err
		}
	}
	cmd := exec.Command("cp", "-r", source, destination)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	done <- true
	return err
}

// Untar ...
func Untar(directory, file string) error {
	cmd := exec.Command("tar", "-C", directory, "-xf", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Tar ...
func Tar(directory, file string, done chan bool) error {
	if err := os.Chdir(directory); err != nil {
		return err
	}
	cmd := exec.Command("tar", "cpjf", file, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	done <- true
	return err
}

// Wget ...
func Wget(url, filepath string) error {
	// run shell `wget URL -O filepath`
	cmd := exec.Command("wget", url, "-O", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
