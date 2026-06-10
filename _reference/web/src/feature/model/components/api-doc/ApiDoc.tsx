import React from 'react'
import {
    Sheet,
    SheetContent,
} from "@/components/ui/sheet"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import { ModelConfig } from '@/types/model'
import { useTranslation } from 'react-i18next'
import CodeBlock from './CodeHight'
import { toast } from 'sonner'
import { TFunction } from 'i18next'

interface ApiDocContent {
    title: string
    endpoint: string
    method: string
    requestExample: string
    responseExample: string
    responseFormat: string
    requestAdditionalInfo?: {
        voices?: string[]
        formats?: string[]
    }
    responseAdditionalInfo?: {
        voices?: string[]
        formats?: string[]
    }
}

interface ApiDocDrawerProps {
    isOpen: boolean
    onClose: () => void
    modelConfig: ModelConfig
}

const getApiDocContent = (
    modelConfig: ModelConfig,
    apiEndpoint: string,
    t: TFunction
): ApiDocContent => {
    switch (modelConfig.type) {
        case 1:
            return {
                title: t('modeType.1'),
                endpoint: '/chat/completions',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/chat/completions \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "messages": [
    {
      "role": "user",
      "content": "What is Sealos"
    }
  ],
  "stream": false,
  "max_tokens": 512,
  "temperature": 0.7
}'`,
                responseExample: `{
  "object": "chat.completion",
  "created": 1729672480,
  "model": "${modelConfig.model}",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": "Sealos is a cloud operating system based on Kubernetes, designed to provide users with a simple, efficient, and scalable cloud-native application deployment and management experience."
      },
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 18,
    "completion_tokens": 52,
    "total_tokens": 70
  }
}`
            }
        case 3:
            return {
                title: t('modeType.3'),
                endpoint: '/embeddings',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/embeddings \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "input": "Your text string goes here",
  "encoding_format": "float"
}'`,
                responseExample: `{
  "object": "list",
  "model": "${modelConfig.model}",
  "data": [
    {
      "object": "embedding",
      "embedding": [
        -0.1082494854927063,
        0.022976752370595932
        ...
      ],
      "index": 0
    }
  ],
  "usage": {
    "prompt_tokens": 4,
    "completion_tokens": 0,
    "total_tokens": 4
  }
}`
            }
        case 7:
            return {
                title: t('modeType.7'),
                endpoint: '/audio/speech',
                method: 'POST',
                responseFormat: 'binary',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/audio/speech \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "input": "The text to generate audio for",
${modelConfig?.config?.support_voices?.length
                        ? `  "voice": "${modelConfig.config.support_voices[0]}",\n`
                        : ''
                    }${modelConfig?.config?.support_formats?.length
                        ? `  "response_format": "${modelConfig.config.support_formats[0]}",\n`
                        : ''
                    }  "stream": true,
  "speed": 1
}' > audio.mp3`,
                responseExample: 'Binary audio data',
                requestAdditionalInfo: {
                    voices: modelConfig?.config?.support_voices,
                    formats: modelConfig?.config?.support_formats
                }
            }
        case 8:
            return {
                title: t('modeType.8'),
                endpoint: '/audio/transcriptions',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/audio/transcriptions \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: multipart/form-data' \\
--form model=${modelConfig.model} \\
--form 'file=@"audio.mp3"'`,
                responseExample: `{
  "text": "<string>"
}`
            }
        case 10:
            return {
                title: t('modeType.10'),
                endpoint: '/rerank',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/rerank \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "query": "Apple",
  "documents": [
    "Apple",
    "Banana",
    "Fruit",
    "Vegetable"
  ],
  "top_n": 4,
  "return_documents": false,
  "max_chunks_per_doc": 1024,
  "overlap_tokens": 80
}'`,
                responseExample: `{
  "results": [
    {
      "index": 0,
      "relevance_score": 0.9953725
    },
    {
      "index": 2,
      "relevance_score": 0.002157342
    },
    {
      "index": 1,
      "relevance_score": 0.00046371284
    },
    {
      "index": 3,
      "relevance_score": 0.000017502925
    }
  ],
  "meta": {
    "tokens": {
      "input_tokens": 28
    }
  }
}`
            }
        case 11:
            return {
                title: t('modeType.11'),
                endpoint: '/parse/pdf',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/parse/pdf \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: multipart/form-data' \\
--form model=${modelConfig.model} \\
--form 'file=@"test.pdf"'`,
                responseExample: `{
  "pages": 1,
  "markdown": "sf ad fda daf da \\\\( f \\\\) ds f sd fs d afdas fsd asfad f\\n\\n\\n\\n![img](data:image/jpeg;base64,/9...)\\n\\n| sadsa |  |  |\\n| --- | --- | --- |\\n|  | sadasdsa | sad |\\n|  |  | dsadsadsa |\\n|  |  |  |\\n\\n\\n\\na fda"
}`
            }
        case 2:
            return {
                title: t('modeType.2'),
                endpoint: '/completions',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/completions \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "prompt": "Write a short product tagline",
  "max_tokens": 64,
  "temperature": 0.7
}'`,
                responseExample: `{
  "object": "text_completion",
  "model": "${modelConfig.model}",
  "choices": [
    {
      "text": "Cloud-native apps, shipped faster.",
      "index": 0,
      "finish_reason": "stop"
    }
  ]
}`
            }
        case 4:
            return {
                title: t('modeType.4'),
                endpoint: '/moderations',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/moderations \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "input": "Text to classify"
}'`,
                responseExample: `{
  "id": "modr_123",
  "model": "${modelConfig.model}",
  "results": [
    {
      "flagged": false,
      "categories": {}
    }
  ]
}`
            }
        case 5:
            return {
                title: t('modeType.5'),
                endpoint: '/images/generations',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/images/generations \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "prompt": "A minimal cloud dashboard illustration",
  "size": "1024x1024",
  "n": 1
}'`,
                responseExample: `{
  "created": 1729672480,
  "data": [
    {
      "url": "https://example.com/image.png"
    }
  ]
}`
            }
        case 6:
            return {
                title: t('modeType.6'),
                endpoint: '/images/edits',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/images/edits \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: multipart/form-data' \\
--form model=${modelConfig.model} \\
--form 'image=@"image.png"' \\
--form 'prompt=Add a clean blue background' \\
--form size=1024x1024`,
                responseExample: `{
  "created": 1729672480,
  "data": [
    {
      "url": "https://example.com/edited-image.png"
    }
  ]
}`
            }
        case 9:
            return {
                title: t('modeType.9'),
                endpoint: '/audio/translations',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/audio/translations \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: multipart/form-data' \\
--form model=${modelConfig.model} \\
--form 'file=@"audio.mp3"'`,
                responseExample: `{
  "text": "Translated transcript text"
}`
            }
        case 12:
            return {
                title: t('modeType.12'),
                endpoint: '/messages',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/messages \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "max_tokens": 512,
  "messages": [
    {
      "role": "user",
      "content": "Summarize this release note"
    }
  ]
}'`,
                responseExample: `{
  "id": "msg_123",
  "type": "message",
  "role": "assistant",
  "content": [
    {
      "type": "text",
      "text": "Summary text"
    }
  ]
}`
            }
        case 13:
            return {
                title: t('modeType.13'),
                endpoint: '/video/generations/jobs',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/video/generations/jobs \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "prompt": "A calm ocean at sunrise",
  "width": 1280,
  "height": 720,
  "n_seconds": 5
}'`,
                responseExample: `{
  "id": "vgjob_123",
  "object": "video.generation.job",
  "status": "queued",
  "model": "${modelConfig.model}"
}`
            }
        case 14:
            return {
                title: t('modeType.14'),
                endpoint: '/video/generations/jobs/{id}',
                method: 'GET',
                responseFormat: 'json',
                requestExample: `curl --request GET \\
--url ${apiEndpoint}/v1/video/generations/jobs/vgjob_123 \\
--header "Authorization: Bearer $token"`,
                responseExample: `{
  "id": "vgjob_123",
  "object": "video.generation.job",
  "status": "succeeded",
  "generations": [
    {
      "id": "video_123"
    }
  ]
}`
            }
        case 15:
            return {
                title: t('modeType.15'),
                endpoint: '/video/generations/{id}/content/video',
                method: 'GET',
                responseFormat: 'binary',
                requestExample: `curl --request GET \\
--url ${apiEndpoint}/v1/video/generations/video_123/content/video \\
--header "Authorization: Bearer $token" \\
--output video.mp4`,
                responseExample: 'Binary video data'
            }
        case 16:
            return {
                title: t('modeType.16'),
                endpoint: '/responses',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/responses \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "input": "Write a concise status update"
}'`,
                responseExample: `{
  "id": "resp_123",
  "object": "response",
  "status": "completed",
  "output_text": "Status update text"
}`
            }
        case 17:
            return {
                title: t('modeType.17'),
                endpoint: '/responses/{response_id}',
                method: 'GET',
                responseFormat: 'json',
                requestExample: `curl --request GET \\
--url ${apiEndpoint}/v1/responses/resp_123 \\
--header "Authorization: Bearer $token"`,
                responseExample: `{
  "id": "resp_123",
  "object": "response",
  "status": "completed"
}`
            }
        case 18:
            return {
                title: t('modeType.18'),
                endpoint: '/responses/{response_id}',
                method: 'DELETE',
                responseFormat: 'json',
                requestExample: `curl --request DELETE \\
--url ${apiEndpoint}/v1/responses/resp_123 \\
--header "Authorization: Bearer $token"`,
                responseExample: `{
  "id": "resp_123",
  "object": "response.deleted",
  "deleted": true
}`
            }
        case 19:
            return {
                title: t('modeType.19'),
                endpoint: '/responses/{response_id}/cancel',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/responses/resp_123/cancel \\
--header "Authorization: Bearer $token"`,
                responseExample: `{
  "id": "resp_123",
  "object": "response",
  "status": "cancelled"
}`
            }
        case 20:
            return {
                title: t('modeType.20'),
                endpoint: '/responses/{response_id}/input_items',
                method: 'GET',
                responseFormat: 'json',
                requestExample: `curl --request GET \\
--url ${apiEndpoint}/v1/responses/resp_123/input_items \\
--header "Authorization: Bearer $token"`,
                responseExample: `{
  "object": "list",
  "data": [
    {
      "id": "item_123",
      "type": "message"
    }
  ]
}`
            }
        case 21:
            return {
                title: t('modeType.21'),
                endpoint: '/models/{model}:generateContent',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1beta/models/${modelConfig.model}:generateContent \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "contents": [
    {
      "role": "user",
      "parts": [
        {
          "text": "Explain Kubernetes in one sentence"
        }
      ]
    }
  ]
}'`,
                responseExample: `{
  "candidates": [
    {
      "content": {
        "parts": [
          {
            "text": "Kubernetes automates deployment, scaling, and management of containerized applications."
          }
        ]
      }
    }
  ]
}`
            }
        case 22:
            return {
                title: t('modeType.22'),
                endpoint: '/videos',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/videos \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "prompt": "A calm ocean at sunrise",
  "seconds": 5,
  "size": "1280x720"
}'`,
                responseExample: `{
  "id": "video_123",
  "object": "video",
  "status": "queued",
  "model": "${modelConfig.model}"
}`
            }
        case 23:
            return {
                title: t('modeType.23'),
                endpoint: '/videos/{video_id}',
                method: 'GET',
                responseFormat: 'json',
                requestExample: `curl --request GET \\
--url ${apiEndpoint}/v1/videos/video_123 \\
--header "Authorization: Bearer $token"`,
                responseExample: `{
  "id": "video_123",
  "object": "video",
  "status": "completed"
}`
            }
        case 24:
            return {
                title: t('modeType.24'),
                endpoint: '/videos/{video_id}/content',
                method: 'GET',
                responseFormat: 'binary',
                requestExample: `curl --request GET \\
--url ${apiEndpoint}/v1/videos/video_123/content \\
--header "Authorization: Bearer $token" \\
--output video.mp4`,
                responseExample: 'Binary video data'
            }
        case 25:
            return {
                title: t('modeType.25'),
                endpoint: '/videos/{video_id}',
                method: 'DELETE',
                responseFormat: 'text',
                requestExample: `curl --request DELETE \\
--url ${apiEndpoint}/v1/videos/video_123 \\
--header "Authorization: Bearer $token"`,
                responseExample: 'No content'
            }
        case 26:
            return {
                title: t('modeType.26'),
                endpoint: '/videos/{video_id}/remix',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/videos/video_123/remix \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "prompt": "Make it cinematic",
  "seconds": 5,
  "size": "1280x720"
}'`,
                responseExample: `{
  "id": "video_456",
  "object": "video",
  "status": "queued"
}`
            }
        case 27:
            return {
                title: t('modeType.27'),
                endpoint: '/models/{model}:predictLongRunning',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1beta/models/${modelConfig.model}:predictLongRunning \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "instances": [
    {
      "prompt": "A calm ocean at sunrise"
    }
  ],
  "parameters": {
    "durationSeconds": 8,
    "resolution": "720p"
  }
}'`,
                responseExample: `{
  "name": "operations/video-operation-123",
  "done": false
}`
            }
        case 28:
            return {
                title: t('modeType.28'),
                endpoint: '/operations/{operation_id}',
                method: 'GET',
                responseFormat: 'json',
                requestExample: `curl --request GET \\
--url ${apiEndpoint}/v1beta/operations/video-operation-123 \\
--header "Authorization: Bearer $token"`,
                responseExample: `{
  "name": "operations/video-operation-123",
  "done": true,
  "response": {}
}`
            }
        case 29:
            return {
                title: t('modeType.29'),
                endpoint: '/models/{model}:generateContent',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1beta/models/${modelConfig.model}:generateContent \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "contents": [
    {
      "parts": [
        {
          "text": "Say hello in a friendly voice"
        }
      ]
    }
  ],
  "generationConfig": {
    "responseModalities": ["AUDIO"]
  }
}'`,
                responseExample: `{
  "candidates": [
    {
      "content": {
        "parts": [
          {
            "inlineData": {
              "mimeType": "audio/wav",
              "data": "BASE64_AUDIO"
            }
          }
        ]
      }
    }
  ]
}`
            }
        case 30:
            return {
                title: t('modeType.30'),
                endpoint: '/models/{model}:generateContent',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1beta/models/${modelConfig.model}:generateContent \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "contents": [
    {
      "parts": [
        {
          "text": "Generate a clean product icon"
        }
      ]
    }
  ],
  "generationConfig": {
    "responseModalities": ["IMAGE"]
  }
}'`,
                responseExample: `{
  "candidates": [
    {
      "content": {
        "parts": [
          {
            "inlineData": {
              "mimeType": "image/png",
              "data": "BASE64_IMAGE"
            }
          }
        ]
      }
    }
  ]
}`
            }
        case 31:
            return {
                title: t('modeType.31'),
                endpoint: '/files/{file}:download',
                method: 'GET',
                responseFormat: 'binary',
                requestExample: `curl --request GET \\
--url ${apiEndpoint}/v1beta/files/abc123:download?alt=media \\
--header "Authorization: Bearer $token" \\
--output video.mp4`,
                responseExample: 'Binary file data'
            }
        case 32:
            return {
                title: t('modeType.32'),
                endpoint: '/videos/edits',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/videos/edits \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: multipart/form-data' \\
--form model=${modelConfig.model} \\
--form 'video=@"source.mp4"' \\
--form 'prompt=Replace the background with a sunrise' \\
--form seconds=5 \\
--form size=1280x720`,
                responseExample: `{
  "id": "video_edited_123",
  "object": "video",
  "status": "queued"
}`
            }
        case 33:
            return {
                title: t('modeType.33'),
                endpoint: '/videos/extensions',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/v1/videos/extensions \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: multipart/form-data' \\
--form model=${modelConfig.model} \\
--form 'video=@"source.mp4"' \\
--form 'prompt=Continue the scene naturally' \\
--form seconds=5 \\
--form size=1280x720`,
                responseExample: `{
  "id": "video_extended_123",
  "object": "video",
  "status": "queued"
}`
            }
        case 34:
            return {
                title: t('modeType.34'),
                endpoint: '/services/aigc/video-generation/video-synthesis',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/api/v1/services/aigc/video-generation/video-synthesis \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "input": {
    "prompt": "A calm ocean at sunrise"
  },
  "parameters": {
    "duration": 5,
    "size": "720P"
  }
}'`,
                responseExample: `{
  "output": {
    "task_id": "ali-task-123",
    "task_status": "PENDING"
  },
  "request_id": "request-123"
}`
            }
        case 35:
            return {
                title: t('modeType.35'),
                endpoint: '/tasks/{task_id}',
                method: 'GET',
                responseFormat: 'json',
                requestExample: `curl --request GET \\
--url ${apiEndpoint}/api/v1/tasks/ali-task-123 \\
--header "Authorization: Bearer $token"`,
                responseExample: `{
  "output": {
    "task_id": "ali-task-123",
    "task_status": "SUCCEEDED",
    "video_url": "https://example.com/video.mp4"
  }
}`
            }
        case 36:
            return {
                title: t('modeType.36'),
                endpoint: '/contents/generations/tasks',
                method: 'POST',
                responseFormat: 'json',
                requestExample: `curl --request POST \\
--url ${apiEndpoint}/api/v3/contents/generations/tasks \\
--header "Authorization: Bearer $token" \\
--header 'Content-Type: application/json' \\
--data '{
  "model": "${modelConfig.model}",
  "content": [
    {
      "type": "text",
      "text": "A calm ocean at sunrise"
    }
  ],
  "duration": 5,
  "resolution": "720p",
  "ratio": "16:9"
}'`,
                responseExample: `{
  "id": "doubao-task-123",
  "model": "${modelConfig.model}",
  "status": "queued"
}`
            }
        case 37:
            return {
                title: t('modeType.37'),
                endpoint: '/contents/generations/tasks/{task_id}',
                method: 'GET',
                responseFormat: 'json',
                requestExample: `curl --request GET \\
--url ${apiEndpoint}/api/v3/contents/generations/tasks/doubao-task-123 \\
--header "Authorization: Bearer $token"`,
                responseExample: `{
  "id": "doubao-task-123",
  "status": "succeeded",
  "content": {
    "video_url": "https://example.com/video.mp4"
  }
}`
            }
        case 38:
            return {
                title: t('modeType.38'),
                endpoint: '/contents/generations/tasks/{task_id}',
                method: 'DELETE',
                responseFormat: 'text',
                requestExample: `curl --request DELETE \\
--url ${apiEndpoint}/api/v3/contents/generations/tasks/doubao-task-123 \\
--header "Authorization: Bearer $token"`,
                responseExample: 'No content'
            }
        default:
            return {
                title: t('modeType.0'),
                endpoint: '',
                method: '',
                responseFormat: 'json',
                requestExample: 'unknown',
                responseExample: 'unknown'
            }
    }
}

