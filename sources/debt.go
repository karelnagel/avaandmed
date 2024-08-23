package sources

import (
	"avaandmed/utils"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
	"gorm.io/gorm"
)

type Maksuvolg struct {
	EttevotteID     string `gorm:"primaryKey"`
	Maksuvõlg       int
	ShVaidlustatud  int
	MaksuvõlgAlates *int64
}

func ParseDebt(db *gorm.DB, batchSize int) error {
	source := utils.Source{
		URL:      "https://ncfailid.emta.ee/index.php/s/BbzKpE7g6WKT3rM/download/maksuvolglaste_nimekiri.csv",
		FilePath: "data/debt.csv",
		ZipPath:  "",
	}
	err := source.Download()
	if err != nil {
		return err
	}

	maksuvolad := make([]Maksuvolg, 0, batchSize)
	bar := progressbar.Default(2930436)

	file, _ := os.Open(source.FilePath)
	defer file.Close()

	utfReader := utils.NewUTF8Reader(file)
	reader := csv.NewReader(utfReader)

	reader.LazyQuotes = true
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	linesToSkip := 3
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

		maksuvolad = append(maksuvolad, Maksuvolg{
			EttevotteID:     record[0],
			Maksuvõlg:       utils.ParseInt(record[2]),
			ShVaidlustatud:  utils.ParseInt(record[3]),
			MaksuvõlgAlates: utils.DatePointer(&record[4]),
		})
		InsertBatch(db, &maksuvolad, batchSize)
	}
	InsertAll(db, &maksuvolad)
	return nil
}
