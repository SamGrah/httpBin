# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY binManagerApp /app

CMD ["/app/binManagerApp"]