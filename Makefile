# Makefile for shenad
#
# Copyright (C) 2023 takfu.cf
#
# This is free software, licensed under the MIT General Public License.
# See /LICENSE for more information.
#

VERSION=1.0.0
GO_GEN=go generate main/shenad.go
NPM_PACK=npm run build:prod
DIST_DIR=./release
GO_CLEAN=go clean
NPM_CLEAN=npm cache clean -f
NODE_PATH=/usr/local
NODE_VERSION=15.12.0
GO_PATH=/usr/local
GO_VERSION=1.15.3



shenad:	.bashrc generation
	mkdir -p $(DIST_DIR) 2>/dev/null
	flags="-X 'main.goversion=$(go version)'"
	go build -ldflags "$$flags" -x -o $(DIST_DIR)/$@-${VERSION} main/shenad.go
	echo "Done"

generation:	.bashrc dist
	source /$(shell whoami)/$<
	$(GO_GEN)

dist:	clean
	$(NPM_PACK)

clean:	modules
	npm install --save

modules:	native
	rm node_$@ -fR
	$(GO_CLEAN)
	$(NPM_CLEAN)

native:	.bashrc
	if [ -z "$(echo $(whereis g++) | tr 'g++:' ' ')" ]; then yum install gcc-c++ -y  ;fi
	yum install python38.x86_64 -y
	source /$(shell whoami)/$^
	go get github.com/rakyll/statik


.ONESHELL:
.bashrc:	world
	person=$(shell whoami)
	mkdir -p /$$person/.gopath >/dev/null
	cat <<- EOF >> /$$person/$@
	GOPATH=/$$person/.gopath
	export GOPATH
	EOF
	echo 'export PATH=$$PATH:$$GOPATH/bin' >> /$$person/$@


world:
	@if [[ "$$OSTYPE" == "linux-gnu" ]]; then \
		echo "your OSTYPE is linux64 , Just waiting until PROGMS gracefully end." ;\
		sleep 2; \
	else \
		echo "The current os is not supported" 1>&2; \
		exit 1; \
	fi
	@if type node >/dev/null 2>&1; then \
		echo 'exists node' ; \
	else \
		echo 'not exists node' ; \
		echo "installing nodejs" ; \
		echo "https://nodejs.org/dist/latest-v15.x/node-v$(NODE_VERSION)-linux-x64.tar.gz" ; \
		curl -jLO "https://nodejs.org/dist/latest-v15.x/node-v$(NODE_VERSION)-linux-x64.tar.gz" ; \
		tar -C $(NODE_PATH) -zxvf node-*.tar.gz ; \
		mv -vf $(NODE_PATH)/node-v$(NODE_VERSION)-* $(NODE_PATH)/node 2>/dev/null ; \
		ln -sfn $(NODE_PATH)/node/bin/node /usr/bin/node ; \
		ln -sfn $(NODE_PATH)/node/bin/npm /usr/bin/npm ; \
		ln -sfn $(NODE_PATH)/node/bin/npx /usr/bin/npx ; \
		rm node-*.tar.gz -fR ; \
		echo "install nodejs done" ; \
	fi
	@if type go >/dev/null 2>&1; then \
		echo 'exists go' ; \
	else \
		echo 'not exists go' ; \
		curl -jLO "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz" ; \
		tar -C $(GO_PATH) -zxvf go$(GO_VERSION).linux*.tar.gz ; \
		ln -sfn $(GO_PATH)/go/bin/go /usr/bin/go ; \
		rm go$(GO_VERSION).linux*.tar.gz -fR ; \
		echo "install golang done" ; \
	fi


.PHONY: shenad generation dist clean modules native world
