package notion

import "time"

// ObjectType represents the type of Notion object
type ObjectType string

const (
	ObjectTypePage     ObjectType = "page"
	ObjectTypeDatabase ObjectType = "database"
	ObjectTypeBlock    ObjectType = "block"
	ObjectTypeList     ObjectType = "list"
	ObjectTypeUser     ObjectType = "user"
	ObjectTypeComment  ObjectType = "comment"
)

// RichTextItemResponse represents a rich text item in Notion API
type RichTextItemResponse struct {
	Type        string           `json:"type"`
	Text        *TextContent     `json:"text,omitempty"`
	Mention     *MentionContent  `json:"mention,omitempty"`
	Annotations *TextAnnotations `json:"annotations,omitempty"`
	PlainText   string           `json:"plain_text,omitempty"`
	Href        *string          `json:"href,omitempty"`
	Equation    *EquationContent `json:"equation,omitempty"`
}

// TextContent represents text content in rich text
type TextContent struct {
	Content string    `json:"content"`
	Link    *LinkInfo `json:"link,omitempty"`
}

// LinkInfo represents link information
type LinkInfo struct {
	URL string `json:"url"`
}

// MentionContent represents mention content
type MentionContent struct {
	Type     string             `json:"type"`
	Database *DatabaseReference `json:"database,omitempty"`
	Date     *DateMention       `json:"date,omitempty"`
	Page     *PageReference     `json:"page,omitempty"`
	User     *UserReference     `json:"user,omitempty"`
	// Add other mention types as needed
}

// DatabaseReference represents a database reference
type DatabaseReference struct {
	ID string `json:"id"`
}

// DateMention represents a date mention
type DateMention struct {
	Start    string  `json:"start"`
	End      *string `json:"end,omitempty"`
	TimeZone *string `json:"time_zone,omitempty"`
}

// PageReference represents a page reference
type PageReference struct {
	ID string `json:"id"`
}

// UserReference represents a user reference
type UserReference struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}

// TextAnnotations represents text styling
type TextAnnotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

// EquationContent represents equation content
type EquationContent struct {
	Expression string `json:"expression"`
}

// BlockType represents the type of a block
type BlockType string

// Block types constants
const (
	BlockTypeParagraph        BlockType = "paragraph"
	BlockTypeHeading1         BlockType = "heading_1"
	BlockTypeHeading2         BlockType = "heading_2"
	BlockTypeHeading3         BlockType = "heading_3"
	BlockTypeBulletedListItem BlockType = "bulleted_list_item"
	BlockTypeNumberedListItem BlockType = "numbered_list_item"
	BlockTypeToDo             BlockType = "to_do"
	BlockTypeToggle           BlockType = "toggle"
	BlockTypeChildPage        BlockType = "child_page"
	BlockTypeChildDatabase    BlockType = "child_database"
	BlockTypeEmbed            BlockType = "embed"
	BlockTypeCallout          BlockType = "callout"
	BlockTypeQuote            BlockType = "quote"
	BlockTypeEquation         BlockType = "equation"
	BlockTypeDivider          BlockType = "divider"
	BlockTypeTableOfContents  BlockType = "table_of_contents"
	BlockTypeColumn           BlockType = "column"
	BlockTypeColumnList       BlockType = "column_list"
	BlockTypeLinkPreview      BlockType = "link_preview"
	BlockTypeSyncedBlock      BlockType = "synced_block"
	BlockTypeTemplate         BlockType = "template"
	BlockTypeLinkToPage       BlockType = "link_to_page"
	BlockTypeAudio            BlockType = "audio"
	BlockTypeBookmark         BlockType = "bookmark"
	BlockTypeBreadcrumb       BlockType = "breadcrumb"
	BlockTypeCode             BlockType = "code"
	BlockTypeFile             BlockType = "file"
	BlockTypeImage            BlockType = "image"
	BlockTypePDF              BlockType = "pdf"
	BlockTypeVideo            BlockType = "video"
	BlockTypeTable            BlockType = "table"
	BlockTypeTableRow         BlockType = "table_row"
	BlockTypeUnsupported      BlockType = "unsupported"
)

