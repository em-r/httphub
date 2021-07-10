FROM golang:latest

ENV home=/home/src/app

RUN mkdir -p ${home}
WORKDIR ${home}
ADD . .

# install dependencies
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

# build bin
RUN go build -o httphub.out httphub/main.go

EXPOSE 80

# run server
CMD CompileDaemon --build="go build -o httphub.out httphub/main.go" --command=./httphub.out
