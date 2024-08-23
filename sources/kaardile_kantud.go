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

func KaardileKantud(db *gorm.DB, batchSize int) error {
	const (
		url       = "https://avaandmed.ariregister.rik.ee/sites/default/files/avaandmed/ettevotja_rekvisiidid__kaardile_kantud_isikud.json.zip"
		fileName  = "data/ettevotja_rekvisiidid__kaardile_kantud_isikud.json.zip"
		jsonFile  = "data/ettevotja_rekvisiidid__kaardile_kantud_isikud.json"
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

	kaardileKantud := make([]database.KaardileKantudIsik, 0, batchSize)

	bar := progressbar.Default(companies)
	for decoder.More() {
		bar.Add(1)
		var value database.KaardileKantudJSON
		decoder.Decode(&value)
		for _, isik := range value.KaardileKantudIsikud {
			kaardileKantud = append(kaardileKantud, database.KaardileKantudIsik{
				KaardileKantudIsikJSON:   isik,
				EttevotteID:              value.AriregistriKood,
				VolitusteLoppemiseKpvInt: utils.DatePointer(isik.VolitusteLoppemiseKpv),
				AlgusKpvInt:              utils.Date(isik.AlgusKpv),
				LoppKpvInt:               utils.DatePointer(isik.LoppKpv),
			})
		}
		database.InsertBatch(db, &kaardileKantud, batchSize)
	}
	database.InsertAll(db, &kaardileKantud)

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("error reading closing bracket: %v", err)
	}

	return nil
}
