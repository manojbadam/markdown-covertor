FROM golang:1.10

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/manojbadam/markdown-convertor

RUN go get github.com/Sirupsen/logrus && \
    go get github.com/spf13/cobra && \
    go install github.com/manojbadam/markdown-convertor

WORKDIR /go/bin