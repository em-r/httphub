FROM golang:1.16-alpine

RUN apk update && apk add --no-cache git

ENV home=/home/src/app

RUN mkdir -p ${home}
WORKDIR ${home}
ADD . .

# install dependencies
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

# build bin
RUN go build -o httphub.out httphub/main.go

RUN apk add libcap && setcap 'cap_net_bind_service=+ep' httphub.out

EXPOSE 80

# run server
CMD CompileDaemon --build="go build -o httphub.out httphub/main.go" --command=./httphub.out
