package figma

import (
	"fmt"
	"strings"
)

// ParseFigmaResponse parses a Figma file response
func ParseFigmaResponse(response FileResponse) *SimplifiedDesign {
	globalVars := GlobalVars{
		Styles: make(map[string]any),
	}

	nodes := parseNodes(response.Document.Children, &globalVars, nil)

	return &SimplifiedDesign{
		Metadata: DesignMetadata{
			Name:         response.Name,
			LastModified: response.LastModified,
			ThumbnailURL: response.ThumbnailURL,
		},
		Nodes:      nodes,
		GlobalVars: globalVars,
	}
}

// ParseFigmaNodesResponse parses a Figma nodes response
func ParseFigmaNodesResponse(response NodesResponse) *SimplifiedDesign {
	globalVars := GlobalVars{
		Styles: make(map[string]any),
	}

	var allNodes []Node
	for _, nodeData := range response.Nodes {
		node := parseNode(nodeData.Document, &globalVars, nil)
		if node != nil {
			allNodes = append(allNodes, *node)
		}
	}

	return &SimplifiedDesign{
		Metadata: DesignMetadata{
			Name:         response.Name,
			LastModified: response.LastModified,
			ThumbnailURL: response.ThumbnailURL,
		},
		Nodes:      allNodes,
		GlobalVars: globalVars,
	}
}

// parseNodes parses multiple Figma nodes
func parseNodes(figmaNodes []FigmaNode, globalVars *GlobalVars, parent *FigmaNode) []Node {
	var nodes []Node
	for _, figmaNode := range figmaNodes {
		if isVisible(figmaNode) {
			if node := parseNode(figmaNode, globalVars, parent); node != nil {
				nodes = append(nodes, *node)
			}
		}
	}

	return nodes
}

// parseNode parses a single Figma node
func parseNode(figmaNode FigmaNode, globalVars *GlobalVars, parent *FigmaNode) *Node {
	node := Node{
		ID:   figmaNode.ID,
		Name: figmaNode.Name,
		Type: figmaNode.Type,
	}

	// Convert VECTOR to IMAGE-SVG
	if node.Type == "VECTOR" {
		node.Type = "IMAGE-SVG"
	}

	// Parse bounding box
	if figmaNode.AbsoluteBoundingBox != nil {
		node.BoundingBox = &BoundingBox{
			X:      figmaNode.AbsoluteBoundingBox.X,
			Y:      figmaNode.AbsoluteBoundingBox.Y,
			Width:  figmaNode.AbsoluteBoundingBox.Width,
			Height: figmaNode.AbsoluteBoundingBox.Height,
		}
	}

	// Parse text
	if figmaNode.Characters != "" {
		node.Text = figmaNode.Characters
	}

	// Parse text style
	if figmaNode.Style != nil && hasTextStyleContent(*figmaNode.Style) {
		styleID := findOrCreateVar(globalVars, parseTextStyle(*figmaNode.Style), "style")
		node.TextStyle = styleID
	}

	// Parse fills
	if len(figmaNode.Fills) > 0 {
		fills := parseFills(figmaNode.Fills)
		if len(fills) > 0 {
			fillID := findOrCreateVar(globalVars, fills, "fill")
			node.Fills = fillID
		}
	}

	// Parse strokes
	if len(figmaNode.Strokes) > 0 {
		strokes := parseStrokes(figmaNode.Strokes, figmaNode.StrokeWeight)
		if colors, ok := strokes["colors"].([]any); ok && len(colors) > 0 {
			strokeID := findOrCreateVar(globalVars, strokes, "stroke")
			node.Strokes = strokeID
		}
	}

	// Parse effects
	if len(figmaNode.Effects) > 0 {
		effects := parseEffects(figmaNode.Effects)
		if len(effects) > 0 {
			effectID := findOrCreateVar(globalVars, effects, "effect")
			node.Effects = effectID
		}
	}

	// Parse opacity
	if figmaNode.Opacity != nil && *figmaNode.Opacity != 1.0 {
		node.Opacity = figmaNode.Opacity
	}

	// Parse border radius
	if figmaNode.CornerRadius != nil {
		node.BorderRadius = fmt.Sprintf("%.1fpx", *figmaNode.CornerRadius)
	} else if len(figmaNode.RectangleCornerRadii) == 4 {
		node.BorderRadius = fmt.Sprintf("%.1fpx %.1fpx %.1fpx %.1fpx",
			figmaNode.RectangleCornerRadii[0],
			figmaNode.RectangleCornerRadii[1],
			figmaNode.RectangleCornerRadii[2],
			figmaNode.RectangleCornerRadii[3])
	}

	// Parse layout
	layout := parseLayout(figmaNode, parent)
	if len(layout) > 1 { // More than just "mode"
		layoutID := findOrCreateVar(globalVars, layout, "layout")
		node.Layout = layoutID
	}

	// Parse component properties
	if figmaNode.ComponentID != "" {
		node.ComponentID = figmaNode.ComponentID
	}

	if len(figmaNode.ComponentProperties) > 0 {
		for name, prop := range figmaNode.ComponentProperties {
			node.Properties = append(node.Properties, ComponentProperty{
				Name:  name,
				Value: prop.Value,
				Type:  prop.Type,
			})
		}
	}

	// Parse children
	if len(figmaNode.Children) > 0 {
		children := parseNodes(figmaNode.Children, globalVars, &figmaNode)
		if len(children) > 0 {
			node.Children = children
		}
	}

	return &node
}

