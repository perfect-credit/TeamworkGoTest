# TeamworkGoTest

A simple Go application to read customer data from a CSV file, counts the occurrences of each email domain, and outputs the sorted list of domains along with their respective counts. This tool is designed to efficiently handle large datasets.

## Features

- Reads customer data from a CSV file.
- Counts occurrences of email domains.
- Sorts domains by count in ascending order.

## Design Principle

```bash
   This project designed by Domain Driven Design (DDD) principle.
```

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
   go build -o main.exe ./cmd/customerimporter
   ```

### Usage

To display in terminal, use the following command:

```bash
   ./main.exe -input ./data/customers.csv
   ./main.exe -input ./data/customers.csv -sort domain
   ./main.exe -input ./data/customers.csv -sort count
```

To run the application, use the following command:

```bash
   ./main.exe -input ./data/customers.csv -sort domain -output ./data/output.csv
   ./main.exe -input ./data/customers.csv -sort count -output ./data/output.csv
```

Replace path/to/your/customers.csv with the actual path to your CSV file.

### CSV Format

The CSV file should have a header row and 5 columns. For example:

```bash
   first_name     last_name            email                            gender            ip_address

   Mildred        Hernandez            mhernandez0@github.io            Female            38.194.51.128
   Bonnie         Ortiz                bortiz1@cyberchimps.com          Male              197.54.209.129
```

However, if some cells are empty or their data types don't match the data context, this project outputs those rows to data/invalid.csv. So clients can confirm missing part in source file.

```bash
   row   first_name     last_name            email                            gender            ip_address

   8     Mildred        Hernandez            github.io                        Female            38.194.51.128
   15    Bonnie                              bortiz1@cyberchimps.com          Male              197.54.209.129
```

### Example Output

After running the application, you will see the counts of email domains, sorted in ascending order by domain or count

```makefile
   domain.com: 2
   example.com: 4
```

### Testing

To run the tests for the application, use the following command:

```bash
   go test ./...
```
