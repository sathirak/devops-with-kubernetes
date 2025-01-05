VERSION ?= latest

build:
	docker build ./ -t sathirak/postgres-backup:${VERSION}

push:
	docker push sathirak/postgres-backup:${VERSION}

deploy:
	kubectl apply -f ./manifests/