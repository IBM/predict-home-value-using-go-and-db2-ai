FROM golang:1.10.3-alpine

# Update packages and install dependency packages for services
RUN apk update && apk add --no-cache bash git

# Change working directory
WORKDIR $GOPATH/src/db2wml/

# Install dependencies
RUN go get -u github.com/golang/dep/...
RUN go get -u github.com/derekparker/delve/cmd/dlv/...
COPY . ./
RUN dep ensure -v

ENV PORT 8080
ENV GIN_MODE release
EXPOSE 8080

CMD ["go", "run", "server.go"]
