package continent

import (
	"github.com/tazuddinleton/go/basicrest/tables/continent"
)

type ReqHandler struct {
	table *continent.Table
}

func NewReqHandler(t *continent.Table) *ReqHandler {
	return &ReqHandler{t}
}

func (rh *ReqHandler) ReadAllContinents() (*[]continent.Continent, error) {
	return rh.table.ReadAll()
}
