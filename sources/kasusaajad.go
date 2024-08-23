package sources

import (
	"avaandmed/utils"
	"encoding/json"
	"fmt"
	"os"
	"gorm.io/gorm"
)

type KasusaajadJSON struct {
	AriregistriKood int64         `json:"ariregistri_kood"`
	Nimi            string        `json:"nimi"`
	Kasusaajad      []KasusaajaJSON `json:"kasusaajad"`
}

type KasusaajaJSON struct {
	KirjeID                       int64   `json:"kirje_id"`
	AlgusKpv                      string  `json:"algus_kpv"`
	LoppKpv                       *string `json:"lopp_kpv"`
	Eesnimi                       string  `json:"eesnimi"`
	Nimi                          string  `json:"nimi"`
	Isikukood                     string  `json:"isikukood"`
	ValisKood                     *string `json:"valis_kood"`
	ValisKoodRiik                 *string `json:"valis_kood_riik"`
	ValisKoodRiikTekstina         *string `json:"valis_kood_riik_tekstina"`
	Synniaeg                      *string `json:"synniaeg"`
	AadressRiik                   string  `json:"aadress_riik"`
	AadressRiikTekstina           string  `json:"aadress_riik_tekstina"`
	KontrolliTeostamiseViis       string  `json:"kontrolli_teostamise_viis"`
	KontrolliTeostamiseViisTekstina string `json:"kontrolli_teostamise_viis_tekstina"`
	LahknevusteadeEsitatud        *string `json:"lahknevusteade_esitatud"`
}

type Kasusaaja struct {
	ID           int `gorm:"primarykey"`
	EttevotteID  int64
	AlgusKpvInt  int64
	LoppKpvInt   *int64
	KasusaajaJSON
}

func ParseKasusaajad(db *gorm.DB, batchSize int) error {
	source := utils.Source{
		URL:      "https://avaandmed.ariregister.rik.ee/sites/default/files/avaandmed/ettevotja_rekvisiidid__kasusaajad.json.zip",
		ZipPath:  "data/ettevotja_rekvisiidid__kasusaajad.json.zip",
		FilePath: "data/ettevotja_rekvisiidid__kasusaajad.json",
	}
	err := source.Download()
	if err != nil {
		return fmt.Errorf("error downloading: %w", err)
	}

	file, err := os.Open(source.FilePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("error reading opening bracket: %v", err)
	}

	kasusaajad := make([]Kasusaaja, 0, batchSize)

	bar := utils.NewProgressBar(utils.COMPANIES, "Processing Kasusaajad")
	for decoder.More() {
		bar.Add(1)
		var value KasusaajadJSON
		decoder.Decode(&value)
		for _, isik := range value.Kasusaajad {
			kasusaajad = append(kasusaajad, Kasusaaja{
				KasusaajaJSON: isik,
				EttevotteID:   value.AriregistriKood,
				AlgusKpvInt:   utils.Date(isik.AlgusKpv),
				LoppKpvInt:    utils.DatePointer(isik.LoppKpv),
			})
		}
		InsertBatch(db, &kasusaajad, batchSize)
	}
	InsertAll(db, &kasusaajad)

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("error reading closing bracket: %v", err)
	}

	return nil
}
