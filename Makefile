.PHONY: demo

demo:
	@scripts/run_demo.sh

swag:
	@swag init --generalInfo pkg/api/api.go
