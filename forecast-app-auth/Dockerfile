FROM golang:1.22.2-alpine3.19 as build

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd

FROM alpine:3.19 as run

WORKDIR /app

COPY --from=build /src/app .

RUN adduser --disabled-password newuser

RUN chown newuser:newuser ./app

USER newuser

CMD ./app -http ${HTTP} -database-address ${DATABASE_ADDRESS} -database-name ${DATABASE_NAME}

