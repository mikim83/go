FROM golang:1.6
ENV GOOS=darwin 
ENV GOARCH=amd64

RUN go get -d -v github.com/hudl/fargo
