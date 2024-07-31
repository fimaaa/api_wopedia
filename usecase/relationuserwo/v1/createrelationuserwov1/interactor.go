package createrelationuserwov1

import (
	"backend_base_app/domain/entity"
	"backend_base_app/shared/dbhelpers"
	"backend_base_app/shared/log"
	"backend_base_app/shared/util"
	"context"
)

type apibaseapporganizerweddingcreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &apibaseapporganizerweddingcreateInteractor{
		outport: outputPort,
	}
}

func (r *apibaseapporganizerweddingcreateInteractor) Execute(ctx context.Context, req entity.RelationUserWoData) (*entity.RelationUserWoData, error) {
	res := &entity.RelationUserWoData{}

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {

		//automapper
		var relationUserWoDataDataRequest entity.RelationUserWoData
		err := util.Automapper(req, &relationUserWoDataDataRequest)
		if err != nil {
			return err
		}

		log.Info(ctx, util.StructToJson(relationUserWoDataDataRequest))

		err = r.outport.CreateRelationUserWoData(ctx, relationUserWoDataDataRequest)
		if err != nil {
			return err
		}

		res = &relationUserWoDataDataRequest

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