// Helper functions

func isVisible(node FigmaNode) bool {
	return node.Visible == nil || *node.Visible
}

func hasTextStyleContent(style TextStyle) bool {
	return style.FontFamily != "" || style.FontWeight != nil || style.FontSize != nil
}

func parseTextStyle(style TextStyle) map[string]any {
	result := make(map[string]any)

	if style.FontFamily != "" {
		result["fontFamily"] = style.FontFamily
	}

	if style.FontWeight != nil {
		result["fontWeight"] = *style.FontWeight
	}

	if style.FontSize != nil {
		result["fontSize"] = *style.FontSize
	}

	if style.LineHeightPx != nil && style.FontSize != nil && *style.FontSize > 0 {
		result["lineHeight"] = fmt.Sprintf("%.2fem", *style.LineHeightPx / *style.FontSize)
	}

	if style.LetterSpacing != nil && *style.LetterSpacing != 0 && style.FontSize != nil &&
		*style.FontSize > 0 {
		result["letterSpacing"] = fmt.Sprintf(
			"%.2f%%",
			(*style.LetterSpacing / *style.FontSize)*100,
		)
	}

	if style.TextCase != "" {
		result["textCase"] = style.TextCase
	}

	if style.TextAlignHorizontal != "" {
		result["textAlignHorizontal"] = style.TextAlignHorizontal
	}

	if style.TextAlignVertical != "" {
		result["textAlignVertical"] = style.TextAlignVertical
	}

	return result
}

func parseFills(fills []Paint) []any {
	var result []any
	for _, fill := range fills {
		if isVisible(FigmaNode{Visible: fill.Visible}) {
			result = append(result, parsePaint(fill))
		}
	}

	return result
}

func parseStrokes(strokes []Paint, strokeWeight *float64) map[string]any {
	result := map[string]any{
		"colors": []any{},
	}

	var colors []any
	for _, stroke := range strokes {
		if isVisible(FigmaNode{Visible: stroke.Visible}) {
			colors = append(colors, parsePaint(stroke))
		}
	}

	result["colors"] = colors

	if strokeWeight != nil && *strokeWeight > 0 {
		result["strokeWeight"] = fmt.Sprintf("%.1fpx", *strokeWeight)
	}

	return result
}

