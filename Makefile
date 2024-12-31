VERSION ?= latest

.PHONY: build-all push-all deploy-all delete-all logs-all

# Build Docker images
build-todo:
	docker build ./todo/ -t sathirak/todo:${VERSION}

build-ping-pong:
	docker build ./ping-pong/ -t sathirak/ping-pong:${VERSION}

build-log-reader:
	docker build ./log-output/reader -t sathirak/log-output-reader:${VERSION}

build-log-writer:
	docker build ./log-output/writer -t sathirak/log-output-writer:${VERSION}

build-all: build-ping-pong build-todo build-log-reader build-log-writer

# Push Docker images
push-todo: build-todo
	docker push sathirak/todo:${VERSION}

push-ping-pong: build-ping-pong
	docker push sathirak/ping-pong:${VERSION}

push-log-reader: build-log-reader
	docker push sathirak/log-output-reader:${VERSION}

push-log-writer: build-log-writer
	docker push sathirak/log-output-writer:${VERSION}

push-all: push-ping-pong push-log-reader push-log-writer

# Kubernetes deployments

deploy-todo:
	kubectl apply -f ./todo/mainfests/

deploy-ping-pong:
	kubectl apply -f ./ping-pong/mainfests/

deploy-log-output:
	kubectl apply -f ./log-output/mainfests/

deploy-all: deploy-ping-pong deploy-todo deploy-log-output

# Delete deployments

delete-todo:
	kubectl delete -f ./todo/mainfests/

delete-ping-pong:
	kubectl delete -f ./ping-pong/mainfests/

delete-log-output:
	kubectl delete -f ./log-output/mainfests/

delete-all: delete-ping-pong delete-todo delete-log-output

# View logs
logs-todo:
	kubectl logs -l app=todo --tail=100 -f

logs-ping-pong:
	kubectl logs -l app=ping-pong --tail=100 -f

logs-log-output:
	kubectl logs -l app=log-output --tail=100 -f

logs-all:
	kubectl logs -l "app in (ping-pong,todo,log-output)" --tail=100 -f

# List images
images:
	docker images | grep sathirak