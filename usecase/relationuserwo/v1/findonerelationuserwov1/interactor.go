package findonerelationuserwov1

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

func (r *apibaseapporganizerweddingcreateInteractor) Execute(ctx context.Context, id string) (*entity.RelationUserWoData, error) {
	res := &entity.RelationUserWoData{}

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		organizerWeddingDataData, err := r.outport.FindOneRelationUserWoDataById(ctx, id)
		if err != nil {
			return err
		}

		res = organizerWeddingDataData

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
