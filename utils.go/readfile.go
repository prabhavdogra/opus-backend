package utils

import "os"

func ReadFile(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	// read the file
	fileSize := stat.Size()
	bytes := make([]byte, fileSize)
	_, err = file.Read(bytes)
	if err != nil {
		panic(err)
	}

	return bytes
}
