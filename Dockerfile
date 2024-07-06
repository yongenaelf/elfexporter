FROM golang

ADD . /go/src/github.com/yongenaelf/elfexporter
RUN cd /go/src/github.com/yongenaelf/elfexporter && go get
RUN go install github.com/yongenaelf/elfexporter

ENV GETH https://mainnet.infura.io
ENV PORT 9015

RUN mkdir /app
WORKDIR /app
ADD addresses.txt /app

EXPOSE 9015

ENTRYPOINT /go/bin/elfexporter
