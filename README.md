# Recipe App - Backend
Taking the recipe_app_next repo and copying over the functionality into a split go api backend and a separate react frontend.

## Getting started
- download this repo
- ensure go 1.24 installed
- from root dir, `go run main.go`

## Current Status

```bash
➜  recipe_api_go git:(main) ✗ curl localhost:8080 -v
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/8.7.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 418 I'm a teapot
< Content-Type: application/json
< Date: Tue, 25 Mar 2025 01:44:04 GMT
< Content-Length: 29
< 
{"success":true,"data":"hi"}
* Connection #0 to host localhost left intact
```

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
