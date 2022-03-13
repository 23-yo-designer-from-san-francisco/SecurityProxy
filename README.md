## Simple HTTP proxy server

# Build
```bash
docker build -t proxy .
```

# Run
```bash
docker run -dp 8080:8080 -p 80:80 proxy
```

# Status
- [x] HTTP Proxy
- [x] HTTPS Proxy
- [x] Query storage
- [x] Vulnerability scanner

# Vulnerability scanner
- GET params
- POST (`application/x-www-form-urlencoded`)
- Cookie (`Cookie: <key>=<value>; <key>=<value>; ...` only)

Example output
```
Found vulnerable query param {test} with exploit {;cat /etc/passwd;}
Found vulnerable query param {test} with exploit {|cat /etc/passwd|}
Found vulnerable query param {test} with exploit {`cat /etc/passwd`}
Found vulnerable cookie {internet} with exploit {;cat /etc/passwd;}
Found vulnerable cookie {internet} with exploit {|cat /etc/passwd|}
Found vulnerable cookie {internet} with exploit {`cat /etc/passwd`}
```
