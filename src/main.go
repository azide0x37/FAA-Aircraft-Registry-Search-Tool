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
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Aircraft struct {
	NNumber        string
	SerialNumber   string
	MfrMdlCode     string
	EngMfrMdl      string
	YearMfr        string
	TypeRegistrant string
	Name           string
	Street         string
	Street2        string
	City           string
	State          string
	ZipCode        string
	Region         string
	County         string
	Country        string
	LastActionDate string
	CertIssueDate  string
	Certification  string
	TypeAircraft   string
	TypeEngine     string
	StatusCode     string
	ModeSCode      string
	FractOwner     string
	AirWorthDate   string
	OtherNames     [5]string
	ExpirationDate string
	UniqueID       string
	KitMfr         string
	KitModel       string
	ModeSCodeHex   string
}

func main() {
	nNumber := flag.String("n-number", "", "N-Number to search")
	flag.Parse()

	if *nNumber == "" {
		fmt.Println("Error: Missing required flag --n-number")
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

	allAircrafts, err := parseCSVFile("MASTER.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	modelCode := findModelCodeByNNumber(*nNumber, allAircrafts)
	kitMfr, kitModel := findKitInfoByNNumber(*nNumber, allAircrafts)

	if modelCode == "" && (kitMfr == "" || kitModel == "") {
		fmt.Println("No aircraft found with the provided N-Number")
		return
	}

	var matchingAircrafts []Aircraft
	for _, aircraft := range allAircrafts {
		if aircraft.MfrMdlCode == modelCode || (aircraft.KitMfr == kitMfr && aircraft.KitModel == kitModel) {
			matchingAircrafts = append(matchingAircrafts, aircraft)
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	// curvy borders

	t.AppendHeader(table.Row{"N-Number", "SerialNumber", "MfrMdlCode", "EngMfrMdl", "YearMfr", "Name", "Street", "City", "State", "ZipCode", "LastActionDate", "CertIssueDate", "Certification", "ModeSCode", "FractOwner", "AirWorthDate", "ExpirationDate", "UniqueID", "KitMfr", "KitModel", "ModeSCodeHex"})

	for _, aircraft := range matchingAircrafts {
		t.AppendRow(table.Row{aircraft.NNumber, aircraft.SerialNumber, aircraft.MfrMdlCode, aircraft.EngMfrMdl, aircraft.YearMfr, aircraft.Name, aircraft.Street, aircraft.City, aircraft.State, aircraft.ZipCode, aircraft.LastActionDate, aircraft.CertIssueDate, aircraft.Certification, aircraft.ModeSCode, aircraft.FractOwner, aircraft.AirWorthDate, aircraft.ExpirationDate, aircraft.UniqueID, aircraft.KitMfr, aircraft.KitModel, aircraft.ModeSCodeHex})
	}

	t.Render()
	fmt.Printf("Found %d matching aircrafts\n", len(matchingAircrafts))

}

func findModelCodeByNNumber(nNumber string, aircrafts []Aircraft) string {
	for _, aircraft := range aircrafts {
		if aircraft.NNumber == nNumber {
			return aircraft.MfrMdlCode
		}
	}
	return ""
}

func findKitInfoByNNumber(nNumber string, aircrafts []Aircraft) (string, string) {
	for _, aircraft := range aircrafts {
		if aircraft.NNumber == nNumber {
			return aircraft.KitMfr, aircraft.KitModel
		}
	}
	return "", ""
}

func parseCSVFile(filename string) ([]Aircraft, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true // Add this line to handle unescaped quotes

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var aircraftList []Aircraft

	for i, record := range records {
		if i == 0 {
			continue // skip header
		}

		otherNames := [5]string{}
		for j := 0; j < 5; j++ {
			otherNames[j] = strings.TrimSpace(record[22+j])
		}

		aircraft := Aircraft{
			NNumber:        strings.TrimSpace(record[0]),
			SerialNumber:   strings.TrimSpace(record[1]),
			MfrMdlCode:     strings.TrimSpace(record[2]),
			EngMfrMdl:      strings.TrimSpace(record[3]),
			YearMfr:        strings.TrimSpace(record[4]),
			TypeRegistrant: strings.TrimSpace(record[5]),
			Name:           strings.TrimSpace(record[6]),
			Street:         strings.TrimSpace(record[7]),
			Street2:        strings.TrimSpace(record[8]),
			City:           strings.TrimSpace(record[9]),
			State:          strings.TrimSpace(record[10]),
			ZipCode:        strings.TrimSpace(record[11]),
			Region:         strings.TrimSpace(record[12]),
			County:         strings.TrimSpace(record[13]),
			Country:        strings.TrimSpace(record[14]),
			LastActionDate: strings.TrimSpace(record[15]),
			CertIssueDate:  strings.TrimSpace(record[16]),
			Certification:  strings.TrimSpace(record[17]),
			TypeAircraft:   strings.TrimSpace(record[18]),
			TypeEngine:     strings.TrimSpace(record[19]),
			StatusCode:     strings.TrimSpace(record[20]),
			ModeSCode:      strings.TrimSpace(record[21]),
			FractOwner:     strings.TrimSpace(record[22]),
			AirWorthDate:   strings.TrimSpace(record[23]),
			OtherNames:     otherNames,
			ExpirationDate: strings.TrimSpace(record[27]),
			UniqueID:       strings.TrimSpace(record[28]),
			KitMfr:         strings.TrimSpace(record[29]),
			KitModel:       strings.TrimSpace(record[30]),
			ModeSCodeHex:   strings.TrimSpace(record[31]),
		}

		aircraftList = append(aircraftList, aircraft)
	}

	return aircraftList, nil
}

func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func unzipFile(zipfile string, outfile string) error {
	reader, err := zip.OpenReader(zipfile)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.Name == outfile {
			rc, err := file.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			out, err := os.Create(outfile)
			if err != nil {
				return err
			}
			defer out.Close()

			_, err = io.Copy(out, rc)
			return err
		}
	}

	return fmt.Errorf("file %s not found in zip file %s", outfile, zipfile)
}
