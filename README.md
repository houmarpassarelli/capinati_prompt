# TESTE EM PYTHON
> O projeto consiste em manipulação de dados via prompt de comando. <br />
> Via prompt, pode-se gerar dados separados por linha ou em uma única linha.<br />
> O resultado final é gerar um arquivo CSV que pode ser usado como relatório.<br />
> No arquivo CSV, será retornado os últimos 5 dados ordenado pelo valor total do maior pra o menor.
<br /><br />
> Os código em python estão separados por versão: 2.7 e 3.x.<br />
> Já o código em Golang está em sua própria pasta em versão única.<br />
> Dentro de cada pasta já existem arquivos de exemplos com dados gerados.

## Melhorias futuras
> Uma das idéias que eu tinha foi a criação de uma interface amigável. A mesma foi realizada no outro repositório, pois como eu expliquei lá no readme do mesmo, eu fiz o processo principal para funcionar via prompt.<br />
> Como melhoria para esse processo é de além de aceitar mais formatos de arquivos com dados, poderia colocar outros tipos de divisões para cada conjunto de informação, como um pipe por exemplo.

## Decisões do projeto

> Como já explicado anteriormente, eu optei por primeiro criar um processo funcional primeiramente via prompt, para depois ter uma interface amigável.<br />
> Isso não reflete no front, mas aqui eu decidi também criar duas formas de gerar modelo de dados, um sem divisão e que sempre separa os conjuntos dados por uma contagem de 25 pontos, e outro que separa cada conjunto de dados por linha.<br >
> E por fim, sendo que existem duas formas de gerar os conjuntos de dados, eu precisei implementar no processo de geração do arquivo final, uma decisão para o usuário de qual arquivo ele quer processar, apesar que o resultado é o mesmo.<br />
> Sobre o Go, eu fiz somente mais do mesmo. A ideia era mostrar meus conhecimentos nas duas linguagens e que me serviu até experiência para mim.

## Desafios encontrados
> O único desafio mesmo foi em questão de ajustes de versão em relação ao python. Usei um macOS e por vezes você pode ter mais de uma versão instalada do python dependendo da forma que instala, pelo brew ou pelo processo instruído pela organização do python.

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