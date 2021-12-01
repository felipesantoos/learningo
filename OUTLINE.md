# Por que Go?
- Criada por Ken Thompson (Unix, B, C), Rob Pike (UTF-8) e Robert Griesemer.
- Em 2006, não tinha uma linguagem de compilação e execução rápida, fácil e eficiente.
- É uma linguagem criada para resolver as questões de performance e complexidade.
- Ótimas bibliotecas-padrão.
- Multiplataforma (cross-compile).
- Coleção de lixo automática (garbage collection).
- É uma linguagem compilada, de tipagem forte e estática.
- Tem pouquíssimas palavras reservadas.
- Recurso matador: em 2006, logo após o primeiro dual core.
    - Thread: 1mb.
    - Goroutine: 2kb.
- É OOP do jeito dela.

## Quando usar?
- Escala.
- Seviços web, redes, servers (machine learning, image processing, crypto, ...)
- Quando precisar de uma linguagem rápida, simples, fácil de aprender, e fácil de usar.
- Usada em: APIs, CLIs, microservices, libraries/framework, processamento de dados.
- É a base dos serviços de cloud e orquestração de containers.

# Como ter sucesso?
- Tenha foco.
- Seja proativo (corre atrás).
- Faça exercícios.
- Com direito.
- Tenha uma atitude positiva.
- Quando estiver procrastinando, para e vá estudar Go.
- Não copie e cole: digite.

# Onde programar?
- [Go Playground](https://go.dev/play)
- Go é uma linguagem idiomática (um formato padrão de código).

# Hello, World!
- package.indentifier
- Uma função variádica permite a passagem de um número indeterminado de parâmentros.
- ... indica uma função variádica.
- _ (blank identifier) descarta um valor retornado por uma função que não vamos utilizar.

# Gopher
- := parece uma marmota (gopher) ou o justiceiro (punisher).
- Tipagem automática.
- Usado para criar variáveis novas.
- != do operador de atribuição.
- Só funciona dentro de codeblocks.

## Terminologia
- Keywords são os termos reservados (que não podem ser usados como identificadores).
- Operadores (ex.: +, -, ×, ÷, etc.).
- Operandos (ex.: 1, 10, 73).
- Expressão é qualquer coisa que produz um resutado.
- Statement (declaração, afirmação): é uma instrução que gera uma ação. É formado por expressões.
- Package level scope: escopo ao nível do pacote.

# Var
- Uma variável criada dentro de um codeblock é `undefined` em outro.
- Variável package-level-scope são vistas por qualquer função pertencente ao pacote.

# Tipos
- Estáticos.
- Pode ser deduzido pelo compilador ou declarado explicitamente.
- Variáveis PLS não inicializadas só podem ser "atribuídas" dentro de um codeblock.
