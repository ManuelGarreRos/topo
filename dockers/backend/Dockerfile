FROM golang:1.21.4

ENV DEBIAN_FRONTED noninteractive

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cosmtrek/air@latest
RUN mkdir /src && mkdir /dst && mkdir /src/pgd && mkdir /src/templates

RUN chmod -R 777 /src
WORKDIR /src
COPY . .

EXPOSE 8000 40000
