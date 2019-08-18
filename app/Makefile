init/dotenv:
	go run cmd/deg/main.go

migrate/new:
	go run cmd/migw/main.go new $(NAME)

migrate/up:
	go run cmd/migw/main.go up

migrate/down:
	go run cmd/migw/main.go down

docker-compose/up:
	make -C ../database docker-compose/up

docker-compose/down:
	make -C ../database docker-compose/down

mysql/create-db:
	go run cmd/mysqlw/main.go create-db

mysql/drop-db:
	go run cmd/mysqlw/main.go drop-db

mysql/get-cli-cmd:
	go run cmd/mysqlw/main.go get-cli-cmd
