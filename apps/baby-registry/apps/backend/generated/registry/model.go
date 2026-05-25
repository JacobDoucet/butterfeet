package registry

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Model struct {
	Id                 string
	CoverImageUrl      string
	Created            actor_trace.Model
	DueDate            time.Time
	IsPublic           bool
	OwnerId            string
	ParentNames        string
	Slug               string
	ThemeColor         string
	Title              string
	Updated            actor_trace.Model
	UpdatedByOwnerUser actor_trace.Model
	WelcomeMessage     string
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
