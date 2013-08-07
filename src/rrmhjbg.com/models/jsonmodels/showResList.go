package jsonmodels

const (
	RoleType   = 1
	DialogType = 2
	SceneType  = 3

	RoleFaceType     = 11
	RoleActionType   = 12
	RoleClothingType = 13
)

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

type DownRes struct {
	PicName           string `json:"picName"`
	SrcType           string `json:"srcType"`
	KeyName           string `json:"keyName"`
	ItemPicName       string `json:"itemPicName"`
	ActionItemPicName string `json:"actionItemPicName"`
	Direction         string `json:"direction"`
	DefaultFace       string `json:"defaultFace"`
	DefaultClothing   string `json:"defaultClothing"`
	RoleName          string `json:"roleName"`
	ClothingGroup     string `json:"clothingGroup"`
	ActionGroup       string `json:"actionGroup"`
	Color             string `json:"color"`
}

type NewDownRes struct {
	FileName  string    `json:"fileName"`
	ImgStruct []DownRes `json:"imgStruct"`
}

type ShowRoleInfo struct {
	OptCode     string `json:"optCode"`
	KeyName     string `json:"keyName"`
	ProfileName string `json:"profileName"`
	ProfilePic  string `json:"profilePic"`
	ProfileText string `json:"profileText"`
	ImgSuffix   string `json:"imgSuffix"`
	TipNum      string `json:"tipNum"`
}
