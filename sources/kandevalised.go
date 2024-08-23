package sources

import (
	"avaandmed/utils"
	"encoding/json"
	"fmt"
	"os"
	"gorm.io/gorm"
)

type KandevalisedJSON struct {
	AriregistriKood     int64                 `json:"ariregistri_kood"`
	Nimi                string                `json:"nimi"`
	KaardivalisedIsikud []KandevalineIsikJSON `json:"kaardivalised_isikud"`
}

type KandevalineIsikJSON struct {
	KirjeID                                          int64   `json:"kirje_id"`
	IsikuTyyp                                        string  `json:"isiku_tyyp"`
	IsikuRoll                                        string  `json:"isiku_roll"`
	IsikuRollTekstina                                string  `json:"isiku_roll_tekstina"`
	Eesnimi                                          string  `json:"eesnimi"`
	NimiArinimi                                      string  `json:"nimi_arinimi"`
	IsikukoodRegistrikood                            string  `json:"isikukood_registrikood"`
	ValisKood                                        *string `json:"valis_kood"`
	ValisKoodRiikTekstina                            *string `json:"valis_kood_riik_tekstina"`
	ValisKoodRiik                                    *string `json:"valis_kood_riik"`
	Synniaeg                                         *string `json:"synniaeg"`
	AadressRiik                                      string  `json:"aadress_riik"`
	AadressRiikTekstina                              string  `json:"aadress_riik_tekstina"`
	AadressEhak                                      *string `json:"aadress_ehak"`
	AadressEhakTekstina                              string  `json:"aadress_ehak_tekstina"`
	AadressTanavMajaKorter                           *string `json:"aadress_tanav_maja_korter"`
	OsaluseProtsent                                  *string `json:"osaluse_protsent"`
	OsaluseSuurus                                    string  `json:"osaluse_suurus"`
	OsaluseValuuta                                   string  `json:"osaluse_valuuta"`
	OsamaksuValuutaTekstina                          string  `json:"osamaksu_valuuta_tekstina"`
	OsaluseOmandiliik                                string  `json:"osaluse_omandiliik"`
	OsaluseOmandiliikTekstina                        string  `json:"osaluse_omandiliik_tekstina"`
	OsaluseMurdosaLugeja                             *string `json:"osaluse_murdosa_lugeja"`
	OsaluseMurdosaNimetaja                           *string `json:"osaluse_murdosa_nimetaja"`
	VolitusteLoppemiseKpv                            string  `json:"volituste_loppemise_kpv"`
	KontrolliAllikas                                 string  `json:"kontrolli_allikas"`
	KontrolliAllikasTekstina                         string  `json:"kontrolli_allikas_tekstina"`
	KontrolliAllikaKpv                               string  `json:"kontrolli_allika_kpv"`
	AlgusKpv                                         string  `json:"algus_kpv"`
	LoppKpv                                          *string `json:"lopp_kpv"`
	Grupp                                            *string `json:"grupp"`
	AadressAdsAdrID                                  *int64  `json:"aadress_ads__adr_id"`
	AadressAdsAdsOid                                 *string `json:"adress_ads__ads_oid"`
	AadressAdsAdsNormaliseeritudTaisaadress          *string `json:"aadress_ads__ads_normaliseeritud_taisaadress"`
	AadressAdsAdsNormaliseeritudTaisaadressTapsustus *string `json:"aadress_ads__ads_normaliseeritud_taisaadress_tapsustus"`
	AadressAdsKoodaadress                            *string `json:"aadress_ads__koodaadress"`
	AadressAdsAdobID                                 *string `json:"aadress_ads__adob_id"`
	AadressAdsTyyp                                   *string `json:"aadress_ads__tyyp"`
}

type KandevalineIsik struct {
	ID                       int `gorm:"primarykey"`
	EttevotteID              int64
	AlgusKpvInt              int64
	LoppKpvInt               *int64
	VolitusteLoppemiseKpvInt *int64
	KontrolliAllikaKpvInt    *int64
	KandevalineIsikJSON
}

func ParseKandevalised(db *gorm.DB, batchSize int) error {
	source := utils.Source{
		URL:      "https://avaandmed.ariregister.rik.ee/sites/default/files/avaandmed/ettevotja_rekvisiidid__kandevalised_isikud.json.zip",
		ZipPath:  "data/ettevotja_rekvisiidid__kandevalised_isikud.json.zip",
		FilePath: "data/ettevotja_rekvisiidid__kandevalised_isikud.json",
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

	kandevalised := make([]KandevalineIsik, 0, batchSize)

	bar := utils.NewProgressBar(utils.COMPANIES, "Processing Kandevalised")
	for decoder.More() {
		bar.Add(1)
		var value KandevalisedJSON
		decoder.Decode(&value)
		for _, isik := range value.KaardivalisedIsikud {
			kandevalised = append(kandevalised, KandevalineIsik{
				KandevalineIsikJSON:      isik,
				EttevotteID:              value.AriregistriKood,
				AlgusKpvInt:              utils.Date(isik.AlgusKpv),
				LoppKpvInt:               utils.DatePointer(isik.LoppKpv),
				VolitusteLoppemiseKpvInt: utils.DatePointer(&isik.VolitusteLoppemiseKpv),
				KontrolliAllikaKpvInt:    utils.DatePointer(&isik.KontrolliAllikaKpv),
			})
		}
		InsertBatch(db, &kandevalised, batchSize)
	}
	InsertAll(db, &kandevalised)

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("error reading closing bracket: %v", err)
	}

	return nil
}
