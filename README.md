# Go-Nurse-mgmt

1. Navigate to the project directory:
cd go-nurse-mgmt

2. Install project dependencies:
go mod tidy

3. Configure the database connection in the project's configuration file (if applicable).

### Usage
Start the server:
go run cmd/main.go

Access the API using a tool like Postman, or develop a front-end application to interact with it.

### Example API endpoints:

GET /nurses: Retrieve a list of nurses.
POST /nurse: Add a new nurse.
PUT /nurse/{id}: Update a nurse by ID.
DELETE /nurse/{id}: Delete a nurse by ID.
POST /login: Login.
POST /signup: Sign up user

## Project Structure
cmd/main.go: The main application entry point.
routes/: Contains the  routes.
database/: Database models.
controllers/: Request handlers.
middlewares/: Custom middleware functions.
utils/: Utility functions.

## Technologies Used
Go (Golang)
Gorilla Mux (for routing)
Database (PostgreSQL)
Authentication ( JWT )
