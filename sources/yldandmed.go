package sources

import (
	"avaandmed/database"
	"avaandmed/utils"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gorm.io/gorm"
)

func parse[T any](
	id int64,
	yldandmed map[string]interface{},
	key string,
	items *[]T,
	parser func(int64, map[string]interface{}) T,
) {
	jsonItems, ok := yldandmed[key].([]interface{})
	if ok {
		for _, item := range jsonItems {
			*items = append(*items, parser(id, item.(map[string]interface{})))
		}
	}
}

const (
	url      = "https://avaandmed.ariregister.rik.ee/sites/default/files/avaandmed/ettevotja_rekvisiidid__yldandmed.json.zip"
	fileName = "data/downloaded_file.zip"
	jsonFile = "data/ettevotja_rekvisiidid__yldandmed.json"
)

func Yldandmed(db *gorm.DB, batchSize int) error {

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

	i := 0
	t := time.Now()
	ettevotted := make([]database.Ettevote, 0, batchSize)
	aadressid := make([]database.Aadress, 0, batchSize)
	arinimed := make([]database.Arinimi, 0, batchSize)
	kapitalid := make([]database.Kapital, 0, batchSize)
	majandusaastad := make([]database.Majandusaasta, 0, batchSize)
	markusedKaardil := make([]database.MarkusedKaardil, 0, batchSize)
	oiguslikudVormid := make([]database.OiguslikVorm, 0, batchSize)
	sidevahendid := make([]database.Sidevahend, 0, batchSize)
	staatused := make([]database.Staatus, 0, batchSize)
	teatatudTegevusalad := make([]database.TeatatudTegevusala, 0, batchSize)
	pohikirjad := make([]database.Pohikiri, 0, batchSize)

	for decoder.More() {
		var value map[string]interface{}
		err := decoder.Decode(&value)
		if err != nil {
			return fmt.Errorf("error decoding JSON: %v", err)
		}

		id := utils.Int(value["ariregistri_kood"])
		yldandmed := value["yldandmed"].(map[string]interface{})

		ettevotted = append(ettevotted, database.Ettevote{
			ID:                            id,
			Nimi:                          utils.String(value["nimi"]),
			AsutatudSissemaksetTegemata:   utils.Bool(yldandmed["asutatud_sissemakset_tegemata"]),
			EsitabKasusaajad:              utils.Bool(yldandmed["esitab_kasusaajad"]),
			EsmaregistreerimiseKpv:        utils.Date(yldandmed["esmaregistreerimise_kpv"]),
			EttevotteregistriNr:           utils.StringPointer(yldandmed["ettevotteregistri_nr"]),
			EvksRegistreeritud:            utils.BoolPointer(yldandmed["evks_registreeritud"]),
			EvksRegistreeritudKandeKpv:    utils.DatePointer(yldandmed["evks_registreeritud_kande_kpv"]),
			KustutamiseKpv:                utils.DatePointer(yldandmed["kustutamise_kpv"]),
			LahknevusteadePuudumisest:     utils.StringPointer(yldandmed["lahknevusteade_puudumisest"]),
			LoobunudVorminouetest:         utils.BoolPointer(yldandmed["loobunud_vorminouetest"]),
			OnRaamatupidamiskohustuslane:  utils.Bool(yldandmed["on_raamatupidamiskohustuslane"]),
			OiguslikVorm:                  utils.String(yldandmed["oiguslik_vorm"]),
			OiguslikVormNr:                utils.IntPointer(yldandmed["oiguslik_vorm_nr"]),
			OiguslikVormTekstina:          utils.String(yldandmed["oiguslik_vorm_tekstina"]),
			OiguslikuVormiAlaliik:         utils.StringPointer(yldandmed["oigusliku_vormi_alaliik"]),
			OiguslikuVormiAlaliikTekstina: utils.String(yldandmed["oigusliku_vormi_alaliik_tekstina"]),
			Piirkond:                      utils.IntPointer(yldandmed["piirkond"]),
			PiirkondTekstina:              utils.String(yldandmed["piirkond_tekstina"]),
			PiirkondTekstinaPikk:          utils.String(yldandmed["piirkond_tekstina_pikk"]),
			Staatus:                       utils.String(yldandmed["staatus"]),
			StaatusTekstina:               utils.String(yldandmed["staatus_tekstina"]),
			Tegutseb:                      utils.BoolPointer(yldandmed["tegutseb"]),
			TegutsebTekstina:              utils.String(yldandmed["tegutseb_tekstina"]),
		})

		parse(id, yldandmed, "aadressid", &aadressid, func(id int64, aadress map[string]interface{}) database.Aadress {
			return database.Aadress{
				EttevotteID:                             id,
				AadressAdsAdobID:                        utils.StringPointer(aadress["aadress_ads__adob_id"]),
				AadressAdsAdrID:                         utils.IntPointer(aadress["aadress_ads__adr_id"]),
				AadressAdsAdsNormaliseeritudTaisaadress: utils.StringPointer(aadress["aadress_ads__ads_normaliseeritud_taisaadress"]),
				AadressAdsAdsNormaliseeritudTaisaadressTapsustus: utils.StringPointer(aadress["aadress_ads__ads_normaliseeritud_taisaadress_tapsustus"]),
				AadressAdsAdsOid:      utils.StringPointer(aadress["aadress_ads__ads_oid"]),
				AadressAdsKoodaadress: utils.StringPointer(aadress["aadress_ads__koodaadress"]),
				AadressAdsTyyp:        utils.StringPointer(aadress["aadress_ads__tyyp"]),
				AlgusKpv:              utils.Date(aadress["algus_kpv"]),
				Ehak:                  utils.StringPointer(aadress["ehak"]),
				EhakNimetus:           utils.StringPointer(aadress["ehak_nimetus"]),
				KaardiNr:              utils.Int(aadress["kaardi_nr"]),
				KaardiPiirkond:        utils.Int(aadress["kaardi_piirkond"]),
				KaardiTyyp:            utils.String(aadress["kaardi_tyyp"]),
				KandeNr:               utils.Int(aadress["kande_nr"]),
				KirjeID:               utils.Int(aadress["kirje_id"]),
				LoppKpv:               utils.DatePointer(aadress["lopp_kpv"]),
				Postiindeks:           utils.StringPointer(aadress["postiindeks"]),
				Riik:                  utils.StringPointer(aadress["riik"]),
				RiikTekstina:          utils.StringPointer(aadress["riik_tekstina"]),
				TanavMajaKorter:       utils.StringPointer(aadress["tanav_maja_korter"]),
			}
		})

		parse(id, yldandmed, "arinimed", &arinimed, func(id int64, arinimi map[string]interface{}) database.Arinimi {
			return database.Arinimi{
				EttevotteID:    id,
				AlgusKpv:       utils.DatePointer(arinimi["algus_kpv"]),
				KaardiNr:       utils.IntPointer(arinimi["kaardi_nr"]),
				KaardiPiirkond: utils.IntPointer(arinimi["kaardi_piirkond"]),
				KaardiTyyp:     utils.StringPointer(arinimi["kaardi_tyyp"]),
				KandeNr:        utils.IntPointer(arinimi["kande_nr"]),
				KirjeID:        utils.IntPointer(arinimi["kirje_id"]),
				LoppKpv:        utils.DatePointer(arinimi["lopp_kpv"]),
				Sisu:           utils.StringPointer(arinimi["sisu"]),
			}
		})

		parse(id, yldandmed, "kapitalid", &kapitalid, func(id int64, kapital map[string]interface{}) database.Kapital {
			return database.Kapital{
				EttevotteID:             id,
				AlgusKpv:                utils.DatePointer(kapital["algus_kpv"]),
				KaardiNr:                utils.IntPointer(kapital["kaardi_nr"]),
				KaardiPiirkond:          utils.IntPointer(kapital["kaardi_piirkond"]),
				KaardiTyyp:              utils.StringPointer(kapital["kaardi_tyyp"]),
				KandeNr:                 utils.IntPointer(kapital["kande_nr"]),
				KapitaliSuurus:          utils.StringPointer(kapital["kapitali_suurus"]),
				KapitaliValuuta:         utils.StringPointer(kapital["kapitali_valuuta"]),
				KapitaliValuutaTekstina: utils.StringPointer(kapital["kapitali_valuuta_tekstina"]),
				KirjeID:                 utils.IntPointer(kapital["kirje_id"]),
				LoppKpv:                 utils.DatePointer(kapital["lopp_kpv"]),
			}
		})

		parse(id, yldandmed, "majandusaastad", &majandusaastad, func(id int64, majandusaasta map[string]interface{}) database.Majandusaasta {
			return database.Majandusaasta{
				EttevotteID:    id,
				AlgusKpv:       utils.DatePointer(majandusaasta["algus_kpv"]),
				KaardiNr:       utils.IntPointer(majandusaasta["kaardi_nr"]),
				KaardiPiirkond: utils.IntPointer(majandusaasta["kaardi_piirkond"]),
				KaardiTyyp:     utils.StringPointer(majandusaasta["kaardi_tyyp"]),
				KandeNr:        utils.IntPointer(majandusaasta["kande_nr"]),
				MajAastaAlgus:  utils.StringPointer(majandusaasta["maj_aasta_algus"]),
				MajAastaLopp:   utils.StringPointer(majandusaasta["maj_aasta_lopp"]),
				KirjeID:        utils.IntPointer(majandusaasta["kirje_id"]),
				LoppKpv:        utils.DatePointer(majandusaasta["lopp_kpv"]),
			}
		})

		parse(id, yldandmed, "markused_kaardil", &markusedKaardil, func(id int64, markusedKaardil map[string]interface{}) database.MarkusedKaardil {
			return database.MarkusedKaardil{
				EttevotteID:    id,
				AlgusKpv:       utils.DatePointer(markusedKaardil["algus_kpv"]),
				KaardiNr:       utils.IntPointer(markusedKaardil["kaardi_nr"]),
				KaardiPiirkond: utils.IntPointer(markusedKaardil["kaardi_piirkond"]),
				KaardiTyyp:     utils.StringPointer(markusedKaardil["kaardi_tyyp"]),
				KandeNr:        utils.IntPointer(markusedKaardil["kande_nr"]),
				KirjeID:        utils.Int(markusedKaardil["kirje_id"]),
				LoppKpv:        utils.DatePointer(markusedKaardil["lopp_kpv"]),
				Sisu:           utils.StringPointer(markusedKaardil["sisu"]),
				Tyyp:           utils.StringPointer(markusedKaardil["tyyp"]),
				TyypTekstina:   utils.StringPointer(markusedKaardil["tyyp_tekstina"]),
				VeergNr:        utils.IntPointer(markusedKaardil["veerg_nr"]),
			}
		})

		parse(id, yldandmed, "oiguslikud_vormid", &oiguslikudVormid, func(id int64, oiguslikVorm map[string]interface{}) database.OiguslikVorm {
			return database.OiguslikVorm{
				EttevotteID:    id,
				AlgusKpv:       utils.DatePointer(oiguslikVorm["algus_kpv"]),
				KaardiNr:       utils.IntPointer(oiguslikVorm["kaardi_nr"]),
				KaardiPiirkond: utils.IntPointer(oiguslikVorm["kaardi_piirkond"]),
				KaardiTyyp:     utils.StringPointer(oiguslikVorm["kaardi_tyyp"]),
				KandeNr:        utils.IntPointer(oiguslikVorm["kande_nr"]),
				KirjeID:        utils.Int(oiguslikVorm["kirje_id"]),
				LoppKpv:        utils.DatePointer(oiguslikVorm["lopp_kpv"]),
				Sisu:           utils.StringPointer(oiguslikVorm["sisu"]),
				SisuNr:         utils.IntPointer(oiguslikVorm["sisu_nr"]),
				SisuTekstina:   utils.StringPointer(oiguslikVorm["sisu_tekstina"]),
			}
		})

		parse(id, yldandmed, "sidevahendid", &sidevahendid, func(id int64, sidevahend map[string]interface{}) database.Sidevahend {
			return database.Sidevahend{
				EttevotteID:    id,
				KaardiNr:       utils.IntPointer(sidevahend["kaardi_nr"]),
				KaardiPiirkond: utils.IntPointer(sidevahend["kaardi_piirkond"]),
				KaardiTyyp:     utils.StringPointer(sidevahend["kaardi_tyyp"]),
				KandeNr:        utils.IntPointer(sidevahend["kande_nr"]),
				KirjeID:        utils.Int(sidevahend["kirje_id"]),
				LoppKpv:        utils.DatePointer(sidevahend["lopp_kpv"]),
				Sisu:           utils.StringPointer(sidevahend["sisu"]),
				Liik:           utils.StringPointer(sidevahend["liik"]),
				LiikTekstina:   utils.StringPointer(sidevahend["liik_tekstina"]),
			}
		})

		parse(id, yldandmed, "staatused", &staatused, func(id int64, staatus map[string]interface{}) database.Staatus {
			return database.Staatus{
				EttevotteID:     id,
				AlgusKpv:        utils.DatePointer(staatus["algus_kpv"]),
				KaardiNr:        utils.IntPointer(staatus["kaardi_nr"]),
				KaardiPiirkond:  utils.IntPointer(staatus["kaardi_piirkond"]),
				KaardiTyyp:      utils.StringPointer(staatus["kaardi_tyyp"]),
				KandeNr:         utils.IntPointer(staatus["kande_nr"]),
				Staatus:         utils.StringPointer(staatus["staatus"]),
				StaatusTekstina: utils.StringPointer(staatus["staatus_tekstina"]),
			}
		})

		parse(id, yldandmed, "teatatud_tegevusalad", &teatatudTegevusalad, func(id int64, teatatudTegevusala map[string]interface{}) database.TeatatudTegevusala {
			return database.TeatatudTegevusala{
				EttevotteID:           id,
				AlgusKpv:              utils.DatePointer(teatatudTegevusala["algus_kpv"]),
				KirjeID:               utils.Int(teatatudTegevusala["kirje_id"]),
				LoppKpv:               utils.DatePointer(teatatudTegevusala["lopp_kpv"]),
				EmtakKood:             utils.StringPointer(teatatudTegevusala["emtak_kood"]),
				EmtakTekstina:         utils.StringPointer(teatatudTegevusala["emtak_tekstina"]),
				EmtakVersioon:         utils.IntPointer(teatatudTegevusala["emtak_versioon"]),
				EmtakVersioonTekstina: utils.StringPointer(teatatudTegevusala["emtak_versioon_tekstina"]),
				NaceKood:              utils.StringPointer(teatatudTegevusala["nace_kood"]),
				OnPohitegevusala:      utils.BoolPointer(teatatudTegevusala["on_pohitegevusala"]),
			}
		})

		parse(id, yldandmed, "pohikirjad", &pohikirjad, func(id int64, pohikiri map[string]interface{}) database.Pohikiri {
			return database.Pohikiri{
				EttevotteID:       id,
				AlgusKpv:          utils.DatePointer(pohikiri["algus_kpv"]),
				KaardiNr:          utils.IntPointer(pohikiri["kaardi_nr"]),
				KaardiPiirkond:    utils.IntPointer(pohikiri["kaardi_piirkond"]),
				KaardiTyyp:        utils.StringPointer(pohikiri["kaardi_tyyp"]),
				KandeNr:           utils.IntPointer(pohikiri["kande_nr"]),
				KirjeID:           utils.Int(pohikiri["kirje_id"]),
				LoppKpv:           utils.DatePointer(pohikiri["lopp_kpv"]),
				KinnitamiseKpv:    utils.DatePointer(pohikiri["kinnitamise_kpv"]),
				MuutmiseKpv:       utils.DatePointer(pohikiri["muutmise_kpv"]),
				Selgitus:          utils.StringPointer(pohikiri["selgitus"]),
				SisaldabErioigusi: utils.BoolPointer(pohikiri["sisaldab_erioigusi"]),
			}
		})

		if len(ettevotted) == batchSize {
			i++
			fmt.Printf("Count %d, time: %d \n", i*batchSize, time.Since(t).Milliseconds())
			t = time.Now()
		}

		database.InsertBatch(db, &ettevotted, batchSize)
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
	}

	database.InsertAll(db, &ettevotted)
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

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("error reading closing bracket: %v", err)
	}

	return nil
}
