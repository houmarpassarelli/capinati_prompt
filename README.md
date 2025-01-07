# TESTE EM PYTHON
> O projeto consiste em manipulação de dados via prompt de comando. <br />
> Via prompt, pode-se gerar dados separados por linha ou em uma única linha.<br />
> O resultado final é gerar um arquivo CSV que pode ser usado como relatório.<br />
> No arquivo CSV, será retornado os últimos 5 dados ordenado pelo valor total do maior pra o menor.
<br /><br />
> Os código em python estão separados por versão: 2.7 e 3.x.<br />
> Já o código em Golang está em sua própria pasta em versão única.<br />
> Dentro de cada pasta já existem arquivos de exemplos com dados gerados.

## Requisitos

> python 2.7<br />
> python 3.x<br />
> golang (últimas versão)

## Execução do projeto

### Executar versão python 2.7

```bash
# Acesse a pasta python2.7 
$ cd python2.7

# Para gerar um arquivo de dados execute o script payloadgen2.py
$ python payloadgen2.py

# Após o arquivo de dados gerado, o arquivo pode ser processado e gerado o relatório em CSV
$ python csvgen2.py
```

### Executar versão python 3.x

```bash
# Acesse a pasta python3
$ cd python3

# Para gerar um arquivo de dados execute o script payloadgen3.py
$ python payloadgen3.py

# Após o arquivo de dados gerado, o arquivo pode ser processado e gerado o relatório em CSV
$ python csvgen3.py
```

### Executar versão golang

```bash
# Acesse a pasta golang
$ cd golang

# Para gerar um arquivo de dados execute o script payloadgen.go
$ go run payloadgen.go

# Após o arquivo de dados gerado, o arquivo pode ser processado e gerado o relatório em CSV
$ go run csvgen.go
```