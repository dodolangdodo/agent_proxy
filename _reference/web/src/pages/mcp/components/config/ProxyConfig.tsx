import { useState } from 'react'
import { PublicMCPProxyConfig, PublicMCPProxyReusingParam } from '@/api/mcp'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { Textarea } from '@/components/ui/textarea'
import { Switch } from '@/components/ui/switch'

interface ProxyConfigProps {
  config: PublicMCPProxyConfig | undefined
  onChange: (config: PublicMCPProxyConfig) => void
}

const ProxyConfig = ({ config, onChange }: ProxyConfigProps) => {
  const [proxyConfig, setProxyConfig] = useState<PublicMCPProxyConfig>(
    config || {
      url: '',
      headers: {},
      querys: {},
      reusing: {}
    }
  )

  // 添加键值对的临时状态
  const [newHeaderKey, setNewHeaderKey] = useState('')
  const [newHeaderValue, setNewHeaderValue] = useState('')
  const [newQueryKey, setNewQueryKey] = useState('')
  const [newQueryValue, setNewQueryValue] = useState('')
  const [newReusingKey, setNewReusingKey] = useState('')
  const [newReusingParam, setNewReusingParam] = useState<PublicMCPProxyReusingParam>({
    name: '',
    description: '',
    required: false,
    type: 'header'
  })

  const handleURLChange = (url: string) => {
    const newConfig = { ...proxyConfig, url }
    setProxyConfig(newConfig)
    onChange(newConfig)
  }

  const addHeader = () => {
    if (!newHeaderKey.trim()) return
    
    const newHeaders = { 
      ...proxyConfig.headers, 
      [newHeaderKey]: newHeaderValue 
    }
    
    const newConfig = { 
      ...proxyConfig, 
      headers: newHeaders 
    }
    
    setProxyConfig(newConfig)
    onChange(newConfig)
    setNewHeaderKey('')
    setNewHeaderValue('')
  }

  const removeHeader = (key: string) => {
    const newHeaders = { ...proxyConfig.headers }
    delete newHeaders[key]
    
    const newConfig = { 
      ...proxyConfig, 
      headers: newHeaders 
    }
    
    setProxyConfig(newConfig)
    onChange(newConfig)
  }

  const addQuery = () => {
    if (!newQueryKey.trim()) return
    
    const newQuerys = { 
      ...proxyConfig.querys, 
      [newQueryKey]: newQueryValue 
    }
    
    const newConfig = { 
      ...proxyConfig, 
      querys: newQuerys 
    }
    
    setProxyConfig(newConfig)
    onChange(newConfig)
    setNewQueryKey('')
    setNewQueryValue('')
  }

  const removeQuery = (key: string) => {
    const newQuerys = { ...proxyConfig.querys }
    delete newQuerys[key]
    
    const newConfig = { 
      ...proxyConfig, 
      querys: newQuerys 
    }
    
    setProxyConfig(newConfig)
    onChange(newConfig)
  }

  const addReusingParam = () => {
    if (!newReusingKey.trim() || !newReusingParam.name.trim()) return
    
    const newReusingParams = { 
      ...proxyConfig.reusing, 
      [newReusingKey]: { ...newReusingParam } 
    }
    
    const newConfig = { 
      ...proxyConfig, 
      reusing: newReusingParams 
    }
    
    setProxyConfig(newConfig)
    onChange(newConfig)
    setNewReusingKey('')
    setNewReusingParam({
      name: '',
      description: '',
      required: false,
      type: 'header'
    })
  }

  const removeReusingParam = (key: string) => {
    const newReusingParams = { ...proxyConfig.reusing }
    delete newReusingParams[key]
    
    const newConfig = { 
      ...proxyConfig, 
      reusing: newReusingParams 
    }
    
    setProxyConfig(newConfig)
    onChange(newConfig)
  }

  return (
    <div className="space-y-6">
      <div className="space-y-2">
        <Label htmlFor="url">Backend URL <span className="text-red-500">*</span></Label>
        <Input
          id="url"
          value={proxyConfig.url}
          onChange={(e) => handleURLChange(e.target.value)}
          placeholder="https://example.com/api"
        />
        <p className="text-xs text-muted-foreground">The backend URL to proxy requests to</p>
      </div>

      <Tabs defaultValue="headers">
        <TabsList className="grid grid-cols-3">
          <TabsTrigger value="headers">Headers</TabsTrigger>
          <TabsTrigger value="query">Query Parameters</TabsTrigger>
          <TabsTrigger value="reusing">Reusing Parameters</TabsTrigger>
        </TabsList>

        <TabsContent value="headers" className="space-y-4 pt-4">
          <div className="space-y-2">
            <div className="flex gap-2">
              <Input
                placeholder="Header Name"
                value={newHeaderKey}
                onChange={(e) => setNewHeaderKey(e.target.value)}
              />
              <Input
                placeholder="Header Value"
                value={newHeaderValue}
                onChange={(e) => setNewHeaderValue(e.target.value)}
              />
              <Button type="button" onClick={addHeader}>Add</Button>
            </div>
          </div>

          {Object.keys(proxyConfig.headers).length === 0 ? (
            <div className="text-center text-muted-foreground py-4">
              No headers configured
            </div>
          ) : (
            <div className="space-y-2">
              {Object.entries(proxyConfig.headers).map(([key, value]) => (
                <div key={key} className="flex items-center gap-2 p-2 bg-muted rounded-md">
                  <div className="flex-1">
                    <div className="font-medium">{key}</div>
                    <div className="text-sm text-muted-foreground">{value}</div>
                  </div>
                  <Button variant="ghost" size="sm" onClick={() => removeHeader(key)}>
                    Remove
                  </Button>
                </div>
              ))}
            </div>
          )}
        </TabsContent>

        <TabsContent value="query" className="space-y-4 pt-4">
          <div className="space-y-2">
            <div className="flex gap-2">
              <Input
                placeholder="Parameter Name"
                value={newQueryKey}
                onChange={(e) => setNewQueryKey(e.target.value)}
              />
              <Input
                placeholder="Parameter Value"
                value={newQueryValue}
                onChange={(e) => setNewQueryValue(e.target.value)}
              />
              <Button type="button" onClick={addQuery}>Add</Button>
            </div>
          </div>

          {Object.keys(proxyConfig.querys).length === 0 ? (
            <div className="text-center text-muted-foreground py-4">
              No query parameters configured
            </div>
          ) : (
            <div className="space-y-2">
              {Object.entries(proxyConfig.querys).map(([key, value]) => (
                <div key={key} className="flex items-center gap-2 p-2 bg-muted rounded-md">
                  <div className="flex-1">
                    <div className="font-medium">{key}</div>
                    <div className="text-sm text-muted-foreground">{value}</div>
                  </div>
                  <Button variant="ghost" size="sm" onClick={() => removeQuery(key)}>
                    Remove
                  </Button>
                </div>
              ))}
            </div>
          )}
        </TabsContent>

        <TabsContent value="reusing" className="space-y-4 pt-4">
          <div className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="reusingKey">Parameter Key</Label>
              <Input
                id="reusingKey"
                placeholder="e.g., api_key"
                value={newReusingKey}
                onChange={(e) => setNewReusingKey(e.target.value)}
              />
            </div>
            
            <div className="space-y-2">
              <Label htmlFor="reusingName">Display Name</Label>
              <Input
                id="reusingName"
                placeholder="e.g., API Key"
                value={newReusingParam.name}
                onChange={(e) => setNewReusingParam({...newReusingParam, name: e.target.value})}
              />
            </div>
            
            <div className="space-y-2">
              <Label htmlFor="reusingDescription">Description</Label>
              <Textarea
                id="reusingDescription"
                placeholder="Describe what this parameter is for"
                value={newReusingParam.description}
                onChange={(e) => setNewReusingParam({...newReusingParam, description: e.target.value})}
              />
            </div>
            
            <div className="space-y-2">
              <Label htmlFor="reusingType">Parameter Type</Label>
              <Select
                value={newReusingParam.type}
                onValueChange={(value: 'header' | 'query') => setNewReusingParam({...newReusingParam, type: value})}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Select parameter type" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="header">Header</SelectItem>
                  <SelectItem value="query">Query Parameter</SelectItem>
                </SelectContent>
              </Select>
            </div>
            
            <div className="flex items-center space-x-2">
              <Switch
                id="required"
                checked={newReusingParam.required}
                onCheckedChange={(checked) => setNewReusingParam({...newReusingParam, required: checked})}
              />
              <Label htmlFor="required">Required</Label>
            </div>
            
            <Button type="button" onClick={addReusingParam}>
              Add Reusing Parameter
            </Button>
          </div>

          {Object.keys(proxyConfig.reusing).length === 0 ? (
            <div className="text-center text-muted-foreground py-4">
              No reusing parameters configured
            </div>
          ) : (
            <div className="space-y-2">
              {Object.entries(proxyConfig.reusing).map(([key, param]) => (
                <Card key={key}>
                  <CardHeader>
                    <CardTitle className="text-base flex justify-between">
                      <span>{key}</span>
                      <Button variant="ghost" size="sm" onClick={() => removeReusingParam(key)}>
                        Remove
                      </Button>
                    </CardTitle>
                  </CardHeader>
                  <CardContent>
                    <div className="space-y-1">
                      <div className="text-sm">
                        <span className="font-medium">Name:</span> {param.name}
                      </div>
                      <div className="text-sm">
                        <span className="font-medium">Type:</span> {param.type}
                      </div>
                      <div className="text-sm">
                        <span className="font-medium">Required:</span> {param.required ? 'Yes' : 'No'}
                      </div>
                      {param.description && (
                        <div className="text-sm">
                          <span className="font-medium">Description:</span> {param.description}
                        </div>
                      )}
                    </div>
                  </CardContent>
                </Card>
              ))}
            </div>
          )}
        </TabsContent>
      </Tabs>
    </div>
  )
}

export default ProxyConfig 