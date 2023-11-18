# Build stage

FROM golang:1.21.4-alpine3.18 AS builder

ENV GOOS=linux
ENV CGO_ENABLED=0

WORKDIR /app
COPY . .
RUN apk add --no-cache make
RUN make prepare
RUN make all

# Final stage

FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/bin/application .
ENV PATH="/app:${PATH}"

ENTRYPOINT ["application"]
