# Elytron API - Vulnerable

Esta API faz parte do workshop da Elytron Security de desenvolvimento seguro. Ela é propositalmente vulnerável a vulnerabilidades listadas em algumas das categorias da [OWASP Top 10 API](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) e serve como base prática para implementação de correções das principais vulnerabilidades encontradas em APIs. 

Ela possui somente um CRUD de usuários e sua intenção é demonstrar como algumas vulnerabilidades existentes no OWASP Top 10 API acontecem e, além de sugerir possíveis correções, instigar o desenvolvimento de correções mais robustas.

## 🛠 Tecnologias Utilizadas

- **🔵 Go**: Linguagem base da API.
- **🐳 Docker**: Solução para desenvolvimento e execução de aplicativos em contêineres.
- **📦 MySQL**: Banco de dados relacional.

## Estrutura da API

Nossa API terá a estrutura dividida da seguinte forma:

### Pacotes Principais

- **Main**: Pacote principal que será executado;
- **Router**: Configura todas as rotas da API;
- **Controllers**: Contém todas as funções que lidam com as requisições HTTP;
- **Models**: Guarda as entidades utilizadas na aplicação;
- **Repo**: O pacote que faz a interação com o banco de dados.


### Pacotes Auxiliares

- **Config**: Lida com configurações de variáveis de ambiente;
- **Database**: Abre a conexão com o banco de dados;
- **Authentication**: Lida com login, criação e validação de *tokens*;
- **Middleware**: Uma camada que fica entre a requisição e a resposta, valida se o usuário já está autenticado;
- **Security**: Lida com *hashes* de senha e valida a autenticação no banco;
- **Responses**: Lida com as respostas da API.

## 🐳 Execução com Docker

Para executar a aplicação, utilize o Docker-Compose:

```bash
docker-compose up --build -d
```

## 📦 Criação do banco de dados

Após os containeres estarem em execução, é preciso criar o banco de dados de usuários que receberá os dados:

```bash
docker exec -i api-db sh -c 'exec mysql -uroot -p"$MYSQL_ROOT_PASSWORD"' < sql/sql.sql
```

Este comando não só cria o banco de dados, como o reseta em execuções posteriores.

Ao criar o banco de dados, 3 usuários são criados automaticamente:

- `admin@api.com:SenhaSupers3cret4!@` - Perfil admin.
- `user1@api.com:SenhaSupers3cret4!@` - Perfil de usuário.
- `user2@api.com:SenhaSupers3cret4!@` - Perfil de usuário.


## 📡 API Endpoints

### Usuários

- **📝 POST /users** - Registra um novo usuário (role deve ser **1 para admin** ou **2 para usuário comum**).

**Body:**

```json
{
    "name": "Username",
    "nick": "Apelido",
    "email": "email@email.com",
    "password": "senhasupersecreta",
    "role": <1,2>
}
```

- **👀 GET /users** - Retorna detalhes de todos os usuários.
- **👀 GET /users/:userId** - Retorna detalhes de um usuário.
- **🔧 PUT /users/:userId** - Atualiza um usuário existente.

**Body:**

```json
{
    "name": "NovoUsername",
    "nick": "NovoApelido",
    "email": "novoemail@email.com"
}
```

- **🪣 DELETE /users/:userId** - Deleta um usuário existente.
- **🔧 PUT /users/:userId/update-pass** - Atualiza a senha de um usuário existente.

**Body:**

```json
{
    "old": "senhaAntiga",
    "new": "senhaNova"
}
```

### Login

- **📝 POST /login** - Efetua o login e recebe o *token* de acesso.

**Body:**

```json
{
    "email": "email@email.com",
    "password": "senhasupersecreta"
}
```

- **📝 POST /forgot-pass** - Envia um e-mail de recuperação de senha.

**Body:**

```json
{
    "email": "email@email.com",
}
```


---

# Vulnerabilidades

Esta API propositalmente vulnerável a vulnerabilidades listadas em algumas das categorias da [OWASP Top 10 API](https://owasp.org/API-Security/editions/2023/en/0x11-t10/). Listaremos as vulnerabilidades e impactos adicionados, assim como sugestões simples de como evitá-las e até mesmo corrigí-las. O intuito é que esta API sirva como complemento para o *workshop* de desenvolvimento seguro e seja utilizada para correções mais robustas e bem elaboradas.

## Material Auxiliar

- [APIs](docs/base/APIs#apis)
- [JWT](docs/base/JWT#jwt)
- [Rate-Limiting](docs/base/Rate-Limit#rate-limiting)
- [Protocolo HTTP](docs/base/Protocolo-HTTP#protocolo-http)

## Vulnerabilidades

- [API1:2023 Broken Object Level Authorization](docs/API01-2023/)
- [API2:2023 Broken Authentication](docs/API02-2023/)
- [API3:2023 Broken Object Property Level Authorization](docs/API03-2023/)
- [API4:2023 Unrestricted Resource Consumption](docs/API04-2023/)
- [API5:2023 Broken Function Level Authorization](docs/API05-2023/)
- [API8:2023 Security Misconfiguration](docs/API08-2023/)
