package createeventweddingv1

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

func (r *apibaseapporganizerweddingcreateInteractor) Execute(ctx context.Context, req entity.EventWeddingData) (*entity.EventWeddingData, error) {
	res := &entity.EventWeddingData{}

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {

		//automapper
		var organizerWeddingDataDataRequest entity.EventWeddingData
		err := util.Automapper(req, &organizerWeddingDataDataRequest)
		if err != nil {
			return err
		}

		log.Info(ctx, util.StructToJson(organizerWeddingDataDataRequest))

		err = r.outport.CreateEventWeddingData(ctx, organizerWeddingDataDataRequest)
		if err != nil {
			return err
		}

		res = &organizerWeddingDataDataRequest

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
