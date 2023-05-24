package q3

import "errors"

//Você está desenvolvendo um sistema de gerenciamento de estoque para uma loja. Cada produto possui um código único, nome,
//preço unitário e quantidade em estoque. Você decidiu usar uma struct para representar as informações de cada produto.
//
//Agora, você precisa implementar uma função chamada "updateStock" que recebe como parâmetro um ponteiro para um objeto do
//tipo Product e um mapa que representa as vendas realizadas, onde as chaves são os códigos dos produtos vendidos e os
//valores são as quantidades vendidas. A função deve atualizar a quantidade em estoque do produto com base nas vendas
//realizadas.
//
//Sua tarefa é escrever a função "updateStock" e usá-la para atualizar a quantidade em estoque dos produtos vendidos.

type Product struct {
	Code     string
	Name     string
	Price    float64
	Quantity int
}

func UpdateStock(product *Product, sales map[string]int) error {
	// Verifica se o produto é nulo
	if product == nil {
		return errors.New("Product is nil")
	}

	// Verifica se o mapa de vendas é nulo
	if sales == nil {
		return errors.New("Sales map is nil")
	}

	// Obtém a quantidade do produto que foi vendido
	quantitySold, ok := sales[product.Code]
	if !ok {
		return errors.New("Product code not found in sales map")
	}

	// Verifica se a quantidade do produto vendido é maior que a quantidade atual do produto em estoque
	if product.Quantity < quantitySold {
		return errors.New("Quantity of product that was sold is greater than the product's current stock quantity")
	}

	// Atualiza a quantidade de estoque do produto
	product.Quantity -= quantitySold

	// Retorna nulo se a atualização foi bem-sucedida
	return nil
}
