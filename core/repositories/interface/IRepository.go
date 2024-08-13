package _interface

import _interface "pince/core/models/interface"

type IRepository interface {
	Create(model _interface.IModel) error
	GetById(model _interface.IModel) error
	GetByIds(models []_interface.IModel, ids []any) error
	UpdateById(model *_interface.IModel, id any) error
	DeleteById(model *_interface.IModel, id any) error
}
