# FAA Aircraft Registry Search Tool

This command-line tool allows you to search for aircraft in the FAA Aircraft Registry database using their N-Number. The tool will download the latest available registry database, extract the relevant information, and display the results in a formatted table.

## Features

- Download and extract the FAA Aircraft Registry database
- Search for an aircraft using its N-Number
- Display the aircraft details in a formatted table

## Prerequisites

- Golang (version 1.16 or later)

## Installation

1. Clone the repository:

```sh
git clone https://github.com/user/faa-aircraft-registry-search.git
```

2. Change to the project directory:

```
cd faa-aircraft-registry-search
```

3. Build the project:
```sh
go build
```

## Usage
To search for an aircraft using its N-Number, run the following command:

```sh
./faa-aircraft-registry-search --n-number `<N_NUMBER>`
```

Replace `<N_NUMBER>` with the N-Number of the aircraft you are looking for.

Example:

```sh
./faa-aircraft-registry-search --n-number N12345
```

The tool will download the latest FAA Aircraft Registry database, extract the relevant information, and display the matching aircraft details in a formatted table.

License
This project is licensed under the MIT License. See the LICENSE file for details.
