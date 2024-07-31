package deleteoneeventtestimonyv1

import (
	"backend_base_app/gateway/apibaseappgateway"
	"backend_base_app/shared/dbhelpers"
)

type Outport interface {
	apibaseappgateway.CreateEventTestimonyDataRepo
	dbhelpers.WithoutTransactionDB
}
