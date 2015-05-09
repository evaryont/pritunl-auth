FROM pritunl/archlinux:2015-05-02
MAINTAINER Pritunl <contact@pritunl.com>

RUN pacman -S --noconfirm go git bzr

ENV GOPATH /go
ENV PATH $PATH:/go/bin

RUN go get -u github.com/pritunl/pritunl-auth # eac25c54e84b83e46d47f13285f22a3465c6566b

CMD ["pritunl-auth", "app"]
