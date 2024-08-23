package sources

import (
	"avaandmed/utils"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"gorm.io/gorm"
)

type Majandusaasta struct {
	ReportID                        string `gorm:"primaryKey"`
	TaidetudAruanneReportID         string
	Registrikood                    string
	OiguslikVorm                    string
	Staatus                         string
	Aruandeaasta                    string
	KasKonsolideeritud              string
	PeriodStart                     string
	PeriodEnd                       string
	EsitatudKpv                     string
	KasAuditeeritud                 string
	ValitudAruanneKategooria        string
	MinimaalneKategooriaAndmetest   string
	AuditiToovottuLiik              string
	AudiitoriOtsuseTyyp             string
	ModifikatsioonAsjaoluRohutamine string
	ModifikatsioonMuuAsjaolu        string
	ModifikatsioonTegevuseJatkuvus  string
	EsitatudKpvInt                  *int64

	Assets                                                        *int
	AverageNumberOfEmployeesInFullTimeEquivalentUnits             *int
	CashAndCashEquivalents                                        *int
	CurrentAssets                                                 *int
	CurrentLiabilities                                            *int
	DepreciationAndImpairmentLossReversal                         *int
	EmployeeExpense                                               *int
	Equity                                                        *int
	IssuedCapital                                                 *int
	LaborExpense                                                  *int
	NonCurrentAssets                                              *int
	NonCurrentLiabilities                                         *int
	RetainedEarningsLoss                                          *int
	Revenue                                                       *int
	TotalAnnualPeriodProfitLoss                                   *int
	TotalProfitLoss                                               *int
	NetAssets                                                     *int
	SurplusDeficitFromOperatingActivities                         *int
	NetSurplusDeficitForPeriod                                    *int
	TotalRevenue                                                  *int
	DepreciationAndImpairmentLossReversalNeg                      *int
	IssuedCapital2                                                *int
	RetainedEarningsLossConsolidated                              *int
	CurrentAssetsConsolidated                                     *int
	CurrentLiabilitiesConsolidated                                *int
	EquityConsolidated                                            *int
	IssuedCapitalConsolidated                                     *int
	NonCurrentLiabilitiesConsolidated                             *int
	NonCurrentAssetsConsolidated                                  *int
	CashAndCashEquivalentsConsolidated                            *int
	AssetsConsolidated                                            *int
	TotalAnnualPeriodProfitLossConsolidated                       *int
	RevenueConsolidated                                           *int
	DepreciationAndImpairmentLossReversalConsolidated             *int
	EmployeeExpenseConsolidated                                   *int
	TotalProfitLossConsolidated                                   *int
	DepreciationAndImpairmentLossReversalNegConsolidated          *int
	LaborExpenseConsolidated                                      *int
	AverageNumberOfEmployeesInFullTimeEquivalentUnitsConsolidated *int
	IssuedCapital2Consolidated                                    *int
}

