package cart

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_cart_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_payment_method_type"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Model struct {
	Id                 string
	AmountCents        int
	ClaimedAt          time.Time
	ContributorEmail   string
	ContributorName    string
	Created            actor_trace.Model
	Currency           string
	DecidedAt          time.Time
	DecisionReason     string
	Message            string
	MethodDisplayName  string
	MethodType         enum_payment_method_type.Value
	OwnerId            string
	PaymentMethodId    string
	ReferenceCode      string
	RegistryId         string
	Status             enum_cart_status.Value
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
	if projection.AmountCents {
		elemamountCents0 := m.AmountCents
		r.AmountCents = &elemamountCents0
	}
	if projection.ClaimedAt {
		elemclaimedAt0 := m.ClaimedAt
		r.ClaimedAt = &elemclaimedAt0
	}
	if projection.ContributorEmail {
		elemcontributorEmail0 := m.ContributorEmail
		r.ContributorEmail = &elemcontributorEmail0
	}
	if projection.ContributorName {
		elemcontributorName0 := m.ContributorName
		r.ContributorName = &elemcontributorName0
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
	if projection.DecidedAt {
		elemdecidedAt0 := m.DecidedAt
		r.DecidedAt = &elemdecidedAt0
	}
	if projection.DecisionReason {
		elemdecisionReason0 := m.DecisionReason
		r.DecisionReason = &elemdecisionReason0
	}
	if projection.Message {
		elemmessage0 := m.Message
		r.Message = &elemmessage0
	}
	if projection.MethodDisplayName {
		elemmethodDisplayName0 := m.MethodDisplayName
		r.MethodDisplayName = &elemmethodDisplayName0
	}
	if projection.MethodType {
		elemmethodType0 := m.MethodType
		r.MethodType = &elemmethodType0
	}
	if projection.OwnerId && m.OwnerId != "" {
		elemownerId0, err := primitive.ObjectIDFromHex(m.OwnerId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.OwnerId"), err)
		}
		r.OwnerId = &elemownerId0
	}
	if projection.PaymentMethodId && m.PaymentMethodId != "" {
		elempaymentMethodId0, err := primitive.ObjectIDFromHex(m.PaymentMethodId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.PaymentMethodId"), err)
		}
		r.PaymentMethodId = &elempaymentMethodId0
	}
	if projection.ReferenceCode {
		elemreferenceCode0 := m.ReferenceCode
		r.ReferenceCode = &elemreferenceCode0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0, err := primitive.ObjectIDFromHex(m.RegistryId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.RegistryId"), err)
		}
		r.RegistryId = &elemregistryId0
	}
	if projection.Status {
		elemstatus0 := m.Status
		r.Status = &elemstatus0
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
	if projection.AmountCents {
		elemamountCents0 := m.AmountCents
		r.AmountCents = &elemamountCents0
	}
	if projection.ClaimedAt {
		elemclaimedAt0 := m.ClaimedAt
		r.ClaimedAt = &elemclaimedAt0
	}
	if projection.ContributorEmail {
		elemcontributorEmail0 := m.ContributorEmail
		r.ContributorEmail = &elemcontributorEmail0
	}
	if projection.ContributorName {
		elemcontributorName0 := m.ContributorName
		r.ContributorName = &elemcontributorName0
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
	if projection.DecidedAt {
		elemdecidedAt0 := m.DecidedAt
		r.DecidedAt = &elemdecidedAt0
	}
	if projection.DecisionReason {
		elemdecisionReason0 := m.DecisionReason
		r.DecisionReason = &elemdecisionReason0
	}
	if projection.Message {
		elemmessage0 := m.Message
		r.Message = &elemmessage0
	}
	if projection.MethodDisplayName {
		elemmethodDisplayName0 := m.MethodDisplayName
		r.MethodDisplayName = &elemmethodDisplayName0
	}
	if projection.MethodType {
		elemmethodType0 := m.MethodType
		r.MethodType = &elemmethodType0
	}
	if projection.OwnerId && m.OwnerId != "" {
		elemownerId0 := m.OwnerId
		r.OwnerId = &elemownerId0
	}
	if projection.PaymentMethodId && m.PaymentMethodId != "" {
		elempaymentMethodId0 := m.PaymentMethodId
		r.PaymentMethodId = &elempaymentMethodId0
	}
	if projection.ReferenceCode {
		elemreferenceCode0 := m.ReferenceCode
		r.ReferenceCode = &elemreferenceCode0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0 := m.RegistryId
		r.RegistryId = &elemregistryId0
	}
	if projection.Status {
		elemstatus0 := m.Status
		r.Status = &elemstatus0
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
type SelectByReferenceUniqueQuery struct {
	ReferenceCode string
}

type WhereClause struct {
	// id (Ref<Cart>) search options
	IdEq     *string
	IdIn     *[]string
	IdNin    *[]string
	IdExists *bool
	// amountCents (int) search options
	AmountCentsEq     *int
	AmountCentsNe     *int
	AmountCentsGt     *int
	AmountCentsGte    *int
	AmountCentsLt     *int
	AmountCentsLte    *int
	AmountCentsIn     *[]int
	AmountCentsNin    *[]int
	AmountCentsExists *bool
	// claimedAt (timestamp) search options
	ClaimedAtEq     *time.Time
	ClaimedAtNe     *time.Time
	ClaimedAtGt     *time.Time
	ClaimedAtGte    *time.Time
	ClaimedAtLt     *time.Time
	ClaimedAtLte    *time.Time
	ClaimedAtIn     *[]time.Time
	ClaimedAtNin    *[]time.Time
	ClaimedAtExists *bool
	// contributorEmail (string) search options
	ContributorEmailEq     *string
	ContributorEmailNe     *string
	ContributorEmailGt     *string
	ContributorEmailGte    *string
	ContributorEmailLt     *string
	ContributorEmailLte    *string
	ContributorEmailIn     *[]string
	ContributorEmailNin    *[]string
	ContributorEmailExists *bool
	ContributorEmailLike   *string
	ContributorEmailNlike  *string
	// contributorName (string) search options
	ContributorNameEq     *string
	ContributorNameNe     *string
	ContributorNameGt     *string
	ContributorNameGte    *string
	ContributorNameLt     *string
	ContributorNameLte    *string
	ContributorNameIn     *[]string
	ContributorNameNin    *[]string
	ContributorNameExists *bool
	ContributorNameLike   *string
	ContributorNameNlike  *string
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
	// decidedAt (timestamp) search options
	DecidedAtEq     *time.Time
	DecidedAtNe     *time.Time
	DecidedAtGt     *time.Time
	DecidedAtGte    *time.Time
	DecidedAtLt     *time.Time
	DecidedAtLte    *time.Time
	DecidedAtIn     *[]time.Time
	DecidedAtNin    *[]time.Time
	DecidedAtExists *bool
	// decisionReason (string) search options
	DecisionReasonEq     *string
	DecisionReasonNe     *string
	DecisionReasonGt     *string
	DecisionReasonGte    *string
	DecisionReasonLt     *string
	DecisionReasonLte    *string
	DecisionReasonIn     *[]string
	DecisionReasonNin    *[]string
	DecisionReasonExists *bool
	DecisionReasonLike   *string
	DecisionReasonNlike  *string
	// message (string) search options
	MessageEq     *string
	MessageNe     *string
	MessageGt     *string
	MessageGte    *string
	MessageLt     *string
	MessageLte    *string
	MessageIn     *[]string
	MessageNin    *[]string
	MessageExists *bool
	MessageLike   *string
	MessageNlike  *string
	// methodDisplayName (string) search options
	MethodDisplayNameEq     *string
	MethodDisplayNameNe     *string
	MethodDisplayNameGt     *string
	MethodDisplayNameGte    *string
	MethodDisplayNameLt     *string
	MethodDisplayNameLte    *string
	MethodDisplayNameIn     *[]string
	MethodDisplayNameNin    *[]string
	MethodDisplayNameExists *bool
	MethodDisplayNameLike   *string
	MethodDisplayNameNlike  *string
	// methodType (PaymentMethodType) search options
	MethodTypeEq     *enum_payment_method_type.Value
	MethodTypeNe     *enum_payment_method_type.Value
	MethodTypeGt     *enum_payment_method_type.Value
	MethodTypeGte    *enum_payment_method_type.Value
	MethodTypeLt     *enum_payment_method_type.Value
	MethodTypeLte    *enum_payment_method_type.Value
	MethodTypeIn     *[]enum_payment_method_type.Value
	MethodTypeNin    *[]enum_payment_method_type.Value
	MethodTypeExists *bool
	// ownerId (Ref<OwnerUser>) search options
	OwnerIdEq     *string
	OwnerIdIn     *[]string
	OwnerIdNin    *[]string
	OwnerIdExists *bool
	// paymentMethodId (Ref<RegistryPaymentMethod>) search options
	PaymentMethodIdEq     *string
	PaymentMethodIdIn     *[]string
	PaymentMethodIdNin    *[]string
	PaymentMethodIdExists *bool
	// referenceCode (string) search options
	ReferenceCodeEq     *string
	ReferenceCodeNe     *string
	ReferenceCodeGt     *string
	ReferenceCodeGte    *string
	ReferenceCodeLt     *string
	ReferenceCodeLte    *string
	ReferenceCodeIn     *[]string
	ReferenceCodeNin    *[]string
	ReferenceCodeExists *bool
	ReferenceCodeLike   *string
	ReferenceCodeNlike  *string
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *string
	RegistryIdIn     *[]string
	RegistryIdNin    *[]string
	RegistryIdExists *bool
	// status (CartStatus) search options
	StatusEq     *enum_cart_status.Value
	StatusNe     *enum_cart_status.Value
	StatusGt     *enum_cart_status.Value
	StatusGte    *enum_cart_status.Value
	StatusLt     *enum_cart_status.Value
	StatusLte    *enum_cart_status.Value
	StatusIn     *[]enum_cart_status.Value
	StatusNin    *[]enum_cart_status.Value
	StatusExists *bool
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
func (o SelectByReferenceUniqueQuery) ToMongoSelectByReferenceUniqueQuery() (MongoSelectByReferenceUniqueQuery, error) {
	to := MongoSelectByReferenceUniqueQuery{}
	elemreferenceCode0 := o.ReferenceCode
	to.ReferenceCode = elemreferenceCode0
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
	if o.AmountCentsEq != nil {
		elemamountCentsEq0 := o.AmountCentsEq
		to.AmountCentsEq = elemamountCentsEq0
	}
	if o.AmountCentsNe != nil {
		elemamountCentsNe0 := o.AmountCentsNe
		to.AmountCentsNe = elemamountCentsNe0
	}
	if o.AmountCentsGt != nil {
		elemamountCentsGt0 := o.AmountCentsGt
		to.AmountCentsGt = elemamountCentsGt0
	}
	if o.AmountCentsGte != nil {
		elemamountCentsGte0 := o.AmountCentsGte
		to.AmountCentsGte = elemamountCentsGte0
	}
	if o.AmountCentsLt != nil {
		elemamountCentsLt0 := o.AmountCentsLt
		to.AmountCentsLt = elemamountCentsLt0
	}
	if o.AmountCentsLte != nil {
		elemamountCentsLte0 := o.AmountCentsLte
		to.AmountCentsLte = elemamountCentsLte0
	}
	if o.AmountCentsIn != nil {
		elemamountCentsIn0 := make([]int, 0)
		for _, oamountCentsIn0 := range *o.AmountCentsIn {
			elemamountCentsIn1 := oamountCentsIn0
			elemamountCentsIn0 = append(elemamountCentsIn0, elemamountCentsIn1)
		}
		to.AmountCentsIn = &elemamountCentsIn0
	}
	if o.AmountCentsNin != nil {
		elemamountCentsNin0 := make([]int, 0)
		for _, oamountCentsNin0 := range *o.AmountCentsNin {
			elemamountCentsNin1 := oamountCentsNin0
			elemamountCentsNin0 = append(elemamountCentsNin0, elemamountCentsNin1)
		}
		to.AmountCentsNin = &elemamountCentsNin0
	}
	if o.AmountCentsExists != nil {
		elemamountCentsExists0 := o.AmountCentsExists
		to.AmountCentsExists = elemamountCentsExists0
	}
	if o.ClaimedAtEq != nil {
		elemclaimedAtEq0 := o.ClaimedAtEq
		to.ClaimedAtEq = elemclaimedAtEq0
	}
	if o.ClaimedAtNe != nil {
		elemclaimedAtNe0 := o.ClaimedAtNe
		to.ClaimedAtNe = elemclaimedAtNe0
	}
	if o.ClaimedAtGt != nil {
		elemclaimedAtGt0 := o.ClaimedAtGt
		to.ClaimedAtGt = elemclaimedAtGt0
	}
	if o.ClaimedAtGte != nil {
		elemclaimedAtGte0 := o.ClaimedAtGte
		to.ClaimedAtGte = elemclaimedAtGte0
	}
	if o.ClaimedAtLt != nil {
		elemclaimedAtLt0 := o.ClaimedAtLt
		to.ClaimedAtLt = elemclaimedAtLt0
	}
	if o.ClaimedAtLte != nil {
		elemclaimedAtLte0 := o.ClaimedAtLte
		to.ClaimedAtLte = elemclaimedAtLte0
	}
	if o.ClaimedAtIn != nil {
		elemclaimedAtIn0 := make([]time.Time, 0)
		for _, oclaimedAtIn0 := range *o.ClaimedAtIn {
			elemclaimedAtIn1 := oclaimedAtIn0
			elemclaimedAtIn0 = append(elemclaimedAtIn0, elemclaimedAtIn1)
		}
		to.ClaimedAtIn = &elemclaimedAtIn0
	}
	if o.ClaimedAtNin != nil {
		elemclaimedAtNin0 := make([]time.Time, 0)
		for _, oclaimedAtNin0 := range *o.ClaimedAtNin {
			elemclaimedAtNin1 := oclaimedAtNin0
			elemclaimedAtNin0 = append(elemclaimedAtNin0, elemclaimedAtNin1)
		}
		to.ClaimedAtNin = &elemclaimedAtNin0
	}
	if o.ClaimedAtExists != nil {
		elemclaimedAtExists0 := o.ClaimedAtExists
		to.ClaimedAtExists = elemclaimedAtExists0
	}
	if o.ContributorEmailEq != nil {
		elemcontributorEmailEq0 := o.ContributorEmailEq
		to.ContributorEmailEq = elemcontributorEmailEq0
	}
	if o.ContributorEmailNe != nil {
		elemcontributorEmailNe0 := o.ContributorEmailNe
		to.ContributorEmailNe = elemcontributorEmailNe0
	}
	if o.ContributorEmailGt != nil {
		elemcontributorEmailGt0 := o.ContributorEmailGt
		to.ContributorEmailGt = elemcontributorEmailGt0
	}
	if o.ContributorEmailGte != nil {
		elemcontributorEmailGte0 := o.ContributorEmailGte
		to.ContributorEmailGte = elemcontributorEmailGte0
	}
	if o.ContributorEmailLt != nil {
		elemcontributorEmailLt0 := o.ContributorEmailLt
		to.ContributorEmailLt = elemcontributorEmailLt0
	}
	if o.ContributorEmailLte != nil {
		elemcontributorEmailLte0 := o.ContributorEmailLte
		to.ContributorEmailLte = elemcontributorEmailLte0
	}
	if o.ContributorEmailIn != nil {
		elemcontributorEmailIn0 := make([]string, 0)
		for _, ocontributorEmailIn0 := range *o.ContributorEmailIn {
			elemcontributorEmailIn1 := ocontributorEmailIn0
			elemcontributorEmailIn0 = append(elemcontributorEmailIn0, elemcontributorEmailIn1)
		}
		to.ContributorEmailIn = &elemcontributorEmailIn0
	}
	if o.ContributorEmailNin != nil {
		elemcontributorEmailNin0 := make([]string, 0)
		for _, ocontributorEmailNin0 := range *o.ContributorEmailNin {
			elemcontributorEmailNin1 := ocontributorEmailNin0
			elemcontributorEmailNin0 = append(elemcontributorEmailNin0, elemcontributorEmailNin1)
		}
		to.ContributorEmailNin = &elemcontributorEmailNin0
	}
	if o.ContributorEmailExists != nil {
		elemcontributorEmailExists0 := o.ContributorEmailExists
		to.ContributorEmailExists = elemcontributorEmailExists0
	}
	if o.ContributorEmailLike != nil {
		elemcontributorEmailLike0 := o.ContributorEmailLike
		to.ContributorEmailLike = elemcontributorEmailLike0
	}
	if o.ContributorEmailNlike != nil {
		elemcontributorEmailNlike0 := o.ContributorEmailNlike
		to.ContributorEmailNlike = elemcontributorEmailNlike0
	}
	if o.ContributorNameEq != nil {
		elemcontributorNameEq0 := o.ContributorNameEq
		to.ContributorNameEq = elemcontributorNameEq0
	}
	if o.ContributorNameNe != nil {
		elemcontributorNameNe0 := o.ContributorNameNe
		to.ContributorNameNe = elemcontributorNameNe0
	}
	if o.ContributorNameGt != nil {
		elemcontributorNameGt0 := o.ContributorNameGt
		to.ContributorNameGt = elemcontributorNameGt0
	}
	if o.ContributorNameGte != nil {
		elemcontributorNameGte0 := o.ContributorNameGte
		to.ContributorNameGte = elemcontributorNameGte0
	}
	if o.ContributorNameLt != nil {
		elemcontributorNameLt0 := o.ContributorNameLt
		to.ContributorNameLt = elemcontributorNameLt0
	}
	if o.ContributorNameLte != nil {
		elemcontributorNameLte0 := o.ContributorNameLte
		to.ContributorNameLte = elemcontributorNameLte0
	}
	if o.ContributorNameIn != nil {
		elemcontributorNameIn0 := make([]string, 0)
		for _, ocontributorNameIn0 := range *o.ContributorNameIn {
			elemcontributorNameIn1 := ocontributorNameIn0
			elemcontributorNameIn0 = append(elemcontributorNameIn0, elemcontributorNameIn1)
		}
		to.ContributorNameIn = &elemcontributorNameIn0
	}
	if o.ContributorNameNin != nil {
		elemcontributorNameNin0 := make([]string, 0)
		for _, ocontributorNameNin0 := range *o.ContributorNameNin {
			elemcontributorNameNin1 := ocontributorNameNin0
			elemcontributorNameNin0 = append(elemcontributorNameNin0, elemcontributorNameNin1)
		}
		to.ContributorNameNin = &elemcontributorNameNin0
	}
	if o.ContributorNameExists != nil {
		elemcontributorNameExists0 := o.ContributorNameExists
		to.ContributorNameExists = elemcontributorNameExists0
	}
	if o.ContributorNameLike != nil {
		elemcontributorNameLike0 := o.ContributorNameLike
		to.ContributorNameLike = elemcontributorNameLike0
	}
	if o.ContributorNameNlike != nil {
		elemcontributorNameNlike0 := o.ContributorNameNlike
		to.ContributorNameNlike = elemcontributorNameNlike0
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
	if o.DecidedAtEq != nil {
		elemdecidedAtEq0 := o.DecidedAtEq
		to.DecidedAtEq = elemdecidedAtEq0
	}
	if o.DecidedAtNe != nil {
		elemdecidedAtNe0 := o.DecidedAtNe
		to.DecidedAtNe = elemdecidedAtNe0
	}
	if o.DecidedAtGt != nil {
		elemdecidedAtGt0 := o.DecidedAtGt
		to.DecidedAtGt = elemdecidedAtGt0
	}
	if o.DecidedAtGte != nil {
		elemdecidedAtGte0 := o.DecidedAtGte
		to.DecidedAtGte = elemdecidedAtGte0
	}
	if o.DecidedAtLt != nil {
		elemdecidedAtLt0 := o.DecidedAtLt
		to.DecidedAtLt = elemdecidedAtLt0
	}
	if o.DecidedAtLte != nil {
		elemdecidedAtLte0 := o.DecidedAtLte
		to.DecidedAtLte = elemdecidedAtLte0
	}
	if o.DecidedAtIn != nil {
		elemdecidedAtIn0 := make([]time.Time, 0)
		for _, odecidedAtIn0 := range *o.DecidedAtIn {
			elemdecidedAtIn1 := odecidedAtIn0
			elemdecidedAtIn0 = append(elemdecidedAtIn0, elemdecidedAtIn1)
		}
		to.DecidedAtIn = &elemdecidedAtIn0
	}
	if o.DecidedAtNin != nil {
		elemdecidedAtNin0 := make([]time.Time, 0)
		for _, odecidedAtNin0 := range *o.DecidedAtNin {
			elemdecidedAtNin1 := odecidedAtNin0
			elemdecidedAtNin0 = append(elemdecidedAtNin0, elemdecidedAtNin1)
		}
		to.DecidedAtNin = &elemdecidedAtNin0
	}
	if o.DecidedAtExists != nil {
		elemdecidedAtExists0 := o.DecidedAtExists
		to.DecidedAtExists = elemdecidedAtExists0
	}
	if o.DecisionReasonEq != nil {
		elemdecisionReasonEq0 := o.DecisionReasonEq
		to.DecisionReasonEq = elemdecisionReasonEq0
	}
	if o.DecisionReasonNe != nil {
		elemdecisionReasonNe0 := o.DecisionReasonNe
		to.DecisionReasonNe = elemdecisionReasonNe0
	}
	if o.DecisionReasonGt != nil {
		elemdecisionReasonGt0 := o.DecisionReasonGt
		to.DecisionReasonGt = elemdecisionReasonGt0
	}
	if o.DecisionReasonGte != nil {
		elemdecisionReasonGte0 := o.DecisionReasonGte
		to.DecisionReasonGte = elemdecisionReasonGte0
	}
	if o.DecisionReasonLt != nil {
		elemdecisionReasonLt0 := o.DecisionReasonLt
		to.DecisionReasonLt = elemdecisionReasonLt0
	}
	if o.DecisionReasonLte != nil {
		elemdecisionReasonLte0 := o.DecisionReasonLte
		to.DecisionReasonLte = elemdecisionReasonLte0
	}
	if o.DecisionReasonIn != nil {
		elemdecisionReasonIn0 := make([]string, 0)
		for _, odecisionReasonIn0 := range *o.DecisionReasonIn {
			elemdecisionReasonIn1 := odecisionReasonIn0
			elemdecisionReasonIn0 = append(elemdecisionReasonIn0, elemdecisionReasonIn1)
		}
		to.DecisionReasonIn = &elemdecisionReasonIn0
	}
	if o.DecisionReasonNin != nil {
		elemdecisionReasonNin0 := make([]string, 0)
		for _, odecisionReasonNin0 := range *o.DecisionReasonNin {
			elemdecisionReasonNin1 := odecisionReasonNin0
			elemdecisionReasonNin0 = append(elemdecisionReasonNin0, elemdecisionReasonNin1)
		}
		to.DecisionReasonNin = &elemdecisionReasonNin0
	}
	if o.DecisionReasonExists != nil {
		elemdecisionReasonExists0 := o.DecisionReasonExists
		to.DecisionReasonExists = elemdecisionReasonExists0
	}
	if o.DecisionReasonLike != nil {
		elemdecisionReasonLike0 := o.DecisionReasonLike
		to.DecisionReasonLike = elemdecisionReasonLike0
	}
	if o.DecisionReasonNlike != nil {
		elemdecisionReasonNlike0 := o.DecisionReasonNlike
		to.DecisionReasonNlike = elemdecisionReasonNlike0
	}
	if o.MessageEq != nil {
		elemmessageEq0 := o.MessageEq
		to.MessageEq = elemmessageEq0
	}
	if o.MessageNe != nil {
		elemmessageNe0 := o.MessageNe
		to.MessageNe = elemmessageNe0
	}
	if o.MessageGt != nil {
		elemmessageGt0 := o.MessageGt
		to.MessageGt = elemmessageGt0
	}
	if o.MessageGte != nil {
		elemmessageGte0 := o.MessageGte
		to.MessageGte = elemmessageGte0
	}
	if o.MessageLt != nil {
		elemmessageLt0 := o.MessageLt
		to.MessageLt = elemmessageLt0
	}
	if o.MessageLte != nil {
		elemmessageLte0 := o.MessageLte
		to.MessageLte = elemmessageLte0
	}
	if o.MessageIn != nil {
		elemmessageIn0 := make([]string, 0)
		for _, omessageIn0 := range *o.MessageIn {
			elemmessageIn1 := omessageIn0
			elemmessageIn0 = append(elemmessageIn0, elemmessageIn1)
		}
		to.MessageIn = &elemmessageIn0
	}
	if o.MessageNin != nil {
		elemmessageNin0 := make([]string, 0)
		for _, omessageNin0 := range *o.MessageNin {
			elemmessageNin1 := omessageNin0
			elemmessageNin0 = append(elemmessageNin0, elemmessageNin1)
		}
		to.MessageNin = &elemmessageNin0
	}
	if o.MessageExists != nil {
		elemmessageExists0 := o.MessageExists
		to.MessageExists = elemmessageExists0
	}
	if o.MessageLike != nil {
		elemmessageLike0 := o.MessageLike
		to.MessageLike = elemmessageLike0
	}
	if o.MessageNlike != nil {
		elemmessageNlike0 := o.MessageNlike
		to.MessageNlike = elemmessageNlike0
	}
	if o.MethodDisplayNameEq != nil {
		elemmethodDisplayNameEq0 := o.MethodDisplayNameEq
		to.MethodDisplayNameEq = elemmethodDisplayNameEq0
	}
	if o.MethodDisplayNameNe != nil {
		elemmethodDisplayNameNe0 := o.MethodDisplayNameNe
		to.MethodDisplayNameNe = elemmethodDisplayNameNe0
	}
	if o.MethodDisplayNameGt != nil {
		elemmethodDisplayNameGt0 := o.MethodDisplayNameGt
		to.MethodDisplayNameGt = elemmethodDisplayNameGt0
	}
	if o.MethodDisplayNameGte != nil {
		elemmethodDisplayNameGte0 := o.MethodDisplayNameGte
		to.MethodDisplayNameGte = elemmethodDisplayNameGte0
	}
	if o.MethodDisplayNameLt != nil {
		elemmethodDisplayNameLt0 := o.MethodDisplayNameLt
		to.MethodDisplayNameLt = elemmethodDisplayNameLt0
	}
	if o.MethodDisplayNameLte != nil {
		elemmethodDisplayNameLte0 := o.MethodDisplayNameLte
		to.MethodDisplayNameLte = elemmethodDisplayNameLte0
	}
	if o.MethodDisplayNameIn != nil {
		elemmethodDisplayNameIn0 := make([]string, 0)
		for _, omethodDisplayNameIn0 := range *o.MethodDisplayNameIn {
			elemmethodDisplayNameIn1 := omethodDisplayNameIn0
			elemmethodDisplayNameIn0 = append(elemmethodDisplayNameIn0, elemmethodDisplayNameIn1)
		}
		to.MethodDisplayNameIn = &elemmethodDisplayNameIn0
	}
	if o.MethodDisplayNameNin != nil {
		elemmethodDisplayNameNin0 := make([]string, 0)
		for _, omethodDisplayNameNin0 := range *o.MethodDisplayNameNin {
			elemmethodDisplayNameNin1 := omethodDisplayNameNin0
			elemmethodDisplayNameNin0 = append(elemmethodDisplayNameNin0, elemmethodDisplayNameNin1)
		}
		to.MethodDisplayNameNin = &elemmethodDisplayNameNin0
	}
	if o.MethodDisplayNameExists != nil {
		elemmethodDisplayNameExists0 := o.MethodDisplayNameExists
		to.MethodDisplayNameExists = elemmethodDisplayNameExists0
	}
	if o.MethodDisplayNameLike != nil {
		elemmethodDisplayNameLike0 := o.MethodDisplayNameLike
		to.MethodDisplayNameLike = elemmethodDisplayNameLike0
	}
	if o.MethodDisplayNameNlike != nil {
		elemmethodDisplayNameNlike0 := o.MethodDisplayNameNlike
		to.MethodDisplayNameNlike = elemmethodDisplayNameNlike0
	}
	if o.MethodTypeEq != nil {
		elemmethodTypeEq0 := o.MethodTypeEq
		to.MethodTypeEq = elemmethodTypeEq0
	}
	if o.MethodTypeNe != nil {
		elemmethodTypeNe0 := o.MethodTypeNe
		to.MethodTypeNe = elemmethodTypeNe0
	}
	if o.MethodTypeGt != nil {
		elemmethodTypeGt0 := o.MethodTypeGt
		to.MethodTypeGt = elemmethodTypeGt0
	}
	if o.MethodTypeGte != nil {
		elemmethodTypeGte0 := o.MethodTypeGte
		to.MethodTypeGte = elemmethodTypeGte0
	}
	if o.MethodTypeLt != nil {
		elemmethodTypeLt0 := o.MethodTypeLt
		to.MethodTypeLt = elemmethodTypeLt0
	}
	if o.MethodTypeLte != nil {
		elemmethodTypeLte0 := o.MethodTypeLte
		to.MethodTypeLte = elemmethodTypeLte0
	}
	if o.MethodTypeIn != nil {
		elemmethodTypeIn0 := make([]enum_payment_method_type.Value, 0)
		for _, omethodTypeIn0 := range *o.MethodTypeIn {
			elemmethodTypeIn1 := omethodTypeIn0
			elemmethodTypeIn0 = append(elemmethodTypeIn0, elemmethodTypeIn1)
		}
		to.MethodTypeIn = &elemmethodTypeIn0
	}
	if o.MethodTypeNin != nil {
		elemmethodTypeNin0 := make([]enum_payment_method_type.Value, 0)
		for _, omethodTypeNin0 := range *o.MethodTypeNin {
			elemmethodTypeNin1 := omethodTypeNin0
			elemmethodTypeNin0 = append(elemmethodTypeNin0, elemmethodTypeNin1)
		}
		to.MethodTypeNin = &elemmethodTypeNin0
	}
	if o.MethodTypeExists != nil {
		elemmethodTypeExists0 := o.MethodTypeExists
		to.MethodTypeExists = elemmethodTypeExists0
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
	if o.PaymentMethodIdEq != nil {
		elempaymentMethodIdEq0, err := primitive.ObjectIDFromHex(*o.PaymentMethodIdEq)
		if err != nil {
			return to, errors.Join(errors.New("invalid o.PaymentMethodIdEq"), err)
		}
		to.PaymentMethodIdEq = &elempaymentMethodIdEq0
	}
	if o.PaymentMethodIdIn != nil {
		elempaymentMethodIdIn0 := make([]primitive.ObjectID, 0)
		for _, opaymentMethodIdIn0 := range *o.PaymentMethodIdIn {
			elempaymentMethodIdIn1, err := primitive.ObjectIDFromHex(opaymentMethodIdIn0)
			if err != nil {
				return to, errors.Join(errors.New("invalid opaymentMethodIdIn0"), err)
			}
			elempaymentMethodIdIn0 = append(elempaymentMethodIdIn0, elempaymentMethodIdIn1)
		}
		to.PaymentMethodIdIn = &elempaymentMethodIdIn0
	}
	if o.PaymentMethodIdNin != nil {
		elempaymentMethodIdNin0 := make([]primitive.ObjectID, 0)
		for _, opaymentMethodIdNin0 := range *o.PaymentMethodIdNin {
			elempaymentMethodIdNin1, err := primitive.ObjectIDFromHex(opaymentMethodIdNin0)
			if err != nil {
				return to, errors.Join(errors.New("invalid opaymentMethodIdNin0"), err)
			}
			elempaymentMethodIdNin0 = append(elempaymentMethodIdNin0, elempaymentMethodIdNin1)
		}
		to.PaymentMethodIdNin = &elempaymentMethodIdNin0
	}
	if o.PaymentMethodIdExists != nil {
		elempaymentMethodIdExists0 := o.PaymentMethodIdExists
		to.PaymentMethodIdExists = elempaymentMethodIdExists0
	}
	if o.ReferenceCodeEq != nil {
		elemreferenceCodeEq0 := o.ReferenceCodeEq
		to.ReferenceCodeEq = elemreferenceCodeEq0
	}
	if o.ReferenceCodeNe != nil {
		elemreferenceCodeNe0 := o.ReferenceCodeNe
		to.ReferenceCodeNe = elemreferenceCodeNe0
	}
	if o.ReferenceCodeGt != nil {
		elemreferenceCodeGt0 := o.ReferenceCodeGt
		to.ReferenceCodeGt = elemreferenceCodeGt0
	}
	if o.ReferenceCodeGte != nil {
		elemreferenceCodeGte0 := o.ReferenceCodeGte
		to.ReferenceCodeGte = elemreferenceCodeGte0
	}
	if o.ReferenceCodeLt != nil {
		elemreferenceCodeLt0 := o.ReferenceCodeLt
		to.ReferenceCodeLt = elemreferenceCodeLt0
	}
	if o.ReferenceCodeLte != nil {
		elemreferenceCodeLte0 := o.ReferenceCodeLte
		to.ReferenceCodeLte = elemreferenceCodeLte0
	}
	if o.ReferenceCodeIn != nil {
		elemreferenceCodeIn0 := make([]string, 0)
		for _, oreferenceCodeIn0 := range *o.ReferenceCodeIn {
			elemreferenceCodeIn1 := oreferenceCodeIn0
			elemreferenceCodeIn0 = append(elemreferenceCodeIn0, elemreferenceCodeIn1)
		}
		to.ReferenceCodeIn = &elemreferenceCodeIn0
	}
	if o.ReferenceCodeNin != nil {
		elemreferenceCodeNin0 := make([]string, 0)
		for _, oreferenceCodeNin0 := range *o.ReferenceCodeNin {
			elemreferenceCodeNin1 := oreferenceCodeNin0
			elemreferenceCodeNin0 = append(elemreferenceCodeNin0, elemreferenceCodeNin1)
		}
		to.ReferenceCodeNin = &elemreferenceCodeNin0
	}
	if o.ReferenceCodeExists != nil {
		elemreferenceCodeExists0 := o.ReferenceCodeExists
		to.ReferenceCodeExists = elemreferenceCodeExists0
	}
	if o.ReferenceCodeLike != nil {
		elemreferenceCodeLike0 := o.ReferenceCodeLike
		to.ReferenceCodeLike = elemreferenceCodeLike0
	}
	if o.ReferenceCodeNlike != nil {
		elemreferenceCodeNlike0 := o.ReferenceCodeNlike
		to.ReferenceCodeNlike = elemreferenceCodeNlike0
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
	if o.StatusEq != nil {
		elemstatusEq0 := o.StatusEq
		to.StatusEq = elemstatusEq0
	}
	if o.StatusNe != nil {
		elemstatusNe0 := o.StatusNe
		to.StatusNe = elemstatusNe0
	}
	if o.StatusGt != nil {
		elemstatusGt0 := o.StatusGt
		to.StatusGt = elemstatusGt0
	}
	if o.StatusGte != nil {
		elemstatusGte0 := o.StatusGte
		to.StatusGte = elemstatusGte0
	}
	if o.StatusLt != nil {
		elemstatusLt0 := o.StatusLt
		to.StatusLt = elemstatusLt0
	}
	if o.StatusLte != nil {
		elemstatusLte0 := o.StatusLte
		to.StatusLte = elemstatusLte0
	}
	if o.StatusIn != nil {
		elemstatusIn0 := make([]enum_cart_status.Value, 0)
		for _, ostatusIn0 := range *o.StatusIn {
			elemstatusIn1 := ostatusIn0
			elemstatusIn0 = append(elemstatusIn0, elemstatusIn1)
		}
		to.StatusIn = &elemstatusIn0
	}
	if o.StatusNin != nil {
		elemstatusNin0 := make([]enum_cart_status.Value, 0)
		for _, ostatusNin0 := range *o.StatusNin {
			elemstatusNin1 := ostatusNin0
			elemstatusNin0 = append(elemstatusNin0, elemstatusNin1)
		}
		to.StatusNin = &elemstatusNin0
	}
	if o.StatusExists != nil {
		elemstatusExists0 := o.StatusExists
		to.StatusExists = elemstatusExists0
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
	ContributorEmail int8
	CreatedAt        int8
	OwnerId          int8
	ReferenceCode    int8
	RegistryId       int8
	UpdatedAt        int8
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	to.ContributorEmail = s.ContributorEmail
	to.CreatedAt = s.CreatedAt
	to.OwnerId = s.OwnerId
	to.ReferenceCode = s.ReferenceCode
	to.RegistryId = s.RegistryId
	to.UpdatedAt = s.UpdatedAt
	return to
}
