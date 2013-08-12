package jsonmodels

type VersionInfo struct {
	HasNewVer      string   `json:"hasNewVer"`
	ImgPrefix      string   `json:"imgPrefix"`
	ImgTip         string   `json:"imgTip"`
	VerText        []string `json:"verText"`
	VerInt         string   `json:"verInt"`
	VerAndroidDown string   `json:"verAndroidDown"`
	VerIosDown     string   `json:"verIosDown"`
}
