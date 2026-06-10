# 百度地图 MCP Server

> <https://github.com/baidu-maps/mcp>

## 概述

百度地图API现已全面兼容[MCP协议](https://modelcontextprotocol.io/)，是国内首家兼容MCP协议的地图服务商。

百度地图提供的MCP Server，包含10个符合MCP协议标准的API接口，涵盖逆地理编码、地点检索、路线规划等。

依赖`MCP Python SDK`和`MCP Typescript SDK`开发，任意支持MCP协议的智能体助手（如`Claude`、`Cursor`以及`千帆AppBuilder`等）都可以快速接入。

**强烈推荐通过[SSE](https://lbsyun.baidu.com/faq/api?title=mcpserver/quickstart)接入百度地图MCP Server, 以获得更低的延迟和更高的稳定性。请不要忘记在[控制台](https://lbsyun.baidu.com/apiconsole/key)为你的AK勾选上`MCP(SSE)`服务。**

## 工具

1. 地理编码 `map_geocode`
    - 描述: 将地址解析为对应的位置坐标, 地址结构越完整, 地址内容越准确, 解析的坐标精度越高
    - 参数: `address` 地址信息
    - 输出: `location` 纬经度坐标
  
2. 逆地理编码 `map_reverse_geocode`
    - 描述: 根据纬经度坐标, 获取对应位置的地址描述, 所在行政区划, 道路以及相关POI等信息
    - 参数:
      - `latitude` 纬度坐标
      - `longitude`经度坐标
    - 输出: `formatted_address`, `uid`, `addressComponent` 等语义化地址信息

3. 地点检索 `map_search_places`
    - 描述: 支持检索城市内的地点信息(最小到`city`级别), 也可支持圆形区域内的周边地点信息检索
    - 参数:
      - `query` 检索关键词, 可用名称或类型, 多关键字使用英文逗号隔开, 如: `query=天安门,美食`
      - `tag` 检索的类型偏好, 格式为`tag=美食`或者`tag=美食,酒店`
      - `region` 检索的行政区划, 格式为`region=cityname`或`region=citycode`
      - `location` 圆形检索中心点纬经度坐标, 格式为`location=lat,lng`
      - `radius` 圆形检索的半径
    - 输出: POI列表, 包含`name`, `location`, `address`等

4. 地点详情检索 `map_place_details`
    - 描述: 根据POI的uid，检索其相关的详情信息, 如评分、营业时间等（不同类型POI对应不同类别详情数据）
    - 参数: `uid`POI的唯一标识
    - 输出: POI详情, 包含`name`, `location`, `address`, `brand`, `price`等
  
5. 批量算路 `map_directions_matrix`
    - 描述: 根据起点和终点坐标计算路线规划距离和行驶时间，支持驾车、骑行、步行。步行时任意起终点之间的距离不得超过200KM，驾车批量算路一次最多计算100条路线，起终点个数之积不能超过100。
    - 参数:
      - `origins` 起点纬经度列表, 格式为`origins=lat,lng`，多个起点用`|`分隔
      - `destinations` 终点纬经度列表, 格式为`destinations=lat,lng`，多个终点用`|`分隔
      - `model` 算路类型，可选取值包括 `driving`, `walking`, `riding`，默认使用`driving`
    - 输出: 每条路线的耗时和距离, 包含`distance`, `duration`等

6. 路线规划 `map_directions`
    - 描述: 根据起终点位置名称或经纬度坐标规划出行路线和耗时, 可指定驾车、步行、骑行、公交等出行方式
    - 参数:
      - `origin` 起点位置名称或纬经度, 格式为`origin=lat,lng`
      - `destination` 终点位置名称或纬经度, 格式为`destination=lat,lng`
      - `model` 出行类型, 可选取值包括 `driving`, `walking`, `riding`, `transit`, 默认使用`driving`
    - 输出: 路线详情,包含`steps`, `distance`, `duration`等
  
7. 天气查询 `map_weather`
    - 描述: 通过行政区划或是经纬度坐标查询实时天气信息及未来5天天气预报
    - 参数:
      - `district_id` 行政区划编码
      - `location` 经纬度坐标, 格式为`location=lng, lat`
    - 输出: 天气信息, 包含`temperature`, `weather`, `wind`等

8. IP定位 `map_ip_location`
    - 描述: 通过所给IP获取具体位置信息和城市名称, 可用于定位IP或用户当前位置。可选参数`ip`，如果为空则获取本机IP地址（支持IPv4和IPv6）。
    - 参数:
      - `ip`（可选）需要定位的IP地址
    - 输出: 当前所在城市和城市中点`location`

9. 实时路况查询 `map_road_traffic`
    - 描述: 查询实时交通拥堵情况, 可通过指定道路名和区域形状(矩形, 多边形, 圆形)进行实时路况查询。
    - 参数:
      - `model` 路况查询类型 (可选值包括`road`, `bound`, `polygon`, `around`, 默认使用`road`)
      - `road_name` 道路名称和道路方向, `model=road`时必传 (如:`朝阳路南向北`)
      - `city` 城市名称或城市adcode, `model=road`时必传 (如:`北京市`)
      - `bounds` 区域左下角和右上角的纬经度坐标, `model=bound`时必传 (如:`39.9,116.4;39.9,116.4`)
      - `vertexes` 多边形区域的顶点纬经度坐标, `model=polygon`时必传 (如:`39.9,116.4;39.9,116.4;39.9,116.4;39.9,116.4`)
      - `center` 圆形区域的中心点纬经度坐标, `model=around`时必传 (如:`39.912078,116.464303`)
      - `radius` 圆形区域的半径(米), 取值`[1,1000]`, `model=around`时必传 (如:`200`)
    - 输出: 路况信息, 包含`road_name`, `traffic_condition`等

10. POI智能提取 `map_poi_extract`
    - 描述: 当所给的`API_KEY`带有**高级权限**才可使用, 根据所给文本内容提取其中的相关POI信息。
    - 参数: `text_content` 用于提取POI的文本描述信息 (完整的旅游路线，行程规划，景点推荐描述等文本内容, 例如: 新疆独库公路和塔里木湖太美了, 从独山子大峡谷到天山神秘大峡谷也是很不错的体验)
    - 输出：相关的POI信息，包含`name`, `location`等
