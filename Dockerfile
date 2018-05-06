FROM alpine:3.6

ENV GOPATH /go
ENV LEGO_VERSION nicky-fork

RUN apk update && apk add --no-cache --virtual run-dependencies ca-certificates && \
    apk add --no-cache --virtual build-dependencies go git musl-dev

RUN go get -u github.com/nicky-dev/lego

RUN cd ${GOPATH}/src/github.com/nicky-dev/lego && \
    git checkout ${LEGO_VERSION} && \
    go build -o /usr/bin/lego . && \
    apk del build-dependencies && \
    rm -rf ${GOPATH}

ENTRYPOINT [ "/usr/bin/lego" ]
