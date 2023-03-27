/*
Description of MASTER.txt file

AIRCRAFT REGISTRATION MASTER FILE
Contains the records of all U.S. Civil Aircraft maintained by the FAA, Civil Aviation Registry, Aircraft
Registration Branch, AFS-750
Element
Position
Location in
Record
Number of
Positions Descriptions
N-Number 1-5 5 Identification number assigned to aircraft.
Serial Number 7-36 30 The complete aircraft serial number
assigned to the aircraft by the
manufacturer.
Aircraft Mfr Model
Code
38-44 7 A code assigned to the aircraft
manufacturer, model and series.
Positions (38-40) - Manufacturer Code
Positions (41-42) - Model Code
Positions (43-44) - Series Code
Engine Mfr Mode Code 46-50 5 A code assigned to the engine
manufacturer and model.
Positions (46-48) - Manufacturer Code
Positions (49-50) - Model Code
Year Mfr 52-55 4 Year manufactured.
Type Registrant 57 1 1 - Individual
2 - Partnership
3 - Corporation
4 - Co-Owned
5 – Government
7 - LLC
8 - Non Citizen Corporation
9 - Non Citizen Co-Owned
Registrant’s Name 59-108 50
The first registrant’s name which appears
on the Application for Registration, AC
Form 8050-1.
Street1 110-142 33 The street address which appears on the
Application for Registration, AC Form
8050-1 or the latest street address
reported.
Street2 144-176 33 The 2nd street address which appears on
the Application for Registration, AC,
Form 8050-1, or the latest street address
reported.
12/3/2020 T:/ARS-DEV/ONLINE-SUPPORT/WEBFILES.DOC
Element
Position
Location in
Record
Number of
Positions Descriptions
Registrant’s City 178-195 18 The city name which appears on the
Application for Registration, AC Form
8050-1 or the latest address reported.
Registrant’s State 197-198 2 The state name which appears on the
Application for Registration, AC Form
8050-1 or the latest address reported.
Registrant’s Zip Code 200-209 10 The postal Zip Code which appears on the
Application for Registration, AC Form
8050-1 or the latest address reported.
Registrant’s Region 211 1 1 - Eastern
2 - Southwestern
3 - Central
4 - Western-Pacific
5 - Alaskan
7 - Southern
8 - European
C- Great Lakes
E - New England
S - Northwest Mountain
County Mail 213-215 3 A code representing the county which
appears on the Application for
Registration.
Country Mail 217-218 2 A code representing the country which
appears on the Application for
Registration.
Last Activity Date 220-227 8 Format YYYY/MM/DD
Certificate Issue Date 229-236 8 Format YYYY/MM/DD
Certification requested
and uses:

	A - Airworthiness
	Classification Code

238 1 The airworthiness certificate class which
is reported on the Application for
Airworthiness, FAA Form 8130-6.
1 - Standard
2 - Limited
3 - Restricted
4 - Experimental
5 - Provisional
6 – Multiple
7 - Primary
8 - Special Flight Permit
12/3/2020 T:/ARS-DEV/ONLINE-SUPPORT/WEBFILES.DOC
Element
Position
Location in
Record
Number of
Positions Descriptions
9 – Light Sport

	B - Approved
	Operation Codes

239-247 9 The approved operations (up to a
maximum of six) which are reported on
the Application for Airworthiness, FAA
Form 8130-6.

	Standard 238 1

239 1 Blank
N - Normal
U - Utility
A - Acrobatic
T - Transport
G - Glider
B - Balloon
C - Commuter
O - Other
240-247 8 Positions are blank.

	Limited 238 1

239-247 9 Positions are blank.

	Restricted 238 1

239-245 7 May contain a code of 0-7.
0 - Other
1 - Agriculture and Pest Control
2 - Aerial Surveying
3 - Aerial Advertising
4 - Forest
5 - Patrolling
6 - Weather Control
7 - Carriage of Cargo
246-247 2 Positions are blank

	Experimental 238 1

239-245 7 May contain a code of 0-9.
0 - To show compliance with FAR
1 - Research and Development
2 - Amateur Built
3 - Exhibition
4 - Racing
5 - Crew Training
6 - Market Survey
12/3/2020 T:/ARS-DEV/ONLINE-SUPPORT/WEBFILES.DOC
Element
Position
Location in
Record
Number of
Positions Descriptions
7 - Operating Kit Built Aircraft
8A - Reg. Prior to 01/31/08
8B - Operating Light-Sport Kit-Built
8C - Operating Light-Sport Previously

	issued cert under 21.190

9A - Unmanned Aircraft - Research and
Development
9B - Unmanned Aircraft - Market Survey
9C - Unmanned Aircraft - Crew Training
9D - Unmanned Aircraft – Exhibition
9E - Unmanned Aircraft – Compliance

	With CFR

246-247 2 Positions are blank

	Provisional 238 1

239 1 1 - Class I
2 - Class II
240-247 8 Positions are blank

	Multiple 238 1

239-240 2 May contain a code of 1-3.
1 - Standard
2 - Limited
3 - Restricted
241-247 7 May be blank or contain:
0 - Other
1 - Agriculture and Pest Control
2 - Aerial Surveying
3 - Aerial Advertising
4 - Forest
5 - Patrolling
6 - Weather Control
7 - Carriage of Cargo

	Primary 238 1

239-247 9 Positions are blank

	Special Flight
	Permit

238 1
239-247 9 May contain a code of 1-6.
1 - Ferry flight for repairs, alterations,
maintenance or storage
2 - Evacuate from area of impending
12/3/2020 T:/ARS-DEV/ONLINE-SUPPORT/WEBFILES.DOC
Element
Position
Location in
Record
Number of
Positions Descriptions
danger
3 - Operation in excess of maximum
certificated
4 - Delivery or export
5 - Production flight testing
6 - Customer Demo

	Light Sport 238 1

239 1 May contain a code of A-W.
A - Airplane
G - Glider
L - Lighter than Air
P - Power-Parachute
W- Weight-Shift-Control
240-247 8 Positions are blank
Type Aircraft 249 1 1 - Glider
2 - Balloon
3 - Blimp/Dirigible
4 - Fixed wing single engine
5 - Fixed wing multi engine
6 - Rotorcraft
7 - Weight-shift-control
8 - Powered Parachute
9 - Gyroplane
H - Hybrid Lift
O - Other
Type Engine 251-252 2 0 - None
1 - Reciprocating
2 - Turbo-prop
3 - Turbo-shaft
4 - Turbo-jet
5 - Turbo-fan
6 - Ramjet
7 - 2 Cycle
8 - 4 Cycle
9 – Unknown
10 – Electric
11 - Rotary
*/
package main

