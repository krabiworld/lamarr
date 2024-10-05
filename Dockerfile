FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -C cmd/module -o /app/module

FROM gcr.io/distroless/static-debian12
COPY --from=build /app/module .
CMD ["/module"]
