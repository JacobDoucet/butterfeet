package registry_item

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_item_source"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id                 string
	AffiliateUrl       string
	CanonicalUrl       string
	Category           string
	Created            actor_trace.Model
	Currency           string
	Description        string
	ImageBgColor       string
	ImageUrl           string
	NoSubstitutes      bool
	Notes              string
	OriginalUrl        string
	OwnerPurchased     bool
	ParentItemId       string
	Position           int
	PriceCents         int
	ProductUrl         string
	Quantity           int
	QuantityUnlimited  bool
	RegistryId         string
	Retailer           string
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
	if projection.AffiliateUrl {
		elemaffiliateUrl0 := m.AffiliateUrl
		r.AffiliateUrl = &elemaffiliateUrl0
	}
	if projection.CanonicalUrl {
		elemcanonicalUrl0 := m.CanonicalUrl
		r.CanonicalUrl = &elemcanonicalUrl0
	}
	if projection.Category {
		elemcategory0 := m.Category
		r.Category = &elemcategory0
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
	if projection.ImageBgColor {
		elemimageBgColor0 := m.ImageBgColor
		r.ImageBgColor = &elemimageBgColor0
	}
	if projection.ImageUrl {
		elemimageUrl0 := m.ImageUrl
		r.ImageUrl = &elemimageUrl0
	}
	if projection.NoSubstitutes {
		elemnoSubstitutes0 := m.NoSubstitutes
		r.NoSubstitutes = &elemnoSubstitutes0
	}
	if projection.Notes {
		elemnotes0 := m.Notes
		r.Notes = &elemnotes0
	}
	if projection.OriginalUrl {
		elemoriginalUrl0 := m.OriginalUrl
		r.OriginalUrl = &elemoriginalUrl0
	}
	if projection.OwnerPurchased {
		elemownerPurchased0 := m.OwnerPurchased
		r.OwnerPurchased = &elemownerPurchased0
	}
	if projection.ParentItemId && m.ParentItemId != "" {
		elemparentItemId0, err := primitive.ObjectIDFromHex(m.ParentItemId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.ParentItemId"), err)
		}
		r.ParentItemId = &elemparentItemId0
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
	if projection.QuantityUnlimited {
		elemquantityUnlimited0 := m.QuantityUnlimited
		r.QuantityUnlimited = &elemquantityUnlimited0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0, err := primitive.ObjectIDFromHex(m.RegistryId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.RegistryId"), err)
		}
		r.RegistryId = &elemregistryId0
	}
	if projection.Retailer {
		elemretailer0 := m.Retailer
		r.Retailer = &elemretailer0
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
	if projection.AffiliateUrl {
		elemaffiliateUrl0 := m.AffiliateUrl
		r.AffiliateUrl = &elemaffiliateUrl0
	}
	if projection.CanonicalUrl {
		elemcanonicalUrl0 := m.CanonicalUrl
		r.CanonicalUrl = &elemcanonicalUrl0
	}
	if projection.Category {
		elemcategory0 := m.Category
		r.Category = &elemcategory0
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
	if projection.ImageBgColor {
		elemimageBgColor0 := m.ImageBgColor
		r.ImageBgColor = &elemimageBgColor0
	}
	if projection.ImageUrl {
		elemimageUrl0 := m.ImageUrl
		r.ImageUrl = &elemimageUrl0
	}
	if projection.NoSubstitutes {
		elemnoSubstitutes0 := m.NoSubstitutes
		r.NoSubstitutes = &elemnoSubstitutes0
	}
	if projection.Notes {
		elemnotes0 := m.Notes
		r.Notes = &elemnotes0
	}
	if projection.OriginalUrl {
		elemoriginalUrl0 := m.OriginalUrl
		r.OriginalUrl = &elemoriginalUrl0
	}
	if projection.OwnerPurchased {
		elemownerPurchased0 := m.OwnerPurchased
		r.OwnerPurchased = &elemownerPurchased0
	}
	if projection.ParentItemId && m.ParentItemId != "" {
		elemparentItemId0 := m.ParentItemId
		r.ParentItemId = &elemparentItemId0
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
	if projection.QuantityUnlimited {
		elemquantityUnlimited0 := m.QuantityUnlimited
		r.QuantityUnlimited = &elemquantityUnlimited0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0 := m.RegistryId
		r.RegistryId = &elemregistryId0
	}
	if projection.Retailer {
		elemretailer0 := m.Retailer
		r.Retailer = &elemretailer0
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
	// affiliateUrl (string) search options
	AffiliateUrlEq     *string
	AffiliateUrlNe     *string
	AffiliateUrlGt     *string
	AffiliateUrlGte    *string
	AffiliateUrlLt     *string
	AffiliateUrlLte    *string
	AffiliateUrlIn     *[]string
	AffiliateUrlNin    *[]string
	AffiliateUrlExists *bool
	AffiliateUrlLike   *string
	AffiliateUrlNlike  *string
	// canonicalUrl (string) search options
	CanonicalUrlEq     *string
	CanonicalUrlNe     *string
	CanonicalUrlGt     *string
	CanonicalUrlGte    *string
	CanonicalUrlLt     *string
	CanonicalUrlLte    *string
	CanonicalUrlIn     *[]string
	CanonicalUrlNin    *[]string
	CanonicalUrlExists *bool
	CanonicalUrlLike   *string
	CanonicalUrlNlike  *string
	// category (string) search options
	CategoryEq     *string
	CategoryNe     *string
	CategoryGt     *string
	CategoryGte    *string
	CategoryLt     *string
	CategoryLte    *string
	CategoryIn     *[]string
	CategoryNin    *[]string
	CategoryExists *bool
	CategoryLike   *string
	CategoryNlike  *string
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
	// imageBgColor (string) search options
	ImageBgColorEq     *string
	ImageBgColorNe     *string
	ImageBgColorGt     *string
	ImageBgColorGte    *string
	ImageBgColorLt     *string
	ImageBgColorLte    *string
	ImageBgColorIn     *[]string
	ImageBgColorNin    *[]string
	ImageBgColorExists *bool
	ImageBgColorLike   *string
	ImageBgColorNlike  *string
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
	// noSubstitutes (bool) search options
	NoSubstitutesEq     *bool
	NoSubstitutesNe     *bool
	NoSubstitutesGt     *bool
	NoSubstitutesGte    *bool
	NoSubstitutesLt     *bool
	NoSubstitutesLte    *bool
	NoSubstitutesIn     *[]bool
	NoSubstitutesNin    *[]bool
	NoSubstitutesExists *bool
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
	// originalUrl (string) search options
	OriginalUrlEq     *string
	OriginalUrlNe     *string
	OriginalUrlGt     *string
	OriginalUrlGte    *string
	OriginalUrlLt     *string
	OriginalUrlLte    *string
	OriginalUrlIn     *[]string
	OriginalUrlNin    *[]string
	OriginalUrlExists *bool
	OriginalUrlLike   *string
	OriginalUrlNlike  *string
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
	// parentItemId (Ref<RegistryItem>) search options
	ParentItemIdEq     *string
	ParentItemIdIn     *[]string
	ParentItemIdNin    *[]string
	ParentItemIdExists *bool
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
	// quantityUnlimited (bool) search options
	QuantityUnlimitedEq     *bool
	QuantityUnlimitedNe     *bool
	QuantityUnlimitedGt     *bool
	QuantityUnlimitedGte    *bool
	QuantityUnlimitedLt     *bool
	QuantityUnlimitedLte    *bool
	QuantityUnlimitedIn     *[]bool
	QuantityUnlimitedNin    *[]bool
	QuantityUnlimitedExists *bool
	// registryId (ParentRef<Registry>) search options
	RegistryIdEq     *string
	RegistryIdIn     *[]string
	RegistryIdNin    *[]string
	RegistryIdExists *bool
	// retailer (string) search options
	RetailerEq     *string
	RetailerNe     *string
	RetailerGt     *string
	RetailerGte    *string
	RetailerLt     *string
	RetailerLte    *string
	RetailerIn     *[]string
	RetailerNin    *[]string
	RetailerExists *bool
	RetailerLike   *string
	RetailerNlike  *string
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
	if o.AffiliateUrlEq != nil {
		elemaffiliateUrlEq0 := o.AffiliateUrlEq
		to.AffiliateUrlEq = elemaffiliateUrlEq0
	}
	if o.AffiliateUrlNe != nil {
		elemaffiliateUrlNe0 := o.AffiliateUrlNe
		to.AffiliateUrlNe = elemaffiliateUrlNe0
	}
	if o.AffiliateUrlGt != nil {
		elemaffiliateUrlGt0 := o.AffiliateUrlGt
		to.AffiliateUrlGt = elemaffiliateUrlGt0
	}
	if o.AffiliateUrlGte != nil {
		elemaffiliateUrlGte0 := o.AffiliateUrlGte
		to.AffiliateUrlGte = elemaffiliateUrlGte0
	}
	if o.AffiliateUrlLt != nil {
		elemaffiliateUrlLt0 := o.AffiliateUrlLt
		to.AffiliateUrlLt = elemaffiliateUrlLt0
	}
	if o.AffiliateUrlLte != nil {
		elemaffiliateUrlLte0 := o.AffiliateUrlLte
		to.AffiliateUrlLte = elemaffiliateUrlLte0
	}
	if o.AffiliateUrlIn != nil {
		elemaffiliateUrlIn0 := make([]string, 0)
		for _, oaffiliateUrlIn0 := range *o.AffiliateUrlIn {
			elemaffiliateUrlIn1 := oaffiliateUrlIn0
			elemaffiliateUrlIn0 = append(elemaffiliateUrlIn0, elemaffiliateUrlIn1)
		}
		to.AffiliateUrlIn = &elemaffiliateUrlIn0
	}
	if o.AffiliateUrlNin != nil {
		elemaffiliateUrlNin0 := make([]string, 0)
		for _, oaffiliateUrlNin0 := range *o.AffiliateUrlNin {
			elemaffiliateUrlNin1 := oaffiliateUrlNin0
			elemaffiliateUrlNin0 = append(elemaffiliateUrlNin0, elemaffiliateUrlNin1)
		}
		to.AffiliateUrlNin = &elemaffiliateUrlNin0
	}
	if o.AffiliateUrlExists != nil {
		elemaffiliateUrlExists0 := o.AffiliateUrlExists
		to.AffiliateUrlExists = elemaffiliateUrlExists0
	}
	if o.AffiliateUrlLike != nil {
		elemaffiliateUrlLike0 := o.AffiliateUrlLike
		to.AffiliateUrlLike = elemaffiliateUrlLike0
	}
	if o.AffiliateUrlNlike != nil {
		elemaffiliateUrlNlike0 := o.AffiliateUrlNlike
		to.AffiliateUrlNlike = elemaffiliateUrlNlike0
	}
	if o.CanonicalUrlEq != nil {
		elemcanonicalUrlEq0 := o.CanonicalUrlEq
		to.CanonicalUrlEq = elemcanonicalUrlEq0
	}
	if o.CanonicalUrlNe != nil {
		elemcanonicalUrlNe0 := o.CanonicalUrlNe
		to.CanonicalUrlNe = elemcanonicalUrlNe0
	}
	if o.CanonicalUrlGt != nil {
		elemcanonicalUrlGt0 := o.CanonicalUrlGt
		to.CanonicalUrlGt = elemcanonicalUrlGt0
	}
	if o.CanonicalUrlGte != nil {
		elemcanonicalUrlGte0 := o.CanonicalUrlGte
		to.CanonicalUrlGte = elemcanonicalUrlGte0
	}
	if o.CanonicalUrlLt != nil {
		elemcanonicalUrlLt0 := o.CanonicalUrlLt
		to.CanonicalUrlLt = elemcanonicalUrlLt0
	}
	if o.CanonicalUrlLte != nil {
		elemcanonicalUrlLte0 := o.CanonicalUrlLte
		to.CanonicalUrlLte = elemcanonicalUrlLte0
	}
	if o.CanonicalUrlIn != nil {
		elemcanonicalUrlIn0 := make([]string, 0)
		for _, ocanonicalUrlIn0 := range *o.CanonicalUrlIn {
			elemcanonicalUrlIn1 := ocanonicalUrlIn0
			elemcanonicalUrlIn0 = append(elemcanonicalUrlIn0, elemcanonicalUrlIn1)
		}
		to.CanonicalUrlIn = &elemcanonicalUrlIn0
	}
	if o.CanonicalUrlNin != nil {
		elemcanonicalUrlNin0 := make([]string, 0)
		for _, ocanonicalUrlNin0 := range *o.CanonicalUrlNin {
			elemcanonicalUrlNin1 := ocanonicalUrlNin0
			elemcanonicalUrlNin0 = append(elemcanonicalUrlNin0, elemcanonicalUrlNin1)
		}
		to.CanonicalUrlNin = &elemcanonicalUrlNin0
	}
	if o.CanonicalUrlExists != nil {
		elemcanonicalUrlExists0 := o.CanonicalUrlExists
		to.CanonicalUrlExists = elemcanonicalUrlExists0
	}
	if o.CanonicalUrlLike != nil {
		elemcanonicalUrlLike0 := o.CanonicalUrlLike
		to.CanonicalUrlLike = elemcanonicalUrlLike0
	}
	if o.CanonicalUrlNlike != nil {
		elemcanonicalUrlNlike0 := o.CanonicalUrlNlike
		to.CanonicalUrlNlike = elemcanonicalUrlNlike0
	}
	if o.CategoryEq != nil {
		elemcategoryEq0 := o.CategoryEq
		to.CategoryEq = elemcategoryEq0
	}
	if o.CategoryNe != nil {
		elemcategoryNe0 := o.CategoryNe
		to.CategoryNe = elemcategoryNe0
	}
	if o.CategoryGt != nil {
		elemcategoryGt0 := o.CategoryGt
		to.CategoryGt = elemcategoryGt0
	}
	if o.CategoryGte != nil {
		elemcategoryGte0 := o.CategoryGte
		to.CategoryGte = elemcategoryGte0
	}
	if o.CategoryLt != nil {
		elemcategoryLt0 := o.CategoryLt
		to.CategoryLt = elemcategoryLt0
	}
	if o.CategoryLte != nil {
		elemcategoryLte0 := o.CategoryLte
		to.CategoryLte = elemcategoryLte0
	}
	if o.CategoryIn != nil {
		elemcategoryIn0 := make([]string, 0)
		for _, ocategoryIn0 := range *o.CategoryIn {
			elemcategoryIn1 := ocategoryIn0
			elemcategoryIn0 = append(elemcategoryIn0, elemcategoryIn1)
		}
		to.CategoryIn = &elemcategoryIn0
	}
	if o.CategoryNin != nil {
		elemcategoryNin0 := make([]string, 0)
		for _, ocategoryNin0 := range *o.CategoryNin {
			elemcategoryNin1 := ocategoryNin0
			elemcategoryNin0 = append(elemcategoryNin0, elemcategoryNin1)
		}
		to.CategoryNin = &elemcategoryNin0
	}
	if o.CategoryExists != nil {
		elemcategoryExists0 := o.CategoryExists
		to.CategoryExists = elemcategoryExists0
	}
	if o.CategoryLike != nil {
		elemcategoryLike0 := o.CategoryLike
		to.CategoryLike = elemcategoryLike0
	}
	if o.CategoryNlike != nil {
		elemcategoryNlike0 := o.CategoryNlike
		to.CategoryNlike = elemcategoryNlike0
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
	if o.ImageBgColorEq != nil {
		elemimageBgColorEq0 := o.ImageBgColorEq
		to.ImageBgColorEq = elemimageBgColorEq0
	}
	if o.ImageBgColorNe != nil {
		elemimageBgColorNe0 := o.ImageBgColorNe
		to.ImageBgColorNe = elemimageBgColorNe0
	}
	if o.ImageBgColorGt != nil {
		elemimageBgColorGt0 := o.ImageBgColorGt
		to.ImageBgColorGt = elemimageBgColorGt0
	}
	if o.ImageBgColorGte != nil {
		elemimageBgColorGte0 := o.ImageBgColorGte
		to.ImageBgColorGte = elemimageBgColorGte0
	}
	if o.ImageBgColorLt != nil {
		elemimageBgColorLt0 := o.ImageBgColorLt
		to.ImageBgColorLt = elemimageBgColorLt0
	}
	if o.ImageBgColorLte != nil {
		elemimageBgColorLte0 := o.ImageBgColorLte
		to.ImageBgColorLte = elemimageBgColorLte0
	}
	if o.ImageBgColorIn != nil {
		elemimageBgColorIn0 := make([]string, 0)
		for _, oimageBgColorIn0 := range *o.ImageBgColorIn {
			elemimageBgColorIn1 := oimageBgColorIn0
			elemimageBgColorIn0 = append(elemimageBgColorIn0, elemimageBgColorIn1)
		}
		to.ImageBgColorIn = &elemimageBgColorIn0
	}
	if o.ImageBgColorNin != nil {
		elemimageBgColorNin0 := make([]string, 0)
		for _, oimageBgColorNin0 := range *o.ImageBgColorNin {
			elemimageBgColorNin1 := oimageBgColorNin0
			elemimageBgColorNin0 = append(elemimageBgColorNin0, elemimageBgColorNin1)
		}
		to.ImageBgColorNin = &elemimageBgColorNin0
	}
	if o.ImageBgColorExists != nil {
		elemimageBgColorExists0 := o.ImageBgColorExists
		to.ImageBgColorExists = elemimageBgColorExists0
	}
	if o.ImageBgColorLike != nil {
		elemimageBgColorLike0 := o.ImageBgColorLike
		to.ImageBgColorLike = elemimageBgColorLike0
	}
	if o.ImageBgColorNlike != nil {
		elemimageBgColorNlike0 := o.ImageBgColorNlike
		to.ImageBgColorNlike = elemimageBgColorNlike0
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
	if o.NoSubstitutesEq != nil {
		elemnoSubstitutesEq0 := o.NoSubstitutesEq
		to.NoSubstitutesEq = elemnoSubstitutesEq0
	}
	if o.NoSubstitutesNe != nil {
		elemnoSubstitutesNe0 := o.NoSubstitutesNe
		to.NoSubstitutesNe = elemnoSubstitutesNe0
	}
	if o.NoSubstitutesGt != nil {
		elemnoSubstitutesGt0 := o.NoSubstitutesGt
		to.NoSubstitutesGt = elemnoSubstitutesGt0
	}
	if o.NoSubstitutesGte != nil {
		elemnoSubstitutesGte0 := o.NoSubstitutesGte
		to.NoSubstitutesGte = elemnoSubstitutesGte0
	}
	if o.NoSubstitutesLt != nil {
		elemnoSubstitutesLt0 := o.NoSubstitutesLt
		to.NoSubstitutesLt = elemnoSubstitutesLt0
	}
	if o.NoSubstitutesLte != nil {
		elemnoSubstitutesLte0 := o.NoSubstitutesLte
		to.NoSubstitutesLte = elemnoSubstitutesLte0
	}
	if o.NoSubstitutesIn != nil {
		elemnoSubstitutesIn0 := make([]bool, 0)
		for _, onoSubstitutesIn0 := range *o.NoSubstitutesIn {
			elemnoSubstitutesIn1 := onoSubstitutesIn0
			elemnoSubstitutesIn0 = append(elemnoSubstitutesIn0, elemnoSubstitutesIn1)
		}
		to.NoSubstitutesIn = &elemnoSubstitutesIn0
	}
	if o.NoSubstitutesNin != nil {
		elemnoSubstitutesNin0 := make([]bool, 0)
		for _, onoSubstitutesNin0 := range *o.NoSubstitutesNin {
			elemnoSubstitutesNin1 := onoSubstitutesNin0
			elemnoSubstitutesNin0 = append(elemnoSubstitutesNin0, elemnoSubstitutesNin1)
		}
		to.NoSubstitutesNin = &elemnoSubstitutesNin0
	}
	if o.NoSubstitutesExists != nil {
		elemnoSubstitutesExists0 := o.NoSubstitutesExists
		to.NoSubstitutesExists = elemnoSubstitutesExists0
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
	if o.OriginalUrlEq != nil {
		elemoriginalUrlEq0 := o.OriginalUrlEq
		to.OriginalUrlEq = elemoriginalUrlEq0
	}
	if o.OriginalUrlNe != nil {
		elemoriginalUrlNe0 := o.OriginalUrlNe
		to.OriginalUrlNe = elemoriginalUrlNe0
	}
	if o.OriginalUrlGt != nil {
		elemoriginalUrlGt0 := o.OriginalUrlGt
		to.OriginalUrlGt = elemoriginalUrlGt0
	}
	if o.OriginalUrlGte != nil {
		elemoriginalUrlGte0 := o.OriginalUrlGte
		to.OriginalUrlGte = elemoriginalUrlGte0
	}
	if o.OriginalUrlLt != nil {
		elemoriginalUrlLt0 := o.OriginalUrlLt
		to.OriginalUrlLt = elemoriginalUrlLt0
	}
	if o.OriginalUrlLte != nil {
		elemoriginalUrlLte0 := o.OriginalUrlLte
		to.OriginalUrlLte = elemoriginalUrlLte0
	}
	if o.OriginalUrlIn != nil {
		elemoriginalUrlIn0 := make([]string, 0)
		for _, ooriginalUrlIn0 := range *o.OriginalUrlIn {
			elemoriginalUrlIn1 := ooriginalUrlIn0
			elemoriginalUrlIn0 = append(elemoriginalUrlIn0, elemoriginalUrlIn1)
		}
		to.OriginalUrlIn = &elemoriginalUrlIn0
	}
	if o.OriginalUrlNin != nil {
		elemoriginalUrlNin0 := make([]string, 0)
		for _, ooriginalUrlNin0 := range *o.OriginalUrlNin {
			elemoriginalUrlNin1 := ooriginalUrlNin0
			elemoriginalUrlNin0 = append(elemoriginalUrlNin0, elemoriginalUrlNin1)
		}
		to.OriginalUrlNin = &elemoriginalUrlNin0
	}
	if o.OriginalUrlExists != nil {
		elemoriginalUrlExists0 := o.OriginalUrlExists
		to.OriginalUrlExists = elemoriginalUrlExists0
	}
	if o.OriginalUrlLike != nil {
		elemoriginalUrlLike0 := o.OriginalUrlLike
		to.OriginalUrlLike = elemoriginalUrlLike0
	}
	if o.OriginalUrlNlike != nil {
		elemoriginalUrlNlike0 := o.OriginalUrlNlike
		to.OriginalUrlNlike = elemoriginalUrlNlike0
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
	if o.ParentItemIdEq != nil {
		elemparentItemIdEq0, err := primitive.ObjectIDFromHex(*o.ParentItemIdEq)
		if err != nil {
			return to, errors.Join(errors.New("invalid o.ParentItemIdEq"), err)
		}
		to.ParentItemIdEq = &elemparentItemIdEq0
	}
	if o.ParentItemIdIn != nil {
		elemparentItemIdIn0 := make([]primitive.ObjectID, 0)
		for _, oparentItemIdIn0 := range *o.ParentItemIdIn {
			elemparentItemIdIn1, err := primitive.ObjectIDFromHex(oparentItemIdIn0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oparentItemIdIn0"), err)
			}
			elemparentItemIdIn0 = append(elemparentItemIdIn0, elemparentItemIdIn1)
		}
		to.ParentItemIdIn = &elemparentItemIdIn0
	}
	if o.ParentItemIdNin != nil {
		elemparentItemIdNin0 := make([]primitive.ObjectID, 0)
		for _, oparentItemIdNin0 := range *o.ParentItemIdNin {
			elemparentItemIdNin1, err := primitive.ObjectIDFromHex(oparentItemIdNin0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oparentItemIdNin0"), err)
			}
			elemparentItemIdNin0 = append(elemparentItemIdNin0, elemparentItemIdNin1)
		}
		to.ParentItemIdNin = &elemparentItemIdNin0
	}
	if o.ParentItemIdExists != nil {
		elemparentItemIdExists0 := o.ParentItemIdExists
		to.ParentItemIdExists = elemparentItemIdExists0
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
	if o.QuantityUnlimitedEq != nil {
		elemquantityUnlimitedEq0 := o.QuantityUnlimitedEq
		to.QuantityUnlimitedEq = elemquantityUnlimitedEq0
	}
	if o.QuantityUnlimitedNe != nil {
		elemquantityUnlimitedNe0 := o.QuantityUnlimitedNe
		to.QuantityUnlimitedNe = elemquantityUnlimitedNe0
	}
	if o.QuantityUnlimitedGt != nil {
		elemquantityUnlimitedGt0 := o.QuantityUnlimitedGt
		to.QuantityUnlimitedGt = elemquantityUnlimitedGt0
	}
	if o.QuantityUnlimitedGte != nil {
		elemquantityUnlimitedGte0 := o.QuantityUnlimitedGte
		to.QuantityUnlimitedGte = elemquantityUnlimitedGte0
	}
	if o.QuantityUnlimitedLt != nil {
		elemquantityUnlimitedLt0 := o.QuantityUnlimitedLt
		to.QuantityUnlimitedLt = elemquantityUnlimitedLt0
	}
	if o.QuantityUnlimitedLte != nil {
		elemquantityUnlimitedLte0 := o.QuantityUnlimitedLte
		to.QuantityUnlimitedLte = elemquantityUnlimitedLte0
	}
	if o.QuantityUnlimitedIn != nil {
		elemquantityUnlimitedIn0 := make([]bool, 0)
		for _, oquantityUnlimitedIn0 := range *o.QuantityUnlimitedIn {
			elemquantityUnlimitedIn1 := oquantityUnlimitedIn0
			elemquantityUnlimitedIn0 = append(elemquantityUnlimitedIn0, elemquantityUnlimitedIn1)
		}
		to.QuantityUnlimitedIn = &elemquantityUnlimitedIn0
	}
	if o.QuantityUnlimitedNin != nil {
		elemquantityUnlimitedNin0 := make([]bool, 0)
		for _, oquantityUnlimitedNin0 := range *o.QuantityUnlimitedNin {
			elemquantityUnlimitedNin1 := oquantityUnlimitedNin0
			elemquantityUnlimitedNin0 = append(elemquantityUnlimitedNin0, elemquantityUnlimitedNin1)
		}
		to.QuantityUnlimitedNin = &elemquantityUnlimitedNin0
	}
	if o.QuantityUnlimitedExists != nil {
		elemquantityUnlimitedExists0 := o.QuantityUnlimitedExists
		to.QuantityUnlimitedExists = elemquantityUnlimitedExists0
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
	if o.RetailerEq != nil {
		elemretailerEq0 := o.RetailerEq
		to.RetailerEq = elemretailerEq0
	}
	if o.RetailerNe != nil {
		elemretailerNe0 := o.RetailerNe
		to.RetailerNe = elemretailerNe0
	}
	if o.RetailerGt != nil {
		elemretailerGt0 := o.RetailerGt
		to.RetailerGt = elemretailerGt0
	}
	if o.RetailerGte != nil {
		elemretailerGte0 := o.RetailerGte
		to.RetailerGte = elemretailerGte0
	}
	if o.RetailerLt != nil {
		elemretailerLt0 := o.RetailerLt
		to.RetailerLt = elemretailerLt0
	}
	if o.RetailerLte != nil {
		elemretailerLte0 := o.RetailerLte
		to.RetailerLte = elemretailerLte0
	}
	if o.RetailerIn != nil {
		elemretailerIn0 := make([]string, 0)
		for _, oretailerIn0 := range *o.RetailerIn {
			elemretailerIn1 := oretailerIn0
			elemretailerIn0 = append(elemretailerIn0, elemretailerIn1)
		}
		to.RetailerIn = &elemretailerIn0
	}
	if o.RetailerNin != nil {
		elemretailerNin0 := make([]string, 0)
		for _, oretailerNin0 := range *o.RetailerNin {
			elemretailerNin1 := oretailerNin0
			elemretailerNin0 = append(elemretailerNin0, elemretailerNin1)
		}
		to.RetailerNin = &elemretailerNin0
	}
	if o.RetailerExists != nil {
		elemretailerExists0 := o.RetailerExists
		to.RetailerExists = elemretailerExists0
	}
	if o.RetailerLike != nil {
		elemretailerLike0 := o.RetailerLike
		to.RetailerLike = elemretailerLike0
	}
	if o.RetailerNlike != nil {
		elemretailerNlike0 := o.RetailerNlike
		to.RetailerNlike = elemretailerNlike0
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
