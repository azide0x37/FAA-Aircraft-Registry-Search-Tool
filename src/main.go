package main

import (
	"encoding/csv"
	"fmt"
	"io"
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
	// Replace "AircraftData.csv" with the path to your data file
	file, err := os.Open("Aircraft.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Replace "1234567" with the desired Aircraft Mfr Model Code
	aircrafts := parseAircraftData(file, "06001BR")
	sortAircrafts(aircrafts)

	fmt.Printf("%-20s %-10s %-10s %-50s %-100s\n", "Last Activity Date", "N-Number", "Year Mfr", "Registrant Info", "Address")
	for _, aircraft := range aircrafts {
		fmt.Printf("%-20s %-10s %-10s %-50s %-100s\n", aircraft.LastActivityDate, aircraft.NNumber, aircraft.YearMfr, aircraft.RegistrantInfo, aircraft.Address)
	}
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
