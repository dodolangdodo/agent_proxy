package baidumap

// Response 百度地图API基础响应结构
type Response struct {
	Status  int    `json:"status"`
	Msg     string `json:"msg,omitempty"`
	Message string `json:"message,omitempty"`
}

// GeocodeResponse 地理编码响应
type GeocodeResponse struct {
	Response
	Result struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		Precise       int    `json:"precise"`
		Confidence    int    `json:"confidence"`
		Comprehension int    `json:"comprehension"`
		Level         string `json:"level"`
	} `json:"result"`
}

// ReverseGeocodeResponse 逆地理编码响应
type ReverseGeocodeResponse struct {
	Response
	Result struct {
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		FormattedAddress    string `json:"formatted_address"`
		FormattedAddressPOI string `json:"formatted_address_poi"`
		Business            string `json:"business"`
		BusinessInfo        []struct {
			Name     string `json:"name"`
			Location struct {
				Lng float64 `json:"lng"`
				Lat float64 `json:"lat"`
			} `json:"location"`
			Adcode    int     `json:"adcode"`
			Distance  float64 `json:"distance"`
			Direction string  `json:"direction"`
		} `json:"business_info"`
		AddressComponent struct {
			Country         string `json:"country"`
			CountryCode     int    `json:"country_code"`
			CountryCodeISO  string `json:"country_code_iso"`
			CountryCodeISO2 string `json:"country_code_iso2"`
			Province        string `json:"province"`
			City            string `json:"city"`
			CityLevel       int    `json:"city_level"`
			District        string `json:"district"`
			Town            string `json:"town"`
			TownCode        string `json:"town_code"`
			Distance        string `json:"distance"`
			Direction       string `json:"direction"`
			Adcode          string `json:"adcode"`
			Street          string `json:"street"`
			StreetNumber    string `json:"street_number"`
		} `json:"addressComponent"`
		EDZ struct {
			Name string `json:"name"`
		} `json:"edz"`
		POIs               []any  `json:"pois"`
		Roads              []any  `json:"roads"`
		POIRegions         []any  `json:"poiRegions"`
		SematicDescription string `json:"sematic_description"`
		CityCode           int    `json:"cityCode"`
	} `json:"result"`
}

// PlacesSearchResponse 地点搜索响应
type PlacesSearchResponse struct {
	Response
	ResultType string `json:"result_type,omitempty"`
	QueryType  string `json:"query_type,omitempty"`
	Results    []struct {
		Name     string `json:"name"`
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		Address   string `json:"address"`
		Province  string `json:"province"`
		City      string `json:"city"`
		Area      string `json:"area"`
		StreetID  string `json:"street_id,omitempty"`
		Telephone string `json:"telephone,omitempty"`
		Detail    int    `json:"detail"`
		UID       string `json:"uid"`
	} `json:"results,omitempty"`
	Result []struct {
		Name     string `json:"name"`
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		Address   string `json:"address"`
		Province  string `json:"province"`
		City      string `json:"city"`
		Area      string `json:"area"`
		StreetID  string `json:"street_id,omitempty"`
		Telephone string `json:"telephone,omitempty"`
		Detail    int    `json:"detail"`
		UID       string `json:"uid"`
	} `json:"result,omitempty"`
}

// PlaceDetailsResponse 地点详情响应
type PlaceDetailsResponse struct {
	Response
	Result struct {
		UID      string `json:"uid"`
		StreetID string `json:"street_id"`
		Name     string `json:"name"`
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		Address    string `json:"address"`
		Province   string `json:"province"`
		City       string `json:"city"`
		Area       string `json:"area"`
		Detail     int    `json:"detail"`
		DetailInfo *struct {
			Tag          string `json:"tag"`
			NaviLocation struct {
				Lng float64 `json:"lng"`
				Lat float64 `json:"lat"`
			} `json:"navi_location"`
			NewCatalog    string `json:"new_catalog"`
			ShopHours     string `json:"shop_hours"`
			DetailURL     string `json:"detail_url"`
			Type          string `json:"type"`
			OverallRating string `json:"overall_rating"`
			ImageNum      string `json:"image_num"`
			CommentNum    string `json:"comment_num"`
			ContentTag    string `json:"content_tag"`
		} `json:"detail_info,omitempty"`
	} `json:"result"`
}

