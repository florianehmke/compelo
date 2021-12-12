.PHONY: frontend backend all

VERSION?=$(shell git describe --always)
TAG?=$(shell git rev-parse --abbrev-ref HEAD)
DATE?=$(shell date '+%Y-%m-%d %H:%M')
export TAG
export VERSION
export DATE

GOCMD := go
EXECUTEABLE := compelo
FRONTEND_PATH := frontend/compelo

all: generate frontend backend

# Codegen
# =================

generate:
	mkdir -p $(FRONTEND_PATH)/dist
	mkdir -p $(FRONTEND_PATH)/src/generated
	$(GOCMD) generate ./internal/db/scripts
	$(GOCMD) generate ./frontend

# Frontend
# =================

frontend-prepare:
	cd $(FRONTEND_PATH) && npm install

frontend-verify: frontend-prepare
	cd $(FRONTEND_PATH) && npm run format:check
	cd $(FRONTEND_PATH) && npm run lint
	cd $(FRONTEND_PATH) && npm run test:ci

frontend: frontend-verify
	cd $(FRONTEND_PATH) && echo "export const APP_VERSION = '$(VERSION)';" > src/app/version.ts
	cd $(FRONTEND_PATH) && echo "export const APP_BUILD_DATE = '$(DATE)';" >> src/app/version.ts
	cd $(FRONTEND_PATH) && npm run build-prod && git checkout -- src/app/version.ts

# Backend
# =================

backend-prepare: generate

backend-verify: backend-prepare
	$(GOCMD) fmt ./...
	$(GOCMD) vet ./...
	$(GOCMD) test ./...

backend: backend-verify
	$(GOCMD) build -o $(EXECUTEABLE) ./cmd/compelo


# Docker
# =================

docker-build:
	docker build \
	 	-t florianehmke/compelo:latest \
	 	-t florianehmke/compelo:$(TAG) .

docker-push: docker-build
	docker push florianehmke/compelo:latest
	docker push florianehmke/compelo:$(TAG)

# Cleanup
# =================

clean:
	rm -f $(EXECUTEABLE)

distclean: clean
	rm -f -r frontend/compelo/dist
	rm -f frontend/frontend_vfsdata.go
	rm -f internal/db/scripts/scripts_vfsdata.go
	rm frontend/compelo/src/generated/*.models.ts
