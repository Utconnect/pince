package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"pince/common"
	"pince/utils"
	"time"
)

func ConsoleLoggerMiddleware(context *gin.Context) {
	log.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	start := time.Now()
	context.Next()
	end := time.Now()
	duration := time.Since(end)
	color := common.Green
	if context.Writer.Status() >= 500 {
		color = common.Red
	}

	log.Printf(color+"Request - Method: %s | Status: %d | Duration: %v | Start: %s | End: %s"+common.Reset, context.Request.Method, context.Writer.Status(), duration, start, end)
	log.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func FileLoggerMiddleware(context *gin.Context) {
	start := time.Now()
	context.Next()
	end := time.Now()
	duration := time.Since(end)
	fileName := "access_log_" + fmt.Sprintf("%02d", start.Year()) + "-" + fmt.Sprintf("%02d", int(start.Month())) + "-" + fmt.Sprintf("%02d", start.Day()) + ".log"
	utils.WriteFile(fmt.Sprintf("Request - Method: %s | Status: %d | Duration: %v | Start: %s | End: %s\n", context.Request.Method, context.Writer.Status(), duration, start, end), fileName, "./storage/logs/")
}

func GinLoggerCustomFormat() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		color := common.Green
		if param.StatusCode >= 500 {
			color = common.Red
		}

		return fmt.Sprintf(color+"%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n"+common.Reset,
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
