package helper

type PaginationRequest struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

type PaginationResponse struct {
	Data        interface{} `json:"data"`
	CurrentPage int         `json:"current_page"`
	TotalPages  int         `json:"total_pages"`
	TotalItems  int         `json:"total_datas"`
}

func Paginate(data interface{}, page, limit int) PaginationResponse {
	// konversi data ke slice
	var sliceData []interface{}
	switch v := data.(type) {
	case []interface{}:
		sliceData = v
	default:
		// jika data bukan slice, kembalikan response kosong
		return PaginationResponse{
			Data:        []interface{}{},
			CurrentPage: page,
			TotalPages:  0,
			TotalItems:  0,
		}
	}

	totalItems := len(sliceData)
	totalPages := (totalItems + limit - 1) / limit

	// pastikan page tidak melebihi total page
	if page > totalPages {
		page = totalPages
	}

	startIndex := (page - 1) * limit
	endIndex := startIndex + limit
	if endIndex > totalItems {
		endIndex = totalItems
	}

	// ambil data yang sesuai dengan page dan limit
	paginatedData := sliceData[startIndex:endIndex]

	return PaginationResponse{
		Data:        paginatedData,
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}
}
