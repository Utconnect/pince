package utils

import (
	"fmt"
	"os"
)

func WriteFile(content string, fileName string, path string) {
	file, err := os.OpenFile(path+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.WriteString(content)

	if err != nil {
		fmt.Println(err)
		err := file.Close()
		if err != nil {
			return
		}
		return
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
