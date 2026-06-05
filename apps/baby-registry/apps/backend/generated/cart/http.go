package cart

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_cart_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_payment_method_type"
	"time"
)

type HTTPRecord struct {
	Id                 *string                         `json:"id,omitempty"`
	AmountCents        *int                            `json:"amountCents,omitempty"`
	ClaimedAt          *time.Time                      `json:"claimedAt,omitempty"`
	ContributorEmail   *string                         `json:"contributorEmail,omitempty"`
	ContributorName    *string                         `json:"contributorName,omitempty"`
	Created            *actor_trace.HTTPRecord         `json:"created,omitempty"`
	Currency           *string                         `json:"currency,omitempty"`
	DecidedAt          *time.Time                      `json:"decidedAt,omitempty"`
	DecisionReason     *string                         `json:"decisionReason,omitempty"`
	Message            *string                         `json:"message,omitempty"`
	MethodDisplayName  *string                         `json:"methodDisplayName,omitempty"`
	MethodType         *enum_payment_method_type.Value `json:"methodType,omitempty"`
	OwnerId            *string                         `json:"ownerId,omitempty"`
	PaymentMethodId    *string                         `json:"paymentMethodId,omitempty"`
	ReferenceCode      *string                         `json:"referenceCode,omitempty"`
	RegistryId         *string                         `json:"registryId,omitempty"`
	Status             *enum_cart_status.Value         `json:"status,omitempty"`
	Updated            *actor_trace.HTTPRecord         `json:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.HTTPRecord         `json:"updatedByOwnerUser,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id
		m.Id = *elemid0
	}
	if r.AmountCents != nil {
		elemamountCents0 := r.AmountCents
		m.AmountCents = *elemamountCents0
	}
	if r.ClaimedAt != nil {
		elemclaimedAt0 := r.ClaimedAt
		m.ClaimedAt = *elemclaimedAt0
	}
	if r.ContributorEmail != nil {
		elemcontributorEmail0 := r.ContributorEmail
		m.ContributorEmail = *elemcontributorEmail0
	}
	if r.ContributorName != nil {
		elemcontributorName0 := r.ContributorName
		m.ContributorName = *elemcontributorName0
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
	if r.DecidedAt != nil {
		elemdecidedAt0 := r.DecidedAt
		m.DecidedAt = *elemdecidedAt0
	}
	if r.DecisionReason != nil {
		elemdecisionReason0 := r.DecisionReason
		m.DecisionReason = *elemdecisionReason0
	}
	if r.Message != nil {
		elemmessage0 := r.Message
		m.Message = *elemmessage0
	}
	if r.MethodDisplayName != nil {
		elemmethodDisplayName0 := r.MethodDisplayName
		m.MethodDisplayName = *elemmethodDisplayName0
	}
	if r.MethodType != nil {
		elemmethodType0 := r.MethodType
		m.MethodType = *elemmethodType0
	}
	if r.OwnerId != nil {
		elemownerId0 := r.OwnerId
		m.OwnerId = *elemownerId0
	}
	if r.PaymentMethodId != nil {
		elempaymentMethodId0 := r.PaymentMethodId
		m.PaymentMethodId = *elempaymentMethodId0
	}
	if r.ReferenceCode != nil {
		elemreferenceCode0 := r.ReferenceCode
		m.ReferenceCode = *elemreferenceCode0
	}
	if r.RegistryId != nil {
		elemregistryId0 := r.RegistryId
		m.RegistryId = *elemregistryId0
	}
	if r.Status != nil {
		elemstatus0 := r.Status
		m.Status = *elemstatus0
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
	if r.AmountCents != nil {
		p.AmountCents = true
	}
	if r.ClaimedAt != nil {
		p.ClaimedAt = true
	}
	if r.ContributorEmail != nil {
		p.ContributorEmail = true
	}
	if r.ContributorName != nil {
		p.ContributorName = true
	}
	if r.Created != nil {
		p.Created = true
		p.CreatedFields = actor_trace.NewProjection(true)
	}
	if r.Currency != nil {
		p.Currency = true
	}
	if r.DecidedAt != nil {
		p.DecidedAt = true
	}
	if r.DecisionReason != nil {
		p.DecisionReason = true
	}
	if r.Message != nil {
		p.Message = true
	}
	if r.MethodDisplayName != nil {
		p.MethodDisplayName = true
	}
	if r.MethodType != nil {
		p.MethodType = true
	}
	if r.OwnerId != nil {
		p.OwnerId = true
	}
	if r.PaymentMethodId != nil {
		p.PaymentMethodId = true
	}
	if r.ReferenceCode != nil {
		p.ReferenceCode = true
	}
	if r.RegistryId != nil {
		p.RegistryId = true
	}
	if r.Status != nil {
		p.Status = true
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
type HTTPSelectByReferenceUniqueQuery struct {
	ReferenceCode string `json:"referenceCode"`
}

type HTTPWhereClause struct {
	// id (Ref<Cart>) search options
	IdEq     *string   `json:"idEq,omitempty"`
	IdIn     *[]string `json:"idIn,omitempty"`
	IdNin    *[]string `json:"idNin,omitempty"`
	IdExists *bool     `json:"idExists,omitempty"`
	// amountCents (int) search options
	AmountCentsEq     *int   `json:"amountCentsEq,omitempty"`
	AmountCentsNe     *int   `json:"amountCentsNe,omitempty"`
	AmountCentsGt     *int   `json:"amountCentsGt,omitempty"`
	AmountCentsGte    *int   `json:"amountCentsGte,omitempty"`
	AmountCentsLt     *int   `json:"amountCentsLt,omitempty"`
	AmountCentsLte    *int   `json:"amountCentsLte,omitempty"`
	AmountCentsIn     *[]int `json:"amountCentsIn,omitempty"`
	AmountCentsNin    *[]int `json:"amountCentsNin,omitempty"`
	AmountCentsExists *bool  `json:"amountCentsExists,omitempty"`
	// claimedAt (timestamp) search options
	ClaimedAtEq     *time.Time   `json:"claimedAtEq,omitempty"`
	ClaimedAtNe     *time.Time   `json:"claimedAtNe,omitempty"`
	ClaimedAtGt     *time.Time   `json:"claimedAtGt,omitempty"`
	ClaimedAtGte    *time.Time   `json:"claimedAtGte,omitempty"`
	ClaimedAtLt     *time.Time   `json:"claimedAtLt,omitempty"`
	ClaimedAtLte    *time.Time   `json:"claimedAtLte,omitempty"`
	ClaimedAtIn     *[]time.Time `json:"claimedAtIn,omitempty"`
	ClaimedAtNin    *[]time.Time `json:"claimedAtNin,omitempty"`
	ClaimedAtExists *bool        `json:"claimedAtExists,omitempty"`
	// contributorEmail (string) search options
	ContributorEmailEq     *string   `json:"contributorEmailEq,omitempty"`
	ContributorEmailNe     *string   `json:"contributorEmailNe,omitempty"`
	ContributorEmailGt     *string   `json:"contributorEmailGt,omitempty"`
	ContributorEmailGte    *string   `json:"contributorEmailGte,omitempty"`
	ContributorEmailLt     *string   `json:"contributorEmailLt,omitempty"`
	ContributorEmailLte    *string   `json:"contributorEmailLte,omitempty"`
	ContributorEmailIn     *[]string `json:"contributorEmailIn,omitempty"`
	ContributorEmailNin    *[]string `json:"contributorEmailNin,omitempty"`
	ContributorEmailExists *bool     `json:"contributorEmailExists,omitempty"`
	ContributorEmailLike   *string   `json:"contributorEmailLike,omitempty"`
	ContributorEmailNlike  *string   `json:"contributorEmailNlike,omitempty"`
	// contributorName (string) search options
	ContributorNameEq     *string   `json:"contributorNameEq,omitempty"`
	ContributorNameNe     *string   `json:"contributorNameNe,omitempty"`
	ContributorNameGt     *string   `json:"contributorNameGt,omitempty"`
	ContributorNameGte    *string   `json:"contributorNameGte,omitempty"`
	ContributorNameLt     *string   `json:"contributorNameLt,omitempty"`
	ContributorNameLte    *string   `json:"contributorNameLte,omitempty"`
	ContributorNameIn     *[]string `json:"contributorNameIn,omitempty"`
	ContributorNameNin    *[]string `json:"contributorNameNin,omitempty"`
	ContributorNameExists *bool     `json:"contributorNameExists,omitempty"`
	ContributorNameLike   *string   `json:"contributorNameLike,omitempty"`
	ContributorNameNlike  *string   `json:"contributorNameNlike,omitempty"`
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
	// decidedAt (timestamp) search options
	DecidedAtEq     *time.Time   `json:"decidedAtEq,omitempty"`
	DecidedAtNe     *time.Time   `json:"decidedAtNe,omitempty"`
	DecidedAtGt     *time.Time   `json:"decidedAtGt,omitempty"`
	DecidedAtGte    *time.Time   `json:"decidedAtGte,omitempty"`
	DecidedAtLt     *time.Time   `json:"decidedAtLt,omitempty"`
	DecidedAtLte    *time.Time   `json:"decidedAtLte,omitempty"`
	DecidedAtIn     *[]time.Time `json:"decidedAtIn,omitempty"`
	DecidedAtNin    *[]time.Time `json:"decidedAtNin,omitempty"`
	DecidedAtExists *bool        `json:"decidedAtExists,omitempty"`
	// decisionReason (string) search options
	DecisionReasonEq     *string   `json:"decisionReasonEq,omitempty"`
	DecisionReasonNe     *string   `json:"decisionReasonNe,omitempty"`
	DecisionReasonGt     *string   `json:"decisionReasonGt,omitempty"`
	DecisionReasonGte    *string   `json:"decisionReasonGte,omitempty"`
	DecisionReasonLt     *string   `json:"decisionReasonLt,omitempty"`
	DecisionReasonLte    *string   `json:"decisionReasonLte,omitempty"`
	DecisionReasonIn     *[]string `json:"decisionReasonIn,omitempty"`
	DecisionReasonNin    *[]string `json:"decisionReasonNin,omitempty"`
	DecisionReasonExists *bool     `json:"decisionReasonExists,omitempty"`
	DecisionReasonLike   *string   `json:"decisionReasonLike,omitempty"`
	DecisionReasonNlike  *string   `json:"decisionReasonNlike,omitempty"`
	// message (string) search options
	MessageEq     *string   `json:"messageEq,omitempty"`
	MessageNe     *string   `json:"messageNe,omitempty"`
	MessageGt     *string   `json:"messageGt,omitempty"`
	MessageGte    *string   `json:"messageGte,omitempty"`
	MessageLt     *string   `json:"messageLt,omitempty"`
	MessageLte    *string   `json:"messageLte,omitempty"`
	MessageIn     *[]string `json:"messageIn,omitempty"`
	MessageNin    *[]string `json:"messageNin,omitempty"`
	MessageExists *bool     `json:"messageExists,omitempty"`
	MessageLike   *string   `json:"messageLike,omitempty"`
	MessageNlike  *string   `json:"messageNlike,omitempty"`
	// methodDisplayName (string) search options
	MethodDisplayNameEq     *string   `json:"methodDisplayNameEq,omitempty"`
	MethodDisplayNameNe     *string   `json:"methodDisplayNameNe,omitempty"`
	MethodDisplayNameGt     *string   `json:"methodDisplayNameGt,omitempty"`
	MethodDisplayNameGte    *string   `json:"methodDisplayNameGte,omitempty"`
	MethodDisplayNameLt     *string   `json:"methodDisplayNameLt,omitempty"`
	MethodDisplayNameLte    *string   `json:"methodDisplayNameLte,omitempty"`
	MethodDisplayNameIn     *[]string `json:"methodDisplayNameIn,omitempty"`
	MethodDisplayNameNin    *[]string `json:"methodDisplayNameNin,omitempty"`
	MethodDisplayNameExists *bool     `json:"methodDisplayNameExists,omitempty"`
	MethodDisplayNameLike   *string   `json:"methodDisplayNameLike,omitempty"`
	MethodDisplayNameNlike  *string   `json:"methodDisplayNameNlike,omitempty"`
	// methodType (PaymentMethodType) search options
	MethodTypeEq     *enum_payment_method_type.Value   `json:"methodTypeEq,omitempty"`
	MethodTypeNe     *enum_payment_method_type.Value   `json:"methodTypeNe,omitempty"`
	MethodTypeGt     *enum_payment_method_type.Value   `json:"methodTypeGt,omitempty"`
	MethodTypeGte    *enum_payment_method_type.Value   `json:"methodTypeGte,omitempty"`
	MethodTypeLt     *enum_payment_method_type.Value   `json:"methodTypeLt,omitempty"`
	MethodTypeLte    *enum_payment_method_type.Value   `json:"methodTypeLte,omitempty"`
	MethodTypeIn     *[]enum_payment_method_type.Value `json:"methodTypeIn,omitempty"`
	MethodTypeNin    *[]enum_payment_method_type.Value `json:"methodTypeNin,omitempty"`
	MethodTypeExists *bool                             `json:"methodTypeExists,omitempty"`
	// ownerId (Ref<OwnerUser>) search options
	OwnerIdEq     *string   `json:"ownerIdEq,omitempty"`
	OwnerIdIn     *[]string `json:"ownerIdIn,omitempty"`
	OwnerIdNin    *[]string `json:"ownerIdNin,omitempty"`
	OwnerIdExists *bool     `json:"ownerIdExists,omitempty"`
	// paymentMethodId (Ref<RegistryPaymentMethod>) search options
	PaymentMethodIdEq     *string   `json:"paymentMethodIdEq,omitempty"`
	PaymentMethodIdIn     *[]string `json:"paymentMethodIdIn,omitempty"`
	PaymentMethodIdNin    *[]string `json:"paymentMethodIdNin,omitempty"`
	PaymentMethodIdExists *bool     `json:"paymentMethodIdExists,omitempty"`
	// referenceCode (string) search options
	ReferenceCodeEq     *string   `json:"referenceCodeEq,omitempty"`
	ReferenceCodeNe     *string   `json:"referenceCodeNe,omitempty"`
	ReferenceCodeGt     *string   `json:"referenceCodeGt,omitempty"`
	ReferenceCodeGte    *string   `json:"referenceCodeGte,omitempty"`
	ReferenceCodeLt     *string   `json:"referenceCodeLt,omitempty"`
	ReferenceCodeLte    *string   `json:"referenceCodeLte,omitempty"`
	ReferenceCodeIn     *[]string `json:"referenceCodeIn,omitempty"`
	ReferenceCodeNin    *[]string `json:"referenceCodeNin,omitempty"`
	ReferenceCodeExists *bool     `json:"referenceCodeExists,omitempty"`
	ReferenceCodeLike   *string   `json:"referenceCodeLike,omitempty"`
	ReferenceCodeNlike  *string   `json:"referenceCodeNlike,omitempty"`
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *string   `json:"registryIdEq,omitempty"`
	RegistryIdIn     *[]string `json:"registryIdIn,omitempty"`
	RegistryIdNin    *[]string `json:"registryIdNin,omitempty"`
	RegistryIdExists *bool     `json:"registryIdExists,omitempty"`
	// status (CartStatus) search options
	StatusEq     *enum_cart_status.Value   `json:"statusEq,omitempty"`
	StatusNe     *enum_cart_status.Value   `json:"statusNe,omitempty"`
	StatusGt     *enum_cart_status.Value   `json:"statusGt,omitempty"`
	StatusGte    *enum_cart_status.Value   `json:"statusGte,omitempty"`
	StatusLt     *enum_cart_status.Value   `json:"statusLt,omitempty"`
	StatusLte    *enum_cart_status.Value   `json:"statusLte,omitempty"`
	StatusIn     *[]enum_cart_status.Value `json:"statusIn,omitempty"`
	StatusNin    *[]enum_cart_status.Value `json:"statusNin,omitempty"`
	StatusExists *bool                     `json:"statusExists,omitempty"`
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
func (o HTTPSelectByReferenceUniqueQuery) ToSelectByReferenceUniqueQuery() (SelectByReferenceUniqueQuery, error) {
	to := SelectByReferenceUniqueQuery{}
	elemreferenceCode0 := o.ReferenceCode
	to.ReferenceCode = elemreferenceCode0
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
	if o.PaymentMethodIdEq != nil {
		elempaymentMethodIdEq0 := o.PaymentMethodIdEq
		to.PaymentMethodIdEq = elempaymentMethodIdEq0
	}
	if o.PaymentMethodIdIn != nil {
		elempaymentMethodIdIn0 := make([]string, 0)
		for _, opaymentMethodIdIn0 := range *o.PaymentMethodIdIn {
			elempaymentMethodIdIn1 := opaymentMethodIdIn0
			elempaymentMethodIdIn0 = append(elempaymentMethodIdIn0, elempaymentMethodIdIn1)
		}
		to.PaymentMethodIdIn = &elempaymentMethodIdIn0
	}
	if o.PaymentMethodIdNin != nil {
		elempaymentMethodIdNin0 := make([]string, 0)
		for _, opaymentMethodIdNin0 := range *o.PaymentMethodIdNin {
			elempaymentMethodIdNin1 := opaymentMethodIdNin0
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
	ContributorEmail *int8 `json:"contributorEmail,omitempty"`
	CreatedAt        *int8 `json:"createdAt,omitempty"`
	OwnerId          *int8 `json:"ownerId,omitempty"`
	ReferenceCode    *int8 `json:"referenceCode,omitempty"`
	RegistryId       *int8 `json:"registryId,omitempty"`
	UpdatedAt        *int8 `json:"updatedAt,omitempty"`
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	if s.ContributorEmail != nil {
		to.ContributorEmail = *s.ContributorEmail
	}
	if s.CreatedAt != nil {
		to.CreatedAt = *s.CreatedAt
	}
	if s.OwnerId != nil {
		to.OwnerId = *s.OwnerId
	}
	if s.ReferenceCode != nil {
		to.ReferenceCode = *s.ReferenceCode
	}
	if s.RegistryId != nil {
		to.RegistryId = *s.RegistryId
	}
	if s.UpdatedAt != nil {
		to.UpdatedAt = *s.UpdatedAt
	}
	return to
}
