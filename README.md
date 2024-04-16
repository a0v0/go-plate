# Server

## Overview

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

- Tests are written at service level.

## Instructions

- Pass the environment variables in the `.env` file at the root dir or in the environment. See [`.env.example`](.env.example) for reference.

# Reference

- https://github.com/efectn/fiber-boilerplate
