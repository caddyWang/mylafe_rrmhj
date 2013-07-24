package resource

type SrcUserDownloaded struct {
	Id         string "_id"
	Uid        string
	RoleInfo   []string
	DialogInfo []string
	SceneInfo  []string

	RoleFaceInfo     []string
	RoleActionInfo   []string
	RoleClothingInfo []string
}
