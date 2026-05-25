package registry_approved_guest

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_guest_access_level"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_guest_status"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id                 string
	AccessLevel        enum_guest_access_level.Value
	Created            actor_trace.Model
	EmailEnc           string
	EmailHash          string
	Name               string
	OwnerId            string
	RegistryId         string
	Status             enum_guest_status.Value
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
	if projection.AccessLevel {
		elemaccessLevel0 := m.AccessLevel
		r.AccessLevel = &elemaccessLevel0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToMongoRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
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
	if projection.OwnerId && m.OwnerId != "" {
		elemownerId0, err := primitive.ObjectIDFromHex(m.OwnerId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.OwnerId"), err)
		}
		r.OwnerId = &elemownerId0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0, err := primitive.ObjectIDFromHex(m.RegistryId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.RegistryId"), err)
		}
		r.RegistryId = &elemregistryId0
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
	if projection.AccessLevel {
		elemaccessLevel0 := m.AccessLevel
		r.AccessLevel = &elemaccessLevel0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToHTTPRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
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
	if projection.OwnerId && m.OwnerId != "" {
		elemownerId0 := m.OwnerId
		r.OwnerId = &elemownerId0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0 := m.RegistryId
		r.RegistryId = &elemregistryId0
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
type SelectByRegistryEmailUniqueQuery struct {
	RegistryId string
	EmailHash  string
}

type WhereClause struct {
	// id (Ref<RegistryApprovedGuest>) search options
	IdEq     *string
	IdIn     *[]string
	IdNin    *[]string
	IdExists *bool
	// accessLevel (GuestAccessLevel) search options
	AccessLevelEq     *enum_guest_access_level.Value
	AccessLevelNe     *enum_guest_access_level.Value
	AccessLevelGt     *enum_guest_access_level.Value
	AccessLevelGte    *enum_guest_access_level.Value
	AccessLevelLt     *enum_guest_access_level.Value
	AccessLevelLte    *enum_guest_access_level.Value
	AccessLevelIn     *[]enum_guest_access_level.Value
	AccessLevelNin    *[]enum_guest_access_level.Value
	AccessLevelExists *bool
	// created (ActorTrace) search options
	Created *actor_trace.WhereClause
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
	// ownerId (Ref<OwnerUser>) search options
	OwnerIdEq     *string
	OwnerIdIn     *[]string
	OwnerIdNin    *[]string
	OwnerIdExists *bool
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *string
	RegistryIdIn     *[]string
	RegistryIdNin    *[]string
	RegistryIdExists *bool
	// status (GuestStatus) search options
	StatusEq     *enum_guest_status.Value
	StatusNe     *enum_guest_status.Value
	StatusGt     *enum_guest_status.Value
	StatusGte    *enum_guest_status.Value
	StatusLt     *enum_guest_status.Value
	StatusLte    *enum_guest_status.Value
	StatusIn     *[]enum_guest_status.Value
	StatusNin    *[]enum_guest_status.Value
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
func (o SelectByRegistryEmailUniqueQuery) ToMongoSelectByRegistryEmailUniqueQuery() (MongoSelectByRegistryEmailUniqueQuery, error) {
	to := MongoSelectByRegistryEmailUniqueQuery{}
	elemregistryId0, err := primitive.ObjectIDFromHex(o.RegistryId)
	if err != nil {
		return to, errors.Join(errors.New("invalid o.RegistryId"), err)
	}
	to.RegistryId = elemregistryId0
	elememailHash0 := o.EmailHash
	to.EmailHash = elememailHash0
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
		elemcreated0, err := o.Created.ToMongoWhereClause()
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
	UpdatedAt  int8
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	to.CreatedAt = s.CreatedAt
	to.EmailHash = s.EmailHash
	to.OwnerId = s.OwnerId
	to.RegistryId = s.RegistryId
	to.UpdatedAt = s.UpdatedAt
	return to
}
