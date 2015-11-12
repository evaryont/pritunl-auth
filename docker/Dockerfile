FROM pritunl/archlinux:2015-10-24
MAINTAINER Pritunl <contact@pritunl.com>

RUN pacman -S --noconfirm go git bzr xmlsec

ENV GOPATH /go
ENV PATH $PATH:/go/bin

RUN go get -u github.com/pritunl/pritunl-auth # e2505c518c20b493dec273a8df507d4070c91db5

CMD ["pritunl-auth", "app"]
