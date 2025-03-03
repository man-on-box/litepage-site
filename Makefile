dev:
	@echo "Starting dev mode"
	@LP_MODE=serve LP_PORT=3001 air & \
		npx @tailwindcss/cli -i ./style/main.css -o ./public/main.css --watch

serve:
	@npx @tailwindcss/cli -i ./style/main.css -o ./public/main.css & \
		LP_MODE=serve go run cmd/main.go

build:
	@echo "Building for prod"
	go fmt ./...
	go vet ./...
	@go run cmd/main.go
	@npx @tailwindcss/cli -i ./style/main.css -o ./dist/main.css --minify
