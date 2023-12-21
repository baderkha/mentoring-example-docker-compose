
build:
	docker build -t go_example_webserver:latest .

start: build
	docker-compose up
