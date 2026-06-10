import { useState } from "react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import MCPList from "@/pages/mcp/components/MCPList";
import EmbedMCP from "@/pages/mcp/components/EmbedMCP";
import MCPConfig from "@/pages/mcp/components/MCPConfig";

const MCPPage = () => {
  const [activeTab, setActiveTab] = useState("list");

  return (
    <div className="container mx-auto p-4 space-y-4">
      <h1 className="text-2xl font-bold">MCP Management</h1>

      <Tabs value={activeTab} onValueChange={setActiveTab} className="w-full">
        <TabsList className="grid grid-cols-3 w-full max-w-md">
          <TabsTrigger value="list">MCP List</TabsTrigger>
          <TabsTrigger value="embed">Embed MCP</TabsTrigger>
          <TabsTrigger value="config">MCP Config</TabsTrigger>
        </TabsList>

        <TabsContent value="list">
          <MCPList />
        </TabsContent>

        <TabsContent value="embed">
          <EmbedMCP />
        </TabsContent>

        <TabsContent value="config">
          <MCPConfig />
        </TabsContent>
      </Tabs>
    </div>
  );
};

export default MCPPage;