func parseEffects(effects []Effect) map[string]any {
	result := make(map[string]any)

	var (
		boxShadows      []string
		filters         []string
		backdropFilters []string
	)

	for _, effect := range effects {
		if !isVisible(FigmaNode{Visible: effect.Visible}) {
			continue
		}

		switch effect.Type {
		case "DROP_SHADOW":
			if effect.Offset != nil && effect.Radius != nil && effect.Color != nil {
				spread := 0.0
				if effect.Spread != nil {
					spread = *effect.Spread
				}

				shadow := fmt.Sprintf("%.1fpx %.1fpx %.1fpx %.1fpx %s",
					effect.Offset.X, effect.Offset.Y, *effect.Radius, spread,
					formatRGBAColor(*effect.Color))
				boxShadows = append(boxShadows, shadow)
			}
		case "INNER_SHADOW":
			if effect.Offset != nil && effect.Radius != nil && effect.Color != nil {
				spread := 0.0
				if effect.Spread != nil {
					spread = *effect.Spread
				}

				shadow := fmt.Sprintf("inset %.1fpx %.1fpx %.1fpx %.1fpx %s",
					effect.Offset.X, effect.Offset.Y, *effect.Radius, spread,
					formatRGBAColor(*effect.Color))
				boxShadows = append(boxShadows, shadow)
			}
		case "LAYER_BLUR":
			if effect.Radius != nil {
				filters = append(filters, fmt.Sprintf("blur(%.1fpx)", *effect.Radius))
			}
		case "BACKGROUND_BLUR":
			if effect.Radius != nil {
				backdropFilters = append(
					backdropFilters,
					fmt.Sprintf("blur(%.1fpx)", *effect.Radius),
				)
			}
		}
	}

	if len(boxShadows) > 0 {
		result["boxShadow"] = strings.Join(boxShadows, ", ")
	}

	if len(filters) > 0 {
		result["filter"] = strings.Join(filters, " ")
	}

	if len(backdropFilters) > 0 {
		result["backdropFilter"] = strings.Join(backdropFilters, " ")
	}

	return result
}

func parsePaint(paint Paint) any {
	switch paint.Type {
	case "SOLID":
		if paint.Color != nil {
			opacity := 1.0
			if paint.Opacity != nil {
				opacity = *paint.Opacity
			}

			if opacity == 1.0 {
				return convertColorToHex(*paint.Color)
			}

			return formatRGBAColor(*paint.Color)
		}
	case "IMAGE":
		result := map[string]any{
			"type": "IMAGE",
		}
		if paint.ImageRef != "" {
			result["imageRef"] = paint.ImageRef
		}

		if paint.ScaleMode != "" {
			result["scaleMode"] = paint.ScaleMode
		}

		return result
	case "GRADIENT_LINEAR", "GRADIENT_RADIAL", "GRADIENT_ANGULAR", "GRADIENT_DIAMOND":
		result := map[string]any{
			"type": paint.Type,
		}
		if len(paint.GradientHandlePositions) > 0 {
			result["gradientHandlePositions"] = paint.GradientHandlePositions
		}

		if len(paint.GradientStops) > 0 {
			var stops []map[string]any
			for _, stop := range paint.GradientStops {
				stops = append(stops, map[string]any{
					"position": stop.Position,
					"color": map[string]any{
						"hex":     convertColorToHex(stop.Color),
						"opacity": stop.Color.A,
					},
				})
			}

			result["gradientStops"] = stops
		}

		return result
	}

	return nil
}

