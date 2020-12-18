.PHONY: demo build swag

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
	@scripts/run_build.sh

