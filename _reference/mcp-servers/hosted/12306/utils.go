package train12306

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bytedance/sonic"
)

// checkDate checks if the date is not earlier than today
func (s *Server) checkDate(date string) bool {
	location, err := time.LoadLocation(TimeZone)
	if err != nil {
		return false
	}

	now := time.Now().In(location)
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)

	inputTime, err := time.ParseInLocation(time.DateOnly, date, location)
	if err != nil {
		return false
	}

	return inputTime.Unix() >= now.Unix()
}

// getCookie gets cookies from the specified URL
func (s *Server) getCookie(ctx context.Context, urlStr string) (map[string]string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", s.userAgent)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cookies := make(map[string]string)
	for _, cookie := range resp.Cookies() {
		cookies[cookie.Name] = cookie.Value
	}

	return cookies, nil
}

// formatCookies formats cookies map to cookie string
func (s *Server) formatCookies(cookies map[string]string) string {
	parts := make([]string, 0, len(cookies))
	for name, value := range cookies {
		parts = append(parts, fmt.Sprintf("%s=%s", name, value))
	}

	return strings.Join(parts, "; ")
}

// make12306Request makes a request to 12306 API
func (s *Server) make12306Request(
	ctx context.Context,
	urlStr string,
	headers map[string]string,
	result any,
) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", s.userAgent)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	return sonic.ConfigDefault.NewDecoder(resp.Body).Decode(result)
}

var (
	stationNameJSRegex = regexp.MustCompile(`\.(/script/core/common/station_name.+?\.js)`)
	stationDataJSRegex = regexp.MustCompile(`var station_names ='([^']+)'`)
	lcQueryPathRegex   = regexp.MustCompile(`var lc_search_url = '([^']+)'`)
)

// getStations loads station data from 12306
func (s *Server) getStations(ctx context.Context) (map[string]StationData, error) {
	// Get main page to find station JS file
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, WebURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", s.userAgent)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Find station name JS file path
	matches := stationNameJSRegex.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return nil, errors.New("station name JS file not found")
	}

	stationJSURL := WebURL + matches[1]

	// Get station JS file
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, stationJSURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", s.userAgent)

	resp, err = s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Extract station data from JS
	jsContent := string(jsBody)

	matches = stationDataJSRegex.FindStringSubmatch(jsContent)
	if len(matches) < 2 {
		return nil, errors.New("station data not found in JS file")
	}

	stationsData := s.parseStationsData(matches[1])

	// Add missing stations
	for _, station := range MissingStations {
		if _, exists := stationsData[station.StationCode]; !exists {
			stationsData[station.StationCode] = station
		}
	}

	return stationsData, nil
}

// getLCQueryPath gets the LC query path from the init page
func (s *Server) getLCQueryPath(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, LCQueryInitURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", s.userAgent)

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Extract LC query path
	matches := lcQueryPathRegex.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return "", errors.New("LC query path not found")
	}

	return matches[1], nil
}

// parseStationsData parses station data from the raw string
func (s *Server) parseStationsData(rawData string) map[string]StationData {
	result := make(map[string]StationData)
	dataArray := strings.Split(rawData, "|")

	for i := range len(dataArray) / 10 {
		group := dataArray[i*10 : i*10+10]
		if len(group) < 10 {
			continue
		}

		station := StationData{
			StationID:     group[0],
			StationName:   group[1],
			StationCode:   group[2],
			StationPinyin: group[3],
			StationShort:  group[4],
			StationIndex:  group[5],
			Code:          group[6],
			City:          group[7],
			R1:            group[8],
			R2:            group[9],
		}

		if station.StationCode != "" {
			result[station.StationCode] = station
		}
	}

	return result
}

