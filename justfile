
dev: generate
    go run main.go

generate:
    sqlc generate
    templ generate
