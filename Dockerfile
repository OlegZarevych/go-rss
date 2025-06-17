FROM golang:1.24

WORKDIR /app
COPY . .
RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-rss

ENTRYPOINT ["/go-rss"]