package shipping_address_request

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_address_request_status"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id                 string
	Created            actor_trace.Model
	DecisionReason     string
	EmailEnc           string
	EmailHash          string
	Name               string
	Note               string
	OwnerId            string
	PolicyVersion      int
	RegistryId         string
	RegistryItemId     string
	Status             enum_address_request_status.Value
	Updated            actor_trace.Model
	UpdatedByOwnerUser actor_trace.Model
}

func (m *Model) ToMongoRecord(projection Projection) (MongoRecord, error) {
	r := MongoRecord{}
	if m.Id != "" {
		elemid0, err := primitive.ObjectIDFromHex(m.Id)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.Id"), err)
		}
		r.Id = &elemid0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToMongoRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.DecisionReason {
		elemdecisionReason0 := m.DecisionReason
		r.DecisionReason = &elemdecisionReason0
	}
	if projection.EmailEnc {
		elememailEnc0 := m.EmailEnc
		r.EmailEnc = &elememailEnc0
	}
	if projection.EmailHash {
		elememailHash0 := m.EmailHash
		r.EmailHash = &elememailHash0
	}
	if projection.Name {
		elemname0 := m.Name
		r.Name = &elemname0
	}
	if projection.Note {
		elemnote0 := m.Note
		r.Note = &elemnote0
	}
	if projection.OwnerId && m.OwnerId != "" {
		elemownerId0, err := primitive.ObjectIDFromHex(m.OwnerId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.OwnerId"), err)
		}
		r.OwnerId = &elemownerId0
	}
	if projection.PolicyVersion {
		elempolicyVersion0 := m.PolicyVersion
		r.PolicyVersion = &elempolicyVersion0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0, err := primitive.ObjectIDFromHex(m.RegistryId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.RegistryId"), err)
		}
		r.RegistryId = &elemregistryId0
	}
	if projection.RegistryItemId && m.RegistryItemId != "" {
		elemregistryItemId0, err := primitive.ObjectIDFromHex(m.RegistryItemId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.RegistryItemId"), err)
		}
		r.RegistryItemId = &elemregistryItemId0
	}
	if projection.Status {
		elemstatus0 := m.Status
		r.Status = &elemstatus0
	}
	if projection.Updated {
		elemupdated0, err := m.Updated.ToMongoRecord(projection.UpdatedFields)
		if err != nil {
			return r, err
		}
		r.Updated = &elemupdated0
	}
	if projection.UpdatedByOwnerUser {
		elemupdatedByOwnerUser0, err := m.UpdatedByOwnerUser.ToMongoRecord(projection.UpdatedByOwnerUserFields)
		if err != nil {
			return r, err
		}
		r.UpdatedByOwnerUser = &elemupdatedByOwnerUser0
	}
	return r, nil
}

