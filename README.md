# HTTPHub

An HTTP Request/Response service fully written in Go.

## Run Locally

Pull from Docker Hub

```bash
  docker run -p 80:80 iammehdi/httphub
```

## Environment Variables

To run this project in development mode, you will need to add DEV_MODE to your .env file or to set it from the CLI

```bash
  export DEV_MODE=1
```

## Build process

To run tests, run the following command

```bash
  make test
```

To generate Swagger specs, run the following command

```bash
  make swagger
```

## Features

Currently, HTTPHub covers the following topics:

- Authentication
- Caching
- Cookies
- HTTP methods
- Redirections
- Request/Response inspection
- Status codes

## Acknowledgements

- [This project is inspired from Kenneth Reitz' httpbin](https://httpbin.org)
- [Special thanks to MDN for their amazing docs](https://developer.mozilla.org/)
- [README created using readme.so](https://readme.so)
