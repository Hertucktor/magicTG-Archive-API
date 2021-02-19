test:
	go test -v -cover -coverprofile=c.out ./internal/pkg/...; \
    go tool cover -html=c.out -o coverage.html
imggetter:
	go build -o imggetter --mod=vendor ./cmd/imggetter ;\
	./imggetter
importer:
	go build -o importer --mod=vendor ./cmd/importer ;\
    ./importer
api:
	go build -o bin/api --mod=vendor ./cmd/api ;\
    ./bin/api
db:
	docker-compose up