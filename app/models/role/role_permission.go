package role

type RolePermission struct {
	ID 				uint `gorm:"primaryKey" json:"id"`
	RoleID       	uint
	PermissionID 	uint
	Role 			*Role
	Permission		*Permission
}