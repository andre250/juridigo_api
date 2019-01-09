FROM golang:1.8 as builder
ENV ENV=Staging

LABEL MAINTAINER="GuilhermeCaruso"
LABEL COMPANY="Mica"

WORKDIR /go/src/gitlab.com/mica/backend/user

COPY . ./

RUN apt-get update
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build

FROM golang:1.8
WORKDIR /go/src/gitlab.com/mica/backend/user
COPY --from=builder /go/src/gitlab.com/mica/backend/user/user .
CMD ./user
EXPOSE 3030