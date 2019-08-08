package structs

type OrderListing struct {
	CreatedAt            string `gorm:"default:'null'" json:"created_at"`
	Id                   int64  `gorm:"default:'null'" json:"id"`
	IdContact            int64  `gorm:"default:'null'" json:"id_contact"`
	ContactFirstName     string `gorm:"default:'null'" json:"contact_first_name"`
	ContactLastName      string `gorm:"default:'null'" json:"contact_last_name"`
	OrderMstStatusStatus string `gorm:"default:'null'" json:"order_mst_status_status"`
	Plafond              int    `gorm:"default:'null'" json:"plafond"`
	Model                string `gorm:"default:'null'" json:"model"`
}
type OrderCount struct {
	TotalOrder int64 `gorm:"default:'null'" json:"total_order"`
}
type Order struct {
	Id               int64  `gorm:"default:'null'" json:"id"`
	CreatedAt        string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt        string `gorm:"default:'null'" json:"updated_at"`
	IdMstDataSource  int64  `gorm:"default:'null'" json:"id_mst_data_source"`
	Category         string `gorm:"default:'null'" json:"category"`
	IdMstOutlet      int64  `gorm:"default:'null'" json:"id_mst_outlet"`
	IdMstCabangFif   int64  `gorm:"default:'null'" json:"id_mst_cabang_fif"`
	IdContact        int64  `gorm:"default:'null'" json:"id_contact"`
	StatusAddress    string `gorm:"default:'null'" json:"status_address"`
	IdMstProduct     int64  `gorm:"default:'null'" json:"id_mst_product"`
	NoOrder          string `gorm:"default:'null'" json:"no_order"`
	IdOrderMstStatus int64  `gorm:"default:'null'" json:"id_order_mst_status"`
	IdOrderMstReason int64  `gorm:"default:'null'" json:"id_order_mst_reason"`
	Survey           string `gorm:"default:'null'" json:"survey"`
	IdCmsUsers       int64  `gorm:"default:'null'" json:"id_cms_users"`
	UpdatedBy        int64  `gorm:"default:'null'" json:"updated_by"`
}
type OrderUfi struct {
	Id           int64  `gorm:"default:'null'" json:"id"`
	CreatedAt    string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt    string `gorm:"default:'null'" json:"updated_at"`
	IdOrder      int64  `gorm:"default:'null'" json:"id_order"`
	IdMstUnit    int64  `gorm:"default:'null'" json:"id_mst_unit"`
	Nopol        string `gorm:"default:'null'" json:"nopol"`
	TaxStatus    string `gorm:"default:'null'" json:"tax_status"`
	Owner        string `gorm:"default:'null'" json:"owner"`
	OtrTaksasi   int    `json:"otr_taksasi"`
	NomorTaksasi string `gorm:"default:'null'" json:"nomor_taksasi"`
	IdCmsUsers   int64  `gorm:"default:'null'" json:"id_cms_users"`
	UpdatedBy    int64  `gorm:"default:'null'" json:"updated_by"`
}
type OrderLoan struct {
	Id          int64  `gorm:"default:'null'" json:"id"`
	CreatedAt   string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt   string `gorm:"default:'null'" json:"updated_at"`
	IdOrder     int64  `gorm:"default:'null'" json:"id_order"`
	OtrCustom   int    `json:"otr_custom"`
	Plafond     int    `json:"plafond"`
	DownPayment int    `json:"down_payment"`
	Installment int    `json:"installment"`
	Tenor       int    `json:"tenor"`
	Need        string `gorm:"default:'null'" json:"need"`
	IdCmsUsers  int64  `gorm:"default:'null'" json:"id_cms_users"`
	UpdatedBy   int64  `gorm:"default:'null'" json:"updated_by"`
}
type OrderSurety struct {
	Id                         int64  `gorm:"default:'null'" json:"id"`
	CreatedAt                  string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt                  string `gorm:"default:'null'" json:"updated_at"`
	IdOrder                    int64  `gorm:"default:'null'" json:"id_order"`
	Name                       string `gorm:"default:'null'" json:"name"`
	BirthPlace                 string `gorm:"default:'null'" json:"birth_place"`
	BirthDate                  string `gorm:"default:'null'" json:"birth_date"`
	IdMstJob                   int64  `gorm:"default:'null'" json:"id_mst_job"`
	Company                    string `gorm:"default:'null'" json:"company"`
	IdContactMstStatusEmployee int64  `gorm:"default:'null'" json:"id_contact_mst_status_employee"`
	Position                   string `gorm:"default:'null'" json:"position"`
	WorkingTime                int    `json:"working_time"`
	Income                     int    `gorm:"json:"income"`
	Outlay                     int    `gorm:"json:"outlay"`
	IdCmsUsers                 int64  `gorm:"default:'null'" json:"id_cms_users"`
	UpdatedBy                  int64  `gorm:"default:'null'" json:"updated_by"`
}

