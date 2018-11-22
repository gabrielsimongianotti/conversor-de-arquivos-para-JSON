package main

import (
	"alura/udemy/escrever_arquivos/model"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	//abre o arquivo que sera convertido em JSON
	arquivo, err := os.Open("cidade.csv")
	if err != nil {
		fmt.Println("erro ao abrir o arquivo, Erro: ", err.Error())
		return
	}
	defer arquivo.Close()
	//le o arquivo
	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("Erro ao ler o arquivo, Erro: ", err.Error())
		return
	}
	//cria o arquivo cidade.json
	arquivoJSON, err := os.Create("cidade.json")
	if err != nil {
		fmt.Println("Erro ao criar o arquivoJSON. Erro: ", err.Error())
		return
	}
	defer arquivoJSON.Close()
	//salva informações no Json
	escritor := bufio.NewWriter(arquivoJSON)
	escritor.WriteString("[\r\n")
	for indiceLinha, linha := range conteudo {
		fmt.Printf("Linha [%d] é %s\r\n", indiceLinha, linha)
		if indiceLinha != 0 {
			escritor.WriteString(",\r\n")
		}
		for indiceItem, item := range linha {
			//fmt.Printf("Item[%d} é %s\r\n",indiceItem, item)
			dados := strings.Split(item, "/")
			cidade := model.Cidade{}
			cidade.Nome = dados[0]
			cidade.Estado = dados[1]
			fmt.Printf("Cidade: %+v\r\n", cidade)

			cidadeJSON, err := json.Marshal(cidade)
			if err != nil {
				fmt.Println("Erro ao gerar o json do item", item, ". Erro: ", err.Error())
				return
			}
			escritor.WriteString(" " + string(cidadeJSON))
			if (indiceItem + 1) < len(linha) {
				escritor.WriteString(",\r\n")
			}
		}
	}
	//fecha todos os arquivos
	escritor.WriteString("\r\n]")
	escritor.Flush()
}
