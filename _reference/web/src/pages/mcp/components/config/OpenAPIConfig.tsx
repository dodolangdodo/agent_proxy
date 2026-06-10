import { useState } from 'react'
import { MCPOpenAPIConfig } from '@/api/mcp'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Switch } from '@/components/ui/switch'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'

interface OpenAPIConfigProps {
  config: MCPOpenAPIConfig | undefined
  onChange: (config: MCPOpenAPIConfig) => void
}

const OpenAPIConfig = ({ config, onChange }: OpenAPIConfigProps) => {
  const [openApiConfig, setOpenApiConfig] = useState<MCPOpenAPIConfig>(
    config || {
      openapi_spec: '',
      openapi_content: '',
      v2: false,
      server_addr: '',
      authorization: ''
    }
  )

  const handleChange = (field: keyof MCPOpenAPIConfig, value: string | boolean) => {
    const newConfig = { ...openApiConfig, [field]: value }
    setOpenApiConfig(newConfig)
    onChange(newConfig)
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center space-x-2">
        <Switch
          id="v2"
          checked={openApiConfig.v2}
          onCheckedChange={(checked) => handleChange('v2', checked)}
        />
        <Label htmlFor="v2">Use OpenAPI v2 (Swagger)</Label>
      </div>

      <Tabs defaultValue="url">
        <TabsList className="grid grid-cols-2">
          <TabsTrigger value="url">Specification URL</TabsTrigger>
          <TabsTrigger value="content">Specification Content</TabsTrigger>
        </TabsList>

        <TabsContent value="url" className="space-y-4 pt-4">
          <div className="space-y-2">
            <Label htmlFor="openapi_spec">OpenAPI Specification URL</Label>
            <Input
              id="openapi_spec"
              value={openApiConfig.openapi_spec}
              onChange={(e) => handleChange('openapi_spec', e.target.value)}
              placeholder="https://example.com/openapi.json"
            />
            <p className="text-xs text-muted-foreground">URL to your OpenAPI/Swagger specification</p>
          </div>
        </TabsContent>

        <TabsContent value="content" className="space-y-4 pt-4">
          <div className="space-y-2">
            <Label htmlFor="openapi_content">OpenAPI Specification Content</Label>
            <Textarea
              id="openapi_content"
              value={openApiConfig.openapi_content || ''}
              onChange={(e) => handleChange('openapi_content', e.target.value)}
              placeholder="Paste your OpenAPI/Swagger JSON or YAML here"
              className="min-h-[300px] font-mono"
            />
            <p className="text-xs text-muted-foreground">Paste your OpenAPI specification (JSON or YAML format)</p>
          </div>
        </TabsContent>
      </Tabs>

      <div className="space-y-2">
        <Label htmlFor="server_addr">Server Address (Optional)</Label>
        <Input
          id="server_addr"
          value={openApiConfig.server_addr || ''}
          onChange={(e) => handleChange('server_addr', e.target.value)}
          placeholder="https://api.example.com"
        />
        <p className="text-xs text-muted-foreground">Override the server address defined in the specification</p>
      </div>

      <div className="space-y-2">
        <Label htmlFor="authorization">Authorization (Optional)</Label>
        <Input
          id="authorization"
          value={openApiConfig.authorization || ''}
          onChange={(e) => handleChange('authorization', e.target.value)}
          placeholder="Bearer token123"
          type="password"
        />
        <p className="text-xs text-muted-foreground">Default authorization header to include with all requests</p>
      </div>
    </div>
  )
}

export default OpenAPIConfig 