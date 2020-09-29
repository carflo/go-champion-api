build:
	docker build -t go-champion-api .

run: build
	docker run -it -p 80:8080 --name champion-api-server go-champion-api
