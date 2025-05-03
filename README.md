# TeamworkGoTest

A simple Go application to read customer data from a CSV file, count email domains, and sort the counts. This tool is designed to efficiently handle large datasets.

## Features

- Reads customer data from a CSV file.
- Counts occurrences of email domains.
- Sorts domains by count in ascending order.

## Getting Started

### Prerequisites

- Go (version 1.16 or higher) installed on your machine.
- A CSV file containing customer data with at least one email column.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/perfect-credit/TeamworkGoTest.git
   cd TeamworkGoTest
   ```

2. Build the application:

   ```bash
   go build
   ```

### Usage

To run the application, use the following command:

```bash
go run main.go path/to/your/customers.csv
```

Replace path/to/your/customers.csv with the actual path to your CSV file.

### CSV Format

The CSV file should have a header row and 5 columns. For example:

```bash
first_name     last_name            email                            gender            ip_address

Mildred        Hernandez            mhernandez0@github.io            Female            38.194.51.128
Bonnie         Ortiz                bortiz1@cyberchimps.com          Male              197.54.209.129
```

But if some cells are empty, this project will find them and output in invalid.csv.

### Example Output

After running the application, you will see the counts of email domains, sorted in ascending order:

```makefile
domain.com: 2
example.com: 4
```

### Testing

To run the tests for the application, use the following command:

```bash
go test ./...
```
