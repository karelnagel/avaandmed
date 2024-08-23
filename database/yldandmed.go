package database

type YldandmedFileJSON struct {
	AriregistriKood int64                     `json:"ariregistri_kood"`
	Nimi            string                    `json:"nimi"`
	Yldandmed       YldandmedWithChildrenJSON `json:"yldandmed"`
}

type YldandmedJSON struct {
	EttevotteregistriNr           *int    `json:"ettevotteregistri_nr"`
	EsmaregistreerimiseKpv        string  `json:"esmaregistreerimise_kpv"`
	KustutamiseKpv                *string `json:"kustutamise_kpv"`
	Staatus                       string  `json:"staatus"`
	StaatusTekstina               string  `json:"staatus_tekstina"`
	Piirkond                      int     `json:"piirkond"`
	PiirkondTekstina              string  `json:"piirkond_tekstina"`
	PiirkondTekstinaPikk          string  `json:"piirkond_tekstina_pikk"`
	EvksRegistreeritud            *bool   `json:"evks_registreeritud"`
	EvksRegistreeritudKandeKpv    *string `json:"evks_registreeritud_kande_kpv"`
	OiguslikVorm                  string  `json:"oiguslik_vorm"`
	OiguslikVormNr                int     `json:"oiguslik_vorm_nr"`
	OiguslikVormTekstina          string  `json:"oiguslik_vorm_tekstina"`
	OiguslikuVormiAlaliik         *string `json:"oigusliku_vormi_alaliik"`
	LahknevusteadePuudumisest     *bool   `json:"lahknevusteade_puudumisest"`
	OiguslikuVormiAlaliikTekstina string  `json:"oigusliku_vormi_alaliik_tekstina"`
	AsutatudSissemaksetTegemata   bool    `json:"asutatud_sissemakset_tegemata"`
	LoobunudVorminouetest         *bool   `json:"loobunud_vorminouetest"`
	OnRaamatupidamiskohustuslane  bool    `json:"on_raamatupidamiskohustuslane"`
	Tegutseb                      *string `json:"tegutseb"`
	TegutsebTekstina              string  `json:"tegutseb_tekstina"`
	EsitabKasusaajad              bool    `json:"esitab_kasusaajad"`
}
type YldandmedWithChildrenJSON struct {
	YldandmedJSON
	Staatused                      []StaatusJSON                    `json:"staatused"`
	Arinimed                       []ArinimiJSON                    `json:"arinimed"`
	Aadressid                      []AadressJSON                    `json:"aadressid"`
	OiguslikudVormid               []OiguslikVormJSON               `json:"oiguslikud_vormid"`
	Kapitalid                      []KapitalJSON                    `json:"kapitalid"`
	Majandusaastad                 []MajandusaastaJSON              `json:"majandusaastad"`
	Pohikirjad                     []PohikiriJSON                   `json:"pohikirjad"`
	MarkusedKaardil                []MarkusKaardilJSON              `json:"markused_kaardil"`
	Sidevahendid                   []SidevahendJSON                 `json:"sidevahendid"`
	TeatatudTegevusalad            []TeatatudTegevusalaJSON         `json:"teatatud_tegevusalad"`
	InfoMajandusaastaAruandestJSON []InfoMajandusaastaAruandestJSON `json:"info_majandusaasta_aruannetest"`
}
type Yldandmed struct {
	YldandmedJSON
	EttevotteID                   int64 `gorm:"primarykey"`
	Nimi                          string
	EsmaregistreerimiseKpvInt     int64
	KustutamiseKpvInt             *int64
	EvksRegistreeritudKandeKpvInt *int64
}

type StaatusJSON struct {
	KaardiPiirkond  int    `json:"kaardi_piirkond"`
	KaardiNr        int    `json:"kaardi_nr"`
	KaardiTyyp      string `json:"kaardi_tyyp"`
	KandeNr         int    `json:"kande_nr"`
	Staatus         string `json:"staatus"`
	StaatusTekstina string `json:"staatus_tekstina"`
	AlgusKpv        string `json:"algus_kpv"`
}
type Staatus struct {
	StaatusJSON
	ID          int `gorm:"primarykey"`
	EttevotteID int64
	AlgusKpvInt int64
}

type ArinimiJSON struct {
	KirjeID        int64   `json:"kirje_id"`
	KaardiPiirkond int     `json:"kaardi_piirkond"`
	KaardiNr       int     `json:"kaardi_nr"`
	KaardiTyyp     string  `json:"kaardi_tyyp"`
	KandeNr        int     `json:"kande_nr"`
	Sisu           string  `json:"sisu"`
	AlgusKpv       string  `json:"algus_kpv"`
	LoppKpv        *string `json:"lopp_kpv"`
}
type Arinimi struct {
	ArinimiJSON
	EttevotteID int64 `gorm:"primarykey"`
	AlgusKpvInt int64
	LoppKpvInt  *int64
}

