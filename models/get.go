package models

import (
	"fmt"
	"os"
)

func CeeatePathIfNotExists(path string) {
	_, err := os.Stat(path)
	if err == nil {
		fmt.Printf("dir %s 存在\n", path)
	}
	if os.IsNotExist(err) {
		fmt.Printf("dir %s 不存在\n", path)
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
}
