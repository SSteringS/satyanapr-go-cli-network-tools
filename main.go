package main

import (
	"app_linha_de_comando/app"
	"fmt"
	"os"
)

func main() {
	fmt.Println("inicio do programa")

	aplicacao := app.Gerar()
	aplicacao.Run(os.Args)
}