type AadressJSON struct {
	KirjeID                                          int64   `json:"kirje_id"`
	KaardiPiirkond                                   int     `json:"kaardi_piirkond"`
	KaardiNr                                         int     `json:"kaardi_nr"`
	KaardiTyyp                                       string  `json:"kaardi_tyyp"`
	KandeNr                                          int     `json:"kande_nr"`
	Riik                                             string  `json:"riik"`
	RiikTekstina                                     string  `json:"riik_tekstina"`
	Ehak                                             string  `json:"ehak"`
	EhakNimetus                                      string  `json:"ehak_nimetus"`
	TanavMajaKorter                                  string  `json:"tanav_maja_korter"`
	AadressAdsAdsOid                                 string  `json:"aadress_ads__ads_oid"`
	AadressAdsAdrID                                  int     `json:"aadress_ads__adr_id"`
	AadressAdsAdsNormaliseeritudTaisaadress          string  `json:"aadress_ads__ads_normaliseeritud_taisaadress"`
	AadressAdsAdsNormaliseeritudTaisaadressTapsustus *string `json:"aadress_ads__ads_normaliseeritud_taisaadress_tapsustus"`
	AadressAdsKoodaadress                            string  `json:"aadress_ads__koodaadress"`
	AadressAdsAdobID                                 string  `json:"aadress_ads__adob_id"`
	AadressAdsTyyp                                   *string `json:"aadress_ads__tyyp"`
	Postiindeks                                      string  `json:"postiindeks"`
	AlgusKpv                                         string  `json:"algus_kpv"`
	LoppKpv                                          *string `json:"lopp_kpv"`
}
type Aadress struct {
	AadressJSON
	EttevotteID int64 `gorm:"primarykey"`
	AlgusKpvInt int64
	LoppKpvInt  *int64
}

type OiguslikVormJSON struct {
	KirjeID        int64   `json:"kirje_id"`
	KaardiPiirkond int     `json:"kaardi_piirkond"`
	KaardiNr       int     `json:"kaardi_nr"`
	KaardiTyyp     string  `json:"kaardi_tyyp"`
	KandeNr        int     `json:"kande_nr"`
	Sisu           string  `json:"sisu"`
	SisuNr         int     `json:"sisu_nr"`
	SisuTekstina   string  `json:"sisu_tekstina"`
	AlgusKpv       string  `json:"algus_kpv"`
	LoppKpv        *string `json:"lopp_kpv"`
}
type OiguslikVorm struct {
	OiguslikVormJSON
	ID          int `gorm:"primarykey"`
	EttevotteID int64
	AlgusKpvInt int64
	LoppKpvInt  *int64
}

type KapitalJSON struct {
	KirjeID                 int64   `json:"kirje_id"`
	KaardiPiirkond          int     `json:"kaardi_piirkond"`
	KaardiNr                int     `json:"kaardi_nr"`
	KaardiTyyp              string  `json:"kaardi_tyyp"`
	KandeNr                 int     `json:"kande_nr"`
	KapitaliSuurus          string  `json:"kapitali_suurus"`
	KapitaliValuuta         string  `json:"kapitali_valuuta"`
	KapitaliValuutaTekstina string  `json:"kapitali_valuuta_tekstina"`
	AlgusKpv                string  `json:"algus_kpv"`
	LoppKpv                 *string `json:"lopp_kpv"`
}
type Kapital struct {
	KapitalJSON
	EttevotteID int64 `gorm:"primarykey"`
	AlgusKpvInt int64
	LoppKpvInt  *int64
}

type MajandusaastaJSON struct {
	KirjeID        int64   `json:"kirje_id"`
	KaardiPiirkond int     `json:"kaardi_piirkond"`
	KaardiNr       int     `json:"kaardi_nr"`
	KaardiTyyp     string  `json:"kaardi_tyyp"`
	KandeNr        int     `json:"kande_nr"`
	MajAastaAlgus  string  `json:"maj_aasta_algus"`
	MajAastaLopp   string  `json:"maj_aasta_lopp"`
	AlgusKpv       string  `json:"algus_kpv"`
	LoppKpv        *string `json:"lopp_kpv"`
}
type Majandusaasta struct {
	MajandusaastaJSON
	EttevotteID int64 `gorm:"primarykey"`
	AlgusKpvInt int64
	LoppKpvInt  *int64
}

