# JWT

O JWT, ou *JSON Web Token*, é uma tecnologia amplamente utilizada no contexto da segurança de aplicações para facilitar a troca de informações de forma segura entre partes. Entre as formas de validação de identidade, o JWT se destaca como uma solução robusta para a gestão de informações de autenticação.

## O que é o JWT?

O JWT é um padrão de token que permite a troca de informações entre duas partes — como um cliente e um servidor — de maneira compacta e segura. Este token é codificado em formato JSON, o que facilita sua integração com a web e aplicações modernas. A segurança do JWT é garantida através de uma assinatura digital que pode ser verificada por ambas as partes envolvidas na comunicação.

## Conceito Geral de Tokens na Autenticação

Tokens de autenticação são pequenas peças de dados que servem como prova de identidade ou sessão entre um cliente e um servidor. Quando um usuário se autentica em um sistema, ele recebe um token que será usado para acessar recursos protegidos sem a necessidade de reautenticar repetidamente. O token age como um passe seguro, contendo informações essenciais que são validadas a cada requisição para confirmar que o pedido veio de uma fonte autorizada.

Os tokens são particularmente úteis em arquiteturas onde a manutenção de estado é uma desvantagem ou impraticável, como em aplicações *stateless*. Com tokens, as aplicações não precisam armazenar o estado de autenticação em um servidor, o que contribui para melhor escalabilidade e gerenciamento de recursos.

## Vantagens do Uso de JWT

- **Segurança**: O JWT contém um mecanismo de segurança que garante que o token não foi alterado após sua criação. A assinatura digital impede modificações no token, e sua estrutura permite a inclusão de timestamps e outros métodos de verificação de validade.
- **Eficiência**: Como o JWT é autossuficiente, contendo todas as informações necessárias sobre o usuário, ele reduz a necessidade de consultas constantes ao banco de dados ou a servidores de autenticação para validar cada requisição do cliente.
- **Flexibilidade**: O JWT pode ser usado em uma variedade de aplicações, desde autenticação simples até integração complexa entre diferentes serviços, proporcionando uma forma padronizada de segurança entre plataformas diversas.

## Estrutura do JWT: *Header, Payload* e *Signature*

O *JSON Web Token* é constituído por três partes principais, cada uma com um papel crucial na segurança e eficiência do token. Essas partes são: *Header, Payload* e *Signature*.

### *Header*

O *Header* do JWT é a primeira parte do token e serve para descrever as características básicas do token, incluindo o tipo de token e o algoritmo de criptografia usado. Tipicamente, um *Header* de JWT é composto por dois campos principais:

- **typ**: Este campo especifica o tipo do token, que é JWT.
- **alg**: Este campo indica o algoritmo de assinatura utilizado para a criptografia e validação do token. Algoritmos comuns incluem HS256 (HMAC SHA-256) e RS256 (RSA SHA-256).

**Exemplo de um *Header* em JWT:**

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

Este *Header* é então codificado em Base64Url para formar a primeira parte do token.

### *Payload*

O *Payload* do JWT contém as reivindicações (*claims*) que são declarações sobre o usuário e outros dados necessários pela aplicação. As reivindicações podem ser classificadas em três tipos:

- **Reivindicações Registradas**: São um conjunto predefinido de reivindicações que não são obrigatórias, mas recomendadas. Exemplos incluem:

  - **iss** (*issuer*): Emissor do token.
  - **sub** (*subject*): Sujeito do token, geralmente o ID do usuário.
  - **aud** (*audience*): Destinatário do token, como um identificador do aplicativo.
  - **exp** (*expiration time*): Tempo de expiração do token.
  - **iat** (*issued at*): Momento em que o token foi emitido.

- **Reivindicações Públicas**: Podem ser definidas à vontade pelos desenvolvedores, desde que não colidam com as reivindicações registradas ou privadas.

- **Reivindicações Privadas**: São reivindicações criadas para compartilhar informações entre partes que concordaram com o uso dessas reivindicações.

**Exemplo de um *Payload* em JWT:**

```json
{
  "sub": "1234567890",
  "name": "João",
  "admin": true,
  "iat": 1516239022
}
```

Este *Payload* também é codificado em Base64Url para formar a segunda parte do token.

### *Signature*

A *Signature* é a parte final do JWT e é usada para verificar que o token não foi alterado após a emissão. A assinatura é gerada tomando o *Header* codificado em Base64Url, o *Payload* codificado em Base64Url, e utilizando o algoritmo especificado no *Header* para assinar esses dados com uma chave secreta (em algoritmos simétricos) ou uma chave privada (em algoritmos assimétricos).

