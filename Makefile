test:
	go test -v -cover -coverprofile=c.out ./internal/pkg/...; \
    go tool cover -html=c.out -o coverage.html
importer:
	go build -o importer --mod=vendor ./cmd/importer;\
    ./importer
cli:
	go build -o cli --mod=vendor ./cmd/cli;\
    ./cli