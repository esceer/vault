build:
	go build -o build/vault.exe cmd/web/main.go

buildCli:
	go build -o build/vault_cli.exe cmd/cli/main.go

run:
	go run cmd/web/main.go

runCli:
	go run cmd/cli/main.go
