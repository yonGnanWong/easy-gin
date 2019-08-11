FROM alpine
WORKDIR /web/gin
COPY gin .
CMD ./gin
