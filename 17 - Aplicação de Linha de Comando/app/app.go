package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// A função Gerar vai retornar a aplicação de linha de comando prota para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando de Wallace Freire"
	app.Usage = "Busca IPs e nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "udemy.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Buscar Ips de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Buscar o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
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

	servidores, erro := net.LookupNS(host) // NS = Name Server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
