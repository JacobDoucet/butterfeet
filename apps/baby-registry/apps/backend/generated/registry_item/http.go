package registry_item

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_item_source"
)

type HTTPRecord struct {
	Id                 *string                 `json:"id,omitempty"`
	AffiliateUrl       *string                 `json:"affiliateUrl,omitempty"`
	CanonicalUrl       *string                 `json:"canonicalUrl,omitempty"`
	Category           *string                 `json:"category,omitempty"`
	Created            *actor_trace.HTTPRecord `json:"created,omitempty"`
	Currency           *string                 `json:"currency,omitempty"`
	Description        *string                 `json:"description,omitempty"`
	ImageBgColor       *string                 `json:"imageBgColor,omitempty"`
	ImageUrl           *string                 `json:"imageUrl,omitempty"`
	NoSubstitutes      *bool                   `json:"noSubstitutes,omitempty"`
	Notes              *string                 `json:"notes,omitempty"`
	OriginalUrl        *string                 `json:"originalUrl,omitempty"`
	OwnerPurchased     *bool                   `json:"ownerPurchased,omitempty"`
	ParentItemId       *string                 `json:"parentItemId,omitempty"`
	Position           *int                    `json:"position,omitempty"`
	PriceCents         *int                    `json:"priceCents,omitempty"`
	ProductUrl         *string                 `json:"productUrl,omitempty"`
	Quantity           *int                    `json:"quantity,omitempty"`
	QuantityUnlimited  *bool                   `json:"quantityUnlimited,omitempty"`
	RegistryId         *string                 `json:"registryId,omitempty"`
	Retailer           *string                 `json:"retailer,omitempty"`
	Source             *enum_item_source.Value `json:"source,omitempty"`
	Title              *string                 `json:"title,omitempty"`
	Updated            *actor_trace.HTTPRecord `json:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.HTTPRecord `json:"updatedByOwnerUser,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id
		m.Id = *elemid0
	}
	if r.AffiliateUrl != nil {
		elemaffiliateUrl0 := r.AffiliateUrl
		m.AffiliateUrl = *elemaffiliateUrl0
	}
	if r.CanonicalUrl != nil {
		elemcanonicalUrl0 := r.CanonicalUrl
		m.CanonicalUrl = *elemcanonicalUrl0
	}
	if r.Category != nil {
		elemcategory0 := r.Category
		m.Category = *elemcategory0
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
	if r.ImageBgColor != nil {
		elemimageBgColor0 := r.ImageBgColor
		m.ImageBgColor = *elemimageBgColor0
	}
	if r.ImageUrl != nil {
		elemimageUrl0 := r.ImageUrl
		m.ImageUrl = *elemimageUrl0
	}
	if r.NoSubstitutes != nil {
		elemnoSubstitutes0 := r.NoSubstitutes
		m.NoSubstitutes = *elemnoSubstitutes0
	}
	if r.Notes != nil {
		elemnotes0 := r.Notes
		m.Notes = *elemnotes0
	}
	if r.OriginalUrl != nil {
		elemoriginalUrl0 := r.OriginalUrl
		m.OriginalUrl = *elemoriginalUrl0
	}
	if r.OwnerPurchased != nil {
		elemownerPurchased0 := r.OwnerPurchased
		m.OwnerPurchased = *elemownerPurchased0
	}
	if r.ParentItemId != nil {
		elemparentItemId0 := r.ParentItemId
		m.ParentItemId = *elemparentItemId0
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
	if r.QuantityUnlimited != nil {
		elemquantityUnlimited0 := r.QuantityUnlimited
		m.QuantityUnlimited = *elemquantityUnlimited0
	}
	if r.RegistryId != nil {
		elemregistryId0 := r.RegistryId
		m.RegistryId = *elemregistryId0
	}
	if r.Retailer != nil {
		elemretailer0 := r.Retailer
		m.Retailer = *elemretailer0
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

func (r *HTTPRecord) ToProjection() (Projection, error) {
	p := Projection{}
	if r.Id != nil {
		p.Id = true
	}
	if r.AffiliateUrl != nil {
		p.AffiliateUrl = true
	}
	if r.CanonicalUrl != nil {
		p.CanonicalUrl = true
	}
	if r.Category != nil {
		p.Category = true
	}
	if r.Created != nil {
		p.Created = true
		p.CreatedFields = actor_trace.NewProjection(true)
	}
	if r.Currency != nil {
		p.Currency = true
	}
	if r.Description != nil {
		p.Description = true
	}
	if r.ImageBgColor != nil {
		p.ImageBgColor = true
	}
	if r.ImageUrl != nil {
		p.ImageUrl = true
	}
	if r.NoSubstitutes != nil {
		p.NoSubstitutes = true
	}
	if r.Notes != nil {
		p.Notes = true
	}
	if r.OriginalUrl != nil {
		p.OriginalUrl = true
	}
	if r.OwnerPurchased != nil {
		p.OwnerPurchased = true
	}
	if r.ParentItemId != nil {
		p.ParentItemId = true
	}
	if r.Position != nil {
		p.Position = true
	}
	if r.PriceCents != nil {
		p.PriceCents = true
	}
	if r.ProductUrl != nil {
		p.ProductUrl = true
	}
	if r.Quantity != nil {
		p.Quantity = true
	}
	if r.QuantityUnlimited != nil {
		p.QuantityUnlimited = true
	}
	if r.RegistryId != nil {
		p.RegistryId = true
	}
	if r.Retailer != nil {
		p.Retailer = true
	}
	if r.Source != nil {
		p.Source = true
	}
	if r.Title != nil {
		p.Title = true
	}
	if r.Updated != nil {
		p.Updated = true
		p.UpdatedFields = actor_trace.NewProjection(true)
	}
	if r.UpdatedByOwnerUser != nil {
		p.UpdatedByOwnerUser = true
		p.UpdatedByOwnerUserFields = actor_trace.NewProjection(true)
	}
	return p, nil
}

type HTTPSelectByIdQuery struct {
	Id string `json:"id"`
}

type HTTPWhereClause struct {
	// id (Ref<RegistryItem>) search options
	IdEq     *string   `json:"idEq,omitempty"`
	IdIn     *[]string `json:"idIn,omitempty"`
	IdNin    *[]string `json:"idNin,omitempty"`
	IdExists *bool     `json:"idExists,omitempty"`
	// affiliateUrl (string) search options
	AffiliateUrlEq     *string   `json:"affiliateUrlEq,omitempty"`
	AffiliateUrlNe     *string   `json:"affiliateUrlNe,omitempty"`
	AffiliateUrlGt     *string   `json:"affiliateUrlGt,omitempty"`
	AffiliateUrlGte    *string   `json:"affiliateUrlGte,omitempty"`
	AffiliateUrlLt     *string   `json:"affiliateUrlLt,omitempty"`
	AffiliateUrlLte    *string   `json:"affiliateUrlLte,omitempty"`
	AffiliateUrlIn     *[]string `json:"affiliateUrlIn,omitempty"`
	AffiliateUrlNin    *[]string `json:"affiliateUrlNin,omitempty"`
	AffiliateUrlExists *bool     `json:"affiliateUrlExists,omitempty"`
	AffiliateUrlLike   *string   `json:"affiliateUrlLike,omitempty"`
	AffiliateUrlNlike  *string   `json:"affiliateUrlNlike,omitempty"`
	// canonicalUrl (string) search options
	CanonicalUrlEq     *string   `json:"canonicalUrlEq,omitempty"`
	CanonicalUrlNe     *string   `json:"canonicalUrlNe,omitempty"`
	CanonicalUrlGt     *string   `json:"canonicalUrlGt,omitempty"`
	CanonicalUrlGte    *string   `json:"canonicalUrlGte,omitempty"`
	CanonicalUrlLt     *string   `json:"canonicalUrlLt,omitempty"`
	CanonicalUrlLte    *string   `json:"canonicalUrlLte,omitempty"`
	CanonicalUrlIn     *[]string `json:"canonicalUrlIn,omitempty"`
	CanonicalUrlNin    *[]string `json:"canonicalUrlNin,omitempty"`
	CanonicalUrlExists *bool     `json:"canonicalUrlExists,omitempty"`
	CanonicalUrlLike   *string   `json:"canonicalUrlLike,omitempty"`
	CanonicalUrlNlike  *string   `json:"canonicalUrlNlike,omitempty"`
	// category (string) search options
	CategoryEq     *string   `json:"categoryEq,omitempty"`
	CategoryNe     *string   `json:"categoryNe,omitempty"`
	CategoryGt     *string   `json:"categoryGt,omitempty"`
	CategoryGte    *string   `json:"categoryGte,omitempty"`
	CategoryLt     *string   `json:"categoryLt,omitempty"`
	CategoryLte    *string   `json:"categoryLte,omitempty"`
	CategoryIn     *[]string `json:"categoryIn,omitempty"`
	CategoryNin    *[]string `json:"categoryNin,omitempty"`
	CategoryExists *bool     `json:"categoryExists,omitempty"`
	CategoryLike   *string   `json:"categoryLike,omitempty"`
	CategoryNlike  *string   `json:"categoryNlike,omitempty"`
	// created (ActorTrace) search options
	Created *actor_trace.HTTPWhereClause `json:"created,omitempty"`
	// currency (string) search options
	CurrencyEq     *string   `json:"currencyEq,omitempty"`
	CurrencyNe     *string   `json:"currencyNe,omitempty"`
	CurrencyGt     *string   `json:"currencyGt,omitempty"`
	CurrencyGte    *string   `json:"currencyGte,omitempty"`
	CurrencyLt     *string   `json:"currencyLt,omitempty"`
	CurrencyLte    *string   `json:"currencyLte,omitempty"`
	CurrencyIn     *[]string `json:"currencyIn,omitempty"`
	CurrencyNin    *[]string `json:"currencyNin,omitempty"`
	CurrencyExists *bool     `json:"currencyExists,omitempty"`
	CurrencyLike   *string   `json:"currencyLike,omitempty"`
	CurrencyNlike  *string   `json:"currencyNlike,omitempty"`
	// description (string) search options
	DescriptionEq     *string   `json:"descriptionEq,omitempty"`
	DescriptionNe     *string   `json:"descriptionNe,omitempty"`
	DescriptionGt     *string   `json:"descriptionGt,omitempty"`
	DescriptionGte    *string   `json:"descriptionGte,omitempty"`
	DescriptionLt     *string   `json:"descriptionLt,omitempty"`
	DescriptionLte    *string   `json:"descriptionLte,omitempty"`
	DescriptionIn     *[]string `json:"descriptionIn,omitempty"`
	DescriptionNin    *[]string `json:"descriptionNin,omitempty"`
	DescriptionExists *bool     `json:"descriptionExists,omitempty"`
	DescriptionLike   *string   `json:"descriptionLike,omitempty"`
	DescriptionNlike  *string   `json:"descriptionNlike,omitempty"`
	// imageBgColor (string) search options
	ImageBgColorEq     *string   `json:"imageBgColorEq,omitempty"`
	ImageBgColorNe     *string   `json:"imageBgColorNe,omitempty"`
	ImageBgColorGt     *string   `json:"imageBgColorGt,omitempty"`
	ImageBgColorGte    *string   `json:"imageBgColorGte,omitempty"`
	ImageBgColorLt     *string   `json:"imageBgColorLt,omitempty"`
	ImageBgColorLte    *string   `json:"imageBgColorLte,omitempty"`
	ImageBgColorIn     *[]string `json:"imageBgColorIn,omitempty"`
	ImageBgColorNin    *[]string `json:"imageBgColorNin,omitempty"`
	ImageBgColorExists *bool     `json:"imageBgColorExists,omitempty"`
	ImageBgColorLike   *string   `json:"imageBgColorLike,omitempty"`
	ImageBgColorNlike  *string   `json:"imageBgColorNlike,omitempty"`
	// imageUrl (string) search options
	ImageUrlEq     *string   `json:"imageUrlEq,omitempty"`
	ImageUrlNe     *string   `json:"imageUrlNe,omitempty"`
	ImageUrlGt     *string   `json:"imageUrlGt,omitempty"`
	ImageUrlGte    *string   `json:"imageUrlGte,omitempty"`
	ImageUrlLt     *string   `json:"imageUrlLt,omitempty"`
	ImageUrlLte    *string   `json:"imageUrlLte,omitempty"`
	ImageUrlIn     *[]string `json:"imageUrlIn,omitempty"`
	ImageUrlNin    *[]string `json:"imageUrlNin,omitempty"`
	ImageUrlExists *bool     `json:"imageUrlExists,omitempty"`
	ImageUrlLike   *string   `json:"imageUrlLike,omitempty"`
	ImageUrlNlike  *string   `json:"imageUrlNlike,omitempty"`
	// noSubstitutes (bool) search options
	NoSubstitutesEq     *bool   `json:"noSubstitutesEq,omitempty"`
	NoSubstitutesNe     *bool   `json:"noSubstitutesNe,omitempty"`
	NoSubstitutesGt     *bool   `json:"noSubstitutesGt,omitempty"`
	NoSubstitutesGte    *bool   `json:"noSubstitutesGte,omitempty"`
	NoSubstitutesLt     *bool   `json:"noSubstitutesLt,omitempty"`
	NoSubstitutesLte    *bool   `json:"noSubstitutesLte,omitempty"`
	NoSubstitutesIn     *[]bool `json:"noSubstitutesIn,omitempty"`
	NoSubstitutesNin    *[]bool `json:"noSubstitutesNin,omitempty"`
	NoSubstitutesExists *bool   `json:"noSubstitutesExists,omitempty"`
	// notes (string) search options
	NotesEq     *string   `json:"notesEq,omitempty"`
	NotesNe     *string   `json:"notesNe,omitempty"`
	NotesGt     *string   `json:"notesGt,omitempty"`
	NotesGte    *string   `json:"notesGte,omitempty"`
	NotesLt     *string   `json:"notesLt,omitempty"`
	NotesLte    *string   `json:"notesLte,omitempty"`
	NotesIn     *[]string `json:"notesIn,omitempty"`
	NotesNin    *[]string `json:"notesNin,omitempty"`
	NotesExists *bool     `json:"notesExists,omitempty"`
	NotesLike   *string   `json:"notesLike,omitempty"`
	NotesNlike  *string   `json:"notesNlike,omitempty"`
	// originalUrl (string) search options
	OriginalUrlEq     *string   `json:"originalUrlEq,omitempty"`
	OriginalUrlNe     *string   `json:"originalUrlNe,omitempty"`
	OriginalUrlGt     *string   `json:"originalUrlGt,omitempty"`
	OriginalUrlGte    *string   `json:"originalUrlGte,omitempty"`
	OriginalUrlLt     *string   `json:"originalUrlLt,omitempty"`
	OriginalUrlLte    *string   `json:"originalUrlLte,omitempty"`
	OriginalUrlIn     *[]string `json:"originalUrlIn,omitempty"`
	OriginalUrlNin    *[]string `json:"originalUrlNin,omitempty"`
	OriginalUrlExists *bool     `json:"originalUrlExists,omitempty"`
	OriginalUrlLike   *string   `json:"originalUrlLike,omitempty"`
	OriginalUrlNlike  *string   `json:"originalUrlNlike,omitempty"`
	// ownerPurchased (bool) search options
	OwnerPurchasedEq     *bool   `json:"ownerPurchasedEq,omitempty"`
	OwnerPurchasedNe     *bool   `json:"ownerPurchasedNe,omitempty"`
	OwnerPurchasedGt     *bool   `json:"ownerPurchasedGt,omitempty"`
	OwnerPurchasedGte    *bool   `json:"ownerPurchasedGte,omitempty"`
	OwnerPurchasedLt     *bool   `json:"ownerPurchasedLt,omitempty"`
	OwnerPurchasedLte    *bool   `json:"ownerPurchasedLte,omitempty"`
	OwnerPurchasedIn     *[]bool `json:"ownerPurchasedIn,omitempty"`
	OwnerPurchasedNin    *[]bool `json:"ownerPurchasedNin,omitempty"`
	OwnerPurchasedExists *bool   `json:"ownerPurchasedExists,omitempty"`
	// parentItemId (Ref<RegistryItem>) search options
	ParentItemIdEq     *string   `json:"parentItemIdEq,omitempty"`
	ParentItemIdIn     *[]string `json:"parentItemIdIn,omitempty"`
	ParentItemIdNin    *[]string `json:"parentItemIdNin,omitempty"`
	ParentItemIdExists *bool     `json:"parentItemIdExists,omitempty"`
	// position (int) search options
	PositionEq     *int   `json:"positionEq,omitempty"`
	PositionNe     *int   `json:"positionNe,omitempty"`
	PositionGt     *int   `json:"positionGt,omitempty"`
	PositionGte    *int   `json:"positionGte,omitempty"`
	PositionLt     *int   `json:"positionLt,omitempty"`
	PositionLte    *int   `json:"positionLte,omitempty"`
	PositionIn     *[]int `json:"positionIn,omitempty"`
	PositionNin    *[]int `json:"positionNin,omitempty"`
	PositionExists *bool  `json:"positionExists,omitempty"`
	// priceCents (int) search options
	PriceCentsEq     *int   `json:"priceCentsEq,omitempty"`
	PriceCentsNe     *int   `json:"priceCentsNe,omitempty"`
	PriceCentsGt     *int   `json:"priceCentsGt,omitempty"`
	PriceCentsGte    *int   `json:"priceCentsGte,omitempty"`
	PriceCentsLt     *int   `json:"priceCentsLt,omitempty"`
	PriceCentsLte    *int   `json:"priceCentsLte,omitempty"`
	PriceCentsIn     *[]int `json:"priceCentsIn,omitempty"`
	PriceCentsNin    *[]int `json:"priceCentsNin,omitempty"`
	PriceCentsExists *bool  `json:"priceCentsExists,omitempty"`
	// productUrl (string) search options
	ProductUrlEq     *string   `json:"productUrlEq,omitempty"`
	ProductUrlNe     *string   `json:"productUrlNe,omitempty"`
	ProductUrlGt     *string   `json:"productUrlGt,omitempty"`
	ProductUrlGte    *string   `json:"productUrlGte,omitempty"`
	ProductUrlLt     *string   `json:"productUrlLt,omitempty"`
	ProductUrlLte    *string   `json:"productUrlLte,omitempty"`
	ProductUrlIn     *[]string `json:"productUrlIn,omitempty"`
	ProductUrlNin    *[]string `json:"productUrlNin,omitempty"`
	ProductUrlExists *bool     `json:"productUrlExists,omitempty"`
	ProductUrlLike   *string   `json:"productUrlLike,omitempty"`
	ProductUrlNlike  *string   `json:"productUrlNlike,omitempty"`
	// quantity (int) search options
	QuantityEq     *int   `json:"quantityEq,omitempty"`
	QuantityNe     *int   `json:"quantityNe,omitempty"`
	QuantityGt     *int   `json:"quantityGt,omitempty"`
	QuantityGte    *int   `json:"quantityGte,omitempty"`
	QuantityLt     *int   `json:"quantityLt,omitempty"`
	QuantityLte    *int   `json:"quantityLte,omitempty"`
	QuantityIn     *[]int `json:"quantityIn,omitempty"`
	QuantityNin    *[]int `json:"quantityNin,omitempty"`
	QuantityExists *bool  `json:"quantityExists,omitempty"`
	// quantityUnlimited (bool) search options
	QuantityUnlimitedEq     *bool   `json:"quantityUnlimitedEq,omitempty"`
	QuantityUnlimitedNe     *bool   `json:"quantityUnlimitedNe,omitempty"`
	QuantityUnlimitedGt     *bool   `json:"quantityUnlimitedGt,omitempty"`
	QuantityUnlimitedGte    *bool   `json:"quantityUnlimitedGte,omitempty"`
	QuantityUnlimitedLt     *bool   `json:"quantityUnlimitedLt,omitempty"`
	QuantityUnlimitedLte    *bool   `json:"quantityUnlimitedLte,omitempty"`
	QuantityUnlimitedIn     *[]bool `json:"quantityUnlimitedIn,omitempty"`
	QuantityUnlimitedNin    *[]bool `json:"quantityUnlimitedNin,omitempty"`
	QuantityUnlimitedExists *bool   `json:"quantityUnlimitedExists,omitempty"`
	// registryId (ParentRef<Registry>) search options
	RegistryIdEq     *string   `json:"registryIdEq,omitempty"`
	RegistryIdIn     *[]string `json:"registryIdIn,omitempty"`
	RegistryIdNin    *[]string `json:"registryIdNin,omitempty"`
	RegistryIdExists *bool     `json:"registryIdExists,omitempty"`
	// retailer (string) search options
	RetailerEq     *string   `json:"retailerEq,omitempty"`
	RetailerNe     *string   `json:"retailerNe,omitempty"`
	RetailerGt     *string   `json:"retailerGt,omitempty"`
	RetailerGte    *string   `json:"retailerGte,omitempty"`
	RetailerLt     *string   `json:"retailerLt,omitempty"`
	RetailerLte    *string   `json:"retailerLte,omitempty"`
	RetailerIn     *[]string `json:"retailerIn,omitempty"`
	RetailerNin    *[]string `json:"retailerNin,omitempty"`
	RetailerExists *bool     `json:"retailerExists,omitempty"`
	RetailerLike   *string   `json:"retailerLike,omitempty"`
	RetailerNlike  *string   `json:"retailerNlike,omitempty"`
	// source (ItemSource) search options
	SourceEq     *enum_item_source.Value   `json:"sourceEq,omitempty"`
	SourceNe     *enum_item_source.Value   `json:"sourceNe,omitempty"`
	SourceGt     *enum_item_source.Value   `json:"sourceGt,omitempty"`
	SourceGte    *enum_item_source.Value   `json:"sourceGte,omitempty"`
	SourceLt     *enum_item_source.Value   `json:"sourceLt,omitempty"`
	SourceLte    *enum_item_source.Value   `json:"sourceLte,omitempty"`
	SourceIn     *[]enum_item_source.Value `json:"sourceIn,omitempty"`
	SourceNin    *[]enum_item_source.Value `json:"sourceNin,omitempty"`
	SourceExists *bool                     `json:"sourceExists,omitempty"`
	// title (string) search options
	TitleEq     *string   `json:"titleEq,omitempty"`
	TitleNe     *string   `json:"titleNe,omitempty"`
	TitleGt     *string   `json:"titleGt,omitempty"`
	TitleGte    *string   `json:"titleGte,omitempty"`
	TitleLt     *string   `json:"titleLt,omitempty"`
	TitleLte    *string   `json:"titleLte,omitempty"`
	TitleIn     *[]string `json:"titleIn,omitempty"`
	TitleNin    *[]string `json:"titleNin,omitempty"`
	TitleExists *bool     `json:"titleExists,omitempty"`
	TitleLike   *string   `json:"titleLike,omitempty"`
	TitleNlike  *string   `json:"titleNlike,omitempty"`
	// updated (ActorTrace) search options
	Updated *actor_trace.HTTPWhereClause `json:"updated,omitempty"`
	// updatedByOwnerUser (ActorTrace) search options
	UpdatedByOwnerUser *actor_trace.HTTPWhereClause `json:"updatedByOwnerUser,omitempty"`
}

func (o HTTPSelectByIdQuery) ToSelectByIdQuery() (SelectByIdQuery, error) {
	to := SelectByIdQuery{}
	elemid0 := o.Id
	to.Id = elemid0
	return to, nil
}

func (o HTTPWhereClause) ToWhereClause() (WhereClause, error) {
	to := WhereClause{}
	if o.IdEq != nil {
		elemidEq0 := o.IdEq
		to.IdEq = elemidEq0
	}
	if o.IdIn != nil {
		elemidIn0 := make([]string, 0)
		for _, oidIn0 := range *o.IdIn {
			elemidIn1 := oidIn0
			elemidIn0 = append(elemidIn0, elemidIn1)
		}
		to.IdIn = &elemidIn0
	}
	if o.IdNin != nil {
		elemidNin0 := make([]string, 0)
		for _, oidNin0 := range *o.IdNin {
			elemidNin1 := oidNin0
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
		elemcreated0, err := o.Created.ToWhereClause()
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
		elemparentItemIdEq0 := o.ParentItemIdEq
		to.ParentItemIdEq = elemparentItemIdEq0
	}
	if o.ParentItemIdIn != nil {
		elemparentItemIdIn0 := make([]string, 0)
		for _, oparentItemIdIn0 := range *o.ParentItemIdIn {
			elemparentItemIdIn1 := oparentItemIdIn0
			elemparentItemIdIn0 = append(elemparentItemIdIn0, elemparentItemIdIn1)
		}
		to.ParentItemIdIn = &elemparentItemIdIn0
	}
	if o.ParentItemIdNin != nil {
		elemparentItemIdNin0 := make([]string, 0)
		for _, oparentItemIdNin0 := range *o.ParentItemIdNin {
			elemparentItemIdNin1 := oparentItemIdNin0
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
		elemregistryIdEq0 := o.RegistryIdEq
		to.RegistryIdEq = elemregistryIdEq0
	}
	if o.RegistryIdIn != nil {
		elemregistryIdIn0 := make([]string, 0)
		for _, oregistryIdIn0 := range *o.RegistryIdIn {
			elemregistryIdIn1 := oregistryIdIn0
			elemregistryIdIn0 = append(elemregistryIdIn0, elemregistryIdIn1)
		}
		to.RegistryIdIn = &elemregistryIdIn0
	}
	if o.RegistryIdNin != nil {
		elemregistryIdNin0 := make([]string, 0)
		for _, oregistryIdNin0 := range *o.RegistryIdNin {
			elemregistryIdNin1 := oregistryIdNin0
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
		elemupdated0, err := o.Updated.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.Updated = &elemupdated0
	}
	if o.UpdatedByOwnerUser != nil {
		elemupdatedByOwnerUser0, err := o.UpdatedByOwnerUser.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.UpdatedByOwnerUser = &elemupdatedByOwnerUser0
	}
	return to, nil
}

type HTTPSortParams struct {
	CreatedAt  *int8 `json:"createdAt,omitempty"`
	Position   *int8 `json:"position,omitempty"`
	RegistryId *int8 `json:"registryId,omitempty"`
	UpdatedAt  *int8 `json:"updatedAt,omitempty"`
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	if s.CreatedAt != nil {
		to.CreatedAt = *s.CreatedAt
	}
	if s.Position != nil {
		to.Position = *s.Position
	}
	if s.RegistryId != nil {
		to.RegistryId = *s.RegistryId
	}
	if s.UpdatedAt != nil {
		to.UpdatedAt = *s.UpdatedAt
	}
	return to
}