// BlockResponse represents a block response from Notion API
type BlockResponse struct {
	Object         string    `json:"object"`
	ID             string    `json:"id"`
	Type           BlockType `json:"type"`
	CreatedTime    time.Time `json:"created_time"`
	LastEditedTime time.Time `json:"last_edited_time"`
	HasChildren    bool      `json:"has_children,omitempty"`
	Archived       bool      `json:"archived,omitempty"`
	// Block type specific content
	Paragraph        *ParagraphBlock `json:"paragraph,omitempty"`
	Heading1         *HeadingBlock   `json:"heading_1,omitempty"`
	Heading2         *HeadingBlock   `json:"heading_2,omitempty"`
	Heading3         *HeadingBlock   `json:"heading_3,omitempty"`
	BulletedListItem *ListItemBlock  `json:"bulleted_list_item,omitempty"`
	NumberedListItem *ListItemBlock  `json:"numbered_list_item,omitempty"`
	ToDo             *ToDoBlock      `json:"to_do,omitempty"`
	Toggle           *ToggleBlock    `json:"toggle,omitempty"`
	ChildPage        *ChildPageBlock `json:"child_page,omitempty"`
	Image            *ImageBlock     `json:"image,omitempty"`
	Quote            *QuoteBlock     `json:"quote,omitempty"`
	Code             *CodeBlock      `json:"code,omitempty"`
	Callout          *CalloutBlock   `json:"callout,omitempty"`
	Bookmark         *BookmarkBlock  `json:"bookmark,omitempty"`
	// Add other block types as needed
}

// Block content types
type ParagraphBlock struct {
	RichText []RichTextItemResponse `json:"rich_text"`
	Color    string                 `json:"color,omitempty"`
	Children []BlockResponse        `json:"children,omitempty"`
}

type HeadingBlock struct {
	RichText     []RichTextItemResponse `json:"rich_text"`
	Color        string                 `json:"color,omitempty"`
	IsToggleable bool                   `json:"is_toggleable,omitempty"`
}

type ListItemBlock struct {
	RichText []RichTextItemResponse `json:"rich_text"`
	Color    string                 `json:"color,omitempty"`
	Children []BlockResponse        `json:"children,omitempty"`
}

type ToDoBlock struct {
	RichText []RichTextItemResponse `json:"rich_text"`
	Checked  bool                   `json:"checked"`
	Color    string                 `json:"color,omitempty"`
	Children []BlockResponse        `json:"children,omitempty"`
}

type ToggleBlock struct {
	RichText []RichTextItemResponse `json:"rich_text"`
	Color    string                 `json:"color,omitempty"`
	Children []BlockResponse        `json:"children,omitempty"`
}

type ChildPageBlock struct {
	Title string `json:"title"`
}

type ImageBlock struct {
	Type     string                 `json:"type"`
	External *ExternalFile          `json:"external,omitempty"`
	File     *InternalFile          `json:"file,omitempty"`
	Caption  []RichTextItemResponse `json:"caption,omitempty"`
}

type QuoteBlock struct {
	RichText []RichTextItemResponse `json:"rich_text"`
	Color    string                 `json:"color,omitempty"`
}

type CodeBlock struct {
	RichText []RichTextItemResponse `json:"rich_text"`
	Language string                 `json:"language,omitempty"`
}

type CalloutBlock struct {
	RichText []RichTextItemResponse `json:"rich_text"`
	Icon     *Icon                  `json:"icon,omitempty"`
	Color    string                 `json:"color,omitempty"`
}

type BookmarkBlock struct {
	URL     string                 `json:"url"`
	Caption []RichTextItemResponse `json:"caption,omitempty"`
}

// File types
type ExternalFile struct {
	URL string `json:"url"`
}

type InternalFile struct {
	URL        string    `json:"url"`
	ExpiryTime time.Time `json:"expiry_time"`
}

// Icon represents an icon
type Icon struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji,omitempty"`
}

