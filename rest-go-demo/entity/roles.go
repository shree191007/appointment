
package entity
type Role struct {
	roleId   int    `gorm:"primary_key" json:"roleid"`
	roleName string `json:"name"`

}

type Privilege struct {
	privelegeId int `gorm:"primary_key" json:"privilegeid"`
	privilege string `json:"name"`
}

type RolePrivilege struct {
	roleId int `gorm: "ForiegnKey:roleId" json:"roleId"`
	privilegeId int `gorm: "ForiegnKey:privilegeId" json:"privilegeId"`
}
