package figma

// ImageNode represents an image node to download
type ImageNode struct {
	NodeID   string `json:"nodeId"`
	ImageRef string `json:"imageRef,omitempty"`
	FileName string `json:"fileName"`
}

// SVGOptions represents options for SVG export
type SVGOptions struct {
	OutlineText    bool `json:"outlineText"`
	IncludeID      bool `json:"includeId"`
	SimplifyStroke bool `json:"simplifyStroke"`
}

// SimplifiedDesign represents the processed design data
type SimplifiedDesign struct {
	Metadata   DesignMetadata `json:"metadata"`
	Nodes      []Node         `json:"nodes"`
	GlobalVars GlobalVars     `json:"globalVars"`
}

// DesignMetadata represents design metadata
type DesignMetadata struct {
	Name         string `json:"name"`
	LastModified string `json:"lastModified"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

// Node represents a simplified design node
type Node struct {
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	Type         string              `json:"type"`
	BoundingBox  *BoundingBox        `json:"boundingBox,omitempty"`
	Text         string              `json:"text,omitempty"`
	TextStyle    string              `json:"textStyle,omitempty"`
	Fills        string              `json:"fills,omitempty"`
	Styles       string              `json:"styles,omitempty"`
	Strokes      string              `json:"strokes,omitempty"`
	Effects      string              `json:"effects,omitempty"`
	Opacity      *float64            `json:"opacity,omitempty"`
	BorderRadius string              `json:"borderRadius,omitempty"`
	Layout       string              `json:"layout,omitempty"`
	ComponentID  string              `json:"componentId,omitempty"`
	Properties   []ComponentProperty `json:"componentProperties,omitempty"`
	Children     []Node              `json:"children,omitempty"`
}

// BoundingBox represents a bounding box
type BoundingBox struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// ComponentProperty represents a component property
type ComponentProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

// GlobalVars represents global variables
type GlobalVars struct {
	Styles map[string]any `json:"styles"`
}

// Figma API response types
type FileResponse struct {
	Name          string                  `json:"name"`
	LastModified  string                  `json:"lastModified"`
	ThumbnailURL  string                  `json:"thumbnailUrl"`
	Document      FigmaNode               `json:"document"`
	Components    map[string]Component    `json:"components"`
	ComponentSets map[string]ComponentSet `json:"componentSets"`
}

type NodesResponse struct {
	Name         string              `json:"name"`
	LastModified string              `json:"lastModified"`
	ThumbnailURL string              `json:"thumbnailUrl"`
	Nodes        map[string]NodeData `json:"nodes"`
}

type NodeData struct {
	Document      FigmaNode               `json:"document"`
	Components    map[string]Component    `json:"components"`
	ComponentSets map[string]ComponentSet `json:"componentSets"`
}

type FigmaNode struct {
	ID                   string                            `json:"id"`
	Name                 string                            `json:"name"`
	Type                 string                            `json:"type"`
	Visible              *bool                             `json:"visible,omitempty"`
	AbsoluteBoundingBox  *Rectangle                        `json:"absoluteBoundingBox,omitempty"`
	Characters           string                            `json:"characters,omitempty"`
	Style                *TextStyle                        `json:"style,omitempty"`
	Fills                []Paint                           `json:"fills,omitempty"`
	Strokes              []Paint                           `json:"strokes,omitempty"`
	StrokeWeight         *float64                          `json:"strokeWeight,omitempty"`
	Effects              []Effect                          `json:"effects,omitempty"`
	Opacity              *float64                          `json:"opacity,omitempty"`
	CornerRadius         *float64                          `json:"cornerRadius,omitempty"`
	RectangleCornerRadii []float64                         `json:"rectangleCornerRadii,omitempty"`
	ComponentID          string                            `json:"componentId,omitempty"`
	ComponentProperties  map[string]ComponentPropertyValue `json:"componentProperties,omitempty"`
	Children             []FigmaNode                       `json:"children,omitempty"`
	// Layout properties
	LayoutMode             string   `json:"layoutMode,omitempty"`
	PrimaryAxisAlignItems  string   `json:"primaryAxisAlignItems,omitempty"`
	CounterAxisAlignItems  string   `json:"counterAxisAlignItems,omitempty"`
	ItemSpacing            *float64 `json:"itemSpacing,omitempty"`
	PaddingTop             *float64 `json:"paddingTop,omitempty"`
	PaddingRight           *float64 `json:"paddingRight,omitempty"`
	PaddingBottom          *float64 `json:"paddingBottom,omitempty"`
	PaddingLeft            *float64 `json:"paddingLeft,omitempty"`
	LayoutAlign            string   `json:"layoutAlign,omitempty"`
	LayoutGrow             *float64 `json:"layoutGrow,omitempty"`
	LayoutSizingHorizontal string   `json:"layoutSizingHorizontal,omitempty"`
	LayoutSizingVertical   string   `json:"layoutSizingVertical,omitempty"`
	LayoutPositioning      string   `json:"layoutPositioning,omitempty"`
}

type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type TextStyle struct {
	FontFamily          string   `json:"fontFamily,omitempty"`
	FontWeight          *int     `json:"fontWeight,omitempty"`
	FontSize            *float64 `json:"fontSize,omitempty"`
	LineHeightPx        *float64 `json:"lineHeightPx,omitempty"`
	LetterSpacing       *float64 `json:"letterSpacing,omitempty"`
	TextCase            string   `json:"textCase,omitempty"`
	TextAlignHorizontal string   `json:"textAlignHorizontal,omitempty"`
	TextAlignVertical   string   `json:"textAlignVertical,omitempty"`
}

type Paint struct {
	Type                    string         `json:"type"`
	Visible                 *bool          `json:"visible,omitempty"`
	Opacity                 *float64       `json:"opacity,omitempty"`
	Color                   *RGBA          `json:"color,omitempty"`
	ImageRef                string         `json:"imageRef,omitempty"`
	ScaleMode               string         `json:"scaleMode,omitempty"`
	GradientHandlePositions []Vector       `json:"gradientHandlePositions,omitempty"`
	GradientStops           []GradientStop `json:"gradientStops,omitempty"`
}

type RGBA struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
	A float64 `json:"a"`
}

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type GradientStop struct {
	Position float64 `json:"position"`
	Color    RGBA    `json:"color"`
}

type Effect struct {
	Type    string   `json:"type"`
	Visible *bool    `json:"visible,omitempty"`
	Radius  *float64 `json:"radius,omitempty"`
	Color   *RGBA    `json:"color,omitempty"`
	Offset  *Vector  `json:"offset,omitempty"`
	Spread  *float64 `json:"spread,omitempty"`
}

type ComponentPropertyValue struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Component struct {
	Key            string `json:"key"`
	Name           string `json:"name"`
	ComponentSetID string `json:"componentSetId,omitempty"`
}

type ComponentSet struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type ImagesResponse struct {
	Images map[string]string `json:"images"`
}

type ImageFillsResponse struct {
	Meta struct {
		Images map[string]string `json:"images"`
	} `json:"meta"`
}
