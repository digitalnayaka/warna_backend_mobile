package structs

type CmsUsersStruct struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`
	UpdatedAt string `gorm:"default:'null'" json:"updated_at"`
	ApiToken string `gorm:"default:'null'" json:"api_token"`
	Npm string `gorm:"default:'null'" json:"npm"`
	Name string `gorm:"default:'null'" json:"name"`
	Password string `gorm:"default:'null'" json:"password"`
	Email string `gorm:"default:'null'" json:"email"`
	Photo string `gorm:"default:'null'" json:"photo"`
	Status string `gorm:"default:'null'" json:"status"`
	DeviceId string `gorm:"default:'null'" json:"device_id"`
	IdCmsPrivileges int `gorm:"default:'null'" json:"id_cms_privileges"`
	IdMstOutlet int64 `gorm:"default:'null'" json:"id_mst_outlet"`
	CreatedBy int64 `gorm:"default:'null'" json:"created_by"`
	UpdatedBy int64 `gorm:"default:'null'" json:"updated_by"`
	MstOutletOutletName string `gorm:"default:'null'" json:"mst_outlet_outlet_name"`
	IdMstBranch int64 `gorm:"default:'null'" json:"id_mst_branch"`
	CmsPrivilegesName string `gorm:"default:'null'" json:"cms_privileges_name"`
}

type UserAgreement struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`
	Eula string `gorm:"default:'null'" json:"eula"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
}

type UserLogs struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`
	Useragent string `gorm:"default:'null'" json:"useragent"`
	Description string `gorm:"default:'null'" json:"description"`
	Version string `gorm:"default:'null'" json:"version"`
	IdCmsUsers int64 `gorm:"default:'null'" json:"id_cms_users"`
}

type ListLogUser struct {

	Id int64 `gorm:"default:'null'" json:"id"`
	CmsUsersName string `gorm:"default:'null'" json:"cms_users_name"`
	IdModul int `gorm:"default:'null'" json:"id_modul"`
	IdData int64 `gorm:"default:'null'" json:"id_data"`
	Jenis string `gorm:"default:'null'" json:"jenis"`
	CreatedAt string `gorm:"default:'null'" json:"created_at"`
	DataAktivitas string `gorm:"default:'null'" json:"data_aktivitas"`
}

type PerformIndicator struct {

	CountBrosur int `json:"count_brosur"`
	CountTele int `json:"count_tele"`
	CountNewdb int `json:"count_newDB"`
	CountOrder int `json:"count_order"`
	CountBooking int `json:"count_booking"`
	CountTeleblnkemarin int `json:"count_teleBlnKemarin"`
	CountNewdbblnkemarin int `json:"count_newDBBlnKemarin"`
	CountOrderbulankemarin int `json:"count_orderBulanKemarin"`
	CountTeletotal int `json:"count_teleTotal"`
	CountNewdbtotal int `json:"count_newDBTotal"`
	CountOrdertotal int `json:"count_orderTotal"`
	CountBookingtotal int `json:"count_bookingTotal"`
	CountBrosurbulankemarin int `json:"count_brosurBulanKemarin"`
	CountBrosurbulantotal int `json:"count_brosurBulanTotal"`
	CountVisum int `json:"count_visum"`
	CountVisumblnkemarin int `json:"count_visumBlnKemarin"`
	CountVisumtotal int `json:"count_visumTotal"`
	CountBookingbulankemarin int `json:"count_bookingBulanKemarin"`
	
}
type Rekap struct {

	CountBrosur int `json:"count_brosur"`
	CountTele int `json:"count_tele"`
	CountNewdb int `json:"count_newDB"`
	CountOrder int `json:"count_order"`
	CountBooking int `json:"count_booking"`
	CountVisum int `json:"count_visum"`
	CountOrderOutlet int `json:"count_order_outlet"`
	CountBookingOutlet int `json:"count_booking_outlet"`

}



