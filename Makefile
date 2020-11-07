.PHONY: demo

demo:
	@scripts/run_demo.sh

swag:
	@swag init --generalInfo pkg/api/api.go

# end of the file
_app:
	@scripts/new_app.sh
