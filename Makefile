build:
	go build -o bin/kiddy_line ./main.go

run:
	docker-compose up

tests:
	go test

stop:
	docker-compose stop


restart_client:
	docker-compose up -d --no-deps --build client

