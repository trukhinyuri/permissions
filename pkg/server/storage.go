package server

import (
	"context"

	"git.containerum.net/ch/permissions/pkg/dao"
	"git.containerum.net/ch/permissions/pkg/model"
)

type StorageActions interface {
	CreateStorage(ctx context.Context, storage model.Storage) error
	GetStorages(ctx context.Context) ([]model.Storage, error)
	UpdateStorage(ctx context.Context, name string, req model.UpdateStorageRequest) error
	DeleteStorage(ctx context.Context, name string) error
}

func (s *Server) CreateStorage(ctx context.Context, storage model.Storage) error {
	s.log.Infof("create storage %+v", storage)

	err := s.db.Transactional(func(tx *dao.DAO) error {
		return tx.CreateStorage(ctx, &storage)
	})
	return err
}

func (s *Server) GetStorages(ctx context.Context) ([]model.Storage, error) {
	s.log.Infof("get storages")

	return s.db.AllStorages(ctx)
}

func (s *Server) UpdateStorage(ctx context.Context, name string, req model.UpdateStorageRequest) error {
	s.log.Infof("update storage")

	return s.db.Transactional(func(tx *dao.DAO) error {
		storage, getErr := tx.StorageByName(ctx, name)
		if getErr != nil {
			return getErr
		}
		if req.Name != nil {
			storage.Name = *req.Name
		}
		if req.Size != nil {
			storage.Size = *req.Size
		}
		if req.Replicas != nil {
			storage.Replicas = *req.Replicas
		}
		if len(req.IPs) <= 0 {
			storage.IPs = req.IPs
		}

		return tx.UpdateStorage(ctx, name, storage)
	})
}

func (s *Server) DeleteStorage(ctx context.Context, name string) error {
	s.log.WithField("name", name).Infof("delete storage")

	return s.db.Transactional(func(tx *dao.DAO) error {
		return tx.DeleteStorage(ctx, model.Storage{Name: name})
	})
}
