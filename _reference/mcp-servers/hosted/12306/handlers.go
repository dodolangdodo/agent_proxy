package train12306

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/bytedance/sonic"
	"github.com/mark3labs/mcp-go/mcp"
)

// handleGetTickets handles the get tickets tool
func (s *Server) handleGetTickets(
	ctx context.Context,
	request mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	date, ok := args["date"].(string)
	if !ok || date == "" {
		return nil, errors.New("date is required")
	}

	fromStation, ok := args["fromStation"].(string)
	if !ok || fromStation == "" {
		return nil, errors.New("fromStation is required")
	}

	toStation, ok := args["toStation"].(string)
	if !ok || toStation == "" {
		return nil, errors.New("toStation is required")
	}

	trainFilterFlags := ""
	if tff, ok := args["trainFilterFlags"].(string); ok {
		trainFilterFlags = tff
	}

	sortFlag := ""
	if sf, ok := args["sortFlag"].(string); ok {
		sortFlag = sf
	}

	sortReverse := false
	if sr, ok := args["sortReverse"].(bool); ok {
		sortReverse = sr
	}

	limitedNum := 0
	if ln, ok := args["limitedNum"].(float64); ok {
		limitedNum = int(ln)
	}

	// Check if date is valid
	if !s.checkDate(date) {
		return nil, errors.New("the date cannot be earlier than today")
	}

	// Check if stations exist
	if _, exists := s.stations[fromStation]; !exists {
		return nil, errors.New("from station not found")
	}

	if _, exists := s.stations[toStation]; !exists {
		return nil, errors.New("to station not found")
	}

	// Get cookies
	cookies, err := s.getCookie(ctx, APIBase)
	if err != nil {
		return nil, err
	}

	queryURL := fmt.Sprintf(
		"%s/otn/leftTicket/query?leftTicketDTO.train_date=%s&leftTicketDTO.from_station=%s&leftTicketDTO.to_station=%s&purpose_codes=ADULT",
		APIBase,
		url.QueryEscape(date),
		url.QueryEscape(fromStation),
		url.QueryEscape(toStation),
	)
	headers := map[string]string{
		"Cookie": s.formatCookies(cookies),
	}

	var queryResponse LeftTicketsQueryResponse
	if err := s.make12306Request(ctx, queryURL, headers, &queryResponse); err != nil {
		return nil, err
	}

	// Parse tickets data
	resultData, ok := queryResponse.Data["result"].([]any)
	if !ok {
		return nil, errors.New("invalid response format")
	}

	stationMap, ok := queryResponse.Data["map"].(map[string]any)
	if !ok {
		return nil, errors.New("invalid response format")
	}

	ticketsData := s.parseTicketsData(resultData)
	ticketsInfo := s.parseTicketsInfo(ticketsData, stationMap)

	// Filter and sort tickets
	filteredTicketsInfo := s.filterTicketsInfo(
		ticketsInfo,
		trainFilterFlags,
		sortFlag,
		sortReverse,
		limitedNum,
	)

	// Format response
	result := s.formatTicketsInfo(filteredTicketsInfo)

	return mcp.NewToolResultText(result), nil
}

