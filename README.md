# 🚀 Como Rodar o Projeto

## 📋 Pré-requisitos

- [Docker](https://docs.docker.com/get-docker/) e [Docker Compose](https://docs.docker.com/compose/) (forma recomendada)
- [Go 1.25+](https://go.dev/dl/) (apenas se for rodar sem Docker)

## ⚙️ Configuração

O projeto lê as variáveis de ambiente de um arquivo `.env` na raiz. Copie o exemplo:

```bash
cp .env-example .env
```

| Variável | Descrição | Padrão |
|----------|-----------|--------|
| `DB_USER` | Usuário do Postgres | `postgres` |
| `DB_PASSWORD` | Senha do Postgres | `postgres` |
| `DB_NAME` | Nome do banco | `backend_test` |
| `DB_HOST` | Host do banco (use `localhost` fora do Docker) | `localhost` |
| `DB_PORT` | Porta exposta no host | `5434` |
| `DB_SSLMODE` | Modo SSL da conexão | `disable` |

## 🐳 Rodando com Docker (recomendado)

Um único comando sobe o banco, aplica as migrations e builda/sobe a API:

```bash
docker compose up --build
```

O Compose orquestra três serviços na ordem correta:

1. **`postgres`** — sobe o banco e aguarda ficar saudável (`healthcheck`).
2. **`migrate`** — aplica as migrations em `./migrations` e encerra.
3. **`api`** — só inicia após o banco estar saudável e as migrations concluídas.

A API fica disponível em `http://localhost:8080`.

Para parar e remover os containers:

```bash
docker compose down
```

Para remover também os dados do banco (volume):

```bash
docker compose down -v
```

## 💻 Rodando localmente (sem Docker)

1. Suba apenas o Postgres via Docker (ou use uma instância própria):

   ```bash
   docker compose up -d postgres migrate
   ```

2. Ajuste o `.env` para apontar ao banco local (`DB_HOST=localhost`, `DB_PORT=5434`).

3. Rode a aplicação:

   ```bash
   go run ./cmd
   ```

## 🗄️ Migrations

As migrations ficam em `./migrations` no formato do [golang-migrate](https://github.com/golang-migrate/migrate). No fluxo Docker elas são aplicadas automaticamente pelo serviço `migrate`. Para criar uma nova:

```bash
migrate create -ext sql -dir migrations -seq nome_da_migration
```

## 🧪 Testes

```bash
go test ./...
```

Os testes unitários do cálculo de prioridade (incluindo cenários extremos: estoque negativo, venda zero e lead time alto) estão em `internal/domain/part/part_test.go`.

---

# 📡 Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| `POST` | `/part` | Cria uma peça |
| `GET` | `/part` | Lista todas as peças |
| `PUT` | `/part/{id}` | Atualiza uma peça |
| `DELETE` | `/part/{id}` | Remove uma peça |
| `GET` | `/restock/priorities` | Lista as peças ordenadas por prioridade de reposição |

## 📝 Exemplos de Requisição

Exemplos prontos (clicáveis no GoLand/IntelliJ) estão em [`internal/examples/requests.http`](internal/examples/requests.http). Abaixo os equivalentes em `curl`:

**Criar peça:**

```bash
curl -X POST http://localhost:8080/part \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Filtro de Óleo X",
    "category": "engine",
    "currentStock": 15,
    "minimumStock": 20,
    "averageDailySales": 4,
    "leadTimeDays": 5,
    "criticalityLevel": 3,
    "unitCost": 18.50
  }'
```

**Listar peças:**

```bash
curl http://localhost:8080/part
```

**Atualizar peça:**

```bash
curl -X PUT http://localhost:8080/part/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Filtro de Óleo X",
    "category": "engine",
    "currentStock": 50,
    "minimumStock": 30,
    "averageDailySales": 4,
    "leadTimeDays": 5,
    "criticalityLevel": 3,
    "unitCost": 18.50
  }'
```

**Remover peça:**

```bash
curl -X DELETE http://localhost:8080/part/{id}
```

**Prioridades de reposição:**

```bash
curl http://localhost:8080/restock/priorities
```
