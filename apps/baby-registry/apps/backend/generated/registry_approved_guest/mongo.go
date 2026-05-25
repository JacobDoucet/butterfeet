package registry_approved_guest

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_guest_access_level"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_guest_status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRecord struct {
	Id                 *primitive.ObjectID            `bson:"_id,omitempty"`
	AccessLevel        *enum_guest_access_level.Value `bson:"accessLevel,omitempty"`
	Created            *actor_trace.MongoRecord       `bson:"created,omitempty"`
	EmailEnc           *string                        `bson:"emailEnc,omitempty"`
	EmailHash          *string                        `bson:"emailHash,omitempty"`
	Name               *string                        `bson:"name,omitempty"`
	OwnerId            *primitive.ObjectID            `bson:"ownerId,omitempty"`
	RegistryId         *primitive.ObjectID            `bson:"registryId,omitempty"`
	Status             *enum_guest_status.Value       `bson:"status,omitempty"`
	Updated            *actor_trace.MongoRecord       `bson:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.MongoRecord       `bson:"updatedByOwnerUser,omitempty"`
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
		elemownerId0 := r.OwnerId.Hex()
		m.OwnerId = elemownerId0
	}
	if r.RegistryId != nil {
		elemregistryId0 := r.RegistryId.Hex()
		m.RegistryId = elemregistryId0
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

type MongoSelectByIdQuery struct {
	Id primitive.ObjectID
}
type MongoSelectByRegistryEmailUniqueQuery struct {
	RegistryId primitive.ObjectID
	EmailHash  string
}

type MongoWhereClause struct {
	// id (Ref<RegistryApprovedGuest>) search options
	IdEq     *primitive.ObjectID
	IdIn     *[]primitive.ObjectID
	IdNin    *[]primitive.ObjectID
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
	Created *actor_trace.MongoWhereClause
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
	OwnerIdEq     *primitive.ObjectID
	OwnerIdIn     *[]primitive.ObjectID
	OwnerIdNin    *[]primitive.ObjectID
	OwnerIdExists *bool
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *primitive.ObjectID
	RegistryIdIn     *[]primitive.ObjectID
	RegistryIdNin    *[]primitive.ObjectID
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
	if o.AccessLevelEq != nil {
		query := bson.M{}
		query["accessLevel"] = o.AccessLevelEq
		and = append(and, query)
	}
	if o.AccessLevelNe != nil {
		query := bson.M{}
		query["accessLevel"] = bson.M{"$ne": o.AccessLevelNe}
		and = append(and, query)
	}
	if o.AccessLevelGt != nil {
		query := bson.M{}
		query["accessLevel"] = bson.M{"$gt": o.AccessLevelGt}
		and = append(and, query)
	}
	if o.AccessLevelGte != nil {
		query := bson.M{}
		query["accessLevel"] = bson.M{"$gte": o.AccessLevelGte}
		and = append(and, query)
	}
	if o.AccessLevelLt != nil {
		query := bson.M{}
		query["accessLevel"] = bson.M{"$lt": o.AccessLevelLt}
		and = append(and, query)
	}
	if o.AccessLevelLte != nil {
		query := bson.M{}
		query["accessLevel"] = bson.M{"$lte": o.AccessLevelLte}
		and = append(and, query)
	}
	if o.AccessLevelIn != nil {
		query := bson.M{}
		query["accessLevel"] = bson.M{"$in": o.AccessLevelIn}
		and = append(and, query)
	}
	if o.AccessLevelNin != nil {
		query := bson.M{}
		query["accessLevel"] = bson.M{"$nin": o.AccessLevelNin}
		and = append(and, query)
	}
	if o.AccessLevelExists != nil {
		query := bson.M{}
		query["accessLevel"] = bson.M{"$exists": *o.AccessLevelExists}
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
	if o.EmailEncEq != nil {
		query := bson.M{}
		query["emailEnc"] = o.EmailEncEq
		and = append(and, query)
	}
	if o.EmailEncNe != nil {
		query := bson.M{}
		query["emailEnc"] = bson.M{"$ne": o.EmailEncNe}
		and = append(and, query)
	}
	if o.EmailEncGt != nil {
		query := bson.M{}
		query["emailEnc"] = bson.M{"$gt": o.EmailEncGt}
		and = append(and, query)
	}
	if o.EmailEncGte != nil {
		query := bson.M{}
		query["emailEnc"] = bson.M{"$gte": o.EmailEncGte}
		and = append(and, query)
	}
	if o.EmailEncLt != nil {
		query := bson.M{}
		query["emailEnc"] = bson.M{"$lt": o.EmailEncLt}
		and = append(and, query)
	}
	if o.EmailEncLte != nil {
		query := bson.M{}
		query["emailEnc"] = bson.M{"$lte": o.EmailEncLte}
		and = append(and, query)
	}
	if o.EmailEncIn != nil {
		query := bson.M{}
		query["emailEnc"] = bson.M{"$in": o.EmailEncIn}
		and = append(and, query)
	}
	if o.EmailEncNin != nil {
		query := bson.M{}
		query["emailEnc"] = bson.M{"$nin": o.EmailEncNin}
		and = append(and, query)
	}
	if o.EmailEncExists != nil {
		query := bson.M{}
		query["emailEnc"] = bson.M{"$exists": *o.EmailEncExists}
		and = append(and, query)
	}
	if o.EmailEncLike != nil {
		query := bson.M{}
		query["emailEnc"] = bson.M{"$regex": o.EmailEncLike, "$options": "i"}
		and = append(and, query)
	}
	if o.EmailEncNlike != nil {
		query := bson.M{}
		query["emailEnc"] = bson.M{"$not": bson.M{"$regex": o.EmailEncNlike, "$options": "i"}}
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
	if o.NameEq != nil {
		query := bson.M{}
		query["name"] = o.NameEq
		and = append(and, query)
	}
	if o.NameNe != nil {
		query := bson.M{}
		query["name"] = bson.M{"$ne": o.NameNe}
		and = append(and, query)
	}
	if o.NameGt != nil {
		query := bson.M{}
		query["name"] = bson.M{"$gt": o.NameGt}
		and = append(and, query)
	}
	if o.NameGte != nil {
		query := bson.M{}
		query["name"] = bson.M{"$gte": o.NameGte}
		and = append(and, query)
	}
	if o.NameLt != nil {
		query := bson.M{}
		query["name"] = bson.M{"$lt": o.NameLt}
		and = append(and, query)
	}
	if o.NameLte != nil {
		query := bson.M{}
		query["name"] = bson.M{"$lte": o.NameLte}
		and = append(and, query)
	}
	if o.NameIn != nil {
		query := bson.M{}
		query["name"] = bson.M{"$in": o.NameIn}
		and = append(and, query)
	}
	if o.NameNin != nil {
		query := bson.M{}
		query["name"] = bson.M{"$nin": o.NameNin}
		and = append(and, query)
	}
	if o.NameExists != nil {
		query := bson.M{}
		query["name"] = bson.M{"$exists": *o.NameExists}
		and = append(and, query)
	}
	if o.NameLike != nil {
		query := bson.M{}
		query["name"] = bson.M{"$regex": o.NameLike, "$options": "i"}
		and = append(and, query)
	}
	if o.NameNlike != nil {
		query := bson.M{}
		query["name"] = bson.M{"$not": bson.M{"$regex": o.NameNlike, "$options": "i"}}
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
	if o.StatusEq != nil {
		query := bson.M{}
		query["status"] = o.StatusEq
		and = append(and, query)
	}
	if o.StatusNe != nil {
		query := bson.M{}
		query["status"] = bson.M{"$ne": o.StatusNe}
		and = append(and, query)
	}
	if o.StatusGt != nil {
		query := bson.M{}
		query["status"] = bson.M{"$gt": o.StatusGt}
		and = append(and, query)
	}
	if o.StatusGte != nil {
		query := bson.M{}
		query["status"] = bson.M{"$gte": o.StatusGte}
		and = append(and, query)
	}
	if o.StatusLt != nil {
		query := bson.M{}
		query["status"] = bson.M{"$lt": o.StatusLt}
		and = append(and, query)
	}
	if o.StatusLte != nil {
		query := bson.M{}
		query["status"] = bson.M{"$lte": o.StatusLte}
		and = append(and, query)
	}
	if o.StatusIn != nil {
		query := bson.M{}
		query["status"] = bson.M{"$in": o.StatusIn}
		and = append(and, query)
	}
	if o.StatusNin != nil {
		query := bson.M{}
		query["status"] = bson.M{"$nin": o.StatusNin}
		and = append(and, query)
	}
	if o.StatusExists != nil {
		query := bson.M{}
		query["status"] = bson.M{"$exists": *o.StatusExists}
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
	UpdatedAt  int8
}
