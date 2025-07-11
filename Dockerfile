FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" go build -C cmd/lamarr -o /app/lamarr

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/lamarr .
CMD ["/lamarr"]
