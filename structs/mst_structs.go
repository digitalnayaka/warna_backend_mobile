package structs

type MstLogStatus struct {
	Id     int64  `gorm:"default:'null'" json:"id"`
	Status string `gorm:"default:'null'" json:"status"`
}

type MstLogDesc struct {
	Id             int64  `gorm:"default:'null'" json:"id"`
	Description    string `gorm:"default:'null'" json:"description"`
	IdMstLogStatus int64  `gorm:"default:'null'" json:"id_mst_log_status"`
}

type MstLogs struct {
	Id         int64  `gorm:"default:'null'" json:"id"`
	IdCmsUsers int64  `gorm:"default:'null'" json:"id_cms_users"`
	IdModul    int64  `gorm:"default:'null'" json:"id_modul"`
	IdData     int64  `gorm:"default:'null'" json:"id_data"`
	Jenis      string `gorm:"default:'null'" json:"jenis"`
	CreatedAt  string `gorm:"default:'null'" json:"created_at"`
}

type MstVisumStatus struct {
	Id     int64  `gorm:"default:'null'" json:"id"`
	Status string `gorm:"default:'null'" json:"status"`
}

type MstJob struct {
	Id  int64  `gorm:"default:'null'" json:"id"`
	Job string `gorm:"default:'null'" json:"job"`
}
type MstDataSource struct {
	Id         int64  `gorm:"default:'null'" json:"id"`
	Datasource string `gorm:"default:'null'" json:"datasource"`
}

type MstCategoryAddress struct {
	Id       int64  `gorm:"default:'null'" json:"id"`
	Category string `gorm:"default:'null'" json:"category"`
}

type MstUnitUfi struct {
	Id                  int64  `gorm:"default:'null'" json:"id"`
	MstBranchBranchName string `gorm:"default:'null'" json:"mst_branch_branch_name"`
	Year                int    `gorm:"default:'null'" json:"year"`
	KodeUnit            string `gorm:"default:'null'" json:"kode_unit"`
	Merk                string `gorm:"default:'null'" json:"merk"`
	Type                string `gorm:"default:'null'" json:"type"`
	Model               string `gorm:"default:'null'" json:"model"`
	Otr                 int    `gorm:"default:'null'" json:"otr"`
}

type MstUnitYear struct {
	Id   int64 `gorm:"default:'null'" json:"id"`
	Year int   `gorm:"default:'null'" json:"year"`
}

type MstUnitMerk struct {
	Id   int64  `gorm:"default:'null'" json:"id"`
	Merk string `gorm:"default:'null'" json:"merk"`
}

type MstReligion struct {
	Id    int64  `gorm:"default:'null'" json:"id"`
	Agama string `gorm:"default:'null'" json:"agama"`
}

type ContactMstStatusMarital struct {
	Id     int64  `gorm:"default:'null'" json:"id"`
	Status string `gorm:"default:'null'" json:"status"`
}

type MstPlace struct {
	Id     int64  `gorm:"default:'null'" json:"id"`
	Status string `gorm:"default:'null'" json:"status"`
}

type ContactMstStatusEmployee struct {
	Id     int64  `gorm:"default:'null'" json:"id"`
	Status string `gorm:"default:'null'" json:"status"`
}

type MstCabangFif struct {
	Id         int64  `gorm:"default:'null'" json:"id"`
	CodeCabFif string `gorm:"default:'null'" json:"code_cab_fif"`
	BranchName string `gorm:"default:'null'" json:"branch_name"`
	CodePosFif string `gorm:"default:'null'" json:"code_pos_fif"`
	PosName    string `gorm:"default:'null'" json:"pos_name"`
}
type MstNeed struct {
	Id   int64  `gorm:"default:'null'" json:"id"`
	Need string `gorm:"default:'null'" json:"need"`
}
