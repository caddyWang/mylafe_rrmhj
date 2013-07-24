package resource

type SrcUserDownloaded struct {
	Id  string "_id"
	Uid string

	RoleInfo []struct {
		RoleName         string
		RoleFaceInfo     []string
		RoleActionInfo   []string
		RoleClothingInfo []string
	}

	DialogInfo []string
	SceneInfo  []string
}
