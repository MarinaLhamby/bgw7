# Docker Compose Project

This project sets up a MySQL database using Docker Compose. It includes a MySQL service with an initialization script to set up the database schema and data.

## Project Structure

```
docker-compose-project
├── docker-compose.yml
├── init.sql
└── README.md
```

## Getting Started

To build and run the Docker containers, follow these steps:

1. Ensure you have Docker and Docker Compose installed on your machine.

2. Clone this repository or download the project files.

3. Navigate to the project directory:

   ```bash
   cd docker-compose-project
   ```

4. Build and start the containers using Docker Compose:

   ```bash
   docker-compose up
   ```

5. The MySQL service will be available on port 3306. You can connect to it using the following credentials:

   - **Root Password:** rootpassword
   - **Database Name:** movies_db
   - **User:** myuser
   - **Password:** mypassword

## Initialization

The `init.sql` file contains SQL commands that will be executed when the MySQL container is initialized. You can modify this file to set up your initial database schema and data.

## Stopping the Containers

To stop the running containers, use:

```bash
docker-compose down
```

This command will stop and remove the containers defined in the `docker-compose.yml` file.