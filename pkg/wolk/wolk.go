package wolk

import (
	"fmt"
)

func ReadCurrentDir() {

	filePath := getCurrentDirectory()

	if filePath == "" {
		return
	}

	htmlContent, err := readHTMLContent(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = uploadHTMLFile(filePath, htmlContent)
	if err != nil {
		fmt.Println("Upload failed:", err)
		return
	}

	fmt.Println("Upload completed successfully.")

}
