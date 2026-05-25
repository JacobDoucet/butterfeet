package address_access_session

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoRecord struct {
	Id                   *primitive.ObjectID      `bson:"_id,omitempty"`
	ApprovedGuestId      *primitive.ObjectID      `bson:"approvedGuestId,omitempty"`
	Created              *actor_trace.MongoRecord `bson:"created,omitempty"`
	EmailHash            *string                  `bson:"emailHash,omitempty"`
	ExpiresAt            *time.Time               `bson:"expiresAt,omitempty"`
	OwnerId              *primitive.ObjectID      `bson:"ownerId,omitempty"`
	PolicyVersionAtIssue *int                     `bson:"policyVersionAtIssue,omitempty"`
	RegistryId           *primitive.ObjectID      `bson:"registryId,omitempty"`
	TokenHash            *string                  `bson:"tokenHash,omitempty"`
	Updated              *actor_trace.MongoRecord `bson:"updated,omitempty"`
	UpdatedByOwnerUser   *actor_trace.MongoRecord `bson:"updatedByOwnerUser,omitempty"`
}

type MongoUpdateWhereClause struct {
	Id        primitive.ObjectID
	OwnerIdIn *[]primitive.ObjectID
}

func (r *MongoRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id.Hex()
		m.Id = elemid0
	}
	if r.ApprovedGuestId != nil {
		elemapprovedGuestId0 := r.ApprovedGuestId.Hex()
		m.ApprovedGuestId = elemapprovedGuestId0
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
		elemownerId0 := r.OwnerId.Hex()
		m.OwnerId = elemownerId0
	}
	if r.PolicyVersionAtIssue != nil {
		elempolicyVersionAtIssue0 := r.PolicyVersionAtIssue
		m.PolicyVersionAtIssue = *elempolicyVersionAtIssue0
	}
	if r.RegistryId != nil {
		elemregistryId0 := r.RegistryId.Hex()
		m.RegistryId = elemregistryId0
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

type MongoSelectByIdQuery struct {
	Id primitive.ObjectID
}
type MongoSelectByTokenUniqueQuery struct {
	TokenHash string
}

type MongoWhereClause struct {
	// id (Ref<AddressAccessSession>) search options
	IdEq     *primitive.ObjectID
	IdIn     *[]primitive.ObjectID
	IdNin    *[]primitive.ObjectID
	IdExists *bool
	// approvedGuestId (Ref<RegistryApprovedGuest>) search options
	ApprovedGuestIdEq     *primitive.ObjectID
	ApprovedGuestIdIn     *[]primitive.ObjectID
	ApprovedGuestIdNin    *[]primitive.ObjectID
	ApprovedGuestIdExists *bool
	// created (ActorTrace) search options
	Created *actor_trace.MongoWhereClause
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
	OwnerIdEq     *primitive.ObjectID
	OwnerIdIn     *[]primitive.ObjectID
	OwnerIdNin    *[]primitive.ObjectID
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
	RegistryIdEq     *primitive.ObjectID
	RegistryIdIn     *[]primitive.ObjectID
	RegistryIdNin    *[]primitive.ObjectID
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
	Updated *actor_trace.MongoWhereClause
	// updatedByOwnerUser (ActorTrace) search options
	UpdatedByOwnerUser *actor_trace.MongoWhereClause
}

type MongoLookup interface {
	GetQueryParts() (bson.A, error)
	GetLookupQuery() (bson.M, error)
}

func (o MongoWhereClause) GetLookupQuery() (bson.M, error) {
	query := bson.M{}
	and, err := o.GetQueryParts()
	if err != nil {
		return nil, err
	}
	if len(and) > 0 {
		query["$and"] = and
	}
	return query, nil
}

func (o MongoWhereClause) GetQueryParts() (bson.A, error) {
	and := bson.A{}
	if o.IdEq != nil {
		query := bson.M{}
		query["_id"] = o.IdEq
		and = append(and, query)
	}
	if o.IdIn != nil {
		query := bson.M{}
		query["_id"] = bson.M{"$in": o.IdIn}
		and = append(and, query)
	}
	if o.IdNin != nil {
		query := bson.M{}
		query["_id"] = bson.M{"$nin": o.IdNin}
		and = append(and, query)
	}
	if o.IdExists != nil {
		query := bson.M{}
		query["_id"] = bson.M{"$exists": *o.IdExists}
		and = append(and, query)
	}
	if o.ApprovedGuestIdEq != nil {
		query := bson.M{}
		query["approvedGuestId"] = o.ApprovedGuestIdEq
		and = append(and, query)
	}
	if o.ApprovedGuestIdIn != nil {
		query := bson.M{}
		query["approvedGuestId"] = bson.M{"$in": o.ApprovedGuestIdIn}
		and = append(and, query)
	}
	if o.ApprovedGuestIdNin != nil {
		query := bson.M{}
		query["approvedGuestId"] = bson.M{"$nin": o.ApprovedGuestIdNin}
		and = append(and, query)
	}
	if o.ApprovedGuestIdExists != nil {
		query := bson.M{}
		query["approvedGuestId"] = bson.M{"$exists": *o.ApprovedGuestIdExists}
		and = append(and, query)
	}
	if o.Created != nil {
		query := bson.M{}
		createdQuery, err := o.Created.GetQueryParts()
		if err != nil {
			return nil, err
		}
		for _, part := range createdQuery {
			partAsBsonM, ok := part.(bson.M)
			if !ok {
				continue
			}
			for k, v := range partAsBsonM {
				query["created."+k] = v
			}
		}
		and = append(and, query)
	}
	if o.EmailHashEq != nil {
		query := bson.M{}
		query["emailHash"] = o.EmailHashEq
		and = append(and, query)
	}
	if o.EmailHashNe != nil {
		query := bson.M{}
		query["emailHash"] = bson.M{"$ne": o.EmailHashNe}
		and = append(and, query)
	}
	if o.EmailHashGt != nil {
		query := bson.M{}
		query["emailHash"] = bson.M{"$gt": o.EmailHashGt}
		and = append(and, query)
	}
	if o.EmailHashGte != nil {
		query := bson.M{}
		query["emailHash"] = bson.M{"$gte": o.EmailHashGte}
		and = append(and, query)
	}
	if o.EmailHashLt != nil {
		query := bson.M{}
		query["emailHash"] = bson.M{"$lt": o.EmailHashLt}
		and = append(and, query)
	}
	if o.EmailHashLte != nil {
		query := bson.M{}
		query["emailHash"] = bson.M{"$lte": o.EmailHashLte}
		and = append(and, query)
	}
	if o.EmailHashIn != nil {
		query := bson.M{}
		query["emailHash"] = bson.M{"$in": o.EmailHashIn}
		and = append(and, query)
	}
	if o.EmailHashNin != nil {
		query := bson.M{}
		query["emailHash"] = bson.M{"$nin": o.EmailHashNin}
		and = append(and, query)
	}
	if o.EmailHashExists != nil {
		query := bson.M{}
		query["emailHash"] = bson.M{"$exists": *o.EmailHashExists}
		and = append(and, query)
	}
	if o.EmailHashLike != nil {
		query := bson.M{}
		query["emailHash"] = bson.M{"$regex": o.EmailHashLike, "$options": "i"}
		and = append(and, query)
	}
	if o.EmailHashNlike != nil {
		query := bson.M{}
		query["emailHash"] = bson.M{"$not": bson.M{"$regex": o.EmailHashNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.ExpiresAtEq != nil {
		query := bson.M{}
		query["expiresAt"] = o.ExpiresAtEq
		and = append(and, query)
	}
	if o.ExpiresAtNe != nil {
		query := bson.M{}
		query["expiresAt"] = bson.M{"$ne": o.ExpiresAtNe}
		and = append(and, query)
	}
	if o.ExpiresAtGt != nil {
		query := bson.M{}
		query["expiresAt"] = bson.M{"$gt": o.ExpiresAtGt}
		and = append(and, query)
	}
	if o.ExpiresAtGte != nil {
		query := bson.M{}
		query["expiresAt"] = bson.M{"$gte": o.ExpiresAtGte}
		and = append(and, query)
	}
	if o.ExpiresAtLt != nil {
		query := bson.M{}
		query["expiresAt"] = bson.M{"$lt": o.ExpiresAtLt}
		and = append(and, query)
	}
	if o.ExpiresAtLte != nil {
		query := bson.M{}
		query["expiresAt"] = bson.M{"$lte": o.ExpiresAtLte}
		and = append(and, query)
	}
	if o.ExpiresAtIn != nil {
		query := bson.M{}
		query["expiresAt"] = bson.M{"$in": o.ExpiresAtIn}
		and = append(and, query)
	}
	if o.ExpiresAtNin != nil {
		query := bson.M{}
		query["expiresAt"] = bson.M{"$nin": o.ExpiresAtNin}
		and = append(and, query)
	}
	if o.ExpiresAtExists != nil {
		query := bson.M{}
		query["expiresAt"] = bson.M{"$exists": *o.ExpiresAtExists}
		and = append(and, query)
	}
	if o.OwnerIdEq != nil {
		query := bson.M{}
		query["ownerId"] = o.OwnerIdEq
		and = append(and, query)
	}
	if o.OwnerIdIn != nil {
		query := bson.M{}
		query["ownerId"] = bson.M{"$in": o.OwnerIdIn}
		and = append(and, query)
	}
	if o.OwnerIdNin != nil {
		query := bson.M{}
		query["ownerId"] = bson.M{"$nin": o.OwnerIdNin}
		and = append(and, query)
	}
	if o.OwnerIdExists != nil {
		query := bson.M{}
		query["ownerId"] = bson.M{"$exists": *o.OwnerIdExists}
		and = append(and, query)
	}
	if o.PolicyVersionAtIssueEq != nil {
		query := bson.M{}
		query["policyVersionAtIssue"] = o.PolicyVersionAtIssueEq
		and = append(and, query)
	}
	if o.PolicyVersionAtIssueNe != nil {
		query := bson.M{}
		query["policyVersionAtIssue"] = bson.M{"$ne": o.PolicyVersionAtIssueNe}
		and = append(and, query)
	}
	if o.PolicyVersionAtIssueGt != nil {
		query := bson.M{}
		query["policyVersionAtIssue"] = bson.M{"$gt": o.PolicyVersionAtIssueGt}
		and = append(and, query)
	}
	if o.PolicyVersionAtIssueGte != nil {
		query := bson.M{}
		query["policyVersionAtIssue"] = bson.M{"$gte": o.PolicyVersionAtIssueGte}
		and = append(and, query)
	}
	if o.PolicyVersionAtIssueLt != nil {
		query := bson.M{}
		query["policyVersionAtIssue"] = bson.M{"$lt": o.PolicyVersionAtIssueLt}
		and = append(and, query)
	}
	if o.PolicyVersionAtIssueLte != nil {
		query := bson.M{}
		query["policyVersionAtIssue"] = bson.M{"$lte": o.PolicyVersionAtIssueLte}
		and = append(and, query)
	}
	if o.PolicyVersionAtIssueIn != nil {
		query := bson.M{}
		query["policyVersionAtIssue"] = bson.M{"$in": o.PolicyVersionAtIssueIn}
		and = append(and, query)
	}
	if o.PolicyVersionAtIssueNin != nil {
		query := bson.M{}
		query["policyVersionAtIssue"] = bson.M{"$nin": o.PolicyVersionAtIssueNin}
		and = append(and, query)
	}
	if o.PolicyVersionAtIssueExists != nil {
		query := bson.M{}
		query["policyVersionAtIssue"] = bson.M{"$exists": *o.PolicyVersionAtIssueExists}
		and = append(and, query)
	}
	if o.RegistryIdEq != nil {
		query := bson.M{}
		query["registryId"] = o.RegistryIdEq
		and = append(and, query)
	}
	if o.RegistryIdIn != nil {
		query := bson.M{}
		query["registryId"] = bson.M{"$in": o.RegistryIdIn}
		and = append(and, query)
	}
	if o.RegistryIdNin != nil {
		query := bson.M{}
		query["registryId"] = bson.M{"$nin": o.RegistryIdNin}
		and = append(and, query)
	}
	if o.RegistryIdExists != nil {
		query := bson.M{}
		query["registryId"] = bson.M{"$exists": *o.RegistryIdExists}
		and = append(and, query)
	}
	if o.TokenHashEq != nil {
		query := bson.M{}
		query["tokenHash"] = o.TokenHashEq
		and = append(and, query)
	}
	if o.TokenHashNe != nil {
		query := bson.M{}
		query["tokenHash"] = bson.M{"$ne": o.TokenHashNe}
		and = append(and, query)
	}
	if o.TokenHashGt != nil {
		query := bson.M{}
		query["tokenHash"] = bson.M{"$gt": o.TokenHashGt}
		and = append(and, query)
	}
	if o.TokenHashGte != nil {
		query := bson.M{}
		query["tokenHash"] = bson.M{"$gte": o.TokenHashGte}
		and = append(and, query)
	}
	if o.TokenHashLt != nil {
		query := bson.M{}
		query["tokenHash"] = bson.M{"$lt": o.TokenHashLt}
		and = append(and, query)
	}
	if o.TokenHashLte != nil {
		query := bson.M{}
		query["tokenHash"] = bson.M{"$lte": o.TokenHashLte}
		and = append(and, query)
	}
	if o.TokenHashIn != nil {
		query := bson.M{}
		query["tokenHash"] = bson.M{"$in": o.TokenHashIn}
		and = append(and, query)
	}
	if o.TokenHashNin != nil {
		query := bson.M{}
		query["tokenHash"] = bson.M{"$nin": o.TokenHashNin}
		and = append(and, query)
	}
	if o.TokenHashExists != nil {
		query := bson.M{}
		query["tokenHash"] = bson.M{"$exists": *o.TokenHashExists}
		and = append(and, query)
	}
	if o.TokenHashLike != nil {
		query := bson.M{}
		query["tokenHash"] = bson.M{"$regex": o.TokenHashLike, "$options": "i"}
		and = append(and, query)
	}
	if o.TokenHashNlike != nil {
		query := bson.M{}
		query["tokenHash"] = bson.M{"$not": bson.M{"$regex": o.TokenHashNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.Updated != nil {
		query := bson.M{}
		updatedQuery, err := o.Updated.GetQueryParts()
		if err != nil {
			return nil, err
		}
		for _, part := range updatedQuery {
			partAsBsonM, ok := part.(bson.M)
			if !ok {
				continue
			}
			for k, v := range partAsBsonM {
				query["updated."+k] = v
			}
		}
		and = append(and, query)
	}
	if o.UpdatedByOwnerUser != nil {
		query := bson.M{}
		updatedByOwnerUserQuery, err := o.UpdatedByOwnerUser.GetQueryParts()
		if err != nil {
			return nil, err
		}
		for _, part := range updatedByOwnerUserQuery {
			partAsBsonM, ok := part.(bson.M)
			if !ok {
				continue
			}
			for k, v := range partAsBsonM {
				query["updatedByOwnerUser."+k] = v
			}
		}
		and = append(and, query)
	}
	return and, nil
}

type MongoSortParams struct {
	CreatedAt  int8
	EmailHash  int8
	OwnerId    int8
	RegistryId int8
	TokenHash  int8
	UpdatedAt  int8
}
