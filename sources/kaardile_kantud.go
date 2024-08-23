package sources

import (
	"avaandmed/utils"
	"encoding/json"
	"fmt"
	"os"

	"github.com/schollz/progressbar/v3"
	"gorm.io/gorm"
)

type KaardileKantudJSON struct {
	AriregistriKood      int64                    `json:"ariregistri_kood"`
	Nimi                 string                   `json:"nimi"`
	KaardileKantudIsikud []KaardileKantudIsikJSON `json:"kaardile_kantud_isikud"`
}

type KaardileKantudIsikJSON struct {
	KirjeID                                          int64    `json:"kirje_id"`
	KaardiPiirkond                                   int64    `json:"kaardi_piirkond"`
	KaardiNr                                         int64    `json:"kaardi_nr"`
	KaardiTyyp                                       string   `json:"kaardi_tyyp"`
	KandeNr                                          int64    `json:"kande_nr"`
	IsikuTyyp                                        string   `json:"isiku_tyyp"`
	IsikuRoll                                        string   `json:"isiku_roll"`
	IsikuRollTekstina                                string   `json:"isiku_roll_tekstina"`
	Eesnimi                                          *string  `json:"eesnimi"`
	NimiArinimi                                      string   `json:"nimi_arinimi"`
	IsikukoodRegistrikood                            string   `json:"isikukood_registrikood"`
	ValisKood                                        *string  `json:"valis_kood"`
	ValisKoodRiik                                    *string  `json:"valis_kood_riik"`
	ValisKoodRiikTekstina                            *string  `json:"valis_kood_riik_tekstina"`
	Synniaeg                                         *string  `json:"synniaeg"`
	Osamaks                                          *float64 `json:"osamaks"`
	OsamaksuValuuta                                  *string  `json:"osamaksu_valuuta"`
	OsamaksuValuutaTekstina                          *string  `json:"osamaksu_valuuta_tekstina"`
	VolitusteLoppemiseKpv                            *string  `json:"volituste_loppemise_kpv"`
	AadressRiik                                      *string  `json:"aadress_riik"`
	AadressRiikTekstina                              *string  `json:"aadress_riik_tekstina"`
	AadressEhak                                      *string  `json:"aadress_ehak"`
	AadressEhakTekstina                              *string  `json:"aadress_ehak_tekstina"`
	AadressTanavMajaKorter                           *string  `json:"aadress_tanav_maja_korter"`
	AadressPostiindeks                               *string  `json:"aadress_postiindeks"`
	AlgusKpv                                         string   `json:"algus_kpv"`
	LoppKpv                                          *string  `json:"lopp_kpv"`
	Email                                            *string  `json:"email"`
	AadressAdsAdrID                                  *int64   `json:"aadress_ads__adr_id"`
	AadressAdsAdsOid                                 *string  `json:"aadress_ads__ads_oid"`
	AadressAdsAdsNormaliseeritudTaisaadress          *string  `json:"aadress_ads__ads_normaliseeritud_taisaadress"`
	AadressAdsAdsNormaliseeritudTaisaadressTapsustus *string  `json:"aadress_ads__ads_normaliseeritud_taisaadress_tapsustus"`
	AadressAdsKoodaadress                            *string  `json:"aadress_ads__koodaadress"`
	AadressAdsAdobID                                 *string  `json:"aadress_ads__adob_id"`
	AadressAdsTyyp                                   *string  `json:"aadress_ads__tyyp"`
}

type KaardileKantudIsik struct {
	ID                       int `gorm:"primarykey"`
	EttevotteID              int64
	AlgusKpvInt              int64
	LoppKpvInt               *int64
	VolitusteLoppemiseKpvInt *int64
	KaardileKantudIsikJSON
}

func ParseKaardileKantud(db *gorm.DB, batchSize int) error {
	source := utils.Source{
		URL:      "https://avaandmed.ariregister.rik.ee/sites/default/files/avaandmed/ettevotja_rekvisiidid__kaardile_kantud_isikud.json.zip",
		ZipPath:  "data/ettevotja_rekvisiidid__kaardile_kantud_isikud.json.zip",
		FilePath: "data/ettevotja_rekvisiidid__kaardile_kantud_isikud.json",
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

	kaardileKantud := make([]KaardileKantudIsik, 0, batchSize)

	bar := progressbar.Default(utils.COMPANIES)
	for decoder.More() {
		bar.Add(1)
		var value KaardileKantudJSON
		decoder.Decode(&value)
		for _, isik := range value.KaardileKantudIsikud {
			kaardileKantud = append(kaardileKantud, KaardileKantudIsik{
				KaardileKantudIsikJSON:   isik,
				EttevotteID:              value.AriregistriKood,
				VolitusteLoppemiseKpvInt: utils.DatePointer(isik.VolitusteLoppemiseKpv),
				AlgusKpvInt:              utils.Date(isik.AlgusKpv),
				LoppKpvInt:               utils.DatePointer(isik.LoppKpv),
			})
		}
		InsertBatch(db, &kaardileKantud, batchSize)
	}
	InsertAll(db, &kaardileKantud)

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("error reading closing bracket: %v", err)
	}

	return nil
}
