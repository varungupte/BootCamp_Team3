FROM golang:1.14
WORKDIR /usr/go-gin-server
COPY . /usr/go-gin-server
RUN go build -o gin_server ./cmd/gin_server
EXPOSE 7878
#CMD ["go", "run", "cmd/gin_server/main.go"]
CMD ["/usr/go-gin-server/gin_server"]