func (m *Model) ToHTTPRecord(projection Projection) (HTTPRecord, error) {
	r := HTTPRecord{}
	if m.Id != "" {
		elemid0 := m.Id
		r.Id = &elemid0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToHTTPRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.DecisionReason {
		elemdecisionReason0 := m.DecisionReason
		r.DecisionReason = &elemdecisionReason0
	}
	if projection.EmailEnc {
		elememailEnc0 := m.EmailEnc
		r.EmailEnc = &elememailEnc0
	}
	if projection.EmailHash {
		elememailHash0 := m.EmailHash
		r.EmailHash = &elememailHash0
	}
	if projection.Name {
		elemname0 := m.Name
		r.Name = &elemname0
	}
	if projection.Note {
		elemnote0 := m.Note
		r.Note = &elemnote0
	}
	if projection.OwnerId && m.OwnerId != "" {
		elemownerId0 := m.OwnerId
		r.OwnerId = &elemownerId0
	}
	if projection.PolicyVersion {
		elempolicyVersion0 := m.PolicyVersion
		r.PolicyVersion = &elempolicyVersion0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0 := m.RegistryId
		r.RegistryId = &elemregistryId0
	}
	if projection.RegistryItemId && m.RegistryItemId != "" {
		elemregistryItemId0 := m.RegistryItemId
		r.RegistryItemId = &elemregistryItemId0
	}
	if projection.Status {
		elemstatus0 := m.Status
		r.Status = &elemstatus0
	}
	if projection.Updated {
		elemupdated0, err := m.Updated.ToHTTPRecord(projection.UpdatedFields)
		if err != nil {
			return r, err
		}
		r.Updated = &elemupdated0
	}
	if projection.UpdatedByOwnerUser {
		elemupdatedByOwnerUser0, err := m.UpdatedByOwnerUser.ToHTTPRecord(projection.UpdatedByOwnerUserFields)
		if err != nil {
			return r, err
		}
		r.UpdatedByOwnerUser = &elemupdatedByOwnerUser0
	}
	return r, nil
}

type SelectByIdQuery struct {
	Id string
}

type WhereClause struct {
	// id (Ref<ShippingAddressRequest>) search options
	IdEq     *string
	IdIn     *[]string
	IdNin    *[]string
	IdExists *bool
	// created (ActorTrace) search options
	Created *actor_trace.WhereClause
	// decisionReason (string) search options
	DecisionReasonEq     *string
	DecisionReasonNe     *string
	DecisionReasonGt     *string
	DecisionReasonGte    *string
	DecisionReasonLt     *string
	DecisionReasonLte    *string
	DecisionReasonIn     *[]string
	DecisionReasonNin    *[]string
	DecisionReasonExists *bool
	DecisionReasonLike   *string
	DecisionReasonNlike  *string
	// emailEnc (string) search options
	EmailEncEq     *string
	EmailEncNe     *string
	EmailEncGt     *string
	EmailEncGte    *string
	EmailEncLt     *string
	EmailEncLte    *string
	EmailEncIn     *[]string
	EmailEncNin    *[]string
	EmailEncExists *bool
	EmailEncLike   *string
	EmailEncNlike  *string
	// emailHash (string) search options
	EmailHashEq     *string
	EmailHashNe     *string
	EmailHashGt     *string
	EmailHashGte    *string
	EmailHashLt     *string
	EmailHashLte    *string
	EmailHashIn     *[]string
	EmailHashNin    *[]string
	EmailHashExists *bool
	EmailHashLike   *string
	EmailHashNlike  *string
	// name (string) search options
	NameEq     *string
	NameNe     *string
	NameGt     *string
	NameGte    *string
	NameLt     *string
	NameLte    *string
	NameIn     *[]string
	NameNin    *[]string
	NameExists *bool
	NameLike   *string
	NameNlike  *string
	// note (string) search options
	NoteEq     *string
	NoteNe     *string
	NoteGt     *string
	NoteGte    *string
	NoteLt     *string
	NoteLte    *string
	NoteIn     *[]string
	NoteNin    *[]string
	NoteExists *bool
	NoteLike   *string
	NoteNlike  *string
	// ownerId (Ref<OwnerUser>) search options
	OwnerIdEq     *string
	OwnerIdIn     *[]string
	OwnerIdNin    *[]string
	OwnerIdExists *bool
	// policyVersion (int) search options
	PolicyVersionEq     *int
	PolicyVersionNe     *int
	PolicyVersionGt     *int
	PolicyVersionGte    *int
	PolicyVersionLt     *int
	PolicyVersionLte    *int
	PolicyVersionIn     *[]int
	PolicyVersionNin    *[]int
	PolicyVersionExists *bool
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *string
	RegistryIdIn     *[]string
	RegistryIdNin    *[]string
	RegistryIdExists *bool
	// registryItemId (Ref<RegistryItem>) search options
	RegistryItemIdEq     *string
	RegistryItemIdIn     *[]string
	RegistryItemIdNin    *[]string
	RegistryItemIdExists *bool
	// status (AddressRequestStatus) search options
	StatusEq     *enum_address_request_status.Value
	StatusNe     *enum_address_request_status.Value
	StatusGt     *enum_address_request_status.Value
	StatusGte    *enum_address_request_status.Value
	StatusLt     *enum_address_request_status.Value
	StatusLte    *enum_address_request_status.Value
	StatusIn     *[]enum_address_request_status.Value
	StatusNin    *[]enum_address_request_status.Value
	StatusExists *bool
	// updated (ActorTrace) search options
	Updated *actor_trace.WhereClause
	// updatedByOwnerUser (ActorTrace) search options
	UpdatedByOwnerUser *actor_trace.WhereClause
}

func (o SelectByIdQuery) ToMongoSelectByIdQuery() (MongoSelectByIdQuery, error) {
	to := MongoSelectByIdQuery{}
	elemid0, err := primitive.ObjectIDFromHex(o.Id)
	if err != nil {
		return to, errors.Join(errors.New("invalid o.Id"), err)
	}
	to.Id = elemid0
	return to, nil
}

func (o WhereClause) ToMongoWhereClause() (MongoWhereClause, error) {
	to := MongoWhereClause{}
	if o.IdEq != nil {
		elemidEq0, err := primitive.ObjectIDFromHex(*o.IdEq)
		if err != nil {
			return to, errors.Join(errors.New("invalid o.IdEq"), err)
		}
		to.IdEq = &elemidEq0
	}
	if o.IdIn != nil {
		elemidIn0 := make([]primitive.ObjectID, 0)
		for _, oidIn0 := range *o.IdIn {
			elemidIn1, err := primitive.ObjectIDFromHex(oidIn0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oidIn0"), err)
			}
			elemidIn0 = append(elemidIn0, elemidIn1)
		}
		to.IdIn = &elemidIn0
	}
	if o.IdNin != nil {
		elemidNin0 := make([]primitive.ObjectID, 0)
		for _, oidNin0 := range *o.IdNin {
			elemidNin1, err := primitive.ObjectIDFromHex(oidNin0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oidNin0"), err)
			}
			elemidNin0 = append(elemidNin0, elemidNin1)
		}
		to.IdNin = &elemidNin0
	}
	if o.IdExists != nil {
		elemidExists0 := o.IdExists
		to.IdExists = elemidExists0
	}
	if o.Created != nil {
		elemcreated0, err := o.Created.ToMongoWhereClause()
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
		elemownerIdEq0, err := primitive.ObjectIDFromHex(*o.OwnerIdEq)
		if err != nil {
			return to, errors.Join(errors.New("invalid o.OwnerIdEq"), err)
		}
		to.OwnerIdEq = &elemownerIdEq0
	}
	if o.OwnerIdIn != nil {
		elemownerIdIn0 := make([]primitive.ObjectID, 0)
		for _, oownerIdIn0 := range *o.OwnerIdIn {
			elemownerIdIn1, err := primitive.ObjectIDFromHex(oownerIdIn0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oownerIdIn0"), err)
			}
			elemownerIdIn0 = append(elemownerIdIn0, elemownerIdIn1)
		}
		to.OwnerIdIn = &elemownerIdIn0
	}
	if o.OwnerIdNin != nil {
		elemownerIdNin0 := make([]primitive.ObjectID, 0)
		for _, oownerIdNin0 := range *o.OwnerIdNin {
			elemownerIdNin1, err := primitive.ObjectIDFromHex(oownerIdNin0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oownerIdNin0"), err)
			}
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
		elemregistryIdEq0, err := primitive.ObjectIDFromHex(*o.RegistryIdEq)
		if err != nil {
			return to, errors.Join(errors.New("invalid o.RegistryIdEq"), err)
		}
		to.RegistryIdEq = &elemregistryIdEq0
	}
	if o.RegistryIdIn != nil {
		elemregistryIdIn0 := make([]primitive.ObjectID, 0)
		for _, oregistryIdIn0 := range *o.RegistryIdIn {
			elemregistryIdIn1, err := primitive.ObjectIDFromHex(oregistryIdIn0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oregistryIdIn0"), err)
			}
			elemregistryIdIn0 = append(elemregistryIdIn0, elemregistryIdIn1)
		}
		to.RegistryIdIn = &elemregistryIdIn0
	}
	if o.RegistryIdNin != nil {
		elemregistryIdNin0 := make([]primitive.ObjectID, 0)
		for _, oregistryIdNin0 := range *o.RegistryIdNin {
			elemregistryIdNin1, err := primitive.ObjectIDFromHex(oregistryIdNin0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oregistryIdNin0"), err)
			}
			elemregistryIdNin0 = append(elemregistryIdNin0, elemregistryIdNin1)
		}
		to.RegistryIdNin = &elemregistryIdNin0
	}
	if o.RegistryIdExists != nil {
		elemregistryIdExists0 := o.RegistryIdExists
		to.RegistryIdExists = elemregistryIdExists0
	}
	if o.RegistryItemIdEq != nil {
		elemregistryItemIdEq0, err := primitive.ObjectIDFromHex(*o.RegistryItemIdEq)
		if err != nil {
			return to, errors.Join(errors.New("invalid o.RegistryItemIdEq"), err)
		}
		to.RegistryItemIdEq = &elemregistryItemIdEq0
	}
	if o.RegistryItemIdIn != nil {
		elemregistryItemIdIn0 := make([]primitive.ObjectID, 0)
		for _, oregistryItemIdIn0 := range *o.RegistryItemIdIn {
			elemregistryItemIdIn1, err := primitive.ObjectIDFromHex(oregistryItemIdIn0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oregistryItemIdIn0"), err)
			}
			elemregistryItemIdIn0 = append(elemregistryItemIdIn0, elemregistryItemIdIn1)
		}
		to.RegistryItemIdIn = &elemregistryItemIdIn0
	}
	if o.RegistryItemIdNin != nil {
		elemregistryItemIdNin0 := make([]primitive.ObjectID, 0)
		for _, oregistryItemIdNin0 := range *o.RegistryItemIdNin {
			elemregistryItemIdNin1, err := primitive.ObjectIDFromHex(oregistryItemIdNin0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oregistryItemIdNin0"), err)
			}
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
		elemupdated0, err := o.Updated.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.Updated = &elemupdated0
	}
	if o.UpdatedByOwnerUser != nil {
		elemupdatedByOwnerUser0, err := o.UpdatedByOwnerUser.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.UpdatedByOwnerUser = &elemupdatedByOwnerUser0
	}
	return to, nil
}

type SortParams struct {
	CreatedAt  int8
	EmailHash  int8
	OwnerId    int8
	RegistryId int8
	Status     int8
	UpdatedAt  int8
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	to.CreatedAt = s.CreatedAt
	to.EmailHash = s.EmailHash
	to.OwnerId = s.OwnerId
	to.RegistryId = s.RegistryId
	to.Status = s.Status
	to.UpdatedAt = s.UpdatedAt
	return to
}
