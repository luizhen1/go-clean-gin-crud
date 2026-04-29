# Go Clean Architecture CRUD API 🚀

[![Go Version](https://img.shields.io/github/go-mod/go-v/luizhen1/go-clean-gin-crud)](https://go.dev/)
[![Docker](https://img.shields.io/badge/docker-powered-blue.svg)](https://www.docker.com/)
[![Swagger](https://img.shields.io/badge/swagger-OpenAPI-green.svg)](http://swagger.io/)

Este projeto é uma API RESTful robusta desenvolvida em Go, projetada sob os princípios da **Clean Architecture**. O objetivo principal é fornecer um sistema de CRUD (Create, Read, Update, Delete) de produtos, totalmente containerizado e documentado.

---

## 🏗️ System Design

O diagrama abaixo ilustra a arquitetura do sistema, destacando o isolamento dos serviços via Docker e o fluxo de dados entre as camadas da Clean Architecture.

![System Design]
<img width="1121" height="902" alt="image" src="https://github.com/user-attachments/assets/7cec76d3-b13c-4601-9e57-5810e88a5a16" />


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
