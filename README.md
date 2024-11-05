# libCon-Postgre-GO

Esta é uma biblioteca Go para realizar operações básicas de CRUD (Create, Read, Update, Delete) em um banco de dados PostgreSQL. Ela oferece funções para inserir, selecionar, atualizar e deletar dados de tabelas no banco de dados de forma simples e eficaz.

Instalação
Para usar a biblioteca, basta instalar o pacote diretamente utilizando o Go Modules.

Instalação via go get:
```go get github.com/EduardoBackesdev/go-crud-lib```

# Como Usar
1. Configuração do Banco de Dados
Antes de realizar qualquer operação de CRUD, é necessário configurar a conexão com o banco de dados PostgreSQL. A função ConnectionConfig deve ser chamada uma única vez no início da aplicação do usuário para configurar os parâmetros de conexão.

Exemplo de Configuração:
```
package main

import (
	"github.com/EduardoBackesdev/go-crud-lib/lib_functions/Connection"
)

func main() {
	// Configuração da conexão - deve ser feita uma única vez na inicialização da aplicação
	Connection.ConnectionConfig("localhost", "5432", "postgres", "ADMIN", "tester")

	// Agora você pode chamar qualquer operação de CRUD depois dessa configuração
}```
