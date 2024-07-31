package apibaseappcontroller

import (
	"backend_base_app/domain/domerror"
	"backend_base_app/domain/entity"
	"backend_base_app/shared/log"
	"backend_base_app/shared/util"
	"backend_base_app/usecase/relationuserwo/v1/createrelationuserwov1"
	"backend_base_app/usecase/relationuserwo/v1/deleteonerelationuserwov1"
	"backend_base_app/usecase/relationuserwo/v1/deleterelationuserwov1"
	"backend_base_app/usecase/relationuserwo/v1/findallrelationuserwov1"
	"backend_base_app/usecase/relationuserwo/v1/findonerelationuserwov1"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ApiBaseAppRelationUserWoCreate(r *Controller) gin.HandlerFunc {
	var inputPort = createrelationuserwov1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		var req entity.RelationUserWoData
		if err := c.Bind(&req); err != nil {
			newErr := domerror.FailUnmarshalRequestBodyError
			log.Error(ctx, err.Error())
			r.Helper.SendBadRequest(c, err.Error(), newErr, traceID)
			return
		}

		res, err := inputPort.Execute(ctx, req)

		if err != nil {
			log.Error(ctx, err.Error())
			r.Helper.SendBadRequest(c, err.Error(), fmt.Sprintf("file err : %s", err.Error()), traceID)
			return
		}

		r.Helper.SendSuccess(c, "Success", res, traceID)
	}
}

func ApiBaseAppRelationUserWoFindAll(r *Controller) gin.HandlerFunc {
	var inputPort = findallrelationuserwov1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		var req entity.BaseReqFind
		if err := c.BindQuery(&req); err != nil {
			newErr := domerror.FailUnmarshalRequestBodyError
			log.Error(ctx, err.Error())
			r.Helper.SendBadRequest(c, err.Error(), newErr, traceID)
			return
		}
		var reqValue entity.FindRelationUserWoData
		c.BindQuery(&reqValue)
		req.Value = reqValue

		sortByParams := make(map[string]interface{})
		// Loop through all query parameters
		for key, value := range c.Request.URL.Query() {
			// Check if the key contains "sort_by_"
			if strings.HasPrefix(key, "sort_by_") {
				trimmedKey := strings.TrimPrefix(key, "sort_by_")
				sortByParams[trimmedKey] = value
			}
		}
		req.SortBy = sortByParams

		res, count, err := inputPort.Execute(ctx, req)

		if err != nil {
			log.Error(ctx, err.Error())
			r.Helper.SendBadRequest(c, err.Error(), fmt.Sprintf("file err : %s", err.Error()), traceID)
			return
		}

		finalResponse := req.ToResponse(res, count)

		r.Helper.SendSuccess(c, "Success", finalResponse, traceID)
	}
}

func ApiBaseAppRelationUserWoFindOne(r *Controller) gin.HandlerFunc {
	var inputPort = findonerelationuserwov1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		id := c.Param("id")
		if id == "" {
			newErr := domerror.FailUnmarshalRequestBodyError
			r.Helper.SendBadRequest(c, newErr.Error(), newErr, traceID)
			return
		}

		res, err := inputPort.Execute(ctx, id)

		if err != nil {
			log.Error(ctx, err.Error())
			r.Helper.SendBadRequest(c, err.Error(), fmt.Sprintf("file err : %s", err.Error()), traceID)
			return
		}

		r.Helper.SendSuccess(c, "Success", res, traceID)
	}
}

func ApiBaseAppRelationUserWoDeleteOne(r *Controller) gin.HandlerFunc {
	var inputPort = deleteonerelationuserwov1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		id := c.Param("id")
		if id == "" {
			newErr := domerror.FailUnmarshalRequestBodyError
			r.Helper.SendBadRequest(c, newErr.Error(), newErr, traceID)
			return
		}

		res, err := inputPort.Execute(ctx, id)

		if err != nil {
			log.Error(ctx, err.Error())
			r.Helper.SendBadRequest(c, err.Error(), fmt.Sprintf("file err : %s", err.Error()), traceID)
			return
		}

		r.Helper.SendSuccess(c, "Success", res, traceID)
	}
}

func ApiBaseAppRelationUserWoDelete(r *Controller) gin.HandlerFunc {
	var inputPort = deleterelationuserwov1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		var req entity.FindRelationUserWoData
		if err := c.BindQuery(&req); err != nil {
			newErr := domerror.FailUnmarshalRequestBodyError
			log.Error(ctx, err.Error())
			r.Helper.SendBadRequest(c, err.Error(), newErr, traceID)
			return
		}
		var reqValue entity.FindRelationUserWoData
		c.BindQuery(&reqValue)

		res, err := inputPort.Execute(ctx, req)

		if err != nil {
			log.Error(ctx, err.Error())
			r.Helper.SendBadRequest(c, err.Error(), fmt.Sprintf("file err : %s", err.Error()), traceID)
			return
		}

		r.Helper.SendSuccess(c, "Success", res, traceID)
	}
}