// addGetInterlineTicketsTool adds the get interline tickets tool
func (s *Server) addGetInterlineTicketsTool() {
	s.AddTool(
		mcp.Tool{
			Name:        "get-interline-tickets",
			Description: "查询12306中转余票信息。尚且只支持查询前十条。",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]any{
					"date": map[string]any{
						"type":        "string",
						"description": "查询日期，格式为 \"yyyy-MM-dd\"。如果用户提供的是相对日期（如\"明天\"），请务必先调用 `get-current-date` 接口获取当前日期，并计算出目标日期。",
					},
					"fromStation": map[string]any{
						"type":        "string",
						"description": "出发地的 `station_code` 。必须是通过 `get-station-code-by-names` 或 `get-station-code-of-citys` 接口查询得到的编码，严禁直接使用中文地名。",
					},
					"toStation": map[string]any{
						"type":        "string",
						"description": "到达地的 `station_code` 。必须是通过 `get-station-code-by-names` 或 `get-station-code-of-citys` 接口查询得到的编码，严禁直接使用中文地名。",
					},
					"middleStation": map[string]any{
						"type":        "string",
						"description": "中转地的 `station_code` ，可选。必须是通过 `get-station-code-by-names` 或 `get-station-code-of-citys` 接口查询得到的编码，严禁直接使用中文地名。",
						"default":     "",
					},
					"showWZ": map[string]any{
						"type":        "boolean",
						"description": "是否显示无座车，默认不显示无座车。",
						"default":     false,
					},
					"trainFilterFlags": map[string]any{
						"type":        "string",
						"description": "车次筛选条件，默认为空。从以下标志中选取多个条件组合[G(高铁/城际),D(动车),Z(直达特快),T(特快),K(快速),O(其他),F(复兴号),S(智能动车组)]",
						"default":     "",
					},
					"sortFlag": map[string]any{
						"type":        "string",
						"description": "排序方式，默认为空，即不排序。仅支持单一标识。可选标志：[startTime(出发时间从早到晚), arriveTime(抵达时间从早到晚), duration(历时从短到长)]",
						"default":     "",
					},
					"sortReverse": map[string]any{
						"type":        "boolean",
						"description": "是否逆向排序结果，默认为false。仅在设置了sortFlag时生效。",
						"default":     false,
					},
					"limitedNum": map[string]any{
						"type":        "integer",
						"description": "返回的余票数量限制，默认为0，即不限制。",
						"default":     0,
						"minimum":     0,
					},
				},
				Required: []string{"date", "fromStation", "toStation"},
			},
		},
		s.handleGetInterlineTickets,
	)
}

// handleGetInterlineTickets handles the get interline tickets tool
func (s *Server) handleGetInterlineTickets(
	ctx context.Context,
	request mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	date, ok := args["date"].(string)
	if !ok || date == "" {
		return nil, errors.New("date is required")
	}

	fromStation, ok := args["fromStation"].(string)
	if !ok || fromStation == "" {
		return nil, errors.New("fromStation is required")
	}

	toStation, ok := args["toStation"].(string)
	if !ok || toStation == "" {
		return nil, errors.New("toStation is required")
	}

	middleStation := ""
	if ms, ok := args["middleStation"].(string); ok {
		middleStation = ms
	}

	showWZ := false
	if sw, ok := args["showWZ"].(bool); ok {
		showWZ = sw
	}

	trainFilterFlags := ""
	if tff, ok := args["trainFilterFlags"].(string); ok {
		trainFilterFlags = tff
	}

	sortFlag := ""
	if sf, ok := args["sortFlag"].(string); ok {
		sortFlag = sf
	}

	sortReverse := false
	if sr, ok := args["sortReverse"].(bool); ok {
		sortReverse = sr
	}

	limitedNum := 0
	if ln, ok := args["limitedNum"].(float64); ok {
		limitedNum = int(ln)
	}

	// Check if date is valid
	if !s.checkDate(date) {
		return nil, errors.New("the date cannot be earlier than today")
	}

	// Check if stations exist
	if _, exists := s.stations[fromStation]; !exists {
		return nil, errors.New("from station not found")
	}

	if _, exists := s.stations[toStation]; !exists {
		return nil, errors.New("to station not found")
	}

	// Get cookies
	cookies, err := s.getCookie(ctx, APIBase)
	if err != nil {
		return nil, err
	}

	// Query interline tickets
	queryParams := url.Values{
		"train_date":            {date},
		"from_station_telecode": {fromStation},
		"to_station_telecode":   {toStation},
		"middle_station":        {middleStation},
		"result_index":          {"0"},
		"can_query":             {"Y"},
		"isShowWZ":              {map[bool]string{true: "Y", false: "N"}[showWZ]},
		"purpose_codes":         {"00"},
		"channel":               {"E"},
	}

	queryURL := fmt.Sprintf(
		"%s%s?%s",
		APIBase,
		s.lcQueryPath,
		queryParams.Encode(),
	)
	headers := map[string]string{
		"Cookie": s.formatCookies(cookies),
	}

	var queryResponse InterlineQueryResponse
	if err := s.make12306Request(ctx, queryURL, headers, &queryResponse); err != nil {
		return nil, err
	}

	// Check if response contains error
	if _, ok := queryResponse.Data.(string); ok {
		return nil, fmt.Errorf("很抱歉，未查到相关的列车余票。(%s)", queryResponse.ErrorMsg)
	}

	// Parse response data
	dataMap, ok := queryResponse.Data.(map[string]any)
	if !ok {
		return nil, errors.New("invalid response format")
	}

	middleListData, ok := dataMap["middleList"].([]any)
	if !ok {
		return nil, errors.New("invalid response format")
	}

	// Parse interline data
	interlineData := s.parseInterlineData(middleListData)

	interlineInfo := s.parseInterlinesInfo(interlineData)

	// Filter and sort tickets
	filteredInterlineInfo := s.filterInterlineInfo(
		interlineInfo,
		trainFilterFlags,
		sortFlag,
		sortReverse,
		limitedNum,
	)

	// Format response
	result := s.formatInterlinesInfo(filteredInterlineInfo)

	return mcp.NewToolResultText(result), nil
}

