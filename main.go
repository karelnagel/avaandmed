package main

import (
	"avaandmed/database"
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
	DeleteDataDir bool
	Sources       []string
}

const DATA_DIR = "data"
const DEFAULT_SOURCES = "yldandmed,kaardile_kantud,kandevalised,kasusaajad"

func main() {
	// Args
	args := Args{}
	flag.StringVar(&args.SQLitePath, "sqlite", "data/out.db", "Path to the SQLite database")
	flag.IntVar(&args.BatchSize, "batch", 800, "Batch size")
	flag.BoolVar(&args.DeleteDataDir, "delete", false, "Delete data directory")
	srcs := flag.String("sources", DEFAULT_SOURCES, "Sources to process (comma separated), default: "+DEFAULT_SOURCES)

	flag.Parse()

	args.Sources = strings.Split(*srcs, ",")

	// Delete data directory
	if args.DeleteDataDir {
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

	db, err := database.InitDB(args.SQLitePath)
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err))
	}

	for _, source := range args.Sources {
		t := time.Now()
		fmt.Printf("Processing source %s\n", source)
		switch source {
		case "yldandmed":
			err = sources.Yldandmed(db, args.BatchSize)
		case "kaardile_kantud":
			err = sources.KaardileKantud(db, args.BatchSize)
		case "kandevalised":
			err = sources.Kandevalised(db, args.BatchSize)
		case "kasusaajad":
			err = sources.Kasusaajad(db, args.BatchSize)
		default:
			err = fmt.Errorf("unknown source: %s", source)
		}

		if err != nil {
			panic(fmt.Errorf("source %s failed: %w", source, err))
		} else {
			fmt.Printf("Source %s finished in %s\n", source, time.Since(t))
		}
	}
}
