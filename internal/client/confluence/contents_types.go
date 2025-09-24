package confluence

import "atlassian-dc-mcp-go/internal/types"

// GetContentInput represents the input parameters for getting content
type GetContentInput struct {
	PaginationInput
	TypeParam  string   `json:"type,omitempty" jsonschema:"Filter content by type"`
	SpaceKey   string   `json:"spaceKey,omitempty" jsonschema:"Filter content by space key"`
	Title      string   `json:"title,omitempty" jsonschema:"Filter content by title"`
	Status     []string `json:"status,omitempty" jsonschema:"Filter content by status"`
	PostingDay string   `json:"postingDay,omitempty" jsonschema:"Filter content by posting day"`
	Expand     []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
}

// GetContentByIDInput represents the input parameters for getting content by ID
type GetContentByIDInput struct {
	ContentID string   `json:"contentID" jsonschema:"required,The ID of the content to retrieve"`
	Expand    []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
}

// CreateContentInput represents the input parameters for creating content
type CreateContentInput struct {
	Type      string            `json:"type" jsonschema:"required,The type of the content"`
	Title     string            `json:"title" jsonschema:"required,The title of the content"`
	Space     types.MapOutput   `json:"space" jsonschema:"required,The space information"`
	Body      types.MapOutput   `json:"body" jsonschema:"required,The body of the content"`
	Ancestors []types.MapOutput `json:"ancestors,omitempty" jsonschema:"The ancestor information"`
	Metadata  types.MapOutput   `json:"metadata,omitempty" jsonschema:"The metadata information"`
}

// UpdateContentInput represents the input parameters for updating content
type UpdateContentInput struct {
	ContentID   string          `json:"contentID" jsonschema:"required,The ID of the content to update"`
	ContentData types.MapOutput `json:"contentData" jsonschema:"required,The updated content data"`
}

// DeleteContentInput represents the input parameters for deleting content
type DeleteContentInput struct {
	ContentID string `json:"contentID" jsonschema:"required,The ID of the content to delete"`
}

// GetContentHistoryInput represents the input parameters for getting content history
type GetContentHistoryInput struct {
	ContentID string   `json:"contentID" jsonschema:"required,The ID of the content"`
	Expand    []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
}

// GetContentVersionInput represents the input parameters for getting a specific content version
type GetContentVersionInput struct {
	ContentID string   `json:"contentID" jsonschema:"required,The ID of the content"`
	Version   int      `json:"version" jsonschema:"required,The version number to retrieve"`
	Expand    []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
}

// GetContentRestrictionsInput represents the input parameters for getting content restrictions
type GetContentRestrictionsInput struct {
	ContentID string   `json:"contentID" jsonschema:"required,The ID of the content"`
	Expand    []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
}

// GetContentRestrictionsByOperationInput represents the input parameters for getting content restrictions by operation
type GetContentRestrictionsByOperationInput struct {
	ContentID string `json:"contentID" jsonschema:"required,The ID of the content"`
	Operation string `json:"operation" jsonschema:"required,The operation to get restrictions for"`
}

// GetContentLabelsInput represents the input parameters for getting content labels
type GetContentLabelsInput struct {
	PaginationInput
	ContentID string `json:"contentID" jsonschema:"required,The ID of the content"`
}

// AddContentLabelInput represents the input parameters for adding a label to content
type AddContentLabelInput struct {
	ContentID string          `json:"contentID" jsonschema:"required,The ID of the content"`
	LabelData types.MapOutput `json:"labelData" jsonschema:"required,The label data to add"`
}

// RemoveContentLabelInput represents the input parameters for removing a label from content
type RemoveContentLabelInput struct {
	ContentID string `json:"contentID" jsonschema:"required,The ID of the content"`
	Label     string `json:"label" jsonschema:"required,The label to remove"`
}

// GetContentPropertiesInput represents the input parameters for getting content properties
type GetContentPropertiesInput struct {
	PaginationInput
	ContentID string   `json:"contentID" jsonschema:"required,The ID of the content"`
	Expand    []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
}

// GetContentPropertyInput represents the input parameters for getting a specific content property
type GetContentPropertyInput struct {
	ContentID   string   `json:"contentID" jsonschema:"required,The ID of the content"`
	PropertyKey string   `json:"propertyKey" jsonschema:"required,The key of the property to retrieve"`
	Expand      []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
}

// CreateContentPropertyInput represents the input parameters for creating a content property
type CreateContentPropertyInput struct {
	ContentID    string          `json:"contentID" jsonschema:"required,The ID of the content"`
	PropertyData types.MapOutput `json:"propertyData" jsonschema:"required,The property data to create"`
}

// UpdateContentPropertyInput represents the input parameters for updating a content property
type UpdateContentPropertyInput struct {
	ContentID    string          `json:"contentID" jsonschema:"required,The ID of the content"`
	PropertyKey  string          `json:"propertyKey" jsonschema:"required,The key of the property to update"`
	PropertyData types.MapOutput `json:"propertyData" jsonschema:"required,The updated property data"`
}

// DeleteContentPropertyInput represents the input parameters for deleting a content property
type DeleteContentPropertyInput struct {
	ContentID   string `json:"contentID" jsonschema:"required,The ID of the content"`
	PropertyKey string `json:"propertyKey" jsonschema:"required,The key of the property to delete"`
}

// SearchContentInput represents the input parameters for SearchContent method.
type SearchContentInput struct {
	PaginationInput
	CQL        string   `json:"cql,omitempty" jsonschema:"The CQL query string"`
	CQLContext string   `json:"cqlcontext,omitempty" jsonschema:"The context for the CQL query"`
	Expand     []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
}

// GetCommentsInput represents the input parameters for GetComments method.
type GetCommentsInput struct {
	PaginationInput
	ContentID string   `json:"id" jsonschema:"required,The ID of the content to get comments for"`
	Expand    []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
}

// AddCommentInput represents the input parameters for AddComment method.
type AddCommentInput struct {
	ContentID   string `json:"contentId" jsonschema:"required,The ID of the content to add comment to"`
	CommentBody string `json:"commentBody" jsonschema:"required,The body of the comment"`
}

// GetAttachmentsInput represents the input parameters for GetAttachments method.
type GetAttachmentsInput struct {
	PaginationInput
	ContentID string   `json:"id" jsonschema:"required,The ID of the content to get attachments for"`
	Expand    []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
	Filename  string   `json:"filename,omitempty" jsonschema:"The filename to filter by"`
	MediaType string   `json:"mediaType,omitempty" jsonschema:"The media type to filter by"`
}

// GetExtractedTextInput represents the input parameters for GetExtractedText method.
type GetExtractedTextInput struct {
	ContentID    string `json:"contentId" jsonschema:"required,The ID of the content"`
	AttachmentID string `json:"attachmentId" jsonschema:"required,The ID of the attachment"`
}

// ScanContentBySpaceKeyInput represents the input parameters for ScanContentBySpaceKey method.
type ScanContentBySpaceKeyInput struct {
	PaginationInput
	TypeParam  string   `json:"type,omitempty" jsonschema:"The type of content to filter by"`
	SpaceKey   string   `json:"spaceKey,omitempty" jsonschema:"required,The space key to scan content for"`
	Status     []string `json:"status,omitempty" jsonschema:"The status values to filter by"`
	PostingDay string   `json:"postingDay,omitempty" jsonschema:"The posting day to filter by"`
	Expand     []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
	Cursor     string   `json:"cursor,omitempty" jsonschema:"The cursor for pagination"`
}
