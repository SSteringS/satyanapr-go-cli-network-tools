package app

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/urfave/cli"
)

// Gerar cria a aplicação de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App de Linha de Comando"
	app.Usage = "Uma aplicação de linha de comando em Go"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "buscar servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
		{
			Name:   "mx",
			Usage:  "Busca servidores MX de e-mail na internet",
			Flags:  flags,
			Action: buscarMX,
		},
		{
			Name:   "txt",
			Usage:  "Busca registros de texto (TXT) na internet",
			Flags:  flags,
			Action: buscarTxt,
		},
		{
			Name:   "status",
			Usage:  "Verifica o status HTTP de um site",
			Flags:  flags,
			Action: buscarStatus,
		},
	}
	return app

}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}

func buscarMX(c *cli.Context) {
	host := c.String("host")

	mx, erro := net.LookupMX(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range mx {
		fmt.Printf("servidor: %s com prioridade: %v\n", servidor.Host, servidor.Pref)
	}
}

func buscarTxt(c *cli.Context) {
	host := c.String("host")

	txts, erro := net.LookupTXT(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, txt := range txts {
		fmt.Println(txt)
	}
}

func buscarStatus(c *cli.Context) {
	host := c.String("host")
	resp, erro := http.Get("http://" + host)
	if erro != nil {
		log.Fatal(erro)
	}
	defer resp.Body.Close()
	fmt.Printf("Status: %s\n", resp.Status)
}
