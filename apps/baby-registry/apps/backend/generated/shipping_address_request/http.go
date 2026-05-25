package shipping_address_request

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_address_request_status"
)

type HTTPRecord struct {
	Id                 *string                            `json:"id,omitempty"`
	Created            *actor_trace.HTTPRecord            `json:"created,omitempty"`
	DecisionReason     *string                            `json:"decisionReason,omitempty"`
	EmailEnc           *string                            `json:"emailEnc,omitempty"`
	EmailHash          *string                            `json:"emailHash,omitempty"`
	Name               *string                            `json:"name,omitempty"`
	Note               *string                            `json:"note,omitempty"`
	OwnerId            *string                            `json:"ownerId,omitempty"`
	PolicyVersion      *int                               `json:"policyVersion,omitempty"`
	RegistryId         *string                            `json:"registryId,omitempty"`
	RegistryItemId     *string                            `json:"registryItemId,omitempty"`
	Status             *enum_address_request_status.Value `json:"status,omitempty"`
	Updated            *actor_trace.HTTPRecord            `json:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.HTTPRecord            `json:"updatedByOwnerUser,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id
		m.Id = *elemid0
	}
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.DecisionReason != nil {
		elemdecisionReason0 := r.DecisionReason
		m.DecisionReason = *elemdecisionReason0
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
	if r.Note != nil {
		elemnote0 := r.Note
		m.Note = *elemnote0
	}
	if r.OwnerId != nil {
		elemownerId0 := r.OwnerId
		m.OwnerId = *elemownerId0
	}
	if r.PolicyVersion != nil {
		elempolicyVersion0 := r.PolicyVersion
		m.PolicyVersion = *elempolicyVersion0
	}
	if r.RegistryId != nil {
		elemregistryId0 := r.RegistryId
		m.RegistryId = *elemregistryId0
	}
	if r.RegistryItemId != nil {
		elemregistryItemId0 := r.RegistryItemId
		m.RegistryItemId = *elemregistryItemId0
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
	if r.Created != nil {
		p.Created = true
		p.CreatedFields = actor_trace.NewProjection(true)
	}
	if r.DecisionReason != nil {
		p.DecisionReason = true
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
	if r.Note != nil {
		p.Note = true
	}
	if r.OwnerId != nil {
		p.OwnerId = true
	}
	if r.PolicyVersion != nil {
		p.PolicyVersion = true
	}
	if r.RegistryId != nil {
		p.RegistryId = true
	}
	if r.RegistryItemId != nil {
		p.RegistryItemId = true
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

type HTTPWhereClause struct {
	// id (Ref<ShippingAddressRequest>) search options
	IdEq     *string   `json:"idEq,omitempty"`
	IdIn     *[]string `json:"idIn,omitempty"`
	IdNin    *[]string `json:"idNin,omitempty"`
	IdExists *bool     `json:"idExists,omitempty"`
	// created (ActorTrace) search options
	Created *actor_trace.HTTPWhereClause `json:"created,omitempty"`
	// decisionReason (string) search options
	DecisionReasonEq     *string   `json:"decisionReasonEq,omitempty"`
	DecisionReasonNe     *string   `json:"decisionReasonNe,omitempty"`
	DecisionReasonGt     *string   `json:"decisionReasonGt,omitempty"`
	DecisionReasonGte    *string   `json:"decisionReasonGte,omitempty"`
	DecisionReasonLt     *string   `json:"decisionReasonLt,omitempty"`
	DecisionReasonLte    *string   `json:"decisionReasonLte,omitempty"`
	DecisionReasonIn     *[]string `json:"decisionReasonIn,omitempty"`
	DecisionReasonNin    *[]string `json:"decisionReasonNin,omitempty"`
	DecisionReasonExists *bool     `json:"decisionReasonExists,omitempty"`
	DecisionReasonLike   *string   `json:"decisionReasonLike,omitempty"`
	DecisionReasonNlike  *string   `json:"decisionReasonNlike,omitempty"`
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
	// note (string) search options
	NoteEq     *string   `json:"noteEq,omitempty"`
	NoteNe     *string   `json:"noteNe,omitempty"`
	NoteGt     *string   `json:"noteGt,omitempty"`
	NoteGte    *string   `json:"noteGte,omitempty"`
	NoteLt     *string   `json:"noteLt,omitempty"`
	NoteLte    *string   `json:"noteLte,omitempty"`
	NoteIn     *[]string `json:"noteIn,omitempty"`
	NoteNin    *[]string `json:"noteNin,omitempty"`
	NoteExists *bool     `json:"noteExists,omitempty"`
	NoteLike   *string   `json:"noteLike,omitempty"`
	NoteNlike  *string   `json:"noteNlike,omitempty"`
	// ownerId (Ref<OwnerUser>) search options
	OwnerIdEq     *string   `json:"ownerIdEq,omitempty"`
	OwnerIdIn     *[]string `json:"ownerIdIn,omitempty"`
	OwnerIdNin    *[]string `json:"ownerIdNin,omitempty"`
	OwnerIdExists *bool     `json:"ownerIdExists,omitempty"`
	// policyVersion (int) search options
	PolicyVersionEq     *int   `json:"policyVersionEq,omitempty"`
	PolicyVersionNe     *int   `json:"policyVersionNe,omitempty"`
	PolicyVersionGt     *int   `json:"policyVersionGt,omitempty"`
	PolicyVersionGte    *int   `json:"policyVersionGte,omitempty"`
	PolicyVersionLt     *int   `json:"policyVersionLt,omitempty"`
	PolicyVersionLte    *int   `json:"policyVersionLte,omitempty"`
	PolicyVersionIn     *[]int `json:"policyVersionIn,omitempty"`
	PolicyVersionNin    *[]int `json:"policyVersionNin,omitempty"`
	PolicyVersionExists *bool  `json:"policyVersionExists,omitempty"`
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *string   `json:"registryIdEq,omitempty"`
	RegistryIdIn     *[]string `json:"registryIdIn,omitempty"`
	RegistryIdNin    *[]string `json:"registryIdNin,omitempty"`
	RegistryIdExists *bool     `json:"registryIdExists,omitempty"`
	// registryItemId (Ref<RegistryItem>) search options
	RegistryItemIdEq     *string   `json:"registryItemIdEq,omitempty"`
	RegistryItemIdIn     *[]string `json:"registryItemIdIn,omitempty"`
	RegistryItemIdNin    *[]string `json:"registryItemIdNin,omitempty"`
	RegistryItemIdExists *bool     `json:"registryItemIdExists,omitempty"`
	// status (AddressRequestStatus) search options
	StatusEq     *enum_address_request_status.Value   `json:"statusEq,omitempty"`
	StatusNe     *enum_address_request_status.Value   `json:"statusNe,omitempty"`
	StatusGt     *enum_address_request_status.Value   `json:"statusGt,omitempty"`
	StatusGte    *enum_address_request_status.Value   `json:"statusGte,omitempty"`
	StatusLt     *enum_address_request_status.Value   `json:"statusLt,omitempty"`
	StatusLte    *enum_address_request_status.Value   `json:"statusLte,omitempty"`
	StatusIn     *[]enum_address_request_status.Value `json:"statusIn,omitempty"`
	StatusNin    *[]enum_address_request_status.Value `json:"statusNin,omitempty"`
	StatusExists *bool                                `json:"statusExists,omitempty"`
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
	if o.Created != nil {
		elemcreated0, err := o.Created.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.Created = &elemcreated0
	}
	if o.DecisionReasonEq != nil {
		elemdecisionReasonEq0 := o.DecisionReasonEq
		to.DecisionReasonEq = elemdecisionReasonEq0
	}
	if o.DecisionReasonNe != nil {
		elemdecisionReasonNe0 := o.DecisionReasonNe
		to.DecisionReasonNe = elemdecisionReasonNe0
	}
	if o.DecisionReasonGt != nil {
		elemdecisionReasonGt0 := o.DecisionReasonGt
		to.DecisionReasonGt = elemdecisionReasonGt0
	}
	if o.DecisionReasonGte != nil {
		elemdecisionReasonGte0 := o.DecisionReasonGte
		to.DecisionReasonGte = elemdecisionReasonGte0
	}
	if o.DecisionReasonLt != nil {
		elemdecisionReasonLt0 := o.DecisionReasonLt
		to.DecisionReasonLt = elemdecisionReasonLt0
	}
	if o.DecisionReasonLte != nil {
		elemdecisionReasonLte0 := o.DecisionReasonLte
		to.DecisionReasonLte = elemdecisionReasonLte0
	}
	if o.DecisionReasonIn != nil {
		elemdecisionReasonIn0 := make([]string, 0)
		for _, odecisionReasonIn0 := range *o.DecisionReasonIn {
			elemdecisionReasonIn1 := odecisionReasonIn0
			elemdecisionReasonIn0 = append(elemdecisionReasonIn0, elemdecisionReasonIn1)
		}
		to.DecisionReasonIn = &elemdecisionReasonIn0
	}
	if o.DecisionReasonNin != nil {
		elemdecisionReasonNin0 := make([]string, 0)
		for _, odecisionReasonNin0 := range *o.DecisionReasonNin {
			elemdecisionReasonNin1 := odecisionReasonNin0
			elemdecisionReasonNin0 = append(elemdecisionReasonNin0, elemdecisionReasonNin1)
		}
		to.DecisionReasonNin = &elemdecisionReasonNin0
	}
	if o.DecisionReasonExists != nil {
		elemdecisionReasonExists0 := o.DecisionReasonExists
		to.DecisionReasonExists = elemdecisionReasonExists0
	}
	if o.DecisionReasonLike != nil {
		elemdecisionReasonLike0 := o.DecisionReasonLike
		to.DecisionReasonLike = elemdecisionReasonLike0
	}
	if o.DecisionReasonNlike != nil {
		elemdecisionReasonNlike0 := o.DecisionReasonNlike
		to.DecisionReasonNlike = elemdecisionReasonNlike0
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
	if o.NoteEq != nil {
		elemnoteEq0 := o.NoteEq
		to.NoteEq = elemnoteEq0
	}
	if o.NoteNe != nil {
		elemnoteNe0 := o.NoteNe
		to.NoteNe = elemnoteNe0
	}
	if o.NoteGt != nil {
		elemnoteGt0 := o.NoteGt
		to.NoteGt = elemnoteGt0
	}
	if o.NoteGte != nil {
		elemnoteGte0 := o.NoteGte
		to.NoteGte = elemnoteGte0
	}
	if o.NoteLt != nil {
		elemnoteLt0 := o.NoteLt
		to.NoteLt = elemnoteLt0
	}
	if o.NoteLte != nil {
		elemnoteLte0 := o.NoteLte
		to.NoteLte = elemnoteLte0
	}
	if o.NoteIn != nil {
		elemnoteIn0 := make([]string, 0)
		for _, onoteIn0 := range *o.NoteIn {
			elemnoteIn1 := onoteIn0
			elemnoteIn0 = append(elemnoteIn0, elemnoteIn1)
		}
		to.NoteIn = &elemnoteIn0
	}
	if o.NoteNin != nil {
		elemnoteNin0 := make([]string, 0)
		for _, onoteNin0 := range *o.NoteNin {
			elemnoteNin1 := onoteNin0
			elemnoteNin0 = append(elemnoteNin0, elemnoteNin1)
		}
		to.NoteNin = &elemnoteNin0
	}
	if o.NoteExists != nil {
		elemnoteExists0 := o.NoteExists
		to.NoteExists = elemnoteExists0
	}
	if o.NoteLike != nil {
		elemnoteLike0 := o.NoteLike
		to.NoteLike = elemnoteLike0
	}
	if o.NoteNlike != nil {
		elemnoteNlike0 := o.NoteNlike
		to.NoteNlike = elemnoteNlike0
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
	if o.PolicyVersionEq != nil {
		elempolicyVersionEq0 := o.PolicyVersionEq
		to.PolicyVersionEq = elempolicyVersionEq0
	}
	if o.PolicyVersionNe != nil {
		elempolicyVersionNe0 := o.PolicyVersionNe
		to.PolicyVersionNe = elempolicyVersionNe0
	}
	if o.PolicyVersionGt != nil {
		elempolicyVersionGt0 := o.PolicyVersionGt
		to.PolicyVersionGt = elempolicyVersionGt0
	}
	if o.PolicyVersionGte != nil {
		elempolicyVersionGte0 := o.PolicyVersionGte
		to.PolicyVersionGte = elempolicyVersionGte0
	}
	if o.PolicyVersionLt != nil {
		elempolicyVersionLt0 := o.PolicyVersionLt
		to.PolicyVersionLt = elempolicyVersionLt0
	}
	if o.PolicyVersionLte != nil {
		elempolicyVersionLte0 := o.PolicyVersionLte
		to.PolicyVersionLte = elempolicyVersionLte0
	}
	if o.PolicyVersionIn != nil {
		elempolicyVersionIn0 := make([]int, 0)
		for _, opolicyVersionIn0 := range *o.PolicyVersionIn {
			elempolicyVersionIn1 := opolicyVersionIn0
			elempolicyVersionIn0 = append(elempolicyVersionIn0, elempolicyVersionIn1)
		}
		to.PolicyVersionIn = &elempolicyVersionIn0
	}
	if o.PolicyVersionNin != nil {
		elempolicyVersionNin0 := make([]int, 0)
		for _, opolicyVersionNin0 := range *o.PolicyVersionNin {
			elempolicyVersionNin1 := opolicyVersionNin0
			elempolicyVersionNin0 = append(elempolicyVersionNin0, elempolicyVersionNin1)
		}
		to.PolicyVersionNin = &elempolicyVersionNin0
	}
	if o.PolicyVersionExists != nil {
		elempolicyVersionExists0 := o.PolicyVersionExists
		to.PolicyVersionExists = elempolicyVersionExists0
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
	if o.RegistryItemIdEq != nil {
		elemregistryItemIdEq0 := o.RegistryItemIdEq
		to.RegistryItemIdEq = elemregistryItemIdEq0
	}
	if o.RegistryItemIdIn != nil {
		elemregistryItemIdIn0 := make([]string, 0)
		for _, oregistryItemIdIn0 := range *o.RegistryItemIdIn {
			elemregistryItemIdIn1 := oregistryItemIdIn0
			elemregistryItemIdIn0 = append(elemregistryItemIdIn0, elemregistryItemIdIn1)
		}
		to.RegistryItemIdIn = &elemregistryItemIdIn0
	}
	if o.RegistryItemIdNin != nil {
		elemregistryItemIdNin0 := make([]string, 0)
		for _, oregistryItemIdNin0 := range *o.RegistryItemIdNin {
			elemregistryItemIdNin1 := oregistryItemIdNin0
			elemregistryItemIdNin0 = append(elemregistryItemIdNin0, elemregistryItemIdNin1)
		}
		to.RegistryItemIdNin = &elemregistryItemIdNin0
	}
	if o.RegistryItemIdExists != nil {
		elemregistryItemIdExists0 := o.RegistryItemIdExists
		to.RegistryItemIdExists = elemregistryItemIdExists0
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
		elemstatusIn0 := make([]enum_address_request_status.Value, 0)
		for _, ostatusIn0 := range *o.StatusIn {
			elemstatusIn1 := ostatusIn0
			elemstatusIn0 = append(elemstatusIn0, elemstatusIn1)
		}
		to.StatusIn = &elemstatusIn0
	}
	if o.StatusNin != nil {
		elemstatusNin0 := make([]enum_address_request_status.Value, 0)
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
	Status     *int8 `json:"status,omitempty"`
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
	if s.Status != nil {
		to.Status = *s.Status
	}
	if s.UpdatedAt != nil {
		to.UpdatedAt = *s.UpdatedAt
	}
	return to
}
