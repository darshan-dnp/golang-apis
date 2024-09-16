# golang-apis

GET and POST APIs in Golang to post purchase receipt and get reward points

## Endpoint: Process Receipts

### Description

- **Path**: `/receipts/process`
- **Method**: POST
- **Payload**: Receipt JSON
- **Response**: JSON containing an ID for the receipt.

### Points Calculation Rules

1. **Retailer Name**: Award 1 point for every alphanumeric character in the retailer name.
2. **Total Amount**:
   - Award 50 points if the total amount is a round dollar amount with no cents.
   - Award 25 points if the total amount is a multiple of 0.25.
3. **Number of Items**: Award 5 points for every two items on the receipt.
4. **Item Description**:
   - If the trimmed length of the item description is a multiple of 3, multiply the item price by 0.2, round up to the nearest integer, and use this value as the number of points earned for that item.
5. **Purchase Date**: Award 6 points if the day of the purchase date is odd.
6. **Time of Purchase**: Award 10 points if the time of purchase is between 2:00 PM and 4:00 PM.

### Example Payload

```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },
    {
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },
    {
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },
    {
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },
    {
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```

### Example Response Body

```json
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

## Endpoint: Get Points Awarded to Receipt

### Description

- **Path**: `/receipts/{id}/points`
- **Method**: GET
- **Response**: JSON containing Points for the receipt.

### Example Response Body

```json
{ "points": 28 }
```

## Prerequisites

Before running the application, ensure that the following software is installed on your system:

1. **Docker**: Install Docker from the [official website](https://www.docker.com/get-started).
2. **Golang (Optional for development)**: If you plan to work on the API codebase, install Golang from the [official website](https://golang.org/doc/install).

## Installation

To get started with the project, follow these steps:

1. **Clone the repository**:
   Open your terminal and run:
   ```bash
   git clone git@github.com:darshan-dnp/golang-apis.git
   cd golang-apis
   ```

## Docker Setup

To build and run the application using Docker:

1. **Build the Docker image**:
   In the root directory of the project, run the following command:

   ```bash
   docker build . -t golang-apis
   ```

   This command creates a Docker image tagged as golang-apis.

2. **Run the Docker container**:
   Once the image is built, run the container with the following command:

   ```bash
   docker run --env-file .env -p 8080:8080 golang-apis
   ```

   This will start the API and map it to port 8080 on your local machine. You can access the API at http://localhost:8080.

3. **Environment variables in Docker**:
   The environment variables of docker container are being set using `.env` file. Command to run docker container will change as per the changes of PORT in `.env` file.

## Usage

1. **Example API request (using curl)**:

   - **Create receipt**:

   ```bash
   curl -X POST http://localhost:8080/receipts/process \
   -H "Content-Type: application/json" \
   -d '{
   "retailer": "Target",
   "purchaseDate": "2022-01-01",
   "purchaseTime": "13:01",
   "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    }
   ],
   "total": "35.35"
   }'

   ```

   - **Get Receipt Points**

   ```bash
   curl -X GET http://localhost:8080/receipts/ID_RECEIVED_FROM_PREVIOUS_POST_REQ/points
   ```

2. **Use tools like Postman**

## Config

- **Database**: The application is using in-memory sqllite database. The required tables will be created based on models.

- **Rules Config**: The rules mentioned in description are stored in database. As currently the app is using in-memory database, these rules will be stored at start of the application. The config can be changed in codebase.