// addGetTrainRouteStationsTool adds the get train route stations tool
func (s *Server) addGetTrainRouteStationsTool() {
	s.AddTool(
		mcp.Tool{
			Name:        "get-train-route-stations",
			Description: "查询特定列车车次在指定区间内的途径车站、到站时间、出发时间及停留时间等详细经停信息。当用户询问某趟具体列车的经停站时使用此接口。",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]any{
					"trainNo": map[string]any{
						"type":        "string",
						"description": "要查询的实际车次编号 `train_no`，例如 \"240000G10336\"，而非\"G1033\"。此编号通常可以从 `get-tickets` 的查询结果中获取，或者由用户直接提供。",
					},
					"fromStationTelecode": map[string]any{
						"type":        "string",
						"description": "该列车行程的**出发站**的 `station_telecode` (3位字母编码`)。通常来自 `get-tickets` 结果中的 `telecode` 字段，或者通过 `get-station-code-by-names` 得到。",
					},
					"toStationTelecode": map[string]any{
						"type":        "string",
						"description": "该列车行程的**到达站**的 `station_telecode` (3位字母编码)。通常来自 `get-tickets` 结果中的 `telecode` 字段，或者通过 `get-station-code-by-names` 得到。",
					},
					"departDate": map[string]any{
						"type":        "string",
						"description": "列车从 `fromStationTelecode` 指定的车站出发的日期 (格式: yyyy-MM-dd)。如果用户提供的是相对日期，请务必先调用 `get-current-date` 解析。",
					},
				},
				Required: []string{
					"trainNo",
					"fromStationTelecode",
					"toStationTelecode",
					"departDate",
				},
			},
		},
		s.handleGetTrainRouteStations,
	)
}

// handleGetTrainRouteStations handles the get train route stations tool
func (s *Server) handleGetTrainRouteStations(
	ctx context.Context,
	request mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	trainNo, ok := args["trainNo"].(string)
	if !ok || trainNo == "" {
		return nil, errors.New("trainNo is required")
	}

	fromStationTelecode, ok := args["fromStationTelecode"].(string)
	if !ok || fromStationTelecode == "" {
		return nil, errors.New("fromStationTelecode is required")
	}

	toStationTelecode, ok := args["toStationTelecode"].(string)
	if !ok || toStationTelecode == "" {
		return nil, errors.New("toStationTelecode is required")
	}

	departDate, ok := args["departDate"].(string)
	if !ok || departDate == "" {
		return nil, errors.New("departDate is required")
	}

	// Get cookies
	cookies, err := s.getCookie(ctx, APIBase)
	if err != nil {
		return nil, err
	}

	// Query route stations
	queryParams := url.Values{
		"train_no":              {trainNo},
		"from_station_telecode": {fromStationTelecode},
		"to_station_telecode":   {toStationTelecode},
		"depart_date":           {departDate},
	}

	queryURL := fmt.Sprintf(
		"%s/otn/czxx/queryByTrainNo?%s",
		APIBase,
		queryParams.Encode(),
	)
	headers := map[string]string{
		"Cookie": s.formatCookies(cookies),
	}

	var queryResponse RouteQueryResponse
	if err := s.make12306Request(ctx, queryURL, headers, &queryResponse); err != nil {
		return nil, err
	}

	// Parse route stations data
	dataMap, ok := queryResponse.Data["data"].([]any)
	if !ok {
		return nil, errors.New("invalid response format")
	}

	routeStationsData := s.parseRouteStationsData(dataMap)

	routeStationsInfo := s.parseRouteStationsInfo(routeStationsData)

	if len(routeStationsInfo) == 0 {
		return nil, errors.New("未查询到相关车次信息。")
	}

	result, err := sonic.Marshal(routeStationsInfo)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(result)), nil
}
