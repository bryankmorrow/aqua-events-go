############################
# STEP 1 build executable binary
############################
FROM golang:rc-alpine AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Copy code

RUN mkdir -p $GOPATH/src/github.com/BryanKMorrow/aqua-events-go
ADD . $GOPATH/src/github.com/BryanKMorrow/aqua-events-go
WORKDIR $GOPATH/src/github.com/BryanKMorrow/aqua-events-go

# Fetch dependencies.
# Using go get.
RUN go get "github.com/gorilla/mux"; go get "github.com/gorilla/handlers"

# Build the binary.
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o $GOPATH/src/github.com/BryanKMorrow/aqua-events-go/aqua-events-go cmd/aqua-events-go/main.go

############################
# STEP 2 build a small image
############################
FROM scratch
# Import from the builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable.
COPY --from=builder /go/src/github.com/BryanKMorrow/aqua-events-go/aqua-events-go /go/src/github.com/BryanKMorrow/aqua-events-go/aqua-events-go

# Run the binary.
ENTRYPOINT ["/go/src/github.com/BryanKMorrow/aqua-events-go/aqua-events-go"]
