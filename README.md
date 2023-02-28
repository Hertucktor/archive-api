# archive-api
This is a RESTful API service that allows for the manipulation of Magic: The Gathering card data in a MongoDB database. 
The API provides endpoints for all CRUD (Create, Read, Update, Delete) operations for card data.

## Requirements
Go (version 1.16 or later)
MongoDB
Postman (optional, for testing)

## Installation
Clone the repository: git clone https://github.com/Hertucktor/archive-api
Change to the project directory: cd mtg-card-api
Install the required packages: go mod tidy

## Configuration
The application requires a MongoDB database to be running locally or remotely. You will need to set the following environment variables:

DB_HOST - the MongoDB database host
DB_PORT - the MongoDB database port
DB_NAME - the name of the MongoDB database to use
DB_USERNAME - (optional) the username to use to authenticate with the MongoDB database
DB_PASSWORD - (optional) the password to use to authenticate with the MongoDB database

## Usage
Start the API server: go run main.go
Use a tool like Postman to test the API endpoints