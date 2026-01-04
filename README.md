# App de Linha de Comando em Go

Este projeto é uma aplicação de linha de comando (CLI) desenvolvida seguindo o curso **"Aprenda Golang do Zero! Desenvolva uma APLICAÇÃO COMPLETA!"** da Udemy.

O objetivo principal foi aprender a sintaxe da linguagem Go, o uso de pacotes externos e a interação com o sistema operacional.

## Funcionalidades

O projeto original do curso contempla os comandos `ip` e `servidores`. Para praticar e expandir os conhecimentos adquiridos, implementei **por conta própria** os comandos `mx`, `txt` e `status`.

### Lista de Comandos

Para utilizar, execute o programa seguido do comando e da flag `--host`.

- **ip**: Busca os endereços IP (IPv4 e IPv6) associados a um domínio na internet.
- **servidores**: Busca os servidores de nome (Name Servers - NS) responsáveis pelo domínio.
- **mx** *(Extra)*: Busca os registros de troca de e-mail (Mail Exchange) para verificar quem recebe os e-mails do domínio.
- **txt** *(Extra)*: Busca registros de texto (TXT), muito utilizados para validações de segurança (como SPF) e propriedade de domínio.
- **status** *(Extra)*: Realiza uma requisição HTTP para o site e retorna o código de status (ex: 200 OK, 404 Not Found), verificando se o serviço está online.

### Exemplo de Uso

```bash
go run main.go status --host google.com
```

## Bibliotecas Usadas

As seguintes bibliotecas (pacotes) foram utilizadas na construção desta ferramenta:

- **fmt**: Para formatação de entrada e saída (I/O).
- **log**: Para registro de erros e falhas críticas.
- **net**: Para interações de rede, especificamente consultas de DNS (Lookup).
- **net/http**: Para realizar requisições web e verificar status de sites.
- **github.com/urfave/cli**: Framework externo utilizado para facilitar a criação da estrutura da linha de comando, flags e ajuda.