SRVFOLDER=server
CLNTFOLDER=client

all: deps build run	

.PHONY: build
build: clean ui
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
	@docker build -f ./$(SRVFOLDER)/Dockerfile -t go-react .

.PHONY: run
run:
	@ENV=dev go run main.go

.PHONY: ui
ui:
	@cd $(CLNTFOLDER) && npm i && npm run build

.PHONY: update_deps
update_deps:
	@go mod tidy
