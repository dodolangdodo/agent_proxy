package train12306

// TicketData represents raw ticket data from 12306 API
type TicketData struct {
	SecretStr            string `json:"secret_str"`
	ButtonTextInfo       string `json:"button_text_info"`
	TrainNo              string `json:"train_no"`
	StationTrainCode     string `json:"station_train_code"`
	StartStationTelecode string `json:"start_station_telecode"`
	EndStationTelecode   string `json:"end_station_telecode"`
	FromStationTelecode  string `json:"from_station_telecode"`
	ToStationTelecode    string `json:"to_station_telecode"`
	StartTime            string `json:"start_time"`
	ArriveTime           string `json:"arrive_time"`
	Lishi                string `json:"lishi"`
	CanWebBuy            string `json:"can_web_buy"`
	YpInfo               string `json:"yp_info"`
	StartTrainDate       string `json:"start_train_date"`
	TrainSeatFeature     string `json:"train_seat_feature"`
	LocationCode         string `json:"location_code"`
	FromStationNo        string `json:"from_station_no"`
	ToStationNo          string `json:"to_station_no"`
	IsSupportCard        string `json:"is_support_card"`
	ControlledTrainFlag  string `json:"controlled_train_flag"`
	GgNum                string `json:"gg_num"`
	GrNum                string `json:"gr_num"`
	QtNum                string `json:"qt_num"`
	RwNum                string `json:"rw_num"`
	RzNum                string `json:"rz_num"`
	TzNum                string `json:"tz_num"`
	WzNum                string `json:"wz_num"`
	YbNum                string `json:"yb_num"`
	YwNum                string `json:"yw_num"`
	YzNum                string `json:"yz_num"`
	ZeNum                string `json:"ze_num"`
	ZyNum                string `json:"zy_num"`
	SwzNum               string `json:"swz_num"`
	SrrbNum              string `json:"srrb_num"`
	YpEx                 string `json:"yp_ex"`
	SeatTypes            string `json:"seat_types"`
	ExchangeTrainFlag    string `json:"exchange_train_flag"`
	HoubuTrainFlag       string `json:"houbu_train_flag"`
	HoubuSeatLimit       string `json:"houbu_seat_limit"`
	YpInfoNew            string `json:"yp_info_new"`
	DwFlag               string `json:"dw_flag"`
	StopcheckTime        string `json:"stopcheck_time"`
	CountryFlag          string `json:"country_flag"`
	LocalArriveTime      string `json:"local_arrive_time"`
	LocalStartTime       string `json:"local_start_time"`
	BedLevelInfo         string `json:"bed_level_info"`
	SeatDiscountInfo     string `json:"seat_discount_info"`
	SaleTime             string `json:"sale_time"`
}

// TicketInfo represents processed ticket information
type TicketInfo struct {
	TrainNo             string   `json:"train_no"`
	StartTrainCode      string   `json:"start_train_code"`
	StartDate           string   `json:"start_date"`
	StartTime           string   `json:"start_time"`
	ArriveDate          string   `json:"arrive_date"`
	ArriveTime          string   `json:"arrive_time"`
	Lishi               string   `json:"lishi"`
	FromStation         string   `json:"from_station"`
	ToStation           string   `json:"to_station"`
	FromStationTelecode string   `json:"from_station_telecode"`
	ToStationTelecode   string   `json:"to_station_telecode"`
	Prices              []Price  `json:"prices"`
	DwFlag              []string `json:"dw_flag"`
}

// StationData represents station information
type StationData struct {
	StationID     string `json:"station_id"`
	StationName   string `json:"station_name"`
	StationCode   string `json:"station_code"`
	StationPinyin string `json:"station_pinyin"`
	StationShort  string `json:"station_short"`
	StationIndex  string `json:"station_index"`
	Code          string `json:"code"`
	City          string `json:"city"`
	R1            string `json:"r1"`
	R2            string `json:"r2"`
}

// Price represents seat price information
type Price struct {
	SeatName     string  `json:"seat_name"`
	Short        string  `json:"short"`
	SeatTypeCode string  `json:"seat_type_code"`
	Num          string  `json:"num"`
	Price        float64 `json:"price"`
	Discount     *int    `json:"discount"`
}

// RouteStationData represents raw route station data
type RouteStationData struct {
	ArriveTime       string `json:"arrive_time"`
	StationName      string `json:"station_name"`
	IsChina          string `json:"isChina"`
	StartTime        string `json:"start_time"`
	StopoverTime     string `json:"stopover_time"`
	StationNo        string `json:"station_no"`
	CountryCode      string `json:"country_code"`
	CountryName      string `json:"country_name"`
	IsEnabled        bool   `json:"isEnabled"`
	TrainClassName   string `json:"train_class_name,omitempty"`
	ServiceType      string `json:"service_type,omitempty"`
	EndStationName   string `json:"end_station_name,omitempty"`
	StartStationName string `json:"start_station_name,omitempty"`
	StationTrainCode string `json:"station_train_code,omitempty"`
}

// RouteStationInfo represents processed route station information
type RouteStationInfo struct {
	ArriveTime   string `json:"arrive_time"`
	StationName  string `json:"station_name"`
	StopoverTime string `json:"stopover_time"`
	StationNo    int    `json:"station_no"`
}

