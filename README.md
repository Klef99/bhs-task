[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Klef99/bhs-task)
[![Swagger Validator](https://img.shields.io/swagger/valid/3.0?specUrl=https%3A%2F%2Fraw.githubusercontent.com%2FKlef99%2Fbhs-task%2Fmain%2Fdocs%2Fswagger.yaml)


# BHS-task

A test assignment for a backend developer at BHS

## Task

1. Go language, JWT Auth
- Create a basic web application in the Go language (the framework is not important) with a JWT-based authentication system. The user must be able to register, log in and log out of the system.


2. Working with PostgreSQL
- Use PostgreSQL to store information about users and their "assets". Create the users and assets tables linked by a one-to-many relationship. Fields for assets: id, name, description, price. Fields for users: id, username, password_hash. Optionally, you can add additional fields.
- Implement the functionality of adding and removing assets for registered users.
- Implement a mechanism for processing asset purchases. It is important to get rid of possible data races.
- Connect the database to the app on Go. Use the migration mechanism (migrate library). To work with the database, you can use both ORM (for example, GORM) and pure SQL and add-ons on it (for example, SQLX, using squirrel will also be a plus).


3. API and documentation
- Create a RESTful API for working with assets and users. Try to adhere to the rules of clean architecture (using a template will be a plus https://github.com/evrone/go-clean-template ). 
- Create API documentation using Swagger or other similar tools.


4. Tests
- Write unit tests to test all the above-mentioned functions. Preferably test coverage of at least 50%.
- *Optionally write functional tests to verify HTTP requests. It is allowed to use frameworks for testing (https://github.com/ozontech/cute or https://github.com/lamoda/gonkey

5. *Makefile
- Create and write a Makefile with the ability to run the entire application with the infrastructure via docker-compose (make up) and only the infrastructure (in our case, only the database) via docker-compose or a separate container (make dev-up). Additional commands are not required, but they will be a plus (make test, make dev-down, make down, etc.).


## Tech Stack

**Server:** Go, JWT, PostgreSQL, Chi


## Installation

### Install make, go, docker and docker-compose (if it not installed):

For Debian/Ubuntu based systems:
```bash
# Update package index
sudo apt update

# Install make, Go, Docker, Docker Compose, and golangci-lint
sudo apt install -y make golang docker.io docker-compose curl

# Install golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

# Start and enable Docker service
sudo systemctl start docker
sudo systemctl enable docker

# Verify installation
make --version
go version
docker --version
docker-compose --version
golangci-lint --version
```

For CentOS/RHEL based systems:
```bash
# Update package index
sudo yum update -y

# Install make, Go, Docker, Docker Compose, and golangci-lint
sudo yum install -y make golang docker curl
sudo systemctl start docker
sudo systemctl enable docker

# Install Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Install golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

# Verify installation
make --version
go version
docker --version
docker-compose --version
golangci-lint --version
```

For macOS using [Homebrew](https://brew.sh/):
```bash
# Install dependencies using Homebrew
brew install make go docker docker-compose golangci-lint

# Verify installation
make --version
go version
docker --version
docker-compose --version
golangci-lint --version
```
## Run Locally

Clone the project

```bash
  git clone https://github.com/Klef99/bhs-task.git
```

Go to the project directory

```bash
  cd bhs-task
```

Start the server

```bash
  make up
```

### Avaliable commands
```bash
Usage:
  make <target>
  help             Display this help screen
  up               Run docker-compose
  down             Down docker-compose
  dev-up           Up infrastructure
  dev-down         Down infrastructure
  mock-generate    generate mocks
  rm-volume        remove docker volume
  linter-golangci  check by golangci linter
  test             run test
  migrate-create   create new migration
  migrate-up       migration up
```
## Running Tests

To run tests, run the following command

```bash
  make test
```

## API Documentation
[![Swagger](https://img.shields.io/badge/swagger-docs-brightgreen)](http://localhost:8080/swagger/index.html)

The API description is available at ```http://<ip or domain>:<port>/swagger/index.html```. 

If the project is running locally with the standard settings: ```http://localhost:8080/swagger/index.html```
## Acknowledgements

 - [Go Clean templates](https://github.com/evrone/go-clean-template/tree/master)