import (
	"archive/zip"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
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
	matchCode := flag.String("match-code", "", "Model code to search")
	flag.Parse()

	if *matchCode == "" {
		fmt.Println("Error: Missing required flag --match-code")
		return
	}

	err := downloadFile("https://registry.faa.gov/database/ReleasableAircraft.zip", "ReleasableAircraft.zip")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = unzipFile("ReleasableAircraft.zip", "MASTER.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Parse the aircraft data
	file, err := os.Open("MASTER.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	aircrafts := parseAircraftData(file, *matchCode)
	sortAircrafts(aircrafts)

	fmt.Printf("%-20s %-10s %-10s %-50s %-100s\n", "Last Activity Date", "N-Number", "Year Mfr", "Registrant Info", "Address")
	for _, aircraft := range aircrafts {
		fmt.Printf("%-20s %-10s %-10s %-50s %-100s\n", aircraft.LastActivityDate, aircraft.NNumber, aircraft.YearMfr, aircraft.RegistrantInfo, aircraft.Address)
	}

	fmt.Printf("Total: %d\n", len(aircrafts))
}

func parseAircraftData(file *os.File, matchCode string) []Aircraft {
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	var manufacturerCode string
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

		// Check length of record to avoid index out of range error
		if len(record) < 17 {
			continue
		}
		// Check if the model code matches and set the manufacturerCode
		if strings.TrimSpace(record[0]) == matchCode && manufacturerCode == "" {
			manufacturerCode = strings.TrimSpace(record[2])
			fmt.Println("Manufacturer Code:", manufacturerCode)
		}
	}

	// Reset the reader to the beginning of the file
	file.Seek(0, io.SeekStart)
	reader = csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Check length of record to avoid index out of range error
		if len(record) < 17 {
			continue
		}

		// Check if the manufacturer code matches
		if strings.TrimSpace(record[2]) == manufacturerCode {
			registrantInfo := record[6]
			address := record[7] + ", " + record[9] + ", " + record[10] + " " + record[11]

			aircraft := Aircraft{
				LastActivityDate: record[25],
				NNumber:          record[0],
				YearMfr:          record[4],
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

func downloadFile(url, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Download the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func unzipFile(zipfile, targetfile string) error {
	// Open the zip file
	r, err := zip.OpenReader(zipfile)
	if err != nil {
		return err
	}
	defer r.Close()

	// Loop through the files in the zip file
	for _, f := range r.File {
		// Only extract the target file
		if f.Name != targetfile {
			continue
		}

		// Open the file from the zip archive
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Create the target file
		target, err := os.Create(targetfile)
		if err != nil {
			return err
		}
		defer target.Close()

		// Write the file to disk
		_, err = io.Copy(target, rc)
		if err != nil {
			return err
		}

		// Exit the loop after finding the target file
		break
	}

	return nil
}
