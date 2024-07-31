package apibaseappcontroller

import (
	"backend_base_app/domain/domerror"
	"backend_base_app/domain/entity"
	"backend_base_app/shared/log"
	"backend_base_app/shared/util"
	"backend_base_app/usecase/eventtestimony/v1/createeventtestimonyv1"
	"backend_base_app/usecase/eventtestimony/v1/deleteoneeventtestimonyv1"
	"backend_base_app/usecase/eventtestimony/v1/editeventtestimonyv1"
	"backend_base_app/usecase/eventtestimony/v1/findalleventtestimonyv1"
	"backend_base_app/usecase/eventtestimony/v1/findoneeventtestimonyv1"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ApiBaseAppEventTestimonyCreate(r *Controller) gin.HandlerFunc {
	var inputPort = createeventtestimonyv1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		var req entity.EventTestimonyData
		if err := c.Bind(&req); err != nil {
			newErr := domerror.FailUnmarshalRequestBodyError
			log.Error(ctx, err.Error())
			r.Helper.SendBadRequest(c, err.Error(), newErr, traceID)
			return
		}

		// Step 2: Read the raw JSON to extract qty_testimony
		var jsonBody map[string]interface{}
		if err := c.ShouldBindJSON(&jsonBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
			return
		}

		qtyTestimony, ok := jsonBody["qty_testimony"].(int)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "qty_testimony not found or not a number"})
			return
		}

		res, err := inputPort.Execute(ctx, req, qtyTestimony)

		if err != nil {
			log.Error(ctx, err.Error())
			r.Helper.SendBadRequest(c, err.Error(), fmt.Sprintf("file err : %s", err.Error()), traceID)
			return
		}

		r.Helper.SendSuccess(c, "Success", res, traceID)
	}
}

func ApiBaseAppEventTestimonyFindAll(r *Controller) gin.HandlerFunc {
	var inputPort = findalleventtestimonyv1.NewUsecase(r.DataSource)

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
		var reqValue entity.FindEventTestimonyData
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

func ApiBaseAppEventTestimonyFindOne(r *Controller) gin.HandlerFunc {
	var inputPort = findoneeventtestimonyv1.NewUsecase(r.DataSource)

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

func ApiBaseAppEventTestimonyUpdate(r *Controller) gin.HandlerFunc {
	var inputPort = editeventtestimonyv1.NewUsecase(r.DataSource)

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		ctx := log.Context(c.Request.Context(), traceID)

		var req entity.EditEventTestimonyData
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

func ApiBaseAppEventTestimonyDeleteOne(r *Controller) gin.HandlerFunc {
	var inputPort = deleteoneeventtestimonyv1.NewUsecase(r.DataSource)

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
