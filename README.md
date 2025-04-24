# Desafio Pós Go Expert :: Clean Architecture

Criar o usecase de listagem das orders. Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL

Criar as migrações necessárias;
Ciar arquivo order.http com a request para criar e listar as orders (Rest e GraphQL).

Orquestrar o banco de dados MySQL utilizando Docker (Dockerfile e/ou docker-compose.yaml).

## Portas
- RabbitMq : 5672 
- RabbitMq-management: 15672
- Banco MySQL : 3306
- API Rest : 8000
- API GraphQL : 8080
- API gRPC : 50051

## Versão do Go
O projeto foi atualizado para versão 1.24.2 da linguagem.

## Rodando o projeto
O projeto está usando tecnologia de devcontainer para subir um ambiente de desenvolvimento isolado e configurado para a linguagem. A primeira vez irá demorar um pouco pois são instalados todas as ferramentas e dependencias para a execução do desafio.

Estando no container de desenvolvimento, será necessário executar os seguintes passos:

- Executar os containers do MySQL e do HabbitMq:
```bash
make docker-up
```
- Implantar as migrações:
```bash
make migrate-mysql-up
```
- Rodar as APIs:
```bash
make go-run
```

Os endpoints para ciar e listar as orders tanto na API Rest (porta 8000) quanto na API GraphQL (porta 8080), estão no arquivo order.http localizado na pasta api localizado na rapiz do projeto.

Para a API em gRPC, será utlizado o evans como client. Com as apis rodando, abra um novo terminal, certifique-se que você está na raiz do projeto então execute o comando:
```bash
make go-evans
```