package messages

const (
	UnauthorizedErrorMessage          = "O usuário não está autorizado a acessar essa rota!"
	ForbiddenErrorMessage             = "Acesso negado!"
	UnexpectedErrorMessage            = "Ocorreu um erro inesperado!"
	DataSourceUnavailableErrorMessage = "A base de dados não está disponível!"
	ConflictErrorMessage              = "Não é possível realizar a operação pois existem dados conflitantes e/ou duplicados!"
	ConversionErrorMessage            = "Não foi possível realizar a operação pois ocorreu um erro de conversão!"
	FetchingDataErrorMessage          = "Não foi possível buscar as informações pois ocorreu um erro inesperado!"
	InsertingDataErrorMessage         = "Não foi possível inserir as informações pois ocorreu um erro inesperado!"
	NotFoundDataErrorMessage          = "Nenhuma informação foi encontrada para os parâmetros informados!"
)
