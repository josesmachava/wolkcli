package wolk

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func ReadCurrentDir() {

	file, err := os.Open("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	r, _ := http.NewRequest("POST", "https://clownfish-app-5tgjj.ondigitalocean.app/api/upload", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			fmt.Printf("Error occur")
		}

		bodyString := string(body)
		fmt.Printf(bodyString)
	}

}
