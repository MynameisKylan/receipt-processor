# Receipt Processor

Built as part of the [Fetch Rewards Challenge](https://github.com/fetch-rewards/receipt-processor-challenge).

### Prerequisites

- Go 1.22 or later
- Git

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/MynameisKylan/receipt-processor.git
   cd receipt-processor
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

### Running the Server

Start the server on the default port (8181):
```bash
go run main.go
```

To specify a custom port:
```bash
go run main.go -port 3000
```

### API Endpoints

The service exposes two endpoints:

1. **Process Receipts** - Submit a receipt for processing
   - Path: `/receipts/process`
   - Method: POST
   - Request Body: Receipt JSON
   - Response: JSON containing an ID for the receipt

2. **Get Points** - Get the points awarded for a receipt
   - Path: `/receipts/{id}/points`
   - Method: GET
   - Response: JSON containing the points awarded

#### Example
```bash
curl http://localhost:8181/receipts/abc/points
```