// PageResponse represents a page response from Notion API
type PageResponse struct {
	Object         string                  `json:"object"`
	ID             string                  `json:"id"`
	CreatedTime    time.Time               `json:"created_time"`
	LastEditedTime time.Time               `json:"last_edited_time"`
	CreatedBy      *UserReference          `json:"created_by,omitempty"`
	LastEditedBy   *UserReference          `json:"last_edited_by,omitempty"`
	Cover          *Cover                  `json:"cover,omitempty"`
	Icon           *Icon                   `json:"icon,omitempty"`
	Archived       bool                    `json:"archived,omitempty"`
	InTrash        bool                    `json:"in_trash,omitempty"`
	URL            string                  `json:"url,omitempty"`
	PublicURL      *string                 `json:"public_url,omitempty"`
	Parent         Parent                  `json:"parent"`
	Properties     map[string]PageProperty `json:"properties"`
}

// Cover represents a page cover
type Cover struct {
	Type     string        `json:"type"`
	External *ExternalFile `json:"external,omitempty"`
	File     *InternalFile `json:"file,omitempty"`
}

// Parent represents a page parent
type Parent struct {
	Type       string  `json:"type"`
	DatabaseID *string `json:"database_id,omitempty"`
	PageID     *string `json:"page_id,omitempty"`
	Workspace  bool    `json:"workspace,omitempty"`
}

// PageProperty represents a page property
type PageProperty struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	// Property type specific fields
	Title          []RichTextItemResponse `json:"title,omitempty"`
	RichText       []RichTextItemResponse `json:"rich_text,omitempty"`
	Number         *float64               `json:"number,omitempty"`
	Select         *SelectOption          `json:"select,omitempty"`
	MultiSelect    []SelectOption         `json:"multi_select,omitempty"`
	Date           *DateProperty          `json:"date,omitempty"`
	People         []UserReference        `json:"people,omitempty"`
	Files          []FileProperty         `json:"files,omitempty"`
	Checkbox       bool                   `json:"checkbox,omitempty"`
	URL            *string                `json:"url,omitempty"`
	Email          *string                `json:"email,omitempty"`
	PhoneNumber    *string                `json:"phone_number,omitempty"`
	Formula        *FormulaProperty       `json:"formula,omitempty"`
	Status         *StatusProperty        `json:"status,omitempty"`
	Relation       []RelationProperty     `json:"relation,omitempty"`
	Rollup         *RollupProperty        `json:"rollup,omitempty"`
	CreatedBy      *UserReference         `json:"created_by,omitempty"`
	CreatedTime    *time.Time             `json:"created_time,omitempty"`
	LastEditedBy   *UserReference         `json:"last_edited_by,omitempty"`
	LastEditedTime *time.Time             `json:"last_edited_time,omitempty"`
}

// Property types
type SelectOption struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}

type DateProperty struct {
	Start    string  `json:"start"`
	End      *string `json:"end,omitempty"`
	TimeZone *string `json:"time_zone,omitempty"`
}

type FileProperty struct {
	Name     string        `json:"name"`
	Type     string        `json:"type"`
	External *ExternalFile `json:"external,omitempty"`
	File     *InternalFile `json:"file,omitempty"`
}

type FormulaProperty struct {
	Type    string        `json:"type"`
	String  *string       `json:"string,omitempty"`
	Number  *float64      `json:"number,omitempty"`
	Boolean *bool         `json:"boolean,omitempty"`
	Date    *DateProperty `json:"date,omitempty"`
}

type StatusProperty struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}

type RelationProperty struct {
	ID string `json:"id"`
}

type RollupProperty struct {
	Type   string        `json:"type"`
	Number *float64      `json:"number,omitempty"`
	Date   *DateProperty `json:"date,omitempty"`
	String *string       `json:"string,omitempty"`
	Array  []any         `json:"array,omitempty"`
}

