FROM pritunl/archlinux:2015-05-02
MAINTAINER Pritunl <contact@pritunl.com>

RUN pacman -S --noconfirm go git bzr

ENV GOPATH /go
ENV PATH $PATH:/go/bin

RUN go get -u github.com/pritunl/pritunl-auth # 909523d8837903b66961677d92eeec178b2ce3b9

CMD ["pritunl-auth", "app"]
