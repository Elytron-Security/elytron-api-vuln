# API3:2023 Broken Object Property Level Authorization

Ao permitir que um usuário acesse um objeto usando um *endpoint* de API, é importante validar se o usuário tem acesso às propriedades específicas do objeto que está tentando acessar.

Um endpoint de API é vulnerável se:

- O *endpoint* da API expõe propriedades de um objeto que são consideradas confidenciais e não devem ser lidas pelo usuário.
- O *endpoint* da API permite que um usuário altere, adicione/ou exclua o valor da propriedade de um objeto confidencial que o usuário não deveria ser capaz de acessar.

# Exploração

Muitas vezes, o simples fato de um usuário estar autenticado, não é suficiente privilégio para efetiver algumas açÕes. Nesta API, temos o exemplo do *endpoint* que cria usuários `POST /users`.

Este *endpoint* só funciona para usuários autenticados, porém, como uma das informações passadas no *body* é de extrema relevância para a integridade da aplicação `role`, é preciso que este *endpoint* seja bloqueado somente para usuários administradores. De outra forma, qualquer usuário autenticado poderia criar um usuário administrador arbitrário.

## Impacto Técnico

Uma vez que um invasor consiga acessar recursos e dados não permitidos em seu escopo de autorização, ele pode comprometer a confidencialidade, e integridade do sistema.

## Impacto de Negócio

Ao acessar informações e funcionalidades não permitidos em seu escopo de autorização, um invasor pode comprometer de forma crítica a imagem da empresa ao modificar, extrair ou divulgar estes dados. Além de possíveis sanções com órgãos reguladores como a LGPD.

## Recomendação de Correção

- Adotar a política do menor privilégio.
- Implemente um mecanismo de autorização adequado que dependa das políticas e da hierarquia do usuário.
- Use o mecanismo de autorização para verificar se o usuário logado tem acesso para executar a ação solicitada no registro em todas as funções que usam uma entrada do cliente para acessar um registro no banco de dados.

## Sugestão de Correção no Código

Podemos ver que a função `CreateToken()` no script `src/authentication/token.go` recebe o `roleId` e o insere nas *claims* do JWT. Isso significa que os privilégios do usuário já existe no *token*, mas ainda não está sendo validado.

Podemos criar uma função dentro deste mesmo script, que extrai esta informação de um *token* recebido em um *request* e o retorna:

```go
func ExtractRoleId(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		roleId, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["roleId"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return roleId, nil
	}

	return 0, errors.New("Token inválido")
}
```

Em seguida, podemos criar um modelo para receber os dados da tabela `roles` no caminho `src/models/roles.go`:

```go
package models

type Role struct {
	ID   uint64
	Role string
}
```

Agora precisamos fazer a consulta no banco de dados que trará a *role* correspondente ao usuário autenticado, com base no **roleId** contido no *token*. Esta função será criada no script `src/repos/users.go`:

```go
func (repo users) CheckRole(tokenRoleId uint64) (models.Role, error) {
	line, err := repo.db.Query(
		"select role from roles where id = ?",
		tokenRoleId,
	)
	if err != nil {
		return models.Role{}, err
	}
	defer line.Close()

	var role models.Role
	if line.Next() {
		if err = line.Scan(&role.Role); err != nil {
			return models.Role{}, err
		}
	}

	return role, nil
}
```


Em seguida, podemos ir até o módulo `src/controllers/users.go` responsável por lidar com as requisições, e, para cada função que exija a validação de privilégios (principalmente a `CreateUser()`), podemos invocar a função que criamos, e comparar seu retorno com o **role_id** recebida no *token*.

```go
	tokenRoleId, err := authentication.ExtractRoleId(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	role, err := repo.CheckRole(tokenRoleId)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if role.Role != "admin" {
		responses.Err(w, http.StatusForbidden, errors.New("você não tem privilégios para criar um usuário"))
		return
	}
```

Esta é uma solução simples, e de fácil implementação, porém pode não ser a ideal para todos os casos. É recomendado que novas implementações mais robustas sejam desenvolvidas e implementadas.