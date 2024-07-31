package deleteonerelationuserwov1

import (
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

func (r *apibaseapporganizerweddingcreateInteractor) Execute(ctx context.Context, id string) (bool, error) {
	res := false

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		organizerWeddingDataData, err := r.outport.DeleteOneRelationUserWoData(ctx, id)
		if err != nil {
			return err
		}

		res = organizerWeddingDataData

		return nil
	})

	return res, err
}
