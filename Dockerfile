# Use Node.js image to build the UI dependencies
FROM node:18 as ui-builder
WORKDIR /app/ui
COPY ui/package*.json ./
RUN npm ci
COPY ui/ .
RUN npm run build
# Use NPM to get the version from package.json
RUN echo "$(node -p "require('./package.json').version")" > ./version.txt

# Use Go image to build the Go binary
FROM golang:1.20 as go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=ui-builder /app/ui/dist ./ui/dist
COPY --from=ui-builder /app/ui/version.txt ./version.txt
RUN export VERSION=$(echo "$(cat version.txt)")
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -X github.com/ewhanson/bbdb/hooks.Version=$VERSION" -o bbdb bbdb.go

# Use Alpine image to run the application
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=go-builder /app/bbdb ./
EXPOSE 8090
ENTRYPOINT ["./bbdb"]
CMD ["serve", "--http=0.0.0.0:8090"]