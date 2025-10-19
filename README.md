# Golang-App
This is a generic CRUD API project built with Go and MySQL. It includes generic BaseService and BaseController for multiple entities, along with Swagger API documentation.

## Prerequisites
Go >= 1.21  
MySQL Server 

## Setup instructions

1. Clone the repository  
`git clone https://github.com/merimas01/Golang-App.git`  
`cd Golang-App/backend`  

2. Create .env file in the project root (*named: backend*)  
DB_HOST=localhost  
DB_PORT=3306  
DB_USER=root  
DB_PASS=yourMySQLpassword  
DB_NAME=yourMySQLdbname  
SEED_DATA=true  

3. Run the application  
`go run main.go`  

## Notes
> Swagger is available at http://localhost:8080/swagger/index.html

> If `SEED_DATA` in .env is set to ***true***, everytime you run the application, the seeding script will be executed 

> Everytime the seeding script is executed it won't produce duplicate records

