# API5:2023 Broken Function Level Authorization

A melhor maneira de encontrar problemas de autorização em nível de função quebrada é realizar uma análise profunda do mecanismo de autorização, tendo em mente a hierarquia do usuário, as diferentes funções ou grupos no aplicativo e fazendo as seguintes perguntas:

- Um usuário normal pode acessar endpoints administrativos?
- Um usuário pode realizar ações confidenciais (por exemplo, criação, modificação ou exclusão) às quais ele não deveria ter acesso, simplesmente alterando o método HTTP (por exemplo, de GET para DELETE)?
- Um usuário do grupo X pode acessar uma função que deveria ser exposta apenas aos usuários do grupo Y, simplesmente adivinhando a URL e os parâmetros do *endpoint* (por exemplo /api/v1/users/export_all)?

Não presuma que um *endpoint* de API seja regular ou administrativo apenas com base no caminho da URL.

Embora os desenvolvedores possam optar por expor a maioria dos *endpoints* administrativos em um caminho relativo específico, como `/api/admins`, é muito comum encontrar esses *endpoints* administrativos em outros caminhos relativos, juntamente com *endpoints* regulares, como `/api/users`.

# Exploração

Esta falha lembra muito e está entre a [API1:2023 Broken Object Level Authorization](https://github.com/h41stur/api-vuln/tree/main/docs/API01-2023) e a [API3:2023 Broken Object Property Level Authorization](https://github.com/h41stur/api-vuln/tree/main/docs/API03-2023), pois se trata de uma fraquese que se aplica em ambos os casos simultaneamente. 

Nesta API, temos vários *endpoints* que se enquadram nesta vulnerabilidade, em sua maioria a aplicação das correções anteriores em um **nível de revisão geral** corrigem esta falha.

## Impacto Técnico

Uma vez que um invasor consiga acessar recursos e dados não permitidos em seu escopo de autorização, ele pode comprometer a confidencialidade, e integridade do sistema.

## Impacto de Negócio

Ao acessar informações e funcionalidades não permitidos em seu escopo de autorização, um invasor pode comprometer de forma crítica a imagem da empresa ao modificar, extrair ou divulgar estes dados. Além de possíveis sanções com órgãos reguladores como a LGPD.

## Recomendação de Correção

- Adotar a política do menor privilégio.
- Implemente um mecanismo de autorização adequado que dependa das políticas e da hierarquia do usuário.
- Use o mecanismo de autorização para verificar se o usuário logado tem acesso para executar a ação solicitada no registro em todas as funções que usam uma entrada do cliente para acessar um registro no banco de dados.

## Sugestão de Correção no Código

Para corrigir essta vulnerabilidade, recomenda-se que todos os *endpoints* sejam revistos e questionados conforme a sua descrição. Para cada ponto que não esteja em conformidade, a correção correspondente a [API1:2023 Broken Object Level Authorization](https://github.com/h41stur/api-vuln/tree/main/docs/API01-2023) ou a [API3:2023 Broken Object Property Level Authorization](https://github.com/h41stur/api-vuln/tree/main/docs/API03-2023) deve ser implementada.