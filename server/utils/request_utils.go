package utils

import (
	"encoding/json"
	"gym/server/response"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func RequestDecoding(context *gin.Context, data interface{}) {

	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
}

func SetHeader(context *gin.Context){
	context.Writer.Header().Set("Content-Type", "application/json")

}
