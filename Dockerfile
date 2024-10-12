FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -C cmd/lamarr -o /app/lamarr

FROM gcr.io/distroless/static-debian12
COPY --from=build /app/lamarr .
CMD ["/lamarr"]
