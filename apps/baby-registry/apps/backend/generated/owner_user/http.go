package owner_user

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_role"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
)

type HTTPRecord struct {
	Id                 *string                  `json:"id,omitempty"`
	ActorRoles         *[]actor_role.HTTPRecord `json:"actorRoles,omitempty"`
	Created            *actor_trace.HTTPRecord  `json:"created,omitempty"`
	Email              *string                  `json:"email,omitempty"`
	Name               *string                  `json:"name,omitempty"`
	Updated            *actor_trace.HTTPRecord  `json:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.HTTPRecord  `json:"updatedByOwnerUser,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id
		m.Id = *elemid0
	}
	if r.ActorRoles != nil {
		elemactorRoles0 := make([]actor_role.Model, 0)
		for _, ractorRoles0 := range *r.ActorRoles {
			elemactorRoles1, err := ractorRoles0.ToModel()
			if err != nil {
				return m, err
			}
			elemactorRoles0 = append(elemactorRoles0, elemactorRoles1)
		}
		m.ActorRoles = elemactorRoles0
	}
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.Email != nil {
		elememail0 := r.Email
		m.Email = *elememail0
	}
	if r.Name != nil {
		elemname0 := r.Name
		m.Name = *elemname0
	}
	if r.Updated != nil {
		elemupdated0, err := r.Updated.ToModel()
		if err != nil {
			return m, err
		}
		m.Updated = elemupdated0
	}
	if r.UpdatedByOwnerUser != nil {
		elemupdatedByOwnerUser0, err := r.UpdatedByOwnerUser.ToModel()
		if err != nil {
			return m, err
		}
		m.UpdatedByOwnerUser = elemupdatedByOwnerUser0
	}
	return m, nil
}

func (r *HTTPRecord) ToProjection() (Projection, error) {
	p := Projection{}
	if r.Id != nil {
		p.Id = true
	}
	if r.ActorRoles != nil {
		p.ActorRoles = true
		p.ActorRolesFields = actor_role.NewProjection(true)
	}
	if r.Created != nil {
		p.Created = true
		p.CreatedFields = actor_trace.NewProjection(true)
	}
	if r.Email != nil {
		p.Email = true
	}
	if r.Name != nil {
		p.Name = true
	}
	if r.Updated != nil {
		p.Updated = true
		p.UpdatedFields = actor_trace.NewProjection(true)
	}
	if r.UpdatedByOwnerUser != nil {
		p.UpdatedByOwnerUser = true
		p.UpdatedByOwnerUserFields = actor_trace.NewProjection(true)
	}
	return p, nil
}

type HTTPSelectByIdQuery struct {
	Id string `json:"id"`
}
type HTTPSelectByEmailUniqueQuery struct {
	Email string `json:"email"`
}

type HTTPWhereClause struct {
	// id (Ref<OwnerUser>) search options
	IdEq     *string   `json:"idEq,omitempty"`
	IdIn     *[]string `json:"idIn,omitempty"`
	IdNin    *[]string `json:"idNin,omitempty"`
	IdExists *bool     `json:"idExists,omitempty"`
	// actorRoles (List<ActorRole>) search options
	ActorRoles      *actor_role.HTTPWhereClause `json:"actorRoles,omitempty"`
	ActorRolesEmpty *bool                       `json:"actorRolesEmpty,omitempty"`
	// created (ActorTrace) search options
	Created *actor_trace.HTTPWhereClause `json:"created,omitempty"`
	// email (string) search options
	EmailEq     *string   `json:"emailEq,omitempty"`
	EmailNe     *string   `json:"emailNe,omitempty"`
	EmailGt     *string   `json:"emailGt,omitempty"`
	EmailGte    *string   `json:"emailGte,omitempty"`
	EmailLt     *string   `json:"emailLt,omitempty"`
	EmailLte    *string   `json:"emailLte,omitempty"`
	EmailIn     *[]string `json:"emailIn,omitempty"`
	EmailNin    *[]string `json:"emailNin,omitempty"`
	EmailExists *bool     `json:"emailExists,omitempty"`
	EmailLike   *string   `json:"emailLike,omitempty"`
	EmailNlike  *string   `json:"emailNlike,omitempty"`
	// name (string) search options
	NameEq     *string   `json:"nameEq,omitempty"`
	NameNe     *string   `json:"nameNe,omitempty"`
	NameGt     *string   `json:"nameGt,omitempty"`
	NameGte    *string   `json:"nameGte,omitempty"`
	NameLt     *string   `json:"nameLt,omitempty"`
	NameLte    *string   `json:"nameLte,omitempty"`
	NameIn     *[]string `json:"nameIn,omitempty"`
	NameNin    *[]string `json:"nameNin,omitempty"`
	NameExists *bool     `json:"nameExists,omitempty"`
	NameLike   *string   `json:"nameLike,omitempty"`
	NameNlike  *string   `json:"nameNlike,omitempty"`
	// updated (ActorTrace) search options
	Updated *actor_trace.HTTPWhereClause `json:"updated,omitempty"`
	// updatedByOwnerUser (ActorTrace) search options
	UpdatedByOwnerUser *actor_trace.HTTPWhereClause `json:"updatedByOwnerUser,omitempty"`
}

func (o HTTPSelectByIdQuery) ToSelectByIdQuery() (SelectByIdQuery, error) {
	to := SelectByIdQuery{}
	elemid0 := o.Id
	to.Id = elemid0
	return to, nil
}
func (o HTTPSelectByEmailUniqueQuery) ToSelectByEmailUniqueQuery() (SelectByEmailUniqueQuery, error) {
	to := SelectByEmailUniqueQuery{}
	elememail0 := o.Email
	to.Email = elememail0
	return to, nil
}

