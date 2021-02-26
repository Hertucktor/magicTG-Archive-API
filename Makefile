test:
	go test -v -cover -coverprofile=c.out ./internal/pkg/...; \
    go tool cover -html=c.out -o coverage.html
csvReader:
	go build -o bin/csvReader --mod=vendor ./cmd/csvReader ;\
	./bin/csvReader
importer:
	go build -o bin/importer --mod=vendor ./cmd/importer ;\
    ./bin/importer
api:
	go build -o bin/api --mod=vendor ./cmd/api ;\
    ./bin/api
db:
	docker-compose up