# Build

FROM golang:alpine AS build

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

WORKDIR /source

COPY . .

RUN go build -a -o ./service

# Final

FROM alpine

WORKDIR /

COPY --from=build /source/service /usr/local/bin/service

ENV PATH="/usr/local/bin:${PATH}"

ENTRYPOINT ["service"]
