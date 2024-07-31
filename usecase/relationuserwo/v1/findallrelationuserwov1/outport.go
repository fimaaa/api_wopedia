package findallrelationuserwov1

import (
	"backend_base_app/gateway/apibaseappgateway"
	"backend_base_app/shared/dbhelpers"
)

type Outport interface {
	apibaseappgateway.CreateRelationUserWoDataRepo
	dbhelpers.WithoutTransactionDB
}
