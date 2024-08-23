package utils

import (
	"fmt"
	"os"
)

type Source struct {
	URL      string
	ZipPath  string
	FilePath string
}

func (f *Source) Download() error {
	if _, err := os.Stat(f.FilePath); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist, downloading\n", f.FilePath)
		err := DownloadFile(f.URL, f.ZipPath)
		if err != nil {
			return fmt.Errorf("error downloading: %w", err)
		}
		fmt.Println("File downloaded")

		err = Unzip(f.ZipPath)
		if err != nil {
			return fmt.Errorf("error unzipping: %w", err)
		}
		fmt.Println("File unzipped")
	}
	return nil
}
