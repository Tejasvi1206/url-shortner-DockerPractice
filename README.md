# Go URL Shortener

A lightweight URL shortener service built with Go that exposes a REST API for
creating short URLs, redirecting to original URLs, and viewing usage metrics.

## Features

- Shorten long URLs into unique short hashes
- Return the same short URL when the same URL is submitted again
- Redirect short URLs to their original URLs
- In-memory URL storage using a Go map
- Metrics API for identifying the top 3 most frequently shortened domains
- Dockerized application

## How It Works

The service accepts a long URL through the REST API and generates a deterministic
short hash for it.

If the same URL is submitted multiple times, the service returns the previously
generated short URL instead of creating a new one.

When a short URL is accessed, the service resolves the hash and redirects the
request to the original URL.

## Architecture

```text
                ┌──────────────────┐
                │     REST API     │
                └────────┬─────────┘
                         │
             ┌───────────┴───────────┐
             │                       │
        Shorten URL              Redirect URL
             │                       │
             └───────────┬───────────┘
                         │
                  ┌──────▼──────┐
                  │  Go Map     │
                  │ In-Memory DB │
                  └──────┬──────┘
                         │
                  ┌──────▼──────┐
                  │   Metrics   │
                  │  Top Domains│
                  └─────────────┘
```

## Tech Stack
```text
- Language: Go
- API: REST
- Storage: In-memory Go map
- Containerization: Docker
```
## Project Structure
```text
├── Dockerfile
├── README.md
├── go.mod
└── main.go
```
## API Capabilities
- URL Shortening
Accepts a long URL and returns its corresponding shortened URL.
- URL Redirection
Accepts a short URL and redirects the request to the original URL.
- Metrics
Provides usage statistics for the most frequently shortened domains.

## Key Design Decisions
- Deterministic URL Generation
The service generates a deterministic short identifier for a URL. This allows
repeated requests for the same URL to return the same shortened URL.
- In-Memory Storage
URL mappings are stored in a Go map, keeping the implementation lightweight
and avoiding an external database dependency.
- Containerization
The application includes a Dockerfile for running the service in a containerized
environment.

## Future Improvements
- Persistent database storage
- URL expiration
- Custom aliases
- Authentication and authorization
- More detailed analytics
- Concurrent request handling improvements

## Author
Tejasvi Singh
GitHub: Tejasvi1206
