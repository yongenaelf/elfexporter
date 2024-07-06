FROM golang

ADD . /go/src/github.com/yongenaelf/elfexporter
RUN cd /go/src/github.com/yongenaelf/elfexporter && go get
RUN go install github.com/yongenaelf/elfexporter

ENV AELF_URL https://tdvw-test-node.aelf.io
ENV PORT 8080

RUN mkdir /app
WORKDIR /app
ADD addresses.txt /app

EXPOSE 8080

ENTRYPOINT /go/bin/elfexporter
