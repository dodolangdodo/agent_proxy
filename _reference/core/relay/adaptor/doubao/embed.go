package doubao

import (
	"github.com/bytedance/sonic/ast"
	"github.com/labring/aiproxy/core/relay/meta"
)

func patchEmbeddingsVisionInput(node *ast.Node) error {
	inputNode := node.Get("input")
	if !inputNode.Exists() {
		return nil
	}

	switch inputNode.TypeSafe() {
	case ast.V_ARRAY:
		return inputNode.ForEach(func(_ ast.Sequence, item *ast.Node) bool {
			switch item.TypeSafe() {
			case ast.V_STRING:
				text, err := item.String()
				if err != nil {
					return false
				}

				*item = ast.NewObject([]ast.Pair{
					ast.NewPair("type", ast.NewString("text")),
					ast.NewPair("text", ast.NewString(text)),
				})

				return true
			case ast.V_OBJECT:
				textNode := item.Get("text")
				if textNode.Exists() && textNode.TypeSafe() == ast.V_STRING {
					_, err := item.Set("type", ast.NewString("text"))
					return err == nil
				}

				imageNode := item.Get("image")
				if imageNode.Exists() && imageNode.TypeSafe() == ast.V_STRING {
					imageURL, err := imageNode.String()
					if err != nil {
						return false
					}

					return setEmbeddingVisionURLItem(item, "image", "image_url", imageURL) == nil
				}

				videoNode := item.Get("video")
				if videoNode.Exists() && videoNode.TypeSafe() == ast.V_STRING {
					videoURL, err := videoNode.String()
					if err != nil {
						return false
					}

					return setEmbeddingVisionURLItem(item, "video", "video_url", videoURL) == nil
				}

				if item.Get("image_url").Exists() {
					_, err := item.Set("type", ast.NewString("image_url"))
					return err == nil
				}

				if item.Get("video_url").Exists() {
					_, err := item.Set("type", ast.NewString("video_url"))
					return err == nil
				}

				return true
			default:
				return false
			}
		})
	case ast.V_STRING:
		inputText, err := inputNode.String()
		if err != nil {
			return err
		}

		_, err = node.Set("input", ast.NewArray([]ast.Node{
			ast.NewObject([]ast.Pair{
				ast.NewPair("type", ast.NewString("text")),
				ast.NewPair("text", ast.NewString(inputText)),
			}),
		}))

		return err
	default:
		return nil
	}
}

func setEmbeddingVisionURLItem(
	item *ast.Node,
	sourceKey string,
	targetKey string,
	url string,
) error {
	_, err := item.Unset(sourceKey)
	if err != nil {
		return err
	}

	_, err = item.Set("type", ast.NewString(targetKey))
	if err != nil {
		return err
	}

	_, err = item.Set(targetKey, ast.NewObject([]ast.Pair{
		ast.NewPair("url", ast.NewString(url)),
	}))

	return err
}

func embeddingPreHandler(_ *meta.Meta, node *ast.Node) error {
	return patchEmbeddingsVisionResponse(node)
}

func patchEmbeddingsVisionResponse(node *ast.Node) error {
	dataNode := node.Get("data")
	if !dataNode.Exists() {
		return nil
	}

	switch dataNode.TypeSafe() {
	case ast.V_ARRAY:
		return nil
	case ast.V_OBJECT:
		embeddingNode := dataNode.Get("embedding")
		if !embeddingNode.Exists() {
			return nil
		}

		if _, err := dataNode.Set("object", ast.NewString("embedding")); err != nil {
			return err
		}

		if _, err := dataNode.Set("index", ast.NewNumber("0")); err != nil {
			return err
		}

		_, err := node.Set("data", ast.NewArray([]ast.Node{*dataNode}))

		return err
	default:
		return nil
	}
}
