.PHONY: fmt fmt-check lint fmt-go fmt-ts lint-ts

# Format

fmt: fmt-go fmt-ts

fmt-go:
	gofmt -w .

fmt-ts:
	npx biome format --write .

# Lint

lint: lint-ts
	go vet ./...

lint-ts:
	npx biome check .

# Check (CI)

fmt-check:
	gofmt -l . | grep -q . && exit 1 || true
	npx biome check .
