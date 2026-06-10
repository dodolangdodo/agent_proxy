package siliconflow

import (
	"strings"

	"github.com/bytedance/sonic/ast"
	"github.com/labring/aiproxy/core/relay/meta"
)

func isVLEmbeddingModel(meta *meta.Meta) bool {
	if meta == nil {
		return false
	}

	if vision, ok := meta.ModelConfig.SupportVision(); ok && vision {
		return true
	}

	return isVLEmbeddingModelName(meta.OriginModel) ||
		isVLEmbeddingModelName(meta.ActualModel)
}

func isVLEmbeddingModelName(modelName string) bool {
	modelName = strings.ToLower(modelName)
	return strings.Contains(modelName, "multimodal") ||
		strings.Contains(modelName, "vision") ||
		(strings.Contains(modelName, "vl") && strings.Contains(modelName, "embedding"))
}

func patchVLEmbeddingsInput(node *ast.Node) error {
	inputNode := node.Get("input")
	if !inputNode.Exists() {
		return nil
	}

	switch inputNode.TypeSafe() {
	case ast.V_ARRAY:
		var patchErr error

		err := inputNode.ForEach(func(_ ast.Sequence, item *ast.Node) bool {
			patchErr = patchVLEmbeddingsInputItem(item)
			return patchErr == nil
		})
		if err != nil {
			return err
		}

		return patchErr
	case ast.V_OBJECT:
		return patchVLEmbeddingsInputItem(inputNode)
	default:
		return nil
	}
}

func patchVLEmbeddingsInputItem(item *ast.Node) error {
	if item.TypeSafe() != ast.V_OBJECT {
		return nil
	}

	imageURL, ok, err := openAIEmbeddingImageURL(item)
	if err != nil {
		return err
	}

	if ok {
		*item = newVLEmbeddingImageInput(imageURL)
		return nil
	}

	typeNode := item.Get("type")
	if !typeNode.Exists() {
		return nil
	}

	contentType, err := typeNode.String()
	if err != nil {
		return err
	}

	switch contentType {
	case "text":
		textNode := item.Get("text")
		if textNode.Exists() && textNode.TypeSafe() == ast.V_STRING {
			text, err := textNode.String()
			if err != nil {
				return err
			}

			*item = newVLEmbeddingTextInput(text)
		}
	case "image":
		imageNode := item.Get("image")
		if imageNode.Exists() && imageNode.TypeSafe() == ast.V_STRING {
			image, err := imageNode.String()
			if err != nil {
				return err
			}

			*item = newVLEmbeddingImageInput(image)
		}
	}

	return nil
}

func openAIEmbeddingImageURL(item *ast.Node) (string, bool, error) {
	typeNode := item.Get("type")
	if typeNode.Exists() {
		contentType, err := typeNode.String()
		if err != nil {
			return "", false, err
		}

		if contentType != "image_url" {
			return "", false, nil
		}
	}

	imageURLNode := item.Get("image_url")
	if !imageURLNode.Exists() || imageURLNode.TypeSafe() != ast.V_OBJECT {
		return "", false, nil
	}

	urlNode := imageURLNode.Get("url")
	if !urlNode.Exists() || urlNode.TypeSafe() != ast.V_STRING {
		return "", false, nil
	}

	imageURL, err := urlNode.String()
	if err != nil {
		return "", false, err
	}

	return imageURL, true, nil
}

func newVLEmbeddingTextInput(text string) ast.Node {
	return ast.NewObject([]ast.Pair{
		ast.NewPair("text", ast.NewString(text)),
	})
}

func newVLEmbeddingImageInput(image string) ast.Node {
	return ast.NewObject([]ast.Pair{
		ast.NewPair("image", ast.NewString(image)),
	})
}
