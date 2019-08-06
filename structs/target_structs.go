package structs

type TargetListing struct {
	Id                   int64  `gorm:"default:'null'" json:"id"`
	FirstName            string `gorm:"default:'null'" json:"first_name"`
	LastName             string `gorm:"default:'null'" json:"last_name"`
	IdTargetMstStatus    int    `gorm:"default:'null'" json:"id_target_mst_status"`
	Category             string `gorm:"default:'null'" json:"category"`
	IdMstVisumStatus     int64  `gorm:"default:'null'" json:"id_mst_visum_status"`
	Revisit              string `gorm:"default:'null'" json:"revisit"`
	VisitStatus          string `gorm:"default:'null'" json:"visit_status"`
	Recall               string `gorm:"default:'null'" json:"recall"`
	IdMstLogDesc         int    `gorm:"default:'null'" json:"id_mst_log_desc"`
	IdMstLogStatus       int    `gorm:"default:'null'" json:"id_mst_log_status"`
	Description          string `gorm:"default:'null'" json:"description"`
	Status               string `gorm:"default:'null'" json:"status"`
	CreatedAtTarget      string `gorm:"default:'null'" json:"created_at_target_log"`
	CreatedAt            string `gorm:"default:'null'" json:"created_at"`
	CreatedAtTargetVisum string `gorm:"default:'null'" json:"created_at_target_visum"`
}

type DetailTargetTotal struct {
	LogTotal   int    `gorm:"default:'null'" json:"log_total"`
	VisumTotal int    `gorm:"default:'null'" json:"visum_total"`
	NoteTotal  int    `gorm:"default:'null'" json:"note_total"`
	CreatedAt  string `gorm:"default:'null'" json:"created_at"`
}

type DetailTarget struct {
	Id                      int64  `gorm:"default:'null'" json:"id"`
	MstDataSourceDatasource string `gorm:"default:'null'" json:"mst_data_source_datasource"`
	BusinessCode            string `gorm:"default:'null'" json:"business_code"`
	Category                string `gorm:"default:'null'" json:"category"`
	Priority                string `gorm:"default:'null'" json:"priority"`
	NoContract              string `gorm:"default:'null'" json:"no_contract"`
	Nopol                   string `gorm:"default:'null'" json:"nopol"`
	FirstName               string `gorm:"default:'null'" json:"first_name"`
	LastName                string `gorm:"default:'null'" json:"last_name"`
	Hp_1                    string `gorm:"default:'null'" json:"hp_1"`
	Hp_2                    string `gorm:"default:'null'" json:"hp_2"`
	Provider_1              string `gorm:"default:'null'" json:"provider_1"`
	Provider_2              string `gorm:"default:'null'" json:"provider_2"`
	Job                     string `gorm:"default:'null'" json:"job"`
	Address                 string `gorm:"default:'null'" json:"address"`
	Kelurahan               string `gorm:"default:'null'" json:"kelurahan"`
	Kecamatan               string `gorm:"default:'null'" json:"kecamatan"`
	Kabupaten               string `gorm:"default:'null'" json:"kabupaten"`
	Provinsi                string `gorm:"default:'null'" json:"provinsi"`
	IdTargetMstStatus       int64  `gorm:"default:'null'" json:"id_target_mst_status"`
	TargetMstStatusStatus   string `gorm:"default:'null'" json:"target_mst_status_status"`
	CmsUsersName            string `gorm:"default:'null'" json:"cms_users_name"`
	LogTotal                int    `gorm:"default:'null'" json:"log_total"`
	IdMstLogDesc            int64  `gorm:"default:'null'" json:"id_mst_log_desc"`
	Description             string `gorm:"default:'null'" json:"description"`
	IdTargetVisum           int64  `gorm:"default:'null'" json:"id_target_visum"`
	VisumStatus             string `gorm:"default:'null'" json:"visum_status"`
	IdMstLogStatus          int64  `gorm:"default:'null'" json:"id_mst_log_status"`
	VisumTotal              int    `gorm:"default:'null'" json:"visum_total"`
	NoteTotal               int    `gorm:"default:'null'" json:"note_total"`
	CreatedAt               string `gorm:"default:'null'" json:"created_at"`
}

