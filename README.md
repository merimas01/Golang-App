# Golang-App
This is a generic CRUD API project built with Go and MySQL. It includes generic BaseService and BaseController for multiple entities, along with Swagger API documentation.

## Prerequisites
Go >= 1.21  
MySQL Server 

## Setup instructions

1. Clone the repository  
`git clone https://github.com/merimas01/Golang-App.git`  
`cd Golang-App/backend`  

2. Create .env file in the project root  
DB_HOST=localhost  
DB_PORT=3306  
DB_USER=root  
DB_PASS=yourpassword  
DB_NAME=yourdbname  
SEED_DATA=true  

3. Run MySQL   
Make sure your database specified in .env exists:  
`CREATE DATABASE yourdbname;`  

4. Run the application  
`go run main.go`  

## Notes
> Swagger is available at http://localhost:8080/swagger/index.html

> `SEED_DATA=true` in .env will automatically seed sample User records, so make sure you switch it to *false* if you don't want to execute the seeding script everytime you run the application

