.PHONY: demo build swag

# run demo app
demo:
	@scripts/run_demo.sh

# build docker
build:
	@scripts/run_build.sh

# generate swag docs
swag:
	@swag init --generalInfo pkg/api/api.go

# end of the file
_app:
	@scripts/new_app.sh
