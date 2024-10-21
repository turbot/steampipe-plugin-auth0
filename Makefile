STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo
install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/auth0@latest/steampipe-plugin-auth0.plugin -tags "${BUILD_TAGS}" *.go
