package owner_user

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_role"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRecord struct {
	Id                 *primitive.ObjectID       `bson:"_id,omitempty"`
	ActorRoles         *[]actor_role.MongoRecord `bson:"actorRoles,omitempty"`
	Created            *actor_trace.MongoRecord  `bson:"created,omitempty"`
	Email              *string                   `bson:"email,omitempty"`
	Name               *string                   `bson:"name,omitempty"`
	Updated            *actor_trace.MongoRecord  `bson:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.MongoRecord  `bson:"updatedByOwnerUser,omitempty"`
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
	if r.ActorRoles != nil {
		elemactorRoles0 := make([]actor_role.Model, 0)
		for _, ractorRoles0 := range *r.ActorRoles {
			elemactorRoles1, err := ractorRoles0.ToModel()
			if err != nil {
				return m, err
			}
			elemactorRoles0 = append(elemactorRoles0, elemactorRoles1)
		}
		m.ActorRoles = elemactorRoles0
	}
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.Email != nil {
		elememail0 := r.Email
		m.Email = *elememail0
	}
	if r.Name != nil {
		elemname0 := r.Name
		m.Name = *elemname0
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
type MongoSelectByEmailUniqueQuery struct {
	Email string
}

type MongoWhereClause struct {
	// id (Ref<OwnerUser>) search options
	IdEq     *primitive.ObjectID
	IdIn     *[]primitive.ObjectID
	IdNin    *[]primitive.ObjectID
	IdExists *bool
	// actorRoles (List<ActorRole>) search options
	ActorRoles      *actor_role.MongoWhereClause
	ActorRolesEmpty *bool
	// created (ActorTrace) search options
	Created *actor_trace.MongoWhereClause
	// email (string) search options
	EmailEq     *string
	EmailNe     *string
	EmailGt     *string
	EmailGte    *string
	EmailLt     *string
	EmailLte    *string
	EmailIn     *[]string
	EmailNin    *[]string
	EmailExists *bool
	EmailLike   *string
	EmailNlike  *string
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
	if o.ActorRoles != nil {
		query := bson.M{}
		actorRolesQuery, err := o.ActorRoles.GetQueryParts()
		if err != nil {
			return nil, err
		}
		for _, part := range actorRolesQuery {
			partAsBsonM, ok := part.(bson.M)
			if !ok {
				continue
			}
			for k, v := range partAsBsonM {
				query["actorRoles."+k] = v
			}
		}
		and = append(and, query)
	}
	if o.ActorRolesEmpty != nil {
		query := bson.M{}
		if *o.ActorRolesEmpty {
			query["$or"] = bson.A{
				bson.M{"actorRoles": nil},
				bson.M{"actorRoles": bson.A{}},
				bson.M{"actorRoles": bson.M{"$exists": false}},
			}
		} else {
			query["actorRoles"] = bson.M{
				"$ne":     nil,
				"$not":    bson.M{"$size": 0},
				"$exists": true,
			}
		}
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
	if o.EmailEq != nil {
		query := bson.M{}
		query["email"] = o.EmailEq
		and = append(and, query)
	}
	if o.EmailNe != nil {
		query := bson.M{}
		query["email"] = bson.M{"$ne": o.EmailNe}
		and = append(and, query)
	}
	if o.EmailGt != nil {
		query := bson.M{}
		query["email"] = bson.M{"$gt": o.EmailGt}
		and = append(and, query)
	}
	if o.EmailGte != nil {
		query := bson.M{}
		query["email"] = bson.M{"$gte": o.EmailGte}
		and = append(and, query)
	}
	if o.EmailLt != nil {
		query := bson.M{}
		query["email"] = bson.M{"$lt": o.EmailLt}
		and = append(and, query)
	}
	if o.EmailLte != nil {
		query := bson.M{}
		query["email"] = bson.M{"$lte": o.EmailLte}
		and = append(and, query)
	}
	if o.EmailIn != nil {
		query := bson.M{}
		query["email"] = bson.M{"$in": o.EmailIn}
		and = append(and, query)
	}
	if o.EmailNin != nil {
		query := bson.M{}
		query["email"] = bson.M{"$nin": o.EmailNin}
		and = append(and, query)
	}
	if o.EmailExists != nil {
		query := bson.M{}
		query["email"] = bson.M{"$exists": *o.EmailExists}
		and = append(and, query)
	}
	if o.EmailLike != nil {
		query := bson.M{}
		query["email"] = bson.M{"$regex": o.EmailLike, "$options": "i"}
		and = append(and, query)
	}
	if o.EmailNlike != nil {
		query := bson.M{}
		query["email"] = bson.M{"$not": bson.M{"$regex": o.EmailNlike, "$options": "i"}}
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
	CreatedAt int8
	Email     int8
	UpdatedAt int8
}
