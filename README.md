# Receipt Processor - Fetch Assessment

This project implements a Go-based webservice for the Fetch Receipt Processor assessment. It processes receipts via a REST API, calculates points based on specified rules, and stores data in memory. The implementation fulfills the API defined in `api.yml`, with endpoints for submitting receipts and retrieving points.

## Features

- **API Endpoints**:
  - `POST /receipts/process`: Accepts a receipt JSON, assigns a unique ID, and returns the ID.
  - `GET /receipts/{id}/points`: Returns the calculated points for a receipt by its ID.
- **Points Calculation**: Implements all rules, including retailer name, total amount, item counts, item descriptions, purchase date, and time-based conditions.
- **In-Memory Storage**: Uses a Go map to store receipts, ensuring no persistence after restart.
- **Error Handling**: Returns HTTP 400 for invalid receipts with "The receipt is invalid. Please verify input." and HTTP 404 for unknown IDs.
- **Dependencies**: Uses `github.com/google/uuid` for ID generation and `github.com/gorilla/mux` for routing.

## Prerequisites

- Go 1.21 or later
- Git (to clone the repository)
- No external database or software required

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/Ruke3663/receipt-processor
   cd receipt-processor
2. **Install Dependencies**
    ```bash
    go mod tidy

## Running the Application

1. **Verify Package Structure**:
   Ensure main.go and receipt.go are in the same folder (project root) and both start with package main.
   
2. **Start the Server**
    ```bash
    go run .
    go run main.go receipt.go

Server start on http://localhost:8080
