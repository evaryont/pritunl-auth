FROM pritunl/archlinux:2015-05-02
MAINTAINER Pritunl <contact@pritunl.com>

RUN pacman -S --noconfirm go git bzr

ENV GOPATH /go
ENV PATH $PATH:/go/bin

RUN go get -u github.com/pritunl/pritunl-auth # da8aef5ae8061ab63a8c23567f419be8826e6b63

CMD ["pritunl-auth", "app"]
