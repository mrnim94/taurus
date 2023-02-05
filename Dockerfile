FROM golang:1.20.0-alpine

RUN apk update && apk add git && apk add unzip && apk add curl

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
RUN unzip awscliv2.zip
RUN ./aws/install

ENV CGO_ENABLED=0
ENV GO111MODULE=on

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/taurus

COPY . .

RUN go mod download
RUN GOOS=linux go build -o app
ENTRYPOINT ["./app"]

EXPOSE 6969