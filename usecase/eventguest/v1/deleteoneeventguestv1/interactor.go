package deleteoneeventguestv1

import (
	"backend_base_app/shared/dbhelpers"
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

func (r *apibaseappeventguestcreateInteractor) Execute(ctx context.Context, id string) (bool, error) {
	res := false

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		organizerWeddingDataData, err := r.outport.DeleteEventGuestData(ctx, id)
		if err != nil {
			return err
		}

		res = organizerWeddingDataData

		return nil
	})

	return res, err
}
