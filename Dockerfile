FROM golang:1.26.1-alpine AS builder 
RUN apk add --no-cache git ca-certificates 
WORKDIR /app
COPY go.mod go.sum ./ 
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -o manager ./cmd/server 


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/manager .

CMD ["./manager"]
