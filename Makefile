test:
	go test -v -cover -coverprofile=c.out ./internal/pkg/...; \
    go tool cover -html=c.out -o coverage.html
app:
	go build -o app --mod=vendor ./cmd ;\
    ./app
db:
	docker-compose up