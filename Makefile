
.PHONY: govet
govet:
	go vet $$( go list ./internal/...)