//nolint:gocyclo
func parseLayout(node FigmaNode, parent *FigmaNode) map[string]any {
	layout := map[string]any{
		"mode": "none",
	}

	// Check if this is a frame with layout properties
	if node.LayoutMode != "" && node.LayoutMode != "NONE" {
		switch node.LayoutMode {
		case "HORIZONTAL":
			layout["mode"] = "row"
		case "VERTICAL":
			layout["mode"] = "column"
		}

		// Add layout properties
		if node.PrimaryAxisAlignItems != "" {
			layout["justifyContent"] = convertAlign(node.PrimaryAxisAlignItems)
		}

		if node.CounterAxisAlignItems != "" {
			layout["alignItems"] = convertAlign(node.CounterAxisAlignItems)
		}

		if node.ItemSpacing != nil {
			layout["gap"] = fmt.Sprintf("%.1fpx", *node.ItemSpacing)
		}

		// Add padding
		if node.PaddingTop != nil || node.PaddingRight != nil ||
			node.PaddingBottom != nil || node.PaddingLeft != nil {
			padding := generateCSSShorthand(
				getValueOrZero(node.PaddingTop),
				getValueOrZero(node.PaddingRight),
				getValueOrZero(node.PaddingBottom),
				getValueOrZero(node.PaddingLeft),
			)
			if padding != "" {
				layout["padding"] = padding
			}
		}
	}

	// Add sizing information
	sizing := make(map[string]any)
	if node.LayoutSizingHorizontal != "" {
		sizing["horizontal"] = convertSizing(node.LayoutSizingHorizontal)
	}

	if node.LayoutSizingVertical != "" {
		sizing["vertical"] = convertSizing(node.LayoutSizingVertical)
	}

	if len(sizing) > 0 {
		layout["sizing"] = sizing
	}

	// Add positioning for absolute positioned elements
	if node.LayoutPositioning == "ABSOLUTE" {
		layout["position"] = "absolute"

		if node.AbsoluteBoundingBox != nil && parent != nil && parent.AbsoluteBoundingBox != nil {
			layout["locationRelativeToParent"] = map[string]float64{
				"x": node.AbsoluteBoundingBox.X - parent.AbsoluteBoundingBox.X,
				"y": node.AbsoluteBoundingBox.Y - parent.AbsoluteBoundingBox.Y,
			}
		}
	}

	// Add dimensions
	if node.AbsoluteBoundingBox != nil {
		dimensions := make(map[string]float64)

		// Add width and height based on layout mode and sizing
		switch layout["mode"] {
		case "row":
			if node.LayoutSizingHorizontal == "FIXED" {
				dimensions["width"] = node.AbsoluteBoundingBox.Width
			}

			if node.LayoutAlign != "STRETCH" && node.LayoutSizingVertical == "FIXED" {
				dimensions["height"] = node.AbsoluteBoundingBox.Height
			}
		case "column":
			if node.LayoutAlign != "STRETCH" && node.LayoutSizingHorizontal == "FIXED" {
				dimensions["width"] = node.AbsoluteBoundingBox.Width
			}

			if node.LayoutSizingVertical == "FIXED" {
				dimensions["height"] = node.AbsoluteBoundingBox.Height
			}
		default:
			// Not in auto layout
			if node.LayoutSizingHorizontal == "" || node.LayoutSizingHorizontal == "FIXED" {
				dimensions["width"] = node.AbsoluteBoundingBox.Width
			}

			if node.LayoutSizingVertical == "" || node.LayoutSizingVertical == "FIXED" {
				dimensions["height"] = node.AbsoluteBoundingBox.Height
			}
		}

		if len(dimensions) > 0 {
			layout["dimensions"] = dimensions
		}
	}

	return layout
}

// Helper functions for layout parsing

func convertAlign(align string) string {
	switch align {
	case "MIN":
		return "flex-start"
	case "MAX":
		return "flex-end"
	case "CENTER":
		return "center"
	case "SPACE_BETWEEN":
		return "space-between"
	case "BASELINE":
		return "baseline"
	default:
		return "flex-start"
	}
}

func convertSizing(sizing string) string {
	switch sizing {
	case "FIXED":
		return "fixed"
	case "FILL":
		return "fill"
	case "HUG":
		return "hug"
	default:
		return "fixed"
	}
}

func getValueOrZero(value *float64) float64 {
	if value == nil {
		return 0
	}
	return *value
}

func generateCSSShorthand(top, right, bottom, left float64) string {
	if top == 0 && right == 0 && bottom == 0 && left == 0 {
		return ""
	}

	if top == right && right == bottom && bottom == left {
		return fmt.Sprintf("%.1fpx", top)
	}

	if right == left {
		if top == bottom {
			return fmt.Sprintf("%.1fpx %.1fpx", top, right)
		}
		return fmt.Sprintf("%.1fpx %.1fpx %.1fpx", top, right, bottom)
	}

	return fmt.Sprintf("%.1fpx %.1fpx %.1fpx %.1fpx", top, right, bottom, left)
}

func convertColorToHex(color RGBA) string {
	r := int(color.R * 255)
	g := int(color.G * 255)
	b := int(color.B * 255)
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

func formatRGBAColor(color RGBA) string {
	r := int(color.R * 255)
	g := int(color.G * 255)
	b := int(color.B * 255)
	a := color.A

	return fmt.Sprintf("rgba(%d, %d, %d, %.2f)", r, g, b, a)
}

func findOrCreateVar(globalVars *GlobalVars, value any, prefix string) string {
	// Simple implementation - in production you might want to check for existing values
	varID := generateVarID(prefix)
	globalVars.Styles[varID] = value
	return varID
}

func generateVarID(prefix string) string {
	// Simple ID generation - in production you might want something more sophisticated
	return fmt.Sprintf("%s_%d", prefix, len(globalVarCounter))
}

var globalVarCounter = make(map[string]int)