// parseTicketsData parses raw ticket data from API response
func (s *Server) parseTicketsData(rawData []any) []TicketData {
	result := make([]TicketData, 0, len(rawData))

	for _, item := range rawData {
		itemStr, ok := item.(string)
		if !ok {
			continue
		}

		values := strings.Split(itemStr, "|")
		if len(values) < len(TicketDataKeys) {
			continue
		}

		ticket := TicketData{
			SecretStr:            values[0],
			ButtonTextInfo:       values[1],
			TrainNo:              values[2],
			StationTrainCode:     values[3],
			StartStationTelecode: values[4],
			EndStationTelecode:   values[5],
			FromStationTelecode:  values[6],
			ToStationTelecode:    values[7],
			StartTime:            values[8],
			ArriveTime:           values[9],
			Lishi:                values[10],
			CanWebBuy:            values[11],
			YpInfo:               values[12],
			StartTrainDate:       values[13],
			TrainSeatFeature:     values[14],
			LocationCode:         values[15],
			FromStationNo:        values[16],
			ToStationNo:          values[17],
			IsSupportCard:        values[18],
			ControlledTrainFlag:  values[19],
			GgNum:                values[20],
			GrNum:                values[21],
			QtNum:                values[22],
			RwNum:                values[23],
			RzNum:                values[24],
			TzNum:                values[25],
			WzNum:                values[26],
			YbNum:                values[27],
			YwNum:                values[28],
			YzNum:                values[29],
			ZeNum:                values[30],
			ZyNum:                values[31],
			SwzNum:               values[32],
			SrrbNum:              values[33],
			YpEx:                 values[34],
			SeatTypes:            values[35],
			ExchangeTrainFlag:    values[36],
			HoubuTrainFlag:       values[37],
			HoubuSeatLimit:       values[38],
			YpInfoNew:            values[39],
			DwFlag:               values[46],
			StopcheckTime:        values[48],
			CountryFlag:          values[49],
			LocalArriveTime:      values[50],
			LocalStartTime:       values[51],
			BedLevelInfo:         values[53],
			SeatDiscountInfo:     values[54],
			SaleTime:             values[55],
		}

		result = append(result, ticket)
	}

	return result
}

// parseTicketsInfo converts TicketData to TicketInfo
func (s *Server) parseTicketsInfo(
	ticketsData []TicketData,
	stationMap map[string]any,
) []TicketInfo {
	result := make([]TicketInfo, 0, len(ticketsData))

	// Convert station map
	nameMap := make(map[string]string)
	for code, name := range stationMap {
		if nameStr, ok := name.(string); ok {
			nameMap[code] = nameStr
		}
	}

	for _, ticket := range ticketsData {
		prices := s.extractPrices(ticket.YpInfoNew, ticket.SeatDiscountInfo, ticket)
		dwFlag := s.extractDWFlags(ticket.DwFlag)

		// Parse start time
		startTime, err := time.Parse("15:04", ticket.StartTime)
		if err != nil {
			continue
		}

		// Parse duration
		lishiParts := strings.Split(ticket.Lishi, ":")
		if len(lishiParts) != 2 {
			continue
		}

		durationHours, _ := strconv.Atoi(lishiParts[0])
		durationMinutes, _ := strconv.Atoi(lishiParts[1])

		// Parse start date
		startDate, err := time.Parse("20060102", ticket.StartTrainDate)
		if err != nil {
			continue
		}

		// Calculate start and arrive datetime
		startDateTime := time.Date(startDate.Year(), startDate.Month(), startDate.Day(),
			startTime.Hour(), startTime.Minute(), 0, 0, time.UTC)
		arriveDateTime := startDateTime.Add(
			time.Duration(durationHours)*time.Hour + time.Duration(durationMinutes)*time.Minute,
		)

		ticketInfo := TicketInfo{
			TrainNo:             ticket.TrainNo,
			StartTrainCode:      ticket.StationTrainCode,
			StartDate:           startDateTime.Format(time.DateOnly),
			ArriveDate:          arriveDateTime.Format(time.DateOnly),
			StartTime:           ticket.StartTime,
			ArriveTime:          ticket.ArriveTime,
			Lishi:               ticket.Lishi,
			FromStation:         nameMap[ticket.FromStationTelecode],
			ToStation:           nameMap[ticket.ToStationTelecode],
			FromStationTelecode: ticket.FromStationTelecode,
			ToStationTelecode:   ticket.ToStationTelecode,
			Prices:              prices,
			DwFlag:              dwFlag,
		}

		result = append(result, ticketInfo)
	}

	return result
}
