package jina

import (
	"net/http"
	"strings"

	"github.com/bytedance/sonic/ast"
	"github.com/labring/aiproxy/core/relay/adaptor"
	"github.com/labring/aiproxy/core/relay/adaptor/openai"
	"github.com/labring/aiproxy/core/relay/meta"
)

func ConvertEmbeddingsRequest(
	meta *meta.Meta,
	req *http.Request,
) (adaptor.ConvertResult, error) {
	return openai.ConvertEmbeddingsRequest(meta, req, true, func(node *ast.Node) error {
		if _, err := node.Unset("encoding_format"); err != nil {
			return err
		}

		return patchEmbeddingsInput(node)
	})
}

func patchEmbeddingsInput(node *ast.Node) error {
	inputNode := node.Get("input")
	if !inputNode.Exists() {
		return nil
	}

	switch inputNode.TypeSafe() {
	case ast.V_STRING:
		text, err := inputNode.String()
		if err != nil {
			return err
		}

		*inputNode = ast.NewArray([]ast.Node{newJinaTextInput(text)})

		return nil
	case ast.V_ARRAY:
		var patchErr error

		err := inputNode.ForEach(func(_ ast.Sequence, item *ast.Node) bool {
			patchErr = patchEmbeddingsInputItem(item)
			return patchErr == nil
		})
		if err != nil {
			return err
		}

		return patchErr
	default:
		return nil
	}
}

func patchEmbeddingsInputItem(item *ast.Node) error {
	switch item.TypeSafe() {
	case ast.V_STRING:
		text, err := item.String()
		if err != nil {
			return err
		}

		*item = newJinaTextInput(text)

		return nil
	case ast.V_OBJECT:
		imageURL, ok, err := openAIImageURL(item)
		if err != nil {
			return err
		}

		if ok {
			*item = newJinaImageInput(imageURL)
			return nil
		}

		imageNode := item.Get("image")
		if imageNode.Exists() && imageNode.TypeSafe() == ast.V_STRING {
			image, err := imageNode.String()
			if err != nil {
				return err
			}

			normalizedImage := normalizeJinaImage(image)
			if normalizedImage != image || item.Get("type").Exists() {
				*item = newJinaImageInput(normalizedImage)
			}

			return nil
		}

		textNode := item.Get("text")
		if textNode.Exists() && textNode.TypeSafe() == ast.V_STRING && item.Get("type").Exists() {
			text, err := textNode.String()
			if err != nil {
				return err
			}

			*item = newJinaTextInput(text)
		}

		return nil
	default:
		return nil
	}
}

func openAIImageURL(item *ast.Node) (string, bool, error) {
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

func newJinaTextInput(text string) ast.Node {
	return ast.NewObject([]ast.Pair{
		ast.NewPair("text", ast.NewString(text)),
	})
}

func newJinaImageInput(image string) ast.Node {
	return ast.NewObject([]ast.Pair{
		ast.NewPair("image", ast.NewString(normalizeJinaImage(image))),
	})
}

func normalizeJinaImage(image string) string {
	if !strings.HasPrefix(image, "data:image/") {
		return image
	}

	_, base64Data, ok := strings.Cut(image, ";base64,")
	if !ok {
		return image
	}

	if base64Data == "" {
		return image
	}

	return base64Data
}
