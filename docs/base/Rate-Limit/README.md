# Rate Limiting

O *rate-limiting* é uma técnica crucial para a segurança e estabilidade de APIs e serviços web. Ele consiste em limitar a quantidade de requisições que um usuário, aplicação ou endereço IP pode fazer a uma API em um determinado intervalo de tempo. Essa estratégia é essencial para prevenir abusos, garantir a disponibilidade dos serviços e proteger contra uma série de ataques, incluindo ataques de força bruta e negação de serviço (DoS).

## O que é Rate-Limiting?

*Rate-limiting* refere-se ao processo de controle do número de requisições submetidas a um servidor ou API em um período específico. Implementar *rate-limiting* significa estabelecer um limite máximo de operações que podem ser realizadas, ajudando a garantir que os recursos do servidor não sejam sobrecarregados por um excesso de demanda. Este controle é feito através de diversas técnicas e estratégias que permitem gerenciar o tráfego de entrada de forma eficaz.

## Importância do Rate-Limiting

A importância do rate-limiting se manifesta em vários aspectos:

- **Prevenção de Abusos**: Limita o impacto de usuários mal-intencionados ou scripts automatizados que tentam sobrecarregar a API com um volume excessivo de requisições, podendo levar a indisponibilidade do serviço.

- **Proteção contra Ataques de Força Bruta**: Impede que atacantes façam tentativas repetidas e rápidas para adivinhar credenciais ou explorar vulnerabilidades, já que o número de tentativas fica restrito.

- **Equidade de Acesso**: Garante que todos os usuários tenham acesso justo aos recursos, evitando que alguns consumam uma parcela desproporcional da capacidade da API e prejudiquem a experiência de outros.

- **Otimização de Recursos**: Ajuda a gerenciar a carga no servidor, distribuindo o tráfego de maneira mais uniforme e evitando picos que possam levar a falhas ou lentidão.

- **Conformidade Regulatória**: Em alguns casos, o *rate-limiting* é necessário para atender a regulamentações que exigem controle e auditoria do acesso a certos tipos de dados ou serviços.

## Técnicas de Rate-Limiting

Existem várias técnicas de *rate-limiting*, cada uma adequada a diferentes cenários e necessidades. O rate-limiting é um mecanismo essencial para controlar o fluxo de requisições a uma API, protegendo-a contra sobrecarga e ataques mal-intencionados. Entre as principais técnicas, destacam-se o limite de requisições por IP, limites baseados em usuário ou autenticação, e algoritmos como o *Leaky Bucket* e *Token Bucket*. Há também os limites dinâmicos, que se ajustam conforme o comportamento do usuário, e o *rate-limiting* geográfico, que aplica restrições baseadas na localização. Cada abordagem tem suas vantagens e desafios, permitindo uma configuração flexível para garantir a estabilidade e a segurança das APIs.

### Limite de Requisições por IP

Esta técnica é uma das mais simples e amplamente utilizadas. Consiste em limitar o número de requisições que um endereço IP pode fazer a uma API em um determinado período de tempo. Isso é eficaz para prevenir abusos por parte de usuários individuais ou scripts automatizados que operam de um único ponto de acesso.

**Como funciona**: Um contador é associado a cada IP que acessa a API. Este contador é incrementado a cada requisição feita. Se o número de requisições exceder o limite estabelecido dentro do intervalo de tempo especificado (por exemplo, 100 requisições por minuto), as requisições subsequentes são bloqueadas até que o contador seja resetado.

**Vantagens**: Simplicidade de implementação e eficácia imediata contra ataques de volume.

**Desvantagens**: Não diferencia usuários legítimos compartilhando o mesmo IP e pode ser contornado por atacantes usando múltiplos IPs.

### Limites Baseados em Usuário ou Autenticação

Quando os usuários precisam autenticar-se para usar a API, o *rate-limiting* pode ser ajustado individualmente, oferecendo uma abordagem mais granular.

**Como funciona**: A cada usuário ou *token* de autenticação é atribuído um limite específico de requisições. Este método permite que diferentes usuários tenham diferentes cotas, o que é útil em APIs que oferecem diferentes níveis de serviço ou planos de assinatura.

**Vantagens**: Permite personalização e é justo, pois cada usuário é julgado pela sua própria atividade.

**Desvantagens**: Requer autenticação e gerenciamento de estado para cada usuário, aumentando a complexidade.

### Algoritmo *Leaky Bucket*

Este algoritmo é inspirado na ideia de um balde que vaza água a uma taxa constante. É particularmente útil para lidar com tráfego irregular, permitindo pequenos *bursts* de requisições.

**Como funciona**: Um balde (ou *buffer*) enche-se com cada requisição feita. A cada intervalo de tempo, um número fixo de requisições "vaza" ou é processado. Se o balde encher, novas requisições são bloqueadas até que espaço seja liberado.

