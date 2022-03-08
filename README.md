# Wallet Service

A microservice that exposes functionalities for managing a user's wallet.

## Running the application

Make sure you have [Git](https://git-scm.com/downloads) and [Docker](https://docs.docker.com/get-docker/) and [Golang](https://go.dev/doc/install) installed locally.

Also make sure nothing is running on ports **7070**. Or you can change the ports on in the **.env** file.

#### Clone the project locally

```bash
git clone https://github.com/charlesonunze/wallet-service.git && cd wallet-service
```

#### Install packages

```bash
go get -u ./...
```

#### Running the database

```bash
make db
```

#### Running the microservice

```bash
make run
```

#### Running the tests

```bash
make test
```

## Possible Improvements

#### Tests

I did not add tests for the GRPC services. This is bad and should never make it to prod.
