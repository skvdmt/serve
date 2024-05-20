# install mods
install-go-mods:
	go get -u github.com/gin-gonic/gin
	go get -u github.com/labstack/echo/v4

# start development
start:
	go run main.go

# building application
build:
	go build -o ./bin/serve ./main.go
