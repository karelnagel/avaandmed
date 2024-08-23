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

		if f.ZipPath == "" {
			err := DownloadFile(f.URL, f.FilePath)
			if err != nil {
				return fmt.Errorf("error downloading: %w", err)
			}
		} else {
			err := DownloadFile(f.URL, f.ZipPath)
			if err != nil {
				return fmt.Errorf("error downloading: %w", err)
			}

			err = Unzip(f.ZipPath)
			if err != nil {
				return fmt.Errorf("error unzipping: %w", err)
			}
		}
		fmt.Println("File downloaded")
	}
	return nil
}
