## 什么是Zhipu Web Search MCP服务?

> <https://github.com/THUDM>

Zhipu Web Search MCP Server是[智谱开放平台](https://zhipuaishengchan.datasink.sensorsdata.cn/t/ME)（BigModel.cn)推出的一个专给大模型用的搜索引擎，整合了5家搜索引擎供用户灵活对比切换，在传统搜索引擎网页抓取、排序的能力基础上，增强了意图识别能力，返回更适合大模型处理的结果（网页标题、网页URL、网页摘要、网站名称、网站图标等），帮助 AI 应用获得"动态知识获取"与"精准场景适配"的能力。

## 如何使用Zhipu Web Search MCP？

支持运行 MCP 协议的客户端，如Cursor、Cherry Studio等中配置，在[智谱 BigModel开放平台](https://zhipuaishengchan.datasink.sensorsdata.cn/t/ME)复制您的 API 密钥，并按照文档内容设置服务器命令。

## Zhipu Web Search MCP服务的关键特性

**实时联网搜索**： 实时检索全网信息和网页链接。

**意图增强检索**： 结合智谱自研的向量、语义匹配、时效性、内容质量度等模型，针对用户提问进行意图识别，优化提取query 搜索词。同时使用自研 semantic reranker 分析用户问题与给定搜索结果的相关度，结合时效性给出打分和排序，为用户提供准确可靠的搜索结果。

**支持指定域名搜索**：可以通过输入指定的域名搜索站内内容。提升搜索效率与结果相关性，满足专业研究、品牌监控、安全管理等个性化需求。

**灵活条数定制**：可依据需求自由设置 1 - 50 条的搜索结果数量；精准匹配用户对信息量的把控需求，避免信息冗余或不足 。

**精准时间筛选**：可以筛选一天、一周、一个月、一年或者不限时间发布的网站。帮助用户精准检索最新咨询和历史资料。满足用户多条件的检索需求。

**自主摘要调控**：对于网站摘要我们提供两种模式进行摘要的生成。medium模式会总结大概400-600字的网站摘要，而新增的high模式最多可以总结2500字的长上下文摘要共模型生成更加完备的回答。用户可以根据需要进行模式的调整。

**多引擎支持**：  整合智谱自研引擎及主流搜索引擎（Bing/搜狗/夸克/Jina.ai），开发者可以按场景需求灵活调用，发挥不同搜索工具的优势。

| 名称 | 介绍 |
|------|------|
| 智谱自研搜索基础版 | 提供基础搜索能力，超高性价比。 |
| 智谱自研搜索 Pro版 | 支持超长正文、更全面的搜索结果，多引擎降低空结果率提高搜索效率，搜索结果召回率更高、答案更准确。 |
| 搜狗搜索 | 内容全面，可以抓取腾讯生态（腾讯新闻、企鹅号）、知乎的内容；在百科、医疗垂类场景中答案权威度更高。 |
| 夸克搜索 | 支持指定范围搜索，比如可指定在金融、法律行业内搜索，提高召回准确性；全正文输出，提供覆盖 95% 网页的 3000字 长正文服务，无需再次解析URL。 |
| Jina AI 搜索 | 输出结果直接精炼，适合需要明确直接答案的场景；能精准解析复杂 HTML，并将其转换为干净的 Markdown 或 JSON 格式。 |
| Bing | 拥有庞大索引，涵盖数十亿网页及多种内容类型；搜索功能丰富，包括网页、图像、视频、新闻、购物、学术、地图等多种类型搜索。 |

**结构化输出**：返回适合LLM处理的数据格式（含标题/URL/摘要/网站名/图标等）

## Zhipu Web Search MCP服务的使用场景

**专业垂直研究：** 进行高质量网页检索、多源信息整合，如学术研究、法律、金融行业分析报告。

**商业情报雷达：** 实时数据跟踪与分析，如监控行业动态、竞争对手信息、市场趋势等。

**AI 助手/聊天机器人：** 提供实时信息搜索能力，确保回答准确性和时效性。

**消费决策与规划设计：** 根据最新天气、新闻、车票等信息、进行多个选项对比，寻找最优解。

**人才画像与简历优化：** 对于hr，搜索对比市场上的jd招聘信息，帮助生成或优化jd；对于求职者，搜索对比简历与目标岗位JD及相关行业成功简历的模板，帮助生成/优化简历。

## 常见问题解答

**Q：如何获取API Key？**

A：需在[智谱BigModel开放平台](https://zhipuaishengchan.datasink.sensorsdata.cn/t/ME)注册开发者账号，获取 API Key。

**Q：使用Zhipu Web Search MCP是否需要付费？**

A：您在[智谱BigModel开放平台](https://zhipuaishengchan.datasink.sensorsdata.cn/t/ME)可获得免费额度。若免费额度耗尽，则需要付费使用。有任何疑问可在智谱BigModel 开放平台中咨询客服。

**Q：支持哪些搜索引擎？**

A：支持智谱自研引擎及主流搜索引擎。智谱自研引擎：search_std（基础版）、search_pro（高阶版）。第三方引擎：search_pro_sogou （搜狗）、search_pro_quark（夸克）、search_pro_jina（Jina.ai ）、search_pro_bing（Bing）。

## 安装教程

支持运行 MCP 协议的客户端，如Cursor、Cherry Studio等。

点击获取[智谱 BigModel 开放平台的API Key](https://zhipuaishengchan.datasink.sensorsdata.cn/t/ME)。

**在Cursor中使用**

Cursor 0.45.6 版本提供了MCP功能，Cursor将作为MCP服务客户端使用MCP服务，在Cursor中通过简单的配置就可以完成MCP服务的接入。

操作路径：Cursor设置-> 【Features】-> 【MCP Servers】。

**配置MCP服务器**

```json
{
    "mcpServers": {
        "zhipu-web-search-sse": {
            "url": "https://open.bigmodel.cn/api/mcp/web_search/sse?Authorization= YOUR API Key"
        }
    }
}
```
