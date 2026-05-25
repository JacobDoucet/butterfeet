package registry_item

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_item_source"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id                 string
	Created            actor_trace.Model
	Currency           string
	Description        string
	ImageUrl           string
	Notes              string
	OwnerPurchased     bool
	Position           int
	PriceCents         int
	ProductUrl         string
	Quantity           int
	RegistryId         string
	Source             enum_item_source.Value
	Title              string
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
	if projection.Created {
		elemcreated0, err := m.Created.ToMongoRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.Currency {
		elemcurrency0 := m.Currency
		r.Currency = &elemcurrency0
	}
	if projection.Description {
		elemdescription0 := m.Description
		r.Description = &elemdescription0
	}
	if projection.ImageUrl {
		elemimageUrl0 := m.ImageUrl
		r.ImageUrl = &elemimageUrl0
	}
	if projection.Notes {
		elemnotes0 := m.Notes
		r.Notes = &elemnotes0
	}
	if projection.OwnerPurchased {
		elemownerPurchased0 := m.OwnerPurchased
		r.OwnerPurchased = &elemownerPurchased0
	}
	if projection.Position {
		elemposition0 := m.Position
		r.Position = &elemposition0
	}
	if projection.PriceCents {
		elempriceCents0 := m.PriceCents
		r.PriceCents = &elempriceCents0
	}
	if projection.ProductUrl {
		elemproductUrl0 := m.ProductUrl
		r.ProductUrl = &elemproductUrl0
	}
	if projection.Quantity {
		elemquantity0 := m.Quantity
		r.Quantity = &elemquantity0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0, err := primitive.ObjectIDFromHex(m.RegistryId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.RegistryId"), err)
		}
		r.RegistryId = &elemregistryId0
	}
	if projection.Source {
		elemsource0 := m.Source
		r.Source = &elemsource0
	}
	if projection.Title {
		elemtitle0 := m.Title
		r.Title = &elemtitle0
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
	if projection.Created {
		elemcreated0, err := m.Created.ToHTTPRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.Currency {
		elemcurrency0 := m.Currency
		r.Currency = &elemcurrency0
	}
	if projection.Description {
		elemdescription0 := m.Description
		r.Description = &elemdescription0
	}
	if projection.ImageUrl {
		elemimageUrl0 := m.ImageUrl
		r.ImageUrl = &elemimageUrl0
	}
	if projection.Notes {
		elemnotes0 := m.Notes
		r.Notes = &elemnotes0
	}
	if projection.OwnerPurchased {
		elemownerPurchased0 := m.OwnerPurchased
		r.OwnerPurchased = &elemownerPurchased0
	}
	if projection.Position {
		elemposition0 := m.Position
		r.Position = &elemposition0
	}
	if projection.PriceCents {
		elempriceCents0 := m.PriceCents
		r.PriceCents = &elempriceCents0
	}
	if projection.ProductUrl {
		elemproductUrl0 := m.ProductUrl
		r.ProductUrl = &elemproductUrl0
	}
	if projection.Quantity {
		elemquantity0 := m.Quantity
		r.Quantity = &elemquantity0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0 := m.RegistryId
		r.RegistryId = &elemregistryId0
	}
	if projection.Source {
		elemsource0 := m.Source
		r.Source = &elemsource0
	}
	if projection.Title {
		elemtitle0 := m.Title
		r.Title = &elemtitle0
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

type WhereClause struct {
	// id (Ref<RegistryItem>) search options
	IdEq     *string
	IdIn     *[]string
	IdNin    *[]string
	IdExists *bool
	// created (ActorTrace) search options
	Created *actor_trace.WhereClause
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
	// ownerPurchased (bool) search options
	OwnerPurchasedEq     *bool
	OwnerPurchasedNe     *bool
	OwnerPurchasedGt     *bool
	OwnerPurchasedGte    *bool
	OwnerPurchasedLt     *bool
	OwnerPurchasedLte    *bool
	OwnerPurchasedIn     *[]bool
	OwnerPurchasedNin    *[]bool
	OwnerPurchasedExists *bool
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
	RegistryIdEq     *string
	RegistryIdIn     *[]string
	RegistryIdNin    *[]string
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
	if o.Created != nil {
		elemcreated0, err := o.Created.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.Created = &elemcreated0
	}
	if o.CurrencyEq != nil {
		elemcurrencyEq0 := o.CurrencyEq
		to.CurrencyEq = elemcurrencyEq0
	}
	if o.CurrencyNe != nil {
		elemcurrencyNe0 := o.CurrencyNe
		to.CurrencyNe = elemcurrencyNe0
	}
	if o.CurrencyGt != nil {
		elemcurrencyGt0 := o.CurrencyGt
		to.CurrencyGt = elemcurrencyGt0
	}
	if o.CurrencyGte != nil {
		elemcurrencyGte0 := o.CurrencyGte
		to.CurrencyGte = elemcurrencyGte0
	}
	if o.CurrencyLt != nil {
		elemcurrencyLt0 := o.CurrencyLt
		to.CurrencyLt = elemcurrencyLt0
	}
	if o.CurrencyLte != nil {
		elemcurrencyLte0 := o.CurrencyLte
		to.CurrencyLte = elemcurrencyLte0
	}
	if o.CurrencyIn != nil {
		elemcurrencyIn0 := make([]string, 0)
		for _, ocurrencyIn0 := range *o.CurrencyIn {
			elemcurrencyIn1 := ocurrencyIn0
			elemcurrencyIn0 = append(elemcurrencyIn0, elemcurrencyIn1)
		}
		to.CurrencyIn = &elemcurrencyIn0
	}
	if o.CurrencyNin != nil {
		elemcurrencyNin0 := make([]string, 0)
		for _, ocurrencyNin0 := range *o.CurrencyNin {
			elemcurrencyNin1 := ocurrencyNin0
			elemcurrencyNin0 = append(elemcurrencyNin0, elemcurrencyNin1)
		}
		to.CurrencyNin = &elemcurrencyNin0
	}
	if o.CurrencyExists != nil {
		elemcurrencyExists0 := o.CurrencyExists
		to.CurrencyExists = elemcurrencyExists0
	}
	if o.CurrencyLike != nil {
		elemcurrencyLike0 := o.CurrencyLike
		to.CurrencyLike = elemcurrencyLike0
	}
	if o.CurrencyNlike != nil {
		elemcurrencyNlike0 := o.CurrencyNlike
		to.CurrencyNlike = elemcurrencyNlike0
	}
	if o.DescriptionEq != nil {
		elemdescriptionEq0 := o.DescriptionEq
		to.DescriptionEq = elemdescriptionEq0
	}
	if o.DescriptionNe != nil {
		elemdescriptionNe0 := o.DescriptionNe
		to.DescriptionNe = elemdescriptionNe0
	}
	if o.DescriptionGt != nil {
		elemdescriptionGt0 := o.DescriptionGt
		to.DescriptionGt = elemdescriptionGt0
	}
	if o.DescriptionGte != nil {
		elemdescriptionGte0 := o.DescriptionGte
		to.DescriptionGte = elemdescriptionGte0
	}
	if o.DescriptionLt != nil {
		elemdescriptionLt0 := o.DescriptionLt
		to.DescriptionLt = elemdescriptionLt0
	}
	if o.DescriptionLte != nil {
		elemdescriptionLte0 := o.DescriptionLte
		to.DescriptionLte = elemdescriptionLte0
	}
	if o.DescriptionIn != nil {
		elemdescriptionIn0 := make([]string, 0)
		for _, odescriptionIn0 := range *o.DescriptionIn {
			elemdescriptionIn1 := odescriptionIn0
			elemdescriptionIn0 = append(elemdescriptionIn0, elemdescriptionIn1)
		}
		to.DescriptionIn = &elemdescriptionIn0
	}
	if o.DescriptionNin != nil {
		elemdescriptionNin0 := make([]string, 0)
		for _, odescriptionNin0 := range *o.DescriptionNin {
			elemdescriptionNin1 := odescriptionNin0
			elemdescriptionNin0 = append(elemdescriptionNin0, elemdescriptionNin1)
		}
		to.DescriptionNin = &elemdescriptionNin0
	}
	if o.DescriptionExists != nil {
		elemdescriptionExists0 := o.DescriptionExists
		to.DescriptionExists = elemdescriptionExists0
	}
	if o.DescriptionLike != nil {
		elemdescriptionLike0 := o.DescriptionLike
		to.DescriptionLike = elemdescriptionLike0
	}
	if o.DescriptionNlike != nil {
		elemdescriptionNlike0 := o.DescriptionNlike
		to.DescriptionNlike = elemdescriptionNlike0
	}
	if o.ImageUrlEq != nil {
		elemimageUrlEq0 := o.ImageUrlEq
		to.ImageUrlEq = elemimageUrlEq0
	}
	if o.ImageUrlNe != nil {
		elemimageUrlNe0 := o.ImageUrlNe
		to.ImageUrlNe = elemimageUrlNe0
	}
	if o.ImageUrlGt != nil {
		elemimageUrlGt0 := o.ImageUrlGt
		to.ImageUrlGt = elemimageUrlGt0
	}
	if o.ImageUrlGte != nil {
		elemimageUrlGte0 := o.ImageUrlGte
		to.ImageUrlGte = elemimageUrlGte0
	}
	if o.ImageUrlLt != nil {
		elemimageUrlLt0 := o.ImageUrlLt
		to.ImageUrlLt = elemimageUrlLt0
	}
	if o.ImageUrlLte != nil {
		elemimageUrlLte0 := o.ImageUrlLte
		to.ImageUrlLte = elemimageUrlLte0
	}
	if o.ImageUrlIn != nil {
		elemimageUrlIn0 := make([]string, 0)
		for _, oimageUrlIn0 := range *o.ImageUrlIn {
			elemimageUrlIn1 := oimageUrlIn0
			elemimageUrlIn0 = append(elemimageUrlIn0, elemimageUrlIn1)
		}
		to.ImageUrlIn = &elemimageUrlIn0
	}
	if o.ImageUrlNin != nil {
		elemimageUrlNin0 := make([]string, 0)
		for _, oimageUrlNin0 := range *o.ImageUrlNin {
			elemimageUrlNin1 := oimageUrlNin0
			elemimageUrlNin0 = append(elemimageUrlNin0, elemimageUrlNin1)
		}
		to.ImageUrlNin = &elemimageUrlNin0
	}
	if o.ImageUrlExists != nil {
		elemimageUrlExists0 := o.ImageUrlExists
		to.ImageUrlExists = elemimageUrlExists0
	}
	if o.ImageUrlLike != nil {
		elemimageUrlLike0 := o.ImageUrlLike
		to.ImageUrlLike = elemimageUrlLike0
	}
	if o.ImageUrlNlike != nil {
		elemimageUrlNlike0 := o.ImageUrlNlike
		to.ImageUrlNlike = elemimageUrlNlike0
	}
	if o.NotesEq != nil {
		elemnotesEq0 := o.NotesEq
		to.NotesEq = elemnotesEq0
	}
	if o.NotesNe != nil {
		elemnotesNe0 := o.NotesNe
		to.NotesNe = elemnotesNe0
	}
	if o.NotesGt != nil {
		elemnotesGt0 := o.NotesGt
		to.NotesGt = elemnotesGt0
	}
	if o.NotesGte != nil {
		elemnotesGte0 := o.NotesGte
		to.NotesGte = elemnotesGte0
	}
	if o.NotesLt != nil {
		elemnotesLt0 := o.NotesLt
		to.NotesLt = elemnotesLt0
	}
	if o.NotesLte != nil {
		elemnotesLte0 := o.NotesLte
		to.NotesLte = elemnotesLte0
	}
	if o.NotesIn != nil {
		elemnotesIn0 := make([]string, 0)
		for _, onotesIn0 := range *o.NotesIn {
			elemnotesIn1 := onotesIn0
			elemnotesIn0 = append(elemnotesIn0, elemnotesIn1)
		}
		to.NotesIn = &elemnotesIn0
	}
	if o.NotesNin != nil {
		elemnotesNin0 := make([]string, 0)
		for _, onotesNin0 := range *o.NotesNin {
			elemnotesNin1 := onotesNin0
			elemnotesNin0 = append(elemnotesNin0, elemnotesNin1)
		}
		to.NotesNin = &elemnotesNin0
	}
	if o.NotesExists != nil {
		elemnotesExists0 := o.NotesExists
		to.NotesExists = elemnotesExists0
	}
	if o.NotesLike != nil {
		elemnotesLike0 := o.NotesLike
		to.NotesLike = elemnotesLike0
	}
	if o.NotesNlike != nil {
		elemnotesNlike0 := o.NotesNlike
		to.NotesNlike = elemnotesNlike0
	}
	if o.OwnerPurchasedEq != nil {
		elemownerPurchasedEq0 := o.OwnerPurchasedEq
		to.OwnerPurchasedEq = elemownerPurchasedEq0
	}
	if o.OwnerPurchasedNe != nil {
		elemownerPurchasedNe0 := o.OwnerPurchasedNe
		to.OwnerPurchasedNe = elemownerPurchasedNe0
	}
	if o.OwnerPurchasedGt != nil {
		elemownerPurchasedGt0 := o.OwnerPurchasedGt
		to.OwnerPurchasedGt = elemownerPurchasedGt0
	}
	if o.OwnerPurchasedGte != nil {
		elemownerPurchasedGte0 := o.OwnerPurchasedGte
		to.OwnerPurchasedGte = elemownerPurchasedGte0
	}
	if o.OwnerPurchasedLt != nil {
		elemownerPurchasedLt0 := o.OwnerPurchasedLt
		to.OwnerPurchasedLt = elemownerPurchasedLt0
	}
	if o.OwnerPurchasedLte != nil {
		elemownerPurchasedLte0 := o.OwnerPurchasedLte
		to.OwnerPurchasedLte = elemownerPurchasedLte0
	}
	if o.OwnerPurchasedIn != nil {
		elemownerPurchasedIn0 := make([]bool, 0)
		for _, oownerPurchasedIn0 := range *o.OwnerPurchasedIn {
			elemownerPurchasedIn1 := oownerPurchasedIn0
			elemownerPurchasedIn0 = append(elemownerPurchasedIn0, elemownerPurchasedIn1)
		}
		to.OwnerPurchasedIn = &elemownerPurchasedIn0
	}
	if o.OwnerPurchasedNin != nil {
		elemownerPurchasedNin0 := make([]bool, 0)
		for _, oownerPurchasedNin0 := range *o.OwnerPurchasedNin {
			elemownerPurchasedNin1 := oownerPurchasedNin0
			elemownerPurchasedNin0 = append(elemownerPurchasedNin0, elemownerPurchasedNin1)
		}
		to.OwnerPurchasedNin = &elemownerPurchasedNin0
	}
	if o.OwnerPurchasedExists != nil {
		elemownerPurchasedExists0 := o.OwnerPurchasedExists
		to.OwnerPurchasedExists = elemownerPurchasedExists0
	}
	if o.PositionEq != nil {
		elempositionEq0 := o.PositionEq
		to.PositionEq = elempositionEq0
	}
	if o.PositionNe != nil {
		elempositionNe0 := o.PositionNe
		to.PositionNe = elempositionNe0
	}
	if o.PositionGt != nil {
		elempositionGt0 := o.PositionGt
		to.PositionGt = elempositionGt0
	}
	if o.PositionGte != nil {
		elempositionGte0 := o.PositionGte
		to.PositionGte = elempositionGte0
	}
	if o.PositionLt != nil {
		elempositionLt0 := o.PositionLt
		to.PositionLt = elempositionLt0
	}
	if o.PositionLte != nil {
		elempositionLte0 := o.PositionLte
		to.PositionLte = elempositionLte0
	}
	if o.PositionIn != nil {
		elempositionIn0 := make([]int, 0)
		for _, opositionIn0 := range *o.PositionIn {
			elempositionIn1 := opositionIn0
			elempositionIn0 = append(elempositionIn0, elempositionIn1)
		}
		to.PositionIn = &elempositionIn0
	}
	if o.PositionNin != nil {
		elempositionNin0 := make([]int, 0)
		for _, opositionNin0 := range *o.PositionNin {
			elempositionNin1 := opositionNin0
			elempositionNin0 = append(elempositionNin0, elempositionNin1)
		}
		to.PositionNin = &elempositionNin0
	}
	if o.PositionExists != nil {
		elempositionExists0 := o.PositionExists
		to.PositionExists = elempositionExists0
	}
	if o.PriceCentsEq != nil {
		elempriceCentsEq0 := o.PriceCentsEq
		to.PriceCentsEq = elempriceCentsEq0
	}
	if o.PriceCentsNe != nil {
		elempriceCentsNe0 := o.PriceCentsNe
		to.PriceCentsNe = elempriceCentsNe0
	}
	if o.PriceCentsGt != nil {
		elempriceCentsGt0 := o.PriceCentsGt
		to.PriceCentsGt = elempriceCentsGt0
	}
	if o.PriceCentsGte != nil {
		elempriceCentsGte0 := o.PriceCentsGte
		to.PriceCentsGte = elempriceCentsGte0
	}
	if o.PriceCentsLt != nil {
		elempriceCentsLt0 := o.PriceCentsLt
		to.PriceCentsLt = elempriceCentsLt0
	}
	if o.PriceCentsLte != nil {
		elempriceCentsLte0 := o.PriceCentsLte
		to.PriceCentsLte = elempriceCentsLte0
	}
	if o.PriceCentsIn != nil {
		elempriceCentsIn0 := make([]int, 0)
		for _, opriceCentsIn0 := range *o.PriceCentsIn {
			elempriceCentsIn1 := opriceCentsIn0
			elempriceCentsIn0 = append(elempriceCentsIn0, elempriceCentsIn1)
		}
		to.PriceCentsIn = &elempriceCentsIn0
	}
	if o.PriceCentsNin != nil {
		elempriceCentsNin0 := make([]int, 0)
		for _, opriceCentsNin0 := range *o.PriceCentsNin {
			elempriceCentsNin1 := opriceCentsNin0
			elempriceCentsNin0 = append(elempriceCentsNin0, elempriceCentsNin1)
		}
		to.PriceCentsNin = &elempriceCentsNin0
	}
	if o.PriceCentsExists != nil {
		elempriceCentsExists0 := o.PriceCentsExists
		to.PriceCentsExists = elempriceCentsExists0
	}
	if o.ProductUrlEq != nil {
		elemproductUrlEq0 := o.ProductUrlEq
		to.ProductUrlEq = elemproductUrlEq0
	}
	if o.ProductUrlNe != nil {
		elemproductUrlNe0 := o.ProductUrlNe
		to.ProductUrlNe = elemproductUrlNe0
	}
	if o.ProductUrlGt != nil {
		elemproductUrlGt0 := o.ProductUrlGt
		to.ProductUrlGt = elemproductUrlGt0
	}
	if o.ProductUrlGte != nil {
		elemproductUrlGte0 := o.ProductUrlGte
		to.ProductUrlGte = elemproductUrlGte0
	}
	if o.ProductUrlLt != nil {
		elemproductUrlLt0 := o.ProductUrlLt
		to.ProductUrlLt = elemproductUrlLt0
	}
	if o.ProductUrlLte != nil {
		elemproductUrlLte0 := o.ProductUrlLte
		to.ProductUrlLte = elemproductUrlLte0
	}
	if o.ProductUrlIn != nil {
		elemproductUrlIn0 := make([]string, 0)
		for _, oproductUrlIn0 := range *o.ProductUrlIn {
			elemproductUrlIn1 := oproductUrlIn0
			elemproductUrlIn0 = append(elemproductUrlIn0, elemproductUrlIn1)
		}
		to.ProductUrlIn = &elemproductUrlIn0
	}
	if o.ProductUrlNin != nil {
		elemproductUrlNin0 := make([]string, 0)
		for _, oproductUrlNin0 := range *o.ProductUrlNin {
			elemproductUrlNin1 := oproductUrlNin0
			elemproductUrlNin0 = append(elemproductUrlNin0, elemproductUrlNin1)
		}
		to.ProductUrlNin = &elemproductUrlNin0
	}
	if o.ProductUrlExists != nil {
		elemproductUrlExists0 := o.ProductUrlExists
		to.ProductUrlExists = elemproductUrlExists0
	}
	if o.ProductUrlLike != nil {
		elemproductUrlLike0 := o.ProductUrlLike
		to.ProductUrlLike = elemproductUrlLike0
	}
	if o.ProductUrlNlike != nil {
		elemproductUrlNlike0 := o.ProductUrlNlike
		to.ProductUrlNlike = elemproductUrlNlike0
	}
	if o.QuantityEq != nil {
		elemquantityEq0 := o.QuantityEq
		to.QuantityEq = elemquantityEq0
	}
	if o.QuantityNe != nil {
		elemquantityNe0 := o.QuantityNe
		to.QuantityNe = elemquantityNe0
	}
	if o.QuantityGt != nil {
		elemquantityGt0 := o.QuantityGt
		to.QuantityGt = elemquantityGt0
	}
	if o.QuantityGte != nil {
		elemquantityGte0 := o.QuantityGte
		to.QuantityGte = elemquantityGte0
	}
	if o.QuantityLt != nil {
		elemquantityLt0 := o.QuantityLt
		to.QuantityLt = elemquantityLt0
	}
	if o.QuantityLte != nil {
		elemquantityLte0 := o.QuantityLte
		to.QuantityLte = elemquantityLte0
	}
	if o.QuantityIn != nil {
		elemquantityIn0 := make([]int, 0)
		for _, oquantityIn0 := range *o.QuantityIn {
			elemquantityIn1 := oquantityIn0
			elemquantityIn0 = append(elemquantityIn0, elemquantityIn1)
		}
		to.QuantityIn = &elemquantityIn0
	}
	if o.QuantityNin != nil {
		elemquantityNin0 := make([]int, 0)
		for _, oquantityNin0 := range *o.QuantityNin {
			elemquantityNin1 := oquantityNin0
			elemquantityNin0 = append(elemquantityNin0, elemquantityNin1)
		}
		to.QuantityNin = &elemquantityNin0
	}
	if o.QuantityExists != nil {
		elemquantityExists0 := o.QuantityExists
		to.QuantityExists = elemquantityExists0
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
	if o.SourceEq != nil {
		elemsourceEq0 := o.SourceEq
		to.SourceEq = elemsourceEq0
	}
	if o.SourceNe != nil {
		elemsourceNe0 := o.SourceNe
		to.SourceNe = elemsourceNe0
	}
	if o.SourceGt != nil {
		elemsourceGt0 := o.SourceGt
		to.SourceGt = elemsourceGt0
	}
	if o.SourceGte != nil {
		elemsourceGte0 := o.SourceGte
		to.SourceGte = elemsourceGte0
	}
	if o.SourceLt != nil {
		elemsourceLt0 := o.SourceLt
		to.SourceLt = elemsourceLt0
	}
	if o.SourceLte != nil {
		elemsourceLte0 := o.SourceLte
		to.SourceLte = elemsourceLte0
	}
	if o.SourceIn != nil {
		elemsourceIn0 := make([]enum_item_source.Value, 0)
		for _, osourceIn0 := range *o.SourceIn {
			elemsourceIn1 := osourceIn0
			elemsourceIn0 = append(elemsourceIn0, elemsourceIn1)
		}
		to.SourceIn = &elemsourceIn0
	}
	if o.SourceNin != nil {
		elemsourceNin0 := make([]enum_item_source.Value, 0)
		for _, osourceNin0 := range *o.SourceNin {
			elemsourceNin1 := osourceNin0
			elemsourceNin0 = append(elemsourceNin0, elemsourceNin1)
		}
		to.SourceNin = &elemsourceNin0
	}
	if o.SourceExists != nil {
		elemsourceExists0 := o.SourceExists
		to.SourceExists = elemsourceExists0
	}
	if o.TitleEq != nil {
		elemtitleEq0 := o.TitleEq
		to.TitleEq = elemtitleEq0
	}
	if o.TitleNe != nil {
		elemtitleNe0 := o.TitleNe
		to.TitleNe = elemtitleNe0
	}
	if o.TitleGt != nil {
		elemtitleGt0 := o.TitleGt
		to.TitleGt = elemtitleGt0
	}
	if o.TitleGte != nil {
		elemtitleGte0 := o.TitleGte
		to.TitleGte = elemtitleGte0
	}
	if o.TitleLt != nil {
		elemtitleLt0 := o.TitleLt
		to.TitleLt = elemtitleLt0
	}
	if o.TitleLte != nil {
		elemtitleLte0 := o.TitleLte
		to.TitleLte = elemtitleLte0
	}
	if o.TitleIn != nil {
		elemtitleIn0 := make([]string, 0)
		for _, otitleIn0 := range *o.TitleIn {
			elemtitleIn1 := otitleIn0
			elemtitleIn0 = append(elemtitleIn0, elemtitleIn1)
		}
		to.TitleIn = &elemtitleIn0
	}
	if o.TitleNin != nil {
		elemtitleNin0 := make([]string, 0)
		for _, otitleNin0 := range *o.TitleNin {
			elemtitleNin1 := otitleNin0
			elemtitleNin0 = append(elemtitleNin0, elemtitleNin1)
		}
		to.TitleNin = &elemtitleNin0
	}
	if o.TitleExists != nil {
		elemtitleExists0 := o.TitleExists
		to.TitleExists = elemtitleExists0
	}
	if o.TitleLike != nil {
		elemtitleLike0 := o.TitleLike
		to.TitleLike = elemtitleLike0
	}
	if o.TitleNlike != nil {
		elemtitleNlike0 := o.TitleNlike
		to.TitleNlike = elemtitleNlike0
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
	Position   int8
	RegistryId int8
	UpdatedAt  int8
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	to.CreatedAt = s.CreatedAt
	to.Position = s.Position
	to.RegistryId = s.RegistryId
	to.UpdatedAt = s.UpdatedAt
	return to
}