type TargetLog struct {
	Id                    int64  `gorm:"default:'null'" json:"id"`
	TargetFirstName       string `gorm:"default:'null'" json:"target_first_name"`
	TargetLastName        string `gorm:"default:'null'" json:"target_last_name"`
	Duration              string `gorm:"default:'null'" json:"duration"`
	IdMstLogDesc          int64  `gorm:"default:'null'" json:"id_mst_log_desc"`
	MstLogDescDescription string `gorm:"default:'null'" json:"mst_log_desc_description"`
	Recall                string `gorm:"default:'null'" json:"recall"`
	CmsUsersName          string `gorm:"default:'null'" json:"cms_users_name"`
	CreatedAt             string `gorm:"default:'null'" json:"created_at"`
	IdMstLogStatus        int64  `gorm:"default:'null'" json:"id_mst_log_status"`
	Status                string `gorm:"default:'null'" json:"status"`
}
type TargetLogCreate struct {
	Id           int64  `gorm:"default:'null'" json:"id"`
	IdTarget     int64  `gorm:"default:'null'" json:"id_target"`
	Duration     int    `json:"duration"`
	IdMstLogDesc int64  `gorm:"default:'null'" json:"id_mst_log_desc"`
	Recall       string `gorm:"default:'null'" json:"recall"`
	IdCmsUsers   int64  `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt    string `gorm:"default:'null'" json:"created_at"`
}

type TargetVisum struct {
	Id                   int64  `gorm:"default:'null'" json:"id"`
	TargetFirstName      string `gorm:"default:'null'" json:"target_first_name"`
	TargetLastName       string `gorm:"default:'null'" json:"target_last_name"`
	MstVisumStatusStatus string `gorm:"default:'null'" json:"mst_visum_status_status"`
	Revisit              string `gorm:"default:'null'" json:"revisit"`
	IdCmsUsers           int64  `gorm:"default:'null'" json:"id_cms_users"`
	CmsUsersName         string `gorm:"default:'null'" json:"cms_users_name"`
	CreatedAt            string `gorm:"default:'null'" json:"created_at"`
	Photo                string `gorm:"default:'null'" json:"photo"`
}

type TargetNote struct {
	Id           int64  `gorm:"default:'null'" json:"id"`
	Note         string `gorm:"default:'null'" json:"note"`
	CmsUsersName string `gorm:"default:'null'" json:"cms_users_name"`
	CreatedAt    string `gorm:"default:'null'" json:"created_at"`
}
type TargetNoteCreate struct {
	Id         int64  `gorm:"default:'null'" json:"id"`
	IdTarget   int64  `gorm:"default:'null'" json:"id_target"`
	Note       string `gorm:"default:'null'" json:"note"`
	IdCmsUsers int64  `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt  string `gorm:"default:'null'" json:"created_at"`
}
type TargetVisumCreate struct {
	Id               int64  `gorm:"default:'null'" json:"id"`
	IdTarget         int64  `gorm:"default:'null'" json:"id_target"`
	IdMstVisumStatus int64  `gorm:"default:'null'" json:"id_mst_visum_status"`
	Revisit          string `gorm:"default:'null'" json:"revisit"`
	IdCmsUsers       int64  `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt        string `gorm:"default:'null'" json:"created_at"`
}
type TargetVisumPhotoCreate struct {
	Id            int64  `gorm:"default:'null'" json:"id"`
	IdTarget      int64  `gorm:"default:'null'" json:"id_target"`
	IdTargetVisum int64  `gorm:"default:'null'" json:"id_target_visum"`
	Photo         string `gorm:"default:'null'" json:"photo"`
	IdCmsUsers    int64  `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt     string `gorm:"default:'null'" json:"created_at"`
}
type TargetCreate struct {
	Id                int64  `gorm:"default:'null'" json:"id"`
	IdMstOutlet       int64  `gorm:"default:'null'" json:"id_mst_outlet"`
	IdMstDataSource   int64  `gorm:"default:'null'" json:"id_mst_data_source"`
	IdCmsUsers        int64  `gorm:"default:'null'" json:"id_cms_users"`
	FirstName         string `gorm:"default:'null'" json:"first_name"`
	LastName          string `gorm:"default:'null'" json:"last_name"`
	IdMstJob          int64  `gorm:"default:'null'" json:"id_mst_job"`
	IdTargetMstStatus int64  `gorm:"default:'null'" json:"id_target_mst_status"`
	CreatedAt         string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt         string `gorm:"default:'null'" json:"updated_at"`
	UpdatedBy         int64  `gorm:"default:'null'" json:"updated_by"`
}
type TargetAddressCreate struct {
	Id                   int64  `gorm:"default:'null'" json:"id"`
	IdTarget             int64  `gorm:"default:'null'" json:"id_target"`
	IdCmsUsers           int64  `gorm:"default:'null'" json:"id_cms_users"`
	IdMstAddress         int64  `gorm:"default:'null'" json:"id_mst_address"`
	IdMstCategoryAddress int64  `gorm:"default:'null'" json:"id_mst_category_address"`
	Address              string `gorm:"default:'null'" json:"address"`
	CreatedAt            string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt            string `gorm:"default:'null'" json:"updated_at"`
	UpdatedBy            int64  `gorm:"default:'null'" json:"updated_by"`
}

type TargetProductCreate struct {
	Id           int64  `gorm:"default:'null'" json:"id"`
	IdTarget     int64  `gorm:"default:'null'" json:"id_target"`
	IdMstProduct int64  `gorm:"default:'null'" json:"id_mst_product"`
	IdCmsUsers   int64  `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt    string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt    string `gorm:"default:'null'" json:"updated_at"`
	UpdatedBy    int64  `gorm:"default:'null'" json:"updated_by"`
}

type TargetProductDetailCreate struct {
	Id              int64  `gorm:"default:'null'" json:"id"`
	IdTargetProduct int64  `gorm:"default:'null'" json:"id_target_product"`
	IdMstUnit       int64  `gorm:"default:'null'" json:"id_mst_unit"`
	IdCmsUsers      int64  `gorm:"default:'null'" json:"id_cms_users"`
	Nopol           string `gorm:"default:'null'" json:"nopol"`
	TaxStatus       string `gorm:"default:'null'" json:"tax_status"`
	Owner           string `gorm:"default:'null'" json:"owner"`
	CreatedAt       string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt       string `gorm:"default:'null'" json:"updated_at"`
	UpdatedBy       int64  `gorm:"default:'null'" json:"updated_by"`
}
