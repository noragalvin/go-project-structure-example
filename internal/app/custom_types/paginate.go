package custom_types

type Paginate struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
