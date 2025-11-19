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

// GetContentLabelsInput represents the input parameters for getting content labels
type GetContentLabelsInput struct {
	PaginationInput
	ContentID string `json:"contentID" jsonschema:"required,The ID of the content"`
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

type Contents struct {
	Limit   int       `json:"limit,omitempty"`
	Size    int       `json:"size,omitempty"`
	Start   int       `json:"start,omitempty"`
	Links   Links     `json:"_links,omitempty"`
	Content []Content `json:"results,omitempty"`
}

type Content struct {
	Expandable map[string]string `json:"_expandable,omitempty"`
	Links      Links             `json:"_links,omitempty"`
	ID         string            `json:"id,omitempty"`
	Position   int               `json:"position,omitempty"`
	Status     string            `json:"status,omitempty"`
	Title      string            `json:"title,omitempty"`
	Type       string            `json:"type,omitempty"`
	Body       Body              `json:"body,omitempty"`
	Container  Container         `json:"container,omitempty"`
	Metadata   Metadata          `json:"metadata,omitempty"`
	Space      Space             `json:"space,omitempty"`
	Version    Version           `json:"version,omitempty"`
}

type Links struct {
	Base  string `json:"base,omitempty"`
	Self  string `json:"self,omitempty"`
	Webui string `json:"webui,omitempty"`
	Next  string `json:"next,omitempty"`
	Prev  string `json:"prev,omitempty"`
}

type Body struct {
	Expandable map[string]string `json:"_expandable,omitempty"`
	Storage    Storage           `json:"storage,omitempty"`
}

type Storage struct {
	Content StorageContent `json:"content,omitempty"`
	Value   string         `json:"value,omitempty"`
}

type StorageContent struct {
	Expandable map[string]string `json:"_expandable,omitempty"`
	Links      Links             `json:"_links,omitempty"`
	ID         string            `json:"id,omitempty"`
	Position   int               `json:"position,omitempty"`
	Status     string            `json:"status,omitempty"`
	Title      string            `json:"title,omitempty"`
	Type       string            `json:"type,omitempty"`
}

type Container struct {
	Expandable map[string]string `json:"_expandable,omitempty"`
	Links      Links             `json:"_links,omitempty"`
	ID         int               `json:"id,omitempty"`
	Key        string            `json:"key,omitempty"`
	Name       string            `json:"name,omitempty"`
	Status     string            `json:"status,omitempty"`
	Type       string            `json:"type,omitempty"`
}

type Metadata struct {
	Expandable map[string]string `json:"_expandable,omitempty"`
	Labels     Labels            `json:"labels,omitempty"`
}

type Labels struct {
	Links   Links         `json:"_links,omitempty"`
	Limit   int           `json:"limit,omitempty"`
	Results []interface{} `json:"results,omitempty"`
	Size    int           `json:"size,omitempty"`
	Start   int           `json:"start,omitempty"`
}

type Space struct {
	Expandable map[string]string `json:"_expandable,omitempty"`
	Links      Links             `json:"_links,omitempty"`
	ID         int               `json:"id,omitempty"`
	Key        string            `json:"key,omitempty"`
	Name       string            `json:"name,omitempty"`
	Status     string            `json:"status,omitempty"`
	Type       string            `json:"type,omitempty"`
}

type Version struct {
	Links   Links          `json:"_links,omitempty"`
	By      By             `json:"by,omitempty"`
	Content VersionContent `json:"content,omitempty"`
	Number  int            `json:"number,omitempty"`
	When    string         `json:"when,omitempty"`
}

type By struct {
	Expandable  map[string]string `json:"_expandable,omitempty"`
	DisplayName string            `json:"displayName,omitempty"`
}

type VersionContent struct {
	Expandable map[string]string `json:"_expandable,omitempty"`
	Links      Links             `json:"_links,omitempty"`
	ID         string            `json:"id,omitempty"`
	Position   int               `json:"position,omitempty"`
	Status     string            `json:"status,omitempty"`
	Title      string            `json:"title,omitempty"`
	Type       string            `json:"type,omitempty"`
}
