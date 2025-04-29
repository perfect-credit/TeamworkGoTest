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

The CSV file should have a header row and at least one email column. For example:

```graphql
Name,Email
Alice,alice@example.com
Bob,bob@example.com
Charlie,charlie@domain.com
```

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

### Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or features you'd like to add.

### License

This project is licensed under the MIT License. See the LICENSE file for more details.
