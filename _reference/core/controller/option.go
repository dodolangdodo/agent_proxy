package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/common"
	"github.com/labring/aiproxy/core/middleware"
	"github.com/labring/aiproxy/core/model"
)

// GetOptions godoc
//
//	@Summary		Get options
//	@Description	Returns a list of options
//	@Tags			option
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	middleware.APIResponse{data=map[string]string}
//	@Router			/api/option/ [get]
func GetOptions(c *gin.Context) {
	dbOptions, err := model.GetAllOption()
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	options := make(map[string]string, len(dbOptions))
	for _, option := range dbOptions {
		options[option.Key] = option.Value
	}

	middleware.SuccessResponse(c, options)
}

// GetOption godoc
//
//	@Summary		Get option
//	@Description	Returns a single option
//	@Tags			option
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			key	path		string	true	"Option key"
//	@Success		200	{object}	middleware.APIResponse{data=model.Option}
//	@Router			/api/option/{key} [get]
func GetOption(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		middleware.ErrorResponse(c, http.StatusBadRequest, "key is required")
		return
	}

	option, err := model.GetOption(key)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	middleware.SuccessResponse(c, option)
}

// UpdateOption godoc
//
//	@Summary		Update option
//	@Description	Updates a single option
//	@Tags			option
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			value	body		model.Option	true	"Option value"
//	@Success		200		{object}	middleware.APIResponse
//	@Router			/api/option/ [put]
//	@Router			/api/option/ [post]
func UpdateOption(c *gin.Context) {
	var option model.Option

	err := c.BindJSON(&option)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = model.UpdateOption(option.Key, option.Value)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	middleware.SuccessResponse(c, nil)
}

// UpdateOptionByKey godoc
//
//	@Summary		Update option by key
//	@Description	Updates a single option by key
//	@Tags			option
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			key		path		string	true	"Option key"
//	@Param			value	body		string	true	"Option value"
//	@Success		200		{object}	middleware.APIResponse
//	@Router			/api/option/{key} [put]
func UpdateOptionByKey(c *gin.Context) {
	key := c.Param("key")

	body, err := common.GetRequestBody(c.Request)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = model.UpdateOption(key, string(body))
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	middleware.SuccessResponse(c, nil)
}

// UpdateOptions godoc
//
//	@Summary		Update options
//	@Description	Updates multiple options
//	@Tags			option
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			options	body		map[string]string	true	"Options"
//	@Success		200		{object}	middleware.APIResponse
//	@Router			/api/option/batch [post]
func UpdateOptions(c *gin.Context) {
	var options map[string]string

	err := c.BindJSON(&options)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = model.UpdateOptions(options)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	middleware.SuccessResponse(c, nil)
}
