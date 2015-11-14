FROM pritunl/archlinux:2015-10-24
MAINTAINER Pritunl <contact@pritunl.com>

RUN pacman -S --noconfirm go git bzr xmlsec

ENV GOPATH /go
ENV PATH $PATH:/go/bin

ADD . /go/src/github.com/pritunl/pritunl-auth

RUN go get github.com/pritunl/pritunl-auth
RUN go install github.com/pritunl/pritunl-auth

CMD ["pritunl-auth", "app"]
