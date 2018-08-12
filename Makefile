# Makefile to build the command lines and tests in Seele project.
# This Makefile doesn't consider Windows Environment. If you use it in Windows, please be careful.

SHELL := /bin/bash 


BUILD_NAME = gin-quick-api

PROJECT_DIR= gin-quick
BASEDIR = $(shell pwd)

# build with verison infos
versionDir = github.com/wxcsdb88/gin-quick/version

gitBranch = $(shell git symbolic-ref --short -q HEAD)

ifeq ($(gitBranch),)
gitTag = $(shell git describe --tags --abbrev=0)
endif

buildDate = $(shell date "+%FT%T%z")
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)
 
ldflagsOrigin="-X ${versionDir}.gitBranch=${gitBranch} -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} \
 -X ${versionDir}.gitCommit=${gitCommit} \
 -X ${versionDir}.gitTreeState=${gitTreeState}"

# -s -w 
ldflagsRelase="-s -w -X ${versionDir}.gitBranch=${gitBranch} -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} \
 -X ${versionDir}.gitCommit=${gitCommit} \
 -X ${versionDir}.gitTreeState=${gitTreeState}"


all: release
api:
	go build -v -ldflags ${ldflagsOrigin} -o ./build/bin/${BUILD_NAME} ./cmd/api
	@echo "Done gin-quick building debug"

release:
	go build -v -ldflags ${ldflagsRelase} -o ./build/bin/${BUILD_NAME} ./cmd/api
	@echo "Done gin-quick building release"

clean:
	rm ${BASEDIR}/build/bin/*

.PHONY: api release