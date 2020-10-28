FROM golang:1.15.2-buster

ENV URL=http://callhelo-service:9090

WORKDIR /src

COPY index.html /src/index.html
COPY serve /src/serve

CMD ["/src/serve"]

EXPOSE 8080
