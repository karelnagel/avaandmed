package sources

import (
	"avaandmed/database"
	"avaandmed/utils"
	"encoding/json"
	"fmt"
	"os"

	"github.com/schollz/progressbar/v3"
	"gorm.io/gorm"
)

func Kasusaajad(db *gorm.DB, batchSize int) error {
	const (
		url       = "https://avaandmed.ariregister.rik.ee/sites/default/files/avaandmed/ettevotja_rekvisiidid__kasusaajad.json.zip"
		fileName  = "data/ettevotja_rekvisiidid__kasusaajad.json.zip"
		jsonFile  = "data/ettevotja_rekvisiidid__kasusaajad.json"
		companies = 345930
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

	kasusaajad := make([]database.Kasusaaja, 0, batchSize)

	bar := progressbar.Default(companies)
	for decoder.More() {
		bar.Add(1)
		var value database.KasusaajadJSON
		decoder.Decode(&value)
		for _, isik := range value.Kasusaajad {
			kasusaajad = append(kasusaajad, database.Kasusaaja{
				KasusaajaJSON: isik,
				EttevotteID:   value.AriregistriKood,
				AlgusKpvInt:   utils.Date(isik.AlgusKpv),
				LoppKpvInt:    utils.DatePointer(isik.LoppKpv),
			})
		}
		database.InsertBatch(db, &kasusaajad, batchSize)
	}
	database.InsertAll(db, &kasusaajad)

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("error reading closing bracket: %v", err)
	}

	return nil
}
