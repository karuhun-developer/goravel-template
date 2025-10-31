package responses

type PaginationResponse struct {
	Total       int64       `json:"total"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	NextPage    int         `json:"next_page"`
	PrevPage    int         `json:"prev_page"`
	Data        any 		`json:"data"`
}

type SimplePaginationResponse struct {
	PerPage 	int       	`json:"per_page"`
	PrevPage 	int       	`json:"prev_page"`
	NextPage 	int       	`json:"next_page"`
	Data  		any 		`json:"data"`
}

type CursorPaginationResponse struct {
	PerPage       	int         `json:"per_page"`
	NextCursor     	string      `json:"next_cursor"`
	PreviousCursor 	string      `json:"previous_cursor"`
	Data           	any 		`json:"data"`
}