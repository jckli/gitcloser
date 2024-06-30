package index

type IndexResponse struct {
	Status int        `json:"status"`
	Data   *IndexData `json:"data"`
}

type IndexData struct {
	Message string `json:"message"`
}
