.PHONY: fmt fmt-check lint test fmt-go fmt-ts lint-ts test-go test-ts

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

# Test

test: test-go test-ts

test-go:
	go test -race ./...

test-ts:
	node --test '**/*_test.ts'

# Check (CI)

fmt-check:
	gofmt -l . | grep -q . && exit 1 || true
	npx biome check .
