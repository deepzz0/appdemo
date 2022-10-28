.PHONY: demo build swag

REGISTRY=localhost
IMAGE_TAG=`git describe --tags`

usage:
	@echo
	@echo '  Usage for migrate app:'
	@echo '    make swag                   generate all app swagger docs.'
	@echo '    make docker                 build docker image.'
	@echo '    make run app=<demo>         run app locally.'
	@echo '    make protoc                 generate all protocol files'
	@echo '    make _new                   new project at current dir.'

# swagger
swag:
	@scripts/swag.sh $(app)

# build docker
docker:
	@scripts/docker.sh $(REGISTRY) $(IMAGE_TAG) $(app)

# run app
run:
	@scripts/run.sh $(app)

# protoc
protoc:
	@scripts/protoc.sh

# new project
_new:
	@scripts/new_proj.sh
