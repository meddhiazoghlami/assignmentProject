FROM golang:latest AS build

WORKDIR /usr/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create a minimal runtime image
FROM scratch
WORKDIR /usr/src/app

COPY --from=build /usr/src/app/main /usr/src/app
COPY .env /usr/src/app/.env
EXPOSE 8080

CMD ["./main"]
