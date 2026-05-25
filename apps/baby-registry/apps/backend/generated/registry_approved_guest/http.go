package registry_approved_guest

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_guest_access_level"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_guest_status"
)

type HTTPRecord struct {
	Id                 *string                        `json:"id,omitempty"`
	AccessLevel        *enum_guest_access_level.Value `json:"accessLevel,omitempty"`
	Created            *actor_trace.HTTPRecord        `json:"created,omitempty"`
	EmailEnc           *string                        `json:"emailEnc,omitempty"`
	EmailHash          *string                        `json:"emailHash,omitempty"`
	Name               *string                        `json:"name,omitempty"`
	OwnerId            *string                        `json:"ownerId,omitempty"`
	RegistryId         *string                        `json:"registryId,omitempty"`
	Status             *enum_guest_status.Value       `json:"status,omitempty"`
	Updated            *actor_trace.HTTPRecord        `json:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.HTTPRecord        `json:"updatedByOwnerUser,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id
		m.Id = *elemid0
	}
	if r.AccessLevel != nil {
		elemaccessLevel0 := r.AccessLevel
		m.AccessLevel = *elemaccessLevel0
	}
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.EmailEnc != nil {
		elememailEnc0 := r.EmailEnc
		m.EmailEnc = *elememailEnc0
	}
	if r.EmailHash != nil {
		elememailHash0 := r.EmailHash
		m.EmailHash = *elememailHash0
	}
	if r.Name != nil {
		elemname0 := r.Name
		m.Name = *elemname0
	}
	if r.OwnerId != nil {
		elemownerId0 := r.OwnerId
		m.OwnerId = *elemownerId0
	}
	if r.RegistryId != nil {
		elemregistryId0 := r.RegistryId
		m.RegistryId = *elemregistryId0
	}
	if r.Status != nil {
		elemstatus0 := r.Status
		m.Status = *elemstatus0
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
	if r.AccessLevel != nil {
		p.AccessLevel = true
	}
	if r.Created != nil {
		p.Created = true
		p.CreatedFields = actor_trace.NewProjection(true)
	}
	if r.EmailEnc != nil {
		p.EmailEnc = true
	}
	if r.EmailHash != nil {
		p.EmailHash = true
	}
	if r.Name != nil {
		p.Name = true
	}
	if r.OwnerId != nil {
		p.OwnerId = true
	}
	if r.RegistryId != nil {
		p.RegistryId = true
	}
	if r.Status != nil {
		p.Status = true
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
type HTTPSelectByRegistryEmailUniqueQuery struct {
	RegistryId string `json:"registryId"`
	EmailHash  string `json:"emailHash"`
}

type HTTPWhereClause struct {
	// id (Ref<RegistryApprovedGuest>) search options
	IdEq     *string   `json:"idEq,omitempty"`
	IdIn     *[]string `json:"idIn,omitempty"`
	IdNin    *[]string `json:"idNin,omitempty"`
	IdExists *bool     `json:"idExists,omitempty"`
	// accessLevel (GuestAccessLevel) search options
	AccessLevelEq     *enum_guest_access_level.Value   `json:"accessLevelEq,omitempty"`
	AccessLevelNe     *enum_guest_access_level.Value   `json:"accessLevelNe,omitempty"`
	AccessLevelGt     *enum_guest_access_level.Value   `json:"accessLevelGt,omitempty"`
	AccessLevelGte    *enum_guest_access_level.Value   `json:"accessLevelGte,omitempty"`
	AccessLevelLt     *enum_guest_access_level.Value   `json:"accessLevelLt,omitempty"`
	AccessLevelLte    *enum_guest_access_level.Value   `json:"accessLevelLte,omitempty"`
	AccessLevelIn     *[]enum_guest_access_level.Value `json:"accessLevelIn,omitempty"`
	AccessLevelNin    *[]enum_guest_access_level.Value `json:"accessLevelNin,omitempty"`
	AccessLevelExists *bool                            `json:"accessLevelExists,omitempty"`
	// created (ActorTrace) search options
	Created *actor_trace.HTTPWhereClause `json:"created,omitempty"`
	// emailEnc (string) search options
	EmailEncEq     *string   `json:"emailEncEq,omitempty"`
	EmailEncNe     *string   `json:"emailEncNe,omitempty"`
	EmailEncGt     *string   `json:"emailEncGt,omitempty"`
	EmailEncGte    *string   `json:"emailEncGte,omitempty"`
	EmailEncLt     *string   `json:"emailEncLt,omitempty"`
	EmailEncLte    *string   `json:"emailEncLte,omitempty"`
	EmailEncIn     *[]string `json:"emailEncIn,omitempty"`
	EmailEncNin    *[]string `json:"emailEncNin,omitempty"`
	EmailEncExists *bool     `json:"emailEncExists,omitempty"`
	EmailEncLike   *string   `json:"emailEncLike,omitempty"`
	EmailEncNlike  *string   `json:"emailEncNlike,omitempty"`
	// emailHash (string) search options
	EmailHashEq     *string   `json:"emailHashEq,omitempty"`
	EmailHashNe     *string   `json:"emailHashNe,omitempty"`
	EmailHashGt     *string   `json:"emailHashGt,omitempty"`
	EmailHashGte    *string   `json:"emailHashGte,omitempty"`
	EmailHashLt     *string   `json:"emailHashLt,omitempty"`
	EmailHashLte    *string   `json:"emailHashLte,omitempty"`
	EmailHashIn     *[]string `json:"emailHashIn,omitempty"`
	EmailHashNin    *[]string `json:"emailHashNin,omitempty"`
	EmailHashExists *bool     `json:"emailHashExists,omitempty"`
	EmailHashLike   *string   `json:"emailHashLike,omitempty"`
	EmailHashNlike  *string   `json:"emailHashNlike,omitempty"`
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
	// ownerId (Ref<OwnerUser>) search options
	OwnerIdEq     *string   `json:"ownerIdEq,omitempty"`
	OwnerIdIn     *[]string `json:"ownerIdIn,omitempty"`
	OwnerIdNin    *[]string `json:"ownerIdNin,omitempty"`
	OwnerIdExists *bool     `json:"ownerIdExists,omitempty"`
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *string   `json:"registryIdEq,omitempty"`
	RegistryIdIn     *[]string `json:"registryIdIn,omitempty"`
	RegistryIdNin    *[]string `json:"registryIdNin,omitempty"`
	RegistryIdExists *bool     `json:"registryIdExists,omitempty"`
	// status (GuestStatus) search options
	StatusEq     *enum_guest_status.Value   `json:"statusEq,omitempty"`
	StatusNe     *enum_guest_status.Value   `json:"statusNe,omitempty"`
	StatusGt     *enum_guest_status.Value   `json:"statusGt,omitempty"`
	StatusGte    *enum_guest_status.Value   `json:"statusGte,omitempty"`
	StatusLt     *enum_guest_status.Value   `json:"statusLt,omitempty"`
	StatusLte    *enum_guest_status.Value   `json:"statusLte,omitempty"`
	StatusIn     *[]enum_guest_status.Value `json:"statusIn,omitempty"`
	StatusNin    *[]enum_guest_status.Value `json:"statusNin,omitempty"`
	StatusExists *bool                      `json:"statusExists,omitempty"`
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
func (o HTTPSelectByRegistryEmailUniqueQuery) ToSelectByRegistryEmailUniqueQuery() (SelectByRegistryEmailUniqueQuery, error) {
	to := SelectByRegistryEmailUniqueQuery{}
	elemregistryId0 := o.RegistryId
	to.RegistryId = elemregistryId0
	elememailHash0 := o.EmailHash
	to.EmailHash = elememailHash0
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
	if o.AccessLevelEq != nil {
		elemaccessLevelEq0 := o.AccessLevelEq
		to.AccessLevelEq = elemaccessLevelEq0
	}
	if o.AccessLevelNe != nil {
		elemaccessLevelNe0 := o.AccessLevelNe
		to.AccessLevelNe = elemaccessLevelNe0
	}
	if o.AccessLevelGt != nil {
		elemaccessLevelGt0 := o.AccessLevelGt
		to.AccessLevelGt = elemaccessLevelGt0
	}
	if o.AccessLevelGte != nil {
		elemaccessLevelGte0 := o.AccessLevelGte
		to.AccessLevelGte = elemaccessLevelGte0
	}
	if o.AccessLevelLt != nil {
		elemaccessLevelLt0 := o.AccessLevelLt
		to.AccessLevelLt = elemaccessLevelLt0
	}
	if o.AccessLevelLte != nil {
		elemaccessLevelLte0 := o.AccessLevelLte
		to.AccessLevelLte = elemaccessLevelLte0
	}
	if o.AccessLevelIn != nil {
		elemaccessLevelIn0 := make([]enum_guest_access_level.Value, 0)
		for _, oaccessLevelIn0 := range *o.AccessLevelIn {
			elemaccessLevelIn1 := oaccessLevelIn0
			elemaccessLevelIn0 = append(elemaccessLevelIn0, elemaccessLevelIn1)
		}
		to.AccessLevelIn = &elemaccessLevelIn0
	}
	if o.AccessLevelNin != nil {
		elemaccessLevelNin0 := make([]enum_guest_access_level.Value, 0)
		for _, oaccessLevelNin0 := range *o.AccessLevelNin {
			elemaccessLevelNin1 := oaccessLevelNin0
			elemaccessLevelNin0 = append(elemaccessLevelNin0, elemaccessLevelNin1)
		}
		to.AccessLevelNin = &elemaccessLevelNin0
	}
	if o.AccessLevelExists != nil {
		elemaccessLevelExists0 := o.AccessLevelExists
		to.AccessLevelExists = elemaccessLevelExists0
	}
	if o.Created != nil {
		elemcreated0, err := o.Created.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.Created = &elemcreated0
	}
	if o.EmailEncEq != nil {
		elememailEncEq0 := o.EmailEncEq
		to.EmailEncEq = elememailEncEq0
	}
	if o.EmailEncNe != nil {
		elememailEncNe0 := o.EmailEncNe
		to.EmailEncNe = elememailEncNe0
	}
	if o.EmailEncGt != nil {
		elememailEncGt0 := o.EmailEncGt
		to.EmailEncGt = elememailEncGt0
	}
	if o.EmailEncGte != nil {
		elememailEncGte0 := o.EmailEncGte
		to.EmailEncGte = elememailEncGte0
	}
	if o.EmailEncLt != nil {
		elememailEncLt0 := o.EmailEncLt
		to.EmailEncLt = elememailEncLt0
	}
	if o.EmailEncLte != nil {
		elememailEncLte0 := o.EmailEncLte
		to.EmailEncLte = elememailEncLte0
	}
	if o.EmailEncIn != nil {
		elememailEncIn0 := make([]string, 0)
		for _, oemailEncIn0 := range *o.EmailEncIn {
			elememailEncIn1 := oemailEncIn0
			elememailEncIn0 = append(elememailEncIn0, elememailEncIn1)
		}
		to.EmailEncIn = &elememailEncIn0
	}
	if o.EmailEncNin != nil {
		elememailEncNin0 := make([]string, 0)
		for _, oemailEncNin0 := range *o.EmailEncNin {
			elememailEncNin1 := oemailEncNin0
			elememailEncNin0 = append(elememailEncNin0, elememailEncNin1)
		}
		to.EmailEncNin = &elememailEncNin0
	}
	if o.EmailEncExists != nil {
		elememailEncExists0 := o.EmailEncExists
		to.EmailEncExists = elememailEncExists0
	}
	if o.EmailEncLike != nil {
		elememailEncLike0 := o.EmailEncLike
		to.EmailEncLike = elememailEncLike0
	}
	if o.EmailEncNlike != nil {
		elememailEncNlike0 := o.EmailEncNlike
		to.EmailEncNlike = elememailEncNlike0
	}
	if o.EmailHashEq != nil {
		elememailHashEq0 := o.EmailHashEq
		to.EmailHashEq = elememailHashEq0
	}
	if o.EmailHashNe != nil {
		elememailHashNe0 := o.EmailHashNe
		to.EmailHashNe = elememailHashNe0
	}
	if o.EmailHashGt != nil {
		elememailHashGt0 := o.EmailHashGt
		to.EmailHashGt = elememailHashGt0
	}
	if o.EmailHashGte != nil {
		elememailHashGte0 := o.EmailHashGte
		to.EmailHashGte = elememailHashGte0
	}
	if o.EmailHashLt != nil {
		elememailHashLt0 := o.EmailHashLt
		to.EmailHashLt = elememailHashLt0
	}
	if o.EmailHashLte != nil {
		elememailHashLte0 := o.EmailHashLte
		to.EmailHashLte = elememailHashLte0
	}
	if o.EmailHashIn != nil {
		elememailHashIn0 := make([]string, 0)
		for _, oemailHashIn0 := range *o.EmailHashIn {
			elememailHashIn1 := oemailHashIn0
			elememailHashIn0 = append(elememailHashIn0, elememailHashIn1)
		}
		to.EmailHashIn = &elememailHashIn0
	}
	if o.EmailHashNin != nil {
		elememailHashNin0 := make([]string, 0)
		for _, oemailHashNin0 := range *o.EmailHashNin {
			elememailHashNin1 := oemailHashNin0
			elememailHashNin0 = append(elememailHashNin0, elememailHashNin1)
		}
		to.EmailHashNin = &elememailHashNin0
	}
	if o.EmailHashExists != nil {
		elememailHashExists0 := o.EmailHashExists
		to.EmailHashExists = elememailHashExists0
	}
	if o.EmailHashLike != nil {
		elememailHashLike0 := o.EmailHashLike
		to.EmailHashLike = elememailHashLike0
	}
	if o.EmailHashNlike != nil {
		elememailHashNlike0 := o.EmailHashNlike
		to.EmailHashNlike = elememailHashNlike0
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
	if o.OwnerIdEq != nil {
		elemownerIdEq0 := o.OwnerIdEq
		to.OwnerIdEq = elemownerIdEq0
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
	if o.RegistryIdEq != nil {
		elemregistryIdEq0 := o.RegistryIdEq
		to.RegistryIdEq = elemregistryIdEq0
	}
	if o.RegistryIdIn != nil {
		elemregistryIdIn0 := make([]string, 0)
		for _, oregistryIdIn0 := range *o.RegistryIdIn {
			elemregistryIdIn1 := oregistryIdIn0
			elemregistryIdIn0 = append(elemregistryIdIn0, elemregistryIdIn1)
		}
		to.RegistryIdIn = &elemregistryIdIn0
	}
	if o.RegistryIdNin != nil {
		elemregistryIdNin0 := make([]string, 0)
		for _, oregistryIdNin0 := range *o.RegistryIdNin {
			elemregistryIdNin1 := oregistryIdNin0
			elemregistryIdNin0 = append(elemregistryIdNin0, elemregistryIdNin1)
		}
		to.RegistryIdNin = &elemregistryIdNin0
	}
	if o.RegistryIdExists != nil {
		elemregistryIdExists0 := o.RegistryIdExists
		to.RegistryIdExists = elemregistryIdExists0
	}
	if o.StatusEq != nil {
		elemstatusEq0 := o.StatusEq
		to.StatusEq = elemstatusEq0
	}
	if o.StatusNe != nil {
		elemstatusNe0 := o.StatusNe
		to.StatusNe = elemstatusNe0
	}
	if o.StatusGt != nil {
		elemstatusGt0 := o.StatusGt
		to.StatusGt = elemstatusGt0
	}
	if o.StatusGte != nil {
		elemstatusGte0 := o.StatusGte
		to.StatusGte = elemstatusGte0
	}
	if o.StatusLt != nil {
		elemstatusLt0 := o.StatusLt
		to.StatusLt = elemstatusLt0
	}
	if o.StatusLte != nil {
		elemstatusLte0 := o.StatusLte
		to.StatusLte = elemstatusLte0
	}
	if o.StatusIn != nil {
		elemstatusIn0 := make([]enum_guest_status.Value, 0)
		for _, ostatusIn0 := range *o.StatusIn {
			elemstatusIn1 := ostatusIn0
			elemstatusIn0 = append(elemstatusIn0, elemstatusIn1)
		}
		to.StatusIn = &elemstatusIn0
	}
	if o.StatusNin != nil {
		elemstatusNin0 := make([]enum_guest_status.Value, 0)
		for _, ostatusNin0 := range *o.StatusNin {
			elemstatusNin1 := ostatusNin0
			elemstatusNin0 = append(elemstatusNin0, elemstatusNin1)
		}
		to.StatusNin = &elemstatusNin0
	}
	if o.StatusExists != nil {
		elemstatusExists0 := o.StatusExists
		to.StatusExists = elemstatusExists0
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
	CreatedAt  *int8 `json:"createdAt,omitempty"`
	EmailHash  *int8 `json:"emailHash,omitempty"`
	OwnerId    *int8 `json:"ownerId,omitempty"`
	RegistryId *int8 `json:"registryId,omitempty"`
	UpdatedAt  *int8 `json:"updatedAt,omitempty"`
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	if s.CreatedAt != nil {
		to.CreatedAt = *s.CreatedAt
	}
	if s.EmailHash != nil {
		to.EmailHash = *s.EmailHash
	}
	if s.OwnerId != nil {
		to.OwnerId = *s.OwnerId
	}
	if s.RegistryId != nil {
		to.RegistryId = *s.RegistryId
	}
	if s.UpdatedAt != nil {
		to.UpdatedAt = *s.UpdatedAt
	}
	return to
}
