package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
)

type Tabler interface {
	TableName() string
}

type Ettevote struct {
	ID                            int64 `gorm:"primarykey"`
	Nimi                          string
	AsutatudSissemaksetTegemata   bool
	EsitabKasusaajad              bool
	EsmaregistreerimiseKpv        int64
	EttevotteregistriNr           *string
	EvksRegistreeritud            *bool
	EvksRegistreeritudKandeKpv    *int64
	KustutamiseKpv                *int64
	LahknevusteadePuudumisest     *string
	LoobunudVorminouetest         *bool
	OnRaamatupidamiskohustuslane  bool
	OiguslikVorm                  string
	OiguslikVormNr                *int64
	OiguslikVormTekstina          string
	OiguslikuVormiAlaliik         *string
	OiguslikuVormiAlaliikTekstina string
	Piirkond                      *int64
	PiirkondTekstina              string
	PiirkondTekstinaPikk          string
	Staatus                       string
	StaatusTekstina               string
	Tegutseb                      *bool
	TegutsebTekstina              string
}

func (Ettevote) TableName() string {
	return "ettevotted"
}

type Aadress struct {
	EttevotteID                                      int64 `gorm:"primarykey"`
	AadressAdsAdobID                                 *string
	AadressAdsAdrID                                  *int64
	AadressAdsAdsNormaliseeritudTaisaadress          *string
	AadressAdsAdsNormaliseeritudTaisaadressTapsustus *string
	AadressAdsAdsOid                                 *string
	AadressAdsKoodaadress                            *string
	AadressAdsTyyp                                   *string
	AlgusKpv                                         int64
	Ehak                                             *string
	EhakNimetus                                      *string
	KaardiNr                                         int64
	KaardiPiirkond                                   int64
	KaardiTyyp                                       string
	KandeNr                                          int64
	KirjeID                                          int64
	LoppKpv                                          *int64
	Postiindeks                                      *string
	Riik                                             *string
	RiikTekstina                                     *string
	TanavMajaKorter                                  *string
}

func (Aadress) TableName() string {
	return "aadressid"
}

type Arinimi struct {
	EttevotteID    int64 `gorm:"primarykey"`
	AlgusKpv       *int64
	KaardiNr       *int64
	KaardiPiirkond *int64
	KaardiTyyp     *string
	KandeNr        *int64
	KirjeID        *int64
	LoppKpv        *int64
	Sisu           *string
}

func (Arinimi) TableName() string {
	return "arinimed"
}

type Kapital struct {
	EttevotteID             int64 `gorm:"primarykey"`
	AlgusKpv                *int64
	KaardiNr                *int64
	KaardiPiirkond          *int64
	KaardiTyyp              *string
	KandeNr                 *int64
	KapitaliSuurus          *string
	KapitaliValuuta         *string
	KapitaliValuutaTekstina *string
	KirjeID                 *int64
	LoppKpv                 *int64
}

func (Kapital) TableName() string {
	return "kapitalid"
}

type Majandusaasta struct {
	EttevotteID    int64 `gorm:"primarykey"`
	AlgusKpv       *int64
	KaardiNr       *int64
	KaardiPiirkond *int64
	KaardiTyyp     *string
	KandeNr        *int64
	KirjeID        *int64
	LoppKpv        *int64
	MajAastaAlgus  *string
	MajAastaLopp   *string
}

func (Majandusaasta) TableName() string {
	return "majandusaastad"
}

type MarkusedKaardil struct {
	EttevotteID    int64 
	AlgusKpv       *int64
	KaardiNr       *int64
	KaardiPiirkond *int64
	KaardiTyyp     *string
	KandeNr        *int64
	KirjeID        int64 `gorm:"primarykey"`
	LoppKpv        *int64
	Sisu           *string
	Tyyp           *string
	TyypTekstina   *string
	VeergNr        *int64
}

func (MarkusedKaardil) TableName() string {
	return "markused_kaardil"
}

