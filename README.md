# Recipe App - Backend
Taking the recipe_app_next repo and copying over the functionality into a split go api backend and a separate react frontend.

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