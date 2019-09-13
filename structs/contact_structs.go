package structs

type Contact struct {
	Id                        int64  `gorm:"default:'null'" json:"id"`
	Nik                       string `gorm:"default:'null'" json:"nik"`
	FirstName                 string `gorm:"default:'null'" json:"first_name"`
	LastName                  string `gorm:"default:'null'" json:"last_name"`
	BirthPlace                string `gorm:"default:'null'" json:"birth_place"`
	BirthDate                 string `gorm:"default:'null'" json:"birth_date"`
	Gender                    string `gorm:"default:'null'" json:"gender"`
	IdMstReligion             int64  `gorm:"default:'null'" json:"id_mst_religion"`
	IdContactMstStatusMarital int64  `gorm:"default:'null'" json:"id_contact_mst_status_marital"`
	IdMstJob                  int64  `gorm:"default:'null'" json:"id_mst_job"`
	IdMstDataSource           int64  `gorm:"default:'null'" json:"id_mst_data_source"`
	IdCmsUsers                int64  `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt                 string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt                 string `gorm:"default:'null'" json:"updated_at"`
	UpdatedBy                 int64  `gorm:"default:'null'" json:"updated_by"`
}

type ContactList struct {
	Id               int64  `gorm:"default:'null'" json:"id"`
	IdContact        int64  `gorm:"default:'null'" json:"id_contact"`
	ContactFirstName string `gorm:"default:'null'" json:"contact_first_name"`
	ContactLastName  string `gorm:"default:'null'" json:"contact_last_name"`
	ContactGender    string `gorm:"default:'null'" json:"contact_gender"`
	IdOrderMstStatus int64  `gorm:"default:'null'" json:"id_order_mst_status"`
	Status           string `gorm:"default:'null'" json:"status"`
}

type ContactCollaborate struct {
	Id         int64 `gorm:"default:'null'" json:"id"`
	IdContact  int64 `gorm:"default:'null'" json:"id_contact"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
}

type ContactPhone struct {
	Id         int64  `gorm:"default:'null'" json:"id"`
	CreatedAt  string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt  string `gorm:"default:'null'" json:"updated_at"`
	IdContact  int64  `gorm:"default:'null'" json:"id_contact"`
	Number     string `gorm:"default:'null'" json:"number"`
	Status     string `gorm:"default:'null'" json:"status"`
	IdCmsUsers int64  `gorm:"default:'null'" json:"id_cms_users"`
	UpdatedBy  int64  `gorm:"default:'null'" json:"updated_by"`
}

type ContactAddress struct {
	Id                   int64  `gorm:"default:'null'" json:"id"`
	CreatedAt            string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt            string `gorm:"default:'null'" json:"updated_at"`
	IdContact            int64  `gorm:"default:'null'" json:"id_contact"`
	Address              string `gorm:"default:'null'" json:"address"`
	Rt                   string `gorm:"default:'null'" json:"rt"`
	Rw                   string `gorm:"default:'null'" json:"rw"`
	IdMstAddress         int64  `gorm:"default:'null'" json:"id_mst_address"`
	IdMstCategoryAddress int64  `gorm:"default:'null'" json:"id_mst_category_address"`
	IdCmsUsers           int64  `gorm:"default:'null'" json:"id_cms_users"`
}

type ContactAdditional struct {
	Id                         int64  `gorm:"default:'null'" json:"id"`
	CreatedAt                  string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt                  string `gorm:"default:'null'" json:"updated_at"`
	IdContact                  int64  `gorm:"default:'null'" json:"id_contact"`
	Email                      string `gorm:"default:'null'" json:"email"`
	Mother                     string `gorm:"default:'null'" json:"mother"`
	Family                     string `gorm:"default:'null'" json:"family"`
	IdContactMstStatusPlace    int64  `gorm:"default:'null'" json:"id_contact_mst_status_place"`
	Company                    string `gorm:"default:'null'" json:"company"`
	IdContactMstStatusEmployee int64  `gorm:"default:'null'" json:"id_contact_mst_status_employee"`
	Position                   string `gorm:"default:'null'" json:"position"`
	WorkingTime                int    `json:"working_time"`
	Income                     int    `json:"income"`
	Outlay                     int    `json:"outlay"`
	IdCmsUsers                 int64  `gorm:"default:'null'" json:"id_cms_users"`
	UpdatedBy                  int64  `gorm:"default:'null'" json:"updated_by"`
}

type ContactDetailInfo struct {
	ContactInfo          ContactInfo            `gorm:"default:'null'" json:"contact_info"`
	ContactPhoneDetail   []ContactPhoneDetail   `gorm:"default:'null'" json:"contact_phone_detail"`
	ContactAddressDetail []ContactAddressDetail `gorm:"default:'null'" json:"contact_address_detail"`
}

