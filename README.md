## Simple HTTP proxy server

# Build
```bash
docker build -t proxy .
```

# Run
```bash
docker run -dp 8080:8080 proxy
```

# Status
- [x] HTTP Proxy
- [ ] HTTPS Proxy
- [ ] Query storage
- [ ] Vulnerability scanner