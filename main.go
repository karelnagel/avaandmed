package main

import (
	"avaandmed/sources"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Args struct {
	SQLitePath    string
	BatchSize     int
	ForceDownload bool
	FailQuietly   bool
	Sources       []string
}

const DATA_DIR = "data"
const DEFAULT_SOURCES = "yldandmed,kaardile_kantud,kandevalised,kasusaajad,osanikud,majandusaasta,emta,debt,lihtandmed"

func main() {
	// Args
	total := time.Now()
	args := Args{}
	flag.StringVar(&args.SQLitePath, "sqlite", "out.db", "Path to the SQLite database")
	flag.IntVar(&args.BatchSize, "batch", 500, "Batch size")
	flag.BoolVar(&args.ForceDownload, "force-download", false, "Force downloading the latest data again, eg. it deletes the data directory")
	flag.BoolVar(&args.FailQuietly, "fail-quietly", false, "Fail quietly")
	srcs := flag.String("sources", DEFAULT_SOURCES, "Sources to process (comma separated)")

	flag.Parse()

	args.Sources = strings.Split(*srcs, ",")

	// Delete data directory
	if args.ForceDownload {
		if err := os.RemoveAll(DATA_DIR); err != nil {
			panic(fmt.Errorf("error removing data directory: %w", err))
		}
	}

	// Create data directory
	if err := os.MkdirAll(DATA_DIR, 0755); err != nil {
		panic(fmt.Errorf("error creating data directory: %w", err))
	}

	// Database
	os.Remove(args.SQLitePath)

	db, err := sources.InitDB(args.SQLitePath)
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err))
	}

	for _, source := range args.Sources {
		t := time.Now()
		fmt.Printf("Processing source %s\n", source)
		switch source {
		case "yldandmed":
			err = sources.ParseYldandmed(db, args.BatchSize)
		case "kaardile_kantud":
			err = sources.ParseKaardileKantud(db, args.BatchSize)
		case "kandevalised":
			err = sources.ParseKandevalised(db, args.BatchSize)
		case "kasusaajad":
			err = sources.ParseKasusaajad(db, args.BatchSize)
		case "osanikud":
			err = sources.ParseOsanikud(db, args.BatchSize)
		case "majandusaasta":
			err = sources.ParseMajandusaasta(db)
		case "emta":
			err = sources.ParseEMTA(db, args.BatchSize)
		case "debt":
			err = sources.ParseDebt(db, args.BatchSize)
		case "lihtandmed":
			err = sources.ParseLihtandmed(db, args.BatchSize)

		default:
			err = fmt.Errorf("unknown source: %s", source)
		}

		if err != nil {
			if !args.FailQuietly {
				panic(fmt.Errorf("source %s failed: %w", source, err))
			} else {
				fmt.Printf("Source %s failed: %s\n", source, err)
			}
		} else {
			fmt.Printf("Source %s finished in %s\n", source, time.Since(t))
		}
	}
	fmt.Printf("Total time: %s\n", time.Since(total))
}
