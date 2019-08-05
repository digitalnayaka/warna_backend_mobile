package structs

type LeadListing struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	FirstName string `gorm:"default:'null'" json:"first_name"`
	LastName string `gorm:"default:'null'" json:"last_name"`
	IdLeadMstStatus int `gorm:"default:'null'" json:"id_lead_mst_status"`
	Recall string `gorm:"default:'null'" json:"recall"`
	Revisit string `gorm:"default:'null'" json:"revisit"`
	IdMstLogDesc int `gorm:"default:'null'" json:"id_mst_log_desc"`
	IdMstLogStatus int `gorm:"default:'null'" json:"id_mst_log_status"`
	Description string `gorm:"default:'null'" json:"description"`
	Status string `gorm:"default:'null'" json:"status"`
	Favorite string `gorm:"default:'null'" json:"favorite"`
	CreatedAtLead string `gorm:"default:'null'" json:"created_at_lead"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`
	CreatedAtLeadVisum string `gorm:"default:'null'" json:"created_at_lead_visum"`
	VisitStatus string `gorm:"default:'null'" json:"visit_status"`



}

type DetailLeadTotal struct {
	LogTotal int `gorm:"default:'null'" json:"log_total"`
	VisumTotal int `gorm:"default:'null'" json:"visum_total"`
	NoteTotal int `gorm:"default:'null'" json:"note_total"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`

}

type DetailLead struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	IdMstOutlet int64 `gorm:"default:'null'" json:"id_mst_outlet"`
	MstOutletOutletName	string `gorm:"default:'null'" json:"mst_outlet_outlet_name"`
	IdMstDataSource	int64 `gorm:"default:'null'" json:"id_mst_data_source"`
	MstDataSourceDatasource	string `gorm:"default:'null'" json:"mst_data_source_datasource"`
	FirstName string `gorm:"default:'null'" json:"first_name"`
	LastName string `gorm:"default:'null'" json:"last_name"`
	IdMstJob int64 `gorm:"default:'null'" json:"id_mst_job"`
	MstJobJob string `gorm:"default:'null'" json:"mst_job_job"`
	IdLeadMstStatus int64 `gorm:"default:'null'" json:"id_lead_mst_status"`
	LeadMstStatusStatus string `gorm:"default:'null'" json:"lead_mst_status_status"`
	CmsUsersName string `gorm:"default:'null'" json:"cms_users_name"`
	LogTotal int `gorm:"default:'null'" json:"log_total"`
	IdMstLogDesc int64 `gorm:"default:'null'" json:"id_mst_log_desc"`
	Description	string `gorm:"default:'null'" json:"description"`
	IdMstLogStatus int64 `gorm:"default:'null'" json:"id_mst_log_status"`
	VisumTotal int `gorm:"default:'null'" json:"visum_total"`
	NoteTotal int `gorm:"default:'null'" json:"note_total"`
	LeadPhone []LeadPhone `gorm:"default:'null'" json:"lead_phone"`
	LeadAddress []LeadAddress `gorm:"default:'null'" json:"lead_address"`
	LeadUnit []LeadUnit `gorm:"default:'null'" json:"lead_unit"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`

}

type LeadPhone struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	IdLead int64 `gorm:"default:'null'" json:"id_lead"`
	Number string `gorm:"default:'null'" json:"number"`
	Status string `gorm:"default:'null'" json:"status"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`

}

