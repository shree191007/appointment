package entity
type Patient struct {
	patient_id        int `gorm:"primary_key"   json:"id"`
	name string `json:"name"`
	age int `json:"age"`
	mail string `json:"mail"`
	phone string `json:"phone"`

}
