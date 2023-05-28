package bonus

import (
	"fmt"
)

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {

	if len(caminhos) == 0 {

		return "", fmt.Errorf("Não há caminhos.")

	}

	frequenciaDeEntrada := make(map[string]int)
	frequenciaDeSaida := make(map[string]int)

	for _, caminho := range caminhos {

		frequenciaDeEntrada[caminho[0]]++
		frequenciaDeSaida[caminho[1]]++

	}

	for nome := range frequenciaDeSaida {

		if frequenciaDeEntrada[nome] == 0 {

			return nome, nil

		}

	}

	return "", fmt.Errorf("Não há destino.")

}
