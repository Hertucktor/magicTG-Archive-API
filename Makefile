test:
	go test -v -cover -coverprofile=c.out; \
    go tool cover -html=c.out -o coverage.html
importer: test
	go build -o importer --mod=vendor ./cmd/importer;\
    ./importer