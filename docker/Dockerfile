FROM pritunl/archlinux:2015-10-24
MAINTAINER Pritunl <contact@pritunl.com>

RUN pacman -S --noconfirm go git bzr xmlsec

ENV GOPATH /go
ENV PATH $PATH:/go/bin

RUN go get -u github.com/pritunl/pritunl-auth # 7a2d6387c0fb2a2e0763676ff76dd63090dbb27a

CMD ["pritunl-auth", "app"]
