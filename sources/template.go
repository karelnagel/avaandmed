package sources

import (
	"avaandmed/utils"
	"encoding/json"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"gorm.io/gorm"
	"os"
)

func Template(db *gorm.DB, batchSize int) error {
	const (
		url       = "https://avaandmed.ariregister.rik.ee/sites/default/files/avaandmed/ettevotja_rekvisiidid__kaardile_kantud_isikud.json.zip"
		fileName  = "data/ettevotja_rekvisiidid__kaardile_kantud_isikud.json.zip"
		jsonFile  = "data/ettevotja_rekvisiidid__kaardile_kantud_isikud.json"
		companies = 400_000
	)
	// Downloading
	if _, err := os.Stat(jsonFile); os.IsNotExist(err) {
		fmt.Println("File does not exist, downloading")
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
	}

	file, err := os.Open(jsonFile)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("error reading opening bracket: %v", err)
	}

	bar := progressbar.Default(companies)

	for decoder.More() {
		bar.Add(1)
		var value map[string]interface{}
		err := decoder.Decode(&value)
		if err != nil {
			return fmt.Errorf("error decoding JSON: %v", err)
		}

		id := utils.Int(value["ariregistri_kood"])
		fmt.Println(id)
	}

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("error reading closing bracket: %v", err)
	}

	return nil
}