// InterlineData represents raw interline ticket data
type InterlineData struct {
	AllLishi          string                `json:"all_lishi"`
	AllLishiMinutes   int                   `json:"all_lishi_minutes"`
	ArriveDate        string                `json:"arrive_date"`
	ArriveTime        string                `json:"arrive_time"`
	EndStationCode    string                `json:"end_station_code"`
	EndStationName    string                `json:"end_station_name"`
	FirstTrainNo      string                `json:"first_train_no"`
	FromStationCode   string                `json:"from_station_code"`
	FromStationName   string                `json:"from_station_name"`
	FullList          []InterlineTicketData `json:"fullList"`
	IsHeatTrain       string                `json:"isHeatTrain"`
	IsOutStation      string                `json:"isOutStation"`
	LCWaitTime        string                `json:"lCWaitTime"`
	LishiFlag         string                `json:"lishi_flag"`
	MiddleDate        string                `json:"middle_date"`
	MiddleStationCode string                `json:"middle_station_code"`
	MiddleStationName string                `json:"middle_station_name"`
	SameStation       string                `json:"same_station"`
	SameTrain         string                `json:"same_train"`
	Score             float64               `json:"score"`
	ScoreStr          string                `json:"score_str"`
	SecretStr         string                `json:"scretstr"`
	SecondTrainNo     string                `json:"second_train_no"`
	StartTime         string                `json:"start_time"`
	TrainCount        int                   `json:"train_count"`
	TrainDate         string                `json:"train_date"`
	UseTime           string                `json:"use_time"`
	WaitTime          string                `json:"wait_time"`
	WaitTimeMinutes   int                   `json:"wait_time_minutes"`
}

// InterlineInfo represents processed interline information
type InterlineInfo struct {
	Lishi             string       `json:"lishi"`
	StartTime         string       `json:"start_time"`
	StartDate         string       `json:"start_date"`
	MiddleDate        string       `json:"middle_date"`
	ArriveDate        string       `json:"arrive_date"`
	ArriveTime        string       `json:"arrive_time"`
	FromStationCode   string       `json:"from_station_code"`
	FromStationName   string       `json:"from_station_name"`
	MiddleStationCode string       `json:"middle_station_code"`
	MiddleStationName string       `json:"middle_station_name"`
	EndStationCode    string       `json:"end_station_code"`
	EndStationName    string       `json:"end_station_name"`
	StartTrainCode    string       `json:"start_train_code"`
	FirstTrainNo      string       `json:"first_train_no"`
	SecondTrainNo     string       `json:"second_train_no"`
	TrainCount        int          `json:"train_count"`
	TicketList        []TicketInfo `json:"ticket_list"`
	SameStation       bool         `json:"same_station"`
	SameTrain         bool         `json:"same_train"`
	WaitTime          string       `json:"wait_time"`
}

// InterlineTicketData represents raw interline ticket data
type InterlineTicketData struct {
	ArriveTime           string `json:"arrive_time"`
	BedLevelInfo         string `json:"bed_level_info"`
	ControlledTrainFlag  string `json:"controlled_train_flag"`
	CountryFlag          string `json:"country_flag"`
	DayDifference        string `json:"day_difference"`
	DwFlag               string `json:"dw_flag"`
	EndStationName       string `json:"end_station_name"`
	EndStationTelecode   string `json:"end_station_telecode"`
	FromStationName      string `json:"from_station_name"`
	FromStationNo        string `json:"from_station_no"`
	FromStationTelecode  string `json:"from_station_telecode"`
	GgNum                string `json:"gg_num"`
	GrNum                string `json:"gr_num"`
	IsSupportCard        string `json:"is_support_card"`
	Lishi                string `json:"lishi"`
	LocalArriveTime      string `json:"local_arrive_time"`
	LocalStartTime       string `json:"local_start_time"`
	QtNum                string `json:"qt_num"`
	RwNum                string `json:"rw_num"`
	RzNum                string `json:"rz_num"`
	SeatDiscountInfo     string `json:"seat_discount_info"`
	SeatTypes            string `json:"seat_types"`
	SrrbNum              string `json:"srrb_num"`
	StartStationName     string `json:"start_station_name"`
	StartStationTelecode string `json:"start_station_telecode"`
	StartTime            string `json:"start_time"`
	StartTrainDate       string `json:"start_train_date"`
	StationTrainCode     string `json:"station_train_code"`
	SwzNum               string `json:"swz_num"`
	ToStationName        string `json:"to_station_name"`
	ToStationNo          string `json:"to_station_no"`
	ToStationTelecode    string `json:"to_station_telecode"`
	TrainNo              string `json:"train_no"`
	TrainSeatFeature     string `json:"train_seat_feature"`
	TrmsTrainFlag        string `json:"trms_train_flag"`
	TzNum                string `json:"tz_num"`
	WzNum                string `json:"wz_num"`
	YbNum                string `json:"yb_num"`
	YpInfo               string `json:"yp_info"`
	YwNum                string `json:"yw_num"`
	YzNum                string `json:"yz_num"`
	ZeNum                string `json:"ze_num"`
	ZyNum                string `json:"zy_num"`
}

// API response types
type QueryResponse struct {
	Data   any  `json:"data"`
	Status bool `json:"status"`
}

type LeftTicketsQueryResponse struct {
	Data     map[string]any `json:"data"`
	Messages string         `json:"messages"`
}

type InterlineQueryResponse struct {
	Data     any    `json:"data"`
	ErrorMsg string `json:"errorMsg"`
}

type RouteQueryResponse struct {
	Data             map[string]any `json:"data"`
	Messages         []string       `json:"messages"`
	ValidateMessages map[string]any `json:"validateMessages"`
}
