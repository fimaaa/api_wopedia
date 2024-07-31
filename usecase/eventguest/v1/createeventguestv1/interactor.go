package createeventguestv1

import (
	"backend_base_app/domain/entity"
	"backend_base_app/shared/dbhelpers"
	"backend_base_app/shared/log"
	"backend_base_app/shared/util"
	"context"
)

type apibaseappeventguestcreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &apibaseappeventguestcreateInteractor{
		outport: outputPort,
	}
}

func (r *apibaseappeventguestcreateInteractor) Execute(ctx context.Context, req entity.EventGuestData) (*entity.EventGuestData, error) {
	res := &entity.EventGuestData{}

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {

		//automapper
		var organizerWeddingDataDataRequest entity.EventGuestData
		err := util.Automapper(req, &organizerWeddingDataDataRequest)
		if err != nil {
			return err
		}

		log.Info(ctx, util.StructToJson(organizerWeddingDataDataRequest))

		err = r.outport.CreateEventGuestData(ctx, organizerWeddingDataDataRequest)
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
