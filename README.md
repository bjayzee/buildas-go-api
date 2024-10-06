# Go CRUD Application

## Overview

This project is a simple CRUD application built with Go and GORM. The application allows users to register, log in, and perform CRUD operations on user entities. It uses PostgreSQL as the database and is containerized using Docker.

---

## Prerequisites

Before setting up the application, ensure you have the following installed on your local machine:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

---

## Setup and Run Application

## Step 1: Clone the repository

```bash
git clone https://github.com/bjayzee/buildas-go-api
cd buildas-go-api
```


###  Step 2: Create .env file

In the root directory, create a .env file with the following contents:

```bash
DB_HOST=db
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=yourpassword
DB_NAME=go_crud_db
APP_PORT=8081
```

### Step 3: Build and Run the Containers

To start the application using Docker Compose, run the following command in the root of the project:


```bash
docker-compose up --build
```

This command will:

Build the Docker images for both the application and the PostgreSQL database.
Start the application and expose it on port 8081.


## Step 4: Access the Application
Once the containers are up and running, you can access the application by visiting:

```bash
http://localhost:8081/register
```
You can also access the login endpoint at:

```bash
http://localhost:8081/login
```

To get all users

```bash
http://localhost:8081/api/users
```

### Step 5: Stopping the Application
To stop the application, run:

```bash
docker-compose down
```
