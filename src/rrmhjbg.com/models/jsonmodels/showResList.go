package jsonmodels

type ShowResList struct {
	OptCode   string `json:"optCode"`
	SrcType   string `json:"srcType"`
	ListCount string `json:"listCount"`
	PageIndex string `json:"pageIndex"`
	PageSize  string `json:"pageSize"`
	ImgSuffix string `json:"imgSuffix"`
	ListArry  []Res  `json:"listArry"`
}

type Res struct {
	KeyName     string `json:"keyName"`
	ItemPic     string `json:"itemPic"`
	IsDown      string `json:"isDown"`
	TipNum      string `json:"tipNum"`
	ProfileName string `json:"profileName"`
	ProfilePic  string `json:"profilePic"`
	ProfileText string `json:"profileText"`
}
