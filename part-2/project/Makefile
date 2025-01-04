VERSION ?= latest

.PHONY: all build push deploy build-all push-all deploy-all

# Combined command to do everything
all: build-all push-all deploy-all

# Build commands
build-todo:
	docker build ./todo/ -t sathirak/todo:${VERSION}

build-todo-backend:
	docker build ./todo-backend/ -t sathirak/todo-backend:${VERSION}

build-all: build-todo build-todo-backend

# Push commands
push-todo:
	docker push sathirak/todo:${VERSION}

push-todo-backend:
	docker push sathirak/todo-backend:${VERSION}

push-all: push-todo push-todo-backend

# Deploy commands
deploy-cluster:
	kubectl apply -f ./cluster/manifests/

deploy-postgres:
	kubectl apply -f ./postgres/manifests/

deploy-todo:
	kubectl apply -f ./todo/manifests/

deploy-todo-backend:
	kubectl apply -f ./todo-backend/manifests/

deploy-all: deploy-postgres deploy-todo deploy-todo-backend

# Delete commands
delete-all:
	kubectl delete -f ./todo/manifests/
	kubectl delete -f ./todo-backend/manifests/
	kubectl delete -f ./postgres/manifests/
	kubectl delete -f ./cluster/manifests/

# Logs commands
logs-todo:
	kubectl logs -l app=todo -f

logs-backend:
	kubectl logs -l app=todo-backend -f

logs-postgres:
	kubectl logs -l app=postgresapp -f
