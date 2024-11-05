# libCon-Postgre-GO

Esta é uma biblioteca Go para realizar operações básicas de CRUD (Create, Read, Update, Delete) em um banco de dados PostgreSQL. Ela oferece funções para inserir, selecionar, atualizar e deletar dados de tabelas no banco de dados de forma simples e eficaz.

Instalação
Para usar a biblioteca, basta instalar o pacote diretamente utilizando o Go Modules.

Instalação via go get:
```go get github.com/EduardoBackesdev/libCon-postgre-Go```

# Como Usar
1. Configuração do Banco de Dados
Antes de realizar qualquer operação de CRUD, é necessário configurar a conexão com o banco de dados PostgreSQL. A função ConnectionConfig deve ser chamada uma única vez no início da aplicação do usuário para configurar os parâmetros de conexão.

Exemplo de Configuração:
```
package main

import (
	"github.com/EduardoBackesdev/libCon-postgre-Go/lib_functions/Connection"
)

func main() {
	// Configuração da conexão - deve ser feita uma única vez na inicialização da aplicação
	Connection.ConnectionConfig("localhost", "5432", "postgres", "ADMIN", "tester")

	// Agora você pode chamar qualquer operação de CRUD depois dessa configuração
}
```
# Dica Importante:
Recomendamos que a função de configuração (ConnectionConfig) seja chamada apenas uma vez e sempre na parte superior do código de inicialização da aplicação, logo após o main(). Isso garante que a conexão com o banco de dados esteja pronta antes de qualquer operação de CRUD.

# 2. Funções Disponíveis
Aqui estão as funções principais da biblioteca para realizar operações CRUD:

# Insert
A função Insert insere dados em uma tabela do banco de dados PostgreSQL.
``` 
package main

import (
	"fmt"
	"github.com/EduardoBackesdev/libCon-postgre-Go/lib_functions"
)

func main() {
	// Exemplo de uso da função Insert
	table := "clientes"
	columns := []string{"nome", "idade"}
	values := []interface{}{"João", 30}

	lib_functions.Insert(table, columns, values)

	fmt.Println("Dados inseridos com sucesso!")
}
```
# SelectAll
A função SelectAll seleciona todos os registros de uma tabela no banco de dados PostgreSQL.
```
package main

import (
	"fmt"
	"github.com/EduardoBackesdev/libCon-postgre-Go/lib_functions"
)

func main() {
	// Exemplo de uso da função SelectAll
	table := "clientes"
	rows, err := lib_functions.SelectAll(table)
	if err != nil {
		fmt.Println("Erro ao realizar o select:", err)
		return
	}
	defer rows.Close()

	// Exibindo os dados retornados
	for rows.Next() {
		var nome string
		var idade int
		if err := rows.Scan(&nome, &idade); err != nil {
			fmt.Println("Erro ao escanear dados:", err)
			return
		}
		fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
	}
}
```
# Update
A função Update atualiza os dados de uma tabela no banco de dados PostgreSQL.
```
package main

import (
	"fmt"
	"github.com/EduardoBackesdev/libCon-postgre-Go/lib_functions"
)

func main() {
	// Exemplo de uso da função Update
	table := "clientes"
	columns := []string{"nome"}
	values := []interface{}{"Carlos"}
	lib_functions.Update(table, columns, values)

	fmt.Println("Dados atualizados com sucesso!")
}
```
# DeleteAll
A função DeleteAll deleta todos os registros de uma tabela. CUIDADO! Esta operação é irreversível.
```
package main

import (
	"fmt"
	"github.com/EduardoBackesdev/libCon-postgre-Go/lib_functions"
)

func main() {
	// Exemplo de uso da função DeleteAll
	table := "clientes"
	lib_functions.DeleteAll(table)

	fmt.Println("Todos os dados foram deletados com sucesso!")
}
```
# 3. Funções Detalhadas
```Insert```
Insere dados em uma tabela no banco de dados PostgreSQL. Para usar, passe o nome da tabela, os nomes das colunas e os valores a serem inseridos.

Parâmetros:

tabela: Nome da tabela onde os dados serão inseridos.
colunas: Lista de colunas da tabela.
valores: Lista de valores a serem inseridos nas colunas.

```SelectAll```
Seleciona todos os registros de uma tabela no banco de dados PostgreSQL.

Parâmetros:

tabela: Nome da tabela de onde os dados serão selecionados.
Retorno:

*sql.Rows: Um ponteiro para as linhas retornadas pela query, que pode ser iterado para acessar os dados.

```Update```
Atualiza dados em uma tabela PostgreSQL. Para usar, passe o nome da tabela, os nomes das colunas a serem atualizadas e os valores correspondentes.

Parâmetros:
tabela: Nome da tabela onde os dados serão atualizados.
colunas: Lista de colunas que precisam ser atualizadas.
valores: Lista de novos valores a serem atribuídos às colunas.

```DeleteAll```
Deleta todos os registros de uma tabela PostgreSQL.

Parâmetros:
tabela: Nome da tabela onde os dados serão deletados.
Atenção: Esta operação apaga todos os dados da tabela permanentemente.

# 4. Contribuições
Se você encontrou algum problema ou deseja melhorar o código, fique à vontade para enviar issues ou pull requests!


