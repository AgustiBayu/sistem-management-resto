package web

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type WebResponses struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Links []Link      `json:"link"`
}

type WebResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type WebRespon struct {
	Code int `json:"code"`
}
