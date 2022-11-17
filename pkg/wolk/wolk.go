package wolk

import (
	"fmt"
	"log"
	"os"
)

func ReadCurrentDir() {
	file, err := os.Open(".")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		if name == "index.html" {
			fmt.Println("Found html file")

		}

		fmt.Println("index.html not found in directoy")
	}
}
