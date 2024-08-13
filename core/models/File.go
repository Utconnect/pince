package models

import (
	"gorm.io/gorm"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

// swagger:model FileModel
type File struct {
	ID              uint           `gorm:"primaryKey" json:"id" binding:"numeric"`
	Name            string         `json:"name" gorm:"index"`
	Type            string         `json:"type"`
	StoragePlatform string         `json:"storage_platform"`
	Location        string         `json:"location"`
	Size            int            `json:"size"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ActualFile      *os.File       `gorm:"-"`
}

func (f *File) IModelImplement() {
	//TODO implement me
	panic("implement me")
}

func (f *File) ExtractMetaData(file multipart.FileHeader) {
	f.Name = file.Filename
	sliceFileName := strings.Split(f.Name, ".")
	f.Type = sliceFileName[len(sliceFileName)-1]
	f.Size = int(file.Size)
}

func (f *File) Id() any {
	return f.ID
}
