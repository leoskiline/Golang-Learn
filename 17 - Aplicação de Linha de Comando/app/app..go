package app

import "github.com/urfave/cli"

// Gerar retorna uma nova aplicação de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na Internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com.br",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := BuscarIpsDoServidor(host)
	if erro != nil {
		c.App.Writer.Write([]byte("Ocorreu um erro ao buscar os IPs do servidor: " + erro.Error()))
		return
	}

	for _, ip := range ips {
		c.App.Writer.Write([]byte(ip + "\n"))
	}

}
