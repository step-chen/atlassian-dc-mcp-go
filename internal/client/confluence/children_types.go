package confluence

// GetContentChildrenInput represents the input parameters for getting content children
type GetContentChildrenInput struct {
	ContentID     string   `json:"contentID" jsonschema:"required,The ID of the content"`
	Expand        []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
	ParentVersion string   `json:"parentVersion,omitempty" jsonschema:"The version of the parent content"`
}

// GetContentChildrenByTypeInput represents the input parameters for getting content children by type
type GetContentChildrenByTypeInput struct {
	PaginationInput
	ContentID string   `json:"contentID" jsonschema:"required,The ID of the content"`
	ChildType string   `json:"childType" jsonschema:"required,The type of child content to retrieve"`
	Expand    []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
	OrderBy   string   `json:"orderBy,omitempty" jsonschema:"Field to order results by"`
}

// GetContentCommentsInput represents the input parameters for getting content comments
type GetContentCommentsInput struct {
	PaginationInput
	ContentID     string   `json:"contentID" jsonschema:"required,The ID of the content"`
	Expand        []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
	ParentVersion string   `json:"parentVersion,omitempty" jsonschema:"The version of the parent content"`
}
