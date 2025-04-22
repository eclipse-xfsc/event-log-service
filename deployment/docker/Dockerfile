FROM golang as build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /eventlogservice

FROM alpine as runner

WORKDIR /app

COPY --from=build /eventlogservice ./

USER nonroot:nonroot

CMD ["./eventlogservice"]
