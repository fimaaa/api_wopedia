package editorganizerweddingv1

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

func (r *apibaseapporganizerweddingcreateInteractor) Execute(ctx context.Context, req entity.EditOrganizerWeddingData) (*entity.OrganizerWeddingData, error) {
	res := &entity.OrganizerWeddingData{}

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		resUpdate, err := r.outport.UpdateOrganizerWeddingData(ctx, req)
		if err != nil {
			return err
		}
		res = resUpdate

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
