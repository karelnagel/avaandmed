package database

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
