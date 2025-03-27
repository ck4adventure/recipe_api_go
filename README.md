# Recipe App - Backend
Taking the recipe_app_next repo and copying over the functionality into a split go api backend and a separate react frontend.

This golang example will attempt to embody development best practices such as:
- Test Driven Development
- Proper logging, security, and error handling
- Run in any env, containerized for K8s

## Getting started
- download this repo
- ensure go 1.24 installed
- from root dir, `go run main.go`

## Current Routes

- `/` Home, returns {"message":"hi"}
- `/health` HealthCheck, returns {"status":"ok"}


## Eventual Directory Organization
(as suggested by chatgpt)

```
/yourproject
│── /cmd              # Entry points (e.g., different binaries)
│── /config           # Configuration files (optional)
│── /internal         # Internal application logic (not meant to be used outside this project)
│   │── /handlers     # HTTP handlers (route logic)
│   │── /models       # Database models
│   │── /services     # Business logic
│   │── /middleware   # Middleware (e.g., logging, auth)
│── /pkg             # Public reusable code (optional)
│── /api             # API-related docs (e.g., OpenAPI/Swagger)
│── /scripts         # Scripts for automation
│── /test            # Test files
│── main.go          # Application entry point
│── go.mod           # Go module definition
│── go.sum           # Dependency checksums
│── Dockerfile       # Docker setup (optional)
│── .env             # Environment variables
```
