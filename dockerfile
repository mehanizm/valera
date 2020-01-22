############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser.
RUN adduser -D -g '' valerauser

WORKDIR $GOPATH/src/github.com/mehanizm/valera
COPY . .

# Fetch dependencies.
# Using go mod.
RUN go mod download
RUN go mod verify

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/valera .


############################
# STEP 2 build a small image
############################
FROM scratch

# Import the user and group files from the builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /etc/passwd /etc/passwd

# Copy our static executable.
COPY --from=builder /go/bin/valera /go/bin/valera

# Use an unprivileged user.
USER valerauser

# Run the valera binary.
ENTRYPOINT ["/go/bin/valera"]