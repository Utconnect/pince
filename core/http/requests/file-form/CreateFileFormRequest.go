package file_form

import (
	"mime/multipart"
	"pince/core/models"
)

// swagger:parameters CreateFeedback
type CreateFileFormRequest struct {
	// required: true
	// in: formData
	File     multipart.FileHeader `form:"file" binding:"required"`
	FileName string               `form:"file_name"`
}

type Payload struct {
	// required: true
	// in: formData
	Body models.File
}
