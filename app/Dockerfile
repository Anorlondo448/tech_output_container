FROM golang
ADD . /go/src/
RUN /usr/local/go/bin/go build /go/src/main.go
EXPOSE 8081
CMD ["./main"]