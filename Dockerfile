FROM golang:1.20 AS build

COPY . /build

WORKDIR /build

RUN go mod download

RUN go build -o /service-bin

FROM alpine

EXPOSE 8080

WORKDIR /app
COPY --from=build /service-bin .

CMD ["/app/service-bin"]
