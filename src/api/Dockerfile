FROM golang:1.15-alpine3.12 AS compiler
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /builder

ADD go.mod go.sum ./
RUN go mod download

ADD . ./
RUN CGO_ENABLED=0 go build main.go

FROM scratch
COPY --from=compiler /builder/main /main
ENTRYPOINT ["/main"]
