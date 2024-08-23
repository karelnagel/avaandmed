package database

type Tabler interface {
	TableName() string
}

// Yldandmed
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
	EttevotteID     int64  `gorm:"primarykey"`
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

// Kaardile Kantud
