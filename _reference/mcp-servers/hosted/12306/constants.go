package train12306

const (
	VERSION          = "0.3.2"
	APIBase          = "https://kyfw.12306.cn"
	WebURL           = "https://www.12306.cn/index/"
	LCQueryInitURL   = "https://kyfw.12306.cn/otn/lcQuery/init"
	DefaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
	TimeZone         = "Asia/Shanghai"
)

var (
	// Missing stations that are not in the official station list
	MissingStations = []StationData{
		{
			StationID:     "@cdd",
			StationName:   "成  都东",
			StationCode:   "WEI",
			StationPinyin: "chengdudong",
			StationShort:  "cdd",
			StationIndex:  "",
			Code:          "1707",
			City:          "成都",
			R1:            "",
			R2:            "",
		},
	}

	// Seat type mappings
	SeatShortTypes = map[string]string{
		"swz":  "商务座",
		"tz":   "特等座",
		"zy":   "一等座",
		"ze":   "二等座",
		"gr":   "高软卧",
		"srrb": "动卧",
		"rw":   "软卧",
		"yw":   "硬卧",
		"rz":   "软座",
		"yz":   "硬座",
		"wz":   "无座",
		"qt":   "其他",
		"gg":   "",
		"yb":   "",
	}

	SeatTypes = map[string]struct {
		Name  string
		Short string
	}{
		"9":  {"商务座", "swz"},
		"P":  {"特等座", "tz"},
		"M":  {"一等座", "zy"},
		"D":  {"优选一等座", "zy"},
		"O":  {"二等座", "ze"},
		"S":  {"二等包座", "ze"},
		"6":  {"高级软卧", "gr"},
		"A":  {"高级动卧", "gr"},
		"4":  {"软卧", "rw"},
		"I":  {"一等卧", "rw"},
		"F":  {"动卧", "rw"},
		"3":  {"硬卧", "yw"},
		"J":  {"二等卧", "yw"},
		"2":  {"软座", "rz"},
		"1":  {"硬座", "yz"},
		"W":  {"无座", "wz"},
		"WZ": {"无座", "wz"},
		"H":  {"其他", "qt"},
	}

	DWFlags = []string{
		"智能动车组",
		"复兴号",
		"静音车厢",
		"温馨动卧",
		"动感号",
		"支持选铺",
		"老年优惠",
	}

	TicketDataKeys = []string{
		"secret_Sstr", "button_text_info", "train_no", "station_train_code",
		"start_station_telecode", "end_station_telecode", "from_station_telecode",
		"to_station_telecode", "start_time", "arrive_time", "lishi", "canWebBuy",
		"yp_info", "start_train_date", "train_seat_feature", "location_code",
		"from_station_no", "to_station_no", "is_support_card", "controlled_train_flag",
		"gg_num", "gr_num", "qt_num", "rw_num", "rz_num", "tz_num", "wz_num",
		"yb_num", "yw_num", "yz_num", "ze_num", "zy_num", "swz_num", "srrb_num",
		"yp_ex", "seat_types", "exchange_train_flag", "houbu_train_flag",
		"houbu_seat_limit", "yp_info_new", "40", "41", "42", "43", "44", "45",
		"dw_flag", "47", "stopcheckTime", "country_flag", "local_arrive_time",
		"local_start_time", "52", "bed_level_info", "seat_discount_info", "sale_time", "56",
	}

	StationDataKeys = []string{
		"station_id", "station_name", "station_code", "station_pinyin",
		"station_short", "station_index", "code", "city", "r1", "r2",
	}
)

const (
	PriceStrLength    = 10
	DiscountStrLength = 5
)
