package structs

type HumanName struct {
	
	FirstName string `gorm:"default:'null'" json:"first_name"`
	LastName string `gorm:"default:'null'" json:"last_name"`
}
