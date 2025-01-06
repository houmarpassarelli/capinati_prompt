# TESTE EM PYTHON
> O projeto consiste em manipulação de dados via prompt. <br />
> Via prompt, pode-se gerar dados separados por linha ou em uma única linha.<br />
> O resultado final é gerar um arquivo CSV que pode ser usado como relatório.<br />
> No arquivo CSV, será retornado os últimos 5 dados ordenado pelo valor total do maior pra o menor.
<br /><br />
> Está separado em pastas separadas para cada versão: 2.7 e 3.x.
> Dentro de cada pasta já existem arquivos de exemplos com dados gerados.

## Requisitos

> Python 2.7<br />
> Python 3.x<br />

## Execução do projeto

### Executar versão 2.7

```bash
# Acesse a pasta python2.7 
$ cd python2.7

# Para gerar um arquivo de dados execute o script payloadgen2.py
$ python payloadgen2.py

# Após o arquivo de dados gerado, o arquivo pode ser processado e gerado o relatório em CSV
$ python csvgen2.py
```

### Executar versão 3.x

```bash
# Acesse a pasta python3
$ cd python3

# Para gerar um arquivo de dados execute o script payloadgen3.py
$ python payloadgen3.py

# Após o arquivo de dados gerado, o arquivo pode ser processado e gerado o relatório em CSV
$ python csvgen3.py
```