type OiguslikVorm struct {
	EttevotteID    int64 
	AlgusKpv       *int64
	KaardiNr       *int64 
	KaardiPiirkond *int64
	KaardiTyyp     *string
	KandeNr        *int64
	KirjeID        int64 `gorm:"primarykey"`
	LoppKpv        *int64
	Sisu           *string
	SisuNr         *int64
	SisuTekstina   *string
}

func (OiguslikVorm) TableName() string {
	return "oiguslikud_vormid"
}

type Sidevahend struct {
	EttevotteID    int64 
	KaardiNr       *int64
	KaardiPiirkond *int64
	KaardiTyyp     *string
	KandeNr        *int64
	KirjeID        int64 `gorm:"primarykey"`
	Liik           *string
	LiikTekstina   *string
	LoppKpv        *int64
	Sisu           *string
}

func (Sidevahend) TableName() string {
	return "sidevahendid"
}

type Staatus struct {
	EttevotteID     int64 `gorm:"primarykey"`
	AlgusKpv        *int64 `gorm:"primarykey"`
	KaardiNr        *int64 
	KaardiPiirkond  *int64 
	KaardiTyyp      *string 
	KandeNr         *int64  
	Staatus         *string `gorm:"primarykey"`
	StaatusTekstina *string 
}

func (Staatus) TableName() string {
	return "staatused"
}

type TeatatudTegevusala struct {
	EttevotteID           int64 
	AlgusKpv              *int64
	EmtakKood             *string
	EmtakTekstina         *string
	EmtakVersioon         *int64
	EmtakVersioonTekstina *string
	KirjeID               int64 `gorm:"primarykey"`
	LoppKpv               *int64
	NaceKood              *string
	OnPohitegevusala      *bool
}

func (TeatatudTegevusala) TableName() string {
	return "teatatud_tegevusalad"
}

type Pohikiri struct {
	EttevotteID       int64 `gorm:"primarykey"`
	AlgusKpv          *int64
	KaardiNr          *int64
	KaardiPiirkond    *int64
	KaardiTyyp        *string
	KandeNr           *int64
	KinnitamiseKpv    *int64
	KirjeID           int64 
	LoppKpv           *int64
	MuutmiseKpv       *int64
	Selgitus          *string
	SisaldabErioigusi *bool
}

func (Pohikiri) TableName() string {
	return "pohikirjad"
}

func str(v interface{}) string {
	return v.(string)
}
func date(v interface{}) int64 {
	date, _ := time.Parse(format, v.(string))
	return date.Unix()
}
func boolean(v interface{}) bool {
	return v.(bool)
}
func integer(v interface{}) int64 {
	return int64(v.(float64))
}

func strPointer(v interface{}) *string {
	if v == nil {
		return nil
	}
	s, ok := v.(string)
	if !ok {
		return nil
	}
	return &s
}

func boolPointer(v interface{}) *bool {
	if v == nil {
		return nil
	}
	b, ok := v.(bool)
	if !ok {
		return nil
	}
	return &b
}

func integerPointer(v interface{}) *int64 {
	if v == nil {
		return nil
	}
	f, ok := v.(float64)
	if !ok {
		return nil
	}
	i := int64(f)
	return &i
}

const format = "02.01.2006"
const BATCH_SIZE = 1000

func insertBatch[T any](db *gorm.DB, items *[]T) {
	if len(*items) >= BATCH_SIZE {
		db.Create(items)
		*items = (*items)[:0]
	}
}
func insertAll[T any](db *gorm.DB, items *[]T) {
	if len(*items) > 0 {
		db.Create(items)
		*items = (*items)[:0]
	}
}