func ParseMajandusaasta(db *gorm.DB) error {
	const batchSize = 400
	yldandmed := utils.Source{
		URL:      `https://avaandmed.ariregister.rik.ee/sites/default/files/1.aruannete_yldandmed_kuni_31072024.zip`,
		ZipPath:  "data/1.aruannete_yldandmed_kuni_31072024.zip",
		FilePath: "data/1.aruannete_yldandmed_kuni_31072024.csv",
	}
	kuni2019 := utils.Source{
		URL:      `https://avaandmed.ariregister.rik.ee/sites/default/files/4.2019_aruannete_elemendid_kuni_31072024.zip`,
		ZipPath:  "data/4.2019_aruannete_elemendid_kuni_31072024.zip",
		FilePath: "data/4.2019_aruannete_elemendid_kuni_31072024.csv",
	}
	kuni2020 := utils.Source{
		URL:      `https://avaandmed.ariregister.rik.ee/sites/default/files/4.2020_aruannete_elemendid_kuni_31072024.zip`,
		ZipPath:  "data/4.2020_aruannete_elemendid_kuni_31072024.zip",
		FilePath: "data/4.2020_aruannete_elemendid_kuni_31072024.csv",
	}
	kuni2021 := utils.Source{
		URL:      `https://avaandmed.ariregister.rik.ee/sites/default/files/4.2021_aruannete_elemendid_kuni_31072024.zip`,
		ZipPath:  "data/4.2021_aruannete_elemendid_kuni_31072024.zip",
		FilePath: "data/4.2021_aruannete_elemendid_kuni_31072024.csv",
	}
	kuni2022 := utils.Source{
		URL:      `https://avaandmed.ariregister.rik.ee/sites/default/files/4.2022_aruannete_elemendid_kuni_31072024.zip`,
		ZipPath:  "data/4.2022_aruannete_elemendid_kuni_31072024.zip",
		FilePath: "data/4.2022_aruannete_elemendid_kuni_31072024.csv",
	}
	kuni2023 := utils.Source{
		URL:      `https://avaandmed.ariregister.rik.ee/sites/default/files/4.2023_aruannete_elemendid_kuni_31072024.zip`,
		ZipPath:  "data/4.2023_aruannete_elemendid_kuni_31072024.zip",
		FilePath: "data/4.2023_aruannete_elemendid_kuni_31072024.csv",
	}
	years := []utils.Source{kuni2019, kuni2020, kuni2021, kuni2022, kuni2023}

	// Downloading
	err := yldandmed.Download()
	if err != nil {
		return fmt.Errorf("error downloading: %w", err)
	}
	err = kuni2019.Download()
	if err != nil {
		return fmt.Errorf("error downloading: %w", err)
	}
	err = kuni2020.Download()
	if err != nil {
		return fmt.Errorf("error downloading: %w", err)
	}
	err = kuni2021.Download()
	if err != nil {
		return fmt.Errorf("error downloading: %w", err)
	}
	err = kuni2022.Download()
	if err != nil {
		return fmt.Errorf("error downloading: %w", err)
	}
	err = kuni2023.Download()
	if err != nil {
		return fmt.Errorf("error downloading: %w", err)
	}

	file, _ := os.Open(yldandmed.FilePath)
	defer file.Close()

	reader := csv.NewReader(file)

	reader.LazyQuotes = true
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	majandusaastad := make(map[string]Majandusaasta)

	bar := utils.NewProgressBar(16391306, "Processing Majandusaasta")
	isFirst := true
	for {
		bar.Add(1)
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if isFirst {
			isFirst = false
			continue
		}
		if err != nil {
			return fmt.Errorf("error reading CSV record: %v", err)
		}
		majandusaastad[record[0]] = Majandusaasta{
			ReportID:                        record[0],
			TaidetudAruanneReportID:         record[1],
			Registrikood:                    record[2],
			OiguslikVorm:                    record[3],
			Staatus:                         record[4],
			Aruandeaasta:                    record[5],
			KasKonsolideeritud:              record[6],
			PeriodStart:                     record[7],
			PeriodEnd:                       record[8],
			EsitatudKpv:                     record[9],
			KasAuditeeritud:                 record[10],
			ValitudAruanneKategooria:        record[11],
			MinimaalneKategooriaAndmetest:   record[12],
			AuditiToovottuLiik:              record[13],
			AudiitoriOtsuseTyyp:             record[14],
			ModifikatsioonAsjaoluRohutamine: record[15],
			ModifikatsioonMuuAsjaolu:        record[16],
			ModifikatsioonTegevuseJatkuvus:  record[17],
			EsitatudKpvInt:                  utils.DatePointer(&record[9]),
		}
	}

	for _, year := range years {
		file, _ := os.Open(year.FilePath)
		defer file.Close()

		reader := csv.NewReader(file)

		reader.LazyQuotes = true
		reader.Comma = ';'
		reader.FieldsPerRecord = -1

		isFirst := true
		for {
			bar.Add(1)
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if isFirst {
				isFirst = false
				continue
			}
			if err != nil {
				return fmt.Errorf("error reading CSV record: %v", err)
			}

			kandeId := record[0]

			label := record[2]
			key := record[3]
			if key == "elemendi_nimetus" {
				continue
			}

			if key == "" {
				switch strings.ToLower(strings.TrimSpace(label)) {
				case "töötajate keskmine arv taandatuna täistööajale":
					key = "AverageNumberOfEmployeesInFullTimeEquivalentUnits"
				case "töötajate keskmine arv taandatuna täistööajale konsolideeritud":
					key = "AverageNumberOfEmployeesInFullTimeEquivalentUnitsConsolidated"
				case "müügitulu":
					key = "Revenue"
				case "varad":
					key = "Assets"
				case "varad konsolideeritud":
					key = "AssetsConsolidated"
				case "müügitulu konsolideeritud":
					key = "RevenueConsolidated"
				default:
					fmt.Println("Unknown label:", label)
				}
			}
			valueFloat, err := strconv.ParseFloat(record[4], 64)
			if err != nil {
				fmt.Printf("Error converting value to float: %v\n", err)
				continue
			}
			value := int(valueFloat)

			if majandusaasta, ok := majandusaastad[kandeId]; ok {
				switch key {
				case "Assets":
					majandusaasta.Assets = &value
				case "AverageNumberOfEmployeesInFullTimeEquivalentUnits":
					majandusaasta.AverageNumberOfEmployeesInFullTimeEquivalentUnits = &value
				case "CashAndCashEquivalents":
					majandusaasta.CashAndCashEquivalents = &value
				case "CurrentAssets":
					majandusaasta.CurrentAssets = &value
				case "CurrentLiabilities":
					majandusaasta.CurrentLiabilities = &value
				case "DepreciationAndImpairmentLossReversal":
					majandusaasta.DepreciationAndImpairmentLossReversal = &value
				case "EmployeeExpense":
					majandusaasta.EmployeeExpense = &value
				case "Equity":
					majandusaasta.Equity = &value
				case "RetainedEarningsLoss":
					majandusaasta.RetainedEarningsLoss = &value
				case "IssuedCapital":
					majandusaasta.IssuedCapital = &value
				case "LaborExpense":
					majandusaasta.LaborExpense = &value
				case "NonCurrentAssets":
					majandusaasta.NonCurrentAssets = &value
				case "NonCurrentLiabilities":
					majandusaasta.NonCurrentLiabilities = &value
				case "Revenue":
					majandusaasta.Revenue = &value
				case "TotalAnnualPeriodProfitLoss":
					majandusaasta.TotalAnnualPeriodProfitLoss = &value
				case "TotalProfitLoss":
					majandusaasta.TotalProfitLoss = &value
				case "NetAssets":
					majandusaasta.NetAssets = &value
				case "SurplusDeficitFromOperatingActivities":
					majandusaasta.SurplusDeficitFromOperatingActivities = &value
				case "NetSurplusDeficitForPeriod":
					majandusaasta.NetSurplusDeficitForPeriod = &value
				case "DepreciationAndImpairmentLossReversalNeg":
					majandusaasta.DepreciationAndImpairmentLossReversalNeg = &value
				case "IssuedCapital2":
					majandusaasta.IssuedCapital2 = &value
				case "RetainedEarningsLossConsolidated":
					majandusaasta.RetainedEarningsLossConsolidated = &value
				case "CurrentAssetsConsolidated":
					majandusaasta.CurrentAssetsConsolidated = &value
				case "CurrentLiabilitiesConsolidated":
					majandusaasta.CurrentLiabilitiesConsolidated = &value
				case "EquityConsolidated":
					majandusaasta.EquityConsolidated = &value
				case "IssuedCapitalConsolidated":
					majandusaasta.IssuedCapitalConsolidated = &value
				case "NonCurrentLiabilitiesConsolidated":
					majandusaasta.NonCurrentLiabilitiesConsolidated = &value
				case "NonCurrentAssetsConsolidated":
					majandusaasta.NonCurrentAssetsConsolidated = &value
				case "CashAndCashEquivalentsConsolidated":
					majandusaasta.CashAndCashEquivalentsConsolidated = &value
				case "AssetsConsolidated":
					majandusaasta.AssetsConsolidated = &value
				case "TotalAnnualPeriodProfitLossConsolidated":
					majandusaasta.TotalAnnualPeriodProfitLossConsolidated = &value
				case "RevenueConsolidated":
					majandusaasta.RevenueConsolidated = &value
				case "DepreciationAndImpairmentLossReversalConsolidated":
					majandusaasta.DepreciationAndImpairmentLossReversalConsolidated = &value
				case "EmployeeExpenseConsolidated":
					majandusaasta.EmployeeExpenseConsolidated = &value
				case "TotalProfitLossConsolidated":
					majandusaasta.TotalProfitLossConsolidated = &value
				case "DepreciationAndImpairmentLossReversalNegConsolidated":
					majandusaasta.DepreciationAndImpairmentLossReversalNegConsolidated = &value
				case "LaborExpenseConsolidated":
					majandusaasta.LaborExpenseConsolidated = &value
				case "AverageNumberOfEmployeesInFullTimeEquivalentUnitsConsolidated":
					majandusaasta.AverageNumberOfEmployeesInFullTimeEquivalentUnitsConsolidated = &value
				case "TotalRevenue":
					majandusaasta.TotalRevenue = &value
				case "IssuedCapital2Consolidated":
					majandusaasta.IssuedCapital2Consolidated = &value
				default:
					fmt.Println("Unknown key:", key)
				}
				majandusaastad[kandeId] = majandusaasta
			}
		}
	}

	fmt.Printf("Inserting %d majandusaastad\n", len(majandusaastad))

	values := make([]Majandusaasta, 0, len(majandusaastad))
	for _, v := range majandusaastad {
		values = append(values, v)
	}
	db.CreateInBatches(&values, batchSize)

	return nil
}
