package reservation

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_reservation_status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRecord struct {
	Id                 *primitive.ObjectID            `bson:"_id,omitempty"`
	ContactEmail       *string                        `bson:"contactEmail,omitempty"`
	Created            *actor_trace.MongoRecord       `bson:"created,omitempty"`
	IsAnonymous        *bool                          `bson:"isAnonymous,omitempty"`
	ItemId             *primitive.ObjectID            `bson:"itemId,omitempty"`
	Message            *string                        `bson:"message,omitempty"`
	Quantity           *int                           `bson:"quantity,omitempty"`
	RegistryId         *primitive.ObjectID            `bson:"registryId,omitempty"`
	ReserverName       *string                        `bson:"reserverName,omitempty"`
	Status             *enum_reservation_status.Value `bson:"status,omitempty"`
	Updated            *actor_trace.MongoRecord       `bson:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.MongoRecord       `bson:"updatedByOwnerUser,omitempty"`
}

type MongoUpdateWhereClause struct {
	Id primitive.ObjectID
}

func (r *MongoRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id.Hex()
		m.Id = elemid0
	}
	if r.ContactEmail != nil {
		elemcontactEmail0 := r.ContactEmail
		m.ContactEmail = *elemcontactEmail0
	}
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.IsAnonymous != nil {
		elemisAnonymous0 := r.IsAnonymous
		m.IsAnonymous = *elemisAnonymous0
	}
	if r.ItemId != nil {
		elemitemId0 := r.ItemId.Hex()
		m.ItemId = elemitemId0
	}
	if r.Message != nil {
		elemmessage0 := r.Message
		m.Message = *elemmessage0
	}
	if r.Quantity != nil {
		elemquantity0 := r.Quantity
		m.Quantity = *elemquantity0
	}
	if r.RegistryId != nil {
		elemregistryId0 := r.RegistryId.Hex()
		m.RegistryId = elemregistryId0
	}
	if r.ReserverName != nil {
		elemreserverName0 := r.ReserverName
		m.ReserverName = *elemreserverName0
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
	// id (Ref<Reservation>) search options
	IdEq     *primitive.ObjectID
	IdIn     *[]primitive.ObjectID
	IdNin    *[]primitive.ObjectID
	IdExists *bool
	// contactEmail (string) search options
	ContactEmailEq     *string
	ContactEmailNe     *string
	ContactEmailGt     *string
	ContactEmailGte    *string
	ContactEmailLt     *string
	ContactEmailLte    *string
	ContactEmailIn     *[]string
	ContactEmailNin    *[]string
	ContactEmailExists *bool
	ContactEmailLike   *string
	ContactEmailNlike  *string
	// created (ActorTrace) search options
	Created *actor_trace.MongoWhereClause
	// isAnonymous (bool) search options
	IsAnonymousEq     *bool
	IsAnonymousNe     *bool
	IsAnonymousGt     *bool
	IsAnonymousGte    *bool
	IsAnonymousLt     *bool
	IsAnonymousLte    *bool
	IsAnonymousIn     *[]bool
	IsAnonymousNin    *[]bool
	IsAnonymousExists *bool
	// itemId (ParentRef<RegistryItem>) search options
	ItemIdEq     *primitive.ObjectID
	ItemIdIn     *[]primitive.ObjectID
	ItemIdNin    *[]primitive.ObjectID
	ItemIdExists *bool
	// message (string) search options
	MessageEq     *string
	MessageNe     *string
	MessageGt     *string
	MessageGte    *string
	MessageLt     *string
	MessageLte    *string
	MessageIn     *[]string
	MessageNin    *[]string
	MessageExists *bool
	MessageLike   *string
	MessageNlike  *string
	// quantity (int) search options
	QuantityEq     *int
	QuantityNe     *int
	QuantityGt     *int
	QuantityGte    *int
	QuantityLt     *int
	QuantityLte    *int
	QuantityIn     *[]int
	QuantityNin    *[]int
	QuantityExists *bool
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *primitive.ObjectID
	RegistryIdIn     *[]primitive.ObjectID
	RegistryIdNin    *[]primitive.ObjectID
	RegistryIdExists *bool
	// reserverName (string) search options
	ReserverNameEq     *string
	ReserverNameNe     *string
	ReserverNameGt     *string
	ReserverNameGte    *string
	ReserverNameLt     *string
	ReserverNameLte    *string
	ReserverNameIn     *[]string
	ReserverNameNin    *[]string
	ReserverNameExists *bool
	ReserverNameLike   *string
	ReserverNameNlike  *string
	// status (ReservationStatus) search options
	StatusEq     *enum_reservation_status.Value
	StatusNe     *enum_reservation_status.Value
	StatusGt     *enum_reservation_status.Value
	StatusGte    *enum_reservation_status.Value
	StatusLt     *enum_reservation_status.Value
	StatusLte    *enum_reservation_status.Value
	StatusIn     *[]enum_reservation_status.Value
	StatusNin    *[]enum_reservation_status.Value
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
	if o.ContactEmailEq != nil {
		query := bson.M{}
		query["contactEmail"] = o.ContactEmailEq
		and = append(and, query)
	}
	if o.ContactEmailNe != nil {
		query := bson.M{}
		query["contactEmail"] = bson.M{"$ne": o.ContactEmailNe}
		and = append(and, query)
	}
	if o.ContactEmailGt != nil {
		query := bson.M{}
		query["contactEmail"] = bson.M{"$gt": o.ContactEmailGt}
		and = append(and, query)
	}
	if o.ContactEmailGte != nil {
		query := bson.M{}
		query["contactEmail"] = bson.M{"$gte": o.ContactEmailGte}
		and = append(and, query)
	}
	if o.ContactEmailLt != nil {
		query := bson.M{}
		query["contactEmail"] = bson.M{"$lt": o.ContactEmailLt}
		and = append(and, query)
	}
	if o.ContactEmailLte != nil {
		query := bson.M{}
		query["contactEmail"] = bson.M{"$lte": o.ContactEmailLte}
		and = append(and, query)
	}
	if o.ContactEmailIn != nil {
		query := bson.M{}
		query["contactEmail"] = bson.M{"$in": o.ContactEmailIn}
		and = append(and, query)
	}
	if o.ContactEmailNin != nil {
		query := bson.M{}
		query["contactEmail"] = bson.M{"$nin": o.ContactEmailNin}
		and = append(and, query)
	}
	if o.ContactEmailExists != nil {
		query := bson.M{}
		query["contactEmail"] = bson.M{"$exists": *o.ContactEmailExists}
		and = append(and, query)
	}
	if o.ContactEmailLike != nil {
		query := bson.M{}
		query["contactEmail"] = bson.M{"$regex": o.ContactEmailLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ContactEmailNlike != nil {
		query := bson.M{}
		query["contactEmail"] = bson.M{"$not": bson.M{"$regex": o.ContactEmailNlike, "$options": "i"}}
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
	if o.IsAnonymousEq != nil {
		query := bson.M{}
		query["isAnonymous"] = o.IsAnonymousEq
		and = append(and, query)
	}
	if o.IsAnonymousNe != nil {
		query := bson.M{}
		query["isAnonymous"] = bson.M{"$ne": o.IsAnonymousNe}
		and = append(and, query)
	}
	if o.IsAnonymousGt != nil {
		query := bson.M{}
		query["isAnonymous"] = bson.M{"$gt": o.IsAnonymousGt}
		and = append(and, query)
	}
	if o.IsAnonymousGte != nil {
		query := bson.M{}
		query["isAnonymous"] = bson.M{"$gte": o.IsAnonymousGte}
		and = append(and, query)
	}
	if o.IsAnonymousLt != nil {
		query := bson.M{}
		query["isAnonymous"] = bson.M{"$lt": o.IsAnonymousLt}
		and = append(and, query)
	}
	if o.IsAnonymousLte != nil {
		query := bson.M{}
		query["isAnonymous"] = bson.M{"$lte": o.IsAnonymousLte}
		and = append(and, query)
	}
	if o.IsAnonymousIn != nil {
		query := bson.M{}
		query["isAnonymous"] = bson.M{"$in": o.IsAnonymousIn}
		and = append(and, query)
	}
	if o.IsAnonymousNin != nil {
		query := bson.M{}
		query["isAnonymous"] = bson.M{"$nin": o.IsAnonymousNin}
		and = append(and, query)
	}
	if o.IsAnonymousExists != nil {
		query := bson.M{}
		query["isAnonymous"] = bson.M{"$exists": *o.IsAnonymousExists}
		and = append(and, query)
	}
	if o.ItemIdEq != nil {
		query := bson.M{}
		query["itemId"] = o.ItemIdEq
		and = append(and, query)
	}
	if o.ItemIdIn != nil {
		query := bson.M{}
		query["itemId"] = bson.M{"$in": o.ItemIdIn}
		and = append(and, query)
	}
	if o.ItemIdNin != nil {
		query := bson.M{}
		query["itemId"] = bson.M{"$nin": o.ItemIdNin}
		and = append(and, query)
	}
	if o.ItemIdExists != nil {
		query := bson.M{}
		query["itemId"] = bson.M{"$exists": *o.ItemIdExists}
		and = append(and, query)
	}
	if o.MessageEq != nil {
		query := bson.M{}
		query["message"] = o.MessageEq
		and = append(and, query)
	}
	if o.MessageNe != nil {
		query := bson.M{}
		query["message"] = bson.M{"$ne": o.MessageNe}
		and = append(and, query)
	}
	if o.MessageGt != nil {
		query := bson.M{}
		query["message"] = bson.M{"$gt": o.MessageGt}
		and = append(and, query)
	}
	if o.MessageGte != nil {
		query := bson.M{}
		query["message"] = bson.M{"$gte": o.MessageGte}
		and = append(and, query)
	}
	if o.MessageLt != nil {
		query := bson.M{}
		query["message"] = bson.M{"$lt": o.MessageLt}
		and = append(and, query)
	}
	if o.MessageLte != nil {
		query := bson.M{}
		query["message"] = bson.M{"$lte": o.MessageLte}
		and = append(and, query)
	}
	if o.MessageIn != nil {
		query := bson.M{}
		query["message"] = bson.M{"$in": o.MessageIn}
		and = append(and, query)
	}
	if o.MessageNin != nil {
		query := bson.M{}
		query["message"] = bson.M{"$nin": o.MessageNin}
		and = append(and, query)
	}
	if o.MessageExists != nil {
		query := bson.M{}
		query["message"] = bson.M{"$exists": *o.MessageExists}
		and = append(and, query)
	}
	if o.MessageLike != nil {
		query := bson.M{}
		query["message"] = bson.M{"$regex": o.MessageLike, "$options": "i"}
		and = append(and, query)
	}
	if o.MessageNlike != nil {
		query := bson.M{}
		query["message"] = bson.M{"$not": bson.M{"$regex": o.MessageNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.QuantityEq != nil {
		query := bson.M{}
		query["quantity"] = o.QuantityEq
		and = append(and, query)
	}
	if o.QuantityNe != nil {
		query := bson.M{}
		query["quantity"] = bson.M{"$ne": o.QuantityNe}
		and = append(and, query)
	}
	if o.QuantityGt != nil {
		query := bson.M{}
		query["quantity"] = bson.M{"$gt": o.QuantityGt}
		and = append(and, query)
	}
	if o.QuantityGte != nil {
		query := bson.M{}
		query["quantity"] = bson.M{"$gte": o.QuantityGte}
		and = append(and, query)
	}
	if o.QuantityLt != nil {
		query := bson.M{}
		query["quantity"] = bson.M{"$lt": o.QuantityLt}
		and = append(and, query)
	}
	if o.QuantityLte != nil {
		query := bson.M{}
		query["quantity"] = bson.M{"$lte": o.QuantityLte}
		and = append(and, query)
	}
	if o.QuantityIn != nil {
		query := bson.M{}
		query["quantity"] = bson.M{"$in": o.QuantityIn}
		and = append(and, query)
	}
	if o.QuantityNin != nil {
		query := bson.M{}
		query["quantity"] = bson.M{"$nin": o.QuantityNin}
		and = append(and, query)
	}
	if o.QuantityExists != nil {
		query := bson.M{}
		query["quantity"] = bson.M{"$exists": *o.QuantityExists}
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
	if o.ReserverNameEq != nil {
		query := bson.M{}
		query["reserverName"] = o.ReserverNameEq
		and = append(and, query)
	}
	if o.ReserverNameNe != nil {
		query := bson.M{}
		query["reserverName"] = bson.M{"$ne": o.ReserverNameNe}
		and = append(and, query)
	}
	if o.ReserverNameGt != nil {
		query := bson.M{}
		query["reserverName"] = bson.M{"$gt": o.ReserverNameGt}
		and = append(and, query)
	}
	if o.ReserverNameGte != nil {
		query := bson.M{}
		query["reserverName"] = bson.M{"$gte": o.ReserverNameGte}
		and = append(and, query)
	}
	if o.ReserverNameLt != nil {
		query := bson.M{}
		query["reserverName"] = bson.M{"$lt": o.ReserverNameLt}
		and = append(and, query)
	}
	if o.ReserverNameLte != nil {
		query := bson.M{}
		query["reserverName"] = bson.M{"$lte": o.ReserverNameLte}
		and = append(and, query)
	}
	if o.ReserverNameIn != nil {
		query := bson.M{}
		query["reserverName"] = bson.M{"$in": o.ReserverNameIn}
		and = append(and, query)
	}
	if o.ReserverNameNin != nil {
		query := bson.M{}
		query["reserverName"] = bson.M{"$nin": o.ReserverNameNin}
		and = append(and, query)
	}
	if o.ReserverNameExists != nil {
		query := bson.M{}
		query["reserverName"] = bson.M{"$exists": *o.ReserverNameExists}
		and = append(and, query)
	}
	if o.ReserverNameLike != nil {
		query := bson.M{}
		query["reserverName"] = bson.M{"$regex": o.ReserverNameLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ReserverNameNlike != nil {
		query := bson.M{}
		query["reserverName"] = bson.M{"$not": bson.M{"$regex": o.ReserverNameNlike, "$options": "i"}}
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
	ItemId     int8
	RegistryId int8
	UpdatedAt  int8
}
