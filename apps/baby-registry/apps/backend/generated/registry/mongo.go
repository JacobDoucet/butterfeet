package registry

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoRecord struct {
	Id                 *primitive.ObjectID      `bson:"_id,omitempty"`
	CoverImageUrl      *string                  `bson:"coverImageUrl,omitempty"`
	Created            *actor_trace.MongoRecord `bson:"created,omitempty"`
	DueDate            *time.Time               `bson:"dueDate,omitempty"`
	IsPublic           *bool                    `bson:"isPublic,omitempty"`
	OwnerId            *primitive.ObjectID      `bson:"ownerId,omitempty"`
	ParentNames        *string                  `bson:"parentNames,omitempty"`
	Slug               *string                  `bson:"slug,omitempty"`
	ThemeColor         *string                  `bson:"themeColor,omitempty"`
	Title              *string                  `bson:"title,omitempty"`
	Updated            *actor_trace.MongoRecord `bson:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.MongoRecord `bson:"updatedByOwnerUser,omitempty"`
	WelcomeMessage     *string                  `bson:"welcomeMessage,omitempty"`
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
	if r.CoverImageUrl != nil {
		elemcoverImageUrl0 := r.CoverImageUrl
		m.CoverImageUrl = *elemcoverImageUrl0
	}
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.DueDate != nil {
		elemdueDate0 := r.DueDate
		m.DueDate = *elemdueDate0
	}
	if r.IsPublic != nil {
		elemisPublic0 := r.IsPublic
		m.IsPublic = *elemisPublic0
	}
	if r.OwnerId != nil {
		elemownerId0 := r.OwnerId.Hex()
		m.OwnerId = elemownerId0
	}
	if r.ParentNames != nil {
		elemparentNames0 := r.ParentNames
		m.ParentNames = *elemparentNames0
	}
	if r.Slug != nil {
		elemslug0 := r.Slug
		m.Slug = *elemslug0
	}
	if r.ThemeColor != nil {
		elemthemeColor0 := r.ThemeColor
		m.ThemeColor = *elemthemeColor0
	}
	if r.Title != nil {
		elemtitle0 := r.Title
		m.Title = *elemtitle0
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
	if r.WelcomeMessage != nil {
		elemwelcomeMessage0 := r.WelcomeMessage
		m.WelcomeMessage = *elemwelcomeMessage0
	}
	return m, nil
}

type MongoSelectByIdQuery struct {
	Id primitive.ObjectID
}
type MongoSelectBySlugUniqueQuery struct {
	Slug string
}

type MongoWhereClause struct {
	// id (Ref<Registry>) search options
	IdEq     *primitive.ObjectID
	IdIn     *[]primitive.ObjectID
	IdNin    *[]primitive.ObjectID
	IdExists *bool
	// coverImageUrl (string) search options
	CoverImageUrlEq     *string
	CoverImageUrlNe     *string
	CoverImageUrlGt     *string
	CoverImageUrlGte    *string
	CoverImageUrlLt     *string
	CoverImageUrlLte    *string
	CoverImageUrlIn     *[]string
	CoverImageUrlNin    *[]string
	CoverImageUrlExists *bool
	CoverImageUrlLike   *string
	CoverImageUrlNlike  *string
	// created (ActorTrace) search options
	Created *actor_trace.MongoWhereClause
	// dueDate (timestamp) search options
	DueDateEq     *time.Time
	DueDateNe     *time.Time
	DueDateGt     *time.Time
	DueDateGte    *time.Time
	DueDateLt     *time.Time
	DueDateLte    *time.Time
	DueDateIn     *[]time.Time
	DueDateNin    *[]time.Time
	DueDateExists *bool
	// isPublic (bool) search options
	IsPublicEq     *bool
	IsPublicNe     *bool
	IsPublicGt     *bool
	IsPublicGte    *bool
	IsPublicLt     *bool
	IsPublicLte    *bool
	IsPublicIn     *[]bool
	IsPublicNin    *[]bool
	IsPublicExists *bool
	// ownerId (Ref<OwnerUser>) search options
	OwnerIdEq     *primitive.ObjectID
	OwnerIdIn     *[]primitive.ObjectID
	OwnerIdNin    *[]primitive.ObjectID
	OwnerIdExists *bool
	// parentNames (string) search options
	ParentNamesEq     *string
	ParentNamesNe     *string
	ParentNamesGt     *string
	ParentNamesGte    *string
	ParentNamesLt     *string
	ParentNamesLte    *string
	ParentNamesIn     *[]string
	ParentNamesNin    *[]string
	ParentNamesExists *bool
	ParentNamesLike   *string
	ParentNamesNlike  *string
	// slug (string) search options
	SlugEq     *string
	SlugNe     *string
	SlugGt     *string
	SlugGte    *string
	SlugLt     *string
	SlugLte    *string
	SlugIn     *[]string
	SlugNin    *[]string
	SlugExists *bool
	SlugLike   *string
	SlugNlike  *string
	// themeColor (string) search options
	ThemeColorEq     *string
	ThemeColorNe     *string
	ThemeColorGt     *string
	ThemeColorGte    *string
	ThemeColorLt     *string
	ThemeColorLte    *string
	ThemeColorIn     *[]string
	ThemeColorNin    *[]string
	ThemeColorExists *bool
	ThemeColorLike   *string
	ThemeColorNlike  *string
	// title (string) search options
	TitleEq     *string
	TitleNe     *string
	TitleGt     *string
	TitleGte    *string
	TitleLt     *string
	TitleLte    *string
	TitleIn     *[]string
	TitleNin    *[]string
	TitleExists *bool
	TitleLike   *string
	TitleNlike  *string
	// updated (ActorTrace) search options
	Updated *actor_trace.MongoWhereClause
	// updatedByOwnerUser (ActorTrace) search options
	UpdatedByOwnerUser *actor_trace.MongoWhereClause
	// welcomeMessage (string) search options
	WelcomeMessageEq     *string
	WelcomeMessageNe     *string
	WelcomeMessageGt     *string
	WelcomeMessageGte    *string
	WelcomeMessageLt     *string
	WelcomeMessageLte    *string
	WelcomeMessageIn     *[]string
	WelcomeMessageNin    *[]string
	WelcomeMessageExists *bool
	WelcomeMessageLike   *string
	WelcomeMessageNlike  *string
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
	if o.CoverImageUrlEq != nil {
		query := bson.M{}
		query["coverImageUrl"] = o.CoverImageUrlEq
		and = append(and, query)
	}
	if o.CoverImageUrlNe != nil {
		query := bson.M{}
		query["coverImageUrl"] = bson.M{"$ne": o.CoverImageUrlNe}
		and = append(and, query)
	}
	if o.CoverImageUrlGt != nil {
		query := bson.M{}
		query["coverImageUrl"] = bson.M{"$gt": o.CoverImageUrlGt}
		and = append(and, query)
	}
	if o.CoverImageUrlGte != nil {
		query := bson.M{}
		query["coverImageUrl"] = bson.M{"$gte": o.CoverImageUrlGte}
		and = append(and, query)
	}
	if o.CoverImageUrlLt != nil {
		query := bson.M{}
		query["coverImageUrl"] = bson.M{"$lt": o.CoverImageUrlLt}
		and = append(and, query)
	}
	if o.CoverImageUrlLte != nil {
		query := bson.M{}
		query["coverImageUrl"] = bson.M{"$lte": o.CoverImageUrlLte}
		and = append(and, query)
	}
	if o.CoverImageUrlIn != nil {
		query := bson.M{}
		query["coverImageUrl"] = bson.M{"$in": o.CoverImageUrlIn}
		and = append(and, query)
	}
	if o.CoverImageUrlNin != nil {
		query := bson.M{}
		query["coverImageUrl"] = bson.M{"$nin": o.CoverImageUrlNin}
		and = append(and, query)
	}
	if o.CoverImageUrlExists != nil {
		query := bson.M{}
		query["coverImageUrl"] = bson.M{"$exists": *o.CoverImageUrlExists}
		and = append(and, query)
	}
	if o.CoverImageUrlLike != nil {
		query := bson.M{}
		query["coverImageUrl"] = bson.M{"$regex": o.CoverImageUrlLike, "$options": "i"}
		and = append(and, query)
	}
	if o.CoverImageUrlNlike != nil {
		query := bson.M{}
		query["coverImageUrl"] = bson.M{"$not": bson.M{"$regex": o.CoverImageUrlNlike, "$options": "i"}}
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
	if o.DueDateEq != nil {
		query := bson.M{}
		query["dueDate"] = o.DueDateEq
		and = append(and, query)
	}
	if o.DueDateNe != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$ne": o.DueDateNe}
		and = append(and, query)
	}
	if o.DueDateGt != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$gt": o.DueDateGt}
		and = append(and, query)
	}
	if o.DueDateGte != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$gte": o.DueDateGte}
		and = append(and, query)
	}
	if o.DueDateLt != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$lt": o.DueDateLt}
		and = append(and, query)
	}
	if o.DueDateLte != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$lte": o.DueDateLte}
		and = append(and, query)
	}
	if o.DueDateIn != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$in": o.DueDateIn}
		and = append(and, query)
	}
	if o.DueDateNin != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$nin": o.DueDateNin}
		and = append(and, query)
	}
	if o.DueDateExists != nil {
		query := bson.M{}
		query["dueDate"] = bson.M{"$exists": *o.DueDateExists}
		and = append(and, query)
	}
	if o.IsPublicEq != nil {
		query := bson.M{}
		query["isPublic"] = o.IsPublicEq
		and = append(and, query)
	}
	if o.IsPublicNe != nil {
		query := bson.M{}
		query["isPublic"] = bson.M{"$ne": o.IsPublicNe}
		and = append(and, query)
	}
	if o.IsPublicGt != nil {
		query := bson.M{}
		query["isPublic"] = bson.M{"$gt": o.IsPublicGt}
		and = append(and, query)
	}
	if o.IsPublicGte != nil {
		query := bson.M{}
		query["isPublic"] = bson.M{"$gte": o.IsPublicGte}
		and = append(and, query)
	}
	if o.IsPublicLt != nil {
		query := bson.M{}
		query["isPublic"] = bson.M{"$lt": o.IsPublicLt}
		and = append(and, query)
	}
	if o.IsPublicLte != nil {
		query := bson.M{}
		query["isPublic"] = bson.M{"$lte": o.IsPublicLte}
		and = append(and, query)
	}
	if o.IsPublicIn != nil {
		query := bson.M{}
		query["isPublic"] = bson.M{"$in": o.IsPublicIn}
		and = append(and, query)
	}
	if o.IsPublicNin != nil {
		query := bson.M{}
		query["isPublic"] = bson.M{"$nin": o.IsPublicNin}
		and = append(and, query)
	}
	if o.IsPublicExists != nil {
		query := bson.M{}
		query["isPublic"] = bson.M{"$exists": *o.IsPublicExists}
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
	if o.ParentNamesEq != nil {
		query := bson.M{}
		query["parentNames"] = o.ParentNamesEq
		and = append(and, query)
	}
	if o.ParentNamesNe != nil {
		query := bson.M{}
		query["parentNames"] = bson.M{"$ne": o.ParentNamesNe}
		and = append(and, query)
	}
	if o.ParentNamesGt != nil {
		query := bson.M{}
		query["parentNames"] = bson.M{"$gt": o.ParentNamesGt}
		and = append(and, query)
	}
	if o.ParentNamesGte != nil {
		query := bson.M{}
		query["parentNames"] = bson.M{"$gte": o.ParentNamesGte}
		and = append(and, query)
	}
	if o.ParentNamesLt != nil {
		query := bson.M{}
		query["parentNames"] = bson.M{"$lt": o.ParentNamesLt}
		and = append(and, query)
	}
	if o.ParentNamesLte != nil {
		query := bson.M{}
		query["parentNames"] = bson.M{"$lte": o.ParentNamesLte}
		and = append(and, query)
	}
	if o.ParentNamesIn != nil {
		query := bson.M{}
		query["parentNames"] = bson.M{"$in": o.ParentNamesIn}
		and = append(and, query)
	}
	if o.ParentNamesNin != nil {
		query := bson.M{}
		query["parentNames"] = bson.M{"$nin": o.ParentNamesNin}
		and = append(and, query)
	}
	if o.ParentNamesExists != nil {
		query := bson.M{}
		query["parentNames"] = bson.M{"$exists": *o.ParentNamesExists}
		and = append(and, query)
	}
	if o.ParentNamesLike != nil {
		query := bson.M{}
		query["parentNames"] = bson.M{"$regex": o.ParentNamesLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ParentNamesNlike != nil {
		query := bson.M{}
		query["parentNames"] = bson.M{"$not": bson.M{"$regex": o.ParentNamesNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.SlugEq != nil {
		query := bson.M{}
		query["slug"] = o.SlugEq
		and = append(and, query)
	}
	if o.SlugNe != nil {
		query := bson.M{}
		query["slug"] = bson.M{"$ne": o.SlugNe}
		and = append(and, query)
	}
	if o.SlugGt != nil {
		query := bson.M{}
		query["slug"] = bson.M{"$gt": o.SlugGt}
		and = append(and, query)
	}
	if o.SlugGte != nil {
		query := bson.M{}
		query["slug"] = bson.M{"$gte": o.SlugGte}
		and = append(and, query)
	}
	if o.SlugLt != nil {
		query := bson.M{}
		query["slug"] = bson.M{"$lt": o.SlugLt}
		and = append(and, query)
	}
	if o.SlugLte != nil {
		query := bson.M{}
		query["slug"] = bson.M{"$lte": o.SlugLte}
		and = append(and, query)
	}
	if o.SlugIn != nil {
		query := bson.M{}
		query["slug"] = bson.M{"$in": o.SlugIn}
		and = append(and, query)
	}
	if o.SlugNin != nil {
		query := bson.M{}
		query["slug"] = bson.M{"$nin": o.SlugNin}
		and = append(and, query)
	}
	if o.SlugExists != nil {
		query := bson.M{}
		query["slug"] = bson.M{"$exists": *o.SlugExists}
		and = append(and, query)
	}
	if o.SlugLike != nil {
		query := bson.M{}
		query["slug"] = bson.M{"$regex": o.SlugLike, "$options": "i"}
		and = append(and, query)
	}
	if o.SlugNlike != nil {
		query := bson.M{}
		query["slug"] = bson.M{"$not": bson.M{"$regex": o.SlugNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.ThemeColorEq != nil {
		query := bson.M{}
		query["themeColor"] = o.ThemeColorEq
		and = append(and, query)
	}
	if o.ThemeColorNe != nil {
		query := bson.M{}
		query["themeColor"] = bson.M{"$ne": o.ThemeColorNe}
		and = append(and, query)
	}
	if o.ThemeColorGt != nil {
		query := bson.M{}
		query["themeColor"] = bson.M{"$gt": o.ThemeColorGt}
		and = append(and, query)
	}
	if o.ThemeColorGte != nil {
		query := bson.M{}
		query["themeColor"] = bson.M{"$gte": o.ThemeColorGte}
		and = append(and, query)
	}
	if o.ThemeColorLt != nil {
		query := bson.M{}
		query["themeColor"] = bson.M{"$lt": o.ThemeColorLt}
		and = append(and, query)
	}
	if o.ThemeColorLte != nil {
		query := bson.M{}
		query["themeColor"] = bson.M{"$lte": o.ThemeColorLte}
		and = append(and, query)
	}
	if o.ThemeColorIn != nil {
		query := bson.M{}
		query["themeColor"] = bson.M{"$in": o.ThemeColorIn}
		and = append(and, query)
	}
	if o.ThemeColorNin != nil {
		query := bson.M{}
		query["themeColor"] = bson.M{"$nin": o.ThemeColorNin}
		and = append(and, query)
	}
	if o.ThemeColorExists != nil {
		query := bson.M{}
		query["themeColor"] = bson.M{"$exists": *o.ThemeColorExists}
		and = append(and, query)
	}
	if o.ThemeColorLike != nil {
		query := bson.M{}
		query["themeColor"] = bson.M{"$regex": o.ThemeColorLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ThemeColorNlike != nil {
		query := bson.M{}
		query["themeColor"] = bson.M{"$not": bson.M{"$regex": o.ThemeColorNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.TitleEq != nil {
		query := bson.M{}
		query["title"] = o.TitleEq
		and = append(and, query)
	}
	if o.TitleNe != nil {
		query := bson.M{}
		query["title"] = bson.M{"$ne": o.TitleNe}
		and = append(and, query)
	}
	if o.TitleGt != nil {
		query := bson.M{}
		query["title"] = bson.M{"$gt": o.TitleGt}
		and = append(and, query)
	}
	if o.TitleGte != nil {
		query := bson.M{}
		query["title"] = bson.M{"$gte": o.TitleGte}
		and = append(and, query)
	}
	if o.TitleLt != nil {
		query := bson.M{}
		query["title"] = bson.M{"$lt": o.TitleLt}
		and = append(and, query)
	}
	if o.TitleLte != nil {
		query := bson.M{}
		query["title"] = bson.M{"$lte": o.TitleLte}
		and = append(and, query)
	}
	if o.TitleIn != nil {
		query := bson.M{}
		query["title"] = bson.M{"$in": o.TitleIn}
		and = append(and, query)
	}
	if o.TitleNin != nil {
		query := bson.M{}
		query["title"] = bson.M{"$nin": o.TitleNin}
		and = append(and, query)
	}
	if o.TitleExists != nil {
		query := bson.M{}
		query["title"] = bson.M{"$exists": *o.TitleExists}
		and = append(and, query)
	}
	if o.TitleLike != nil {
		query := bson.M{}
		query["title"] = bson.M{"$regex": o.TitleLike, "$options": "i"}
		and = append(and, query)
	}
	if o.TitleNlike != nil {
		query := bson.M{}
		query["title"] = bson.M{"$not": bson.M{"$regex": o.TitleNlike, "$options": "i"}}
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
	if o.WelcomeMessageEq != nil {
		query := bson.M{}
		query["welcomeMessage"] = o.WelcomeMessageEq
		and = append(and, query)
	}
	if o.WelcomeMessageNe != nil {
		query := bson.M{}
		query["welcomeMessage"] = bson.M{"$ne": o.WelcomeMessageNe}
		and = append(and, query)
	}
	if o.WelcomeMessageGt != nil {
		query := bson.M{}
		query["welcomeMessage"] = bson.M{"$gt": o.WelcomeMessageGt}
		and = append(and, query)
	}
	if o.WelcomeMessageGte != nil {
		query := bson.M{}
		query["welcomeMessage"] = bson.M{"$gte": o.WelcomeMessageGte}
		and = append(and, query)
	}
	if o.WelcomeMessageLt != nil {
		query := bson.M{}
		query["welcomeMessage"] = bson.M{"$lt": o.WelcomeMessageLt}
		and = append(and, query)
	}
	if o.WelcomeMessageLte != nil {
		query := bson.M{}
		query["welcomeMessage"] = bson.M{"$lte": o.WelcomeMessageLte}
		and = append(and, query)
	}
	if o.WelcomeMessageIn != nil {
		query := bson.M{}
		query["welcomeMessage"] = bson.M{"$in": o.WelcomeMessageIn}
		and = append(and, query)
	}
	if o.WelcomeMessageNin != nil {
		query := bson.M{}
		query["welcomeMessage"] = bson.M{"$nin": o.WelcomeMessageNin}
		and = append(and, query)
	}
	if o.WelcomeMessageExists != nil {
		query := bson.M{}
		query["welcomeMessage"] = bson.M{"$exists": *o.WelcomeMessageExists}
		and = append(and, query)
	}
	if o.WelcomeMessageLike != nil {
		query := bson.M{}
		query["welcomeMessage"] = bson.M{"$regex": o.WelcomeMessageLike, "$options": "i"}
		and = append(and, query)
	}
	if o.WelcomeMessageNlike != nil {
		query := bson.M{}
		query["welcomeMessage"] = bson.M{"$not": bson.M{"$regex": o.WelcomeMessageNlike, "$options": "i"}}
		and = append(and, query)
	}
	return and, nil
}

type MongoSortParams struct {
	CreatedAt int8
	OwnerId   int8
	Slug      int8
	UpdatedAt int8
}
