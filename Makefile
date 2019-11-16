.PHONY: frontend compelo

TAG?=$(shell git describe --tags)
export TAG

GOCMD := go
EXECUTEABLE := compelo
FRONTEND_PATH := frontend/compelo

frontend-prepare:
	cd $(FRONTEND_PATH) && npm install

frontend-quality: frontend-prepare
	cd $(FRONTEND_PATH) && npm run format:check
	cd $(FRONTEND_PATH) && npm run lint
	cd $(FRONTEND_PATH) && npm run test:ci

frontend: frontend-prepare
	cd $(FRONTEND_PATH) && npm run build-prod

generate:
	mkdir -p $(FRONTEND_PATH)/dist
	$(GOCMD) generate ./frontend
	$(GOCMD) generate ./internal/db/scripts

backend-quality: generate
	$(GOCMD) fmt ./...
	$(GOCMD) vet ./...
	$(GOCMD) test ./...

backend: generate
	$(GOCMD) build -o $(EXECUTEABLE) ./cmd/compelo

# Builds application with dev tag, meaning that sql files and
# frontend will be read from local disk as opposed to being embedded.
backend-dev:
	$(GOCMD) build -o $(EXECUTEABLE) -tags=dev ./cmd/compelo


# Build & publish to docker hub.
# The docker build will run the above 'frontend' and 'backend' steps.
docker-build:
	docker build \
	 	-t florianehmke/compelo:latest \
	 	-t florianehmke/compelo:$(TAG) .

docker-push: build-docker
	docker push florianehmke/compelo:latest
	docker push florianehmke/compelo:$(TAG)

clean:
	rm -f $(EXECUTEABLE)

distclean: clean
	rm -f -r frontend/compelo/dist
	rm -f frontend/frontend_vfsdata.go
	rm -f internal/db/scripts/scripts_vfsdata.go
