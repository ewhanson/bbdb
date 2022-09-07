# Build Instructions

Build for Raspberry Pi

```bash
env GOOS=linux GOARCH=arm GOARM=7 go build -o bbdb-linux-arm-7 bbdb.go 
```

Build for server
```bash
env GOOS=linux GOARCH=amd64 go build -o bbdb-linux-amd64 bbdb.go
```