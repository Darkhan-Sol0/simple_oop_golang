APP = myServ

SOURCE = cmd/app/main.go

PACKAGE =\
				github.com/labstack/echo/v4\

all:

clean:

run:
	go run $(SOURCE)

init:
	go mod init $(APP)

get:
	go get -u $(PACKAGE)