package registry

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_address_access_mode"
	"time"
)

type HTTPRecord struct {
	Id                    *string                         `json:"id,omitempty"`
	AddressAccessMode     *enum_address_access_mode.Value `json:"addressAccessMode,omitempty"`
	CoverImageUrl         *string                         `json:"coverImageUrl,omitempty"`
	Created               *actor_trace.HTTPRecord         `json:"created,omitempty"`
	DueDate               *time.Time                      `json:"dueDate,omitempty"`
	IsPublic              *bool                           `json:"isPublic,omitempty"`
	OwnerId               *string                         `json:"ownerId,omitempty"`
	ParentNames           *string                         `json:"parentNames,omitempty"`
	ShippingCity          *string                         `json:"shippingCity,omitempty"`
	ShippingCountry       *string                         `json:"shippingCountry,omitempty"`
	ShippingDeliveryNotes *string                         `json:"shippingDeliveryNotes,omitempty"`
	ShippingLine1         *string                         `json:"shippingLine1,omitempty"`
	ShippingLine2         *string                         `json:"shippingLine2,omitempty"`
	ShippingPolicyVersion *int                            `json:"shippingPolicyVersion,omitempty"`
	ShippingPostalCode    *string                         `json:"shippingPostalCode,omitempty"`
	ShippingRecipientName *string                         `json:"shippingRecipientName,omitempty"`
	ShippingRegion        *string                         `json:"shippingRegion,omitempty"`
	Slug                  *string                         `json:"slug,omitempty"`
	ThemeColor            *string                         `json:"themeColor,omitempty"`
	Title                 *string                         `json:"title,omitempty"`
	Updated               *actor_trace.HTTPRecord         `json:"updated,omitempty"`
	UpdatedByOwnerUser    *actor_trace.HTTPRecord         `json:"updatedByOwnerUser,omitempty"`
	WelcomeMessage        *string                         `json:"welcomeMessage,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id
		m.Id = *elemid0
	}
	if r.AddressAccessMode != nil {
		elemaddressAccessMode0 := r.AddressAccessMode
		m.AddressAccessMode = *elemaddressAccessMode0
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
		elemownerId0 := r.OwnerId
		m.OwnerId = *elemownerId0
	}
	if r.ParentNames != nil {
		elemparentNames0 := r.ParentNames
		m.ParentNames = *elemparentNames0
	}
	if r.ShippingCity != nil {
		elemshippingCity0 := r.ShippingCity
		m.ShippingCity = *elemshippingCity0
	}
	if r.ShippingCountry != nil {
		elemshippingCountry0 := r.ShippingCountry
		m.ShippingCountry = *elemshippingCountry0
	}
	if r.ShippingDeliveryNotes != nil {
		elemshippingDeliveryNotes0 := r.ShippingDeliveryNotes
		m.ShippingDeliveryNotes = *elemshippingDeliveryNotes0
	}
	if r.ShippingLine1 != nil {
		elemshippingLine10 := r.ShippingLine1
		m.ShippingLine1 = *elemshippingLine10
	}
	if r.ShippingLine2 != nil {
		elemshippingLine20 := r.ShippingLine2
		m.ShippingLine2 = *elemshippingLine20
	}
	if r.ShippingPolicyVersion != nil {
		elemshippingPolicyVersion0 := r.ShippingPolicyVersion
		m.ShippingPolicyVersion = *elemshippingPolicyVersion0
	}
	if r.ShippingPostalCode != nil {
		elemshippingPostalCode0 := r.ShippingPostalCode
		m.ShippingPostalCode = *elemshippingPostalCode0
	}
	if r.ShippingRecipientName != nil {
		elemshippingRecipientName0 := r.ShippingRecipientName
		m.ShippingRecipientName = *elemshippingRecipientName0
	}
	if r.ShippingRegion != nil {
		elemshippingRegion0 := r.ShippingRegion
		m.ShippingRegion = *elemshippingRegion0
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

func (r *HTTPRecord) ToProjection() (Projection, error) {
	p := Projection{}
	if r.Id != nil {
		p.Id = true
	}
	if r.AddressAccessMode != nil {
		p.AddressAccessMode = true
	}
	if r.CoverImageUrl != nil {
		p.CoverImageUrl = true
	}
	if r.Created != nil {
		p.Created = true
		p.CreatedFields = actor_trace.NewProjection(true)
	}
	if r.DueDate != nil {
		p.DueDate = true
	}
	if r.IsPublic != nil {
		p.IsPublic = true
	}
	if r.OwnerId != nil {
		p.OwnerId = true
	}
	if r.ParentNames != nil {
		p.ParentNames = true
	}
	if r.ShippingCity != nil {
		p.ShippingCity = true
	}
	if r.ShippingCountry != nil {
		p.ShippingCountry = true
	}
	if r.ShippingDeliveryNotes != nil {
		p.ShippingDeliveryNotes = true
	}
	if r.ShippingLine1 != nil {
		p.ShippingLine1 = true
	}
	if r.ShippingLine2 != nil {
		p.ShippingLine2 = true
	}
	if r.ShippingPolicyVersion != nil {
		p.ShippingPolicyVersion = true
	}
	if r.ShippingPostalCode != nil {
		p.ShippingPostalCode = true
	}
	if r.ShippingRecipientName != nil {
		p.ShippingRecipientName = true
	}
	if r.ShippingRegion != nil {
		p.ShippingRegion = true
	}
	if r.Slug != nil {
		p.Slug = true
	}
	if r.ThemeColor != nil {
		p.ThemeColor = true
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
	if r.WelcomeMessage != nil {
		p.WelcomeMessage = true
	}
	return p, nil
}

type HTTPSelectByIdQuery struct {
	Id string `json:"id"`
}
type HTTPSelectBySlugUniqueQuery struct {
	Slug string `json:"slug"`
}

type HTTPWhereClause struct {
	// id (Ref<Registry>) search options
	IdEq     *string   `json:"idEq,omitempty"`
	IdIn     *[]string `json:"idIn,omitempty"`
	IdNin    *[]string `json:"idNin,omitempty"`
	IdExists *bool     `json:"idExists,omitempty"`
	// addressAccessMode (AddressAccessMode) search options
	AddressAccessModeEq     *enum_address_access_mode.Value   `json:"addressAccessModeEq,omitempty"`
	AddressAccessModeNe     *enum_address_access_mode.Value   `json:"addressAccessModeNe,omitempty"`
	AddressAccessModeGt     *enum_address_access_mode.Value   `json:"addressAccessModeGt,omitempty"`
	AddressAccessModeGte    *enum_address_access_mode.Value   `json:"addressAccessModeGte,omitempty"`
	AddressAccessModeLt     *enum_address_access_mode.Value   `json:"addressAccessModeLt,omitempty"`
	AddressAccessModeLte    *enum_address_access_mode.Value   `json:"addressAccessModeLte,omitempty"`
	AddressAccessModeIn     *[]enum_address_access_mode.Value `json:"addressAccessModeIn,omitempty"`
	AddressAccessModeNin    *[]enum_address_access_mode.Value `json:"addressAccessModeNin,omitempty"`
	AddressAccessModeExists *bool                             `json:"addressAccessModeExists,omitempty"`
	// coverImageUrl (string) search options
	CoverImageUrlEq     *string   `json:"coverImageUrlEq,omitempty"`
	CoverImageUrlNe     *string   `json:"coverImageUrlNe,omitempty"`
	CoverImageUrlGt     *string   `json:"coverImageUrlGt,omitempty"`
	CoverImageUrlGte    *string   `json:"coverImageUrlGte,omitempty"`
	CoverImageUrlLt     *string   `json:"coverImageUrlLt,omitempty"`
	CoverImageUrlLte    *string   `json:"coverImageUrlLte,omitempty"`
	CoverImageUrlIn     *[]string `json:"coverImageUrlIn,omitempty"`
	CoverImageUrlNin    *[]string `json:"coverImageUrlNin,omitempty"`
	CoverImageUrlExists *bool     `json:"coverImageUrlExists,omitempty"`
	CoverImageUrlLike   *string   `json:"coverImageUrlLike,omitempty"`
	CoverImageUrlNlike  *string   `json:"coverImageUrlNlike,omitempty"`
	// created (ActorTrace) search options
	Created *actor_trace.HTTPWhereClause `json:"created,omitempty"`
	// dueDate (timestamp) search options
	DueDateEq     *time.Time   `json:"dueDateEq,omitempty"`
	DueDateNe     *time.Time   `json:"dueDateNe,omitempty"`
	DueDateGt     *time.Time   `json:"dueDateGt,omitempty"`
	DueDateGte    *time.Time   `json:"dueDateGte,omitempty"`
	DueDateLt     *time.Time   `json:"dueDateLt,omitempty"`
	DueDateLte    *time.Time   `json:"dueDateLte,omitempty"`
	DueDateIn     *[]time.Time `json:"dueDateIn,omitempty"`
	DueDateNin    *[]time.Time `json:"dueDateNin,omitempty"`
	DueDateExists *bool        `json:"dueDateExists,omitempty"`
	// isPublic (bool) search options
	IsPublicEq     *bool   `json:"isPublicEq,omitempty"`
	IsPublicNe     *bool   `json:"isPublicNe,omitempty"`
	IsPublicGt     *bool   `json:"isPublicGt,omitempty"`
	IsPublicGte    *bool   `json:"isPublicGte,omitempty"`
	IsPublicLt     *bool   `json:"isPublicLt,omitempty"`
	IsPublicLte    *bool   `json:"isPublicLte,omitempty"`
	IsPublicIn     *[]bool `json:"isPublicIn,omitempty"`
	IsPublicNin    *[]bool `json:"isPublicNin,omitempty"`
	IsPublicExists *bool   `json:"isPublicExists,omitempty"`
	// ownerId (Ref<OwnerUser>) search options
	OwnerIdEq     *string   `json:"ownerIdEq,omitempty"`
	OwnerIdIn     *[]string `json:"ownerIdIn,omitempty"`
	OwnerIdNin    *[]string `json:"ownerIdNin,omitempty"`
	OwnerIdExists *bool     `json:"ownerIdExists,omitempty"`
	// parentNames (string) search options
	ParentNamesEq     *string   `json:"parentNamesEq,omitempty"`
	ParentNamesNe     *string   `json:"parentNamesNe,omitempty"`
	ParentNamesGt     *string   `json:"parentNamesGt,omitempty"`
	ParentNamesGte    *string   `json:"parentNamesGte,omitempty"`
	ParentNamesLt     *string   `json:"parentNamesLt,omitempty"`
	ParentNamesLte    *string   `json:"parentNamesLte,omitempty"`
	ParentNamesIn     *[]string `json:"parentNamesIn,omitempty"`
	ParentNamesNin    *[]string `json:"parentNamesNin,omitempty"`
	ParentNamesExists *bool     `json:"parentNamesExists,omitempty"`
	ParentNamesLike   *string   `json:"parentNamesLike,omitempty"`
	ParentNamesNlike  *string   `json:"parentNamesNlike,omitempty"`
	// shippingCity (string) search options
	ShippingCityEq     *string   `json:"shippingCityEq,omitempty"`
	ShippingCityNe     *string   `json:"shippingCityNe,omitempty"`
	ShippingCityGt     *string   `json:"shippingCityGt,omitempty"`
	ShippingCityGte    *string   `json:"shippingCityGte,omitempty"`
	ShippingCityLt     *string   `json:"shippingCityLt,omitempty"`
	ShippingCityLte    *string   `json:"shippingCityLte,omitempty"`
	ShippingCityIn     *[]string `json:"shippingCityIn,omitempty"`
	ShippingCityNin    *[]string `json:"shippingCityNin,omitempty"`
	ShippingCityExists *bool     `json:"shippingCityExists,omitempty"`
	ShippingCityLike   *string   `json:"shippingCityLike,omitempty"`
	ShippingCityNlike  *string   `json:"shippingCityNlike,omitempty"`
	// shippingCountry (string) search options
	ShippingCountryEq     *string   `json:"shippingCountryEq,omitempty"`
	ShippingCountryNe     *string   `json:"shippingCountryNe,omitempty"`
	ShippingCountryGt     *string   `json:"shippingCountryGt,omitempty"`
	ShippingCountryGte    *string   `json:"shippingCountryGte,omitempty"`
	ShippingCountryLt     *string   `json:"shippingCountryLt,omitempty"`
	ShippingCountryLte    *string   `json:"shippingCountryLte,omitempty"`
	ShippingCountryIn     *[]string `json:"shippingCountryIn,omitempty"`
	ShippingCountryNin    *[]string `json:"shippingCountryNin,omitempty"`
	ShippingCountryExists *bool     `json:"shippingCountryExists,omitempty"`
	ShippingCountryLike   *string   `json:"shippingCountryLike,omitempty"`
	ShippingCountryNlike  *string   `json:"shippingCountryNlike,omitempty"`
	// shippingDeliveryNotes (string) search options
	ShippingDeliveryNotesEq     *string   `json:"shippingDeliveryNotesEq,omitempty"`
	ShippingDeliveryNotesNe     *string   `json:"shippingDeliveryNotesNe,omitempty"`
	ShippingDeliveryNotesGt     *string   `json:"shippingDeliveryNotesGt,omitempty"`
	ShippingDeliveryNotesGte    *string   `json:"shippingDeliveryNotesGte,omitempty"`
	ShippingDeliveryNotesLt     *string   `json:"shippingDeliveryNotesLt,omitempty"`
	ShippingDeliveryNotesLte    *string   `json:"shippingDeliveryNotesLte,omitempty"`
	ShippingDeliveryNotesIn     *[]string `json:"shippingDeliveryNotesIn,omitempty"`
	ShippingDeliveryNotesNin    *[]string `json:"shippingDeliveryNotesNin,omitempty"`
	ShippingDeliveryNotesExists *bool     `json:"shippingDeliveryNotesExists,omitempty"`
	ShippingDeliveryNotesLike   *string   `json:"shippingDeliveryNotesLike,omitempty"`
	ShippingDeliveryNotesNlike  *string   `json:"shippingDeliveryNotesNlike,omitempty"`
	// shippingLine1 (string) search options
	ShippingLine1Eq     *string   `json:"shippingLine1Eq,omitempty"`
	ShippingLine1Ne     *string   `json:"shippingLine1Ne,omitempty"`
	ShippingLine1Gt     *string   `json:"shippingLine1Gt,omitempty"`
	ShippingLine1Gte    *string   `json:"shippingLine1Gte,omitempty"`
	ShippingLine1Lt     *string   `json:"shippingLine1Lt,omitempty"`
	ShippingLine1Lte    *string   `json:"shippingLine1Lte,omitempty"`
	ShippingLine1In     *[]string `json:"shippingLine1In,omitempty"`
	ShippingLine1Nin    *[]string `json:"shippingLine1Nin,omitempty"`
	ShippingLine1Exists *bool     `json:"shippingLine1Exists,omitempty"`
	ShippingLine1Like   *string   `json:"shippingLine1Like,omitempty"`
	ShippingLine1Nlike  *string   `json:"shippingLine1Nlike,omitempty"`
	// shippingLine2 (string) search options
	ShippingLine2Eq     *string   `json:"shippingLine2Eq,omitempty"`
	ShippingLine2Ne     *string   `json:"shippingLine2Ne,omitempty"`
	ShippingLine2Gt     *string   `json:"shippingLine2Gt,omitempty"`
	ShippingLine2Gte    *string   `json:"shippingLine2Gte,omitempty"`
	ShippingLine2Lt     *string   `json:"shippingLine2Lt,omitempty"`
	ShippingLine2Lte    *string   `json:"shippingLine2Lte,omitempty"`
	ShippingLine2In     *[]string `json:"shippingLine2In,omitempty"`
	ShippingLine2Nin    *[]string `json:"shippingLine2Nin,omitempty"`
	ShippingLine2Exists *bool     `json:"shippingLine2Exists,omitempty"`
	ShippingLine2Like   *string   `json:"shippingLine2Like,omitempty"`
	ShippingLine2Nlike  *string   `json:"shippingLine2Nlike,omitempty"`
	// shippingPolicyVersion (int) search options
	ShippingPolicyVersionEq     *int   `json:"shippingPolicyVersionEq,omitempty"`
	ShippingPolicyVersionNe     *int   `json:"shippingPolicyVersionNe,omitempty"`
	ShippingPolicyVersionGt     *int   `json:"shippingPolicyVersionGt,omitempty"`
	ShippingPolicyVersionGte    *int   `json:"shippingPolicyVersionGte,omitempty"`
	ShippingPolicyVersionLt     *int   `json:"shippingPolicyVersionLt,omitempty"`
	ShippingPolicyVersionLte    *int   `json:"shippingPolicyVersionLte,omitempty"`
	ShippingPolicyVersionIn     *[]int `json:"shippingPolicyVersionIn,omitempty"`
	ShippingPolicyVersionNin    *[]int `json:"shippingPolicyVersionNin,omitempty"`
	ShippingPolicyVersionExists *bool  `json:"shippingPolicyVersionExists,omitempty"`
	// shippingPostalCode (string) search options
	ShippingPostalCodeEq     *string   `json:"shippingPostalCodeEq,omitempty"`
	ShippingPostalCodeNe     *string   `json:"shippingPostalCodeNe,omitempty"`
	ShippingPostalCodeGt     *string   `json:"shippingPostalCodeGt,omitempty"`
	ShippingPostalCodeGte    *string   `json:"shippingPostalCodeGte,omitempty"`
	ShippingPostalCodeLt     *string   `json:"shippingPostalCodeLt,omitempty"`
	ShippingPostalCodeLte    *string   `json:"shippingPostalCodeLte,omitempty"`
	ShippingPostalCodeIn     *[]string `json:"shippingPostalCodeIn,omitempty"`
	ShippingPostalCodeNin    *[]string `json:"shippingPostalCodeNin,omitempty"`
	ShippingPostalCodeExists *bool     `json:"shippingPostalCodeExists,omitempty"`
	ShippingPostalCodeLike   *string   `json:"shippingPostalCodeLike,omitempty"`
	ShippingPostalCodeNlike  *string   `json:"shippingPostalCodeNlike,omitempty"`
	// shippingRecipientName (string) search options
	ShippingRecipientNameEq     *string   `json:"shippingRecipientNameEq,omitempty"`
	ShippingRecipientNameNe     *string   `json:"shippingRecipientNameNe,omitempty"`
	ShippingRecipientNameGt     *string   `json:"shippingRecipientNameGt,omitempty"`
	ShippingRecipientNameGte    *string   `json:"shippingRecipientNameGte,omitempty"`
	ShippingRecipientNameLt     *string   `json:"shippingRecipientNameLt,omitempty"`
	ShippingRecipientNameLte    *string   `json:"shippingRecipientNameLte,omitempty"`
	ShippingRecipientNameIn     *[]string `json:"shippingRecipientNameIn,omitempty"`
	ShippingRecipientNameNin    *[]string `json:"shippingRecipientNameNin,omitempty"`
	ShippingRecipientNameExists *bool     `json:"shippingRecipientNameExists,omitempty"`
	ShippingRecipientNameLike   *string   `json:"shippingRecipientNameLike,omitempty"`
	ShippingRecipientNameNlike  *string   `json:"shippingRecipientNameNlike,omitempty"`
	// shippingRegion (string) search options
	ShippingRegionEq     *string   `json:"shippingRegionEq,omitempty"`
	ShippingRegionNe     *string   `json:"shippingRegionNe,omitempty"`
	ShippingRegionGt     *string   `json:"shippingRegionGt,omitempty"`
	ShippingRegionGte    *string   `json:"shippingRegionGte,omitempty"`
	ShippingRegionLt     *string   `json:"shippingRegionLt,omitempty"`
	ShippingRegionLte    *string   `json:"shippingRegionLte,omitempty"`
	ShippingRegionIn     *[]string `json:"shippingRegionIn,omitempty"`
	ShippingRegionNin    *[]string `json:"shippingRegionNin,omitempty"`
	ShippingRegionExists *bool     `json:"shippingRegionExists,omitempty"`
	ShippingRegionLike   *string   `json:"shippingRegionLike,omitempty"`
	ShippingRegionNlike  *string   `json:"shippingRegionNlike,omitempty"`
	// slug (string) search options
	SlugEq     *string   `json:"slugEq,omitempty"`
	SlugNe     *string   `json:"slugNe,omitempty"`
	SlugGt     *string   `json:"slugGt,omitempty"`
	SlugGte    *string   `json:"slugGte,omitempty"`
	SlugLt     *string   `json:"slugLt,omitempty"`
	SlugLte    *string   `json:"slugLte,omitempty"`
	SlugIn     *[]string `json:"slugIn,omitempty"`
	SlugNin    *[]string `json:"slugNin,omitempty"`
	SlugExists *bool     `json:"slugExists,omitempty"`
	SlugLike   *string   `json:"slugLike,omitempty"`
	SlugNlike  *string   `json:"slugNlike,omitempty"`
	// themeColor (string) search options
	ThemeColorEq     *string   `json:"themeColorEq,omitempty"`
	ThemeColorNe     *string   `json:"themeColorNe,omitempty"`
	ThemeColorGt     *string   `json:"themeColorGt,omitempty"`
	ThemeColorGte    *string   `json:"themeColorGte,omitempty"`
	ThemeColorLt     *string   `json:"themeColorLt,omitempty"`
	ThemeColorLte    *string   `json:"themeColorLte,omitempty"`
	ThemeColorIn     *[]string `json:"themeColorIn,omitempty"`
	ThemeColorNin    *[]string `json:"themeColorNin,omitempty"`
	ThemeColorExists *bool     `json:"themeColorExists,omitempty"`
	ThemeColorLike   *string   `json:"themeColorLike,omitempty"`
	ThemeColorNlike  *string   `json:"themeColorNlike,omitempty"`
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
	// welcomeMessage (string) search options
	WelcomeMessageEq     *string   `json:"welcomeMessageEq,omitempty"`
	WelcomeMessageNe     *string   `json:"welcomeMessageNe,omitempty"`
	WelcomeMessageGt     *string   `json:"welcomeMessageGt,omitempty"`
	WelcomeMessageGte    *string   `json:"welcomeMessageGte,omitempty"`
	WelcomeMessageLt     *string   `json:"welcomeMessageLt,omitempty"`
	WelcomeMessageLte    *string   `json:"welcomeMessageLte,omitempty"`
	WelcomeMessageIn     *[]string `json:"welcomeMessageIn,omitempty"`
	WelcomeMessageNin    *[]string `json:"welcomeMessageNin,omitempty"`
	WelcomeMessageExists *bool     `json:"welcomeMessageExists,omitempty"`
	WelcomeMessageLike   *string   `json:"welcomeMessageLike,omitempty"`
	WelcomeMessageNlike  *string   `json:"welcomeMessageNlike,omitempty"`
}

func (o HTTPSelectByIdQuery) ToSelectByIdQuery() (SelectByIdQuery, error) {
	to := SelectByIdQuery{}
	elemid0 := o.Id
	to.Id = elemid0
	return to, nil
}
func (o HTTPSelectBySlugUniqueQuery) ToSelectBySlugUniqueQuery() (SelectBySlugUniqueQuery, error) {
	to := SelectBySlugUniqueQuery{}
	elemslug0 := o.Slug
	to.Slug = elemslug0
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
	if o.AddressAccessModeEq != nil {
		elemaddressAccessModeEq0 := o.AddressAccessModeEq
		to.AddressAccessModeEq = elemaddressAccessModeEq0
	}
	if o.AddressAccessModeNe != nil {
		elemaddressAccessModeNe0 := o.AddressAccessModeNe
		to.AddressAccessModeNe = elemaddressAccessModeNe0
	}
	if o.AddressAccessModeGt != nil {
		elemaddressAccessModeGt0 := o.AddressAccessModeGt
		to.AddressAccessModeGt = elemaddressAccessModeGt0
	}
	if o.AddressAccessModeGte != nil {
		elemaddressAccessModeGte0 := o.AddressAccessModeGte
		to.AddressAccessModeGte = elemaddressAccessModeGte0
	}
	if o.AddressAccessModeLt != nil {
		elemaddressAccessModeLt0 := o.AddressAccessModeLt
		to.AddressAccessModeLt = elemaddressAccessModeLt0
	}
	if o.AddressAccessModeLte != nil {
		elemaddressAccessModeLte0 := o.AddressAccessModeLte
		to.AddressAccessModeLte = elemaddressAccessModeLte0
	}
	if o.AddressAccessModeIn != nil {
		elemaddressAccessModeIn0 := make([]enum_address_access_mode.Value, 0)
		for _, oaddressAccessModeIn0 := range *o.AddressAccessModeIn {
			elemaddressAccessModeIn1 := oaddressAccessModeIn0
			elemaddressAccessModeIn0 = append(elemaddressAccessModeIn0, elemaddressAccessModeIn1)
		}
		to.AddressAccessModeIn = &elemaddressAccessModeIn0
	}
	if o.AddressAccessModeNin != nil {
		elemaddressAccessModeNin0 := make([]enum_address_access_mode.Value, 0)
		for _, oaddressAccessModeNin0 := range *o.AddressAccessModeNin {
			elemaddressAccessModeNin1 := oaddressAccessModeNin0
			elemaddressAccessModeNin0 = append(elemaddressAccessModeNin0, elemaddressAccessModeNin1)
		}
		to.AddressAccessModeNin = &elemaddressAccessModeNin0
	}
	if o.AddressAccessModeExists != nil {
		elemaddressAccessModeExists0 := o.AddressAccessModeExists
		to.AddressAccessModeExists = elemaddressAccessModeExists0
	}
	if o.CoverImageUrlEq != nil {
		elemcoverImageUrlEq0 := o.CoverImageUrlEq
		to.CoverImageUrlEq = elemcoverImageUrlEq0
	}
	if o.CoverImageUrlNe != nil {
		elemcoverImageUrlNe0 := o.CoverImageUrlNe
		to.CoverImageUrlNe = elemcoverImageUrlNe0
	}
	if o.CoverImageUrlGt != nil {
		elemcoverImageUrlGt0 := o.CoverImageUrlGt
		to.CoverImageUrlGt = elemcoverImageUrlGt0
	}
	if o.CoverImageUrlGte != nil {
		elemcoverImageUrlGte0 := o.CoverImageUrlGte
		to.CoverImageUrlGte = elemcoverImageUrlGte0
	}
	if o.CoverImageUrlLt != nil {
		elemcoverImageUrlLt0 := o.CoverImageUrlLt
		to.CoverImageUrlLt = elemcoverImageUrlLt0
	}
	if o.CoverImageUrlLte != nil {
		elemcoverImageUrlLte0 := o.CoverImageUrlLte
		to.CoverImageUrlLte = elemcoverImageUrlLte0
	}
	if o.CoverImageUrlIn != nil {
		elemcoverImageUrlIn0 := make([]string, 0)
		for _, ocoverImageUrlIn0 := range *o.CoverImageUrlIn {
			elemcoverImageUrlIn1 := ocoverImageUrlIn0
			elemcoverImageUrlIn0 = append(elemcoverImageUrlIn0, elemcoverImageUrlIn1)
		}
		to.CoverImageUrlIn = &elemcoverImageUrlIn0
	}
	if o.CoverImageUrlNin != nil {
		elemcoverImageUrlNin0 := make([]string, 0)
		for _, ocoverImageUrlNin0 := range *o.CoverImageUrlNin {
			elemcoverImageUrlNin1 := ocoverImageUrlNin0
			elemcoverImageUrlNin0 = append(elemcoverImageUrlNin0, elemcoverImageUrlNin1)
		}
		to.CoverImageUrlNin = &elemcoverImageUrlNin0
	}
	if o.CoverImageUrlExists != nil {
		elemcoverImageUrlExists0 := o.CoverImageUrlExists
		to.CoverImageUrlExists = elemcoverImageUrlExists0
	}
	if o.CoverImageUrlLike != nil {
		elemcoverImageUrlLike0 := o.CoverImageUrlLike
		to.CoverImageUrlLike = elemcoverImageUrlLike0
	}
	if o.CoverImageUrlNlike != nil {
		elemcoverImageUrlNlike0 := o.CoverImageUrlNlike
		to.CoverImageUrlNlike = elemcoverImageUrlNlike0
	}
	if o.Created != nil {
		elemcreated0, err := o.Created.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.Created = &elemcreated0
	}
	if o.DueDateEq != nil {
		elemdueDateEq0 := o.DueDateEq
		to.DueDateEq = elemdueDateEq0
	}
	if o.DueDateNe != nil {
		elemdueDateNe0 := o.DueDateNe
		to.DueDateNe = elemdueDateNe0
	}
	if o.DueDateGt != nil {
		elemdueDateGt0 := o.DueDateGt
		to.DueDateGt = elemdueDateGt0
	}
	if o.DueDateGte != nil {
		elemdueDateGte0 := o.DueDateGte
		to.DueDateGte = elemdueDateGte0
	}
	if o.DueDateLt != nil {
		elemdueDateLt0 := o.DueDateLt
		to.DueDateLt = elemdueDateLt0
	}
	if o.DueDateLte != nil {
		elemdueDateLte0 := o.DueDateLte
		to.DueDateLte = elemdueDateLte0
	}
	if o.DueDateIn != nil {
		elemdueDateIn0 := make([]time.Time, 0)
		for _, odueDateIn0 := range *o.DueDateIn {
			elemdueDateIn1 := odueDateIn0
			elemdueDateIn0 = append(elemdueDateIn0, elemdueDateIn1)
		}
		to.DueDateIn = &elemdueDateIn0
	}
	if o.DueDateNin != nil {
		elemdueDateNin0 := make([]time.Time, 0)
		for _, odueDateNin0 := range *o.DueDateNin {
			elemdueDateNin1 := odueDateNin0
			elemdueDateNin0 = append(elemdueDateNin0, elemdueDateNin1)
		}
		to.DueDateNin = &elemdueDateNin0
	}
	if o.DueDateExists != nil {
		elemdueDateExists0 := o.DueDateExists
		to.DueDateExists = elemdueDateExists0
	}
	if o.IsPublicEq != nil {
		elemisPublicEq0 := o.IsPublicEq
		to.IsPublicEq = elemisPublicEq0
	}
	if o.IsPublicNe != nil {
		elemisPublicNe0 := o.IsPublicNe
		to.IsPublicNe = elemisPublicNe0
	}
	if o.IsPublicGt != nil {
		elemisPublicGt0 := o.IsPublicGt
		to.IsPublicGt = elemisPublicGt0
	}
	if o.IsPublicGte != nil {
		elemisPublicGte0 := o.IsPublicGte
		to.IsPublicGte = elemisPublicGte0
	}
	if o.IsPublicLt != nil {
		elemisPublicLt0 := o.IsPublicLt
		to.IsPublicLt = elemisPublicLt0
	}
	if o.IsPublicLte != nil {
		elemisPublicLte0 := o.IsPublicLte
		to.IsPublicLte = elemisPublicLte0
	}
	if o.IsPublicIn != nil {
		elemisPublicIn0 := make([]bool, 0)
		for _, oisPublicIn0 := range *o.IsPublicIn {
			elemisPublicIn1 := oisPublicIn0
			elemisPublicIn0 = append(elemisPublicIn0, elemisPublicIn1)
		}
		to.IsPublicIn = &elemisPublicIn0
	}
	if o.IsPublicNin != nil {
		elemisPublicNin0 := make([]bool, 0)
		for _, oisPublicNin0 := range *o.IsPublicNin {
			elemisPublicNin1 := oisPublicNin0
			elemisPublicNin0 = append(elemisPublicNin0, elemisPublicNin1)
		}
		to.IsPublicNin = &elemisPublicNin0
	}
	if o.IsPublicExists != nil {
		elemisPublicExists0 := o.IsPublicExists
		to.IsPublicExists = elemisPublicExists0
	}
	if o.OwnerIdEq != nil {
		elemownerIdEq0 := o.OwnerIdEq
		to.OwnerIdEq = elemownerIdEq0
	}
	if o.OwnerIdIn != nil {
		elemownerIdIn0 := make([]string, 0)
		for _, oownerIdIn0 := range *o.OwnerIdIn {
			elemownerIdIn1 := oownerIdIn0
			elemownerIdIn0 = append(elemownerIdIn0, elemownerIdIn1)
		}
		to.OwnerIdIn = &elemownerIdIn0
	}
	if o.OwnerIdNin != nil {
		elemownerIdNin0 := make([]string, 0)
		for _, oownerIdNin0 := range *o.OwnerIdNin {
			elemownerIdNin1 := oownerIdNin0
			elemownerIdNin0 = append(elemownerIdNin0, elemownerIdNin1)
		}
		to.OwnerIdNin = &elemownerIdNin0
	}
	if o.OwnerIdExists != nil {
		elemownerIdExists0 := o.OwnerIdExists
		to.OwnerIdExists = elemownerIdExists0
	}
	if o.ParentNamesEq != nil {
		elemparentNamesEq0 := o.ParentNamesEq
		to.ParentNamesEq = elemparentNamesEq0
	}
	if o.ParentNamesNe != nil {
		elemparentNamesNe0 := o.ParentNamesNe
		to.ParentNamesNe = elemparentNamesNe0
	}
	if o.ParentNamesGt != nil {
		elemparentNamesGt0 := o.ParentNamesGt
		to.ParentNamesGt = elemparentNamesGt0
	}
	if o.ParentNamesGte != nil {
		elemparentNamesGte0 := o.ParentNamesGte
		to.ParentNamesGte = elemparentNamesGte0
	}
	if o.ParentNamesLt != nil {
		elemparentNamesLt0 := o.ParentNamesLt
		to.ParentNamesLt = elemparentNamesLt0
	}
	if o.ParentNamesLte != nil {
		elemparentNamesLte0 := o.ParentNamesLte
		to.ParentNamesLte = elemparentNamesLte0
	}
	if o.ParentNamesIn != nil {
		elemparentNamesIn0 := make([]string, 0)
		for _, oparentNamesIn0 := range *o.ParentNamesIn {
			elemparentNamesIn1 := oparentNamesIn0
			elemparentNamesIn0 = append(elemparentNamesIn0, elemparentNamesIn1)
		}
		to.ParentNamesIn = &elemparentNamesIn0
	}
	if o.ParentNamesNin != nil {
		elemparentNamesNin0 := make([]string, 0)
		for _, oparentNamesNin0 := range *o.ParentNamesNin {
			elemparentNamesNin1 := oparentNamesNin0
			elemparentNamesNin0 = append(elemparentNamesNin0, elemparentNamesNin1)
		}
		to.ParentNamesNin = &elemparentNamesNin0
	}
	if o.ParentNamesExists != nil {
		elemparentNamesExists0 := o.ParentNamesExists
		to.ParentNamesExists = elemparentNamesExists0
	}
	if o.ParentNamesLike != nil {
		elemparentNamesLike0 := o.ParentNamesLike
		to.ParentNamesLike = elemparentNamesLike0
	}
	if o.ParentNamesNlike != nil {
		elemparentNamesNlike0 := o.ParentNamesNlike
		to.ParentNamesNlike = elemparentNamesNlike0
	}
	if o.ShippingCityEq != nil {
		elemshippingCityEq0 := o.ShippingCityEq
		to.ShippingCityEq = elemshippingCityEq0
	}
	if o.ShippingCityNe != nil {
		elemshippingCityNe0 := o.ShippingCityNe
		to.ShippingCityNe = elemshippingCityNe0
	}
	if o.ShippingCityGt != nil {
		elemshippingCityGt0 := o.ShippingCityGt
		to.ShippingCityGt = elemshippingCityGt0
	}
	if o.ShippingCityGte != nil {
		elemshippingCityGte0 := o.ShippingCityGte
		to.ShippingCityGte = elemshippingCityGte0
	}
	if o.ShippingCityLt != nil {
		elemshippingCityLt0 := o.ShippingCityLt
		to.ShippingCityLt = elemshippingCityLt0
	}
	if o.ShippingCityLte != nil {
		elemshippingCityLte0 := o.ShippingCityLte
		to.ShippingCityLte = elemshippingCityLte0
	}
	if o.ShippingCityIn != nil {
		elemshippingCityIn0 := make([]string, 0)
		for _, oshippingCityIn0 := range *o.ShippingCityIn {
			elemshippingCityIn1 := oshippingCityIn0
			elemshippingCityIn0 = append(elemshippingCityIn0, elemshippingCityIn1)
		}
		to.ShippingCityIn = &elemshippingCityIn0
	}
	if o.ShippingCityNin != nil {
		elemshippingCityNin0 := make([]string, 0)
		for _, oshippingCityNin0 := range *o.ShippingCityNin {
			elemshippingCityNin1 := oshippingCityNin0
			elemshippingCityNin0 = append(elemshippingCityNin0, elemshippingCityNin1)
		}
		to.ShippingCityNin = &elemshippingCityNin0
	}
	if o.ShippingCityExists != nil {
		elemshippingCityExists0 := o.ShippingCityExists
		to.ShippingCityExists = elemshippingCityExists0
	}
	if o.ShippingCityLike != nil {
		elemshippingCityLike0 := o.ShippingCityLike
		to.ShippingCityLike = elemshippingCityLike0
	}
	if o.ShippingCityNlike != nil {
		elemshippingCityNlike0 := o.ShippingCityNlike
		to.ShippingCityNlike = elemshippingCityNlike0
	}
	if o.ShippingCountryEq != nil {
		elemshippingCountryEq0 := o.ShippingCountryEq
		to.ShippingCountryEq = elemshippingCountryEq0
	}
	if o.ShippingCountryNe != nil {
		elemshippingCountryNe0 := o.ShippingCountryNe
		to.ShippingCountryNe = elemshippingCountryNe0
	}
	if o.ShippingCountryGt != nil {
		elemshippingCountryGt0 := o.ShippingCountryGt
		to.ShippingCountryGt = elemshippingCountryGt0
	}
	if o.ShippingCountryGte != nil {
		elemshippingCountryGte0 := o.ShippingCountryGte
		to.ShippingCountryGte = elemshippingCountryGte0
	}
	if o.ShippingCountryLt != nil {
		elemshippingCountryLt0 := o.ShippingCountryLt
		to.ShippingCountryLt = elemshippingCountryLt0
	}
	if o.ShippingCountryLte != nil {
		elemshippingCountryLte0 := o.ShippingCountryLte
		to.ShippingCountryLte = elemshippingCountryLte0
	}
	if o.ShippingCountryIn != nil {
		elemshippingCountryIn0 := make([]string, 0)
		for _, oshippingCountryIn0 := range *o.ShippingCountryIn {
			elemshippingCountryIn1 := oshippingCountryIn0
			elemshippingCountryIn0 = append(elemshippingCountryIn0, elemshippingCountryIn1)
		}
		to.ShippingCountryIn = &elemshippingCountryIn0
	}
	if o.ShippingCountryNin != nil {
		elemshippingCountryNin0 := make([]string, 0)
		for _, oshippingCountryNin0 := range *o.ShippingCountryNin {
			elemshippingCountryNin1 := oshippingCountryNin0
			elemshippingCountryNin0 = append(elemshippingCountryNin0, elemshippingCountryNin1)
		}
		to.ShippingCountryNin = &elemshippingCountryNin0
	}
	if o.ShippingCountryExists != nil {
		elemshippingCountryExists0 := o.ShippingCountryExists
		to.ShippingCountryExists = elemshippingCountryExists0
	}
	if o.ShippingCountryLike != nil {
		elemshippingCountryLike0 := o.ShippingCountryLike
		to.ShippingCountryLike = elemshippingCountryLike0
	}
	if o.ShippingCountryNlike != nil {
		elemshippingCountryNlike0 := o.ShippingCountryNlike
		to.ShippingCountryNlike = elemshippingCountryNlike0
	}
	if o.ShippingDeliveryNotesEq != nil {
		elemshippingDeliveryNotesEq0 := o.ShippingDeliveryNotesEq
		to.ShippingDeliveryNotesEq = elemshippingDeliveryNotesEq0
	}
	if o.ShippingDeliveryNotesNe != nil {
		elemshippingDeliveryNotesNe0 := o.ShippingDeliveryNotesNe
		to.ShippingDeliveryNotesNe = elemshippingDeliveryNotesNe0
	}
	if o.ShippingDeliveryNotesGt != nil {
		elemshippingDeliveryNotesGt0 := o.ShippingDeliveryNotesGt
		to.ShippingDeliveryNotesGt = elemshippingDeliveryNotesGt0
	}
	if o.ShippingDeliveryNotesGte != nil {
		elemshippingDeliveryNotesGte0 := o.ShippingDeliveryNotesGte
		to.ShippingDeliveryNotesGte = elemshippingDeliveryNotesGte0
	}
	if o.ShippingDeliveryNotesLt != nil {
		elemshippingDeliveryNotesLt0 := o.ShippingDeliveryNotesLt
		to.ShippingDeliveryNotesLt = elemshippingDeliveryNotesLt0
	}
	if o.ShippingDeliveryNotesLte != nil {
		elemshippingDeliveryNotesLte0 := o.ShippingDeliveryNotesLte
		to.ShippingDeliveryNotesLte = elemshippingDeliveryNotesLte0
	}
	if o.ShippingDeliveryNotesIn != nil {
		elemshippingDeliveryNotesIn0 := make([]string, 0)
		for _, oshippingDeliveryNotesIn0 := range *o.ShippingDeliveryNotesIn {
			elemshippingDeliveryNotesIn1 := oshippingDeliveryNotesIn0
			elemshippingDeliveryNotesIn0 = append(elemshippingDeliveryNotesIn0, elemshippingDeliveryNotesIn1)
		}
		to.ShippingDeliveryNotesIn = &elemshippingDeliveryNotesIn0
	}
	if o.ShippingDeliveryNotesNin != nil {
		elemshippingDeliveryNotesNin0 := make([]string, 0)
		for _, oshippingDeliveryNotesNin0 := range *o.ShippingDeliveryNotesNin {
			elemshippingDeliveryNotesNin1 := oshippingDeliveryNotesNin0
			elemshippingDeliveryNotesNin0 = append(elemshippingDeliveryNotesNin0, elemshippingDeliveryNotesNin1)
		}
		to.ShippingDeliveryNotesNin = &elemshippingDeliveryNotesNin0
	}
	if o.ShippingDeliveryNotesExists != nil {
		elemshippingDeliveryNotesExists0 := o.ShippingDeliveryNotesExists
		to.ShippingDeliveryNotesExists = elemshippingDeliveryNotesExists0
	}
	if o.ShippingDeliveryNotesLike != nil {
		elemshippingDeliveryNotesLike0 := o.ShippingDeliveryNotesLike
		to.ShippingDeliveryNotesLike = elemshippingDeliveryNotesLike0
	}
	if o.ShippingDeliveryNotesNlike != nil {
		elemshippingDeliveryNotesNlike0 := o.ShippingDeliveryNotesNlike
		to.ShippingDeliveryNotesNlike = elemshippingDeliveryNotesNlike0
	}
	if o.ShippingLine1Eq != nil {
		elemshippingLine1Eq0 := o.ShippingLine1Eq
		to.ShippingLine1Eq = elemshippingLine1Eq0
	}
	if o.ShippingLine1Ne != nil {
		elemshippingLine1Ne0 := o.ShippingLine1Ne
		to.ShippingLine1Ne = elemshippingLine1Ne0
	}
	if o.ShippingLine1Gt != nil {
		elemshippingLine1Gt0 := o.ShippingLine1Gt
		to.ShippingLine1Gt = elemshippingLine1Gt0
	}
	if o.ShippingLine1Gte != nil {
		elemshippingLine1Gte0 := o.ShippingLine1Gte
		to.ShippingLine1Gte = elemshippingLine1Gte0
	}
	if o.ShippingLine1Lt != nil {
		elemshippingLine1Lt0 := o.ShippingLine1Lt
		to.ShippingLine1Lt = elemshippingLine1Lt0
	}
	if o.ShippingLine1Lte != nil {
		elemshippingLine1Lte0 := o.ShippingLine1Lte
		to.ShippingLine1Lte = elemshippingLine1Lte0
	}
	if o.ShippingLine1In != nil {
		elemshippingLine1In0 := make([]string, 0)
		for _, oshippingLine1In0 := range *o.ShippingLine1In {
			elemshippingLine1In1 := oshippingLine1In0
			elemshippingLine1In0 = append(elemshippingLine1In0, elemshippingLine1In1)
		}
		to.ShippingLine1In = &elemshippingLine1In0
	}
	if o.ShippingLine1Nin != nil {
		elemshippingLine1Nin0 := make([]string, 0)
		for _, oshippingLine1Nin0 := range *o.ShippingLine1Nin {
			elemshippingLine1Nin1 := oshippingLine1Nin0
			elemshippingLine1Nin0 = append(elemshippingLine1Nin0, elemshippingLine1Nin1)
		}
		to.ShippingLine1Nin = &elemshippingLine1Nin0
	}
	if o.ShippingLine1Exists != nil {
		elemshippingLine1Exists0 := o.ShippingLine1Exists
		to.ShippingLine1Exists = elemshippingLine1Exists0
	}
	if o.ShippingLine1Like != nil {
		elemshippingLine1Like0 := o.ShippingLine1Like
		to.ShippingLine1Like = elemshippingLine1Like0
	}
	if o.ShippingLine1Nlike != nil {
		elemshippingLine1Nlike0 := o.ShippingLine1Nlike
		to.ShippingLine1Nlike = elemshippingLine1Nlike0
	}
	if o.ShippingLine2Eq != nil {
		elemshippingLine2Eq0 := o.ShippingLine2Eq
		to.ShippingLine2Eq = elemshippingLine2Eq0
	}
	if o.ShippingLine2Ne != nil {
		elemshippingLine2Ne0 := o.ShippingLine2Ne
		to.ShippingLine2Ne = elemshippingLine2Ne0
	}
	if o.ShippingLine2Gt != nil {
		elemshippingLine2Gt0 := o.ShippingLine2Gt
		to.ShippingLine2Gt = elemshippingLine2Gt0
	}
	if o.ShippingLine2Gte != nil {
		elemshippingLine2Gte0 := o.ShippingLine2Gte
		to.ShippingLine2Gte = elemshippingLine2Gte0
	}
	if o.ShippingLine2Lt != nil {
		elemshippingLine2Lt0 := o.ShippingLine2Lt
		to.ShippingLine2Lt = elemshippingLine2Lt0
	}
	if o.ShippingLine2Lte != nil {
		elemshippingLine2Lte0 := o.ShippingLine2Lte
		to.ShippingLine2Lte = elemshippingLine2Lte0
	}
	if o.ShippingLine2In != nil {
		elemshippingLine2In0 := make([]string, 0)
		for _, oshippingLine2In0 := range *o.ShippingLine2In {
			elemshippingLine2In1 := oshippingLine2In0
			elemshippingLine2In0 = append(elemshippingLine2In0, elemshippingLine2In1)
		}
		to.ShippingLine2In = &elemshippingLine2In0
	}
	if o.ShippingLine2Nin != nil {
		elemshippingLine2Nin0 := make([]string, 0)
		for _, oshippingLine2Nin0 := range *o.ShippingLine2Nin {
			elemshippingLine2Nin1 := oshippingLine2Nin0
			elemshippingLine2Nin0 = append(elemshippingLine2Nin0, elemshippingLine2Nin1)
		}
		to.ShippingLine2Nin = &elemshippingLine2Nin0
	}
	if o.ShippingLine2Exists != nil {
		elemshippingLine2Exists0 := o.ShippingLine2Exists
		to.ShippingLine2Exists = elemshippingLine2Exists0
	}
	if o.ShippingLine2Like != nil {
		elemshippingLine2Like0 := o.ShippingLine2Like
		to.ShippingLine2Like = elemshippingLine2Like0
	}
	if o.ShippingLine2Nlike != nil {
		elemshippingLine2Nlike0 := o.ShippingLine2Nlike
		to.ShippingLine2Nlike = elemshippingLine2Nlike0
	}
	if o.ShippingPolicyVersionEq != nil {
		elemshippingPolicyVersionEq0 := o.ShippingPolicyVersionEq
		to.ShippingPolicyVersionEq = elemshippingPolicyVersionEq0
	}
	if o.ShippingPolicyVersionNe != nil {
		elemshippingPolicyVersionNe0 := o.ShippingPolicyVersionNe
		to.ShippingPolicyVersionNe = elemshippingPolicyVersionNe0
	}
	if o.ShippingPolicyVersionGt != nil {
		elemshippingPolicyVersionGt0 := o.ShippingPolicyVersionGt
		to.ShippingPolicyVersionGt = elemshippingPolicyVersionGt0
	}
	if o.ShippingPolicyVersionGte != nil {
		elemshippingPolicyVersionGte0 := o.ShippingPolicyVersionGte
		to.ShippingPolicyVersionGte = elemshippingPolicyVersionGte0
	}
	if o.ShippingPolicyVersionLt != nil {
		elemshippingPolicyVersionLt0 := o.ShippingPolicyVersionLt
		to.ShippingPolicyVersionLt = elemshippingPolicyVersionLt0
	}
	if o.ShippingPolicyVersionLte != nil {
		elemshippingPolicyVersionLte0 := o.ShippingPolicyVersionLte
		to.ShippingPolicyVersionLte = elemshippingPolicyVersionLte0
	}
	if o.ShippingPolicyVersionIn != nil {
		elemshippingPolicyVersionIn0 := make([]int, 0)
		for _, oshippingPolicyVersionIn0 := range *o.ShippingPolicyVersionIn {
			elemshippingPolicyVersionIn1 := oshippingPolicyVersionIn0
			elemshippingPolicyVersionIn0 = append(elemshippingPolicyVersionIn0, elemshippingPolicyVersionIn1)
		}
		to.ShippingPolicyVersionIn = &elemshippingPolicyVersionIn0
	}
	if o.ShippingPolicyVersionNin != nil {
		elemshippingPolicyVersionNin0 := make([]int, 0)
		for _, oshippingPolicyVersionNin0 := range *o.ShippingPolicyVersionNin {
			elemshippingPolicyVersionNin1 := oshippingPolicyVersionNin0
			elemshippingPolicyVersionNin0 = append(elemshippingPolicyVersionNin0, elemshippingPolicyVersionNin1)
		}
		to.ShippingPolicyVersionNin = &elemshippingPolicyVersionNin0
	}
	if o.ShippingPolicyVersionExists != nil {
		elemshippingPolicyVersionExists0 := o.ShippingPolicyVersionExists
		to.ShippingPolicyVersionExists = elemshippingPolicyVersionExists0
	}
	if o.ShippingPostalCodeEq != nil {
		elemshippingPostalCodeEq0 := o.ShippingPostalCodeEq
		to.ShippingPostalCodeEq = elemshippingPostalCodeEq0
	}
	if o.ShippingPostalCodeNe != nil {
		elemshippingPostalCodeNe0 := o.ShippingPostalCodeNe
		to.ShippingPostalCodeNe = elemshippingPostalCodeNe0
	}
	if o.ShippingPostalCodeGt != nil {
		elemshippingPostalCodeGt0 := o.ShippingPostalCodeGt
		to.ShippingPostalCodeGt = elemshippingPostalCodeGt0
	}
	if o.ShippingPostalCodeGte != nil {
		elemshippingPostalCodeGte0 := o.ShippingPostalCodeGte
		to.ShippingPostalCodeGte = elemshippingPostalCodeGte0
	}
	if o.ShippingPostalCodeLt != nil {
		elemshippingPostalCodeLt0 := o.ShippingPostalCodeLt
		to.ShippingPostalCodeLt = elemshippingPostalCodeLt0
	}
	if o.ShippingPostalCodeLte != nil {
		elemshippingPostalCodeLte0 := o.ShippingPostalCodeLte
		to.ShippingPostalCodeLte = elemshippingPostalCodeLte0
	}
	if o.ShippingPostalCodeIn != nil {
		elemshippingPostalCodeIn0 := make([]string, 0)
		for _, oshippingPostalCodeIn0 := range *o.ShippingPostalCodeIn {
			elemshippingPostalCodeIn1 := oshippingPostalCodeIn0
			elemshippingPostalCodeIn0 = append(elemshippingPostalCodeIn0, elemshippingPostalCodeIn1)
		}
		to.ShippingPostalCodeIn = &elemshippingPostalCodeIn0
	}
	if o.ShippingPostalCodeNin != nil {
		elemshippingPostalCodeNin0 := make([]string, 0)
		for _, oshippingPostalCodeNin0 := range *o.ShippingPostalCodeNin {
			elemshippingPostalCodeNin1 := oshippingPostalCodeNin0
			elemshippingPostalCodeNin0 = append(elemshippingPostalCodeNin0, elemshippingPostalCodeNin1)
		}
		to.ShippingPostalCodeNin = &elemshippingPostalCodeNin0
	}
	if o.ShippingPostalCodeExists != nil {
		elemshippingPostalCodeExists0 := o.ShippingPostalCodeExists
		to.ShippingPostalCodeExists = elemshippingPostalCodeExists0
	}
	if o.ShippingPostalCodeLike != nil {
		elemshippingPostalCodeLike0 := o.ShippingPostalCodeLike
		to.ShippingPostalCodeLike = elemshippingPostalCodeLike0
	}
	if o.ShippingPostalCodeNlike != nil {
		elemshippingPostalCodeNlike0 := o.ShippingPostalCodeNlike
		to.ShippingPostalCodeNlike = elemshippingPostalCodeNlike0
	}
	if o.ShippingRecipientNameEq != nil {
		elemshippingRecipientNameEq0 := o.ShippingRecipientNameEq
		to.ShippingRecipientNameEq = elemshippingRecipientNameEq0
	}
	if o.ShippingRecipientNameNe != nil {
		elemshippingRecipientNameNe0 := o.ShippingRecipientNameNe
		to.ShippingRecipientNameNe = elemshippingRecipientNameNe0
	}
	if o.ShippingRecipientNameGt != nil {
		elemshippingRecipientNameGt0 := o.ShippingRecipientNameGt
		to.ShippingRecipientNameGt = elemshippingRecipientNameGt0
	}
	if o.ShippingRecipientNameGte != nil {
		elemshippingRecipientNameGte0 := o.ShippingRecipientNameGte
		to.ShippingRecipientNameGte = elemshippingRecipientNameGte0
	}
	if o.ShippingRecipientNameLt != nil {
		elemshippingRecipientNameLt0 := o.ShippingRecipientNameLt
		to.ShippingRecipientNameLt = elemshippingRecipientNameLt0
	}
	if o.ShippingRecipientNameLte != nil {
		elemshippingRecipientNameLte0 := o.ShippingRecipientNameLte
		to.ShippingRecipientNameLte = elemshippingRecipientNameLte0
	}
	if o.ShippingRecipientNameIn != nil {
		elemshippingRecipientNameIn0 := make([]string, 0)
		for _, oshippingRecipientNameIn0 := range *o.ShippingRecipientNameIn {
			elemshippingRecipientNameIn1 := oshippingRecipientNameIn0
			elemshippingRecipientNameIn0 = append(elemshippingRecipientNameIn0, elemshippingRecipientNameIn1)
		}
		to.ShippingRecipientNameIn = &elemshippingRecipientNameIn0
	}
	if o.ShippingRecipientNameNin != nil {
		elemshippingRecipientNameNin0 := make([]string, 0)
		for _, oshippingRecipientNameNin0 := range *o.ShippingRecipientNameNin {
			elemshippingRecipientNameNin1 := oshippingRecipientNameNin0
			elemshippingRecipientNameNin0 = append(elemshippingRecipientNameNin0, elemshippingRecipientNameNin1)
		}
		to.ShippingRecipientNameNin = &elemshippingRecipientNameNin0
	}
	if o.ShippingRecipientNameExists != nil {
		elemshippingRecipientNameExists0 := o.ShippingRecipientNameExists
		to.ShippingRecipientNameExists = elemshippingRecipientNameExists0
	}
	if o.ShippingRecipientNameLike != nil {
		elemshippingRecipientNameLike0 := o.ShippingRecipientNameLike
		to.ShippingRecipientNameLike = elemshippingRecipientNameLike0
	}
	if o.ShippingRecipientNameNlike != nil {
		elemshippingRecipientNameNlike0 := o.ShippingRecipientNameNlike
		to.ShippingRecipientNameNlike = elemshippingRecipientNameNlike0
	}
	if o.ShippingRegionEq != nil {
		elemshippingRegionEq0 := o.ShippingRegionEq
		to.ShippingRegionEq = elemshippingRegionEq0
	}
	if o.ShippingRegionNe != nil {
		elemshippingRegionNe0 := o.ShippingRegionNe
		to.ShippingRegionNe = elemshippingRegionNe0
	}
	if o.ShippingRegionGt != nil {
		elemshippingRegionGt0 := o.ShippingRegionGt
		to.ShippingRegionGt = elemshippingRegionGt0
	}
	if o.ShippingRegionGte != nil {
		elemshippingRegionGte0 := o.ShippingRegionGte
		to.ShippingRegionGte = elemshippingRegionGte0
	}
	if o.ShippingRegionLt != nil {
		elemshippingRegionLt0 := o.ShippingRegionLt
		to.ShippingRegionLt = elemshippingRegionLt0
	}
	if o.ShippingRegionLte != nil {
		elemshippingRegionLte0 := o.ShippingRegionLte
		to.ShippingRegionLte = elemshippingRegionLte0
	}
	if o.ShippingRegionIn != nil {
		elemshippingRegionIn0 := make([]string, 0)
		for _, oshippingRegionIn0 := range *o.ShippingRegionIn {
			elemshippingRegionIn1 := oshippingRegionIn0
			elemshippingRegionIn0 = append(elemshippingRegionIn0, elemshippingRegionIn1)
		}
		to.ShippingRegionIn = &elemshippingRegionIn0
	}
	if o.ShippingRegionNin != nil {
		elemshippingRegionNin0 := make([]string, 0)
		for _, oshippingRegionNin0 := range *o.ShippingRegionNin {
			elemshippingRegionNin1 := oshippingRegionNin0
			elemshippingRegionNin0 = append(elemshippingRegionNin0, elemshippingRegionNin1)
		}
		to.ShippingRegionNin = &elemshippingRegionNin0
	}
	if o.ShippingRegionExists != nil {
		elemshippingRegionExists0 := o.ShippingRegionExists
		to.ShippingRegionExists = elemshippingRegionExists0
	}
	if o.ShippingRegionLike != nil {
		elemshippingRegionLike0 := o.ShippingRegionLike
		to.ShippingRegionLike = elemshippingRegionLike0
	}
	if o.ShippingRegionNlike != nil {
		elemshippingRegionNlike0 := o.ShippingRegionNlike
		to.ShippingRegionNlike = elemshippingRegionNlike0
	}
	if o.SlugEq != nil {
		elemslugEq0 := o.SlugEq
		to.SlugEq = elemslugEq0
	}
	if o.SlugNe != nil {
		elemslugNe0 := o.SlugNe
		to.SlugNe = elemslugNe0
	}
	if o.SlugGt != nil {
		elemslugGt0 := o.SlugGt
		to.SlugGt = elemslugGt0
	}
	if o.SlugGte != nil {
		elemslugGte0 := o.SlugGte
		to.SlugGte = elemslugGte0
	}
	if o.SlugLt != nil {
		elemslugLt0 := o.SlugLt
		to.SlugLt = elemslugLt0
	}
	if o.SlugLte != nil {
		elemslugLte0 := o.SlugLte
		to.SlugLte = elemslugLte0
	}
	if o.SlugIn != nil {
		elemslugIn0 := make([]string, 0)
		for _, oslugIn0 := range *o.SlugIn {
			elemslugIn1 := oslugIn0
			elemslugIn0 = append(elemslugIn0, elemslugIn1)
		}
		to.SlugIn = &elemslugIn0
	}
	if o.SlugNin != nil {
		elemslugNin0 := make([]string, 0)
		for _, oslugNin0 := range *o.SlugNin {
			elemslugNin1 := oslugNin0
			elemslugNin0 = append(elemslugNin0, elemslugNin1)
		}
		to.SlugNin = &elemslugNin0
	}
	if o.SlugExists != nil {
		elemslugExists0 := o.SlugExists
		to.SlugExists = elemslugExists0
	}
	if o.SlugLike != nil {
		elemslugLike0 := o.SlugLike
		to.SlugLike = elemslugLike0
	}
	if o.SlugNlike != nil {
		elemslugNlike0 := o.SlugNlike
		to.SlugNlike = elemslugNlike0
	}
	if o.ThemeColorEq != nil {
		elemthemeColorEq0 := o.ThemeColorEq
		to.ThemeColorEq = elemthemeColorEq0
	}
	if o.ThemeColorNe != nil {
		elemthemeColorNe0 := o.ThemeColorNe
		to.ThemeColorNe = elemthemeColorNe0
	}
	if o.ThemeColorGt != nil {
		elemthemeColorGt0 := o.ThemeColorGt
		to.ThemeColorGt = elemthemeColorGt0
	}
	if o.ThemeColorGte != nil {
		elemthemeColorGte0 := o.ThemeColorGte
		to.ThemeColorGte = elemthemeColorGte0
	}
	if o.ThemeColorLt != nil {
		elemthemeColorLt0 := o.ThemeColorLt
		to.ThemeColorLt = elemthemeColorLt0
	}
	if o.ThemeColorLte != nil {
		elemthemeColorLte0 := o.ThemeColorLte
		to.ThemeColorLte = elemthemeColorLte0
	}
	if o.ThemeColorIn != nil {
		elemthemeColorIn0 := make([]string, 0)
		for _, othemeColorIn0 := range *o.ThemeColorIn {
			elemthemeColorIn1 := othemeColorIn0
			elemthemeColorIn0 = append(elemthemeColorIn0, elemthemeColorIn1)
		}
		to.ThemeColorIn = &elemthemeColorIn0
	}
	if o.ThemeColorNin != nil {
		elemthemeColorNin0 := make([]string, 0)
		for _, othemeColorNin0 := range *o.ThemeColorNin {
			elemthemeColorNin1 := othemeColorNin0
			elemthemeColorNin0 = append(elemthemeColorNin0, elemthemeColorNin1)
		}
		to.ThemeColorNin = &elemthemeColorNin0
	}
	if o.ThemeColorExists != nil {
		elemthemeColorExists0 := o.ThemeColorExists
		to.ThemeColorExists = elemthemeColorExists0
	}
	if o.ThemeColorLike != nil {
		elemthemeColorLike0 := o.ThemeColorLike
		to.ThemeColorLike = elemthemeColorLike0
	}
	if o.ThemeColorNlike != nil {
		elemthemeColorNlike0 := o.ThemeColorNlike
		to.ThemeColorNlike = elemthemeColorNlike0
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
	if o.WelcomeMessageEq != nil {
		elemwelcomeMessageEq0 := o.WelcomeMessageEq
		to.WelcomeMessageEq = elemwelcomeMessageEq0
	}
	if o.WelcomeMessageNe != nil {
		elemwelcomeMessageNe0 := o.WelcomeMessageNe
		to.WelcomeMessageNe = elemwelcomeMessageNe0
	}
	if o.WelcomeMessageGt != nil {
		elemwelcomeMessageGt0 := o.WelcomeMessageGt
		to.WelcomeMessageGt = elemwelcomeMessageGt0
	}
	if o.WelcomeMessageGte != nil {
		elemwelcomeMessageGte0 := o.WelcomeMessageGte
		to.WelcomeMessageGte = elemwelcomeMessageGte0
	}
	if o.WelcomeMessageLt != nil {
		elemwelcomeMessageLt0 := o.WelcomeMessageLt
		to.WelcomeMessageLt = elemwelcomeMessageLt0
	}
	if o.WelcomeMessageLte != nil {
		elemwelcomeMessageLte0 := o.WelcomeMessageLte
		to.WelcomeMessageLte = elemwelcomeMessageLte0
	}
	if o.WelcomeMessageIn != nil {
		elemwelcomeMessageIn0 := make([]string, 0)
		for _, owelcomeMessageIn0 := range *o.WelcomeMessageIn {
			elemwelcomeMessageIn1 := owelcomeMessageIn0
			elemwelcomeMessageIn0 = append(elemwelcomeMessageIn0, elemwelcomeMessageIn1)
		}
		to.WelcomeMessageIn = &elemwelcomeMessageIn0
	}
	if o.WelcomeMessageNin != nil {
		elemwelcomeMessageNin0 := make([]string, 0)
		for _, owelcomeMessageNin0 := range *o.WelcomeMessageNin {
			elemwelcomeMessageNin1 := owelcomeMessageNin0
			elemwelcomeMessageNin0 = append(elemwelcomeMessageNin0, elemwelcomeMessageNin1)
		}
		to.WelcomeMessageNin = &elemwelcomeMessageNin0
	}
	if o.WelcomeMessageExists != nil {
		elemwelcomeMessageExists0 := o.WelcomeMessageExists
		to.WelcomeMessageExists = elemwelcomeMessageExists0
	}
	if o.WelcomeMessageLike != nil {
		elemwelcomeMessageLike0 := o.WelcomeMessageLike
		to.WelcomeMessageLike = elemwelcomeMessageLike0
	}
	if o.WelcomeMessageNlike != nil {
		elemwelcomeMessageNlike0 := o.WelcomeMessageNlike
		to.WelcomeMessageNlike = elemwelcomeMessageNlike0
	}
	return to, nil
}

type HTTPSortParams struct {
	CreatedAt *int8 `json:"createdAt,omitempty"`
	OwnerId   *int8 `json:"ownerId,omitempty"`
	Slug      *int8 `json:"slug,omitempty"`
	UpdatedAt *int8 `json:"updatedAt,omitempty"`
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	if s.CreatedAt != nil {
		to.CreatedAt = *s.CreatedAt
	}
	if s.OwnerId != nil {
		to.OwnerId = *s.OwnerId
	}
	if s.Slug != nil {
		to.Slug = *s.Slug
	}
	if s.UpdatedAt != nil {
		to.UpdatedAt = *s.UpdatedAt
	}
	return to
}
