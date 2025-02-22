package deleteoneeventweddingv1

import (
	"backend_base_app/shared/dbhelpers"
	"context"
)

type apibaseappeventweddingcreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &apibaseappeventweddingcreateInteractor{
		outport: outputPort,
	}
}

func (r *apibaseappeventweddingcreateInteractor) Execute(ctx context.Context, id string) (bool, error) {
	res := false

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		organizerWeddingDataData, err := r.outport.DeleteEventWeddingData(ctx, id)
		if err != nil {
			return err
		}

		res = organizerWeddingDataData

		return nil
	})

	return res, err
}