// DistanceMatrixResponse 距离矩阵响应
type DistanceMatrixResponse struct {
	Response
	Result []struct {
		Distance struct {
			Text  string `json:"text"`
			Value string `json:"value"`
		} `json:"distance"`
		Duration struct {
			Text  string `json:"text"`
			Value string `json:"value"`
		} `json:"duration"`
	} `json:"result"`
}

// DirectionsResponse 路线规划响应
type DirectionsResponse struct {
	Response
	Result struct {
		Routes []struct {
			Distance int `json:"distance"`
			Duration int `json:"duration"`
			Steps    []struct {
				Instruction string `json:"instruction"`
			} `json:"steps"`
		} `json:"routes"`
	} `json:"result"`
}

// WeatherResponse 天气响应
type WeatherResponse struct {
	Response
	Result struct {
		Location struct {
			Province string `json:"province"`
			City     string `json:"city"`
			Name     string `json:"name"`
		} `json:"location"`
		Now struct {
			Text      string `json:"text"`
			Temp      int    `json:"temp"`
			FeelsLike int    `json:"feels_like"`
			RH        int    `json:"rh"`
			WindClass string `json:"wind_class"`
			WindDir   string `json:"wind_dir"`
			Uptime    int64  `json:"uptime"`
		} `json:"now"`
		Forecasts []struct {
			TextDay   string `json:"text_day"`
			TextNight string `json:"text_night"`
			High      int    `json:"high"`
			Low       int    `json:"low"`
			WCDay     string `json:"wc_day"`
			WDDay     string `json:"wd_day"`
			WCNight   string `json:"wc_night"`
			WDNight   string `json:"wd_night"`
			Date      string `json:"date"`
			Week      string `json:"week"`
		} `json:"forecasts"`
		Indexes []struct {
			Name   string `json:"name"`
			Brief  string `json:"brief"`
			Detail string `json:"detail"`
		} `json:"indexes,omitempty"`
		Alerts []struct {
			Type  string `json:"type"`
			Level string `json:"level"`
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"alerts"`
		ForecastHours []struct {
			Text      string  `json:"text"`
			TempFC    int     `json:"temp_fc"`
			WindClass string  `json:"wind_class"`
			RH        int     `json:"rh"`
			Prec1H    float64 `json:"prec_1h"`
			Clouds    int     `json:"clouds"`
			DataTime  int64   `json:"data_time"`
		} `json:"forecast_hours,omitempty"`
	} `json:"result"`
}

// IPLocationResponse IP定位响应
type IPLocationResponse struct {
	Response
	Address string `json:"address"`
	Content struct {
		Address       string `json:"address"`
		AddressDetail struct {
			City     string `json:"city"`
			CityCode int    `json:"city_code"`
			Province string `json:"province"`
		} `json:"address_detail"`
		Point struct {
			X string `json:"x"`
			Y string `json:"y"`
		} `json:"point"`
	} `json:"content"`
}

// RoadTrafficResponse 路况响应
type RoadTrafficResponse struct {
	Response
	Description string `json:"description"`
	Evaluation  struct {
		Status     int    `json:"status"`
		StatusDesc string `json:"status_desc"`
	} `json:"evaluation"`
	RoadTraffic struct {
		RoadName           string `json:"road_name"`
		CongestionSections []struct {
			SectionDesc        string  `json:"section_desc"`
			Status             int     `json:"status"`
			Speed              float64 `json:"speed"`
			CongestionDistance int     `json:"congestion_distance"`
			CongestionTrend    string  `json:"congestion_trend"`
		} `json:"congestion_sections,omitempty"`
	} `json:"road_traffic"`
}

// MarkSubmitResponse POI标注提交响应
type MarkSubmitResponse struct {
	Response
	Result struct {
		SessionID string `json:"session_id"`
		MapID     string `json:"map_id"`
	} `json:"result"`
}

// MarkResultResponse POI标注结果响应
type MarkResultResponse struct {
	Response
	Result struct {
		Data []struct {
			AnswerType string `json:"answer_type"`
			CreateTime string `json:"create_time"`
			Link       struct {
				Title   string `json:"title"`
				Desc    string `json:"desc"`
				JumpURL string `json:"jump_url"`
				Image   string `json:"image"`
				POI     []struct {
					UID       string  `json:"uid"`
					Name      string  `json:"name"`
					Location  any     `json:"location"`
					AdminInfo any     `json:"admin_info"`
					Price     float64 `json:"price"`
					ShopHours string  `json:"shop_hours"`
				} `json:"poi"`
			} `json:"link"`
		} `json:"data"`
	} `json:"result"`
}
