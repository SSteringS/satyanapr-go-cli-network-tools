# 🌐 Network Tools CLI - Go

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Este projeto é uma aplicação de linha de comando (CLI) desenvolvida seguindo o curso **"Aprenda Golang do Zero! Desenvolva uma APLICAÇÃO COMPLETA!"** da Udemy.

O objetivo principal foi aprender a sintaxe da linguagem Go, o uso de pacotes externos e a interação com o sistema operacional.

## ✨ Funcionalidades

O projeto original do curso contempla os comandos `ip` e `servidores`. Para praticar e expandir os conhecimentos adquiridos, implementei **por conta própria** os comandos `mx`, `txt` e `status`.

### 📋 Lista de Comandos

Para utilizar, execute o programa seguido do comando e da flag `--host`.

- **ip**: Busca os endereços IP (IPv4 e IPv6) associados a um domínio na internet.
- **servidores**: Busca os servidores de nome (Name Servers - NS) responsáveis pelo domínio.
- **mx** *(Extra)*: Busca os registros de troca de e-mail (Mail Exchange) para verificar quem recebe os e-mails do domínio.
- **txt** *(Extra)*: Busca registros de texto (TXT), muito utilizados para validações de segurança (como SPF) e propriedade de domínio.
- **status** *(Extra)*: Realiza uma requisição HTTP para o site e retorna o código de status (ex: 200 OK, 404 Not Found), verificando se o serviço está online.

## 🚀 Instalação

### Pré-requisitos
- Go 1.20 ou superior

### Clonar e executar
```bash
git clone https://github.com/SSteringS/satyanapr-go-cli-network-tools.git
cd satyanapr-go-cli-network-tools
go mod download
go run main.go [comando] --host [domínio]
```

### Compilar binário
```bash
go build -o network-tools
./network-tools ip --host google.com
```

## 📝 Exemplos de Uso e Saída

### Buscar IPs
```bash
$ go run main.go ip --host google.com
142.250.185.46
2800:3f0:4001:830::200e
```

### Buscar Servidores DNS
```bash
$ go run main.go servidores --host google.com
ns1.google.com.
ns2.google.com.
ns3.google.com.
ns4.google.com.
```

### Verificar Status HTTP
```bash
$ go run main.go status --host google.com
Status: 200 OK
```

### Buscar Registros TXT
```bash
$ go run main.go txt --host google.com
v=spf1 include:_spf.google.com ~all
```

## 🗂️ Estrutura do Projeto
```
app_linha_de_comando/
├── app/
│   └── app.go          # Lógica dos comandos CLI
├── main.go             # Entry point da aplicação
├── go.mod              # Gerenciamento de dependências
├── go.sum              # Checksums das dependências
├── .gitignore          # Arquivos ignorados pelo Git
└── README.md           # Documentação do projeto
```
## 📚 Bibliotecas Usadas

As seguintes bibliotecas (pacotes) foram utilizadas na construção desta ferramenta:

- **fmt**: Para formatação de entrada e saída (I/O).
- **log**: Para registro de erros e falhas críticas.
- **net**: Para interações de rede, especificamente consultas de DNS (Lookup).
- **net/http**: Para realizar requisições web e verificar status de sites.
- **github.com/urfave/cli**: Framework externo utilizado para facilitar a criação da estrutura da linha de comando, flags e ajuda.

## 🛣️ Roadmap

Features planejadas para futuras versões:

- [ ] **ping** - Testar latência e conectividade com hosts
- [ ] **portas** - Scanner de portas abertas (80, 443, 22, etc.)
- [ ] **whois** - Informações de registro de domínio
- [ ] **ssl** - Verificação de certificados SSL/TLS
- [ ] **trace** - Traceroute para visualizar caminho até o host

## 🤝 Contribuindo

Pull requests são bem-vindos! Para mudanças importantes, abra uma issue primeiro para discutir o que você gostaria de mudar.

1. Fork o projeto
2. Crie sua feature branch (`git checkout -b feature/NovaFeature`)
3. Commit suas mudanças (`git commit -m 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/NovaFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto é open source e está disponível para fins educacionais.

## 🙏 Agradecimentos

- Curso Udemy: "Aprenda Golang do Zero! Desenvolva uma APLICAÇÃO COMPLETA!"
- Biblioteca [urfave/cli](https://github.com/urfave/cli) pela excelente ferramenta de criação de CLIs
