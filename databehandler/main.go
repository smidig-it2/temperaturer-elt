// Package main leser timebaserte temperaturer fra CSV og beregner daglige gjennomsnitt.
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

const (
	dataDir    = "../data"
	mainCSV    = "data.csv"
	dummyCSV   = "dummy.csv"
	outputJSON = "data.json"
)

// DagTemp representerer gjennomsnittstemperatur for én dag.
type DagTemp struct {
	Dato         string  `json:"dato"`
	Gjennomsnitt float64 `json:"gjennomsnitt"`
}

// main leser CSV-fil (data.csv eller dummy.csv), beregner daglige
// gjennomsnittstemperaturer og skriver resultat til data.json.
func main() {
	inputPath := filepath.Join(dataDir, mainCSV)
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		inputPath = filepath.Join(dataDir, dummyCSV)
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			log.Fatalf("Fant verken %s eller %s i %s", mainCSV, dummyCSV, dataDir)
		}
	}
	outputPath := filepath.Join(dataDir, outputJSON)

	dager, err := lesOgAggreger(inputPath)
	if err != nil {
		log.Fatalf("Feil: %v", err)
	}

	if err := skrivJSON(outputPath, dager); err != nil {
		log.Fatalf("Feil: %v", err)
	}

	fmt.Printf("Skrev %d dager til %s\n", len(dager), outputPath)
}

// lesOgAggreger leser en CSV-fil med tidsstempel (RFC3339) og temperatur,
// og returnerer daglige gjennomsnittstemperaturer sortert etter dato.
func lesOgAggreger(csvPath string) ([]DagTemp, error) {
	f, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) <= 1 {
		return nil, fmt.Errorf("ingen data rader i CSV")
	}

	// Akkumulator for sum og antall målinger per dato
	type akk struct {
		sum    float64
		antall int
	}
	perDag := make(map[string]*akk)

	for _, row := range records[1:] {
		if len(row) < 2 {
			continue
		}

		t, err := time.Parse(time.RFC3339, row[0])
		if err != nil {
			continue
		}
		dato := t.Format("2006-01-02")

		temp, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			continue
		}

		a := perDag[dato]
		if a == nil {
			a = &akk{}
			perDag[dato] = a
		}
		a.sum += temp
		a.antall++
	}

	res := make([]DagTemp, 0, len(perDag))
	for dato, a := range perDag {
		res = append(res, DagTemp{
			Dato:         dato,
			Gjennomsnitt: a.sum / float64(a.antall),
		})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Dato < res[j].Dato
	})

	return res, nil
}

// skrivJSON serialiserer dagdata til JSON og skriver til fil.
func skrivJSON(path string, dager []DagTemp) error {
	data, err := json.MarshalIndent(dager, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// Smidig IT-2 © TIP AS, 2025