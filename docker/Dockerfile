FROM pritunl/archlinux:2015-10-24
MAINTAINER Pritunl <contact@pritunl.com>

RUN pacman -S --noconfirm go git bzr xmlsec

ENV GOPATH /go
ENV PATH $PATH:/go/bin

RUN go get -u github.com/pritunl/pritunl-auth # 4fd29a8d645fea07ab1efe69eaf4fefe148eb10c

CMD ["pritunl-auth", "app"]
