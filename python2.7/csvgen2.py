# -*- coding: utf-8 -*-

import csv

def generateCSV():
    
    data = []

    file_input = raw_input("Qual arquivo deseja gerar relatório?: ").strip().lower()
    find = file_input.count("_por_linha")
    concat = ""

    if find > 0:
        concat = "_por_linha"

    with open(file_input, "r") as f:
        content = f.read()

    if "\n" in content:
        rows = content.split("\n")
        for row in rows:
            if row:
                data.append(processData(row))
    else:
        for i in range(0, len(content), 25):
            data.append(processData(content[i:i+25]))

    ordened_data = sorted(data, key=lambda x: x['total'], reverse=True)

    max_data = ordened_data[:5]

    file = "dados_processados{}.csv".format(concat)

    with open(file, "wb") as csvfile:
        field_names = ['codigo', 'quantidade', 'preco', 'total']
        writer = csv.DictWriter(csvfile, fieldnames=field_names)

        writer.writeheader()

        for item in max_data:
            writer.writerow(item)
    
    print("Arquivo {} criado com sucesso!".format(file))

def processData(payload):
    reference = payload[:10]
    quantity = int(payload[10:15])
    price = float(payload[15:]) / 100
    
    formated_price = "{:.2f}".format(price)

    total = quantity * price

    formated_total = "{:.2f}".format(total)

    return {
        'codigo': reference,
        'quantidade': quantity,
        'preco': formated_price,
        'total': formated_total
    }


generateCSV()