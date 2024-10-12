# Makefile

SHELL := /bin/bash
TAG := $(shell git describe --tags --abbrev=0)
VERSION := TAG
NEW_VERSION := $(shell echo $(TAG) | awk -F. '{$$NF = $$NF + 1;} 1' | sed 's/ /./g')

define colorecho
      @tput setaf 6
      @echo $1
      @tput sgr0
endef

.PHONY: run

changelog: test
	$(call colorecho, "Preparing CHANGELOG...")
	git-changelog -c angular -s docs,feat,test,chore --output CHANGELOG.MD

test:
	go test ./...
check:
	@echo Old tag: $(TAG) New tag: $(NEW_VERSION)

pre-release:
	$(call colorecho, "Preparing Release with new version...")
	git add .
	git tag $(NEW_VERSION)
	$(MAKE) changelog
	git tag -d $(NEW_VERSION)
	git add . && git commit -m "bump: version $(TAG) -> $(NEW_VERSION)" && git tag $(NEW_VERSION)

release: pre-release
	$(call colorecho, "Push to origin with TAGs...")
	git push
	git push --tags

swagger:
	swag init