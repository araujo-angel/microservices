# Microservices - Order, Payment & Shipping Services

Projeto de microserviços desenvolvido para a disciplina de **Sistemas Distribuídos** no **IFPB 2025.2**.

## Arquitetura

Este projeto utiliza **Arquitetura Hexagonal (Ports and Adapters)**, proporcionando:
- Separação clara entre lógica de negócio e infraestrutura
- Independência de frameworks e tecnologias externas
- Comunicação entre microserviços via gRPC
- Facilidade para testes e manutenção


### Estrutura

```
microservices/
├── docker-compose.yml        # Orquestração dos containers
├── init.sql                  # Script de inicialização dos bancos
├── order/                    # Microserviço de Pedidos
│   ├── Dockerfile
│   ├── cmd/                  # Ponto de entrada
│   ├── config/               # Configurações
│   └── internal/
│       ├── adapters/         # Implementações (gRPC, DB, REST, Payment Client, Shipping Client)
│       ├── application/      
│       └── ports/            
├── payment/                  # Microserviço de Pagamentos
│   ├── Dockerfile
│   ├── cmd/                  
│   ├── config/               
│   └── internal/
│       ├── adapters/         # Implementações (gRPC, DB)
│       ├── application/      
│       └── ports/           
└── shipping/                 # Microserviço de Entregas
    ├── Dockerfile
    ├── cmd/                  
    ├── config/               
    └── internal/
        ├── adapters/         # Implementações (gRPC)
        ├── application/      
        └── ports/            
```

## Tecnologias

- **Go** - Linguagem de programação
- **gRPC** - Comunicação entre serviços
- **GORM** - ORM para acesso ao banco de dados
- **MySQL** - Banco de dados relacional
- **Protocol Buffers** - Definidos em [microservices-proto](https://github.com/araujo-angel/microservices-proto)

> **Nota sobre Docker:** Os `go.mod` usam `replace` com caminhos locais para os arquivos proto. Como o contexto do Docker é isolado, os Dockerfiles clonam o repositório proto durante o build para satisfazer essas dependências.

## Como Executar

### Pré-requisitos
- Docker
- Docker Compose
- grpcurl (para testes)

### Passos

#### **1. Subir todos os serviços**
```powershell
cd microservices\microservices
docker-compose up --build
```
Aguarde até que todos os serviços estejam prontos.

#### **2. Teste com grpcurl**
```powershell
grpcurl -d '{\"costumer_id\": 123, \"order_items\": [{\"product_code\": \"prod\", \"quantity\": 4, \"unit_price\": 12.0}]}' -plaintext localhost:8080 Order/Create
```

**Resposta esperada:**
```json
{
  "orderId": 1,
  "deliveryDays": 1
}
```

### Testes de Tolerância a Falhas

O sistema possui retry automático (5 tentativas, 1s de intervalo) e timeout de 15s.

#### **Teste 1: Retry com Payment indisponível**
```powershell
docker stop payment
grpcurl -d '{\"costumer_id\": 123, \"order_items\": [{\"product_code\": \"prod\", \"quantity\": 4, \"unit_price\": 12.0}]}' -plaintext localhost:8080 Order/Create
docker start payment
```
**Esperado:** Erro `Unavailable` após ~5 segundos (5 tentativas de retry).

#### **Teste 2: Retry com Shipping indisponível**
```powershell
docker stop shipping
grpcurl -d '{\"costumer_id\": 123, \"order_items\": [{\"product_code\": \"prod\", \"quantity\": 4, \"unit_price\": 12.0}]}' -plaintext localhost:8080 Order/Create
docker start shipping
```
**Esperado:** Erro `Unavailable` após ~5 segundos.

#### **Acompanhar logs**
```powershell
docker logs -f order
```

### Variáveis de Ambiente

| Serviço | Variável | Descrição |
|---------|----------|-----------|
| Order | `DATA_SOURCE_URL` | Conexão MySQL: `root:s3cr3t@tcp(mysql:3306)/order?charset=utf8mb4&parseTime=True&loc=Local` |
| Order | `APPLICATION_PORT` | Porta do serviço: `8080` |
| Order | `PAYMENT_SERVICE_URL` | Endereço do Payment: `payment:8081` |
| Order | `SHIPPING_SERVICE_URL` | Endereço do Shipping: `shipping:8082` |
| Order | `ENV` | Ambiente: `development` (habilita gRPC reflection) |
| Payment | `DATA_SOURCE_URL` | Conexão MySQL: `root:s3cr3t@tcp(mysql:3306)/payment?charset=utf8mb4&parseTime=True&loc=Local` |
| Payment | `APPLICATION_PORT` | Porta do serviço: `8081` |
| Payment | `ENV` | Ambiente: `development` (habilita gRPC reflection) |
| Shipping | `APPLICATION_PORT` | Porta do serviço: `8082` |
| Shipping | `ENV` | Ambiente: `development` (habilita gRPC reflection) |

> **Nota:** O parâmetro `parseTime=True` na URL de conexão MySQL é necessário para que o driver converta corretamente campos de data/hora para `time.Time` do Go.

### Fluxo de Comunicação

```
[grpcurl]
    ↓ (cria pedido)
[Order Service :8080] ← MySQL :3306/order
    ├──→ (processa pagamento)
    │    [Payment Service :8081] ← MySQL :3306/payment
    │
    └──→ (calcula entrega)
         [Shipping Service :8082]
```

**IFPB - Instituto Federal da Paraíba** | Sistemas Distribuídos 2025.2