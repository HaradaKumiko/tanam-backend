FROM golang:1.22-alpine AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.

RUN apk update && \
    apk add --no-cache git ca-certificates tzdata && update-ca-certificates

WORKDIR $GOPATH/src/compile/tanam-backend
COPY . .

# Fetch dependencies.
# RUN go get -d -v
RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/tanam-backend .

#####################
# MAKE SMALL BINARY #
#####################
FROM alpine:3.16

RUN apk update

# COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

#Copy the executable.
COPY --from=builder /go/bin/tanam-backend /go/bin/tanam-backend


ENTRYPOINT ["/go/bin/tanam-backend"]

EXPOSE 1323

#CMD ["./main"]