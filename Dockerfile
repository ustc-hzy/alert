FROM golang:1.16.6-alpine

#RUN apk update && apk add git

WORKDIR /home/lan/local/alert
COPY . .

RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy
RUN go env -w GO111MODULE=on

RUN go build -o crud-server .
CMD ./crud-server

#WORKDIR /home/lan/local/alert/schedule
#RUN go build -o schedule .
#CMD ./schedule

