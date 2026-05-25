package address_access_session

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Model struct {
	Id                   string
	ApprovedGuestId      string
	Created              actor_trace.Model
	EmailHash            string
	ExpiresAt            time.Time
	OwnerId              string
	PolicyVersionAtIssue int
	RegistryId           string
	TokenHash            string
	Updated              actor_trace.Model
	UpdatedByOwnerUser   actor_trace.Model
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
	if projection.ApprovedGuestId && m.ApprovedGuestId != "" {
		elemapprovedGuestId0, err := primitive.ObjectIDFromHex(m.ApprovedGuestId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.ApprovedGuestId"), err)
		}
		r.ApprovedGuestId = &elemapprovedGuestId0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToMongoRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.EmailHash {
		elememailHash0 := m.EmailHash
		r.EmailHash = &elememailHash0
	}
	if projection.ExpiresAt {
		elemexpiresAt0 := m.ExpiresAt
		r.ExpiresAt = &elemexpiresAt0
	}
	if projection.OwnerId && m.OwnerId != "" {
		elemownerId0, err := primitive.ObjectIDFromHex(m.OwnerId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.OwnerId"), err)
		}
		r.OwnerId = &elemownerId0
	}
	if projection.PolicyVersionAtIssue {
		elempolicyVersionAtIssue0 := m.PolicyVersionAtIssue
		r.PolicyVersionAtIssue = &elempolicyVersionAtIssue0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0, err := primitive.ObjectIDFromHex(m.RegistryId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.RegistryId"), err)
		}
		r.RegistryId = &elemregistryId0
	}
	if projection.TokenHash {
		elemtokenHash0 := m.TokenHash
		r.TokenHash = &elemtokenHash0
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
	if projection.ApprovedGuestId && m.ApprovedGuestId != "" {
		elemapprovedGuestId0 := m.ApprovedGuestId
		r.ApprovedGuestId = &elemapprovedGuestId0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToHTTPRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.EmailHash {
		elememailHash0 := m.EmailHash
		r.EmailHash = &elememailHash0
	}
	if projection.ExpiresAt {
		elemexpiresAt0 := m.ExpiresAt
		r.ExpiresAt = &elemexpiresAt0
	}
	if projection.OwnerId && m.OwnerId != "" {
		elemownerId0 := m.OwnerId
		r.OwnerId = &elemownerId0
	}
	if projection.PolicyVersionAtIssue {
		elempolicyVersionAtIssue0 := m.PolicyVersionAtIssue
		r.PolicyVersionAtIssue = &elempolicyVersionAtIssue0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0 := m.RegistryId
		r.RegistryId = &elemregistryId0
	}
	if projection.TokenHash {
		elemtokenHash0 := m.TokenHash
		r.TokenHash = &elemtokenHash0
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
type SelectByTokenUniqueQuery struct {
	TokenHash string
}

type WhereClause struct {
	// id (Ref<AddressAccessSession>) search options
	IdEq     *string
	IdIn     *[]string
	IdNin    *[]string
	IdExists *bool
	// approvedGuestId (Ref<RegistryApprovedGuest>) search options
	ApprovedGuestIdEq     *string
	ApprovedGuestIdIn     *[]string
	ApprovedGuestIdNin    *[]string
	ApprovedGuestIdExists *bool
	// created (ActorTrace) search options
	Created *actor_trace.WhereClause
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
	// expiresAt (timestamp) search options
	ExpiresAtEq     *time.Time
	ExpiresAtNe     *time.Time
	ExpiresAtGt     *time.Time
	ExpiresAtGte    *time.Time
	ExpiresAtLt     *time.Time
	ExpiresAtLte    *time.Time
	ExpiresAtIn     *[]time.Time
	ExpiresAtNin    *[]time.Time
	ExpiresAtExists *bool
	// ownerId (Ref<OwnerUser>) search options
	OwnerIdEq     *string
	OwnerIdIn     *[]string
	OwnerIdNin    *[]string
	OwnerIdExists *bool
	// policyVersionAtIssue (int) search options
	PolicyVersionAtIssueEq     *int
	PolicyVersionAtIssueNe     *int
	PolicyVersionAtIssueGt     *int
	PolicyVersionAtIssueGte    *int
	PolicyVersionAtIssueLt     *int
	PolicyVersionAtIssueLte    *int
	PolicyVersionAtIssueIn     *[]int
	PolicyVersionAtIssueNin    *[]int
	PolicyVersionAtIssueExists *bool
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *string
	RegistryIdIn     *[]string
	RegistryIdNin    *[]string
	RegistryIdExists *bool
	// tokenHash (string) search options
	TokenHashEq     *string
	TokenHashNe     *string
	TokenHashGt     *string
	TokenHashGte    *string
	TokenHashLt     *string
	TokenHashLte    *string
	TokenHashIn     *[]string
	TokenHashNin    *[]string
	TokenHashExists *bool
	TokenHashLike   *string
	TokenHashNlike  *string
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
func (o SelectByTokenUniqueQuery) ToMongoSelectByTokenUniqueQuery() (MongoSelectByTokenUniqueQuery, error) {
	to := MongoSelectByTokenUniqueQuery{}
	elemtokenHash0 := o.TokenHash
	to.TokenHash = elemtokenHash0
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
	if o.ApprovedGuestIdEq != nil {
		elemapprovedGuestIdEq0, err := primitive.ObjectIDFromHex(*o.ApprovedGuestIdEq)
		if err != nil {
			return to, errors.Join(errors.New("invalid o.ApprovedGuestIdEq"), err)
		}
		to.ApprovedGuestIdEq = &elemapprovedGuestIdEq0
	}
	if o.ApprovedGuestIdIn != nil {
		elemapprovedGuestIdIn0 := make([]primitive.ObjectID, 0)
		for _, oapprovedGuestIdIn0 := range *o.ApprovedGuestIdIn {
			elemapprovedGuestIdIn1, err := primitive.ObjectIDFromHex(oapprovedGuestIdIn0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oapprovedGuestIdIn0"), err)
			}
			elemapprovedGuestIdIn0 = append(elemapprovedGuestIdIn0, elemapprovedGuestIdIn1)
		}
		to.ApprovedGuestIdIn = &elemapprovedGuestIdIn0
	}
	if o.ApprovedGuestIdNin != nil {
		elemapprovedGuestIdNin0 := make([]primitive.ObjectID, 0)
		for _, oapprovedGuestIdNin0 := range *o.ApprovedGuestIdNin {
			elemapprovedGuestIdNin1, err := primitive.ObjectIDFromHex(oapprovedGuestIdNin0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oapprovedGuestIdNin0"), err)
			}
			elemapprovedGuestIdNin0 = append(elemapprovedGuestIdNin0, elemapprovedGuestIdNin1)
		}
		to.ApprovedGuestIdNin = &elemapprovedGuestIdNin0
	}
	if o.ApprovedGuestIdExists != nil {
		elemapprovedGuestIdExists0 := o.ApprovedGuestIdExists
		to.ApprovedGuestIdExists = elemapprovedGuestIdExists0
	}
	if o.Created != nil {
		elemcreated0, err := o.Created.ToMongoWhereClause()
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
	TokenHash  int8
	UpdatedAt  int8
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	to.CreatedAt = s.CreatedAt
	to.EmailHash = s.EmailHash
	to.OwnerId = s.OwnerId
	to.RegistryId = s.RegistryId
	to.TokenHash = s.TokenHash
	to.UpdatedAt = s.UpdatedAt
	return to
}
