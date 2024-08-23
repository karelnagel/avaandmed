package sources

import (
	"avaandmed/utils"
	"encoding/csv"
	"fmt"
	"gorm.io/gorm"
	"io"
	"os"
)

type EMTA struct {
	EttevotteID                           string `gorm:"primaryKey"`
	Liik                                  string
	RegistreeritudKaibemaksukohustuslaste string
	EMTAKTegevusvaldkond                  string
	Maakond                               string
	RiiklikudMaksud                       float64
	ToojouMaksudJaMaksed                  float64
	Kaive                                 float64
	Tootajaid                             int
	Year                                  int `gorm:"primaryKey"`
	Quarter                               int `gorm:"primaryKey"`
}

type Quarter struct {
	Year    int
	Quarter int
	ID      string
}

func ParseEMTA(db *gorm.DB, batchSize int) error {
	quarters := []Quarter{
		{Year: 2020, Quarter: 1, ID: "ZoL8tcJzG9fYydR"},
		{Year: 2020, Quarter: 2, ID: "p4GYRXiFAzn8998"},
		{Year: 2020, Quarter: 3, ID: "AdZ9PBMtP757ed4"},
		{Year: 2020, Quarter: 4, ID: "Lgf2ykEHJpBtidL"},
		{Year: 2021, Quarter: 1, ID: "xoJwzGRJZKZJ5Dg"},
		{Year: 2021, Quarter: 2, ID: "nJYASodEoqYLEc2"},
		{Year: 2021, Quarter: 3, ID: "jcgm4WLirRGDsNf"},
		{Year: 2021, Quarter: 4, ID: "5Eeq9pzAmQtraJn"},
		{Year: 2022, Quarter: 1, ID: "HoGwciekJozgwNr"},
		{Year: 2022, Quarter: 2, ID: "jYYAz7rQHcNXHSx"},
		{Year: 2022, Quarter: 3, ID: "ctJqwY6EBrWNPEb"},
		{Year: 2022, Quarter: 4, ID: "EMjnxk9yXJpbSDq"},
		{Year: 2023, Quarter: 1, ID: "dzY8diTB6k6wt5J"},
		{Year: 2023, Quarter: 2, ID: "4KxRo99pznKMQ3j"},
		{Year: 2023, Quarter: 3, ID: "tJCj8SwFyNDmwW2"},
		{Year: 2023, Quarter: 4, ID: "tJRx2FDL4LERtcK"},
		{Year: 2024, Quarter: 1, ID: "4eHAiZcnygNDPHC"},
		{Year: 2024, Quarter: 2, ID: "KmMp6iEsJabKz8x"},
	}

	emtas := make([]EMTA, 0, batchSize)
	bar := utils.NewProgressBar(2930436, "Processing EMTA")

	for _, quarter := range quarters {
		QUARTERS := []string{"I", "II", "III", "IV"}
		q := QUARTERS[quarter.Quarter-1]
		source := utils.Source{
			URL:      fmt.Sprintf("https://ncfailid.emta.ee/s/%s/download/tasutud_maksud_%d_%s_kvartal.csv", quarter.ID, quarter.Year, q),
			FilePath: fmt.Sprintf("data/emta_%d_%s.csv", quarter.Year, q),
			ZipPath:  "",
		}
		err := source.Download()
		if err != nil {
			return err
		}

		file, _ := os.Open(source.FilePath)
		defer file.Close()

		utfReader := utils.NewUTF8Reader(file)
		reader := csv.NewReader(utfReader)

		reader.LazyQuotes = true
		reader.Comma = ';'
		reader.FieldsPerRecord = -1

		isFirst := true
		for {
			bar.Add(1)
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if isFirst {
				isFirst = false
				continue
			}
			if err != nil {
				return fmt.Errorf("error reading CSV record: %v", err)
			}

			emtas = append(emtas, EMTA{
				EttevotteID:                           record[0],
				Liik:                                  record[2],
				RegistreeritudKaibemaksukohustuslaste: record[3],
				EMTAKTegevusvaldkond:                  record[4],
				Maakond:                               record[5],
				RiiklikudMaksud:                       utils.ParseFloat(record[6]),
				ToojouMaksudJaMaksed:                  utils.ParseFloat(record[7]),
				Kaive:                                 utils.ParseFloat(record[8]),
				Tootajaid:                             utils.ParseInt(record[9]),
				Year:                                  quarter.Year,
				Quarter:                               quarter.Quarter,
			})

			InsertBatch(db, &emtas, batchSize)
		}
	}
	InsertAll(db, &emtas)
	return nil
}
