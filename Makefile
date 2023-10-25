MKFILE := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
PARENT_MKFILE   := $(HOME)/.Makefile

include $(PARENT_MKFILE)

#where the go.mod files are located
PATH_TO_MOD := ./github/carlosrodlop/go-snippets
#where the go.work file is located
PATH_TO_WS  := ../../../

#https://go.dev/doc/tutorial/workspaces
.PHONY: go-new
go-new: ## Create a new module passed as parameter (MOD). Eg: MOD=04_Receivers make go-new
go-new: guard-MOD
	mkdir $(MOD)
	cd $(MOD) && \
		echo "package main" > main.go && \
		echo "func main() {}" >> main.go && \
		go mod init $(MOD)
	cd $(PATH_TO_WS) && go work use $(PATH_TO_MOD)/$(MOD)

.PHONY: go-run
go-run: ## Run an existing module passed as parameter (MOD). Eg: MOD=02_pointer make go-run
go-run: guard-MOD
	cd $(MOD) && go run main.go

.PHONY: go-test
go-test: ## Test an existing module passed as parameter (MOD). Eg: MOD=02_pointer make go-test
go-test: guard-MOD
	@echo "Test Result $(MOD)"
	cd $(MOD) && go test -v
	@echo "Test Coverage $(MOD)"
	cd $(MOD) && go test -coverprofile=coverage.out && \
		go tool cover -html=coverage.out
