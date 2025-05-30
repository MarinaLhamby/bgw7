package main

import "fmt"

func main() {
	fmt.Println(seletorComSwitch(10))
	fmt.Println(seletorComSwitch(13))
	fmt.Println(seletorComSwitch(2))
	fmt.Println(seletorComSwitch(-1))

	fmt.Println(seletorComIf(1))
	fmt.Println(seletorComIf(0))
	fmt.Println(seletorComIf(8))
	fmt.Println(seletorComIf(12))
}

// Escolheria essa abordagem por ser mais legível
func seletorComSwitch(numMes int) (mensagem string) {
	switch numMes {
	case 1:
		mensagem = "Janeiro"
	case 2:
		mensagem = "Fevereiro"
	case 3:
		mensagem = "Março"
	case 4:
		mensagem = "Abril"
	case 5:
		mensagem = "Maio"
	case 6:
		mensagem = "Junho"
	case 7:
		mensagem = "Julho"
	case 8:
		mensagem = "Agosto"
	case 9:
		mensagem = "Setembro"
	case 10:
		mensagem = "Outubro"
	case 11:
		mensagem = "Novembro"
	case 12:
		mensagem = "Dezembro"
	default:
		mensagem = "Mês inexistente"
	}
	return
}

func seletorComIf(mes int) string {
	if mes == 1 {
		return "Janeiro"
	}
	if mes == 2 {
		return "Fevereiro"
	}
	if mes == 3 {
		return "Março"
	}
	if mes == 4 {
		return "Abril"
	}
	if mes == 5 {
		return "Maio"
	}
	if mes == 6 {
		return "Junho"
	}
	if mes == 7 {
		return "Julho"
	}
	if mes == 8 {
		return "Agosto"
	}
	if mes == 9 {
		return "Setembro"
	}
	if mes == 10 {
		return "Outubro"
	}
	if mes == 11 {
		return "Novembro"
	}
	if mes == 12 {
		return "Dezembro"
	}
	return "Mês inexistente"
}
