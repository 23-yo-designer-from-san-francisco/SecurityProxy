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
- [ ] Vulnerability scanner