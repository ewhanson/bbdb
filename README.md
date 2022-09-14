## Required env files and values

For UI submodule, a `.env.production` file is required, and should contain the following. These values are used at compile time.
```dotenv
VITE_POCKETBASE_URL=https://your-url-here.com
VITE_VIEWER_USER=user@your-url-here.com
VITE_VERSION=0.1.0
```

On the server, an `app.env` file will be required for `bbdb` executable config, and should contain the following. These values are used at run time.
```dotenv
NOTIFICATION_TIME=08:00
SMTP_HOST=smtp.your-url-here.com
SMTP_PORT=587
SMTP_USERNAME=username
SMTP_PASSWORD=password
SMTP_SENDER=noreply@your-url-here.com
LOCAL_LOCATION=America/Vancouver
API_VERSION=0.1.0
```


## Build Instructions

Build for Raspberry Pi

Build for local dev
```bash
git submodule update --init
cd ui && npm run build
cd .. && go build -o ./build/bbdb bbdb.go
```

```bash
git submodule update --init
cd ui && npm run build
cd .. && env GOOS=linux GOARCH=arm GOARM=7 go build -o ./build/bbdb-linux-arm-7 bbdb.go 
```

Build for server
```bash
git submodule update --init
cd ui && npm run build
cd .. && env GOOS=linux GOARCH=amd64 go build -o ./build/bbdb-linux-amd64 bbdb.go
```