type PohikiriJSON struct {
	KirjeID           int64   `json:"kirje_id"`
	KaardiPiirkond    int     `json:"kaardi_piirkond"`
	KaardiNr          int     `json:"kaardi_nr"`
	KaardiTyyp        string  `json:"kaardi_tyyp"`
	KandeNr           int     `json:"kande_nr"`
	KinnitamiseKpv    *string `json:"kinnitamise_kpv"`
	MuutmiseKpv       *string `json:"muutmise_kpv"`
	Selgitus          *string `json:"selgitus"`
	AlgusKpv          string  `json:"algus_kpv"`
	LoppKpv           *string `json:"lopp_kpv"`
	SisaldabErioigusi bool    `json:"sisaldab_erioigusi"`
}
type Pohikiri struct {
	PohikiriJSON
	EttevotteID int64 `gorm:"primarykey"`
	AlgusKpvInt int64
	LoppKpvInt  *int64
}

type MarkusKaardilJSON struct {
	KirjeID        int64   `json:"kirje_id"`
	KaardiPiirkond int     `json:"kaardi_piirkond"`
	KaardiNr       int     `json:"kaardi_nr"`
	KaardiTyyp     string  `json:"kaardi_tyyp"`
	KandeNr        int     `json:"kande_nr"`
	VeergNr        int     `json:"veerg_nr"`
	Tyyp           string  `json:"tyyp"`
	TyypTekstina   string  `json:"tyyp_tekstina"`
	Sisu           string  `json:"sisu"`
	AlgusKpv       string  `json:"algus_kpv"`
	LoppKpv        *string `json:"lopp_kpv"`
}
type MarkusKaardil struct {
	MarkusKaardilJSON
	ID          int `gorm:"primarykey"`
	EttevotteID int64
	AlgusKpvInt int64
	LoppKpvInt  *int64
}

type SidevahendJSON struct {
	KirjeID        int64   `json:"kirje_id"`
	Liik           string  `json:"liik"`
	LiikTekstina   string  `json:"liik_tekstina"`
	Sisu           string  `json:"sisu"`
	LoppKpv        *string `json:"lopp_kpv"`
	KaardiPiirkond int     `json:"kaardi_piirkond"`
	KaardiNr       int     `json:"kaardi_nr"`
	KaardiTyyp     string  `json:"kaardi_tyyp"`
	KandeNr        int     `json:"kande_nr"`
}
type Sidevahend struct {
	SidevahendJSON
	ID          int `gorm:"primarykey"`
	EttevotteID int64
	LoppKpvInt  *int64
}

type TeatatudTegevusalaJSON struct {
	KirjeID               int64   `json:"kirje_id"`
	EmtakKood             string  `json:"emtak_kood"`
	EmtakTekstina         string  `json:"emtak_tekstina"`
	EmtakVersioon         int     `json:"emtak_versioon"`
	EmtakVersioonTekstina *string `json:"emtak_versioon_tekstina"`
	NaceKood              string  `json:"nace_kood"`
	OnPohitegevusala      bool    `json:"on_pohitegevusala"`
	AlgusKpv              string  `json:"algus_kpv"`
	LoppKpv               *string `json:"lopp_kpv"`
}
type TeatatudTegevusala struct {
	TeatatudTegevusalaJSON
	ID          int `gorm:"primarykey"`
	EttevotteID int64
	AlgusKpvInt int64
	LoppKpvInt  *int64
}

type InfoMajandusaastaAruandestJSON struct {
	KirjeID                         int64   `json:"kirje_id"`
	MajandusaastaPeriodiAlgusKpv    string  `json:"majandusaasta_perioodi_algus_kpv"`
	MajandusaastaPeriodiLoppKpv     string  `json:"majandusaasta_perioodi_lopp_kpv"`
	TootajateArv                    string  `json:"tootajate_arv"`
	EttevotjaAadressAruandes        string  `json:"ettevotja_aadress_aruandes"`
	TegevusalaEmtakKood             *string `json:"tegevusala_emtak_kood"`
	TegevusalaEmtakTekstina         *string `json:"tegevusala_emtak_tekstina"`
	TegevusalaEmtakVersioon         *string `json:"tegevusala_emtak_versioon"`
	TegevusalaEmtakVersioonTekstina *string `json:"tegevusala_emtak_versioon_tekstina"`
	TegevusalaNaceKood              *string `json:"tegevusala_nace_kood"`
}
type InfoMajandusaastaAruandest struct {
	InfoMajandusaastaAruandestJSON
	ID                              int `gorm:"primarykey"`
	EttevotteID                     int64
	MajandusaastaPeriodiAlgusKpvInt int64
	MajandusaastaPeriodiLoppKpvInt  *int64
}
