package common

func CalcPaginationStartAndFinPoint(perPage, page int) (int, int) {
	return perPage*(page-1) + 1, perPage * page
}
