<h3>User Service</h3>

<h3>Description</h3>

<p>This repository will be used to manage user and auth</p>

<h3>Directory Structure</h3>

```
skeleton
    L cmd                            → Contains the main entry point or initial configuration of the application
    L common                         → Stores common functions used throughout the application
    L config                         → Contains application configurations such as environment variables and other settings
    L constants                      → Stores global constant values used across the application
    L controllers                    → Manages control logic for handling HTTP requests
    L database                       → Contains files related to database management
        L seeders                    → Scripts for populating initial (seed) data into the database
    L domain                         → The application's domain module containing core domain elements
        L dto                        → Data Transfer Objects, used to define the structure of transferred data
        L models                     → Object models representing the application's or database's data structure
    L middlewares                    → Contains middleware for processing requests/responses before or after reaching the controller
    L repositories                   → Contains data access logic for interacting with the database
    L routes                         → Contains API route definitions
    L services                       → Stores the application's core business logic
```

## How to setup

```
- Clone this repository
- go mod tidy
```

## How to run

```bash
go run main.go
```

## How to build image docker
```bash
docker compose build
```

## How to run with docker

```bash
docker compose up -d
```



