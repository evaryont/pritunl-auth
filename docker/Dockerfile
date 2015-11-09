FROM pritunl/archlinux:2015-10-24
MAINTAINER Pritunl <contact@pritunl.com>

RUN pacman -S --noconfirm go git bzr xmlsec

ENV GOPATH /go
ENV PATH $PATH:/go/bin

RUN go get -u github.com/pritunl/pritunl-auth # 78fe1348d3a0606b94d94849e61c5b998180d94d

CMD ["pritunl-auth", "app"]
