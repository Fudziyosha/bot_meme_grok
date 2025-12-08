FROM golang:1.25.5 as builder
ARG CGO_ENABLED=0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o ./app ./cmd/main.go


FROM alpine:latest
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]