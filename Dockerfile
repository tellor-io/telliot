FROM golang:buster AS builder
WORKDIR /go/src
RUN apt-get update
COPY ./ .
RUN make build

FROM debian:buster-slim  
WORKDIR /
RUN apt-get update
RUN apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /go/src/telliot .
ENTRYPOINT ["./telliot"]