type ContactInfo struct {
	Id                             int64  `gorm:"default:'null'" json:"id"`
	Nik                            string `gorm:"default:'null'" json:"nik"`
	FirstName                      string `gorm:"default:'null'" json:"first_name"`
	LastName                       string `gorm:"default:'null'" json:"last_name"`
	BirthPlace                     string `gorm:"default:'null'" json:"birth_place"`
	BirthDate                      string `gorm:"default:'null'" json:"birth_date"`
	Gender                         string `gorm:"default:'null'" json:"gender"`
	IdMstReligion                  int64  `gorm:"default:'null'" json:"id_mst_religion"`
	IdContactMstStatusMarital      int64  `gorm:"default:'null'" json:"id_contact_mst_status_marital"`
	IdMstJob                       int64  `gorm:"default:'null'" json:"id_mst_job"`
	IdMstDataSource                int64  `gorm:"default:'null'" json:"id_mst_data_source"`
	IdOrderMstStatus               int64  `gorm:"default:'null'" json:"id_order_mst_status"`
	Status                         string `gorm:"default:'null'" json:"status"`
	PotentialOrder                 int    `gorm:"default:'null'" json:"potential_order"`
	WonOrder                       int    `gorm:"default:'null'" json:"won_order"`
	MstReligionAgama               string `gorm:"default:'null'" json:"mst_religion_agama"`
	ContactMstStatusMaritalStatus  string `gorm:"default:'null'" json:"contact_mst_status_marital_status"`
	MstJobJob                      string `gorm:"default:'null'" json:"mst_job_job"`
	CmsUsersName                   string `gorm:"default:'null'" json:"cms_users_name"`
	MstDataSourceDatasource        string `gorm:"default:'null'" json:"mst_data_source_datasource"`
	Email                          string `gorm:"default:'null'" json:"email"`
	Mother                         string `gorm:"default:'null'" json:"mother"`
	Family                         string `gorm:"default:'null'" json:"family"`
	ContactMstStatusPlaceStatus    string `gorm:"default:'null'" json:"contact_mst_status_place_status"`
	Company                        string `gorm:"default:'null'" json:"company"`
	ContactMstStatusEmployeeStatus string `gorm:"default:'null'" json:"contact_mst_status_employee_status"`
	Position                       string `gorm:"default:'null'" json:"position"`
	WorkingTime                    int    `json:"working_time"`
	Income                         int    `json:"income"`
	Outlay                         int    `json:"outlay"`
}

type ContactPhoneDetail struct {
	Id     int64  `gorm:"default:'null'" json:"id"`
	Number string `gorm:"default:'null'" json:"number"`
	Status string `gorm:"default:'null'" json:"status"`
}

type ContactAddressDetail struct {
	Id                         int64  `gorm:"default:'null'" json:"id"`
	CreatedAt                  string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt                  string `gorm:"default:'null'" json:"updated_at"`
	Address                    string `gorm:"default:'null'" json:"address"`
	IdMstAddress               int64  `gorm:"default:'null'" json:"id_mst_address"`
	IdMstCategoryAddress       int64  `gorm:"default:'null'" json:"id_mst_category_address"`
	IdCmsUsers                 int64  `gorm:"default:'null'" json:"id_cms_users"`
	UpdatedBy                  int64  `gorm:"default:'null'" json:"updated_by"`
	Kelurahan                  string `gorm:"default:'null'" json:"kelurahan"`
	Kecamatan                  string `gorm:"default:'null'" json:"kecamatan"`
	Kabupaten                  string `gorm:"default:'null'" json:"kabupaten"`
	Provinsi                   string `gorm:"default:'null'" json:"provinsi"`
	Kodepos                    string `gorm:"default:'null'" json:"kodepos"`
	Rt                         string `gorm:"default:'null'" json:"rt"`
	Rw                         string `gorm:"default:'null'" json:"rw"`
	MstCategoryAddressCategory string `gorm:"default:'null'" json:"mst_category_address_category"`
}

type ContactNote struct {
	Id           int64  `gorm:"default:'null'" json:"id"`
	Note         string `gorm:"default:'null'" json:"note"`
	CmsUsersName string `gorm:"default:'null'" json:"cms_users_name"`
	CreatedAt    string `gorm:"default:'null'" json:"created_at"`
}

type ContactNoteCreate struct {
	Id         int64  `gorm:"default:'null'" json:"id"`
	IdContact  int64  `gorm:"default:'null'" json:"id_target"`
	Note       string `gorm:"default:'null'" json:"note"`
	IdCmsUsers int64  `gorm:"default:'null'" json:"id_cms_users"`
	CreatedAt  string `gorm:"default:'null'" json:"created_at"`
}

type ContactCollabList struct {
	Id           int64  `gorm:"default:'null'" json:"id"`
	CmsUsersName string `gorm:"default:'null'" json:"cms_users_name"`
	CmsUsersNpm  string `gorm:"default:'null'" json:"cms_users_npm"`
}

type ContactCollabCreate struct {
	Id         int64 `json:"id"`
	IdContact  int64 `json:"id_contact"`
	IdCmsUsers int64 `json:"id_cms_users"`
}
