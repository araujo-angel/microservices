# Microservices - Order Service

Projeto de microserviços desenvolvido para a disciplina de **Sistemas Distribuídos** no **IFPB 2025.2**.

## Arquitetura

Este serviço utiliza **Arquitetura Hexagonal (Ports and Adapters)**, proporcionando:
- Separação clara entre lógica de negócio e infraestrutura
- Independência de frameworks e tecnologias externas
- Facilidade para testes e manutenção

![Arquitetura](diagrama.png)

### Estrutura

```
order/
├── cmd/                    # Ponto de entrada da aplicação
├── internal/
│   ├── adapters/          # Implementações concretas (gRPC, DB)
│   ├── application/core/  # Lógica de negócio
│   │   ├── api/          # Orquestração
│   │   └── domain/       # Entidades de domínio
│   └── ports/            # Interfaces (contratos)
└── config/               # Configurações
```

## Tecnologias

- **Go** - Linguagem de programação
- **gRPC** - Comunicação entre serviços
- **GORM** - ORM para acesso ao banco de dados
- **MySQL** - Banco de dados relacional
- **Protocol Buffers** - Definidos em [microservices-proto](https://github.com/araujo-angel/microservices-proto)

## Como Executar

### Pré-requisitos
- Docker
- Go 1.21+
- grpcurl (para testes)

### Passos

1. **Iniciar MySQL no Docker:**
```bash
docker run --name mysql-orders -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=orders -p 3306:3306 -d mysql:8.0
```

2. **Configurar variáveis de ambiente:**
```bash
export DATA_SOURCE_URL="root:root@tcp(127.0.0.1:3306)/orders?charset=utf8mb4&parseTime=True&loc=Local"
export APPLICATION_PORT=3000
export ENV=development
```

3. **Executar o serviço:**
```bash
cd order
go run cmd/main.go
```

4. **Testar com grpcurl:**
```powershell
echo '{"costumer_id": 123, "order_items": [{"product_code": "prod", "quantity": 4, "unit_price": 12}]}' | grpcurl -d '@' -plaintext localhost:3000 Order/Create
```

**IFPB - Instituto Federal da Paraíba** | Sistemas Distribuídos 2025.2