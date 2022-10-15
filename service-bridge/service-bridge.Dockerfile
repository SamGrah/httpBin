# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY serviceBridgeApp /app

CMD ["/app/serviceBridgeApp"]