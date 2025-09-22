package confluence

// GetRelatedLabelsInput represents the input parameters for getting related labels
type GetRelatedLabelsInput struct {
	PaginationInput
	LabelName string `json:"labelName" jsonschema:"required,The name of the label"`
}

// GetLabelsInput represents the input parameters for getting labels
type GetLabelsInput struct {
	PaginationInput
	LabelName string `json:"labelName,omitempty" jsonschema:"The name of the label to filter by"`
	Owner     string `json:"owner,omitempty" jsonschema:"The owner of the labels"`
	Namespace string `json:"namespace,omitempty" jsonschema:"The namespace of the labels"`
	SpaceKey  string `json:"spaceKey,omitempty" jsonschema:"The key of the space to filter by"`
}
