package findallrelationuserwov1

import (
	"backend_base_app/domain/entity"
	"backend_base_app/shared/dbhelpers"
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

func (r *apibaseapporganizerweddingcreateInteractor) Execute(ctx context.Context, req entity.BaseReqFind) ([]entity.RelationUserWoData, int64, error) {
	var response = []entity.RelationUserWoData{}
	var totalRecords = int64(-1)
	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {

		//automapper
		var organizerWeddingDataRequest entity.BaseReqFind
		err := util.Automapper(req, &organizerWeddingDataRequest)
		if err != nil {
			return err
		}

		res, count, err := r.outport.FindAllRelationUserWoData(ctx, req)
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