const ApiDocDrawer: React.FC<ApiDocDrawerProps> = ({ isOpen, onClose, modelConfig }) => {
    const { t } = useTranslation()
    const apiDoc = getApiDocContent(modelConfig, "", t)

    return (
        <Sheet open={isOpen} onOpenChange={(open) => {
            if (!open) onClose()
        }}>
            <SheetContent 
                side="right"
                className="p-0 overflow-hidden w-[400px] max-w-md"
            >
                <div className="flex flex-col gap-3 overflow-y-auto p-6 pb-3 h-[calc(100%-48px)]">
                    {/* method */}
                    <div className="flex gap-2.5 items-center">
                        <Badge className="flex px-2 py-0.5 justify-center items-center gap-0.5 rounded bg-blue-100 dark:bg-blue-900">
                            <span className="text-blue-600 dark:text-blue-300 font-medium text-xs leading-4 tracking-wide">
                                {apiDoc.method}
                            </span>
                        </Badge>
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="2"
                            height="18"
                            viewBox="0 0 2 18"
                            fill="none">
                            <path d="M1 1L1 17" stroke="#F0F1F6" strokeLinecap="round" className="dark:stroke-zinc-700" />
                        </svg>

                        <div className="flex gap-1 items-center justify-start">
                            {apiDoc.endpoint
                                .split('/')
                                .filter(Boolean)
                                .map((segment, index) => (
                                    <React.Fragment key={index}>
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            width="5"
                                            height="12"
                                            viewBox="0 0 5 12"
                                            fill="none">
                                            <path
                                                d="M4.42017 1.30151L0.999965 10.6984"
                                                stroke="#C4CBD7"
                                                strokeLinecap="round"
                                                className="dark:stroke-zinc-600"
                                            />
                                        </svg>
                                        <span className="text-gray-600 dark:text-gray-400 font-medium text-xs leading-4 tracking-wide">
                                            {segment}
                                        </span>
                                    </React.Fragment>
                                ))}
                        </div>
                    </div>

                    {/* request example and response example */}
                    <div className="flex flex-col gap-4 items-start w-full">
                        {/* request example */}
                        <div className="flex flex-col gap-2 items-start w-full">
                            <span className="text-gray-900 dark:text-gray-100 font-medium text-xs leading-4 tracking-wide">
                                {t('apiDoc.requestExample')}
                            </span>

                            {/* code */}
                            <div className="flex flex-col items-start justify-center w-full rounded-md overflow-hidden">
                                <div className="flex w-full p-2.5 justify-between items-center bg-[#232833]">
                                    <span className="text-white font-medium text-xs leading-4 tracking-wide">
                                        {'bash'}
                                    </span>

                                    <Button
                                        onClick={() => {
                                            navigator.clipboard.writeText(apiDoc.requestExample).then(
                                                () => {
                                                    toast.success(t('common.copied'))
                                                },
                                                (err) => {
                                                    toast.error(err?.message || t('common.copyFailed'))
                                                }
                                            )
                                        }}
                                        variant="ghost"
                                        size="icon"
                                        className="inline-flex p-1 min-w-0 h-[22px] w-[22px] justify-center items-center rounded bg-transparent hover:bg-white/10">
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            width="14"
                                            height="14"
                                            viewBox="0 0 14 14"
                                            fill="none">
                                            <path
                                                fillRule="evenodd"
                                                clipRule="evenodd"
                                                d="M2.86483 2.30131C2.73937 2.30131 2.61904 2.35115 2.53032 2.43987C2.44161 2.52859 2.39176 2.64891 2.39176 2.77438V7.5282C2.39176 7.65366 2.44161 7.77399 2.53032 7.86271C2.61904 7.95142 2.73937 8.00127 2.86483 8.00127H3.39304C3.7152 8.00127 3.97637 8.26243 3.97637 8.5846C3.97637 8.90676 3.7152 9.16793 3.39304 9.16793H2.86483C2.42995 9.16793 2.01288 8.99517 1.70537 8.68766C1.39786 8.38015 1.2251 7.96308 1.2251 7.5282V2.77438C1.2251 2.3395 1.39786 1.92242 1.70537 1.61491C2.01288 1.3074 2.42995 1.13464 2.86483 1.13464H7.61865C8.05354 1.13464 8.47061 1.3074 8.77812 1.61491C9.08563 1.92242 9.25839 2.33949 9.25839 2.77438V3.30258C9.25839 3.62475 8.99722 3.88592 8.67505 3.88592C8.35289 3.88592 8.09172 3.62475 8.09172 3.30258V2.77438C8.09172 2.64891 8.04188 2.52859 7.95316 2.43987C7.86444 2.35115 7.74412 2.30131 7.61865 2.30131H2.86483ZM6.56225 5.99872C6.30098 5.99872 6.08918 6.21052 6.08918 6.47179V11.2256C6.08918 11.4869 6.30098 11.6987 6.56225 11.6987H11.3161C11.5773 11.6987 11.7891 11.4869 11.7891 11.2256V6.47179C11.7891 6.21052 11.5773 5.99872 11.3161 5.99872H6.56225ZM4.92251 6.47179C4.92251 5.56619 5.65664 4.83206 6.56225 4.83206H11.3161C12.2217 4.83206 12.9558 5.56619 12.9558 6.47179V11.2256C12.9558 12.1312 12.2217 12.8653 11.3161 12.8653H6.56225C5.65664 12.8653 4.92251 12.1312 4.92251 11.2256V6.47179Z"
                                                fill="white"
                                                fillOpacity="0.8"
                                            />
                                        </svg>
                                    </Button>
                                </div>
                                <div className="p-3 bg-[#14181E] w-full">
                                    <CodeBlock code={apiDoc.requestExample} language="bash" />
                                </div>
                            </div>

                            {/* additional info */}
                            {apiDoc?.requestAdditionalInfo?.voices &&
                                apiDoc?.requestAdditionalInfo?.voices?.length > 0 && (
                                    <div className="flex flex-col p-2.5 w-full gap-2 items-start rounded-md border border-gray-200 dark:border-gray-700">
                                        <div className="flex gap-2">
                                            <span className="text-blue-600 dark:text-blue-400 font-medium text-xs leading-4 tracking-wide">
                                                {'voice'}
                                            </span>
                                            <div className="flex gap-1">
                                                <Badge variant="outline" className="bg-gray-100 dark:bg-gray-800 text-gray-500 dark:text-gray-400 font-medium text-xs">
                                                    {'enum<string>'}
                                                </Badge>
                                                <Badge className="bg-amber-50 dark:bg-amber-900/30 text-amber-600 dark:text-amber-400 font-medium text-xs border-0">
                                                    {t('apiDoc.voice')}
                                                </Badge>
                                            </div>
                                        </div>

                                        <div className="flex flex-col gap-1 items-start w-full">
                                            <span className="text-gray-500 dark:text-gray-400 font-medium text-xs leading-4 tracking-wide">
                                                {t('apiDoc.voiceValues')}
                                            </span>

                                            <div className="flex flex-wrap gap-2">
                                                {apiDoc?.requestAdditionalInfo?.voices?.map((voice) => (
                                                    <Badge
                                                        key={voice}
                                                        variant="outline"
                                                        className="bg-gray-100 dark:bg-gray-800 text-gray-900 dark:text-gray-300 font-medium text-xs">
                                                        {voice}
                                                    </Badge>
                                                ))}
                                            </div>
                                        </div>
                                    </div>
                                )}

                            {apiDoc?.requestAdditionalInfo?.formats &&
                                apiDoc?.requestAdditionalInfo?.formats?.length > 0 && (
                                    <div className="flex flex-col p-2.5 w-full gap-2 items-start rounded-md border border-gray-200 dark:border-gray-700">
                                        <div className="flex gap-2">
                                            <span className="text-blue-600 dark:text-blue-400 font-medium text-xs leading-4 tracking-wide">
                                                {'response_format'}
                                            </span>
                                            <div className="flex gap-1">
                                                <Badge variant="outline" className="bg-gray-100 dark:bg-gray-800 text-gray-500 dark:text-gray-400 font-medium text-xs">
                                                    {'enum<string>'}
                                                </Badge>
                                                <Badge variant="outline" className="bg-gray-100 dark:bg-gray-800 text-gray-500 dark:text-gray-400 font-medium text-xs">
                                                    {'default:mp3'}
                                                </Badge>
                                            </div>
                                        </div>

                                        <div className="flex flex-col gap-1 items-start w-full">
                                            <span className="text-gray-500 dark:text-gray-400 font-medium text-xs leading-4 tracking-wide">
                                                {t('apiDoc.responseFormatValues')}
                                            </span>

                                            <div className="flex flex-wrap gap-2">
                                                {apiDoc?.requestAdditionalInfo?.formats?.map((format) => (
                                                    <Badge
                                                        key={format}
                                                        variant="outline"
                                                        className="bg-gray-100 dark:bg-gray-800 text-gray-900 dark:text-gray-300 font-medium text-xs">
                                                        {format}
                                                    </Badge>
                                                ))}
                                            </div>
                                        </div>
                                    </div>
                                )}
                        </div>

                        {/* response example */}
                        <div className="flex flex-col gap-2 items-start w-full">
                            <span className="text-gray-900 dark:text-gray-100 font-medium text-xs leading-4 tracking-wide">
                                {t('apiDoc.responseExample')}
                            </span>

                            {/* code */}
                            <div className="flex flex-col items-start justify-center w-full rounded-md overflow-hidden">
                                <div className="flex w-full p-2.5 justify-between items-center bg-[#232833]">
                                    <span className="text-white font-medium text-xs leading-4 tracking-wide">
                                        {apiDoc.responseFormat}
                                    </span>

                                    <Button
                                        onClick={() => {
                                            navigator.clipboard.writeText(apiDoc.responseExample).then(
                                                () => {
                                                    toast.success(t('common.copied'))
                                                },
                                                (err) => {
                                                    toast.error(err?.message || t('common.copyFailed'))
                                                }
                                            )
                                        }}
                                        variant="ghost"
                                        size="icon"
                                        className="inline-flex p-1 min-w-0 h-[22px] w-[22px] justify-center items-center rounded bg-transparent hover:bg-white/10">
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            width="14"
                                            height="14"
                                            viewBox="0 0 14 14"
                                            fill="none">
                                            <path
                                                fillRule="evenodd"
                                                clipRule="evenodd"
                                                d="M2.86483 2.30131C2.73937 2.30131 2.61904 2.35115 2.53032 2.43987C2.44161 2.52859 2.39176 2.64891 2.39176 2.77438V7.5282C2.39176 7.65366 2.44161 7.77399 2.53032 7.86271C2.61904 7.95142 2.73937 8.00127 2.86483 8.00127H3.39304C3.7152 8.00127 3.97637 8.26243 3.97637 8.5846C3.97637 8.90676 3.7152 9.16793 3.39304 9.16793H2.86483C2.42995 9.16793 2.01288 8.99517 1.70537 8.68766C1.39786 8.38015 1.2251 7.96308 1.2251 7.5282V2.77438C1.2251 2.3395 1.39786 1.92242 1.70537 1.61491C2.01288 1.3074 2.42995 1.13464 2.86483 1.13464H7.61865C8.05354 1.13464 8.47061 1.3074 8.77812 1.61491C9.08563 1.92242 9.25839 2.33949 9.25839 2.77438V3.30258C9.25839 3.62475 8.99722 3.88592 8.67505 3.88592C8.35289 3.88592 8.09172 3.62475 8.09172 3.30258V2.77438C8.09172 2.64891 8.04188 2.52859 7.95316 2.43987C7.86444 2.35115 7.74412 2.30131 7.61865 2.30131H2.86483ZM6.56225 5.99872C6.30098 5.99872 6.08918 6.21052 6.08918 6.47179V11.2256C6.08918 11.4869 6.30098 11.6987 6.56225 11.6987H11.3161C11.5773 11.6987 11.7891 11.4869 11.7891 11.2256V6.47179C11.7891 6.21052 11.5773 5.99872 11.3161 5.99872H6.56225ZM4.92251 6.47179C4.92251 5.56619 5.65664 4.83206 6.56225 4.83206H11.3161C12.2217 4.83206 12.9558 5.56619 12.9558 6.47179V11.2256C12.9558 12.1312 12.2217 12.8653 11.3161 12.8653H6.56225C5.65664 12.8653 4.92251 12.1312 4.92251 11.2256V6.47179Z"
                                                fill="white"
                                                fillOpacity="0.8"
                                            />
                                        </svg>
                                    </Button>
                                </div>
                                <div className="p-3 bg-[#14181E] w-full">
                                    <CodeBlock code={apiDoc.responseExample} language={apiDoc.responseFormat === 'json' ? 'json' : 'text'} />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </SheetContent>
        </Sheet>
    )
}

export default ApiDocDrawer
