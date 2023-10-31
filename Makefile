all: go

go:
	go run -v main.go api

admin:
	go run main.go admin

test_go:
	go run main.go api --config ./etc/configs_sqlite.json

test_admin:
	go run main.go admin --config ./etc/configs_sqlite.json

buildx:
	go build -o mycube main.go


install:
	go mod tidy
	go mod vendor
	go build -o -v mycube main.go
	systemctl stop mycube.service
	cp mycube /usr/local/bin/mycube
	systemctl restart mycube.service
