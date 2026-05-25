package actor_role

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_role"
)

type Model struct {
	OwnerId string
	Role    enum_role.Value
}

func (m *Model) ToMongoRecord(projection Projection) (MongoRecord, error) {
	r := MongoRecord{}
	if projection.OwnerId {
		elemownerId0 := m.OwnerId
		r.OwnerId = &elemownerId0
	}
	if projection.Role {
		elemrole0 := m.Role
		r.Role = &elemrole0
	}
	return r, nil
}

func (m *Model) ToHTTPRecord(projection Projection) (HTTPRecord, error) {
	r := HTTPRecord{}
	if projection.OwnerId {
		elemownerId0 := m.OwnerId
		r.OwnerId = &elemownerId0
	}
	if projection.Role {
		elemrole0 := m.Role
		r.Role = &elemrole0
	}
	return r, nil
}

type WhereClause struct {
	// ownerId (string) search options
	OwnerIdEq     *string
	OwnerIdNe     *string
	OwnerIdGt     *string
	OwnerIdGte    *string
	OwnerIdLt     *string
	OwnerIdLte    *string
	OwnerIdIn     *[]string
	OwnerIdNin    *[]string
	OwnerIdExists *bool
	OwnerIdLike   *string
	OwnerIdNlike  *string
	// role (Role) search options
	RoleEq     *enum_role.Value
	RoleNe     *enum_role.Value
	RoleGt     *enum_role.Value
	RoleGte    *enum_role.Value
	RoleLt     *enum_role.Value
	RoleLte    *enum_role.Value
	RoleIn     *[]enum_role.Value
	RoleNin    *[]enum_role.Value
	RoleExists *bool
}

func (o WhereClause) ToMongoWhereClause() (MongoWhereClause, error) {
	to := MongoWhereClause{}
	if o.OwnerIdEq != nil {
		elemownerIdEq0 := o.OwnerIdEq
		to.OwnerIdEq = elemownerIdEq0
	}
	if o.OwnerIdNe != nil {
		elemownerIdNe0 := o.OwnerIdNe
		to.OwnerIdNe = elemownerIdNe0
	}
	if o.OwnerIdGt != nil {
		elemownerIdGt0 := o.OwnerIdGt
		to.OwnerIdGt = elemownerIdGt0
	}
	if o.OwnerIdGte != nil {
		elemownerIdGte0 := o.OwnerIdGte
		to.OwnerIdGte = elemownerIdGte0
	}
	if o.OwnerIdLt != nil {
		elemownerIdLt0 := o.OwnerIdLt
		to.OwnerIdLt = elemownerIdLt0
	}
	if o.OwnerIdLte != nil {
		elemownerIdLte0 := o.OwnerIdLte
		to.OwnerIdLte = elemownerIdLte0
	}
	if o.OwnerIdIn != nil {
		elemownerIdIn0 := make([]string, 0)
		for _, oownerIdIn0 := range *o.OwnerIdIn {
			elemownerIdIn1 := oownerIdIn0
			elemownerIdIn0 = append(elemownerIdIn0, elemownerIdIn1)
		}
		to.OwnerIdIn = &elemownerIdIn0
	}
	if o.OwnerIdNin != nil {
		elemownerIdNin0 := make([]string, 0)
		for _, oownerIdNin0 := range *o.OwnerIdNin {
			elemownerIdNin1 := oownerIdNin0
			elemownerIdNin0 = append(elemownerIdNin0, elemownerIdNin1)
		}
		to.OwnerIdNin = &elemownerIdNin0
	}
	if o.OwnerIdExists != nil {
		elemownerIdExists0 := o.OwnerIdExists
		to.OwnerIdExists = elemownerIdExists0
	}
	if o.OwnerIdLike != nil {
		elemownerIdLike0 := o.OwnerIdLike
		to.OwnerIdLike = elemownerIdLike0
	}
	if o.OwnerIdNlike != nil {
		elemownerIdNlike0 := o.OwnerIdNlike
		to.OwnerIdNlike = elemownerIdNlike0
	}
	if o.RoleEq != nil {
		elemroleEq0 := o.RoleEq
		to.RoleEq = elemroleEq0
	}
	if o.RoleNe != nil {
		elemroleNe0 := o.RoleNe
		to.RoleNe = elemroleNe0
	}
	if o.RoleGt != nil {
		elemroleGt0 := o.RoleGt
		to.RoleGt = elemroleGt0
	}
	if o.RoleGte != nil {
		elemroleGte0 := o.RoleGte
		to.RoleGte = elemroleGte0
	}
	if o.RoleLt != nil {
		elemroleLt0 := o.RoleLt
		to.RoleLt = elemroleLt0
	}
	if o.RoleLte != nil {
		elemroleLte0 := o.RoleLte
		to.RoleLte = elemroleLte0
	}
	if o.RoleIn != nil {
		elemroleIn0 := make([]enum_role.Value, 0)
		for _, oroleIn0 := range *o.RoleIn {
			elemroleIn1 := oroleIn0
			elemroleIn0 = append(elemroleIn0, elemroleIn1)
		}
		to.RoleIn = &elemroleIn0
	}
	if o.RoleNin != nil {
		elemroleNin0 := make([]enum_role.Value, 0)
		for _, oroleNin0 := range *o.RoleNin {
			elemroleNin1 := oroleNin0
			elemroleNin0 = append(elemroleNin0, elemroleNin1)
		}
		to.RoleNin = &elemroleNin0
	}
	if o.RoleExists != nil {
		elemroleExists0 := o.RoleExists
		to.RoleExists = elemroleExists0
	}
	return to, nil
}

type SortParams struct {
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	return to
}
