build:
	@docker build -t laurkan/otus-docker:v0.0.1 .
run:
	@docker run --name otus-docker -p 8000:8000 -d laurkan/otus-docker:v0.0.1