package shipping_address_request

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_address_request_status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRecord struct {
	Id                 *primitive.ObjectID                `bson:"_id,omitempty"`
	Created            *actor_trace.MongoRecord           `bson:"created,omitempty"`
	DecisionReason     *string                            `bson:"decisionReason,omitempty"`
	EmailEnc           *string                            `bson:"emailEnc,omitempty"`
	EmailHash          *string                            `bson:"emailHash,omitempty"`
	Name               *string                            `bson:"name,omitempty"`
	Note               *string                            `bson:"note,omitempty"`
	OwnerId            *primitive.ObjectID                `bson:"ownerId,omitempty"`
	PolicyVersion      *int                               `bson:"policyVersion,omitempty"`
	RegistryId         *primitive.ObjectID                `bson:"registryId,omitempty"`
	RegistryItemId     *primitive.ObjectID                `bson:"registryItemId,omitempty"`
	Status             *enum_address_request_status.Value `bson:"status,omitempty"`
	Updated            *actor_trace.MongoRecord           `bson:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.MongoRecord           `bson:"updatedByOwnerUser,omitempty"`
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
		elemownerId0 := r.OwnerId.Hex()
		m.OwnerId = elemownerId0
	}
	if r.PolicyVersion != nil {
		elempolicyVersion0 := r.PolicyVersion
		m.PolicyVersion = *elempolicyVersion0
	}
	if r.RegistryId != nil {
		elemregistryId0 := r.RegistryId.Hex()
		m.RegistryId = elemregistryId0
	}
	if r.RegistryItemId != nil {
		elemregistryItemId0 := r.RegistryItemId.Hex()
		m.RegistryItemId = elemregistryItemId0
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

type MongoWhereClause struct {
	// id (Ref<ShippingAddressRequest>) search options
	IdEq     *primitive.ObjectID
	IdIn     *[]primitive.ObjectID
	IdNin    *[]primitive.ObjectID
	IdExists *bool
	// created (ActorTrace) search options
	Created *actor_trace.MongoWhereClause
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
	OwnerIdEq     *primitive.ObjectID
	OwnerIdIn     *[]primitive.ObjectID
	OwnerIdNin    *[]primitive.ObjectID
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
	RegistryIdEq     *primitive.ObjectID
	RegistryIdIn     *[]primitive.ObjectID
	RegistryIdNin    *[]primitive.ObjectID
	RegistryIdExists *bool
	// registryItemId (Ref<RegistryItem>) search options
	RegistryItemIdEq     *primitive.ObjectID
	RegistryItemIdIn     *[]primitive.ObjectID
	RegistryItemIdNin    *[]primitive.ObjectID
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
	if o.DecisionReasonEq != nil {
		query := bson.M{}
		query["decisionReason"] = o.DecisionReasonEq
		and = append(and, query)
	}
	if o.DecisionReasonNe != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$ne": o.DecisionReasonNe}
		and = append(and, query)
	}
	if o.DecisionReasonGt != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$gt": o.DecisionReasonGt}
		and = append(and, query)
	}
	if o.DecisionReasonGte != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$gte": o.DecisionReasonGte}
		and = append(and, query)
	}
	if o.DecisionReasonLt != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$lt": o.DecisionReasonLt}
		and = append(and, query)
	}
	if o.DecisionReasonLte != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$lte": o.DecisionReasonLte}
		and = append(and, query)
	}
	if o.DecisionReasonIn != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$in": o.DecisionReasonIn}
		and = append(and, query)
	}
	if o.DecisionReasonNin != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$nin": o.DecisionReasonNin}
		and = append(and, query)
	}
	if o.DecisionReasonExists != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$exists": *o.DecisionReasonExists}
		and = append(and, query)
	}
	if o.DecisionReasonLike != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$regex": o.DecisionReasonLike, "$options": "i"}
		and = append(and, query)
	}
	if o.DecisionReasonNlike != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$not": bson.M{"$regex": o.DecisionReasonNlike, "$options": "i"}}
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
	if o.NoteEq != nil {
		query := bson.M{}
		query["note"] = o.NoteEq
		and = append(and, query)
	}
	if o.NoteNe != nil {
		query := bson.M{}
		query["note"] = bson.M{"$ne": o.NoteNe}
		and = append(and, query)
	}
	if o.NoteGt != nil {
		query := bson.M{}
		query["note"] = bson.M{"$gt": o.NoteGt}
		and = append(and, query)
	}
	if o.NoteGte != nil {
		query := bson.M{}
		query["note"] = bson.M{"$gte": o.NoteGte}
		and = append(and, query)
	}
	if o.NoteLt != nil {
		query := bson.M{}
		query["note"] = bson.M{"$lt": o.NoteLt}
		and = append(and, query)
	}
	if o.NoteLte != nil {
		query := bson.M{}
		query["note"] = bson.M{"$lte": o.NoteLte}
		and = append(and, query)
	}
	if o.NoteIn != nil {
		query := bson.M{}
		query["note"] = bson.M{"$in": o.NoteIn}
		and = append(and, query)
	}
	if o.NoteNin != nil {
		query := bson.M{}
		query["note"] = bson.M{"$nin": o.NoteNin}
		and = append(and, query)
	}
	if o.NoteExists != nil {
		query := bson.M{}
		query["note"] = bson.M{"$exists": *o.NoteExists}
		and = append(and, query)
	}
	if o.NoteLike != nil {
		query := bson.M{}
		query["note"] = bson.M{"$regex": o.NoteLike, "$options": "i"}
		and = append(and, query)
	}
	if o.NoteNlike != nil {
		query := bson.M{}
		query["note"] = bson.M{"$not": bson.M{"$regex": o.NoteNlike, "$options": "i"}}
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
	if o.PolicyVersionEq != nil {
		query := bson.M{}
		query["policyVersion"] = o.PolicyVersionEq
		and = append(and, query)
	}
	if o.PolicyVersionNe != nil {
		query := bson.M{}
		query["policyVersion"] = bson.M{"$ne": o.PolicyVersionNe}
		and = append(and, query)
	}
	if o.PolicyVersionGt != nil {
		query := bson.M{}
		query["policyVersion"] = bson.M{"$gt": o.PolicyVersionGt}
		and = append(and, query)
	}
	if o.PolicyVersionGte != nil {
		query := bson.M{}
		query["policyVersion"] = bson.M{"$gte": o.PolicyVersionGte}
		and = append(and, query)
	}
	if o.PolicyVersionLt != nil {
		query := bson.M{}
		query["policyVersion"] = bson.M{"$lt": o.PolicyVersionLt}
		and = append(and, query)
	}
	if o.PolicyVersionLte != nil {
		query := bson.M{}
		query["policyVersion"] = bson.M{"$lte": o.PolicyVersionLte}
		and = append(and, query)
	}
	if o.PolicyVersionIn != nil {
		query := bson.M{}
		query["policyVersion"] = bson.M{"$in": o.PolicyVersionIn}
		and = append(and, query)
	}
	if o.PolicyVersionNin != nil {
		query := bson.M{}
		query["policyVersion"] = bson.M{"$nin": o.PolicyVersionNin}
		and = append(and, query)
	}
	if o.PolicyVersionExists != nil {
		query := bson.M{}
		query["policyVersion"] = bson.M{"$exists": *o.PolicyVersionExists}
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
	if o.RegistryItemIdEq != nil {
		query := bson.M{}
		query["registryItemId"] = o.RegistryItemIdEq
		and = append(and, query)
	}
	if o.RegistryItemIdIn != nil {
		query := bson.M{}
		query["registryItemId"] = bson.M{"$in": o.RegistryItemIdIn}
		and = append(and, query)
	}
	if o.RegistryItemIdNin != nil {
		query := bson.M{}
		query["registryItemId"] = bson.M{"$nin": o.RegistryItemIdNin}
		and = append(and, query)
	}
	if o.RegistryItemIdExists != nil {
		query := bson.M{}
		query["registryItemId"] = bson.M{"$exists": *o.RegistryItemIdExists}
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
	Status     int8
	UpdatedAt  int8
}
