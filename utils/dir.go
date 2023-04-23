package utils

import "os"

func GetTopDirectory() string {
	TopDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return TopDir
}
