# morse-code

> **Implementação em Go para decodificar mensagens em código Morse**

![Go](https://img.shields.io/badge/Language-Go-00ADD8.svg)

---

## 🔍 Visão Geral

O **morse-code** é uma ferramenta de linha de comando implementada em Go que recebe uma string no formato Morse e retorna o texto decodificado correspondente. A aplicação:

* **Ignora espaços extras** antes ou depois do código.
* Utiliza **1 espaço** para separar caracteres e **3 espaços** para separar palavras.
* Suporta letras (A–Z) e números (0–9).

## 📚 Contexto do Código Morse

O código Morse codifica cada caractere como uma sequência de “pontos” (`.`) e “traços” (`-`).

| Símbolo | Caractere |     | Símbolo | Caractere |
| :-----: | :-------: | :-: | :-----: | :-------: |
|   `.-`  |     A     |     | `-----` |     0     |
|  `-...` |     B     |     | `.----` |     1     |
|  `-.-.` |     C     |     | `..---` |     2     |
|  `-..`  |     D     |     | `...--` |     3     |
|   `.`   |     E     |     | `....-` |     4     |
|  `..-.` |     F     |     | `.....` |     5     |
|  `--.`  |     G     |     | `-....` |     6     |
|  `....` |     H     |     | `--...` |     7     |
|   `..`  |     I     |     | `---..` |     8     |
|  `.---` |     J     |     | `----.` |     9     |
|  `-.-`  |     K     |     |  `.-..` |     L     |
|   `--`  |     M     |     |  `-.--` |     Y     |
|   `-.`  |     N     |     |  `--..` |     Z     |

## 🚀 Descrição técnica

* **Sem dependências externas**: usa apenas bibliotecas padrão do Go.
* **Cross-platform**: compatível com Linux, macOS e Windows.
* **CI/CD**: integrado ao GitHub Actions para build e testes automáticos.
* **Alta cobertura**: testes unitários em todo o core da aplicação.

## 🛠️ Estrutura do Projeto

```text
├── cmd/cli
│   └── main.go         # entrypoint da CLI
├── internal
│   ├── application
│   │   ├── dto         # Data Transfer Objects
│   │   └── handler     # Handler da lógica de aplicação
│   └── core
│       └── domain      # Regras de negócio
├── Makefile            # scripts de build, teste e relatório
├── .github
│   └── workflows       # configuração do GitHub Actions
└── README.md           # documentação do projeto
```

## ⚙️ Requisitos

* Go 1.23+ instalado
* Make instalado (para atalho de comandos)

## 📦 Build e Execução

Use o `Makefile` para facilitar os comandos:

```bash
# Mostra opções disponíveis
make help

# Build para Linux
make build-linux
# Build para macOS
make build-macos
# Build para Windows
make build-windows
# Build para todas as plataformas
make build
```

Os binários serão gerados em `out/<os>/morse-code`.

Para executar por exemplo em um máquina Linux:

```bash
./out/linux/morse-code ".... . -.--   .--- ..- -.. ."
# Saída esperada: HEY JUDE
```

## 🔍 Testes e Cobertura

```bash
# Executa todos os testes e gera relatório de cobertura
make test-report
```

O relatório em HTML ficará na raiz em `cover.html`.