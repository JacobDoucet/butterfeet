package registry_item

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_item_source"
)

type HTTPRecord struct {
	Id                 *string                 `json:"id,omitempty"`
	Created            *actor_trace.HTTPRecord `json:"created,omitempty"`
	Currency           *string                 `json:"currency,omitempty"`
	Description        *string                 `json:"description,omitempty"`
	ImageUrl           *string                 `json:"imageUrl,omitempty"`
	Notes              *string                 `json:"notes,omitempty"`
	Position           *int                    `json:"position,omitempty"`
	PriceCents         *int                    `json:"priceCents,omitempty"`
	ProductUrl         *string                 `json:"productUrl,omitempty"`
	Quantity           *int                    `json:"quantity,omitempty"`
	RegistryId         *string                 `json:"registryId,omitempty"`
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
		elemregistryId0 := r.RegistryId
		m.RegistryId = *elemregistryId0
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
	if r.ImageUrl != nil {
		p.ImageUrl = true
	}
	if r.Notes != nil {
		p.Notes = true
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
	if r.RegistryId != nil {
		p.RegistryId = true
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
	// registryId (ParentRef<Registry>) search options
	RegistryIdEq     *string   `json:"registryIdEq,omitempty"`
	RegistryIdIn     *[]string `json:"registryIdIn,omitempty"`
	RegistryIdNin    *[]string `json:"registryIdNin,omitempty"`
	RegistryIdExists *bool     `json:"registryIdExists,omitempty"`
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
