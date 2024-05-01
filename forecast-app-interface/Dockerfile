FROM golang:1.22.2-alpine3.19 as build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o app ./cmd

FROM alpine:3.19 as run

WORKDIR /app

COPY --from=build /src/app .

RUN adduser -D user

RUN chown user:user ./app

USER user

CMD ./app -http ${HTTP} -auth-service-address ${AUTH-SERVICE-ADDRESS}

