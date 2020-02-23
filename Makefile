SRVFOLDER=server
CLNTFOLDER=client
GOROOT=$(shell go env GOROOT)
all: deps build run	

.PHONY: build
build: clean ui pre-build
	mkdir -p ./$(SRVFOLDER)/bin && \
	CGO_ENABLED=0 go build -o ./$(SRVFOLDER)/bin/service && \
	echo "[✔️] Build complete!"

.PHONY: clean
clean:
	@rm -rf ./$(SRVFOLDER)/bin
	@rm -rf ./$(CLNTFOLDER)/build
	@echo "[✔️] Clean complete!"

.PHONY: deps
deps:
	@go mod download

.PHONY: docker
docker:
	@cd $(SRVFOLDER) && GOROOT=$(GOROOT) rice embed-go
	@docker build -f ./$(SRVFOLDER)/Dockerfile -t go-react .
	@cd $(SRVFOLDER) && rm rice-box.go

.PHONY: run
run:
	@cd $(SRVFOLDER) && ENV=dev go run main.go

.PHONY: ui
ui:
	@cd $(CLNTFOLDER) && npm i && npm run build

.PHONY: pre-build
pre-build:
	go get github.com/GeertJohan/go.rice
	go get github.com/GeertJohan/go.rice/rice

.PHONY: update_deps
update_deps:
	@go mod tidy
