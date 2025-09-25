package confluence

// GetSpaceInput represents the input parameters for getting a specific space
type GetSpaceInput struct {
	SpaceKey string   `json:"spaceKey" jsonschema:"required,The key of the space to retrieve"`
	Expand   []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
}

// GetContentsInSpaceInput represents the input parameters for getting contents in a space
type GetContentsInSpaceInput struct {
	PaginationInput
	SpaceKey string   `json:"spaceKey" jsonschema:"required,The key of the space"`
	Expand   []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
}

// GetContentsByTypeInput represents the input parameters for getting contents by type in a space
type GetContentsByTypeInput struct {
	PaginationInput
	SpaceKey    string   `json:"spaceKey" jsonschema:"required,The key of the space"`
	ContentType string   `json:"contentType" jsonschema:"required,The type of content to retrieve"`
	Expand      []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
}

// GetSpacesByKeyInput represents the input parameters for getting spaces by key
type GetSpacesByKeyInput struct {
	PaginationInput
	Keys               []string `json:"keys,omitempty" jsonschema:"Space keys to filter by"`
	Expand             []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
	SpaceIds           []string `json:"spaceIds,omitempty" jsonschema:"Space IDs to filter by"`
	SpaceKeys          string   `json:"spaceKeys,omitempty" jsonschema:"Space keys string"`
	SpaceId            []string `json:"spaceId,omitempty" jsonschema:"Space ID array"`
	SpaceKeySingle     string   `json:"spaceKeySingle,omitempty" jsonschema:"Single space key"`
	Type               string   `json:"type,omitempty" jsonschema:"Type of spaces to retrieve"`
	Status             string   `json:"status,omitempty" jsonschema:"Status of spaces to retrieve"`
	Label              []string `json:"label,omitempty" jsonschema:"Labels to filter by"`
	ContentLabel       []string `json:"contentLabel,omitempty" jsonschema:"Content labels to filter by"`
	Favourite          *bool    `json:"favourite,omitempty" jsonschema:"Filter favourite spaces"`
	HasRetentionPolicy *bool    `json:"hasRetentionPolicy,omitempty" jsonschema:"Filter spaces with retention policy"`
}
