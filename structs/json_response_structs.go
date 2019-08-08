package structs

type JsonResponse struct {
	ApiStatus        int         `gorm:"default:'null'" json:"api_status"`
	ApiMessage       string      `gorm:"default:'null'" json:"api_message"`
	ApiAuthorization string      `gorm:"default:'null'" json:"api_authorization"`
	ApiHttp          int         `gorm:"default:'null'" json:"api_http"`
	Data             interface{} `gorm:"default:'null'" json:"data"`
}

type JsonResponseToken struct {
	ApiStatus        int         `gorm:"default:'null'" json:"api_status"`
	ApiMessage       string      `gorm:"default:'null'" json:"api_message"`
	ApiAuthorization string      `gorm:"default:'null'" json:"api_authorization"`
	ApiHttp          int         `gorm:"default:'null'" json:"api_http"`
	Data             interface{} `gorm:"default:'null'" json:"data"`
	Token            string      `json:"token"`
}
