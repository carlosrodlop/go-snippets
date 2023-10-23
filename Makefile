MKFILE := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
PARENT_MKFILE   := $(HOME)/.Makefile

include $(PARENT_MKFILE)

PATH_TO_WORK := ./carlosrodlop/go-snippets

.PHONY: newModule
newModule: ## Create a new module. Eg: MOD=04_Receivers make newModule
newModule:
	mkdir $(MOD)
	cd $(MOD) && \
		echo "package main" > main.go && \
		echo "func main() {}" >> main.go && \
		go mod init $(MOD)
	echo "add to go.work $(PATH_TO_WORK)/$(MOD)"