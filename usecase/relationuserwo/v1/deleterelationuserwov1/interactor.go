package deleterelationuserwov1

import (
	"backend_base_app/domain/entity"
	"backend_base_app/shared/dbhelpers"
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

func (r *apibaseapporganizerweddingcreateInteractor) Execute(ctx context.Context, req entity.FindRelationUserWoData) (bool, error) {
	res := false

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		resUpdate, err := r.outport.DeleteRelationUserWoData(ctx, req)
		if err != nil {
			return err
		}
		res = resUpdate

		return nil
	})
	if err != nil {
		return false, err
	}

	return res, nil
}
