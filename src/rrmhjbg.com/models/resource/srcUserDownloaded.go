package resource

type SrcUserDownloaded struct {
	Id  string "_id"
	Uid string

	RoleInfo         []string "roleInfo"
	RoleFaceInfo     []string "roleFaceInfo"
	RoleActionInfo   []string "roleActionInfo"
	RoleClothingInfo []string "roleClothingInfo"

	DialogInfo []string "dialogInfo"
	SceneInfo  []string "sceneInfo"
}
