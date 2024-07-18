package model

type Response struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	PageCount int    `json:"pagecount"`
	Total     int    `json:"total"`
	List      []Vod  `json:"list"`
	// Page      interface{} `json:"page"`
}

type Vod struct {
	VodID          interface{} `json:"vod_id"`
	TypeId         interface{} `json:"type_id"`
	VodName        string      `json:"vod_name"`
	VodSub         string      `json:"vod_sub"`
	VodEnname      string      `json:"vod_enname"`
	VodEnFF        string      `json:"vod_en"`
	VodStatus      int         `json:"vod_status"`
	VodTag         string      `json:"vod_tag"`
	VodClass       string      `json:"vod_class"`
	VodPic         string      `json:"vod_pic"`
	VodActor       string      `json:"vod_actor"`
	VodDirector    string      `json:"vod_director"`
	VodWriter      string      `json:"vod_writer"`
	VodBlurb       string      `json:"vod_blurb"`
	VodRemarks     string      `json:"vod_remarks"`
	VodTotal       int         `json:"vod_total"`
	VodArea        string      `json:"vod_area"`
	VodLang        string      `json:"vod_lang"`
	VodYear        string      `json:"vod_year"`
	VodState       string      `json:"vod_state"`
	VodScore       string      `json:"vod_score"`
	VodDoubanID    int         `json:"vod_douban_id"`
	VodDoubanScore string      `json:"vod_douban_score"`
	VodContent     string      `json:"vod_content"`
	VodPlayFrom    string      `json:"vod_play_from"`
	VodPlayNote    string      `json:"vod_play_note"`
	VodPlayURL     string      `json:"vod_play_url"`
	TypeName       string      `json:"type_name"`
}
