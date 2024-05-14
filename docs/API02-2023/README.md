# API2:2023 Broken Authentication

Os *endpoints* e fluxos de autenticação são ativos que precisam ser protegidos. Além disso, "Esqueci a senha/redefinir senha" deve ser tratado da mesma forma que os mecanismos de autenticação.

Uma API é vulnerável se:

- Permite o preenchimento de credenciais onde o invasor usa força bruta com uma lista de nomes de usuário e senhas válidos.
- Permite que invasores realizem um ataque de força bruta na mesma conta de usuário, sem apresentar mecanismo de captcha/bloqueio de conta.
- Permite senhas fracas.
- Envia detalhes confidenciais de autenticação, como tokens de autenticação e senhas na URL.
- Permite que os usuários alterem seu endereço de e-mail, senha atual ou realizem qualquer outra operação confidencial sem solicitar confirmação de senha.
- Não valida a autenticidade dos tokens.
- Aceita tokens JWT não assinados/assinados fracamente (`{"alg":"none"}`)
- Não valida a data de expiração do JWT.
- Usa senhas de texto simples, não criptografadas ou com hash fraco.
- Usa chaves de criptografia fracas.


Além disso, um microsserviço é vulnerável se:

- Outros microsserviços podem acessá-lo sem autenticação.
- Usa tokens fracos ou previsíveis para impor autenticação


# Exploração

Ao criarmos um usuário, a API não faz nenhum tipo de validação de tamanho e complexidade da senha informada, fazendo com que um usuário possa insarir uma senha de um único caractere. Este tipo de fraqueza induz usuários a utilizarem senhas fracas de fácil adivinhação, além de complexidade que permite rápida descoberta em ataques de força-bruta.

## Impacto Técnico

A vulnerabilidade de política fraca de senhas pode ter uma série de impactos técnicos prejudiciais. Tais como:

- **Acesso não autorizado:** Se um atacante conseguir quebrar ou adivinhar uma senha, ele pode ganhar acesso a sistemas e dados para os quais o usuário tem permissões. Isso pode incluir acesso a sistemas críticos, dados sensíveis, recursos de rede e serviços.
- **Elevação de privilégios:** Um atacante que ganha acesso a uma conta pode conseguir elevar seus privilégios na rede, possivelmente até o nível de administrador. Isso pode permitir que eles alterem configurações ou causem outros danos.

## Impacto de Negócio

Alguns desses impactos incluem:

- **Perda de Dados:** Se um invasor ganhar acesso a informações confidenciais, ele poderá roubar ou destruir esses dados. Isso pode resultar em perdas financeiras significativas e danos à reputação da empresa.
- **Interrupção dos Negócios:** Um ataque bem-sucedido pode interromper as operações diárias, levando a uma perda de produtividade. Isso pode ser particularmente prejudicial se os sistemas críticos forem afetados.
- **Danos à Reputação:** Uma violação de segurança pode levar a uma perda de confiança dos clientes, parceiros e partes interessadas. Isso pode resultar em perda de negócios e danos à marca da empresa a longo prazo.
- **Responsabilidade Legal e Conformidade:** Dependendo da natureza dos dados comprometidos e das leis e regulamentos aplicáveis, a empresa pode enfrentar multas significativas, ações judiciais e outros custos legais.

## Recomendação de Correção

- Utilizar uma lista de senhas proibidas para utilização;
- Configurar um valor aceitável de tentativas de autenticação fracassada antes do bloqueio, tomando cuidado para não ser extremamente restritivo, permitindo que um atacante efetue um ataque de bloqueio de todos os usuários;
- Impor regras de tamanho e complexidade para criação de senhas.

## Sugestão de Correção no Código

Existem várias formas de validar o tamanho e complexidade de uma senha, com uso de funções, bibliotecas e regex. A solução sugerida, utiliza uma lógica aplicada a uma biblioteca do próprio Go. No pacote `models` podemos criar a função `passwordFormat()` no script `src/models/users.go` que faz a validação dos seguintes requisitos:

- A senha precisa ter no mínimo 8 caracteres;
- A senha precisa ter no mínimo 2 letras maiúsculas;
- A senha precisa ter no mínimo 2 letras minúsculas;
- A senha precisa ter no mínimo 2 símbolols.

```go
func (user *User) passwordFormat(password string) bool {
	if len(password) < 8 {
		return false
	}

	uppercaseCount := 0
	numberCount := 0
	symbolCount := 0

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			uppercaseCount++
		case unicode.IsDigit(char):
			numberCount++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			symbolCount++
		}
	}

	return uppercaseCount >= 2 && numberCount >= 2 && symbolCount >= 2
}
```

Logo em seguida, invocamos esta função dentro da função `validate()` que está no mesmo script:

```go
func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("O nome é obrigatório!")
	}

	if user.Nick == "" {
		return errors.New("O nick é obrigatório!")
	}

	if user.Email == "" {
		return errors.New("O email é obrigatório!")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("O email inserido é inválido!")
	}

	if step == "register" && !user.passwordFormat(user.Password) {
		return errors.New("A senha deve conter no mínimo 8 caracteres, 2 letras maiúsculas, 2 letras minúsculas e 2 símbolos!")
	}

	return nil
}
```


Esta é uma solução simples, e de fácil implementação, porém pode não ser a ideal para todos os casos. É recomendado que novas implementações mais robustas sejam desenvolvidas e implementadas.
