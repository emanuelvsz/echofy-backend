package response

type ErrorMessage struct {
	StatusCode       int            `json:"status_code"`
	Message          string         `json:"message"`
	InvalidFields    []InvalidField `json:"invalid_fields,omitempty"`
	DuplicatedFields []string       `json:"duplicated_fields,omitempty"`
}

type InvalidField struct {
	FieldName   string `json:"field_name"`
	Description string `json:"description"`
}
