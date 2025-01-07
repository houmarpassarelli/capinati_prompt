package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func generatePayload() {
	var data []string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Deseja que a separação de dados seja em linha? (s/n): ")
	rowBreak, _ := reader.ReadString('\n')
	rowBreak = strings.TrimSpace(strings.ToLower(rowBreak))

	for {
		fmt.Print("Digite o código (até 10 números): ")
		reference, _ := reader.ReadString('\n')
		reference = strings.TrimSpace(reference)
		reference = fmt.Sprintf("%010s", reference)

		fmt.Print("Digite a quantidade (até 5 números): ")
		quantity, _ := reader.ReadString('\n')
		quantity = strings.TrimSpace(quantity)
		quantity = fmt.Sprintf("%05s", quantity)

		fmt.Print("Digite o preço (Ex.: 15.00): ")
		price, _ := reader.ReadString('\n')
		price = strings.TrimSpace(price)
		price = strings.ReplaceAll(price, ".", "")
		price = fmt.Sprintf("%010s", price)

		row := reference + quantity + price
		if rowBreak == "s" {
			row += "\n"
		}

		data = append(data, row)

		fmt.Print("Deseja adicionar mais dados? (s/n): ")
		finalStep, _ := reader.ReadString('\n')
		finalStep = strings.TrimSpace(strings.ToLower(finalStep))
		if finalStep != "s" {
			break
		}
	}

	fileName := "dados_posicionais"
	if rowBreak == "s" {
		fileName += "_por_linha"
	}
	fileName += ".txt"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	file.WriteString(strings.Join(data, ""))

	fmt.Printf("Arquivo %s criado com sucesso!\n", fileName)
}

func main() {
	generatePayload()
}