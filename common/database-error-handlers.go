package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ErrorHandlerHttpResponse(err error) (int, gin.H) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Record not found"}
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return http.StatusConflict, gin.H{"code": http.StatusConflict, "message": "Duplicate primary key"}
	}

	return http.StatusInternalServerError, nil
}
