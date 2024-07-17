FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o consul-node-manager

FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/consul-node-manager .
ENTRYPOINT ["./consul-node-manager"]
CMD ["-node", "", "-address", "http://localhost:8500"]
