IMAGE_VERSION := 1.x

.PHONY: docker-build
docker-build:
	docker build -f env/Dockerfile \
	             -t kamilsk/passport:$(IMAGE_VERSION) \
	             -t kamilsk/passport:latest \
	             --force-rm --no-cache --pull --rm \
	             .

.PHONY: docker-push
docker-push:
	docker push kamilsk/passport:$(IMAGE_VERSION)
	docker push kamilsk/passport:latest

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
	           --publish 8080:8080 \
	           --publish 8090:8090 \
	           --publish 8091:8091 \
	           kamilsk/passport:$(IMAGE_VERSION)

.PHONY: docker-logs
docker-logs:
	docker logs -f passport-dev

.PHONY: docker-stop
docker-stop:
	docker stop passport-dev
