FROM golang:latest as build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o app ./cmd

FROM alpine:latest as run

ENV HTTP _
ENV AUTH-SERVICE-ADDRESS _

WORKDIR /app

COPY --from=build /src/app .

CMD ./app -http ${HTTP} -auth-service-address ${AUTH-SERVICE-ADDRESS}
