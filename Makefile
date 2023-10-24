MKFILE := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
PARENT_MKFILE   := $(HOME)/.Makefile

include $(PARENT_MKFILE)

#where the go.mod files are located
PATH_TO_MOD := ./github/carlosrodlop/go-snippets
#where the go.work file is located
PATH_TO_WS  := ../../../

#https://go.dev/doc/tutorial/workspaces
.PHONY: newModule
newModule: ## Create a new module passed as parameter (MOD). Eg: MOD=04_Receivers make newModule
newModule:
	mkdir $(MOD)
	cd $(MOD) && \
		echo "package main" > main.go && \
		echo "func main() {}" >> main.go && \
		go mod init $(MOD)
	cd $(PATH_TO_WS) && go work use $(PATH_TO_MOD)/$(MOD)

.PHONY: newModule
run: ## Run an existing module passed as parameter (MOD). Eg: MOD=02_pointer make run
run:
	cd $(MOD) && go run main.go
