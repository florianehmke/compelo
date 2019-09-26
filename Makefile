.PHONY: frontend compelo

TAG?=$(shell git describe --tags)
export TAG

GOCMD := go
EXECUTEABLE := compelo
FRONTEND_PATH := frontend/compelo

all: build

quality:
	$(GOCMD) fmt ./...
	$(GOCMD) vet -tags=dev ./...
	$(GOCMD) test -tags=dev ./...
	cd $(FRONTEND_PATH) && npm run format:check
	cd $(FRONTEND_PATH) && npm run lint

build:
	cd $(FRONTEND_PATH) && npm install
	cd $(FRONTEND_PATH) && ng build --prod --base-href /app/
	$(GOCMD) generate ./frontend
	$(GOCMD) generate ./db
	$(GOCMD) build -o $(EXECUTEABLE) ./cmd/compelo

build-dev:
	$(GOCMD) build -o $(EXECUTEABLE) -tags=dev ./cmd/compelo

build-docker:
	docker build \
	 	-t florianehmke/compelo:latest \
	 	-t florianehmke/compelo:$(TAG) .

push-docker: build-docker
	docker push florianehmke/compelo:latest
	docker push florianehmke/compelo:$(TAG)

clean:
	rm -f $(EXECUTEABLE)

distclean: clean
	rm -f -r frontend/compelo/dist
	rm -f frontend/frontend_vfsdata.go
	rm -f db/scripts_vfsdata.go
