.PHONY: demo

demo:
	@scripts/run_demo.sh

app:
	@scripts/new_app.sh

swag:
	@swag init --generalInfo pkg/api/api.go
