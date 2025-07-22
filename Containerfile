# Build

FROM golang:1.24.5 AS build

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

WORKDIR /source

COPY . .

RUN go mod download
RUN go build -a -o ./service

# Final

FROM gcr.io/distroless/static-debian12

WORKDIR /

COPY --from=build /source/service /usr/local/bin/service

ENV PATH="/usr/local/bin:${PATH}"

ENTRYPOINT ["service"]
