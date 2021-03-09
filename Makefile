test:
	go test -v -cover -coverprofile=c.out ./internal/pkg/...; \
    go tool cover -html=c.out -o coverage.html
csvReader:
	go build -o bin/csvReader --mod=vendor ./cmd/csvReader ;\
	./bin/csvReader
importer:
	go build -o bin/importer --mod=vendor ./cmd/importer ;\
    ./bin/importer
setNames:
	go build -o bin/setNames --mod=vendor ./cmd/setNames ;\
    ./bin/setNames
api:
	go build -o bin/api --mod=vendor ./cmd/api ;\
    ./bin/api
ui:
	go build -o bin/ui --mod=vendor ./cmd/ui ;\
    ./bin/ui
db:
	docker-compose up