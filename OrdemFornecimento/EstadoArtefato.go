package OrdemFornecimento

// Tipo base do enum
type EstadoArtefato byte

// Conteúdo do enum
const (
	D EstadoArtefato = iota + 1 // Deleted  | 1
	M                           // Modified | 2
	A                           // Updated  | 3
)

// Devolve uma descrição legível equivalente ao literal retornado pelo GIT
func (estadoEnum EstadoArtefato) String() string {
	return [...]string{"Remoção", "Alteração", "Criação"}[estadoEnum-1]
}

// Devolve o índice
func (estadoEnum EstadoArtefato) Index() byte {
	return byte(estadoEnum)
}

/**
- Devolve uma enumeração a partir de um literal válido
*/
func EstadoArtefatoFromString(literal string) EstadoArtefato {

	// O conteúdo deste array deve bater com o coteúdo do enum, posicionalmente inclusive
	enumKeys := []string{"D", "M", "A"}

	// Efetua uma busca linear simples e devolve a posição somando 1
	for index, key := range enumKeys {
		if literal == key {
			return EstadoArtefato(index + 1)
		}
	}

	// Retorna 0 se nada for encontrado
	return 0

}
