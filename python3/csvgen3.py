import csv

def generate_csv():
    data = []

    file_input = input("Qual arquivo deseja gerar relatório?: ").strip().lower()
    find = "_por_linha" in file_input
    concat = "_por_linha" if find else ""

    try:
        with open(file_input, "r", encoding="utf-8") as f:
            content = f.read()
    except FileNotFoundError:
        print(f"Arquivo {file_input} não encontrado.")
        return

    if "\n" in content:
        rows = content.splitlines()
        for row in rows:
            if row.strip():
                data.append(process_data(row))
    else:
        for i in range(0, len(content), 25):
            data.append(process_data(content[i:i+25]))

    ordered_data = sorted(data, key=lambda x: x['total'], reverse=True)
    max_data = ordered_data[:5]

    output_file = f"dados_processados{concat}.csv"

    with open(output_file, "w", encoding="utf-8", newline="") as csvfile:
        field_names = ['codigo', 'quantidade', 'preco', 'total']
        writer = csv.DictWriter(csvfile, fieldnames=field_names)

        writer.writeheader()
        for item in max_data:
            writer.writerow(item)

    print(f"Arquivo {output_file} criado com sucesso!")

def process_data(payload):
    reference = payload[:10]
    quantity = int(payload[10:15])
    price = float(payload[15:]) / 100

    total = quantity * price

    return {
        'codigo': reference,
        'quantidade': quantity,
        'preco': f"{price:.2f}",
        'total': f"{total:.2f}"
    }

if __name__ == "__main__":
    generate_csv()