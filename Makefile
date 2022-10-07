all: build

build:
	docker build -t mheers/kubeyaml:latest .

push:
	docker push mheers/kubeyaml:latest
