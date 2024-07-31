package findalleventtestimonyv1

import (
	"backend_base_app/domain/entity"
	"backend_base_app/shared/dbhelpers"
	"backend_base_app/shared/util"
	"context"
)

type apibaseappeventtestimonycreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &apibaseappeventtestimonycreateInteractor{
		outport: outputPort,
	}
}

func (r *apibaseappeventtestimonycreateInteractor) Execute(ctx context.Context, req entity.BaseReqFind) ([]entity.EventTestimonyData, int64, error) {
	var response = []entity.EventTestimonyData{}
	var totalRecords = int64(-1)
	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {

		//automapper
		var organizerWeddingDataRequest entity.BaseReqFind
		err := util.Automapper(req, &organizerWeddingDataRequest)
		if err != nil {
			return err
		}

		res, count, err := r.outport.FindAllEventTestimonyData(ctx, req)
		if err != nil {
			return err
		}

		for _, organizerWedding := range res {
			response = append(response, *organizerWedding)
		}

		totalRecords = count

		return nil
	})
	return response, totalRecords, err
}
