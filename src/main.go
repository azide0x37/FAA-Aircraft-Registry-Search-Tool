package main

import (
	"archive/zip"
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Aircraft struct {
	LastActivityDate string
	NNumber          string
	YearMfr          string
	RegistrantInfo   string
	Address          string
}

func main() {
	var matchType string
	flag.StringVar(&matchType, "match-type", "", "NNumber to match")
	flag.Parse()

	if matchType == "" {
		fmt.Println("Error: Please provide an NNumber using the --match-type flag")
		return
	}

	err := downloadAndUnzip("https://registry.faa.gov/database/ReleasableAircraft.zip", "MASTER.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	file, err := os.Open("MASTER.TXT")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	modelCode, err := findModelCode(file, matchType)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	aircrafts := parseAircraftData(file, modelCode)
	sortAircrafts(aircrafts)

	fmt.Printf("%-20s %-10s %-10s %-50s %-100s\n", "Last Activity Date", "N-Number", "Year Mfr", "Registrant Info", "Address")
	for _, aircraft := range aircrafts {
		fmt.Printf("%-20s %-10s %-10s %-50s %-100s\n", aircraft.LastActivityDate, aircraft.NNumber, aircraft.YearMfr, aircraft.RegistrantInfo, aircraft.Address)
	}
	fmt.Printf("\nTotal aircraft: %d\n", len(aircrafts))
}

// Rest of the functions remain the same
func downloadAndUnzip(url, targetFileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tmpFile, err := ioutil.TempFile("", "faa_data_*.zip")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return err
	}

	reader, err := zip.OpenReader(tmpFile.Name())
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, f := range reader.File {
		if f.Name == targetFileName {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			outputFile, err := os.Create(targetFileName)
			if err != nil {
				return err
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, rc)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("file %s not found in the archive", targetFileName)
}

func findModelCode(file *os.File, nNumber string) (string, error) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		record := strings.Split(line, "|")

		if len(record) < 3 {
			continue
		}

		if record[0] == nNumber {
			return record[2], nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("NNumber %s not found", nNumber)
}

func parseAircraftData(file io.Reader, modelCode string) []Aircraft {
	reader := csv.NewReader(file)
	aircrafts := []Aircraft{}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if record[2] == modelCode {
			registrantInfo := record[8] + " " + record[10] + " " + record[11] + " " + record[12] + " " + record[13]
			address := record[9] + ", " + record[14] + ", " + record[15] + " " + record[16]

			aircraft := Aircraft{
				LastActivityDate: record[21],
				NNumber:          record[0],
				YearMfr:          record[5],
				RegistrantInfo:   registrantInfo,
				Address:          address,
			}
			aircrafts = append(aircrafts, aircraft)
		}
	}

	return aircrafts
}

func sortAircrafts(aircrafts []Aircraft) {
	sort.Slice(aircrafts, func(i, j int) bool {
		date1 := strings.Split(aircrafts[i].LastActivityDate, "/")
		date2 := strings.Split(aircrafts[j].LastActivityDate, "/")

		if len(date1) < 3 || len(date2) < 3 {
			return false
		}

		year1, _ := strconv.Atoi(date1[0])
		year2, _ := strconv.Atoi(date2[0])
		month1, _ := strconv.Atoi(date1[1])
		month2, _ := strconv.Atoi(date2[1])
		day1, _ := strconv.Atoi(date1[2])
		day2, _ := strconv.Atoi(date2[2])

		if year1 != year2 {
			return year1 < year2
		}
		if month1 != month2 {
			return month1 < month2
		}
		return day1 < day2
	})
}
