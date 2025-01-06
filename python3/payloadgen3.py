def generatePayload() -> None:
    
    data = []

    row_break = input("Deseja que a separação de dados seja em linha? (s/n): ").strip().lower()

    while True:
        reference = input("Digite o código (até 10 números): ").zfill(10)
        quantity = input("Digite a quantidade (até 5 números): ").zfill(5)
        price = input("Digite o preço (Ex.: 15.00): ").replace(".", "").zfill(10)

        row = reference + quantity + price
        if row_break == "s":
            row += "\n"

        data.append(row)

        final_step = input("Deseja adicionar mais dados? (s/n): ").strip().lower()
        if final_step != "s":
            break

    file_name = f"dados_posicionais{'_por_linha' if row_break == 's' else ''}.txt"

    with open(file_name, "w", encoding="utf-8") as file:
        file.write("".join(data))

    print(f"Arquivo {file_name} criado com sucesso!")


if __name__ == "__main__":
    generatePayload()