**Exemplo de como a assinatura é gerada:**

```plaintext
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
```

A assinatura garante a integridade e a autenticidade do token, permitindo que o receptor do token valide se o token foi realmente emitido por uma fonte confiável e não foi modificado em trânsito.

## Como o JWT Funciona

O processo de uso de um JWT envolve três etapas principais: geração do token, transmissão do token e uso do token em sessões. 

### Geração do Token

A geração de um JWT é um processo que envolve a criação dos componentes *Header, Payload* e *Signature*:

1. **Definição do *Header***: O H*eader típico de um JWT inclui o tipo do token ("JWT") e o algoritmo de assinatura utilizado ("HS256", "RS256", etc.). Esse *Header* é então codificado em Base64Url.
2. **Definição do *Payload***: O *Payload* contém as reivindicações que representam as informações sobre o usuário ou outras informações necessárias pela aplicação. Reivindicações comuns incluem identificador do usuário (sub), tempo de emissão (iat) e tempo de expiração (exp). Assim como o *Header*, o *Payload* também é codificado em Base64Url.
3. **Geração da Assinatura**: A Signature é gerada pela aplicação de uma função *hash* sobre a concatenação do *Header* codificado e do *Payload* codificado, utilizando uma chave secreta (para algoritmos simétricos) ou uma chave privada (para algoritmos assimétricos). O resultado é também codificado em Base64Url.

O token final é a concatenação do *Header* codificado, do *Payload* codificado e da *Signature*, separados por pontos ('.').

### Transmissão do Token

Uma vez gerado, o JWT precisa ser transmitido do cliente para o servidor para autenticação ou de servidor para servidor para delegação de autoridade. A transmissão geralmente ocorre da seguinte maneira:

- **No cabeçalho HTTP**: O JWT é frequentemente incluído no cabeçalho HTTP de requisições feitas ao servidor. O cabeçalho mais comum usado para isso é o `Authorization`, com o token precedido pelo prefixo 'Bearer'. Por exemplo:
  
  ```http
  Authorization: Bearer <token>
  ```

Este método garante que apenas a aplicação cliente e o servidor tenham acesso ao token, protegendo-o contra interceptações no tráfego de rede, especialmente se HTTPS estiver sendo usado.

### Uso do Token em Sessões

Em aplicações *stateless*, onde o servidor não mantém o estado da sessão do usuário, o JWT desempenha um papel crucial ao manter o estado necessário dentro do próprio token. Isso é realizado através das reivindicações no *Payload*, que podem incluir informações de identificação do usuário e quaisquer outras informações de estado necessárias. Cada vez que o cliente faz uma requisição, ele envia o JWT, permitindo ao servidor:

- **Validar o Token**: O servidor verifica a *Signature* do token para assegurar que ele não foi alterado e confirma que o token ainda é válido verificando seu tempo de expiração (exp) e outras reivindicações de validade.
- **Extrair Informações do Usuário**: Uma vez validado, o servidor pode extrair informações do *Payload* para determinar a identidade do usuário e quaisquer outros dados relevantes para processar a requisição.

Dessa forma, o JWT permite que aplicações *stateless* verifiquem rapidamente a autenticidade e a autoridade de um usuário sem a necessidade de interações repetidas com um banco de dados ou um sistema de autenticação interno, aumentando a eficiência e a escalabilidade da aplicação.

# Conclusão

O JWT é uma ferramenta essencial na segurança de aplicações modernas, oferecendo uma metodologia robusta para a troca de informações de forma segura e eficiente entre partes. Por meio de sua estrutura composta por *Header, Payload* e *Signature*, o JWT não só simplifica o processo de autenticação e autorização em sistemas distribuídos como também melhora a escalabilidade e o gerenciamento de recursos ao eliminar a necessidade de armazenamento de estado no servidor.

Além disso, o JWT proporciona uma camada adicional de segurança através de mecanismos como assinatura digital e criptografia, garantindo a integridade e confidencialidade das informações trafegadas. Sua capacidade de ser incorporado em diferentes plataformas e a flexibilidade de uso em variados contextos de aplicação fazem do JWT uma escolha prevalente entre desenvolvedores que buscam soluções eficazes para gestão de identidade e segurança de dados.

Com seu design autossuficiente e a facilidade com que pode ser integrado em diferentes fluxos de autenticação e autorização, o JWT continua a ser uma peça chave no desenvolvimento de aplicações seguras e eficientes, reforçando seu papel vital na arquitetura de aplicações web e móveis contemporâneas.