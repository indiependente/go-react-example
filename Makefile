SRVFOLDER=server
CLNTFOLDER=client
GOROOT=$(shell go env GOROOT)
all: deps build run

.PHONY: build
build: clean ui pre-build
	@mkdir -p ./$(SRVFOLDER)/bin && \
	cd $(SRVFOLDER) && \
	GOROOT=$(GOROOT) rice embed-go && \
	CGO_ENABLED=0 go build -o ./bin/service . && \
	rm rice-box.go

.PHONY: clean
clean:
	@rm -rf ./$(SRVFOLDER)/bin
	@rm -rf ./$(CLNTFOLDER)/dist

.PHONY: deps
deps:
	@npm i
	@go mod download

.PHONY: docker
docker: ui
	@docker build -f ./$(SRVFOLDER)/Dockerfile -t go-react .

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: run
run:
	@cd $(SRVFOLDER) && go run main.go

.PHONY: ui
ui:
	@cd $(CLNTFOLDER) && npm run build

.PHONY: pre-build
pre-build:
	@command -v rice || \
	(go get github.com/GeertJohan/go.rice && \
	go get github.com/GeertJohan/go.rice/rice)

.PHONY: update_deps
update_deps:
	@go mod tidy
