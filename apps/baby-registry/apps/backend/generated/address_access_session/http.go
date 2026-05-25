package address_access_session

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"time"
)

type HTTPRecord struct {
	Id                   *string                 `json:"id,omitempty"`
	ApprovedGuestId      *string                 `json:"approvedGuestId,omitempty"`
	Created              *actor_trace.HTTPRecord `json:"created,omitempty"`
	EmailHash            *string                 `json:"emailHash,omitempty"`
	ExpiresAt            *time.Time              `json:"expiresAt,omitempty"`
	OwnerId              *string                 `json:"ownerId,omitempty"`
	PolicyVersionAtIssue *int                    `json:"policyVersionAtIssue,omitempty"`
	RegistryId           *string                 `json:"registryId,omitempty"`
	TokenHash            *string                 `json:"tokenHash,omitempty"`
	Updated              *actor_trace.HTTPRecord `json:"updated,omitempty"`
	UpdatedByOwnerUser   *actor_trace.HTTPRecord `json:"updatedByOwnerUser,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id
		m.Id = *elemid0
	}
	if r.ApprovedGuestId != nil {
		elemapprovedGuestId0 := r.ApprovedGuestId
		m.ApprovedGuestId = *elemapprovedGuestId0
	}
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.EmailHash != nil {
		elememailHash0 := r.EmailHash
		m.EmailHash = *elememailHash0
	}
	if r.ExpiresAt != nil {
		elemexpiresAt0 := r.ExpiresAt
		m.ExpiresAt = *elemexpiresAt0
	}
	if r.OwnerId != nil {
		elemownerId0 := r.OwnerId
		m.OwnerId = *elemownerId0
	}
	if r.PolicyVersionAtIssue != nil {
		elempolicyVersionAtIssue0 := r.PolicyVersionAtIssue
		m.PolicyVersionAtIssue = *elempolicyVersionAtIssue0
	}
	if r.RegistryId != nil {
		elemregistryId0 := r.RegistryId
		m.RegistryId = *elemregistryId0
	}
	if r.TokenHash != nil {
		elemtokenHash0 := r.TokenHash
		m.TokenHash = *elemtokenHash0
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
	if r.ApprovedGuestId != nil {
		p.ApprovedGuestId = true
	}
	if r.Created != nil {
		p.Created = true
		p.CreatedFields = actor_trace.NewProjection(true)
	}
	if r.EmailHash != nil {
		p.EmailHash = true
	}
	if r.ExpiresAt != nil {
		p.ExpiresAt = true
	}
	if r.OwnerId != nil {
		p.OwnerId = true
	}
	if r.PolicyVersionAtIssue != nil {
		p.PolicyVersionAtIssue = true
	}
	if r.RegistryId != nil {
		p.RegistryId = true
	}
	if r.TokenHash != nil {
		p.TokenHash = true
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
type HTTPSelectByTokenUniqueQuery struct {
	TokenHash string `json:"tokenHash"`
}

type HTTPWhereClause struct {
	// id (Ref<AddressAccessSession>) search options
	IdEq     *string   `json:"idEq,omitempty"`
	IdIn     *[]string `json:"idIn,omitempty"`
	IdNin    *[]string `json:"idNin,omitempty"`
	IdExists *bool     `json:"idExists,omitempty"`
	// approvedGuestId (Ref<RegistryApprovedGuest>) search options
	ApprovedGuestIdEq     *string   `json:"approvedGuestIdEq,omitempty"`
	ApprovedGuestIdIn     *[]string `json:"approvedGuestIdIn,omitempty"`
	ApprovedGuestIdNin    *[]string `json:"approvedGuestIdNin,omitempty"`
	ApprovedGuestIdExists *bool     `json:"approvedGuestIdExists,omitempty"`
	// created (ActorTrace) search options
	Created *actor_trace.HTTPWhereClause `json:"created,omitempty"`
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
	// expiresAt (timestamp) search options
	ExpiresAtEq     *time.Time   `json:"expiresAtEq,omitempty"`
	ExpiresAtNe     *time.Time   `json:"expiresAtNe,omitempty"`
	ExpiresAtGt     *time.Time   `json:"expiresAtGt,omitempty"`
	ExpiresAtGte    *time.Time   `json:"expiresAtGte,omitempty"`
	ExpiresAtLt     *time.Time   `json:"expiresAtLt,omitempty"`
	ExpiresAtLte    *time.Time   `json:"expiresAtLte,omitempty"`
	ExpiresAtIn     *[]time.Time `json:"expiresAtIn,omitempty"`
	ExpiresAtNin    *[]time.Time `json:"expiresAtNin,omitempty"`
	ExpiresAtExists *bool        `json:"expiresAtExists,omitempty"`
	// ownerId (Ref<OwnerUser>) search options
	OwnerIdEq     *string   `json:"ownerIdEq,omitempty"`
	OwnerIdIn     *[]string `json:"ownerIdIn,omitempty"`
	OwnerIdNin    *[]string `json:"ownerIdNin,omitempty"`
	OwnerIdExists *bool     `json:"ownerIdExists,omitempty"`
	// policyVersionAtIssue (int) search options
	PolicyVersionAtIssueEq     *int   `json:"policyVersionAtIssueEq,omitempty"`
	PolicyVersionAtIssueNe     *int   `json:"policyVersionAtIssueNe,omitempty"`
	PolicyVersionAtIssueGt     *int   `json:"policyVersionAtIssueGt,omitempty"`
	PolicyVersionAtIssueGte    *int   `json:"policyVersionAtIssueGte,omitempty"`
	PolicyVersionAtIssueLt     *int   `json:"policyVersionAtIssueLt,omitempty"`
	PolicyVersionAtIssueLte    *int   `json:"policyVersionAtIssueLte,omitempty"`
	PolicyVersionAtIssueIn     *[]int `json:"policyVersionAtIssueIn,omitempty"`
	PolicyVersionAtIssueNin    *[]int `json:"policyVersionAtIssueNin,omitempty"`
	PolicyVersionAtIssueExists *bool  `json:"policyVersionAtIssueExists,omitempty"`
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *string   `json:"registryIdEq,omitempty"`
	RegistryIdIn     *[]string `json:"registryIdIn,omitempty"`
	RegistryIdNin    *[]string `json:"registryIdNin,omitempty"`
	RegistryIdExists *bool     `json:"registryIdExists,omitempty"`
	// tokenHash (string) search options
	TokenHashEq     *string   `json:"tokenHashEq,omitempty"`
	TokenHashNe     *string   `json:"tokenHashNe,omitempty"`
	TokenHashGt     *string   `json:"tokenHashGt,omitempty"`
	TokenHashGte    *string   `json:"tokenHashGte,omitempty"`
	TokenHashLt     *string   `json:"tokenHashLt,omitempty"`
	TokenHashLte    *string   `json:"tokenHashLte,omitempty"`
	TokenHashIn     *[]string `json:"tokenHashIn,omitempty"`
	TokenHashNin    *[]string `json:"tokenHashNin,omitempty"`
	TokenHashExists *bool     `json:"tokenHashExists,omitempty"`
	TokenHashLike   *string   `json:"tokenHashLike,omitempty"`
	TokenHashNlike  *string   `json:"tokenHashNlike,omitempty"`
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
func (o HTTPSelectByTokenUniqueQuery) ToSelectByTokenUniqueQuery() (SelectByTokenUniqueQuery, error) {
	to := SelectByTokenUniqueQuery{}
	elemtokenHash0 := o.TokenHash
	to.TokenHash = elemtokenHash0
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
	if o.ApprovedGuestIdEq != nil {
		elemapprovedGuestIdEq0 := o.ApprovedGuestIdEq
		to.ApprovedGuestIdEq = elemapprovedGuestIdEq0
	}
	if o.ApprovedGuestIdIn != nil {
		elemapprovedGuestIdIn0 := make([]string, 0)
		for _, oapprovedGuestIdIn0 := range *o.ApprovedGuestIdIn {
			elemapprovedGuestIdIn1 := oapprovedGuestIdIn0
			elemapprovedGuestIdIn0 = append(elemapprovedGuestIdIn0, elemapprovedGuestIdIn1)
		}
		to.ApprovedGuestIdIn = &elemapprovedGuestIdIn0
	}
	if o.ApprovedGuestIdNin != nil {
		elemapprovedGuestIdNin0 := make([]string, 0)
		for _, oapprovedGuestIdNin0 := range *o.ApprovedGuestIdNin {
			elemapprovedGuestIdNin1 := oapprovedGuestIdNin0
			elemapprovedGuestIdNin0 = append(elemapprovedGuestIdNin0, elemapprovedGuestIdNin1)
		}
		to.ApprovedGuestIdNin = &elemapprovedGuestIdNin0
	}
	if o.ApprovedGuestIdExists != nil {
		elemapprovedGuestIdExists0 := o.ApprovedGuestIdExists
		to.ApprovedGuestIdExists = elemapprovedGuestIdExists0
	}
	if o.Created != nil {
		elemcreated0, err := o.Created.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.Created = &elemcreated0
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
	if o.ExpiresAtEq != nil {
		elemexpiresAtEq0 := o.ExpiresAtEq
		to.ExpiresAtEq = elemexpiresAtEq0
	}
	if o.ExpiresAtNe != nil {
		elemexpiresAtNe0 := o.ExpiresAtNe
		to.ExpiresAtNe = elemexpiresAtNe0
	}
	if o.ExpiresAtGt != nil {
		elemexpiresAtGt0 := o.ExpiresAtGt
		to.ExpiresAtGt = elemexpiresAtGt0
	}
	if o.ExpiresAtGte != nil {
		elemexpiresAtGte0 := o.ExpiresAtGte
		to.ExpiresAtGte = elemexpiresAtGte0
	}
	if o.ExpiresAtLt != nil {
		elemexpiresAtLt0 := o.ExpiresAtLt
		to.ExpiresAtLt = elemexpiresAtLt0
	}
	if o.ExpiresAtLte != nil {
		elemexpiresAtLte0 := o.ExpiresAtLte
		to.ExpiresAtLte = elemexpiresAtLte0
	}
	if o.ExpiresAtIn != nil {
		elemexpiresAtIn0 := make([]time.Time, 0)
		for _, oexpiresAtIn0 := range *o.ExpiresAtIn {
			elemexpiresAtIn1 := oexpiresAtIn0
			elemexpiresAtIn0 = append(elemexpiresAtIn0, elemexpiresAtIn1)
		}
		to.ExpiresAtIn = &elemexpiresAtIn0
	}
	if o.ExpiresAtNin != nil {
		elemexpiresAtNin0 := make([]time.Time, 0)
		for _, oexpiresAtNin0 := range *o.ExpiresAtNin {
			elemexpiresAtNin1 := oexpiresAtNin0
			elemexpiresAtNin0 = append(elemexpiresAtNin0, elemexpiresAtNin1)
		}
		to.ExpiresAtNin = &elemexpiresAtNin0
	}
	if o.ExpiresAtExists != nil {
		elemexpiresAtExists0 := o.ExpiresAtExists
		to.ExpiresAtExists = elemexpiresAtExists0
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
	if o.PolicyVersionAtIssueEq != nil {
		elempolicyVersionAtIssueEq0 := o.PolicyVersionAtIssueEq
		to.PolicyVersionAtIssueEq = elempolicyVersionAtIssueEq0
	}
	if o.PolicyVersionAtIssueNe != nil {
		elempolicyVersionAtIssueNe0 := o.PolicyVersionAtIssueNe
		to.PolicyVersionAtIssueNe = elempolicyVersionAtIssueNe0
	}
	if o.PolicyVersionAtIssueGt != nil {
		elempolicyVersionAtIssueGt0 := o.PolicyVersionAtIssueGt
		to.PolicyVersionAtIssueGt = elempolicyVersionAtIssueGt0
	}
	if o.PolicyVersionAtIssueGte != nil {
		elempolicyVersionAtIssueGte0 := o.PolicyVersionAtIssueGte
		to.PolicyVersionAtIssueGte = elempolicyVersionAtIssueGte0
	}
	if o.PolicyVersionAtIssueLt != nil {
		elempolicyVersionAtIssueLt0 := o.PolicyVersionAtIssueLt
		to.PolicyVersionAtIssueLt = elempolicyVersionAtIssueLt0
	}
	if o.PolicyVersionAtIssueLte != nil {
		elempolicyVersionAtIssueLte0 := o.PolicyVersionAtIssueLte
		to.PolicyVersionAtIssueLte = elempolicyVersionAtIssueLte0
	}
	if o.PolicyVersionAtIssueIn != nil {
		elempolicyVersionAtIssueIn0 := make([]int, 0)
		for _, opolicyVersionAtIssueIn0 := range *o.PolicyVersionAtIssueIn {
			elempolicyVersionAtIssueIn1 := opolicyVersionAtIssueIn0
			elempolicyVersionAtIssueIn0 = append(elempolicyVersionAtIssueIn0, elempolicyVersionAtIssueIn1)
		}
		to.PolicyVersionAtIssueIn = &elempolicyVersionAtIssueIn0
	}
	if o.PolicyVersionAtIssueNin != nil {
		elempolicyVersionAtIssueNin0 := make([]int, 0)
		for _, opolicyVersionAtIssueNin0 := range *o.PolicyVersionAtIssueNin {
			elempolicyVersionAtIssueNin1 := opolicyVersionAtIssueNin0
			elempolicyVersionAtIssueNin0 = append(elempolicyVersionAtIssueNin0, elempolicyVersionAtIssueNin1)
		}
		to.PolicyVersionAtIssueNin = &elempolicyVersionAtIssueNin0
	}
	if o.PolicyVersionAtIssueExists != nil {
		elempolicyVersionAtIssueExists0 := o.PolicyVersionAtIssueExists
		to.PolicyVersionAtIssueExists = elempolicyVersionAtIssueExists0
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
	if o.TokenHashEq != nil {
		elemtokenHashEq0 := o.TokenHashEq
		to.TokenHashEq = elemtokenHashEq0
	}
	if o.TokenHashNe != nil {
		elemtokenHashNe0 := o.TokenHashNe
		to.TokenHashNe = elemtokenHashNe0
	}
	if o.TokenHashGt != nil {
		elemtokenHashGt0 := o.TokenHashGt
		to.TokenHashGt = elemtokenHashGt0
	}
	if o.TokenHashGte != nil {
		elemtokenHashGte0 := o.TokenHashGte
		to.TokenHashGte = elemtokenHashGte0
	}
	if o.TokenHashLt != nil {
		elemtokenHashLt0 := o.TokenHashLt
		to.TokenHashLt = elemtokenHashLt0
	}
	if o.TokenHashLte != nil {
		elemtokenHashLte0 := o.TokenHashLte
		to.TokenHashLte = elemtokenHashLte0
	}
	if o.TokenHashIn != nil {
		elemtokenHashIn0 := make([]string, 0)
		for _, otokenHashIn0 := range *o.TokenHashIn {
			elemtokenHashIn1 := otokenHashIn0
			elemtokenHashIn0 = append(elemtokenHashIn0, elemtokenHashIn1)
		}
		to.TokenHashIn = &elemtokenHashIn0
	}
	if o.TokenHashNin != nil {
		elemtokenHashNin0 := make([]string, 0)
		for _, otokenHashNin0 := range *o.TokenHashNin {
			elemtokenHashNin1 := otokenHashNin0
			elemtokenHashNin0 = append(elemtokenHashNin0, elemtokenHashNin1)
		}
		to.TokenHashNin = &elemtokenHashNin0
	}
	if o.TokenHashExists != nil {
		elemtokenHashExists0 := o.TokenHashExists
		to.TokenHashExists = elemtokenHashExists0
	}
	if o.TokenHashLike != nil {
		elemtokenHashLike0 := o.TokenHashLike
		to.TokenHashLike = elemtokenHashLike0
	}
	if o.TokenHashNlike != nil {
		elemtokenHashNlike0 := o.TokenHashNlike
		to.TokenHashNlike = elemtokenHashNlike0
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
	TokenHash  *int8 `json:"tokenHash,omitempty"`
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
	if s.TokenHash != nil {
		to.TokenHash = *s.TokenHash
	}
	if s.UpdatedAt != nil {
		to.UpdatedAt = *s.UpdatedAt
	}
	return to
}
