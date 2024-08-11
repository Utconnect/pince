package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"pince/common"
	file_form "pince/core/http/requests/file-form"
	"pince/core/models"
	"pince/core/repositories"
)

type FileController struct {
	Repository repositories.FileRepository
}

func (controller *FileController) Create(context *gin.Context) {
	input := file_form.CreateFileFormRequest{}
	file := models.File{
		StoragePlatform: "local",
	}

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file.ExtractMetaData(input.File)

	file.Location = "./storage/app/uploaded-files/"
	if input.FileName != "" {
		file.Location += input.FileName + "" + file.Type
		file.Name = input.FileName
	} else {
		file.Location += file.Name
	}

	err := context.SaveUploadedFile(&input.File, file.Location)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can not save file!"})
		return
	}

	err = controller.Repository.Create(&file)
	if err != nil {
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": file})
}

func (controller *FileController) ReadData(context *gin.Context) {
	fileName := context.Query("file_name")
	file := models.File{}
	err := controller.Repository.GetByName(&file, fileName)
	if err != nil {
		context.JSON(common.ErrorHandlerHttpResponse(err))
		return
	}

	file.ActualFile, err = os.OpenFile(file.Location, os.O_RDONLY, 0644)
	//_, err = io.CopyN(context.Writer, file.ActualFile, int64(file.Size))
	if err != nil {
		return
	}

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="` + file.Name + `"`,
	}

	context.DataFromReader(http.StatusOK, int64(file.Size), "application/octet-stream", file.ActualFile, extraHeaders)
}
