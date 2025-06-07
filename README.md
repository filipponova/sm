# AWS Session Manager TUI (Go)

Terminal application to list and connect to EC2 instances via AWS Session Manager.

## Requirements

- Go 1.22+
- AWS CLI configured with profiles

## Installation

```bash
git clone https://github.com/filipponova/sm.git
cd sm
go mod tidy
go build -o sm
```

## Usage

```bash
./sm list --region us-east-1 --profile default
```

## Estrutura do Projeto

- `cmd/` - Comandos cobra (root, list, etc)
- `internal/` - Lógica de negócio (AWS, sessão, tipos)
- `main.go` - Inicialização do CLI
