package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"pince/common"
	file_form "pince/core/http/requests/file-form"
	"pince/core/models"
	"pince/core/repositories"
	"strconv"
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
	id, _ := strconv.Atoi(context.Param("id"))

	file := models.File{
		ID: uint(id),
	}
	err := controller.Repository.GetById(&file)
	if err != nil {
		context.JSON(common.ErrorHandlerHttpResponse(err))
		return
	}

	file.ActualFile, err = os.OpenFile(file.Location, os.O_RDONLY, 0644)
	if err != nil {
		context.JSON(http.StatusInternalServerError, map[string]string{})
		return
	}

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="` + file.Name + `"`,
	}

	//binary.Write(context.Writer, binary.LittleEndian, int64(file.Size))
	context.DataFromReader(http.StatusOK, int64(file.Size), "application/octet-stream", file.ActualFile, extraHeaders)
}
