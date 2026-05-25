package registry_item

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_item_source"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRecord struct {
	Id                 *primitive.ObjectID      `bson:"_id,omitempty"`
	Created            *actor_trace.MongoRecord `bson:"created,omitempty"`
	Currency           *string                  `bson:"currency,omitempty"`
	Description        *string                  `bson:"description,omitempty"`
	ImageUrl           *string                  `bson:"imageUrl,omitempty"`
	Notes              *string                  `bson:"notes,omitempty"`
	Position           *int                     `bson:"position,omitempty"`
	PriceCents         *int                     `bson:"priceCents,omitempty"`
	ProductUrl         *string                  `bson:"productUrl,omitempty"`
	Quantity           *int                     `bson:"quantity,omitempty"`
	RegistryId         *primitive.ObjectID      `bson:"registryId,omitempty"`
	Source             *enum_item_source.Value  `bson:"source,omitempty"`
	Title              *string                  `bson:"title,omitempty"`
	Updated            *actor_trace.MongoRecord `bson:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.MongoRecord `bson:"updatedByOwnerUser,omitempty"`
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
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.Currency != nil {
		elemcurrency0 := r.Currency
		m.Currency = *elemcurrency0
	}
	if r.Description != nil {
		elemdescription0 := r.Description
		m.Description = *elemdescription0
	}
	if r.ImageUrl != nil {
		elemimageUrl0 := r.ImageUrl
		m.ImageUrl = *elemimageUrl0
	}
	if r.Notes != nil {
		elemnotes0 := r.Notes
		m.Notes = *elemnotes0
	}
	if r.Position != nil {
		elemposition0 := r.Position
		m.Position = *elemposition0
	}
	if r.PriceCents != nil {
		elempriceCents0 := r.PriceCents
		m.PriceCents = *elempriceCents0
	}
	if r.ProductUrl != nil {
		elemproductUrl0 := r.ProductUrl
		m.ProductUrl = *elemproductUrl0
	}
	if r.Quantity != nil {
		elemquantity0 := r.Quantity
		m.Quantity = *elemquantity0
	}
	if r.RegistryId != nil {
		elemregistryId0 := r.RegistryId.Hex()
		m.RegistryId = elemregistryId0
	}
	if r.Source != nil {
		elemsource0 := r.Source
		m.Source = *elemsource0
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
	return m, nil
}

type MongoSelectByIdQuery struct {
	Id primitive.ObjectID
}

type MongoWhereClause struct {
	// id (Ref<RegistryItem>) search options
	IdEq     *primitive.ObjectID
	IdIn     *[]primitive.ObjectID
	IdNin    *[]primitive.ObjectID
	IdExists *bool
	// created (ActorTrace) search options
	Created *actor_trace.MongoWhereClause
	// currency (string) search options
	CurrencyEq     *string
	CurrencyNe     *string
	CurrencyGt     *string
	CurrencyGte    *string
	CurrencyLt     *string
	CurrencyLte    *string
	CurrencyIn     *[]string
	CurrencyNin    *[]string
	CurrencyExists *bool
	CurrencyLike   *string
	CurrencyNlike  *string
	// description (string) search options
	DescriptionEq     *string
	DescriptionNe     *string
	DescriptionGt     *string
	DescriptionGte    *string
	DescriptionLt     *string
	DescriptionLte    *string
	DescriptionIn     *[]string
	DescriptionNin    *[]string
	DescriptionExists *bool
	DescriptionLike   *string
	DescriptionNlike  *string
	// imageUrl (string) search options
	ImageUrlEq     *string
	ImageUrlNe     *string
	ImageUrlGt     *string
	ImageUrlGte    *string
	ImageUrlLt     *string
	ImageUrlLte    *string
	ImageUrlIn     *[]string
	ImageUrlNin    *[]string
	ImageUrlExists *bool
	ImageUrlLike   *string
	ImageUrlNlike  *string
	// notes (string) search options
	NotesEq     *string
	NotesNe     *string
	NotesGt     *string
	NotesGte    *string
	NotesLt     *string
	NotesLte    *string
	NotesIn     *[]string
	NotesNin    *[]string
	NotesExists *bool
	NotesLike   *string
	NotesNlike  *string
	// position (int) search options
	PositionEq     *int
	PositionNe     *int
	PositionGt     *int
	PositionGte    *int
	PositionLt     *int
	PositionLte    *int
	PositionIn     *[]int
	PositionNin    *[]int
	PositionExists *bool
	// priceCents (int) search options
	PriceCentsEq     *int
	PriceCentsNe     *int
	PriceCentsGt     *int
	PriceCentsGte    *int
	PriceCentsLt     *int
	PriceCentsLte    *int
	PriceCentsIn     *[]int
	PriceCentsNin    *[]int
	PriceCentsExists *bool
	// productUrl (string) search options
	ProductUrlEq     *string
	ProductUrlNe     *string
	ProductUrlGt     *string
	ProductUrlGte    *string
	ProductUrlLt     *string
	ProductUrlLte    *string
	ProductUrlIn     *[]string
	ProductUrlNin    *[]string
	ProductUrlExists *bool
	ProductUrlLike   *string
	ProductUrlNlike  *string
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
	// registryId (ParentRef<Registry>) search options
	RegistryIdEq     *primitive.ObjectID
	RegistryIdIn     *[]primitive.ObjectID
	RegistryIdNin    *[]primitive.ObjectID
	RegistryIdExists *bool
	// source (ItemSource) search options
	SourceEq     *enum_item_source.Value
	SourceNe     *enum_item_source.Value
	SourceGt     *enum_item_source.Value
	SourceGte    *enum_item_source.Value
	SourceLt     *enum_item_source.Value
	SourceLte    *enum_item_source.Value
	SourceIn     *[]enum_item_source.Value
	SourceNin    *[]enum_item_source.Value
	SourceExists *bool
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
	if o.CurrencyEq != nil {
		query := bson.M{}
		query["currency"] = o.CurrencyEq
		and = append(and, query)
	}
	if o.CurrencyNe != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$ne": o.CurrencyNe}
		and = append(and, query)
	}
	if o.CurrencyGt != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$gt": o.CurrencyGt}
		and = append(and, query)
	}
	if o.CurrencyGte != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$gte": o.CurrencyGte}
		and = append(and, query)
	}
	if o.CurrencyLt != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$lt": o.CurrencyLt}
		and = append(and, query)
	}
	if o.CurrencyLte != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$lte": o.CurrencyLte}
		and = append(and, query)
	}
	if o.CurrencyIn != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$in": o.CurrencyIn}
		and = append(and, query)
	}
	if o.CurrencyNin != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$nin": o.CurrencyNin}
		and = append(and, query)
	}
	if o.CurrencyExists != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$exists": *o.CurrencyExists}
		and = append(and, query)
	}
	if o.CurrencyLike != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$regex": o.CurrencyLike, "$options": "i"}
		and = append(and, query)
	}
	if o.CurrencyNlike != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$not": bson.M{"$regex": o.CurrencyNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.DescriptionEq != nil {
		query := bson.M{}
		query["description"] = o.DescriptionEq
		and = append(and, query)
	}
	if o.DescriptionNe != nil {
		query := bson.M{}
		query["description"] = bson.M{"$ne": o.DescriptionNe}
		and = append(and, query)
	}
	if o.DescriptionGt != nil {
		query := bson.M{}
		query["description"] = bson.M{"$gt": o.DescriptionGt}
		and = append(and, query)
	}
	if o.DescriptionGte != nil {
		query := bson.M{}
		query["description"] = bson.M{"$gte": o.DescriptionGte}
		and = append(and, query)
	}
	if o.DescriptionLt != nil {
		query := bson.M{}
		query["description"] = bson.M{"$lt": o.DescriptionLt}
		and = append(and, query)
	}
	if o.DescriptionLte != nil {
		query := bson.M{}
		query["description"] = bson.M{"$lte": o.DescriptionLte}
		and = append(and, query)
	}
	if o.DescriptionIn != nil {
		query := bson.M{}
		query["description"] = bson.M{"$in": o.DescriptionIn}
		and = append(and, query)
	}
	if o.DescriptionNin != nil {
		query := bson.M{}
		query["description"] = bson.M{"$nin": o.DescriptionNin}
		and = append(and, query)
	}
	if o.DescriptionExists != nil {
		query := bson.M{}
		query["description"] = bson.M{"$exists": *o.DescriptionExists}
		and = append(and, query)
	}
	if o.DescriptionLike != nil {
		query := bson.M{}
		query["description"] = bson.M{"$regex": o.DescriptionLike, "$options": "i"}
		and = append(and, query)
	}
	if o.DescriptionNlike != nil {
		query := bson.M{}
		query["description"] = bson.M{"$not": bson.M{"$regex": o.DescriptionNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.ImageUrlEq != nil {
		query := bson.M{}
		query["imageUrl"] = o.ImageUrlEq
		and = append(and, query)
	}
	if o.ImageUrlNe != nil {
		query := bson.M{}
		query["imageUrl"] = bson.M{"$ne": o.ImageUrlNe}
		and = append(and, query)
	}
	if o.ImageUrlGt != nil {
		query := bson.M{}
		query["imageUrl"] = bson.M{"$gt": o.ImageUrlGt}
		and = append(and, query)
	}
	if o.ImageUrlGte != nil {
		query := bson.M{}
		query["imageUrl"] = bson.M{"$gte": o.ImageUrlGte}
		and = append(and, query)
	}
	if o.ImageUrlLt != nil {
		query := bson.M{}
		query["imageUrl"] = bson.M{"$lt": o.ImageUrlLt}
		and = append(and, query)
	}
	if o.ImageUrlLte != nil {
		query := bson.M{}
		query["imageUrl"] = bson.M{"$lte": o.ImageUrlLte}
		and = append(and, query)
	}
	if o.ImageUrlIn != nil {
		query := bson.M{}
		query["imageUrl"] = bson.M{"$in": o.ImageUrlIn}
		and = append(and, query)
	}
	if o.ImageUrlNin != nil {
		query := bson.M{}
		query["imageUrl"] = bson.M{"$nin": o.ImageUrlNin}
		and = append(and, query)
	}
	if o.ImageUrlExists != nil {
		query := bson.M{}
		query["imageUrl"] = bson.M{"$exists": *o.ImageUrlExists}
		and = append(and, query)
	}
	if o.ImageUrlLike != nil {
		query := bson.M{}
		query["imageUrl"] = bson.M{"$regex": o.ImageUrlLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ImageUrlNlike != nil {
		query := bson.M{}
		query["imageUrl"] = bson.M{"$not": bson.M{"$regex": o.ImageUrlNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.NotesEq != nil {
		query := bson.M{}
		query["notes"] = o.NotesEq
		and = append(and, query)
	}
	if o.NotesNe != nil {
		query := bson.M{}
		query["notes"] = bson.M{"$ne": o.NotesNe}
		and = append(and, query)
	}
	if o.NotesGt != nil {
		query := bson.M{}
		query["notes"] = bson.M{"$gt": o.NotesGt}
		and = append(and, query)
	}
	if o.NotesGte != nil {
		query := bson.M{}
		query["notes"] = bson.M{"$gte": o.NotesGte}
		and = append(and, query)
	}
	if o.NotesLt != nil {
		query := bson.M{}
		query["notes"] = bson.M{"$lt": o.NotesLt}
		and = append(and, query)
	}
	if o.NotesLte != nil {
		query := bson.M{}
		query["notes"] = bson.M{"$lte": o.NotesLte}
		and = append(and, query)
	}
	if o.NotesIn != nil {
		query := bson.M{}
		query["notes"] = bson.M{"$in": o.NotesIn}
		and = append(and, query)
	}
	if o.NotesNin != nil {
		query := bson.M{}
		query["notes"] = bson.M{"$nin": o.NotesNin}
		and = append(and, query)
	}
	if o.NotesExists != nil {
		query := bson.M{}
		query["notes"] = bson.M{"$exists": *o.NotesExists}
		and = append(and, query)
	}
	if o.NotesLike != nil {
		query := bson.M{}
		query["notes"] = bson.M{"$regex": o.NotesLike, "$options": "i"}
		and = append(and, query)
	}
	if o.NotesNlike != nil {
		query := bson.M{}
		query["notes"] = bson.M{"$not": bson.M{"$regex": o.NotesNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.PositionEq != nil {
		query := bson.M{}
		query["position"] = o.PositionEq
		and = append(and, query)
	}
	if o.PositionNe != nil {
		query := bson.M{}
		query["position"] = bson.M{"$ne": o.PositionNe}
		and = append(and, query)
	}
	if o.PositionGt != nil {
		query := bson.M{}
		query["position"] = bson.M{"$gt": o.PositionGt}
		and = append(and, query)
	}
	if o.PositionGte != nil {
		query := bson.M{}
		query["position"] = bson.M{"$gte": o.PositionGte}
		and = append(and, query)
	}
	if o.PositionLt != nil {
		query := bson.M{}
		query["position"] = bson.M{"$lt": o.PositionLt}
		and = append(and, query)
	}
	if o.PositionLte != nil {
		query := bson.M{}
		query["position"] = bson.M{"$lte": o.PositionLte}
		and = append(and, query)
	}
	if o.PositionIn != nil {
		query := bson.M{}
		query["position"] = bson.M{"$in": o.PositionIn}
		and = append(and, query)
	}
	if o.PositionNin != nil {
		query := bson.M{}
		query["position"] = bson.M{"$nin": o.PositionNin}
		and = append(and, query)
	}
	if o.PositionExists != nil {
		query := bson.M{}
		query["position"] = bson.M{"$exists": *o.PositionExists}
		and = append(and, query)
	}
	if o.PriceCentsEq != nil {
		query := bson.M{}
		query["priceCents"] = o.PriceCentsEq
		and = append(and, query)
	}
	if o.PriceCentsNe != nil {
		query := bson.M{}
		query["priceCents"] = bson.M{"$ne": o.PriceCentsNe}
		and = append(and, query)
	}
	if o.PriceCentsGt != nil {
		query := bson.M{}
		query["priceCents"] = bson.M{"$gt": o.PriceCentsGt}
		and = append(and, query)
	}
	if o.PriceCentsGte != nil {
		query := bson.M{}
		query["priceCents"] = bson.M{"$gte": o.PriceCentsGte}
		and = append(and, query)
	}
	if o.PriceCentsLt != nil {
		query := bson.M{}
		query["priceCents"] = bson.M{"$lt": o.PriceCentsLt}
		and = append(and, query)
	}
	if o.PriceCentsLte != nil {
		query := bson.M{}
		query["priceCents"] = bson.M{"$lte": o.PriceCentsLte}
		and = append(and, query)
	}
	if o.PriceCentsIn != nil {
		query := bson.M{}
		query["priceCents"] = bson.M{"$in": o.PriceCentsIn}
		and = append(and, query)
	}
	if o.PriceCentsNin != nil {
		query := bson.M{}
		query["priceCents"] = bson.M{"$nin": o.PriceCentsNin}
		and = append(and, query)
	}
	if o.PriceCentsExists != nil {
		query := bson.M{}
		query["priceCents"] = bson.M{"$exists": *o.PriceCentsExists}
		and = append(and, query)
	}
	if o.ProductUrlEq != nil {
		query := bson.M{}
		query["productUrl"] = o.ProductUrlEq
		and = append(and, query)
	}
	if o.ProductUrlNe != nil {
		query := bson.M{}
		query["productUrl"] = bson.M{"$ne": o.ProductUrlNe}
		and = append(and, query)
	}
	if o.ProductUrlGt != nil {
		query := bson.M{}
		query["productUrl"] = bson.M{"$gt": o.ProductUrlGt}
		and = append(and, query)
	}
	if o.ProductUrlGte != nil {
		query := bson.M{}
		query["productUrl"] = bson.M{"$gte": o.ProductUrlGte}
		and = append(and, query)
	}
	if o.ProductUrlLt != nil {
		query := bson.M{}
		query["productUrl"] = bson.M{"$lt": o.ProductUrlLt}
		and = append(and, query)
	}
	if o.ProductUrlLte != nil {
		query := bson.M{}
		query["productUrl"] = bson.M{"$lte": o.ProductUrlLte}
		and = append(and, query)
	}
	if o.ProductUrlIn != nil {
		query := bson.M{}
		query["productUrl"] = bson.M{"$in": o.ProductUrlIn}
		and = append(and, query)
	}
	if o.ProductUrlNin != nil {
		query := bson.M{}
		query["productUrl"] = bson.M{"$nin": o.ProductUrlNin}
		and = append(and, query)
	}
	if o.ProductUrlExists != nil {
		query := bson.M{}
		query["productUrl"] = bson.M{"$exists": *o.ProductUrlExists}
		and = append(and, query)
	}
	if o.ProductUrlLike != nil {
		query := bson.M{}
		query["productUrl"] = bson.M{"$regex": o.ProductUrlLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ProductUrlNlike != nil {
		query := bson.M{}
		query["productUrl"] = bson.M{"$not": bson.M{"$regex": o.ProductUrlNlike, "$options": "i"}}
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
	if o.SourceEq != nil {
		query := bson.M{}
		query["source"] = o.SourceEq
		and = append(and, query)
	}
	if o.SourceNe != nil {
		query := bson.M{}
		query["source"] = bson.M{"$ne": o.SourceNe}
		and = append(and, query)
	}
	if o.SourceGt != nil {
		query := bson.M{}
		query["source"] = bson.M{"$gt": o.SourceGt}
		and = append(and, query)
	}
	if o.SourceGte != nil {
		query := bson.M{}
		query["source"] = bson.M{"$gte": o.SourceGte}
		and = append(and, query)
	}
	if o.SourceLt != nil {
		query := bson.M{}
		query["source"] = bson.M{"$lt": o.SourceLt}
		and = append(and, query)
	}
	if o.SourceLte != nil {
		query := bson.M{}
		query["source"] = bson.M{"$lte": o.SourceLte}
		and = append(and, query)
	}
	if o.SourceIn != nil {
		query := bson.M{}
		query["source"] = bson.M{"$in": o.SourceIn}
		and = append(and, query)
	}
	if o.SourceNin != nil {
		query := bson.M{}
		query["source"] = bson.M{"$nin": o.SourceNin}
		and = append(and, query)
	}
	if o.SourceExists != nil {
		query := bson.M{}
		query["source"] = bson.M{"$exists": *o.SourceExists}
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
	return and, nil
}

type MongoSortParams struct {
	CreatedAt  int8
	Position   int8
	RegistryId int8
	UpdatedAt  int8
}