type OrderSuretyDetail struct {
	Id                             int64  `gorm:"default:'null'" json:"id"`
	CreatedAt                      string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt                      string `gorm:"default:'null'" json:"updated_at"`
	IdOrder                        int64  `gorm:"default:'null'" json:"id_order"`
	Name                           string `gorm:"default:'null'" json:"name"`
	BirthPlace                     string `gorm:"default:'null'" json:"birth_place"`
	BirthDate                      string `gorm:"default:'null'" json:"birth_date"`
	IdMstJob                       int64  `gorm:"default:'null'" json:"id_mst_job"`
	Company                        string `gorm:"default:'null'" json:"company"`
	IdContactMstStatusEmployee     int64  `gorm:"default:'null'" json:"id_contact_mst_status_employee"`
	Position                       string `gorm:"default:'null'" json:"position"`
	WorkingTime                    int    `json:"working_time"`
	Income                         int    `gorm:"json:"income"`
	Outlay                         int    `gorm:"json:"outlay"`
	IdCmsUsers                     int64  `gorm:"default:'null'" json:"id_cms_users"`
	UpdatedBy                      int64  `gorm:"default:'null'" json:"updated_by"`
	MstJobJob                      string `json:"mst_job_job"`
	ContactMstStatusEmployeeStatus string `json:"contact_mst_status_employee_status"`
}

type OrderInfoDetail struct {
	Id                      int64  `json:"id"`
	IdContact               int64  `json:"id_contact"`
	MstDataSourceDatasource string `json:"mst_data_source_datasource"`
	MstOutletOutletName     string `json:"mst_outlet_outlet_name"`
	MstProductNama          string `json:"mst_product_nama"`
	NoOrder                 string `json:"no_order"`
	OrderMstStatusStatus    string `json:"order_mst_status_status"`
	CmsUsersName            string `json:"cms_users_name"`
	CreatedAt               string `json:"created_at"`
	Category                string `json:"category"`
	StatusAddress           string `json:"status_address"`
	Survey                  string `json:"survey"`
	IdMstCabangFif          int64  `json:"id_mst_cabang_fif"`
	CabangFif               string `json:"cabang_fif"`
	PosFif                  string `json:"pos_fif"`
	Reason                  string `json:"reason"`
}
type OrderProductUfiDetail struct {
	Id              int64  `json:"id"`
	MstUnitYear     int    `json:"mst_unit_year"`
	MstUnitKodeUnit string `json:"mst_unit_kode_unit"`
	MstUnitMerk     string `json:"mst_unit_merk"`
	MstUnitType     string `json:"mst_unit_type"`
	MstUnitModel    string `json:"mst_unit_model"`
	MstUnitOtr      int    `json:"mst_unit_otr"`
	Nopol           string `json:"nopol"`
	TaxStatus       string `json:"tax_status"`
	Owner           string `json:"owner"`
	OtrTaksasi      int    `json:"otr_taksasi"`
	NomorTaksasi    string `json:"nomor_taksasi"`
}
type OrderDetail struct {
	OrderInfoDetail       OrderInfoDetail       `json:"order_info_detail"`
	OrderProductUfiDetail OrderProductUfiDetail `json:"order_product_ufi_detail"`
	OrderLoan             OrderLoan             `json:"order_loan"`
	OrderSurety           OrderSuretyDetail     `json:"order_surety"`
	ContactDetailInfo     ContactDetailInfo     `json:"contact_detail_info"`
}
type OrderPhoto struct {
	Id          int64  `gorm:"default:'null'" json:"id"`
	CreatedAt   string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt   string `gorm:"default:'null'" json:"updated_at"`
	IdOrder     int64  `gorm:"default:'null'" json:"id_order"`
	Photo       string `gorm:"default:'null'" json:"photo"`
	Description string `gorm:"default:'null'" json:"description"`
	IdCmsUsers  int64  `gorm:"default:'null'" json:"id_cms_users"`
	UpdatedBy   int64  `gorm:"default:'null'" json:"updated_by"`
}

type OrderNoteCreate struct {
	Id         int64  `gorm:"default:'null'" json:"id"`
	IdOrder    int64  `gorm:"default:'null'" json:"id_order"`
	Note       string `gorm:"default:'null'" json:"note"`
	IdCmsUsers int64  `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt  string `gorm:"default:'null'" json:"created_at"`
}

type OrderDocument struct {
	Id          int64  `gorm:"default:'null'" json:"id"`
	IdOrder     int64  `gorm:"default:'null'" json:"id_order"`
	Photo       string `gorm:"default:'null'" json:"photo"`
	IdCmsUsers  int64  `gorm:"default:'null'" json:"id_cms_users"`
	Description string `gorm:"default:'null'" json:"description"`
	CreatedAt   string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt   string `gorm:"default:'null'" json:"updated_at"`
}
