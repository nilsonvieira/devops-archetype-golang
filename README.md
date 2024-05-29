# My API

Esta é uma API básica desenvolvida em Go (Golang) com integração ao banco de dados MariaDB. A estrutura do projeto é modular e organizada para facilitar a manutenção e a expansão.

## Estrutura do Projeto
```
go-api/
├── cmd/
│ └── main.go
├── config/
│ └── config.go
├── controllers/
│ └── user_controller.go
├── models/
│ └── user.go
├── repositories/
│ └── user_repository.go
├── services/
│ └── user_service.go
├── routes/
│ └── routes.go
├── go.mod
└── go.sum
```

### Descrição das Pastas e Arquivos

- **cmd/**: Contém o ponto de entrada da aplicação.
  - `main.go`: Arquivo principal que inicia a aplicação.

- **config/**: Configurações da aplicação.
  - `config.go`: Carrega as configurações e estabelece a conexão com o banco de dados.

- **controllers/**: Controladores que lidam com as requisições HTTP.
  - `user_controller.go`: Controlador responsável por gerenciar as requisições relacionadas aos usuários.

- **models/**: Definições das estruturas de dados.
  - `user.go`: Define a estrutura do modelo `User`.

- **repositories/**: Manipulação dos dados e interação com o banco de dados.
  - `user_repository.go`: Repositório responsável por acessar e manipular os dados dos usuários no banco de dados.

- **services/**: Contém a lógica de negócio da aplicação.
  - `user_service.go`: Serviço responsável pela lógica de negócio relacionada aos usuários.

- **routes/**: Define as rotas da API.
  - `routes.go`: Configura as rotas da API utilizando o Gorilla Mux.

- **go.mod**: Arquivo de gerenciamento de dependências do Go.
- **go.sum**: Arquivo que contém checksums das dependências.

## Configuração do Banco de Dados

1. Crie o banco de dados e a tabela `users` no MariaDB:
   ```sql
   CREATE DATABASE goapi;
   USE goapi;

   CREATE TABLE users (
       id INT AUTO_INCREMENT,
       name VARCHAR(100),
       email VARCHAR(100),
       PRIMARY KEY (id)
   );

Defina as variáveis de ambiente para conexão ao banco de dados:
```sh
export DB_USER='goapi_user'
export DB_PASSWORD='goapi_password'
export DB_HOST='localhost'
export DB_NAME='goapi'
```
## Executando a API
Instale as dependências do projeto:

```sh
go mod tidy
```
## Execute a aplicação:

```sh
go run cmd/main.go
```

## Endpoints da API
`GET /users`: Retorna todos os usuários cadastrados.

# Dependências
Gorilla Mux: Pacote para roteamento e manipulação de requisições HTTP.
MySQL Driver: Driver para conexão com o banco de dados MariaDB/MySQL.

# Contribuição
Sinta-se à vontade para contribuir com este projeto. Faça um fork, crie uma branch com sua feature ou correção de bug e envie um pull request.

# Licença
Este projeto está licenciado sob a Licença MIT - veja o arquivo LICENSE para mais detalhes.
