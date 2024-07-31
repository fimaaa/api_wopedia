package apibaseappcontroller

import (
	"backend_base_app/domain/domerror"
	"backend_base_app/domain/entity"
	"backend_base_app/shared/log"
	"backend_base_app/shared/util"
	"backend_base_app/usecase/eventwedding/v1/createeventweddingv1"
	"backend_base_app/usecase/eventwedding/v1/deleteoneeventweddingv1"
	"backend_base_app/usecase/eventwedding/v1/editeventweddingv1"
	"backend_base_app/usecase/eventwedding/v1/findalleventweddingv1"
	"backend_base_app/usecase/eventwedding/v1/findoneeventweddingv1"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ApiBaseAppEventWeddingCreate(r *Controller) gin.HandlerFunc {
	var inputPort = createeventweddingv1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		var req entity.EventWeddingData
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

func ApiBaseAppEventWeddingFindAll(r *Controller) gin.HandlerFunc {
	var inputPort = findalleventweddingv1.NewUsecase(r.DataSource)

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
		var reqValue entity.FindEventWeddingData
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

func ApiBaseAppEventWeddingFindOne(r *Controller) gin.HandlerFunc {
	var inputPort = findoneeventweddingv1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		id := c.Param("id_event")
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

func ApiBaseAppEventWeddingUpdate(r *Controller) gin.HandlerFunc {
	var inputPort = editeventweddingv1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		var req entity.EditEventWeddingData
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

func ApiBaseAppEventWeddingDeleteOne(r *Controller) gin.HandlerFunc {
	var inputPort = deleteoneeventweddingv1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		id := c.Param("id_event")
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
