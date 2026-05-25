package registry

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_address_access_mode"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoRecord struct {
	Id                    *primitive.ObjectID             `bson:"_id,omitempty"`
	AddressAccessMode     *enum_address_access_mode.Value `bson:"addressAccessMode,omitempty"`
	CoverImageUrl         *string                         `bson:"coverImageUrl,omitempty"`
	Created               *actor_trace.MongoRecord        `bson:"created,omitempty"`
	DueDate               *time.Time                      `bson:"dueDate,omitempty"`
	IsPublic              *bool                           `bson:"isPublic,omitempty"`
	OwnerId               *primitive.ObjectID             `bson:"ownerId,omitempty"`
	ParentNames           *string                         `bson:"parentNames,omitempty"`
	ShippingCity          *string                         `bson:"shippingCity,omitempty"`
	ShippingCountry       *string                         `bson:"shippingCountry,omitempty"`
	ShippingDeliveryNotes *string                         `bson:"shippingDeliveryNotes,omitempty"`
	ShippingLine1         *string                         `bson:"shippingLine1,omitempty"`
	ShippingLine2         *string                         `bson:"shippingLine2,omitempty"`
	ShippingPolicyVersion *int                            `bson:"shippingPolicyVersion,omitempty"`
	ShippingPostalCode    *string                         `bson:"shippingPostalCode,omitempty"`
	ShippingRecipientName *string                         `bson:"shippingRecipientName,omitempty"`
	ShippingRegion        *string                         `bson:"shippingRegion,omitempty"`
	Slug                  *string                         `bson:"slug,omitempty"`
	ThemeColor            *string                         `bson:"themeColor,omitempty"`
	Title                 *string                         `bson:"title,omitempty"`
	Updated               *actor_trace.MongoRecord        `bson:"updated,omitempty"`
	UpdatedByOwnerUser    *actor_trace.MongoRecord        `bson:"updatedByOwnerUser,omitempty"`
	WelcomeMessage        *string                         `bson:"welcomeMessage,omitempty"`
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
		elemownerId0 := r.OwnerId.Hex()
		m.OwnerId = elemownerId0
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
	if o.AddressAccessModeEq != nil {
		query := bson.M{}
		query["addressAccessMode"] = o.AddressAccessModeEq
		and = append(and, query)
	}
	if o.AddressAccessModeNe != nil {
		query := bson.M{}
		query["addressAccessMode"] = bson.M{"$ne": o.AddressAccessModeNe}
		and = append(and, query)
	}
	if o.AddressAccessModeGt != nil {
		query := bson.M{}
		query["addressAccessMode"] = bson.M{"$gt": o.AddressAccessModeGt}
		and = append(and, query)
	}
	if o.AddressAccessModeGte != nil {
		query := bson.M{}
		query["addressAccessMode"] = bson.M{"$gte": o.AddressAccessModeGte}
		and = append(and, query)
	}
	if o.AddressAccessModeLt != nil {
		query := bson.M{}
		query["addressAccessMode"] = bson.M{"$lt": o.AddressAccessModeLt}
		and = append(and, query)
	}
	if o.AddressAccessModeLte != nil {
		query := bson.M{}
		query["addressAccessMode"] = bson.M{"$lte": o.AddressAccessModeLte}
		and = append(and, query)
	}
	if o.AddressAccessModeIn != nil {
		query := bson.M{}
		query["addressAccessMode"] = bson.M{"$in": o.AddressAccessModeIn}
		and = append(and, query)
	}
	if o.AddressAccessModeNin != nil {
		query := bson.M{}
		query["addressAccessMode"] = bson.M{"$nin": o.AddressAccessModeNin}
		and = append(and, query)
	}
	if o.AddressAccessModeExists != nil {
		query := bson.M{}
		query["addressAccessMode"] = bson.M{"$exists": *o.AddressAccessModeExists}
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
	if o.ShippingCityEq != nil {
		query := bson.M{}
		query["shippingCity"] = o.ShippingCityEq
		and = append(and, query)
	}
	if o.ShippingCityNe != nil {
		query := bson.M{}
		query["shippingCity"] = bson.M{"$ne": o.ShippingCityNe}
		and = append(and, query)
	}
	if o.ShippingCityGt != nil {
		query := bson.M{}
		query["shippingCity"] = bson.M{"$gt": o.ShippingCityGt}
		and = append(and, query)
	}
	if o.ShippingCityGte != nil {
		query := bson.M{}
		query["shippingCity"] = bson.M{"$gte": o.ShippingCityGte}
		and = append(and, query)
	}
	if o.ShippingCityLt != nil {
		query := bson.M{}
		query["shippingCity"] = bson.M{"$lt": o.ShippingCityLt}
		and = append(and, query)
	}
	if o.ShippingCityLte != nil {
		query := bson.M{}
		query["shippingCity"] = bson.M{"$lte": o.ShippingCityLte}
		and = append(and, query)
	}
	if o.ShippingCityIn != nil {
		query := bson.M{}
		query["shippingCity"] = bson.M{"$in": o.ShippingCityIn}
		and = append(and, query)
	}
	if o.ShippingCityNin != nil {
		query := bson.M{}
		query["shippingCity"] = bson.M{"$nin": o.ShippingCityNin}
		and = append(and, query)
	}
	if o.ShippingCityExists != nil {
		query := bson.M{}
		query["shippingCity"] = bson.M{"$exists": *o.ShippingCityExists}
		and = append(and, query)
	}
	if o.ShippingCityLike != nil {
		query := bson.M{}
		query["shippingCity"] = bson.M{"$regex": o.ShippingCityLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ShippingCityNlike != nil {
		query := bson.M{}
		query["shippingCity"] = bson.M{"$not": bson.M{"$regex": o.ShippingCityNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.ShippingCountryEq != nil {
		query := bson.M{}
		query["shippingCountry"] = o.ShippingCountryEq
		and = append(and, query)
	}
	if o.ShippingCountryNe != nil {
		query := bson.M{}
		query["shippingCountry"] = bson.M{"$ne": o.ShippingCountryNe}
		and = append(and, query)
	}
	if o.ShippingCountryGt != nil {
		query := bson.M{}
		query["shippingCountry"] = bson.M{"$gt": o.ShippingCountryGt}
		and = append(and, query)
	}
	if o.ShippingCountryGte != nil {
		query := bson.M{}
		query["shippingCountry"] = bson.M{"$gte": o.ShippingCountryGte}
		and = append(and, query)
	}
	if o.ShippingCountryLt != nil {
		query := bson.M{}
		query["shippingCountry"] = bson.M{"$lt": o.ShippingCountryLt}
		and = append(and, query)
	}
	if o.ShippingCountryLte != nil {
		query := bson.M{}
		query["shippingCountry"] = bson.M{"$lte": o.ShippingCountryLte}
		and = append(and, query)
	}
	if o.ShippingCountryIn != nil {
		query := bson.M{}
		query["shippingCountry"] = bson.M{"$in": o.ShippingCountryIn}
		and = append(and, query)
	}
	if o.ShippingCountryNin != nil {
		query := bson.M{}
		query["shippingCountry"] = bson.M{"$nin": o.ShippingCountryNin}
		and = append(and, query)
	}
	if o.ShippingCountryExists != nil {
		query := bson.M{}
		query["shippingCountry"] = bson.M{"$exists": *o.ShippingCountryExists}
		and = append(and, query)
	}
	if o.ShippingCountryLike != nil {
		query := bson.M{}
		query["shippingCountry"] = bson.M{"$regex": o.ShippingCountryLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ShippingCountryNlike != nil {
		query := bson.M{}
		query["shippingCountry"] = bson.M{"$not": bson.M{"$regex": o.ShippingCountryNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.ShippingDeliveryNotesEq != nil {
		query := bson.M{}
		query["shippingDeliveryNotes"] = o.ShippingDeliveryNotesEq
		and = append(and, query)
	}
	if o.ShippingDeliveryNotesNe != nil {
		query := bson.M{}
		query["shippingDeliveryNotes"] = bson.M{"$ne": o.ShippingDeliveryNotesNe}
		and = append(and, query)
	}
	if o.ShippingDeliveryNotesGt != nil {
		query := bson.M{}
		query["shippingDeliveryNotes"] = bson.M{"$gt": o.ShippingDeliveryNotesGt}
		and = append(and, query)
	}
	if o.ShippingDeliveryNotesGte != nil {
		query := bson.M{}
		query["shippingDeliveryNotes"] = bson.M{"$gte": o.ShippingDeliveryNotesGte}
		and = append(and, query)
	}
	if o.ShippingDeliveryNotesLt != nil {
		query := bson.M{}
		query["shippingDeliveryNotes"] = bson.M{"$lt": o.ShippingDeliveryNotesLt}
		and = append(and, query)
	}
	if o.ShippingDeliveryNotesLte != nil {
		query := bson.M{}
		query["shippingDeliveryNotes"] = bson.M{"$lte": o.ShippingDeliveryNotesLte}
		and = append(and, query)
	}
	if o.ShippingDeliveryNotesIn != nil {
		query := bson.M{}
		query["shippingDeliveryNotes"] = bson.M{"$in": o.ShippingDeliveryNotesIn}
		and = append(and, query)
	}
	if o.ShippingDeliveryNotesNin != nil {
		query := bson.M{}
		query["shippingDeliveryNotes"] = bson.M{"$nin": o.ShippingDeliveryNotesNin}
		and = append(and, query)
	}
	if o.ShippingDeliveryNotesExists != nil {
		query := bson.M{}
		query["shippingDeliveryNotes"] = bson.M{"$exists": *o.ShippingDeliveryNotesExists}
		and = append(and, query)
	}
	if o.ShippingDeliveryNotesLike != nil {
		query := bson.M{}
		query["shippingDeliveryNotes"] = bson.M{"$regex": o.ShippingDeliveryNotesLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ShippingDeliveryNotesNlike != nil {
		query := bson.M{}
		query["shippingDeliveryNotes"] = bson.M{"$not": bson.M{"$regex": o.ShippingDeliveryNotesNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.ShippingLine1Eq != nil {
		query := bson.M{}
		query["shippingLine1"] = o.ShippingLine1Eq
		and = append(and, query)
	}
	if o.ShippingLine1Ne != nil {
		query := bson.M{}
		query["shippingLine1"] = bson.M{"$ne": o.ShippingLine1Ne}
		and = append(and, query)
	}
	if o.ShippingLine1Gt != nil {
		query := bson.M{}
		query["shippingLine1"] = bson.M{"$gt": o.ShippingLine1Gt}
		and = append(and, query)
	}
	if o.ShippingLine1Gte != nil {
		query := bson.M{}
		query["shippingLine1"] = bson.M{"$gte": o.ShippingLine1Gte}
		and = append(and, query)
	}
	if o.ShippingLine1Lt != nil {
		query := bson.M{}
		query["shippingLine1"] = bson.M{"$lt": o.ShippingLine1Lt}
		and = append(and, query)
	}
	if o.ShippingLine1Lte != nil {
		query := bson.M{}
		query["shippingLine1"] = bson.M{"$lte": o.ShippingLine1Lte}
		and = append(and, query)
	}
	if o.ShippingLine1In != nil {
		query := bson.M{}
		query["shippingLine1"] = bson.M{"$in": o.ShippingLine1In}
		and = append(and, query)
	}
	if o.ShippingLine1Nin != nil {
		query := bson.M{}
		query["shippingLine1"] = bson.M{"$nin": o.ShippingLine1Nin}
		and = append(and, query)
	}
	if o.ShippingLine1Exists != nil {
		query := bson.M{}
		query["shippingLine1"] = bson.M{"$exists": *o.ShippingLine1Exists}
		and = append(and, query)
	}
	if o.ShippingLine1Like != nil {
		query := bson.M{}
		query["shippingLine1"] = bson.M{"$regex": o.ShippingLine1Like, "$options": "i"}
		and = append(and, query)
	}
	if o.ShippingLine1Nlike != nil {
		query := bson.M{}
		query["shippingLine1"] = bson.M{"$not": bson.M{"$regex": o.ShippingLine1Nlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.ShippingLine2Eq != nil {
		query := bson.M{}
		query["shippingLine2"] = o.ShippingLine2Eq
		and = append(and, query)
	}
	if o.ShippingLine2Ne != nil {
		query := bson.M{}
		query["shippingLine2"] = bson.M{"$ne": o.ShippingLine2Ne}
		and = append(and, query)
	}
	if o.ShippingLine2Gt != nil {
		query := bson.M{}
		query["shippingLine2"] = bson.M{"$gt": o.ShippingLine2Gt}
		and = append(and, query)
	}
	if o.ShippingLine2Gte != nil {
		query := bson.M{}
		query["shippingLine2"] = bson.M{"$gte": o.ShippingLine2Gte}
		and = append(and, query)
	}
	if o.ShippingLine2Lt != nil {
		query := bson.M{}
		query["shippingLine2"] = bson.M{"$lt": o.ShippingLine2Lt}
		and = append(and, query)
	}
	if o.ShippingLine2Lte != nil {
		query := bson.M{}
		query["shippingLine2"] = bson.M{"$lte": o.ShippingLine2Lte}
		and = append(and, query)
	}
	if o.ShippingLine2In != nil {
		query := bson.M{}
		query["shippingLine2"] = bson.M{"$in": o.ShippingLine2In}
		and = append(and, query)
	}
	if o.ShippingLine2Nin != nil {
		query := bson.M{}
		query["shippingLine2"] = bson.M{"$nin": o.ShippingLine2Nin}
		and = append(and, query)
	}
	if o.ShippingLine2Exists != nil {
		query := bson.M{}
		query["shippingLine2"] = bson.M{"$exists": *o.ShippingLine2Exists}
		and = append(and, query)
	}
	if o.ShippingLine2Like != nil {
		query := bson.M{}
		query["shippingLine2"] = bson.M{"$regex": o.ShippingLine2Like, "$options": "i"}
		and = append(and, query)
	}
	if o.ShippingLine2Nlike != nil {
		query := bson.M{}
		query["shippingLine2"] = bson.M{"$not": bson.M{"$regex": o.ShippingLine2Nlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.ShippingPolicyVersionEq != nil {
		query := bson.M{}
		query["shippingPolicyVersion"] = o.ShippingPolicyVersionEq
		and = append(and, query)
	}
	if o.ShippingPolicyVersionNe != nil {
		query := bson.M{}
		query["shippingPolicyVersion"] = bson.M{"$ne": o.ShippingPolicyVersionNe}
		and = append(and, query)
	}
	if o.ShippingPolicyVersionGt != nil {
		query := bson.M{}
		query["shippingPolicyVersion"] = bson.M{"$gt": o.ShippingPolicyVersionGt}
		and = append(and, query)
	}
	if o.ShippingPolicyVersionGte != nil {
		query := bson.M{}
		query["shippingPolicyVersion"] = bson.M{"$gte": o.ShippingPolicyVersionGte}
		and = append(and, query)
	}
	if o.ShippingPolicyVersionLt != nil {
		query := bson.M{}
		query["shippingPolicyVersion"] = bson.M{"$lt": o.ShippingPolicyVersionLt}
		and = append(and, query)
	}
	if o.ShippingPolicyVersionLte != nil {
		query := bson.M{}
		query["shippingPolicyVersion"] = bson.M{"$lte": o.ShippingPolicyVersionLte}
		and = append(and, query)
	}
	if o.ShippingPolicyVersionIn != nil {
		query := bson.M{}
		query["shippingPolicyVersion"] = bson.M{"$in": o.ShippingPolicyVersionIn}
		and = append(and, query)
	}
	if o.ShippingPolicyVersionNin != nil {
		query := bson.M{}
		query["shippingPolicyVersion"] = bson.M{"$nin": o.ShippingPolicyVersionNin}
		and = append(and, query)
	}
	if o.ShippingPolicyVersionExists != nil {
		query := bson.M{}
		query["shippingPolicyVersion"] = bson.M{"$exists": *o.ShippingPolicyVersionExists}
		and = append(and, query)
	}
	if o.ShippingPostalCodeEq != nil {
		query := bson.M{}
		query["shippingPostalCode"] = o.ShippingPostalCodeEq
		and = append(and, query)
	}
	if o.ShippingPostalCodeNe != nil {
		query := bson.M{}
		query["shippingPostalCode"] = bson.M{"$ne": o.ShippingPostalCodeNe}
		and = append(and, query)
	}
	if o.ShippingPostalCodeGt != nil {
		query := bson.M{}
		query["shippingPostalCode"] = bson.M{"$gt": o.ShippingPostalCodeGt}
		and = append(and, query)
	}
	if o.ShippingPostalCodeGte != nil {
		query := bson.M{}
		query["shippingPostalCode"] = bson.M{"$gte": o.ShippingPostalCodeGte}
		and = append(and, query)
	}
	if o.ShippingPostalCodeLt != nil {
		query := bson.M{}
		query["shippingPostalCode"] = bson.M{"$lt": o.ShippingPostalCodeLt}
		and = append(and, query)
	}
	if o.ShippingPostalCodeLte != nil {
		query := bson.M{}
		query["shippingPostalCode"] = bson.M{"$lte": o.ShippingPostalCodeLte}
		and = append(and, query)
	}
	if o.ShippingPostalCodeIn != nil {
		query := bson.M{}
		query["shippingPostalCode"] = bson.M{"$in": o.ShippingPostalCodeIn}
		and = append(and, query)
	}
	if o.ShippingPostalCodeNin != nil {
		query := bson.M{}
		query["shippingPostalCode"] = bson.M{"$nin": o.ShippingPostalCodeNin}
		and = append(and, query)
	}
	if o.ShippingPostalCodeExists != nil {
		query := bson.M{}
		query["shippingPostalCode"] = bson.M{"$exists": *o.ShippingPostalCodeExists}
		and = append(and, query)
	}
	if o.ShippingPostalCodeLike != nil {
		query := bson.M{}
		query["shippingPostalCode"] = bson.M{"$regex": o.ShippingPostalCodeLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ShippingPostalCodeNlike != nil {
		query := bson.M{}
		query["shippingPostalCode"] = bson.M{"$not": bson.M{"$regex": o.ShippingPostalCodeNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.ShippingRecipientNameEq != nil {
		query := bson.M{}
		query["shippingRecipientName"] = o.ShippingRecipientNameEq
		and = append(and, query)
	}
	if o.ShippingRecipientNameNe != nil {
		query := bson.M{}
		query["shippingRecipientName"] = bson.M{"$ne": o.ShippingRecipientNameNe}
		and = append(and, query)
	}
	if o.ShippingRecipientNameGt != nil {
		query := bson.M{}
		query["shippingRecipientName"] = bson.M{"$gt": o.ShippingRecipientNameGt}
		and = append(and, query)
	}
	if o.ShippingRecipientNameGte != nil {
		query := bson.M{}
		query["shippingRecipientName"] = bson.M{"$gte": o.ShippingRecipientNameGte}
		and = append(and, query)
	}
	if o.ShippingRecipientNameLt != nil {
		query := bson.M{}
		query["shippingRecipientName"] = bson.M{"$lt": o.ShippingRecipientNameLt}
		and = append(and, query)
	}
	if o.ShippingRecipientNameLte != nil {
		query := bson.M{}
		query["shippingRecipientName"] = bson.M{"$lte": o.ShippingRecipientNameLte}
		and = append(and, query)
	}
	if o.ShippingRecipientNameIn != nil {
		query := bson.M{}
		query["shippingRecipientName"] = bson.M{"$in": o.ShippingRecipientNameIn}
		and = append(and, query)
	}
	if o.ShippingRecipientNameNin != nil {
		query := bson.M{}
		query["shippingRecipientName"] = bson.M{"$nin": o.ShippingRecipientNameNin}
		and = append(and, query)
	}
	if o.ShippingRecipientNameExists != nil {
		query := bson.M{}
		query["shippingRecipientName"] = bson.M{"$exists": *o.ShippingRecipientNameExists}
		and = append(and, query)
	}
	if o.ShippingRecipientNameLike != nil {
		query := bson.M{}
		query["shippingRecipientName"] = bson.M{"$regex": o.ShippingRecipientNameLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ShippingRecipientNameNlike != nil {
		query := bson.M{}
		query["shippingRecipientName"] = bson.M{"$not": bson.M{"$regex": o.ShippingRecipientNameNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.ShippingRegionEq != nil {
		query := bson.M{}
		query["shippingRegion"] = o.ShippingRegionEq
		and = append(and, query)
	}
	if o.ShippingRegionNe != nil {
		query := bson.M{}
		query["shippingRegion"] = bson.M{"$ne": o.ShippingRegionNe}
		and = append(and, query)
	}
	if o.ShippingRegionGt != nil {
		query := bson.M{}
		query["shippingRegion"] = bson.M{"$gt": o.ShippingRegionGt}
		and = append(and, query)
	}
	if o.ShippingRegionGte != nil {
		query := bson.M{}
		query["shippingRegion"] = bson.M{"$gte": o.ShippingRegionGte}
		and = append(and, query)
	}
	if o.ShippingRegionLt != nil {
		query := bson.M{}
		query["shippingRegion"] = bson.M{"$lt": o.ShippingRegionLt}
		and = append(and, query)
	}
	if o.ShippingRegionLte != nil {
		query := bson.M{}
		query["shippingRegion"] = bson.M{"$lte": o.ShippingRegionLte}
		and = append(and, query)
	}
	if o.ShippingRegionIn != nil {
		query := bson.M{}
		query["shippingRegion"] = bson.M{"$in": o.ShippingRegionIn}
		and = append(and, query)
	}
	if o.ShippingRegionNin != nil {
		query := bson.M{}
		query["shippingRegion"] = bson.M{"$nin": o.ShippingRegionNin}
		and = append(and, query)
	}
	if o.ShippingRegionExists != nil {
		query := bson.M{}
		query["shippingRegion"] = bson.M{"$exists": *o.ShippingRegionExists}
		and = append(and, query)
	}
	if o.ShippingRegionLike != nil {
		query := bson.M{}
		query["shippingRegion"] = bson.M{"$regex": o.ShippingRegionLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ShippingRegionNlike != nil {
		query := bson.M{}
		query["shippingRegion"] = bson.M{"$not": bson.M{"$regex": o.ShippingRegionNlike, "$options": "i"}}
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
