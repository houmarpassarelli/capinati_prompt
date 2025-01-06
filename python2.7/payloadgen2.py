# -*- coding: utf-8 -*-

def generatePayload():
    
    data = []

    row_break = raw_input("Deseja que a separação de dados seja em linha? (s/n): ").strip().lower()
    
    while True:
        
        reference = raw_input("Digite o código (até 10 números): ").zfill(10)
        quantity = raw_input("Digite a quantidade (até 5 números): ").zfill(5)
        
        price = raw_input("Digite o preço (Ex.: 15.00): ")
        price = price.replace('.', '').zfill(10)
        
        row = reference + quantity + price
        if row_break == "s":
            row += "\n"
        
        data.append(row)
        
        final_step = raw_input('Deseja adicionar mais dados? (s/n): ').strip().lower()
        if final_step != 's':
            break
    
    file = "dados_posicionais{concat}.txt".format(concat="_por_linha" if row_break == "s" else "")

    with open("{}".format(file), "w") as f:
        f.write(''.join(data))
    
    print("Arquivo {} criado com sucesso!".format(file))


generatePayload()