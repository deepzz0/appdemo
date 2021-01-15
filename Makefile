.PHONY: demo build swag

REGISTRY=localhost
IMAGE_TAG=`git describe --tags`

swag:
	@scripts/swag_init.sh

_app:
	@scripts/new_app.sh

# below you should write

# run demo app
demo:
	@scripts/run_app.sh demo

# build docker
build:
	@scripts/run_build.sh $(REGISTRY) $(IMAGE_TAG)

# protoc
protoc:
	@cd pkg/proto && make protoc
