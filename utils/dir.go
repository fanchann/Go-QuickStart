package utils

import "os"

func GetWorkingDirectory() string {
	directory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return directory
}
