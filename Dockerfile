FROM alpine
WORKDIR /web/gin
COPY gin .
COPY ./Config .
CMD ./gin
