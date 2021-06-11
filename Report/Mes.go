package Report

type Mes byte

const (
	Janeiro Mes = 1 + iota
	Fevereiro
	Marco
	Abril
	Maio
	Junho
	Julho
	Agosto
	Setembro
	Outubro
	Novembro
	Dezembro
)

func (estadoEnum Mes) String() string {
	return [...]string{
		"Janeiro", "Fevereiro", "Março", "Abril",
		"Maio", "Junho", "Julho", "Agosto",
		"Setembro", "Outubro", "Novembro", "Dezembro",
	}[estadoEnum-1]
}
