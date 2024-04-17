# go_plate

Yet another go boilerplate. Simple and scalable boilerplate to build powerful and organized REST projects with Fiber.

## Prerequisites

- [Go](https://golang.org/dl/) (1.22+)
- [Taskfile](https://taskfile.dev/#/installation)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Usage

1. Rename the `.env.example` to `.env` and update the values.
1. Install the dependencies

   ```bash
   task install
   ```

1. To generate go code

   ```bash
   task generate
   ```

1. Start the server
   ```bash
   task start
   ```

## Instructions

- Tests are written at service
- Pass the environment variables in the `.env` file at the root dir or in the environment. See [`.env.example`](.env.example) for reference.
- `Service Layer`: The service layer is the core of the server. This layer communicates with the business layer.
  - Handles all the requests from the client with minamal validation.
  - Handles authentication.
- `Business Layer`: The business layer is responsible for handling all the business logic. This layer communicates with the Database layer.
  - Handles all the business logic.
  - Handles authorization.
  - Handles validation.
  - Filtering of data.
- `Database Layer`: It is the layer that communicates with the database.
  - Handles all the database operations.
  - DUMB layer.

`Service Layer` -> `Business Layer` -> `Database Layer`

# Reference

- https://github.com/efectn/fiber-boilerplate
