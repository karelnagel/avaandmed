package sources

import (
	"avaandmed/utils"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"os"
)

type OsanikudJSON struct {
	AriregistriKood                     int64                                `json:"ariregistri_kood"`
	Nimi                                string                               `json:"nimi"`
	Osanikud                            []OsanikJSON                         `json:"osanikud"`
	OsapandidTingimuslikudVoorandamised []OsapantTingimuslikVoorandamineJSON `json:"osapandid_tingimuslikud_voorandamised"`
}

type OsanikJSON struct {
	KirjeID                                          int64   `json:"kirje_id"`
	IsikuTyyp                                        string  `json:"isiku_tyyp"`
	IsikuRoll                                        string  `json:"isiku_roll"`
	IsikuRollTekstina                                string  `json:"isiku_roll_tekstina"`
	Eesnimi                                          *string `json:"eesnimi"`
	NimiArinimi                                      string  `json:"nimi_arinimi"`
	IsikukoodRegistrikood                            string  `json:"isikukood_registrikood"`
	ValisKood                                        *string `json:"valis_kood"`
	ValisKoodRiikTekstina                            *string `json:"valis_kood_riik_tekstina"`
	ValisKoodRiik                                    *string `json:"valis_kood_riik"`
	Synniaeg                                         *string `json:"synniaeg"`
	AadressRiik                                      *string `json:"aadress_riik"`
	AadressRiikTekstina                              *string `json:"aadress_riik_tekstina"`
	AadressEhak                                      *string `json:"aadress_ehak"`
	AadressEhakTekstina                              string  `json:"aadress_ehak_tekstina"`
	AadressTanavMajaKorter                           *string `json:"aadress_tanav_maja_korter"`
	OsaluseProtsent                                  string  `json:"osaluse_protsent"`
	OsaluseSuurus                                    string  `json:"osaluse_suurus"`
	OsaluseValuuta                                   string  `json:"osaluse_valuuta"`
	OsamaksuValuutaTekstina                          string  `json:"osamaksu_valuuta_tekstina"`
	OsaluseOmandiliik                                string  `json:"osaluse_omandiliik"`
	OsaluseOmandiliikTekstina                        string  `json:"osaluse_omandiliik_tekstina"`
	OsaluseMurdosaLugeja                             *string `json:"osaluse_murdosa_lugeja"`
	OsaluseMurdosaNimetaja                           *string `json:"osaluse_murdosa_nimetaja"`
	VolitusteLoppemiseKpv                            *string `json:"volituste_loppemise_kpv"`
	KontrolliAllikas                                 *string `json:"kontrolli_allikas"`
	KontrolliAllikasTekstina                         *string `json:"kontrolli_allikas_tekstina"`
	KontrolliAllikaKpv                               *string `json:"kontrolli_allika_kpv"`
	AlgusKpv                                         string  `json:"algus_kpv"`
	LoppKpv                                          *string `json:"lopp_kpv"`
	Grupp                                            int     `json:"grupp"`
	AadressAdsAdrID                                  *string `json:"aadress_ads__adr_id"`
	AdressAdsAdsOid                                  *string `json:"adress_ads__ads_oid"`
	AadressAdsAdsNormaliseeritudTaisaadress          *string `json:"aadress_ads__ads_normaliseeritud_taisaadress"`
	AadressAdsAdsNormaliseeritudTaisaadressTapsustus *string `json:"aadress_ads__ads_normaliseeritud_taisaadress_tapsustus"`
	AadressAdsKoodaadress                            *string `json:"aadress_ads__koodaadress"`
	AadressAdsAdobID                                 *string `json:"aadress_ads__adob_id"`
	AadressAdsTyyp                                   *string `json:"aadress_ads__tyyp"`
}

type OsapantTingimuslikVoorandamineJSON struct {
}

type Osanik struct {
	ID                       int `gorm:"primarykey"`
	EttevotteID              int64
	AlgusKpvInt              int64
	LoppKpvInt               *int64
	VolitusteLoppemiseKpvInt *int64
	KontrolliAllikaKpvInt    *int64
	OsanikJSON
}

func ParseOsanikud(db *gorm.DB, batchSize int) error {
	source := utils.Source{
		URL:      "https://avaandmed.ariregister.rik.ee/sites/default/files/avaandmed/ettevotja_rekvisiidid__osanikud.json.zip",
		ZipPath:  "data/ettevotja_rekvisiidid__osanikud.json.zip",
		FilePath: "data/ettevotja_rekvisiidid__osanikud.json",
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

	osanikud := make([]Osanik, 0, batchSize)
	isikud := make([]Isik, 0, batchSize)

	bar := utils.NewProgressBar(utils.COMPANIES, "Processing Osanikud")
	for decoder.More() {
		bar.Add(1)
		var value OsanikudJSON
		decoder.Decode(&value)
		for _, isik := range value.Osanikud {
			osanikud = append(osanikud, Osanik{
				OsanikJSON:               isik,
				EttevotteID:              value.AriregistriKood,
				AlgusKpvInt:              utils.Date(isik.AlgusKpv),
				LoppKpvInt:               utils.DatePointer(isik.LoppKpv),
				VolitusteLoppemiseKpvInt: utils.DatePointer(isik.VolitusteLoppemiseKpv),
				KontrolliAllikaKpvInt:    utils.DatePointer(isik.KontrolliAllikaKpv),
			})
			if isik.IsikuTyyp == "F" {
				isik := CreateIsik(&isik.IsikukoodRegistrikood, isik.Eesnimi, &isik.NimiArinimi)
				if isik != nil {
					isikud = append(isikud, *isik)
				}
			}
		}
		InsertBatch(db, &osanikud, batchSize)
		InsertBatch(db, &isikud, batchSize)
	}
	InsertAll(db, &osanikud)
	InsertAll(db, &isikud)

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("error reading closing bracket: %v", err)
	}

	return nil
}
