# desafio-cierre-db
# Makefile Targets Documentation
#
# start-load: Starts the load-job process by running the Go application with the 'load-job' argument.
#             Loads environment variables from the .env file and ensures Docker services are up.
#
# start-api:  Starts the API server by running the Go application with the 'api' argument.
#             Loads environment variables from the .env file and ensures Docker services are up.
#
# docker-up:  Brings up the Docker services defined in docker/docker-compose.yml in detached mode.
#
# docker-down: Stops and removes the Docker services and associated volumes defined in docker/docker-compose.yml.
#
# docker-test-up: Brings up the test Docker services defined in docker/docker-compose.test.yml in detached mode,
#                 recreating containers even if their configuration and image haven't changed.
#
# test:       Runs Go tests with environment variables from .env.test, ensuring test Docker services are up first.