package errors

type ErroPuppy struct {
}

func (e ErroPuppy) Error() string {
	return "Erro de comunicação com o RecipePuppy"
}

type ErroGiphy struct {
}

func (e ErroGiphy) Error() string {
	return "Erro de comunicação com o Giphy"
}
