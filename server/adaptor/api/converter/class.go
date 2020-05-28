package converter

import (
	"github.com/k-ueki/AGPlus/server/adaptor/api/input"
	"github.com/k-ueki/AGPlus/server/domain/query"
)

func ConvertListClassParamToQuery(param *input.ListClass) *query.ListPaginationQuery {
	return &query.ListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffSet(),
	}
}