**Vantagens**: Oferece uma maneira estável de lidar com picos de requisições sem sobrecarregar o servidor.

**Desvantagens**: Pode não ser adequado para cenários onde os *bursts* de tráfego são muito grandes ou muito frequentes.


### Algoritmo *Token Bucket*

Similar ao ***Leaky Bucket***, mas com maior flexibilidade. Ele permite controlar não apenas a taxa média, mas também o tamanho do *burst* permitido.

**Como funciona**: Requisições consomem *tokens* de um balde, e novos tokens são adicionados em intervalos regulares. Se um usuário esgota seus tokens, precisa esperar até que novos tokens sejam adicionados para fazer mais requisições.

**Vantagens**: Permite *bursts* controlados e é eficaz para uma ampla gama de padrões de tráfego.

**Desvantagens**: A configuração dos parâmetros de tokens e taxa de reposição pode ser desafiadora para balancear corretamente.

### Limites Dinâmicos

Estes limites ajustam-se automaticamente com base no comportamento do usuário ou nas condições de rede.

**Como funciona**: Sistemas de monitoramento e análise de tráfego detectam padrões anormais ou picos de uso e ajustam os limites de *rate-limiting* em tempo real para mitigar potenciais ameaças ou sobrecarga.

**Vantagens**: Altamente adaptativo e eficaz contra ataques sofisticados ou mudanças repentinas no padrão de uso.

**Desvantagens**: Requer sistemas avançados de monitoramento e análise, além de ser mais complexo para implementar.

### *Rate-Limiting* Geográfico

Diferencia os limites com base na localização geográfica do usuário, o que é útil para balancear carga globalmente ou restringir regiões com alta atividade maliciosa.

**Como funciona**: A API identifica a localização geográfica do usuário (geralmente através do endereço IP) e aplica limites específicos para essa região.

**Vantagens**: Permite uma gestão eficiente do tráfego internacional e a imposição de medidas de segurança regionais.

**Desvantagens**: Pode ser contornado por atacantes usando VPNs ou proxies para mascarar sua localização real.

### *Rate-Limiting* de Infraestrutura

Utiliza componentes de infraestrutura para impor limites antes que o tráfego atinja a aplicação, protegendo-a de sobrecargas.

**Como funciona**: *Firewalls*, *proxies* e balanceadores de carga são configurados para detectar e limitar o tráfego excessivo baseado em IP, usuário ou outros critérios antes que chegue ao servidor da aplicação.

**Vantagens**: Protege a aplicação antes mesmo que o tráfego excessivo possa causar impacto.

**Desvantagens**: Pode requerer *hardware* ou serviços adicionais e a configuração pode ser complexa.

### *Rate-Limiting* Adaptativo

Implementa algoritmos que se ajustam em tempo real observando padrões de tráfego.

**Como funciona**: A API ou sistema monitora continuamente o tráfego e ajusta os limites de *rate-limiting* automaticamente para responder a mudanças na demanda ou detectar comportamentos suspeitos.

**Vantagens**: Muito flexível e responsivo, ideal para ambientes dinâmicos.

**Desvantagens**: Requer uma infraestrutura de monitoramento e análise robusta e bem integrada.

### *Headers* de *Rate-Limiting*

Envolve a comunicação do estado do *rate-limiting* para o cliente através de *headers* HTTP.

**Como funciona**: A API inclui *headers* HTTP nas respostas que informam o cliente sobre o limite atual, quantas requisições foram feitas e quando o contador será resetado.

**Vantagens**: Oferece transparência para os desenvolvedores, ajudando-os a ajustar suas requisições de acordo com os limites.

**Desvantagens**: Depende da cooperação dos clientes para ser efetivo e não protege contra clientes mal-intencionados que ignoram esses headers.


## Conclusão

Implementar o *rate-limiting* de forma eficaz não só previne abusos e ataques, mas também promove uma distribuição equitativa dos recursos, garantindo que todos os usuários tenham acesso justo e contínuo às funcionalidades oferecidas. Ao entender e aplicar as técnicas corretas de *rate-limiting*, é possível criar um ambiente digital mais robusto, responsivo e seguro.

Por fim, a escolha e combinação certa de estratégias de *rate-limiting* dependerá do entendimento profundo das necessidades específicas de cada API e do ambiente em que ela opera. Com isso em mente, a jornada para uma API segura e eficiente é uma combinação de conhecimento técnico, planejamento cuidadoso e monitoramento constante. Assim, os desenvolvedores podem não apenas proteger seus sistemas, mas também otimizar a experiência para os usuários finais, mantendo a integridade e a disponibilidade dos serviços essenciais no ecossistema digital de hoje.