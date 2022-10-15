.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t pocket-bot-image:v0.1 .

run-conteiner:
	docker run --name pocket-bot -p 80:80 --env-file .env pocket-bot-image:v0.1

start-contener:
	docker start pocket-bot
