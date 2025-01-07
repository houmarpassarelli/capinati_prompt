package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	Codigo     string
	Quantidade int
	Preco      float64
	Total      float64
}

func main() {
	generateCSV()
}

func generateCSV() {
	var data []Data

	fmt.Print("Qual arquivo deseja gerar relatório?: ")
	var fileInput string
	fmt.Scanln(&fileInput)
	fileInput = strings.TrimSpace(strings.ToLower(fileInput))
	
	find := strings.Contains(fileInput, "_por_linha")
	concat := ""
	if find {
		concat = "_por_linha"
	}

	file, err := os.Open(fileInput)
	if err != nil {
		fmt.Printf("Arquivo %s não encontrado.\n", fileInput)
		return
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if strings.Contains(content, "\n") {
		rows := strings.Split(strings.TrimSpace(content), "\n")
		for _, row := range rows {
			if strings.TrimSpace(row) != "" {
				data = append(data, processData(row))
			}
		}
	} else {
		for i := 0; i < len(content); i += 25 {
			end := i + 25
			if end > len(content) {
				end = len(content)
			}
			data = append(data, processData(content[i:end]))
		}
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Total > data[j].Total
	})

	maxData := data
	if len(data) > 5 {
		maxData = data[:5]
	}

	outputFile := fmt.Sprintf("dados_processados%s.csv", concat)

	csvFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	fieldNames := []string{"codigo", "quantidade", "preco", "total"}
	if err := writer.Write(fieldNames); err != nil {
		fmt.Println("Erro ao escrever cabeçalhos no arquivo CSV:", err)
		return
	}

	for _, item := range maxData {
		row := []string{
			item.Codigo,
			strconv.Itoa(item.Quantidade),
			fmt.Sprintf("%.2f", item.Preco),
			fmt.Sprintf("%.2f", item.Total),
		}
		if err := writer.Write(row); err != nil {
			fmt.Println("Erro ao escrever linha no arquivo CSV:", err)
			return
		}
	}

	fmt.Printf("Arquivo %s criado com sucesso!\n", outputFile)
}

func processData(payload string) Data {
	reference := payload[:10]
	quantity, _ := strconv.Atoi(strings.TrimSpace(payload[10:15]))
	price, _ := strconv.ParseFloat(strings.TrimSpace(payload[15:]), 64)
	price /= 100
	total := float64(quantity) * price

	return Data{
		Codigo:     reference,
		Quantidade: quantity,
		Preco:      price,
		Total:      total,
	}
}
