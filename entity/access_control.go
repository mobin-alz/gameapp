package entity

type AccessControl struct {
	ID           uint
	ActorID      uint
	ActorType    ActorType
	ResourceID   uint
	PermissionID uint
}

type ActorType string

const (
	RoleActorType = "role"
	UserActorType = "user"
)
