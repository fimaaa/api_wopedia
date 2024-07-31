package deleteoneeventtestimonyv1

import (
	"backend_base_app/shared/dbhelpers"
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

func (r *apibaseappeventtestimonycreateInteractor) Execute(ctx context.Context, id string) (bool, error) {
	res := false

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		organizerWeddingDataData, err := r.outport.DeleteEventTestimonyData(ctx, id)
		if err != nil {
			return err
		}

		res = organizerWeddingDataData

		return nil
	})

	return res, err
}
