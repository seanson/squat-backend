FROM golang:alpine as builder

WORKDIR /go/src/app
COPY . /go/src/app/

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -a \
             -installsuffix cgo \
             -ldflags="-w -s" \
             -o /go/bin/squat-backend

FROM scratch

ENV GIN_MODE=release

COPY --from=builder /go/bin/squat-backend /squat-backend
COPY --from=builder /tmp /tmp

ENTRYPOINT ["/squat-backend"]