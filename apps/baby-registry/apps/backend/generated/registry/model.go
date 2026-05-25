package registry

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_address_access_mode"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Model struct {
	Id                    string
	AddressAccessMode     enum_address_access_mode.Value
	CoverImageUrl         string
	Created               actor_trace.Model
	DueDate               time.Time
	IsPublic              bool
	OwnerId               string
	ParentNames           string
	ShippingCity          string
	ShippingCountry       string
	ShippingDeliveryNotes string
	ShippingLine1         string
	ShippingLine2         string
	ShippingPolicyVersion int
	ShippingPostalCode    string
	ShippingRecipientName string
	ShippingRegion        string
	Slug                  string
	ThemeColor            string
	Title                 string
	Updated               actor_trace.Model
	UpdatedByOwnerUser    actor_trace.Model
	WelcomeMessage        string
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
	if projection.AddressAccessMode {
		elemaddressAccessMode0 := m.AddressAccessMode
		r.AddressAccessMode = &elemaddressAccessMode0
	}
	if projection.CoverImageUrl {
		elemcoverImageUrl0 := m.CoverImageUrl
		r.CoverImageUrl = &elemcoverImageUrl0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToMongoRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.DueDate {
		elemdueDate0 := m.DueDate
		r.DueDate = &elemdueDate0
	}
	if projection.IsPublic {
		elemisPublic0 := m.IsPublic
		r.IsPublic = &elemisPublic0
	}
	if projection.OwnerId && m.OwnerId != "" {
		elemownerId0, err := primitive.ObjectIDFromHex(m.OwnerId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.OwnerId"), err)
		}
		r.OwnerId = &elemownerId0
	}
	if projection.ParentNames {
		elemparentNames0 := m.ParentNames
		r.ParentNames = &elemparentNames0
	}
	if projection.ShippingCity {
		elemshippingCity0 := m.ShippingCity
		r.ShippingCity = &elemshippingCity0
	}
	if projection.ShippingCountry {
		elemshippingCountry0 := m.ShippingCountry
		r.ShippingCountry = &elemshippingCountry0
	}
	if projection.ShippingDeliveryNotes {
		elemshippingDeliveryNotes0 := m.ShippingDeliveryNotes
		r.ShippingDeliveryNotes = &elemshippingDeliveryNotes0
	}
	if projection.ShippingLine1 {
		elemshippingLine10 := m.ShippingLine1
		r.ShippingLine1 = &elemshippingLine10
	}
	if projection.ShippingLine2 {
		elemshippingLine20 := m.ShippingLine2
		r.ShippingLine2 = &elemshippingLine20
	}
	if projection.ShippingPolicyVersion {
		elemshippingPolicyVersion0 := m.ShippingPolicyVersion
		r.ShippingPolicyVersion = &elemshippingPolicyVersion0
	}
	if projection.ShippingPostalCode {
		elemshippingPostalCode0 := m.ShippingPostalCode
		r.ShippingPostalCode = &elemshippingPostalCode0
	}
	if projection.ShippingRecipientName {
		elemshippingRecipientName0 := m.ShippingRecipientName
		r.ShippingRecipientName = &elemshippingRecipientName0
	}
	if projection.ShippingRegion {
		elemshippingRegion0 := m.ShippingRegion
		r.ShippingRegion = &elemshippingRegion0
	}
	if projection.Slug {
		elemslug0 := m.Slug
		r.Slug = &elemslug0
	}
	if projection.ThemeColor {
		elemthemeColor0 := m.ThemeColor
		r.ThemeColor = &elemthemeColor0
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
	if projection.WelcomeMessage {
		elemwelcomeMessage0 := m.WelcomeMessage
		r.WelcomeMessage = &elemwelcomeMessage0
	}
	return r, nil
}

func (m *Model) ToHTTPRecord(projection Projection) (HTTPRecord, error) {
	r := HTTPRecord{}
	if m.Id != "" {
		elemid0 := m.Id
		r.Id = &elemid0
	}
	if projection.AddressAccessMode {
		elemaddressAccessMode0 := m.AddressAccessMode
		r.AddressAccessMode = &elemaddressAccessMode0
	}
	if projection.CoverImageUrl {
		elemcoverImageUrl0 := m.CoverImageUrl
		r.CoverImageUrl = &elemcoverImageUrl0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToHTTPRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.DueDate {
		elemdueDate0 := m.DueDate
		r.DueDate = &elemdueDate0
	}
	if projection.IsPublic {
		elemisPublic0 := m.IsPublic
		r.IsPublic = &elemisPublic0
	}
	if projection.OwnerId && m.OwnerId != "" {
		elemownerId0 := m.OwnerId
		r.OwnerId = &elemownerId0
	}
	if projection.ParentNames {
		elemparentNames0 := m.ParentNames
		r.ParentNames = &elemparentNames0
	}
	if projection.ShippingCity {
		elemshippingCity0 := m.ShippingCity
		r.ShippingCity = &elemshippingCity0
	}
	if projection.ShippingCountry {
		elemshippingCountry0 := m.ShippingCountry
		r.ShippingCountry = &elemshippingCountry0
	}
	if projection.ShippingDeliveryNotes {
		elemshippingDeliveryNotes0 := m.ShippingDeliveryNotes
		r.ShippingDeliveryNotes = &elemshippingDeliveryNotes0
	}
	if projection.ShippingLine1 {
		elemshippingLine10 := m.ShippingLine1
		r.ShippingLine1 = &elemshippingLine10
	}
	if projection.ShippingLine2 {
		elemshippingLine20 := m.ShippingLine2
		r.ShippingLine2 = &elemshippingLine20
	}
	if projection.ShippingPolicyVersion {
		elemshippingPolicyVersion0 := m.ShippingPolicyVersion
		r.ShippingPolicyVersion = &elemshippingPolicyVersion0
	}
	if projection.ShippingPostalCode {
		elemshippingPostalCode0 := m.ShippingPostalCode
		r.ShippingPostalCode = &elemshippingPostalCode0
	}
	if projection.ShippingRecipientName {
		elemshippingRecipientName0 := m.ShippingRecipientName
		r.ShippingRecipientName = &elemshippingRecipientName0
	}
	if projection.ShippingRegion {
		elemshippingRegion0 := m.ShippingRegion
		r.ShippingRegion = &elemshippingRegion0
	}
	if projection.Slug {
		elemslug0 := m.Slug
		r.Slug = &elemslug0
	}
	if projection.ThemeColor {
		elemthemeColor0 := m.ThemeColor
		r.ThemeColor = &elemthemeColor0
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
	if projection.WelcomeMessage {
		elemwelcomeMessage0 := m.WelcomeMessage
		r.WelcomeMessage = &elemwelcomeMessage0
	}
	return r, nil
}

type SelectByIdQuery struct {
	Id string
}
type SelectBySlugUniqueQuery struct {
	Slug string
}

type WhereClause struct {
	// id (Ref<Registry>) search options
	IdEq     *string
	IdIn     *[]string
	IdNin    *[]string
	IdExists *bool
	// addressAccessMode (AddressAccessMode) search options
	AddressAccessModeEq     *enum_address_access_mode.Value
	AddressAccessModeNe     *enum_address_access_mode.Value
	AddressAccessModeGt     *enum_address_access_mode.Value
	AddressAccessModeGte    *enum_address_access_mode.Value
	AddressAccessModeLt     *enum_address_access_mode.Value
	AddressAccessModeLte    *enum_address_access_mode.Value
	AddressAccessModeIn     *[]enum_address_access_mode.Value
	AddressAccessModeNin    *[]enum_address_access_mode.Value
	AddressAccessModeExists *bool
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
	Created *actor_trace.WhereClause
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
	OwnerIdEq     *string
	OwnerIdIn     *[]string
	OwnerIdNin    *[]string
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
	// shippingCity (string) search options
	ShippingCityEq     *string
	ShippingCityNe     *string
	ShippingCityGt     *string
	ShippingCityGte    *string
	ShippingCityLt     *string
	ShippingCityLte    *string
	ShippingCityIn     *[]string
	ShippingCityNin    *[]string
	ShippingCityExists *bool
	ShippingCityLike   *string
	ShippingCityNlike  *string
	// shippingCountry (string) search options
	ShippingCountryEq     *string
	ShippingCountryNe     *string
	ShippingCountryGt     *string
	ShippingCountryGte    *string
	ShippingCountryLt     *string
	ShippingCountryLte    *string
	ShippingCountryIn     *[]string
	ShippingCountryNin    *[]string
	ShippingCountryExists *bool
	ShippingCountryLike   *string
	ShippingCountryNlike  *string
	// shippingDeliveryNotes (string) search options
	ShippingDeliveryNotesEq     *string
	ShippingDeliveryNotesNe     *string
	ShippingDeliveryNotesGt     *string
	ShippingDeliveryNotesGte    *string
	ShippingDeliveryNotesLt     *string
	ShippingDeliveryNotesLte    *string
	ShippingDeliveryNotesIn     *[]string
	ShippingDeliveryNotesNin    *[]string
	ShippingDeliveryNotesExists *bool
	ShippingDeliveryNotesLike   *string
	ShippingDeliveryNotesNlike  *string
	// shippingLine1 (string) search options
	ShippingLine1Eq     *string
	ShippingLine1Ne     *string
	ShippingLine1Gt     *string
	ShippingLine1Gte    *string
	ShippingLine1Lt     *string
	ShippingLine1Lte    *string
	ShippingLine1In     *[]string
	ShippingLine1Nin    *[]string
	ShippingLine1Exists *bool
	ShippingLine1Like   *string
	ShippingLine1Nlike  *string
	// shippingLine2 (string) search options
	ShippingLine2Eq     *string
	ShippingLine2Ne     *string
	ShippingLine2Gt     *string
	ShippingLine2Gte    *string
	ShippingLine2Lt     *string
	ShippingLine2Lte    *string
	ShippingLine2In     *[]string
	ShippingLine2Nin    *[]string
	ShippingLine2Exists *bool
	ShippingLine2Like   *string
	ShippingLine2Nlike  *string
	// shippingPolicyVersion (int) search options
	ShippingPolicyVersionEq     *int
	ShippingPolicyVersionNe     *int
	ShippingPolicyVersionGt     *int
	ShippingPolicyVersionGte    *int
	ShippingPolicyVersionLt     *int
	ShippingPolicyVersionLte    *int
	ShippingPolicyVersionIn     *[]int
	ShippingPolicyVersionNin    *[]int
	ShippingPolicyVersionExists *bool
	// shippingPostalCode (string) search options
	ShippingPostalCodeEq     *string
	ShippingPostalCodeNe     *string
	ShippingPostalCodeGt     *string
	ShippingPostalCodeGte    *string
	ShippingPostalCodeLt     *string
	ShippingPostalCodeLte    *string
	ShippingPostalCodeIn     *[]string
	ShippingPostalCodeNin    *[]string
	ShippingPostalCodeExists *bool
	ShippingPostalCodeLike   *string
	ShippingPostalCodeNlike  *string
	// shippingRecipientName (string) search options
	ShippingRecipientNameEq     *string
	ShippingRecipientNameNe     *string
	ShippingRecipientNameGt     *string
	ShippingRecipientNameGte    *string
	ShippingRecipientNameLt     *string
	ShippingRecipientNameLte    *string
	ShippingRecipientNameIn     *[]string
	ShippingRecipientNameNin    *[]string
	ShippingRecipientNameExists *bool
	ShippingRecipientNameLike   *string
	ShippingRecipientNameNlike  *string
	// shippingRegion (string) search options
	ShippingRegionEq     *string
	ShippingRegionNe     *string
	ShippingRegionGt     *string
	ShippingRegionGte    *string
	ShippingRegionLt     *string
	ShippingRegionLte    *string
	ShippingRegionIn     *[]string
	ShippingRegionNin    *[]string
	ShippingRegionExists *bool
	ShippingRegionLike   *string
	ShippingRegionNlike  *string
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
	Updated *actor_trace.WhereClause
	// updatedByOwnerUser (ActorTrace) search options
	UpdatedByOwnerUser *actor_trace.WhereClause
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

func (o SelectByIdQuery) ToMongoSelectByIdQuery() (MongoSelectByIdQuery, error) {
	to := MongoSelectByIdQuery{}
	elemid0, err := primitive.ObjectIDFromHex(o.Id)
	if err != nil {
		return to, errors.Join(errors.New("invalid o.Id"), err)
	}
	to.Id = elemid0
	return to, nil
}
func (o SelectBySlugUniqueQuery) ToMongoSelectBySlugUniqueQuery() (MongoSelectBySlugUniqueQuery, error) {
	to := MongoSelectBySlugUniqueQuery{}
	elemslug0 := o.Slug
	to.Slug = elemslug0
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
		elemcreated0, err := o.Created.ToMongoWhereClause()
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
		elemownerIdEq0, err := primitive.ObjectIDFromHex(*o.OwnerIdEq)
		if err != nil {
			return to, errors.Join(errors.New("invalid o.OwnerIdEq"), err)
		}
		to.OwnerIdEq = &elemownerIdEq0
	}
	if o.OwnerIdIn != nil {
		elemownerIdIn0 := make([]primitive.ObjectID, 0)
		for _, oownerIdIn0 := range *o.OwnerIdIn {
			elemownerIdIn1, err := primitive.ObjectIDFromHex(oownerIdIn0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oownerIdIn0"), err)
			}
			elemownerIdIn0 = append(elemownerIdIn0, elemownerIdIn1)
		}
		to.OwnerIdIn = &elemownerIdIn0
	}
	if o.OwnerIdNin != nil {
		elemownerIdNin0 := make([]primitive.ObjectID, 0)
		for _, oownerIdNin0 := range *o.OwnerIdNin {
			elemownerIdNin1, err := primitive.ObjectIDFromHex(oownerIdNin0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oownerIdNin0"), err)
			}
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

type SortParams struct {
	CreatedAt int8
	OwnerId   int8
	Slug      int8
	UpdatedAt int8
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	to.CreatedAt = s.CreatedAt
	to.OwnerId = s.OwnerId
	to.Slug = s.Slug
	to.UpdatedAt = s.UpdatedAt
	return to
}
