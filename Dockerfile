#FROM golang:alpine
#
#WORKDIR /
#ENV PORT=8000
#CMD["go","run","main.go"]

#FROM golang:alpine
#ADD ./src /go/src/app
#WORKDIR /go/src/app
#ENV PORT=3001
##CMD["go","build", "./src/main.go"]
#CMD ["go", "run", "main.go"]

FROM golang
ENV GO111MODULE=on
#ADD ./httpserver /go/src/app
WORKDIR /go/src/app
#WORKDIR /app
# COPY ./go.mod .
# COPY ./go.sum .
RUN go mod download
COPY . .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
ENV PORT=3001
EXPOSE 3001
#ENTRYPOINT ["/app/httpserver"]
ENTRYPOINT ["go","run","main.go"]
#CMD ["go","run","main.go"]