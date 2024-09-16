package sources

import (
	"avaandmed/utils"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"os"
)

func ParseYldandmed(db *gorm.DB, batchSize int) error {
	source := utils.Source{
		URL:      "https://avaandmed.ariregister.rik.ee/sites/default/files/avaandmed/ettevotja_rekvisiidid__yldandmed.json.zip",
		ZipPath:  "data/ettevotja_rekvisiidid__yldandmed.json.zip",
		FilePath: "data/ettevotja_rekvisiidid__yldandmed.json",
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

	ettevotted := make([]Ettevote, 0, batchSize)
	yldandmed := make([]Yldandmed, 0, batchSize)
	aadressid := make([]Aadress, 0, batchSize)
	arinimed := make([]Arinimi, 0, batchSize)
	kapitalid := make([]Kapital, 0, batchSize)
	majandusaastad := make([]YldandmedMajandusaasta, 0, batchSize)
	markusedKaardil := make([]MarkusKaardil, 0, batchSize)
	oiguslikudVormid := make([]OiguslikVorm, 0, batchSize)
	sidevahendid := make([]Sidevahend, 0, batchSize)
	staatused := make([]Staatus, 0, batchSize)
	teatatudTegevusalad := make([]TeatatudTegevusala, 0, batchSize)
	pohikirjad := make([]Pohikiri, 0, batchSize)
	infoMajandusaastaAruannetest := make([]InfoMajandusaastaAruandest, 0, batchSize)

	bar := utils.NewProgressBar(utils.COMPANIES, "Processing Yldandmed")
	for decoder.More() {
		bar.Add(1)
		var value YldandmedFileJSON
		err := decoder.Decode(&value)
		if err != nil {
			return fmt.Errorf("error decoding JSON: %v", err)
		}
		ettevotted = append(ettevotted, Ettevote{
			ID:   value.AriregistriKood,
			Name: value.Nimi,
		})

		yldandmed = append(yldandmed, Yldandmed{
			YldandmedJSON:                 value.Yldandmed.YldandmedJSON,
			EttevotteID:                   value.AriregistriKood,
			Nimi:                          value.Nimi,
			EsmaregistreerimiseKpvInt:     utils.Date(value.Yldandmed.EsmaregistreerimiseKpv),
			KustutamiseKpvInt:             utils.DatePointer(value.Yldandmed.KustutamiseKpv),
			EvksRegistreeritudKandeKpvInt: utils.DatePointer(value.Yldandmed.EvksRegistreeritudKandeKpv),
		})

		for _, aadress := range value.Yldandmed.Aadressid {
			aadressid = append(aadressid, Aadress{
				AadressJSON: aadress,
				EttevotteID: value.AriregistriKood,
				AlgusKpvInt: utils.Date(aadress.AlgusKpv),
				LoppKpvInt:  utils.DatePointer(aadress.LoppKpv),
			})
		}

		for _, arinimi := range value.Yldandmed.Arinimed {
			arinimed = append(arinimed, Arinimi{
				ArinimiJSON: arinimi,
				EttevotteID: value.AriregistriKood,
				AlgusKpvInt: utils.Date(arinimi.AlgusKpv),
				LoppKpvInt:  utils.DatePointer(arinimi.LoppKpv),
			})
		}

		for _, kapital := range value.Yldandmed.Kapitalid {
			kapitalid = append(kapitalid, Kapital{
				KapitalJSON: kapital,
				EttevotteID: value.AriregistriKood,
				AlgusKpvInt: utils.Date(kapital.AlgusKpv),
				LoppKpvInt:  utils.DatePointer(kapital.LoppKpv),
			})
		}

		for _, majandusaasta := range value.Yldandmed.YldandmedMajandusaastad {
			majandusaastad = append(majandusaastad, YldandmedMajandusaasta{
				YldandmedMajandusaastaJSON: majandusaasta,
				EttevotteID:                value.AriregistriKood,
				AlgusKpvInt:                utils.Date(majandusaasta.AlgusKpv),
				LoppKpvInt:                 utils.DatePointer(majandusaasta.LoppKpv),
			})
		}

		for _, pohikiri := range value.Yldandmed.Pohikirjad {
			pohikirjad = append(pohikirjad, Pohikiri{
				PohikiriJSON: pohikiri,
				EttevotteID:  value.AriregistriKood,
				AlgusKpvInt:  utils.Date(pohikiri.AlgusKpv),
				LoppKpvInt:   utils.DatePointer(pohikiri.LoppKpv),
			})
		}

		for _, staatus := range value.Yldandmed.Staatused {
			staatused = append(staatused, Staatus{
				StaatusJSON: staatus,
				EttevotteID: value.AriregistriKood,
				AlgusKpvInt: utils.Date(staatus.AlgusKpv),
			})
		}

		for _, teatatudTegevusala := range value.Yldandmed.TeatatudTegevusalad {
			teatatudTegevusalad = append(teatatudTegevusalad, TeatatudTegevusala{
				TeatatudTegevusalaJSON: teatatudTegevusala,
				EttevotteID:            value.AriregistriKood,
				AlgusKpvInt:            utils.Date(teatatudTegevusala.AlgusKpv),
				LoppKpvInt:             utils.DatePointer(teatatudTegevusala.LoppKpv),
			})
		}

		for _, info := range value.Yldandmed.InfoMajandusaastaAruandestJSON {
			infoMajandusaastaAruannetest = append(infoMajandusaastaAruannetest, InfoMajandusaastaAruandest{
				InfoMajandusaastaAruandestJSON:  info,
				EttevotteID:                     value.AriregistriKood,
				MajandusaastaPeriodiAlgusKpvInt: utils.Date(info.MajandusaastaPeriodiAlgusKpv),
				MajandusaastaPeriodiLoppKpvInt:  utils.DatePointer(&info.MajandusaastaPeriodiLoppKpv),
			})
		}

		for _, x := range value.Yldandmed.MarkusedKaardil {
			markusedKaardil = append(markusedKaardil, MarkusKaardil{
				MarkusKaardilJSON: x,
				EttevotteID:       value.AriregistriKood,
				AlgusKpvInt:       utils.Date(x.AlgusKpv),
				LoppKpvInt:        utils.DatePointer(x.LoppKpv),
			})

		}
		for _, x := range value.Yldandmed.OiguslikudVormid {
			oiguslikudVormid = append(oiguslikudVormid, OiguslikVorm{
				OiguslikVormJSON: x,
				EttevotteID:      value.AriregistriKood,
				AlgusKpvInt:      utils.Date(x.AlgusKpv),
				LoppKpvInt:       utils.DatePointer(x.LoppKpv),
			})
		}
		for _, x := range value.Yldandmed.Sidevahendid {
			sidevahendid = append(sidevahendid, Sidevahend{
				SidevahendJSON: x,
				EttevotteID:    value.AriregistriKood,
				LoppKpvInt:     utils.DatePointer(x.LoppKpv),
			})
		}

		InsertBatch(db, &ettevotted, batchSize)
		InsertBatch(db, &yldandmed, batchSize)
		InsertBatch(db, &aadressid, batchSize)
		InsertBatch(db, &arinimed, batchSize)
		InsertBatch(db, &kapitalid, batchSize)
		InsertBatch(db, &majandusaastad, batchSize)
		InsertBatch(db, &markusedKaardil, batchSize)
		InsertBatch(db, &oiguslikudVormid, batchSize)
		InsertBatch(db, &sidevahendid, batchSize)
		InsertBatch(db, &staatused, batchSize)
		InsertBatch(db, &teatatudTegevusalad, batchSize)
		InsertBatch(db, &pohikirjad, batchSize)
		InsertBatch(db, &infoMajandusaastaAruannetest, batchSize)
	}

	InsertAll(db, &ettevotted)
	InsertAll(db, &yldandmed)
	InsertAll(db, &aadressid)
	InsertAll(db, &arinimed)
	InsertAll(db, &kapitalid)
	InsertAll(db, &majandusaastad)
	InsertAll(db, &markusedKaardil)
	InsertAll(db, &oiguslikudVormid)
	InsertAll(db, &sidevahendid)
	InsertAll(db, &staatused)
	InsertAll(db, &teatatudTegevusalad)
	InsertAll(db, &pohikirjad)
	InsertAll(db, &infoMajandusaastaAruannetest)

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("error reading closing bracket: %v", err)
	}

	return nil
}
