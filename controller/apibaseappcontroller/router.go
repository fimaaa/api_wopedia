package apibaseappcontroller

import (
	cfg "backend_base_app/config/env"
	"backend_base_app/gateway/apibaseappgateway"
	"backend_base_app/shared/helper"
	"backend_base_app/usecase/member/v1/getmemberv1"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Router     gin.IRouter
	Helper     helper.HTTPHelper
	Config     cfg.Config
	DataSource *apibaseappgateway.GatewayApiBaseApp
}

func (r *Controller) handlerAuthMember() gin.HandlerFunc {
	inputPort := getmemberv1.NewUsecase(r.DataSource)
	return r.authorized(inputPort)
}

func (r *Controller) handlerRefreshAuth() gin.HandlerFunc {
	inputPort := getmemberv1.NewUsecase(r.DataSource)
	return r.authorizedRefreshToken(inputPort)
}

func (r *Controller) RegisterRouter() {
	group := r.Router.Group("/api")
	r.RegisterGroupV1(group)
}

func (r *Controller) RegisterGroupV1(groupParent *gin.RouterGroup) {
	group := groupParent.Group("/v1")
	r.RegisterGroupV1Auth(group)
	r.RegisterGroupV1Member(group)
	r.RegisterGroupV1OrganizerWedding(group)
	r.RegisterGroupV1RelationUserWo(group)
	r.RegisterGroupV1EventWedding(group)
	r.RegisterGroupV1EventGuest(group)
	r.RegisterGroupV1EventTestimony(group)
}

func (r *Controller) RegisterGroupV1Auth(groupParent *gin.RouterGroup) {
	group := groupParent.Group("/auth")

	group.POST("/login", ApiBaseAppAuthMember(r))
	group.POST("/refresh", r.handlerRefreshAuth(), ApiBaseRefreshAuthMember(r))
}

func (r *Controller) RegisterGroupV1Member(groupParent *gin.RouterGroup) {
	group := groupParent.Group("/member")

	group.POST("/create", ApiBaseAppMemberCreate(r))
	group.GET("", r.handlerAuthMember(), ApiBaseAppMemberFindAll(r))
	group.GET("/:id", r.handlerAuthMember(), ApiBaseAppMemberFindOne(r))
}

func (r *Controller) RegisterGroupV1OrganizerWedding(groupParent *gin.RouterGroup) {
	group := groupParent.Group("/wedding/organizer")

	group.POST("/create", r.handlerAuthMember(), ApiBaseAppOrganizerWeddingCreate(r))
	group.GET("", r.handlerAuthMember(), ApiBaseAppOrganizerWeddingFindAll(r))
	group.GET("/:id", r.handlerAuthMember(), ApiBaseAppOrganizerWeddingFindOne(r))
	group.PUT("/", r.handlerAuthMember(), ApiBaseAppOrganizerWeddingUpdate(r))
	group.DELETE("/:id", r.handlerAuthMember(), ApiBaseAppOrganizerWeddingDeleteOne(r))
}

func (r *Controller) RegisterGroupV1RelationUserWo(groupParent *gin.RouterGroup) {
	group := groupParent.Group("/relation/user-wo")

	group.POST("/create", r.handlerAuthMember(), ApiBaseAppRelationUserWoCreate(r))
	group.GET("", r.handlerAuthMember(), ApiBaseAppRelationUserWoFindAll(r))
	group.GET("/:id", r.handlerAuthMember(), ApiBaseAppRelationUserWoFindOne(r))
	group.DELETE("/:id", r.handlerAuthMember(), ApiBaseAppRelationUserWoDeleteOne(r))
	group.DELETE("/", r.handlerAuthMember(), ApiBaseAppRelationUserWoDelete(r))
}

func (r *Controller) RegisterGroupV1EventWedding(groupParent *gin.RouterGroup) {
	group := groupParent.Group("/wedding/event")

	group.POST("/create", r.handlerAuthMember(), ApiBaseAppEventWeddingCreate(r))
	group.GET("", r.handlerAuthMember(), ApiBaseAppEventWeddingFindAll(r))
	group.GET("/:id_event", r.handlerAuthMember(), ApiBaseAppEventWeddingFindOne(r))
	group.PUT("/", r.handlerAuthMember(), ApiBaseAppEventWeddingUpdate(r))
	group.DELETE("/:id_event", r.handlerAuthMember(), ApiBaseAppEventWeddingDeleteOne(r))
}

func (r *Controller) RegisterGroupV1EventGuest(groupParent *gin.RouterGroup) {
	group := groupParent.Group("/wedding/event/:id_event/guest")

	group.POST("/create", r.handlerAuthMember(), ApiBaseAppEventGuestCreate(r))
	group.GET("", r.handlerAuthMember(), ApiBaseAppEventGuestFindAll(r))
	group.GET("/:id_guest", r.handlerAuthMember(), ApiBaseAppEventGuestFindOne(r))
	group.PUT("/", r.handlerAuthMember(), ApiBaseAppEventGuestUpdate(r))
	group.DELETE("/:id_guest", r.handlerAuthMember(), ApiBaseAppEventGuestDeleteOne(r))
}

func (r *Controller) RegisterGroupV1EventTestimony(groupParent *gin.RouterGroup) {
	group := groupParent.Group("/wedding/event/:id_event/guest/:id_guest/testimony")

	group.POST("/create", r.handlerAuthMember(), ApiBaseAppEventTestimonyCreate(r))
	group.GET("", r.handlerAuthMember(), ApiBaseAppEventTestimonyFindAll(r))
	group.GET("/:id", r.handlerAuthMember(), ApiBaseAppEventTestimonyFindOne(r))
	group.PUT("/", r.handlerAuthMember(), ApiBaseAppEventTestimonyUpdate(r))
	group.DELETE("/:id", r.handlerAuthMember(), ApiBaseAppEventTestimonyDeleteOne(r))
}
