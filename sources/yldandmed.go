package sources

import (
	"avaandmed/database"
	"avaandmed/utils"
	"encoding/json"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"gorm.io/gorm"
	"os"
)

func Yldandmed(db *gorm.DB, batchSize int) error {
	const (
		url       = "https://avaandmed.ariregister.rik.ee/sites/default/files/avaandmed/ettevotja_rekvisiidid__yldandmed.json.zip"
		fileName  = "data/ettevotja_rekvisiidid__yldandmed.zip"
		jsonFile  = "data/ettevotja_rekvisiidid__yldandmed.json"
		companies = 346698
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

	// Opening file
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

	yldandmed := make([]database.Yldandmed, 0, batchSize)
	aadressid := make([]database.Aadress, 0, batchSize)
	arinimed := make([]database.Arinimi, 0, batchSize)
	kapitalid := make([]database.Kapital, 0, batchSize)
	majandusaastad := make([]database.Majandusaasta, 0, batchSize)
	markusedKaardil := make([]database.MarkusKaardil, 0, batchSize)
	oiguslikudVormid := make([]database.OiguslikVorm, 0, batchSize)
	sidevahendid := make([]database.Sidevahend, 0, batchSize)
	staatused := make([]database.Staatus, 0, batchSize)
	teatatudTegevusalad := make([]database.TeatatudTegevusala, 0, batchSize)
	pohikirjad := make([]database.Pohikiri, 0, batchSize)
	infoMajandusaastaAruannetest := make([]database.InfoMajandusaastaAruandest, 0, batchSize)

	bar := progressbar.Default(companies)

	for decoder.More() {
		bar.Add(1)
		var value database.YldandmedFileJSON
		err := decoder.Decode(&value)
		if err != nil {
			return fmt.Errorf("error decoding JSON: %v", err)
		}

		yldandmed = append(yldandmed, database.Yldandmed{
			YldandmedJSON:                 value.Yldandmed.YldandmedJSON,
			EttevotteID:                   value.AriregistriKood,
			Nimi:                          value.Nimi,
			EsmaregistreerimiseKpvInt:     utils.Date(value.Yldandmed.EsmaregistreerimiseKpv),
			KustutamiseKpvInt:             utils.DatePointer(value.Yldandmed.KustutamiseKpv),
			EvksRegistreeritudKandeKpvInt: utils.DatePointer(value.Yldandmed.EvksRegistreeritudKandeKpv),
		})

		for _, aadress := range value.Yldandmed.Aadressid {
			aadressid = append(aadressid, database.Aadress{
				AadressJSON: aadress,
				EttevotteID: value.AriregistriKood,
				AlgusKpvInt: utils.Date(aadress.AlgusKpv),
				LoppKpvInt:  utils.DatePointer(aadress.LoppKpv),
			})
		}

		for _, arinimi := range value.Yldandmed.Arinimed {
			arinimed = append(arinimed, database.Arinimi{
				ArinimiJSON: arinimi,
				EttevotteID: value.AriregistriKood,
				AlgusKpvInt: utils.Date(arinimi.AlgusKpv),
				LoppKpvInt:  utils.DatePointer(arinimi.LoppKpv),
			})
		}

		for _, kapital := range value.Yldandmed.Kapitalid {
			kapitalid = append(kapitalid, database.Kapital{
				KapitalJSON: kapital,
				EttevotteID: value.AriregistriKood,
				AlgusKpvInt: utils.Date(kapital.AlgusKpv),
				LoppKpvInt:  utils.DatePointer(kapital.LoppKpv),
			})
		}

		for _, majandusaasta := range value.Yldandmed.Majandusaastad {
			majandusaastad = append(majandusaastad, database.Majandusaasta{
				MajandusaastaJSON: majandusaasta,
				EttevotteID:       value.AriregistriKood,
				AlgusKpvInt:       utils.Date(majandusaasta.AlgusKpv),
				LoppKpvInt:        utils.DatePointer(majandusaasta.LoppKpv),
			})
		}

		for _, pohikiri := range value.Yldandmed.Pohikirjad {
			pohikirjad = append(pohikirjad, database.Pohikiri{
				PohikiriJSON: pohikiri,
				EttevotteID:  value.AriregistriKood,
				AlgusKpvInt:  utils.Date(pohikiri.AlgusKpv),
				LoppKpvInt:   utils.DatePointer(pohikiri.LoppKpv),
			})
		}

		for _, staatus := range value.Yldandmed.Staatused {
			staatused = append(staatused, database.Staatus{
				StaatusJSON: staatus,
				EttevotteID: value.AriregistriKood,
				AlgusKpvInt: utils.Date(staatus.AlgusKpv),
			})
		}

		for _, teatatudTegevusala := range value.Yldandmed.TeatatudTegevusalad {
			teatatudTegevusalad = append(teatatudTegevusalad, database.TeatatudTegevusala{
				TeatatudTegevusalaJSON: teatatudTegevusala,
				EttevotteID:            value.AriregistriKood,
				AlgusKpvInt:            utils.Date(teatatudTegevusala.AlgusKpv),
				LoppKpvInt:             utils.DatePointer(teatatudTegevusala.LoppKpv),
			})
		}

		for _, info := range value.Yldandmed.InfoMajandusaastaAruandestJSON {
			infoMajandusaastaAruannetest = append(infoMajandusaastaAruannetest, database.InfoMajandusaastaAruandest{
				InfoMajandusaastaAruandestJSON:  info,
				EttevotteID:                     value.AriregistriKood,
				MajandusaastaPeriodiAlgusKpvInt: utils.Date(info.MajandusaastaPeriodiAlgusKpv),
				MajandusaastaPeriodiLoppKpvInt:  utils.DatePointer(&info.MajandusaastaPeriodiLoppKpv),
			})
		}

		database.InsertBatch(db, &yldandmed, batchSize)
		database.InsertBatch(db, &aadressid, batchSize)
		database.InsertBatch(db, &arinimed, batchSize)
		database.InsertBatch(db, &kapitalid, batchSize)
		database.InsertBatch(db, &majandusaastad, batchSize)
		database.InsertBatch(db, &markusedKaardil, batchSize)
		database.InsertBatch(db, &oiguslikudVormid, batchSize)
		database.InsertBatch(db, &sidevahendid, batchSize)
		database.InsertBatch(db, &staatused, batchSize)
		database.InsertBatch(db, &teatatudTegevusalad, batchSize)
		database.InsertBatch(db, &pohikirjad, batchSize)
		database.InsertBatch(db, &infoMajandusaastaAruannetest, batchSize)
	}

	database.InsertAll(db, &yldandmed)
	database.InsertAll(db, &aadressid)
	database.InsertAll(db, &arinimed)
	database.InsertAll(db, &kapitalid)
	database.InsertAll(db, &majandusaastad)
	database.InsertAll(db, &markusedKaardil)
	database.InsertAll(db, &oiguslikudVormid)
	database.InsertAll(db, &sidevahendid)
	database.InsertAll(db, &staatused)
	database.InsertAll(db, &teatatudTegevusalad)
	database.InsertAll(db, &pohikirjad)
	database.InsertAll(db, &infoMajandusaastaAruannetest)

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("error reading closing bracket: %v", err)
	}

	return nil
}
