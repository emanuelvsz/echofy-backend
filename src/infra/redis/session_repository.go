package redis

import (
	"echofy_backend/src/core/domain/authorization"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/interfaces/repository"
	"echofy_backend/src/core/messages"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var _ repository.SessionLoader = (*redisSessionRepository)(nil)

type redisSessionRepository struct {
	connectorManager
}

func (repo redisSessionRepository) Close(key uuid.UUID) errors.Error {
	conn, err := repo.getConnection()
	if err != nil {
		return errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer repo.closeConnection(conn)

	return nil
}

func (repo redisSessionRepository) CreateSession(key uuid.UUID) errors.Error {
	conn, err := repo.getConnection()
	if err != nil {
		return errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer repo.closeConnection(conn)

	err = conn.Set(fmt.Sprintf("state:%s", key.String()), key.String(), authorization.SessionTimeout).Err()
	if err != nil {
		return errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	return nil
}

func (repo redisSessionRepository) StateExists(key uuid.UUID) (bool, errors.Error) {
	conn, err := repo.getConnection()
	if err != nil {
		return false, errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer repo.closeConnection(conn)

	stateExists, err := conn.Exists(fmt.Sprintf("state:%s", key.String())).Result()
	if err != nil {
		log.Error().Msg(err.Error())
		return false, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, err)
	}

	if stateExists <= 0 {
		return false, nil
	}

	return stateExists > 0, nil
}

func (repo redisSessionRepository) StoreAuth(key uuid.UUID, auth authorization.Authorization) errors.Error {
	conn, err := repo.getConnection()
	if err != nil {
		return errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer repo.closeConnection(conn)

	err = conn.Set(fmt.Sprintf("access:%s", key.String()), auth.Token(), auth.Expiry()).Err()
	if err != nil {
		return errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	err = conn.Set(fmt.Sprintf("refresh:%s", key.String()), auth.RefreshToken(), authorization.RefreshTokenTimeout).Err()
	if err != nil {
		return errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	return nil
}

func NewSessionRepository(manager connectorManager) *redisSessionRepository {
	return &redisSessionRepository{connectorManager: manager}
}
