package main

import (
	"avaandmed/utils"
	"fmt"
	"os"
)

const (
	folder   = "data"
	url      = "https://avaandmed.ariregister.rik.ee/sites/default/files/avaandmed/ettevotja_rekvisiidid__yldandmed.json.zip"
	fileName = "data/downloaded_file.zip"
	jsonFile = "data/ettevotja_rekvisiidid__yldandmed.json"
)

func yldandmed() error {
	fmt.Println("Hello, World!")

	if err := os.RemoveAll(folder); err != nil {
		return fmt.Errorf("error removing data directory: %w", err)
	}
	if err := os.Mkdir(folder, 0755); err != nil {
		return fmt.Errorf("error creating data directory: %w", err)
	}

	err := utils.DownloadFile(url, fileName)
	if err != nil {
		return fmt.Errorf("error downloading: %w", err)
	}
	fmt.Println("File downloaded")

	err = utils.Unzip(fileName)
	if err != nil {
		return fmt.Errorf("error unzipping: %w", err)
	}
	fmt.Println("File unzipped")

	return nil
}

func main() {
	err := ParseYldandmed(jsonFile)
	if err != nil {
		fmt.Println("Error parsing JSON: %w", err)
	}
}
