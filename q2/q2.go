package q2

import (
	"errors"
	"strings"
)

//Você é um desenvolvedor de software em uma empresa financeira e está trabalhando em um sistema de folha de pagamento.
//Cada funcionário possui um ID único, nome, cargo, salário base e um conjunto de bônus mensais. Você decidiu usar uma
//struct para representar as informações de cada funcionário.
//
//Agora, você precisa implementar uma função chamada "calculateTotalSalary" que recebe como parâmetro um ponteiro para um
//objeto do tipo Employee e calcula o salário total do funcionário, considerando o salário base e a soma dos bônus
//mensais. Caso a soma dos bônus mensais seja maior que 1500.0, a titulação do funcionário deve ser atualizada com o
//prefixo "Senior".

type Employee struct {
	ID         int
	Name       string
	Title      string
	BaseSalary float64
	Bonuses    []float64
}

func CalculateTotalSalary(employee *Employee) (float64, error) {
	sum := 0.0
	if employee == nil {
		return 0, errors.New("Funcionário inválido")
	} else {
		if employee.Bonuses == nil {
			return employee.BaseSalary, nil
		} else {
			for i := 0; i < len(employee.Bonuses); i++ {
				sum += employee.Bonuses[i]
			}
			if sum > 1500 {
				employee.Title = strings.ReplaceAll(employee.Title, string(employee.Title), "Senior "+string(employee.Title))
				return employee.BaseSalary + sum, nil
			} else {
				return employee.BaseSalary + sum, nil
			}
		}
	}

}
