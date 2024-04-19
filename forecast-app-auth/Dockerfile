FROM golang:latest as build

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd

FROM alpine:latest as run

WORKDIR /app

COPY --from=build /src/app .

CMD ./app -http ${HTTP} -database-address ${DATABASE_ADDRESS} -database-name ${DATABASE_NAME}
