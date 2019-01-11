FROM golang:1.8 as builder

LABEL MAINTAINER="GuilhermeCaruso"
LABEL COMPANY="Juridigo"

WORKDIR /go/src/github.com/juridigo/juridigo_api_usuario

COPY . ./

RUN apt-get update
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN /bin/bash -c "source .env"
RUN go build

FROM golang:1.8
WORKDIR /go/src/github.com/juridigo/juridigo_api_usuario
COPY --from=builder /go/src/github.com/juridigo/juridigo_api_usuario/juridigo_api_usuario .
COPY --from=builder /go/src/github.com/juridigo/juridigo_api_usuario/.env .

CMD /bin/bash -c "source .env && ./juridigo_api_usuario"
EXPOSE 3030