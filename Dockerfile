FROM daocloud.io/golang:1.8.1-onbuild
RUN go get github.com/Focinfi/sqs-master/
RUN CGO_ENABLED=0 go install -a github.com/Focinfi/sqs-master/
WORKDIR $GOPATH/src/github.com/Focinfi/sqs-master/
EXPOSE 8080
CMD ["sqs-master"]
