# desafio-cierre-db

---

## Welcome to desafio-cierre-db!

This repository contains the necessary components to run our database challenge application, including a data loading job and an API server. We leverage Docker for easy environment setup and `Makefile` targets to streamline common operations.

---

## Getting Started

To get started, you'll need to have Docker and `make` installed on your system.

2.  **Environment Variables:**
    Create a `.env` file in the root directory for your application's environment variables.
    For testing, create a `.env.test` file.

    *For the course purpose I uploaded this files into the repository even though it`s not the best practice*

---

## Makefile Targets Documentation

Our `Makefile` provides convenient commands to manage your application and its Docker environment.

### Application Control

* **`make start-load`**
    Starts the **data loading job**. This command runs the Go application with the `'load-job'` argument which is responsible for loading the json files into the db, ensuring all necessary Docker services are up and environment variables from `.env` are loaded.

* **`make start-api`**
    Starts the **API server**. This command runs the Go application with the `'api'` argument, ensuring Docker services are up and environment variables from `.env` are loaded.

### Docker Management

* **`make docker-up`**
    Brings up the **main Docker services** defined in `docker/docker-compose.yml` in detached mode. Use this to get your database and other services running for development.

* **`make docker-down`**
    Stops and removes all **main Docker services** and their associated volumes defined in `docker/docker-compose.yml`. This is useful for a clean shutdown or to free up resources.

* **`make docker-test-up`**
    Brings up the **test Docker services** defined in `docker/docker-compose.test.yml` in detached mode. This command forces recreation of containers (`--force-recreate`), even if their configuration and image haven't changed, ensuring a fresh test environment.

### Testing

* **`make test`**
    Runs all **Go tests**. This command first ensures that the test Docker services are up (`docker-test-up`) and then executes the Go tests using environment variables loaded from `.env.test`.

---

## Project Structure

I maintained the structure suggested on the base repository for this challenge.