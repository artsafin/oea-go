package codatypes

type StructuredValue struct {
	Context        string `json:"@context"`
	Type           string `json:"@type"`
	AdditionalType string `json:"additionalType"`
	Name           string `json:"name"`
	Url            string `json:"url"`
	TableId        string `json:"tableId"`
	RowId          string `json:"rowId"`
	TableUrl       string `json:"tableUrl"`
}
