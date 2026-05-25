package owner_user

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_role"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id                 string
	ActorRoles         []actor_role.Model
	Created            actor_trace.Model
	Email              string
	Name               string
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
	if projection.ActorRoles {
		elemactorRoles0 := make([]actor_role.MongoRecord, 0)
		for _, mactorRoles0 := range m.ActorRoles {
			elemactorRoles1, err := mactorRoles0.ToMongoRecord(projection.ActorRolesFields)
			if err != nil {
				return r, err
			}
			elemactorRoles0 = append(elemactorRoles0, elemactorRoles1)
		}
		r.ActorRoles = &elemactorRoles0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToMongoRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.Email {
		elememail0 := m.Email
		r.Email = &elememail0
	}
	if projection.Name {
		elemname0 := m.Name
		r.Name = &elemname0
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
	if projection.ActorRoles {
		elemactorRoles0 := make([]actor_role.HTTPRecord, 0)
		for _, mactorRoles0 := range m.ActorRoles {
			elemactorRoles1, err := mactorRoles0.ToHTTPRecord(projection.ActorRolesFields)
			if err != nil {
				return r, err
			}
			elemactorRoles0 = append(elemactorRoles0, elemactorRoles1)
		}
		r.ActorRoles = &elemactorRoles0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToHTTPRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.Email {
		elememail0 := m.Email
		r.Email = &elememail0
	}
	if projection.Name {
		elemname0 := m.Name
		r.Name = &elemname0
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
type SelectByEmailUniqueQuery struct {
	Email string
}

type WhereClause struct {
	// id (Ref<OwnerUser>) search options
	IdEq     *string
	IdIn     *[]string
	IdNin    *[]string
	IdExists *bool
	// actorRoles (List<ActorRole>) search options
	ActorRoles      *actor_role.WhereClause
	ActorRolesEmpty *bool
	// created (ActorTrace) search options
	Created *actor_trace.WhereClause
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
func (o SelectByEmailUniqueQuery) ToMongoSelectByEmailUniqueQuery() (MongoSelectByEmailUniqueQuery, error) {
	to := MongoSelectByEmailUniqueQuery{}
	elememail0 := o.Email
	to.Email = elememail0
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
	if o.ActorRoles != nil {
		elemactorRoles0, err := o.ActorRoles.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.ActorRoles = &elemactorRoles0
	}
	if o.ActorRolesEmpty != nil {
		elemactorRolesEmpty0 := o.ActorRolesEmpty
		to.ActorRolesEmpty = elemactorRolesEmpty0
	}
	if o.Created != nil {
		elemcreated0, err := o.Created.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.Created = &elemcreated0
	}
	if o.EmailEq != nil {
		elememailEq0 := o.EmailEq
		to.EmailEq = elememailEq0
	}
	if o.EmailNe != nil {
		elememailNe0 := o.EmailNe
		to.EmailNe = elememailNe0
	}
	if o.EmailGt != nil {
		elememailGt0 := o.EmailGt
		to.EmailGt = elememailGt0
	}
	if o.EmailGte != nil {
		elememailGte0 := o.EmailGte
		to.EmailGte = elememailGte0
	}
	if o.EmailLt != nil {
		elememailLt0 := o.EmailLt
		to.EmailLt = elememailLt0
	}
	if o.EmailLte != nil {
		elememailLte0 := o.EmailLte
		to.EmailLte = elememailLte0
	}
	if o.EmailIn != nil {
		elememailIn0 := make([]string, 0)
		for _, oemailIn0 := range *o.EmailIn {
			elememailIn1 := oemailIn0
			elememailIn0 = append(elememailIn0, elememailIn1)
		}
		to.EmailIn = &elememailIn0
	}
	if o.EmailNin != nil {
		elememailNin0 := make([]string, 0)
		for _, oemailNin0 := range *o.EmailNin {
			elememailNin1 := oemailNin0
			elememailNin0 = append(elememailNin0, elememailNin1)
		}
		to.EmailNin = &elememailNin0
	}
	if o.EmailExists != nil {
		elememailExists0 := o.EmailExists
		to.EmailExists = elememailExists0
	}
	if o.EmailLike != nil {
		elememailLike0 := o.EmailLike
		to.EmailLike = elememailLike0
	}
	if o.EmailNlike != nil {
		elememailNlike0 := o.EmailNlike
		to.EmailNlike = elememailNlike0
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
	CreatedAt int8
	Email     int8
	UpdatedAt int8
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	to.CreatedAt = s.CreatedAt
	to.Email = s.Email
	to.UpdatedAt = s.UpdatedAt
	return to
}
