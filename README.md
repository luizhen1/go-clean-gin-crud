# Go Clean Architecture CRUD API 🚀

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go&logoColor=white)](https://go.dev/)
[![Docker](https://img.shields.io/badge/docker-powered-blue.svg)](https://www.docker.com/)
[![Swagger](https://img.shields.io/badge/swagger-OpenAPI-green.svg)](http://swagger.io/)

Este projeto é uma API RESTful robusta desenvolvida em Go, projetada sob os princípios da **Clean Architecture**. O objetivo principal é fornecer um sistema de CRUD (Create, Read, Update, Delete) de produtos, totalmente containerizado e documentado.

---

## 🏗️ System Design

O diagrama abaixo ilustra a arquitetura do sistema, destacando o isolamento dos serviços via Docker e o fluxo de dados entre as camadas da Clean Architecture.

<img width="1059" height="818" alt="image" src="https://github.com/user-attachments/assets/6be0e2c7-521f-421e-a4bd-09316a9cc368" />



---

## ✨ Funcionalidades

- **CRUD de Produtos**: Gerenciamento completo de registros no banco de dados.
- **Clean Architecture**: Separação clara entre Domínio, Casos de Uso (Services), Adaptadores (Repositories) e Infraestrutura (Routes).
- **Docker Compose**: Orquestração completa que sobe a API e o Banco de Dados com um único comando.
- **Database Migrations**: Inicialização automática do esquema do banco de dados MySQL via scripts SQL.
- **Swagger UI**: Documentação interativa para teste de endpoints diretamente pelo navegador.
- **Resiliência**: Lógica de auto-retry na conexão com o banco de dados durante o startup.

---

## 🛠️ Tecnologias e Ferramentas

- **Linguagem**: [Go 1.25+](https://go.dev/)
- **Web Framework**: [Gin Gonic](https://gin-gonic.com/)
- **Banco de Dados**: [MySQL 8.0](https://www.mysql.com/)
- **Documentação**: [Swaggo](https://github.com/swaggo/swag)
- **Container**: [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)

---

## 🚀 Como Rodar o Projeto

### Pré-requisitos
Você precisará apenas do **Docker** e **Docker Compose** instalados em sua máquina.

### Passos para execução
1. **Clone o repositório:**
   ```bash
   git clone [https://github.com/luizhen1/go-clean-gin-crud.git](https://github.com/luizhen1/go-clean-gin-crud.git)
   cd go-clean-gin-crud
   ´´´
2. **Suba os containers:**
   ```bash
   docker-compose up --build
   OBS: Este comando irá compilar a aplicação, configurar a rede interna e rodar o script de criação da tabela no MySQL.
   ´´´
3. **Documentação e Testes (Swagger):**
Para acessar a documentação interativa e testar os endpoints:
http://localhost:8080/swagger/index.html
   ```bash
   http://localhost:8080/swagger/index.html
   ´´´
4. **Endpoints principais:**
POST /products: Cria um novo produto.
GET /products: Lista todos os produtos.
GET /products/:id: Busca um produto específico.
PUT /products/:id: Atualiza um produto existente.
DELETE /products/:id: Remove um produto.

5. **Estrutura de Pastas:**
├── cmd/api/            # Ponto de entrada (Main)
├── docs/               # Documentação gerada pelo Swagger
├── internal/
│   ├── database/       # Configuração e conexão do banco (MySQL)
│   ├── domain/         # Entidades e Interfaces (Core do Negócio)
│   ├── repositories/   # Implementação de acesso a dados (SQL)
│   ├── routes/         # Handlers HTTP e definição de rotas (Gin)
│   └── services/       # Lógica de negócio (Use Cases)
├── migrations/         # Scripts SQL de inicialização
├── Dockerfile          # Receita de build da imagem Go
└── docker-compose.yml  # Orquestração de containers

6. **🤝 Contribuição**
Faça um Fork do projeto.
Crie uma nova Branch (git checkout -b feature/nome-da-feature).
Faça o Commit de suas alterações (git commit -m 'Add some feature').
Envie para o seu repositório (Push) (git push origin feature/nome-da-feature).
Abra um Pull Request.

### Dicas Finais:
1.  **Imagens**: Certifique-se de que o arquivo da imagem do seu diagrama (`CRUD.drawio.png`) esteja na raiz do repositório para que o link `![System Design](./CRUD.drawio.png)` funcione.
2.  **Badges**: As "badges" no topo (Go Version, Docker, etc) dão um toque visual muito bacana.
3.  **Markdown**: O GitHub renderiza isso automaticamente. Se precisar editar algo, use o editor do próprio GitHub ou a pré-visualização do seu VS Code (`Ctrl + Shift + V`).



  