// DatabaseResponse represents a database response from Notion API
type DatabaseResponse struct {
	Object         string                            `json:"object"`
	ID             string                            `json:"id"`
	CreatedTime    time.Time                         `json:"created_time"`
	LastEditedTime time.Time                         `json:"last_edited_time"`
	Title          []RichTextItemResponse            `json:"title"`
	Description    []RichTextItemResponse            `json:"description,omitempty"`
	URL            string                            `json:"url,omitempty"`
	Icon           *Icon                             `json:"icon,omitempty"`
	Cover          *Cover                            `json:"cover,omitempty"`
	Properties     map[string]DatabasePropertyConfig `json:"properties"`
	Parent         *Parent                           `json:"parent,omitempty"`
	Archived       bool                              `json:"archived,omitempty"`
	IsInline       bool                              `json:"is_inline,omitempty"`
}

// DatabasePropertyConfig represents database property configuration
type DatabasePropertyConfig struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	// Property type specific configuration
	Select      *SelectConfig      `json:"select,omitempty"`
	MultiSelect *MultiSelectConfig `json:"multi_select,omitempty"`
	Number      *NumberConfig      `json:"number,omitempty"`
	Formula     *FormulaConfig     `json:"formula,omitempty"`
	Relation    *RelationConfig    `json:"relation,omitempty"`
	Rollup      *RollupConfig      `json:"rollup,omitempty"`
	Status      *StatusConfig      `json:"status,omitempty"`
}

// Property config types
type SelectConfig struct {
	Options []SelectOption `json:"options"`
}

type MultiSelectConfig struct {
	Options []SelectOption `json:"options"`
}

type NumberConfig struct {
	Format string `json:"format"`
}

type FormulaConfig struct {
	Expression string `json:"expression"`
}

type RelationConfig struct {
	DatabaseID         string `json:"database_id"`
	SyncedPropertyName string `json:"synced_property_name,omitempty"`
	SyncedPropertyID   string `json:"synced_property_id,omitempty"`
}

type RollupConfig struct {
	RelationPropertyName string `json:"relation_property_name"`
	RelationPropertyID   string `json:"relation_property_id"`
	RollupPropertyName   string `json:"rollup_property_name"`
	RollupPropertyID     string `json:"rollup_property_id"`
	Function             string `json:"function"`
}

type StatusConfig struct {
	Options []StatusOption `json:"options"`
	Groups  []StatusGroup  `json:"groups"`
}

type StatusOption struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type StatusGroup struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Color     string   `json:"color"`
	OptionIDs []string `json:"option_ids"`
}

// ListResponse represents a list response from Notion API
type ListResponse struct {
	Object         string         `json:"object"`
	Results        []any          `json:"results"`
	NextCursor     *string        `json:"next_cursor"`
	HasMore        bool           `json:"has_more"`
	Type           string         `json:"type,omitempty"`
	PageOrDatabase map[string]any `json:"page_or_database,omitempty"`
}

// UserResponse represents a user response from Notion API
type UserResponse struct {
	Object    string         `json:"object"`
	ID        string         `json:"id"`
	Name      string         `json:"name,omitempty"`
	AvatarURL *string        `json:"avatar_url,omitempty"`
	Type      string         `json:"type,omitempty"`
	Person    *PersonInfo    `json:"person,omitempty"`
	Bot       map[string]any `json:"bot,omitempty"`
}

// PersonInfo represents person information
type PersonInfo struct {
	Email string `json:"email"`
}

// CommentResponse represents a comment response from Notion API
type CommentResponse struct {
	Object         string                 `json:"object"`
	ID             string                 `json:"id"`
	Parent         CommentParent          `json:"parent"`
	DiscussionID   string                 `json:"discussion_id"`
	CreatedTime    time.Time              `json:"created_time"`
	LastEditedTime time.Time              `json:"last_edited_time"`
	CreatedBy      UserReference          `json:"created_by"`
	RichText       []RichTextItemResponse `json:"rich_text"`
}

// CommentParent represents a comment parent
type CommentParent struct {
	Type    string  `json:"type"`
	PageID  *string `json:"page_id,omitempty"`
	BlockID *string `json:"block_id,omitempty"`
}

// Tool argument types
type AppendBlockChildrenArgs struct {
	BlockID  string          `json:"block_id"`
	Children []BlockResponse `json:"children"`
	After    *string         `json:"after,omitempty"`
	Format   string          `json:"format,omitempty"`
}

