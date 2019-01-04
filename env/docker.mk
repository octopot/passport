IMAGE_VERSION := 1.x
PACKAGE       := github.com/kamilsk/passport


.PHONY: docker-build
docker-build:
	docker build -f env/Dockerfile \
	             -t kamilsk/passport:$(IMAGE_VERSION) \
	             -t kamilsk/passport:latest \
	             -t quay.io/kamilsk/passport:$(IMAGE_VERSION) \
	             -t quay.io/kamilsk/passport:latest \
	             --build-arg PACKAGE=$(PACKAGE) \
	             --force-rm --no-cache --pull --rm \
	             .

.PHONY: docker-push
docker-push:
	docker push kamilsk/passport:$(IMAGE_VERSION)
	docker push kamilsk/passport:latest
	docker push quay.io/kamilsk/passport:$(IMAGE_VERSION)
	docker push quay.io/kamilsk/passport:latest

.PHONY: docker-refresh
docker-refresh:
	docker images --all \
	| grep '^kamilsk\/passport\s\+' \
	| awk '{print $$3}' \
	| xargs docker rmi -f &>/dev/null || true
	docker pull kamilsk/passport:$(IMAGE_VERSION)



.PHONY: publish
publish: docker-build docker-push



.PHONY: docker-start
docker-start:
	docker run --rm -d \
	           --env-file env/.env.example \
	           --name passport-dev \
	           -p 8080:8080 \
	           -p 8090:8090 \
	           -p 8091:8091 \
	           -p 8092:8092 \
	           -p 8093:8093 \
	           kamilsk/passport:$(IMAGE_VERSION)

.PHONY: docker-logs
docker-logs:
	docker logs -f passport-dev

.PHONY: docker-stop
docker-stop:
	docker stop passport-dev
