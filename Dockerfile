FROM alpine

RUN apk add ca-certificates

ENTRYPOINT ["/kudu-shouter"]
COPY kudu-shouter /