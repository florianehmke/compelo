.PHONY: frontend compelo

TAG?=$(shell git describe --tags)
export TAG

GOCMD := go
EXECUTEABLE := compelo
FRONTEND_PATH := frontend/compelo

all: build

vet:
	$(GOCMD) vet ./...

format:
	$(GOCMD) fmt ./...

test:
	$(GOCMD) test ./...

quality: format vet test

frontend-prepare:
	cd $(FRONTEND_PATH) && npm install

frontend-build:
	cd $(FRONTEND_PATH) && ng build --prod --base-href /app/

frontend: frontend-prepare frontend-build
	$(GOCMD) generate ./frontend

build: frontend
	$(GOCMD) generate ./db
	$(GOCMD) build -o $(EXECUTEABLE) ./cmd/compelo

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
