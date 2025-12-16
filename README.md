# Microservices - Order & Payment Services

Projeto de microserviços desenvolvido para a disciplina de **Sistemas Distribuídos** no **IFPB 2025.2**.

## Arquitetura

Este projeto utiliza **Arquitetura Hexagonal (Ports and Adapters)**, proporcionando:
- Separação clara entre lógica de negócio e infraestrutura
- Independência de frameworks e tecnologias externas
- Comunicação entre microserviços via gRPC
- Facilidade para testes e manutenção

![Arquitetura](diagrama.png)

### Estrutura

```
microservices/
├── order/                 # Microserviço de Pedidos
│   ├── cmd/              # Ponto de entrada
│   ├── internal/
│   │   ├── adapters/     # Implementações (gRPC, DB, Payment Client)
│   │   ├── application/  # Lógica de negócio
│   │   └── ports/        # Interfaces
│   └── config/           # Configurações
├── payment/              # Microserviço de Pagamentos
│   ├── cmd/              # Ponto de entrada
│   ├── internal/
│   │   ├── adapters/     # Implementações (gRPC, DB)
│   │   ├── application/  # Lógica de negócio
│   │   └── ports/        # Interfaces
│   └── config/           # Configurações
└── init.sql              # Script de inicialização dos bancos
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
- Go 1.25+
- grpcurl (para testes)

### Passos (Windows PowerShell)

Execute os comandos em **4 terminais diferentes** na ordem indicada:

#### **Terminal 1: MySQL**
```powershell
cd C:\Users\Angelica\Documents\GitHub\microservices\microservices
docker run -p 3307:3306 -e MYSQL_ROOT_PASSWORD=minhasenha -v "$PWD\init.sql:/docker-entrypoint-initdb.d/init.sql" mysql
```
Aguarde até ver a mensagem: `ready for connections`

#### **Terminal 2: Payment Service**
```powershell
cd C:\Users\Angelica\Documents\GitHub\microservices\microservices\payment
$env:DATA_SOURCE_URL="root:minhasenha@tcp(127.0.0.1:3307)/payment"
$env:APPLICATION_PORT="3001"
$env:ENV="development"
go run cmd/main.go
```
Aguarde até ver: `starting payment service on port 3001 ...`

#### **Terminal 3: Order Service**
```powershell
cd C:\Users\Angelica\Documents\GitHub\microservices\microservices\order
$env:DATA_SOURCE_URL="root:minhasenha@tcp(127.0.0.1:3307)/order"
$env:APPLICATION_PORT="3000"
$env:PAYMENT_SERVICE_URL="localhost:3001"
$env:ENV="development"
go run cmd/main.go
```
Aguarde até ver: `starting order service on port 3000 ...`

#### **Terminal 4: Teste com grpcurl**
```powershell
grpcurl -d '{\"costumer_id\": 123, \"order_items\": [{\"product_code\": \"prod\", \"quantity\": 4, \"unit_price\": 12.0}]}' -plaintext localhost:3000 Order/Create
```

**Resposta esperada:**
```json
{
  "order_id": 1
}
```

### Variáveis de Ambiente

| Serviço | Variável | Descrição |
|---------|----------|-----------|
| Order | `DATA_SOURCE_URL` | Conexão MySQL: `root:minhasenha@tcp(127.0.0.1:3307)/order` |
| Order | `APPLICATION_PORT` | Porta do serviço: `3000` |
| Order | `PAYMENT_SERVICE_URL` | Endereço do Payment: `localhost:3001` |
| Order | `ENV` | Ambiente: `development` (habilita gRPC reflection) |
| Payment | `DATA_SOURCE_URL` | Conexão MySQL: `root:minhasenha@tcp(127.0.0.1:3307)/payment` |
| Payment | `APPLICATION_PORT` | Porta do serviço: `3001` |
| Payment | `ENV` | Ambiente: `development` (habilita gRPC reflection) |

### Fluxo de Comunicação

```
[grpcurl]
    ↓ (cria pedido)
[Order Service :3000] ← MySQL :3307/order
    ↓ (processa pagamento)
[Payment Service :3001] ← MySQL :3307/payment
```

**IFPB - Instituto Federal da Paraíba** | Sistemas Distribuídos 2025.2