func (o HTTPWhereClause) ToWhereClause() (WhereClause, error) {
	to := WhereClause{}
	if o.IdEq != nil {
		elemidEq0 := o.IdEq
		to.IdEq = elemidEq0
	}
	if o.IdIn != nil {
		elemidIn0 := make([]string, 0)
		for _, oidIn0 := range *o.IdIn {
			elemidIn1 := oidIn0
			elemidIn0 = append(elemidIn0, elemidIn1)
		}
		to.IdIn = &elemidIn0
	}
	if o.IdNin != nil {
		elemidNin0 := make([]string, 0)
		for _, oidNin0 := range *o.IdNin {
			elemidNin1 := oidNin0
			elemidNin0 = append(elemidNin0, elemidNin1)
		}
		to.IdNin = &elemidNin0
	}
	if o.IdExists != nil {
		elemidExists0 := o.IdExists
		to.IdExists = elemidExists0
	}
	if o.ActorRoles != nil {
		elemactorRoles0, err := o.ActorRoles.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.ActorRoles = &elemactorRoles0
	}
	if o.ActorRolesEmpty != nil {
		elemactorRolesEmpty0 := o.ActorRolesEmpty
		to.ActorRolesEmpty = elemactorRolesEmpty0
	}
	if o.Created != nil {
		elemcreated0, err := o.Created.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.Created = &elemcreated0
	}
	if o.EmailEq != nil {
		elememailEq0 := o.EmailEq
		to.EmailEq = elememailEq0
	}
	if o.EmailNe != nil {
		elememailNe0 := o.EmailNe
		to.EmailNe = elememailNe0
	}
	if o.EmailGt != nil {
		elememailGt0 := o.EmailGt
		to.EmailGt = elememailGt0
	}
	if o.EmailGte != nil {
		elememailGte0 := o.EmailGte
		to.EmailGte = elememailGte0
	}
	if o.EmailLt != nil {
		elememailLt0 := o.EmailLt
		to.EmailLt = elememailLt0
	}
	if o.EmailLte != nil {
		elememailLte0 := o.EmailLte
		to.EmailLte = elememailLte0
	}
	if o.EmailIn != nil {
		elememailIn0 := make([]string, 0)
		for _, oemailIn0 := range *o.EmailIn {
			elememailIn1 := oemailIn0
			elememailIn0 = append(elememailIn0, elememailIn1)
		}
		to.EmailIn = &elememailIn0
	}
	if o.EmailNin != nil {
		elememailNin0 := make([]string, 0)
		for _, oemailNin0 := range *o.EmailNin {
			elememailNin1 := oemailNin0
			elememailNin0 = append(elememailNin0, elememailNin1)
		}
		to.EmailNin = &elememailNin0
	}
	if o.EmailExists != nil {
		elememailExists0 := o.EmailExists
		to.EmailExists = elememailExists0
	}
	if o.EmailLike != nil {
		elememailLike0 := o.EmailLike
		to.EmailLike = elememailLike0
	}
	if o.EmailNlike != nil {
		elememailNlike0 := o.EmailNlike
		to.EmailNlike = elememailNlike0
	}
	if o.NameEq != nil {
		elemnameEq0 := o.NameEq
		to.NameEq = elemnameEq0
	}
	if o.NameNe != nil {
		elemnameNe0 := o.NameNe
		to.NameNe = elemnameNe0
	}
	if o.NameGt != nil {
		elemnameGt0 := o.NameGt
		to.NameGt = elemnameGt0
	}
	if o.NameGte != nil {
		elemnameGte0 := o.NameGte
		to.NameGte = elemnameGte0
	}
	if o.NameLt != nil {
		elemnameLt0 := o.NameLt
		to.NameLt = elemnameLt0
	}
	if o.NameLte != nil {
		elemnameLte0 := o.NameLte
		to.NameLte = elemnameLte0
	}
	if o.NameIn != nil {
		elemnameIn0 := make([]string, 0)
		for _, onameIn0 := range *o.NameIn {
			elemnameIn1 := onameIn0
			elemnameIn0 = append(elemnameIn0, elemnameIn1)
		}
		to.NameIn = &elemnameIn0
	}
	if o.NameNin != nil {
		elemnameNin0 := make([]string, 0)
		for _, onameNin0 := range *o.NameNin {
			elemnameNin1 := onameNin0
			elemnameNin0 = append(elemnameNin0, elemnameNin1)
		}
		to.NameNin = &elemnameNin0
	}
	if o.NameExists != nil {
		elemnameExists0 := o.NameExists
		to.NameExists = elemnameExists0
	}
	if o.NameLike != nil {
		elemnameLike0 := o.NameLike
		to.NameLike = elemnameLike0
	}
	if o.NameNlike != nil {
		elemnameNlike0 := o.NameNlike
		to.NameNlike = elemnameNlike0
	}
	if o.Updated != nil {
		elemupdated0, err := o.Updated.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.Updated = &elemupdated0
	}
	if o.UpdatedByOwnerUser != nil {
		elemupdatedByOwnerUser0, err := o.UpdatedByOwnerUser.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.UpdatedByOwnerUser = &elemupdatedByOwnerUser0
	}
	return to, nil
}

type HTTPSortParams struct {
	CreatedAt *int8 `json:"createdAt,omitempty"`
	Email     *int8 `json:"email,omitempty"`
	UpdatedAt *int8 `json:"updatedAt,omitempty"`
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	if s.CreatedAt != nil {
		to.CreatedAt = *s.CreatedAt
	}
	if s.Email != nil {
		to.Email = *s.Email
	}
	if s.UpdatedAt != nil {
		to.UpdatedAt = *s.UpdatedAt
	}
	return to
}
