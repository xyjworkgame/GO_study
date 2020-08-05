package utils

import "companyIntroHandler/model"

func CalPageTotal(totalSize int32, pagination *model.Pagination) {
	if totalSize% pagination.PageSize == 0{
		pagination.PageTotal = totalSize/pagination.PageSize
	}else {
		pagination.PageTotal = (totalSize/pagination.PageSize) + 1
	}
}
