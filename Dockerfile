# Build

FROM golang:1.22.5-alpine AS build

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

WORKDIR /opt

COPY . .

RUN go build -a -o ./service

# Final

FROM alpine

WORKDIR /opt

COPY --from=build /opt/bin/service .

ENV PATH="/opt:${PATH}"

ENTRYPOINT ["service"]
