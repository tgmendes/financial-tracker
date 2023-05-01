FROM arigaio/atlas:latest-alpine as atlas

# build stage
FROM golang:1.20-bullseye as build

WORKDIR /build

COPY go.mod go.sum /build/
RUN go mod download

COPY ./pkg /build/pkg
COPY cmd/collector/main.go /build/collector/main.go
COPY cmd/admin/main.go /build/admin/main.go
COPY cmd/cli/main.go /build/cli/main.go

RUN CGO_ENABLED=0 go build -o fin-collector collector/main.go
RUN CGO_ENABLED=0 go build -o fin-admin admin/main.go
RUN CGO_ENABLED=0 go build -o fin-cli cli/main.go

# runtime stage
FROM alpine

# Latest releases available at https://github.com/aptible/supercronic/releases
ENV SUPERCRONIC_URL=https://github.com/aptible/supercronic/releases/download/v0.2.24/supercronic-linux-amd64 \
    SUPERCRONIC=supercronic-linux-amd64 \
    SUPERCRONIC_SHA1SUM=6817299e04457e5d6ec4809c72ee13a43e95ba41

RUN apk add --update --no-cache ca-certificates curl \
     && curl -fsSLO "$SUPERCRONIC_URL" \
     && echo "${SUPERCRONIC_SHA1SUM}  ${SUPERCRONIC}" | sha1sum -c - \
     && chmod +x "$SUPERCRONIC" \
     && mv "$SUPERCRONIC" "/usr/local/bin/${SUPERCRONIC}" \
     && ln -s "/usr/local/bin/${SUPERCRONIC}" /usr/local/bin/supercronic

COPY --from=atlas /atlas /atlas
COPY --from=build /build/fin-collector /fin-collector
COPY --from=build /build/fin-admin /fin-admin
COPY --from=build /build/fin-cli /fin-cli

COPY db /db
COPY crontab ./crontab
COPY atlas.hcl ./atlas.hcl

CMD ["/fin-admin"]