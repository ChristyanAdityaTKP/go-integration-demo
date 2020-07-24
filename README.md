# Go Integration Test Demo

A simple demo of integration testing written in Go with PostgreSQL as our database.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.
### Prerequisites

**Go**: https://golang.org/

**Docker** and **Docker Compose**: https://docs.docker.com/get-docker/

**PostgreSQL Client `psql`**: https://blog.timescale.com/tutorials/how-to-install-psql-on-mac-ubuntu-debian-windows/


### Running the app
To run the integration tests, simply run:
```
make integration-test
```
The above command will spin up a PostgreSQL database using `docker-compose` and run the tests in `integration-tests` directory. It will print out the result of the tests.

**Note**: This command does not automatically terminate the PostgreSQL instance after the tests finishes. This is useful for checking the database entries or debugging after running the test cases. To run tests followed by a cleanup, run 
```
make integration-test-ci
```

To start the application, simply run:
```
make run
```

Have fun experimenting!

This demo was inspired by https://medium.com/@victorsteven/understanding-unit-and-integrationtesting-in-golang-ba60becb778d