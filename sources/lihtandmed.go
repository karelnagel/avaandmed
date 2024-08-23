package sources

import (
	"avaandmed/utils"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"gorm.io/gorm"
)

type Lihtandmed struct {
	Nimi                          string
	EttevotteID                   string `gorm:"primaryKey"`
	OiguslikVorm                  string
	OiguslikuVormiAlaliik         string
	KmKrNr                        string
	Staatus                       string
	StaatusTekstina               string
	EsmakandeKpv                  string
	EsmakandeKpvInt               *int64
	EttevotjaAadress              string
	AsukohtEttevotjaAadressis     string
	AsukohtEhaKood                string
	AsukohtEhaTekstina            string
	IndeksEttevotjaAadressis      string
	AdsAdrId                      string
	AdsAdsOid                     string
	AdsNormaliseeritudTaisaadress string
	TeabesysteemiLink             string
}

func ParseLihtandmed(db *gorm.DB, batchSize int) error {
	source := utils.Source{
		URL:      "https://avaandmed.ariregister.rik.ee/sites/default/files/avaandmed/ettevotja_rekvisiidid__lihtandmed.csv.zip",
		ZipPath:  "data/ettevotja_rekvisiidid__lihtandmed.csv.zip",
		FilePath: "data/ettevotja_rekvisiidid__lihtandmed.csv",
	}
	err := source.Download()
	if err != nil {
		return err
	}

	file, _ := os.Open(source.FilePath)
	defer file.Close()

	reader := csv.NewReader(file)

	reader.LazyQuotes = true
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	linesToSkip := 1

	lihtandmed := make([]Lihtandmed, 0, batchSize)
	bar := utils.NewProgressBar(utils.COMPANIES, "Processing Lihtandmed")
	for {
		bar.Add(1)
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if linesToSkip > 0 {
			linesToSkip--
			continue
		}

		if err != nil {
			return fmt.Errorf("error reading CSV record: %v", err)
		}

		// fmt.Println(record)
		lihtandmed = append(lihtandmed, Lihtandmed{
			Nimi:                          record[0],
			EttevotteID:                   record[1],
			OiguslikVorm:                  record[2],
			OiguslikuVormiAlaliik:         record[3],
			KmKrNr:                        record[4],
			Staatus:                       record[5],
			StaatusTekstina:               record[6],
			EsmakandeKpv:                  record[7],
			EsmakandeKpvInt:               utils.DatePointer(&record[7]),
			EttevotjaAadress:              record[8],
			AsukohtEttevotjaAadressis:     record[9],
			AsukohtEhaKood:                record[10],
			AsukohtEhaTekstina:            record[11],
			IndeksEttevotjaAadressis:      record[12],
			AdsAdrId:                      record[13],
			AdsAdsOid:                     record[14],
			AdsNormaliseeritudTaisaadress: record[15],
			TeabesysteemiLink:             record[16],
		})
		InsertBatch(db, &lihtandmed, batchSize)
	}
	InsertAll(db, &lihtandmed)
	return nil
}