type LeadAddress struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt string `gorm:"default:'null'" json:"updated_at"`
	IdLead	int64 `gorm:"default:'null'" json:"id_lead"`
	Address string `gorm:"default:'null'" json:"address"`
	IdMstAddress int64 `gorm:"default:'null'" json:"id_mst_address"`
	IdMstCategoryAddress int64 `gorm:"default:'null'" json:"id_mst_category_address"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
	UpdatedBy	int64 `gorm:"default:'null'" json:"updated_by"`
	Kelurahan string `gorm:"default:'null'" json:"kelurahan"`
	Kecamatan string `gorm:"default:'null'" json:"kecamatan"`
	Kabupaten string `gorm:"default:'null'" json:"kabupaten"`
	Provinsi string `gorm:"default:'null'" json:"provinsi"`
	Kodepos string `gorm:"default:'null'" json:"kodepos"`

}

type LeadUnit struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	IdMstProduct int64 `gorm:"default:'null'" json:"id_mst_product"`
	MstProductNama string `gorm:"default:'null'" json:"mst_product_nama"`
	MstProductStatus string `gorm:"default:'null'" json:"mst_product_status"`
	IdMstUnit int64 `gorm:"default:'null'" json:"id_mst_unit"`
	Nopol string `gorm:"default:'null'" json:"nopol"`
	TaxStatus string `gorm:"default:'null'" json:"tax_status"`
	Owner string `gorm:"default:'null'" json:"owner"`
	KodeUnit string `gorm:"default:'null'" json:"kode_unit"`
	Merk string `gorm:"default:'null'" json:"merk"`
	Type string `gorm:"default:'null'" json:"type"`
	Model string
	Year int `gorm:"default:'null'" json:"year"`
	Otr int `gorm:"default:'null'" json:"otr"`
	IdLeadProductDetail int64 `gorm:"default:'null'" json:"id_lead_product_detail"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`

}

type LeadLog struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	LeadFirstName string `gorm:"default:'null'" json:"lead_first_name"`
	LeadLastName string `gorm:"default:'null'" json:"lead_last_name"`
	Duration string `gorm:"default:'null'" json:"duration"`
	IdMstLogDesc int64 `gorm:"default:'null'" json:"id_mst_log_desc"`
	MstLogDescDescription string `gorm:"default:'null'" json:"mst_log_desc_description"`
	Recall string `gorm:"default:'null'" json:"recall"`
	CmsUsersName string `gorm:"default:'null'" json:"cms_users_name"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`
	IdMstLogStatus int64 `gorm:"default:'null'" json:"id_mst_log_status"`
	Status string `gorm:"default:'null'" json:"status"`

}
type LeadLogCreate struct {
	Id int64 `gorm:"default:'null'" json:"id"`
	IdLead int64 `gorm:"default:'null'" json:"id_lead"`
	Duration int `gorm:"default:'null'" json:"duration"`
	IdMstLogDesc int64 `gorm:"default:'null'" json:"id_mst_log_desc"`
	Recall string `gorm:"default:'null'" json:"recall"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`
} 

type LeadVisum struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	LeadFirstName string `gorm:"default:'null'" json:"lead_first_name"`
	LeadLastName string `gorm:"default:'null'" json:"lead_last_name"`
	MstVisumStatusStatus string `gorm:"default:'null'" json:"mst_visum_status_status"`
	Revisit string `gorm:"default:'null'" json:"revisit"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
	CmsUsersName string `gorm:"default:'null'" json:"cms_users_name"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`
	Photo string `gorm:"default:'null'" json:"photo"`

}

type LeadNote struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	Note string `gorm:"default:'null'" json:"note"`
	CmsUsersName string `gorm:"default:'null'" json:"cms_users_name"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`

}
type LeadNoteCreate struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	IdLead int64 `gorm:"default:'null'" json:"id_lead"`
	Note string `gorm:"default:'null'" json:"note"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`

}
type LeadVisumCreate struct {
	Id int64 `gorm:"default:'null'" json:"id"`
	IdLead int64 `gorm:"default:'null'" json:"id_lead"`
	IdMstVisumStatus int64 `gorm:"default:'null'" json:"id_mst_visum_status"`
	Revisit string `gorm:"default:'null'" json:"revisit"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`

}
type LeadVisumPhotoCreate struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	IdLead int64 `gorm:"default:'null'" json:"id_lead"`
	IdLeadVisum int64 `gorm:"default:'null'" json:"id_lead_visum"`
	Photo string `gorm:"default:'null'" json:"photo"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`

}
type LeadCreate struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	IdMstOutlet int64 `gorm:"default:'null'" json:"id_mst_outlet"`
	IdMstDataSource int64 `gorm:"default:'null'" json:"id_mst_data_source"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
	FirstName string `gorm:"default:'null'" json:"first_name"`
	LastName string `gorm:"default:'null'" json:"last_name"`
	IdMstJob int64 `gorm:"default:'null'" json:"id_mst_job"`
	IdLeadMstStatus int64 `gorm:"default:'null'" json:"id_lead_mst_status"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt string `gorm:"default:'null'" json:"updated_at"`
	UpdatedBy int64 `gorm:"default:'null'" json:"updated_by"`

}
type LeadAddressCreate struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	IdLead int64 `gorm:"default:'null'" json:"id_lead"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
	IdMstAddress int64 `gorm:"default:'null'" json:"id_mst_address"`
	IdMstCategoryAddress int64 `gorm:"default:'null'" json:"id_mst_category_address"`
	Address string `gorm:"default:'null'" json:"address"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt string `gorm:"default:'null'" json:"updated_at"`
	UpdatedBy int64 `gorm:"default:'null'" json:"updated_by"`

}

type LeadProductCreate struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	IdLead int64 `gorm:"default:'null'" json:"id_lead"`
	IdMstProduct int64 `gorm:"default:'null'" json:"id_mst_product"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt string `gorm:"default:'null'" json:"updated_at"`
	UpdatedBy int64 `gorm:"default:'null'" json:"updated_by"`


}

type LeadProductDetailCreate struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	IdLeadProduct int64 `gorm:"default:'null'" json:"id_lead_product"`
	IdMstUnit int64 `gorm:"default:'null'" json:"id_mst_unit"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
	Nopol string `gorm:"default:'null'" json:"nopol"`
	TaxStatus string `gorm:"default:'null'" json:"tax_status"`
	Owner string `gorm:"default:'null'" json:"owner"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt string `gorm:"default:'null'" json:"updated_at"`
	UpdatedBy int64 `gorm:"default:'null'" json:"updated_by"`

}

type LeadColab struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	IdLead int64 `json:"id_lead"`
	IdCmsUsers int64 `json:"id_cms_users"`
}