func datePointer(v interface{}) *int64 {
	if v == nil {
		return nil
	}
	date, err := time.Parse(format, v.(string))
	if err != nil {
		return nil
	}
	unix := date.Unix()
	return &unix
}
func parseAndAppend[T any](
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

func ParseYldandmed(src string) error {
	// Delete database
	os.Remove("test.db")

	// Open database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Ettevote{})
	db.AutoMigrate(&Aadress{})
	db.AutoMigrate(&Arinimi{})
	db.AutoMigrate(&Kapital{})
	db.AutoMigrate(&Majandusaasta{})
	db.AutoMigrate(&MarkusedKaardil{})
	db.AutoMigrate(&OiguslikVorm{})
	db.AutoMigrate(&Sidevahend{})
	db.AutoMigrate(&Staatus{})
	db.AutoMigrate(&TeatatudTegevusala{})
	db.AutoMigrate(&Pohikiri{})

	file, err := os.Open(src)
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
	ettevotted := make([]Ettevote, 0, BATCH_SIZE)
	aadressid := make([]Aadress, 0, BATCH_SIZE)
	arinimed := make([]Arinimi, 0, BATCH_SIZE)
	kapitalid := make([]Kapital, 0, BATCH_SIZE)
	majandusaastad := make([]Majandusaasta, 0, BATCH_SIZE)
	markusedKaardil := make([]MarkusedKaardil, 0, BATCH_SIZE)
	oiguslikudVormid := make([]OiguslikVorm, 0, BATCH_SIZE)
	sidevahendid := make([]Sidevahend, 0, BATCH_SIZE)
	staatused := make([]Staatus, 0, BATCH_SIZE)
	teatatudTegevusalad := make([]TeatatudTegevusala, 0, BATCH_SIZE)
	pohikirjad := make([]Pohikiri, 0, BATCH_SIZE)

	for decoder.More() {
		var value map[string]interface{}
		err := decoder.Decode(&value)
		if err != nil {
			return fmt.Errorf("error decoding JSON: %v", err)
		}

		id := integer(value["ariregistri_kood"])
		yldandmed := value["yldandmed"].(map[string]interface{})

		ettevotted = append(ettevotted, Ettevote{
			ID:                            id,
			Nimi:                          str(value["nimi"]),
			AsutatudSissemaksetTegemata:   boolean(yldandmed["asutatud_sissemakset_tegemata"]),
			EsitabKasusaajad:              boolean(yldandmed["esitab_kasusaajad"]),
			EsmaregistreerimiseKpv:        date(yldandmed["esmaregistreerimise_kpv"]),
			EttevotteregistriNr:           strPointer(yldandmed["ettevotteregistri_nr"]),
			EvksRegistreeritud:            boolPointer(yldandmed["evks_registreeritud"]),
			EvksRegistreeritudKandeKpv:    datePointer(yldandmed["evks_registreeritud_kande_kpv"]),
			KustutamiseKpv:                datePointer(yldandmed["kustutamise_kpv"]),
			LahknevusteadePuudumisest:     strPointer(yldandmed["lahknevusteade_puudumisest"]),
			LoobunudVorminouetest:         boolPointer(yldandmed["loobunud_vorminouetest"]),
			OnRaamatupidamiskohustuslane:  boolean(yldandmed["on_raamatupidamiskohustuslane"]),
			OiguslikVorm:                  str(yldandmed["oiguslik_vorm"]),
			OiguslikVormNr:                integerPointer(yldandmed["oiguslik_vorm_nr"]),
			OiguslikVormTekstina:          str(yldandmed["oiguslik_vorm_tekstina"]),
			OiguslikuVormiAlaliik:         strPointer(yldandmed["oigusliku_vormi_alaliik"]),
			OiguslikuVormiAlaliikTekstina: str(yldandmed["oigusliku_vormi_alaliik_tekstina"]),
			Piirkond:                      integerPointer(yldandmed["piirkond"]),
			PiirkondTekstina:              str(yldandmed["piirkond_tekstina"]),
			PiirkondTekstinaPikk:          str(yldandmed["piirkond_tekstina_pikk"]),
			Staatus:                       str(yldandmed["staatus"]),
			StaatusTekstina:               str(yldandmed["staatus_tekstina"]),
			Tegutseb:                      boolPointer(yldandmed["tegutseb"]),
			TegutsebTekstina:              str(yldandmed["tegutseb_tekstina"]),
		})

		parseAndAppend(id, yldandmed, "aadressid", &aadressid, func(id int64, aadress map[string]interface{}) Aadress {
			return Aadress{
				EttevotteID:                             id,
				AadressAdsAdobID:                        strPointer(aadress["aadress_ads__adob_id"]),
				AadressAdsAdrID:                         integerPointer(aadress["aadress_ads__adr_id"]),
				AadressAdsAdsNormaliseeritudTaisaadress: strPointer(aadress["aadress_ads__ads_normaliseeritud_taisaadress"]),
				AadressAdsAdsNormaliseeritudTaisaadressTapsustus: strPointer(aadress["aadress_ads__ads_normaliseeritud_taisaadress_tapsustus"]),
				AadressAdsAdsOid:      strPointer(aadress["aadress_ads__ads_oid"]),
				AadressAdsKoodaadress: strPointer(aadress["aadress_ads__koodaadress"]),
				AadressAdsTyyp:        strPointer(aadress["aadress_ads__tyyp"]),
				AlgusKpv:              date(aadress["algus_kpv"]),
				Ehak:                  strPointer(aadress["ehak"]),
				EhakNimetus:           strPointer(aadress["ehak_nimetus"]),
				KaardiNr:              integer(aadress["kaardi_nr"]),
				KaardiPiirkond:        integer(aadress["kaardi_piirkond"]),
				KaardiTyyp:            str(aadress["kaardi_tyyp"]),
				KandeNr:               integer(aadress["kande_nr"]),
				KirjeID:               integer(aadress["kirje_id"]),
				LoppKpv:               datePointer(aadress["lopp_kpv"]),
				Postiindeks:           strPointer(aadress["postiindeks"]),
				Riik:                  strPointer(aadress["riik"]),
				RiikTekstina:          strPointer(aadress["riik_tekstina"]),
				TanavMajaKorter:       strPointer(aadress["tanav_maja_korter"]),
			}
		})

		parseAndAppend(id, yldandmed, "arinimed", &arinimed, func(id int64, arinimi map[string]interface{}) Arinimi {
			return Arinimi{
				EttevotteID:    id,
				AlgusKpv:       datePointer(arinimi["algus_kpv"]),
				KaardiNr:       integerPointer(arinimi["kaardi_nr"]),
				KaardiPiirkond: integerPointer(arinimi["kaardi_piirkond"]),
				KaardiTyyp:     strPointer(arinimi["kaardi_tyyp"]),
				KandeNr:        integerPointer(arinimi["kande_nr"]),
				KirjeID:        integerPointer(arinimi["kirje_id"]),
				LoppKpv:        datePointer(arinimi["lopp_kpv"]),
				Sisu:           strPointer(arinimi["sisu"]),
			}
		})

		parseAndAppend(id, yldandmed, "kapitalid", &kapitalid, func(id int64, kapital map[string]interface{}) Kapital {
			return Kapital{
				EttevotteID:             id,
				AlgusKpv:                datePointer(kapital["algus_kpv"]),
				KaardiNr:                integerPointer(kapital["kaardi_nr"]),
				KaardiPiirkond:          integerPointer(kapital["kaardi_piirkond"]),
				KaardiTyyp:              strPointer(kapital["kaardi_tyyp"]),
				KandeNr:                 integerPointer(kapital["kande_nr"]),
				KapitaliSuurus:          strPointer(kapital["kapitali_suurus"]),
				KapitaliValuuta:         strPointer(kapital["kapitali_valuuta"]),
				KapitaliValuutaTekstina: strPointer(kapital["kapitali_valuuta_tekstina"]),
				KirjeID:                 integerPointer(kapital["kirje_id"]),
				LoppKpv:                 datePointer(kapital["lopp_kpv"]),
			}
		})

		parseAndAppend(id, yldandmed, "majandusaastad", &majandusaastad, func(id int64, majandusaasta map[string]interface{}) Majandusaasta {
			return Majandusaasta{
				EttevotteID:    id,
				AlgusKpv:       datePointer(majandusaasta["algus_kpv"]),
				KaardiNr:       integerPointer(majandusaasta["kaardi_nr"]),
				KaardiPiirkond: integerPointer(majandusaasta["kaardi_piirkond"]),
				KaardiTyyp:     strPointer(majandusaasta["kaardi_tyyp"]),
				KandeNr:        integerPointer(majandusaasta["kande_nr"]),
				MajAastaAlgus:  strPointer(majandusaasta["maj_aasta_algus"]),
				MajAastaLopp:   strPointer(majandusaasta["maj_aasta_lopp"]),
				KirjeID:        integerPointer(majandusaasta["kirje_id"]),
				LoppKpv:        datePointer(majandusaasta["lopp_kpv"]),
			}
		})

		parseAndAppend(id, yldandmed, "markused_kaardil", &markusedKaardil, func(id int64, markusedKaardil map[string]interface{}) MarkusedKaardil {
			return MarkusedKaardil{
				EttevotteID:    id,
				AlgusKpv:       datePointer(markusedKaardil["algus_kpv"]),
				KaardiNr:       integerPointer(markusedKaardil["kaardi_nr"]),
				KaardiPiirkond: integerPointer(markusedKaardil["kaardi_piirkond"]),
				KaardiTyyp:     strPointer(markusedKaardil["kaardi_tyyp"]),
				KandeNr:        integerPointer(markusedKaardil["kande_nr"]),
				KirjeID:        integer(markusedKaardil["kirje_id"]),
				LoppKpv:        datePointer(markusedKaardil["lopp_kpv"]),
				Sisu:           strPointer(markusedKaardil["sisu"]),
				Tyyp:           strPointer(markusedKaardil["tyyp"]),
				TyypTekstina:   strPointer(markusedKaardil["tyyp_tekstina"]),
				VeergNr:        integerPointer(markusedKaardil["veerg_nr"]),
			}
		})

		parseAndAppend(id, yldandmed, "oiguslikud_vormid", &oiguslikudVormid, func(id int64, oiguslikVorm map[string]interface{}) OiguslikVorm {
			return OiguslikVorm{
				EttevotteID:    id,
				AlgusKpv:       datePointer(oiguslikVorm["algus_kpv"]),
				KaardiNr:       integerPointer(oiguslikVorm["kaardi_nr"]),
				KaardiPiirkond: integerPointer(oiguslikVorm["kaardi_piirkond"]),
				KaardiTyyp:     strPointer(oiguslikVorm["kaardi_tyyp"]),
				KandeNr:        integerPointer(oiguslikVorm["kande_nr"]),
				KirjeID:        integer(oiguslikVorm["kirje_id"]),
				LoppKpv:        datePointer(oiguslikVorm["lopp_kpv"]),
				Sisu:           strPointer(oiguslikVorm["sisu"]),
				SisuNr:         integerPointer(oiguslikVorm["sisu_nr"]),
				SisuTekstina:   strPointer(oiguslikVorm["sisu_tekstina"]),
			}
		})

		parseAndAppend(id, yldandmed, "sidevahendid", &sidevahendid, func(id int64, sidevahend map[string]interface{}) Sidevahend {
			return Sidevahend{
				EttevotteID:    id,
				KaardiNr:       integerPointer(sidevahend["kaardi_nr"]),
				KaardiPiirkond: integerPointer(sidevahend["kaardi_piirkond"]),
				KaardiTyyp:     strPointer(sidevahend["kaardi_tyyp"]),
				KandeNr:        integerPointer(sidevahend["kande_nr"]),
				KirjeID:        integer(sidevahend["kirje_id"]),
				LoppKpv:        datePointer(sidevahend["lopp_kpv"]),
				Sisu:           strPointer(sidevahend["sisu"]),
				Liik:           strPointer(sidevahend["liik"]),
				LiikTekstina:   strPointer(sidevahend["liik_tekstina"]),
			}
		})

		parseAndAppend(id, yldandmed, "staatused", &staatused, func(id int64, staatus map[string]interface{}) Staatus {
			return Staatus{
				EttevotteID:     id,
				AlgusKpv:        datePointer(staatus["algus_kpv"]),
				KaardiNr:        integerPointer(staatus["kaardi_nr"]),
				KaardiPiirkond:  integerPointer(staatus["kaardi_piirkond"]),
				KaardiTyyp:      strPointer(staatus["kaardi_tyyp"]),
				KandeNr:         integerPointer(staatus["kande_nr"]),
				Staatus:         strPointer(staatus["staatus"]),
				StaatusTekstina: strPointer(staatus["staatus_tekstina"]),
			}
		})

		parseAndAppend(id, yldandmed, "teatatud_tegevusalad", &teatatudTegevusalad, func(id int64, teatatudTegevusala map[string]interface{}) TeatatudTegevusala {
			return TeatatudTegevusala{
				EttevotteID:           id,
				AlgusKpv:              datePointer(teatatudTegevusala["algus_kpv"]),
				KirjeID:               integer(teatatudTegevusala["kirje_id"]),
				LoppKpv:               datePointer(teatatudTegevusala["lopp_kpv"]),
				EmtakKood:             strPointer(teatatudTegevusala["emtak_kood"]),
				EmtakTekstina:         strPointer(teatatudTegevusala["emtak_tekstina"]),
				EmtakVersioon:         integerPointer(teatatudTegevusala["emtak_versioon"]),
				EmtakVersioonTekstina: strPointer(teatatudTegevusala["emtak_versioon_tekstina"]),
				NaceKood:              strPointer(teatatudTegevusala["nace_kood"]),
				OnPohitegevusala:      boolPointer(teatatudTegevusala["on_pohitegevusala"]),
			}
		})

		parseAndAppend(id, yldandmed, "pohikirjad", &pohikirjad, func(id int64, pohikiri map[string]interface{}) Pohikiri {
			return Pohikiri{
				EttevotteID:       id,
				AlgusKpv:          datePointer(pohikiri["algus_kpv"]),
				KaardiNr:          integerPointer(pohikiri["kaardi_nr"]),
				KaardiPiirkond:    integerPointer(pohikiri["kaardi_piirkond"]),
				KaardiTyyp:        strPointer(pohikiri["kaardi_tyyp"]),
				KandeNr:           integerPointer(pohikiri["kande_nr"]),
				KirjeID:           integer(pohikiri["kirje_id"]),
				LoppKpv:           datePointer(pohikiri["lopp_kpv"]),
				KinnitamiseKpv:    datePointer(pohikiri["kinnitamise_kpv"]),
				MuutmiseKpv:       datePointer(pohikiri["muutmise_kpv"]),
				Selgitus:          strPointer(pohikiri["selgitus"]),
				SisaldabErioigusi: boolPointer(pohikiri["sisaldab_erioigusi"]),
			}
		})

		if len(ettevotted) == BATCH_SIZE {
			i++
			fmt.Printf("Creating companies %d 000\n", i)
		}
		insertBatch(db, &ettevotted)
		insertBatch(db, &aadressid)
		insertBatch(db, &arinimed)
		insertBatch(db, &kapitalid)
		insertBatch(db, &majandusaastad)
		insertBatch(db, &markusedKaardil)
		insertBatch(db, &oiguslikudVormid)
		insertBatch(db, &sidevahendid)
		insertBatch(db, &staatused)
		insertBatch(db, &teatatudTegevusalad)
		insertBatch(db, &pohikirjad)
	}

	insertAll(db, &ettevotted)
	insertAll(db, &aadressid)
	insertAll(db, &arinimed)
	insertAll(db, &kapitalid)
	insertAll(db, &majandusaastad)
	insertAll(db, &markusedKaardil)
	insertAll(db, &oiguslikudVormid)
	insertAll(db, &sidevahendid)
	insertAll(db, &staatused)
	insertAll(db, &teatatudTegevusalad)
	insertAll(db, &pohikirjad)

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("error reading closing bracket: %v", err)
	}

	return nil
}
