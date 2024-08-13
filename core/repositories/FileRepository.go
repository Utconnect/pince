package repositories

import (
	_interface "pince/core/models/interface"
	"pince/database"
)

type FileRepository struct {
	database.Connection
}

func (repository *FileRepository) Create(model _interface.IModel) error {
	return repository.Connection.GormDb.Create(model).Error
}

func (repository *FileRepository) GetById(model _interface.IModel) error {
	return repository.Connection.GormDb.First(model, model.Id()).Error
}

func (repository *FileRepository) GetByIds(models []_interface.IModel, ids []any) error {
	//TODO implement me
	panic("implement me")
}

func (repository *FileRepository) UpdateById(model *_interface.IModel, id any) error {
	//TODO implement me
	panic("implement me")
}

func (repository *FileRepository) DeleteById(model *_interface.IModel, id any) error {
	//TODO implement me
	panic("implement me")
}

func (repository *FileRepository) GetByName(model _interface.IModel, name string) error {
	return repository.Connection.GormDb.Where(map[string]any{"name": name}).First(model).Error
}
