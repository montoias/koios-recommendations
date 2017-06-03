FROM golang:1.8.0-alpine

RUN ln -s /go/src/github.com/montoias/koios-recommendations/development/parameters /root/.uniplaces

# install deps
RUN apk add --no-cache \
    python \
    python-dev \
    py-pip \
    build-base \
    git \
    openssh-client \
    vim \
    curl
WORKDIR /go/src/github.com/montoias/koios-recommendations
RUN go get github.com/golang/lint/golint
RUN go get golang.org/x/tools/cmd/goimports
RUN go get github.com/vektra/mockery/.../
RUN go get github.com/golang/dep/cmd/dep
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"
CMD ["/bin/sh"]
