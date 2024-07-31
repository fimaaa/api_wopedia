package apibaseappcontroller

import (
	"backend_base_app/domain/domerror"
	"backend_base_app/domain/entity"
	"backend_base_app/shared/log"
	"backend_base_app/shared/util"
	"backend_base_app/usecase/eventguest/v1/createeventguestv1"
	"backend_base_app/usecase/eventguest/v1/deleteoneeventguestv1"
	"backend_base_app/usecase/eventguest/v1/editeventguestv1"
	"backend_base_app/usecase/eventguest/v1/findalleventguestv1"
	"backend_base_app/usecase/eventguest/v1/findoneeventguestv1"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ApiBaseAppEventGuestCreate(r *Controller) gin.HandlerFunc {
	var inputPort = createeventguestv1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		var req entity.EventGuestData
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

func ApiBaseAppEventGuestFindAll(r *Controller) gin.HandlerFunc {
	var inputPort = findalleventguestv1.NewUsecase(r.DataSource)

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
		var reqValue entity.FindEventGuestData
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

func ApiBaseAppEventGuestFindOne(r *Controller) gin.HandlerFunc {
	var inputPort = findoneeventguestv1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		id := c.Param("id_guest")
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

func ApiBaseAppEventGuestUpdate(r *Controller) gin.HandlerFunc {
	var inputPort = editeventguestv1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		var req entity.EditEventGuestData
		if err := c.Bind(&req); err != nil {
			newErr := domerror.FailUnmarshalRequestBodyError
			log.Error(ctx, err.Error())
			r.Helper.SendBadRequest(c, err.Error(), newErr, traceID)
			return
		}

		idEvent := c.Param("id_event")
		if idEvent == "" {
			newErr := domerror.FailUnmarshalRequestBodyError
			log.Error(ctx, newErr.Error())
			r.Helper.SendBadRequest(c, newErr.Error(), newErr, traceID)
			return
		}

		res, err := inputPort.Execute(ctx, idEvent, req)

		if err != nil {
			log.Error(ctx, err.Error())
			r.Helper.SendBadRequest(c, err.Error(), fmt.Sprintf("file err : %s", err.Error()), traceID)
			return
		}

		r.Helper.SendSuccess(c, "Success", res, traceID)
	}
}

func ApiBaseAppEventGuestDeleteOne(r *Controller) gin.HandlerFunc {
	var inputPort = deleteoneeventguestv1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		id := c.Param("id_guest")
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