type RetrieveBlockArgs struct {
	BlockID string `json:"block_id"`
	Format  string `json:"format,omitempty"`
}

type RetrieveBlockChildrenArgs struct {
	BlockID     string  `json:"block_id"`
	StartCursor *string `json:"start_cursor,omitempty"`
	PageSize    *int    `json:"page_size,omitempty"`
	Format      string  `json:"format,omitempty"`
}

type DeleteBlockArgs struct {
	BlockID string `json:"block_id"`
	Format  string `json:"format,omitempty"`
}

type UpdateBlockArgs struct {
	BlockID string        `json:"block_id"`
	Block   BlockResponse `json:"block"`
	Format  string        `json:"format,omitempty"`
}

type RetrievePageArgs struct {
	PageID string `json:"page_id"`
	Format string `json:"format,omitempty"`
}

type UpdatePagePropertiesArgs struct {
	PageID     string         `json:"page_id"`
	Properties map[string]any `json:"properties"`
	Format     string         `json:"format,omitempty"`
}

type ListAllUsersArgs struct {
	StartCursor *string `json:"start_cursor,omitempty"`
	PageSize    *int    `json:"page_size,omitempty"`
	Format      string  `json:"format,omitempty"`
}

type RetrieveUserArgs struct {
	UserID string `json:"user_id"`
	Format string `json:"format,omitempty"`
}

type RetrieveBotUserArgs struct {
	RandomString string `json:"random_string"`
	Format       string `json:"format,omitempty"`
}

type CreateDatabaseArgs struct {
	Parent     Parent                 `json:"parent"`
	Title      []RichTextItemResponse `json:"title,omitempty"`
	Properties map[string]any         `json:"properties"`
	Format     string                 `json:"format,omitempty"`
}

type QueryDatabaseArgs struct {
	DatabaseID  string         `json:"database_id"`
	Filter      map[string]any `json:"filter,omitempty"`
	Sorts       []SortObject   `json:"sorts,omitempty"`
	StartCursor *string        `json:"start_cursor,omitempty"`
	PageSize    *int           `json:"page_size,omitempty"`
	Format      string         `json:"format,omitempty"`
}

type SortObject struct {
	Property  *string `json:"property,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Direction string  `json:"direction"`
}

type RetrieveDatabaseArgs struct {
	DatabaseID string `json:"database_id"`
	Format     string `json:"format,omitempty"`
}

type UpdateDatabaseArgs struct {
	DatabaseID  string                 `json:"database_id"`
	Title       []RichTextItemResponse `json:"title,omitempty"`
	Description []RichTextItemResponse `json:"description,omitempty"`
	Properties  map[string]any         `json:"properties,omitempty"`
	Format      string                 `json:"format,omitempty"`
}

type CreateDatabaseItemArgs struct {
	DatabaseID string         `json:"database_id"`
	Properties map[string]any `json:"properties"`
	Format     string         `json:"format,omitempty"`
}

type CreateCommentArgs struct {
	Parent       *CommentParentInput    `json:"parent,omitempty"`
	DiscussionID *string                `json:"discussion_id,omitempty"`
	RichText     []RichTextItemResponse `json:"rich_text"`
	Format       string                 `json:"format,omitempty"`
}

type CommentParentInput struct {
	PageID string `json:"page_id"`
}

type RetrieveCommentsArgs struct {
	BlockID     string  `json:"block_id"`
	StartCursor *string `json:"start_cursor,omitempty"`
	PageSize    *int    `json:"page_size,omitempty"`
	Format      string  `json:"format,omitempty"`
}

type SearchArgs struct {
	Query       *string       `json:"query,omitempty"`
	Filter      *SearchFilter `json:"filter,omitempty"`
	Sort        *SearchSort   `json:"sort,omitempty"`
	StartCursor *string       `json:"start_cursor,omitempty"`
	PageSize    *int          `json:"page_size,omitempty"`
	Format      string        `json:"format,omitempty"`
}

type SearchFilter struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

type SearchSort struct {
	Direction string `json:"direction"`
	Timestamp string `json:"timestamp"`
}
