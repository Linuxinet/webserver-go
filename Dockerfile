FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build ./...

EXPOSE 4040

CMD ["/app/web"]
