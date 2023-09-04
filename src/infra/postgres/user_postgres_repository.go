package postgres

type UserPostgresRepository struct {
	connectorManager
}

func NewUserPostgresRepository(manager connectorManager) *UserPostgresRepository {
	return &UserPostgresRepository{manager}
}
