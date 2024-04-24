# Build stage

FROM golang:1.22.2-alpine3.19 AS builder

ENV GOOS=linux
ENV CGO_ENABLED=0

WORKDIR /app
COPY . .
RUN apk add --no-cache make
RUN make prepare
RUN make all

# Final stage

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/bin/application .
ENV PATH="/app:${PATH}"

ENTRYPOINT ["application"]
