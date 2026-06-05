package registry_payment_method

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_payment_method_type"
)

type HTTPRecord struct {
	Id                 *string                         `json:"id,omitempty"`
	BankAccountName    *string                         `json:"bankAccountName,omitempty"`
	BankAccountNumber  *string                         `json:"bankAccountNumber,omitempty"`
	BankIban           *string                         `json:"bankIban,omitempty"`
	BankName           *string                         `json:"bankName,omitempty"`
	BankRoutingNumber  *string                         `json:"bankRoutingNumber,omitempty"`
	BankSwift          *string                         `json:"bankSwift,omitempty"`
	Created            *actor_trace.HTTPRecord         `json:"created,omitempty"`
	DisplayName        *string                         `json:"displayName,omitempty"`
	Enabled            *bool                           `json:"enabled,omitempty"`
	Instructions       *string                         `json:"instructions,omitempty"`
	OwnerId            *string                         `json:"ownerId,omitempty"`
	PaymentUrl         *string                         `json:"paymentUrl,omitempty"`
	Position           *int                            `json:"position,omitempty"`
	RecipientEmail     *string                         `json:"recipientEmail,omitempty"`
	RecipientPhone     *string                         `json:"recipientPhone,omitempty"`
	RegistryId         *string                         `json:"registryId,omitempty"`
	Type               *enum_payment_method_type.Value `json:"type,omitempty"`
	Updated            *actor_trace.HTTPRecord         `json:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.HTTPRecord         `json:"updatedByOwnerUser,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id
		m.Id = *elemid0
	}
	if r.BankAccountName != nil {
		elembankAccountName0 := r.BankAccountName
		m.BankAccountName = *elembankAccountName0
	}
	if r.BankAccountNumber != nil {
		elembankAccountNumber0 := r.BankAccountNumber
		m.BankAccountNumber = *elembankAccountNumber0
	}
	if r.BankIban != nil {
		elembankIban0 := r.BankIban
		m.BankIban = *elembankIban0
	}
	if r.BankName != nil {
		elembankName0 := r.BankName
		m.BankName = *elembankName0
	}
	if r.BankRoutingNumber != nil {
		elembankRoutingNumber0 := r.BankRoutingNumber
		m.BankRoutingNumber = *elembankRoutingNumber0
	}
	if r.BankSwift != nil {
		elembankSwift0 := r.BankSwift
		m.BankSwift = *elembankSwift0
	}
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.DisplayName != nil {
		elemdisplayName0 := r.DisplayName
		m.DisplayName = *elemdisplayName0
	}
	if r.Enabled != nil {
		elemenabled0 := r.Enabled
		m.Enabled = *elemenabled0
	}
	if r.Instructions != nil {
		eleminstructions0 := r.Instructions
		m.Instructions = *eleminstructions0
	}
	if r.OwnerId != nil {
		elemownerId0 := r.OwnerId
		m.OwnerId = *elemownerId0
	}
	if r.PaymentUrl != nil {
		elempaymentUrl0 := r.PaymentUrl
		m.PaymentUrl = *elempaymentUrl0
	}
	if r.Position != nil {
		elemposition0 := r.Position
		m.Position = *elemposition0
	}
	if r.RecipientEmail != nil {
		elemrecipientEmail0 := r.RecipientEmail
		m.RecipientEmail = *elemrecipientEmail0
	}
	if r.RecipientPhone != nil {
		elemrecipientPhone0 := r.RecipientPhone
		m.RecipientPhone = *elemrecipientPhone0
	}
	if r.RegistryId != nil {
		elemregistryId0 := r.RegistryId
		m.RegistryId = *elemregistryId0
	}
	if r.Type != nil {
		elemtype0 := r.Type
		m.Type = *elemtype0
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
	if r.BankAccountName != nil {
		p.BankAccountName = true
	}
	if r.BankAccountNumber != nil {
		p.BankAccountNumber = true
	}
	if r.BankIban != nil {
		p.BankIban = true
	}
	if r.BankName != nil {
		p.BankName = true
	}
	if r.BankRoutingNumber != nil {
		p.BankRoutingNumber = true
	}
	if r.BankSwift != nil {
		p.BankSwift = true
	}
	if r.Created != nil {
		p.Created = true
		p.CreatedFields = actor_trace.NewProjection(true)
	}
	if r.DisplayName != nil {
		p.DisplayName = true
	}
	if r.Enabled != nil {
		p.Enabled = true
	}
	if r.Instructions != nil {
		p.Instructions = true
	}
	if r.OwnerId != nil {
		p.OwnerId = true
	}
	if r.PaymentUrl != nil {
		p.PaymentUrl = true
	}
	if r.Position != nil {
		p.Position = true
	}
	if r.RecipientEmail != nil {
		p.RecipientEmail = true
	}
	if r.RecipientPhone != nil {
		p.RecipientPhone = true
	}
	if r.RegistryId != nil {
		p.RegistryId = true
	}
	if r.Type != nil {
		p.Type = true
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
	// id (Ref<RegistryPaymentMethod>) search options
	IdEq     *string   `json:"idEq,omitempty"`
	IdIn     *[]string `json:"idIn,omitempty"`
	IdNin    *[]string `json:"idNin,omitempty"`
	IdExists *bool     `json:"idExists,omitempty"`
	// bankAccountName (string) search options
	BankAccountNameEq     *string   `json:"bankAccountNameEq,omitempty"`
	BankAccountNameNe     *string   `json:"bankAccountNameNe,omitempty"`
	BankAccountNameGt     *string   `json:"bankAccountNameGt,omitempty"`
	BankAccountNameGte    *string   `json:"bankAccountNameGte,omitempty"`
	BankAccountNameLt     *string   `json:"bankAccountNameLt,omitempty"`
	BankAccountNameLte    *string   `json:"bankAccountNameLte,omitempty"`
	BankAccountNameIn     *[]string `json:"bankAccountNameIn,omitempty"`
	BankAccountNameNin    *[]string `json:"bankAccountNameNin,omitempty"`
	BankAccountNameExists *bool     `json:"bankAccountNameExists,omitempty"`
	BankAccountNameLike   *string   `json:"bankAccountNameLike,omitempty"`
	BankAccountNameNlike  *string   `json:"bankAccountNameNlike,omitempty"`
	// bankAccountNumber (string) search options
	BankAccountNumberEq     *string   `json:"bankAccountNumberEq,omitempty"`
	BankAccountNumberNe     *string   `json:"bankAccountNumberNe,omitempty"`
	BankAccountNumberGt     *string   `json:"bankAccountNumberGt,omitempty"`
	BankAccountNumberGte    *string   `json:"bankAccountNumberGte,omitempty"`
	BankAccountNumberLt     *string   `json:"bankAccountNumberLt,omitempty"`
	BankAccountNumberLte    *string   `json:"bankAccountNumberLte,omitempty"`
	BankAccountNumberIn     *[]string `json:"bankAccountNumberIn,omitempty"`
	BankAccountNumberNin    *[]string `json:"bankAccountNumberNin,omitempty"`
	BankAccountNumberExists *bool     `json:"bankAccountNumberExists,omitempty"`
	BankAccountNumberLike   *string   `json:"bankAccountNumberLike,omitempty"`
	BankAccountNumberNlike  *string   `json:"bankAccountNumberNlike,omitempty"`
	// bankIban (string) search options
	BankIbanEq     *string   `json:"bankIbanEq,omitempty"`
	BankIbanNe     *string   `json:"bankIbanNe,omitempty"`
	BankIbanGt     *string   `json:"bankIbanGt,omitempty"`
	BankIbanGte    *string   `json:"bankIbanGte,omitempty"`
	BankIbanLt     *string   `json:"bankIbanLt,omitempty"`
	BankIbanLte    *string   `json:"bankIbanLte,omitempty"`
	BankIbanIn     *[]string `json:"bankIbanIn,omitempty"`
	BankIbanNin    *[]string `json:"bankIbanNin,omitempty"`
	BankIbanExists *bool     `json:"bankIbanExists,omitempty"`
	BankIbanLike   *string   `json:"bankIbanLike,omitempty"`
	BankIbanNlike  *string   `json:"bankIbanNlike,omitempty"`
	// bankName (string) search options
	BankNameEq     *string   `json:"bankNameEq,omitempty"`
	BankNameNe     *string   `json:"bankNameNe,omitempty"`
	BankNameGt     *string   `json:"bankNameGt,omitempty"`
	BankNameGte    *string   `json:"bankNameGte,omitempty"`
	BankNameLt     *string   `json:"bankNameLt,omitempty"`
	BankNameLte    *string   `json:"bankNameLte,omitempty"`
	BankNameIn     *[]string `json:"bankNameIn,omitempty"`
	BankNameNin    *[]string `json:"bankNameNin,omitempty"`
	BankNameExists *bool     `json:"bankNameExists,omitempty"`
	BankNameLike   *string   `json:"bankNameLike,omitempty"`
	BankNameNlike  *string   `json:"bankNameNlike,omitempty"`
	// bankRoutingNumber (string) search options
	BankRoutingNumberEq     *string   `json:"bankRoutingNumberEq,omitempty"`
	BankRoutingNumberNe     *string   `json:"bankRoutingNumberNe,omitempty"`
	BankRoutingNumberGt     *string   `json:"bankRoutingNumberGt,omitempty"`
	BankRoutingNumberGte    *string   `json:"bankRoutingNumberGte,omitempty"`
	BankRoutingNumberLt     *string   `json:"bankRoutingNumberLt,omitempty"`
	BankRoutingNumberLte    *string   `json:"bankRoutingNumberLte,omitempty"`
	BankRoutingNumberIn     *[]string `json:"bankRoutingNumberIn,omitempty"`
	BankRoutingNumberNin    *[]string `json:"bankRoutingNumberNin,omitempty"`
	BankRoutingNumberExists *bool     `json:"bankRoutingNumberExists,omitempty"`
	BankRoutingNumberLike   *string   `json:"bankRoutingNumberLike,omitempty"`
	BankRoutingNumberNlike  *string   `json:"bankRoutingNumberNlike,omitempty"`
	// bankSwift (string) search options
	BankSwiftEq     *string   `json:"bankSwiftEq,omitempty"`
	BankSwiftNe     *string   `json:"bankSwiftNe,omitempty"`
	BankSwiftGt     *string   `json:"bankSwiftGt,omitempty"`
	BankSwiftGte    *string   `json:"bankSwiftGte,omitempty"`
	BankSwiftLt     *string   `json:"bankSwiftLt,omitempty"`
	BankSwiftLte    *string   `json:"bankSwiftLte,omitempty"`
	BankSwiftIn     *[]string `json:"bankSwiftIn,omitempty"`
	BankSwiftNin    *[]string `json:"bankSwiftNin,omitempty"`
	BankSwiftExists *bool     `json:"bankSwiftExists,omitempty"`
	BankSwiftLike   *string   `json:"bankSwiftLike,omitempty"`
	BankSwiftNlike  *string   `json:"bankSwiftNlike,omitempty"`
	// created (ActorTrace) search options
	Created *actor_trace.HTTPWhereClause `json:"created,omitempty"`
	// displayName (string) search options
	DisplayNameEq     *string   `json:"displayNameEq,omitempty"`
	DisplayNameNe     *string   `json:"displayNameNe,omitempty"`
	DisplayNameGt     *string   `json:"displayNameGt,omitempty"`
	DisplayNameGte    *string   `json:"displayNameGte,omitempty"`
	DisplayNameLt     *string   `json:"displayNameLt,omitempty"`
	DisplayNameLte    *string   `json:"displayNameLte,omitempty"`
	DisplayNameIn     *[]string `json:"displayNameIn,omitempty"`
	DisplayNameNin    *[]string `json:"displayNameNin,omitempty"`
	DisplayNameExists *bool     `json:"displayNameExists,omitempty"`
	DisplayNameLike   *string   `json:"displayNameLike,omitempty"`
	DisplayNameNlike  *string   `json:"displayNameNlike,omitempty"`
	// enabled (bool) search options
	EnabledEq     *bool   `json:"enabledEq,omitempty"`
	EnabledNe     *bool   `json:"enabledNe,omitempty"`
	EnabledGt     *bool   `json:"enabledGt,omitempty"`
	EnabledGte    *bool   `json:"enabledGte,omitempty"`
	EnabledLt     *bool   `json:"enabledLt,omitempty"`
	EnabledLte    *bool   `json:"enabledLte,omitempty"`
	EnabledIn     *[]bool `json:"enabledIn,omitempty"`
	EnabledNin    *[]bool `json:"enabledNin,omitempty"`
	EnabledExists *bool   `json:"enabledExists,omitempty"`
	// instructions (string) search options
	InstructionsEq     *string   `json:"instructionsEq,omitempty"`
	InstructionsNe     *string   `json:"instructionsNe,omitempty"`
	InstructionsGt     *string   `json:"instructionsGt,omitempty"`
	InstructionsGte    *string   `json:"instructionsGte,omitempty"`
	InstructionsLt     *string   `json:"instructionsLt,omitempty"`
	InstructionsLte    *string   `json:"instructionsLte,omitempty"`
	InstructionsIn     *[]string `json:"instructionsIn,omitempty"`
	InstructionsNin    *[]string `json:"instructionsNin,omitempty"`
	InstructionsExists *bool     `json:"instructionsExists,omitempty"`
	InstructionsLike   *string   `json:"instructionsLike,omitempty"`
	InstructionsNlike  *string   `json:"instructionsNlike,omitempty"`
	// ownerId (Ref<OwnerUser>) search options
	OwnerIdEq     *string   `json:"ownerIdEq,omitempty"`
	OwnerIdIn     *[]string `json:"ownerIdIn,omitempty"`
	OwnerIdNin    *[]string `json:"ownerIdNin,omitempty"`
	OwnerIdExists *bool     `json:"ownerIdExists,omitempty"`
	// paymentUrl (string) search options
	PaymentUrlEq     *string   `json:"paymentUrlEq,omitempty"`
	PaymentUrlNe     *string   `json:"paymentUrlNe,omitempty"`
	PaymentUrlGt     *string   `json:"paymentUrlGt,omitempty"`
	PaymentUrlGte    *string   `json:"paymentUrlGte,omitempty"`
	PaymentUrlLt     *string   `json:"paymentUrlLt,omitempty"`
	PaymentUrlLte    *string   `json:"paymentUrlLte,omitempty"`
	PaymentUrlIn     *[]string `json:"paymentUrlIn,omitempty"`
	PaymentUrlNin    *[]string `json:"paymentUrlNin,omitempty"`
	PaymentUrlExists *bool     `json:"paymentUrlExists,omitempty"`
	PaymentUrlLike   *string   `json:"paymentUrlLike,omitempty"`
	PaymentUrlNlike  *string   `json:"paymentUrlNlike,omitempty"`
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
	// recipientEmail (string) search options
	RecipientEmailEq     *string   `json:"recipientEmailEq,omitempty"`
	RecipientEmailNe     *string   `json:"recipientEmailNe,omitempty"`
	RecipientEmailGt     *string   `json:"recipientEmailGt,omitempty"`
	RecipientEmailGte    *string   `json:"recipientEmailGte,omitempty"`
	RecipientEmailLt     *string   `json:"recipientEmailLt,omitempty"`
	RecipientEmailLte    *string   `json:"recipientEmailLte,omitempty"`
	RecipientEmailIn     *[]string `json:"recipientEmailIn,omitempty"`
	RecipientEmailNin    *[]string `json:"recipientEmailNin,omitempty"`
	RecipientEmailExists *bool     `json:"recipientEmailExists,omitempty"`
	RecipientEmailLike   *string   `json:"recipientEmailLike,omitempty"`
	RecipientEmailNlike  *string   `json:"recipientEmailNlike,omitempty"`
	// recipientPhone (string) search options
	RecipientPhoneEq     *string   `json:"recipientPhoneEq,omitempty"`
	RecipientPhoneNe     *string   `json:"recipientPhoneNe,omitempty"`
	RecipientPhoneGt     *string   `json:"recipientPhoneGt,omitempty"`
	RecipientPhoneGte    *string   `json:"recipientPhoneGte,omitempty"`
	RecipientPhoneLt     *string   `json:"recipientPhoneLt,omitempty"`
	RecipientPhoneLte    *string   `json:"recipientPhoneLte,omitempty"`
	RecipientPhoneIn     *[]string `json:"recipientPhoneIn,omitempty"`
	RecipientPhoneNin    *[]string `json:"recipientPhoneNin,omitempty"`
	RecipientPhoneExists *bool     `json:"recipientPhoneExists,omitempty"`
	RecipientPhoneLike   *string   `json:"recipientPhoneLike,omitempty"`
	RecipientPhoneNlike  *string   `json:"recipientPhoneNlike,omitempty"`
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *string   `json:"registryIdEq,omitempty"`
	RegistryIdIn     *[]string `json:"registryIdIn,omitempty"`
	RegistryIdNin    *[]string `json:"registryIdNin,omitempty"`
	RegistryIdExists *bool     `json:"registryIdExists,omitempty"`
	// type (PaymentMethodType) search options
	TypeEq     *enum_payment_method_type.Value   `json:"typeEq,omitempty"`
	TypeNe     *enum_payment_method_type.Value   `json:"typeNe,omitempty"`
	TypeGt     *enum_payment_method_type.Value   `json:"typeGt,omitempty"`
	TypeGte    *enum_payment_method_type.Value   `json:"typeGte,omitempty"`
	TypeLt     *enum_payment_method_type.Value   `json:"typeLt,omitempty"`
	TypeLte    *enum_payment_method_type.Value   `json:"typeLte,omitempty"`
	TypeIn     *[]enum_payment_method_type.Value `json:"typeIn,omitempty"`
	TypeNin    *[]enum_payment_method_type.Value `json:"typeNin,omitempty"`
	TypeExists *bool                             `json:"typeExists,omitempty"`
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
	if o.BankAccountNameEq != nil {
		elembankAccountNameEq0 := o.BankAccountNameEq
		to.BankAccountNameEq = elembankAccountNameEq0
	}
	if o.BankAccountNameNe != nil {
		elembankAccountNameNe0 := o.BankAccountNameNe
		to.BankAccountNameNe = elembankAccountNameNe0
	}
	if o.BankAccountNameGt != nil {
		elembankAccountNameGt0 := o.BankAccountNameGt
		to.BankAccountNameGt = elembankAccountNameGt0
	}
	if o.BankAccountNameGte != nil {
		elembankAccountNameGte0 := o.BankAccountNameGte
		to.BankAccountNameGte = elembankAccountNameGte0
	}
	if o.BankAccountNameLt != nil {
		elembankAccountNameLt0 := o.BankAccountNameLt
		to.BankAccountNameLt = elembankAccountNameLt0
	}
	if o.BankAccountNameLte != nil {
		elembankAccountNameLte0 := o.BankAccountNameLte
		to.BankAccountNameLte = elembankAccountNameLte0
	}
	if o.BankAccountNameIn != nil {
		elembankAccountNameIn0 := make([]string, 0)
		for _, obankAccountNameIn0 := range *o.BankAccountNameIn {
			elembankAccountNameIn1 := obankAccountNameIn0
			elembankAccountNameIn0 = append(elembankAccountNameIn0, elembankAccountNameIn1)
		}
		to.BankAccountNameIn = &elembankAccountNameIn0
	}
	if o.BankAccountNameNin != nil {
		elembankAccountNameNin0 := make([]string, 0)
		for _, obankAccountNameNin0 := range *o.BankAccountNameNin {
			elembankAccountNameNin1 := obankAccountNameNin0
			elembankAccountNameNin0 = append(elembankAccountNameNin0, elembankAccountNameNin1)
		}
		to.BankAccountNameNin = &elembankAccountNameNin0
	}
	if o.BankAccountNameExists != nil {
		elembankAccountNameExists0 := o.BankAccountNameExists
		to.BankAccountNameExists = elembankAccountNameExists0
	}
	if o.BankAccountNameLike != nil {
		elembankAccountNameLike0 := o.BankAccountNameLike
		to.BankAccountNameLike = elembankAccountNameLike0
	}
	if o.BankAccountNameNlike != nil {
		elembankAccountNameNlike0 := o.BankAccountNameNlike
		to.BankAccountNameNlike = elembankAccountNameNlike0
	}
	if o.BankAccountNumberEq != nil {
		elembankAccountNumberEq0 := o.BankAccountNumberEq
		to.BankAccountNumberEq = elembankAccountNumberEq0
	}
	if o.BankAccountNumberNe != nil {
		elembankAccountNumberNe0 := o.BankAccountNumberNe
		to.BankAccountNumberNe = elembankAccountNumberNe0
	}
	if o.BankAccountNumberGt != nil {
		elembankAccountNumberGt0 := o.BankAccountNumberGt
		to.BankAccountNumberGt = elembankAccountNumberGt0
	}
	if o.BankAccountNumberGte != nil {
		elembankAccountNumberGte0 := o.BankAccountNumberGte
		to.BankAccountNumberGte = elembankAccountNumberGte0
	}
	if o.BankAccountNumberLt != nil {
		elembankAccountNumberLt0 := o.BankAccountNumberLt
		to.BankAccountNumberLt = elembankAccountNumberLt0
	}
	if o.BankAccountNumberLte != nil {
		elembankAccountNumberLte0 := o.BankAccountNumberLte
		to.BankAccountNumberLte = elembankAccountNumberLte0
	}
	if o.BankAccountNumberIn != nil {
		elembankAccountNumberIn0 := make([]string, 0)
		for _, obankAccountNumberIn0 := range *o.BankAccountNumberIn {
			elembankAccountNumberIn1 := obankAccountNumberIn0
			elembankAccountNumberIn0 = append(elembankAccountNumberIn0, elembankAccountNumberIn1)
		}
		to.BankAccountNumberIn = &elembankAccountNumberIn0
	}
	if o.BankAccountNumberNin != nil {
		elembankAccountNumberNin0 := make([]string, 0)
		for _, obankAccountNumberNin0 := range *o.BankAccountNumberNin {
			elembankAccountNumberNin1 := obankAccountNumberNin0
			elembankAccountNumberNin0 = append(elembankAccountNumberNin0, elembankAccountNumberNin1)
		}
		to.BankAccountNumberNin = &elembankAccountNumberNin0
	}
	if o.BankAccountNumberExists != nil {
		elembankAccountNumberExists0 := o.BankAccountNumberExists
		to.BankAccountNumberExists = elembankAccountNumberExists0
	}
	if o.BankAccountNumberLike != nil {
		elembankAccountNumberLike0 := o.BankAccountNumberLike
		to.BankAccountNumberLike = elembankAccountNumberLike0
	}
	if o.BankAccountNumberNlike != nil {
		elembankAccountNumberNlike0 := o.BankAccountNumberNlike
		to.BankAccountNumberNlike = elembankAccountNumberNlike0
	}
	if o.BankIbanEq != nil {
		elembankIbanEq0 := o.BankIbanEq
		to.BankIbanEq = elembankIbanEq0
	}
	if o.BankIbanNe != nil {
		elembankIbanNe0 := o.BankIbanNe
		to.BankIbanNe = elembankIbanNe0
	}
	if o.BankIbanGt != nil {
		elembankIbanGt0 := o.BankIbanGt
		to.BankIbanGt = elembankIbanGt0
	}
	if o.BankIbanGte != nil {
		elembankIbanGte0 := o.BankIbanGte
		to.BankIbanGte = elembankIbanGte0
	}
	if o.BankIbanLt != nil {
		elembankIbanLt0 := o.BankIbanLt
		to.BankIbanLt = elembankIbanLt0
	}
	if o.BankIbanLte != nil {
		elembankIbanLte0 := o.BankIbanLte
		to.BankIbanLte = elembankIbanLte0
	}
	if o.BankIbanIn != nil {
		elembankIbanIn0 := make([]string, 0)
		for _, obankIbanIn0 := range *o.BankIbanIn {
			elembankIbanIn1 := obankIbanIn0
			elembankIbanIn0 = append(elembankIbanIn0, elembankIbanIn1)
		}
		to.BankIbanIn = &elembankIbanIn0
	}
	if o.BankIbanNin != nil {
		elembankIbanNin0 := make([]string, 0)
		for _, obankIbanNin0 := range *o.BankIbanNin {
			elembankIbanNin1 := obankIbanNin0
			elembankIbanNin0 = append(elembankIbanNin0, elembankIbanNin1)
		}
		to.BankIbanNin = &elembankIbanNin0
	}
	if o.BankIbanExists != nil {
		elembankIbanExists0 := o.BankIbanExists
		to.BankIbanExists = elembankIbanExists0
	}
	if o.BankIbanLike != nil {
		elembankIbanLike0 := o.BankIbanLike
		to.BankIbanLike = elembankIbanLike0
	}
	if o.BankIbanNlike != nil {
		elembankIbanNlike0 := o.BankIbanNlike
		to.BankIbanNlike = elembankIbanNlike0
	}
	if o.BankNameEq != nil {
		elembankNameEq0 := o.BankNameEq
		to.BankNameEq = elembankNameEq0
	}
	if o.BankNameNe != nil {
		elembankNameNe0 := o.BankNameNe
		to.BankNameNe = elembankNameNe0
	}
	if o.BankNameGt != nil {
		elembankNameGt0 := o.BankNameGt
		to.BankNameGt = elembankNameGt0
	}
	if o.BankNameGte != nil {
		elembankNameGte0 := o.BankNameGte
		to.BankNameGte = elembankNameGte0
	}
	if o.BankNameLt != nil {
		elembankNameLt0 := o.BankNameLt
		to.BankNameLt = elembankNameLt0
	}
	if o.BankNameLte != nil {
		elembankNameLte0 := o.BankNameLte
		to.BankNameLte = elembankNameLte0
	}
	if o.BankNameIn != nil {
		elembankNameIn0 := make([]string, 0)
		for _, obankNameIn0 := range *o.BankNameIn {
			elembankNameIn1 := obankNameIn0
			elembankNameIn0 = append(elembankNameIn0, elembankNameIn1)
		}
		to.BankNameIn = &elembankNameIn0
	}
	if o.BankNameNin != nil {
		elembankNameNin0 := make([]string, 0)
		for _, obankNameNin0 := range *o.BankNameNin {
			elembankNameNin1 := obankNameNin0
			elembankNameNin0 = append(elembankNameNin0, elembankNameNin1)
		}
		to.BankNameNin = &elembankNameNin0
	}
	if o.BankNameExists != nil {
		elembankNameExists0 := o.BankNameExists
		to.BankNameExists = elembankNameExists0
	}
	if o.BankNameLike != nil {
		elembankNameLike0 := o.BankNameLike
		to.BankNameLike = elembankNameLike0
	}
	if o.BankNameNlike != nil {
		elembankNameNlike0 := o.BankNameNlike
		to.BankNameNlike = elembankNameNlike0
	}
	if o.BankRoutingNumberEq != nil {
		elembankRoutingNumberEq0 := o.BankRoutingNumberEq
		to.BankRoutingNumberEq = elembankRoutingNumberEq0
	}
	if o.BankRoutingNumberNe != nil {
		elembankRoutingNumberNe0 := o.BankRoutingNumberNe
		to.BankRoutingNumberNe = elembankRoutingNumberNe0
	}
	if o.BankRoutingNumberGt != nil {
		elembankRoutingNumberGt0 := o.BankRoutingNumberGt
		to.BankRoutingNumberGt = elembankRoutingNumberGt0
	}
	if o.BankRoutingNumberGte != nil {
		elembankRoutingNumberGte0 := o.BankRoutingNumberGte
		to.BankRoutingNumberGte = elembankRoutingNumberGte0
	}
	if o.BankRoutingNumberLt != nil {
		elembankRoutingNumberLt0 := o.BankRoutingNumberLt
		to.BankRoutingNumberLt = elembankRoutingNumberLt0
	}
	if o.BankRoutingNumberLte != nil {
		elembankRoutingNumberLte0 := o.BankRoutingNumberLte
		to.BankRoutingNumberLte = elembankRoutingNumberLte0
	}
	if o.BankRoutingNumberIn != nil {
		elembankRoutingNumberIn0 := make([]string, 0)
		for _, obankRoutingNumberIn0 := range *o.BankRoutingNumberIn {
			elembankRoutingNumberIn1 := obankRoutingNumberIn0
			elembankRoutingNumberIn0 = append(elembankRoutingNumberIn0, elembankRoutingNumberIn1)
		}
		to.BankRoutingNumberIn = &elembankRoutingNumberIn0
	}
	if o.BankRoutingNumberNin != nil {
		elembankRoutingNumberNin0 := make([]string, 0)
		for _, obankRoutingNumberNin0 := range *o.BankRoutingNumberNin {
			elembankRoutingNumberNin1 := obankRoutingNumberNin0
			elembankRoutingNumberNin0 = append(elembankRoutingNumberNin0, elembankRoutingNumberNin1)
		}
		to.BankRoutingNumberNin = &elembankRoutingNumberNin0
	}
	if o.BankRoutingNumberExists != nil {
		elembankRoutingNumberExists0 := o.BankRoutingNumberExists
		to.BankRoutingNumberExists = elembankRoutingNumberExists0
	}
	if o.BankRoutingNumberLike != nil {
		elembankRoutingNumberLike0 := o.BankRoutingNumberLike
		to.BankRoutingNumberLike = elembankRoutingNumberLike0
	}
	if o.BankRoutingNumberNlike != nil {
		elembankRoutingNumberNlike0 := o.BankRoutingNumberNlike
		to.BankRoutingNumberNlike = elembankRoutingNumberNlike0
	}
	if o.BankSwiftEq != nil {
		elembankSwiftEq0 := o.BankSwiftEq
		to.BankSwiftEq = elembankSwiftEq0
	}
	if o.BankSwiftNe != nil {
		elembankSwiftNe0 := o.BankSwiftNe
		to.BankSwiftNe = elembankSwiftNe0
	}
	if o.BankSwiftGt != nil {
		elembankSwiftGt0 := o.BankSwiftGt
		to.BankSwiftGt = elembankSwiftGt0
	}
	if o.BankSwiftGte != nil {
		elembankSwiftGte0 := o.BankSwiftGte
		to.BankSwiftGte = elembankSwiftGte0
	}
	if o.BankSwiftLt != nil {
		elembankSwiftLt0 := o.BankSwiftLt
		to.BankSwiftLt = elembankSwiftLt0
	}
	if o.BankSwiftLte != nil {
		elembankSwiftLte0 := o.BankSwiftLte
		to.BankSwiftLte = elembankSwiftLte0
	}
	if o.BankSwiftIn != nil {
		elembankSwiftIn0 := make([]string, 0)
		for _, obankSwiftIn0 := range *o.BankSwiftIn {
			elembankSwiftIn1 := obankSwiftIn0
			elembankSwiftIn0 = append(elembankSwiftIn0, elembankSwiftIn1)
		}
		to.BankSwiftIn = &elembankSwiftIn0
	}
	if o.BankSwiftNin != nil {
		elembankSwiftNin0 := make([]string, 0)
		for _, obankSwiftNin0 := range *o.BankSwiftNin {
			elembankSwiftNin1 := obankSwiftNin0
			elembankSwiftNin0 = append(elembankSwiftNin0, elembankSwiftNin1)
		}
		to.BankSwiftNin = &elembankSwiftNin0
	}
	if o.BankSwiftExists != nil {
		elembankSwiftExists0 := o.BankSwiftExists
		to.BankSwiftExists = elembankSwiftExists0
	}
	if o.BankSwiftLike != nil {
		elembankSwiftLike0 := o.BankSwiftLike
		to.BankSwiftLike = elembankSwiftLike0
	}
	if o.BankSwiftNlike != nil {
		elembankSwiftNlike0 := o.BankSwiftNlike
		to.BankSwiftNlike = elembankSwiftNlike0
	}
	if o.Created != nil {
		elemcreated0, err := o.Created.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.Created = &elemcreated0
	}
	if o.DisplayNameEq != nil {
		elemdisplayNameEq0 := o.DisplayNameEq
		to.DisplayNameEq = elemdisplayNameEq0
	}
	if o.DisplayNameNe != nil {
		elemdisplayNameNe0 := o.DisplayNameNe
		to.DisplayNameNe = elemdisplayNameNe0
	}
	if o.DisplayNameGt != nil {
		elemdisplayNameGt0 := o.DisplayNameGt
		to.DisplayNameGt = elemdisplayNameGt0
	}
	if o.DisplayNameGte != nil {
		elemdisplayNameGte0 := o.DisplayNameGte
		to.DisplayNameGte = elemdisplayNameGte0
	}
	if o.DisplayNameLt != nil {
		elemdisplayNameLt0 := o.DisplayNameLt
		to.DisplayNameLt = elemdisplayNameLt0
	}
	if o.DisplayNameLte != nil {
		elemdisplayNameLte0 := o.DisplayNameLte
		to.DisplayNameLte = elemdisplayNameLte0
	}
	if o.DisplayNameIn != nil {
		elemdisplayNameIn0 := make([]string, 0)
		for _, odisplayNameIn0 := range *o.DisplayNameIn {
			elemdisplayNameIn1 := odisplayNameIn0
			elemdisplayNameIn0 = append(elemdisplayNameIn0, elemdisplayNameIn1)
		}
		to.DisplayNameIn = &elemdisplayNameIn0
	}
	if o.DisplayNameNin != nil {
		elemdisplayNameNin0 := make([]string, 0)
		for _, odisplayNameNin0 := range *o.DisplayNameNin {
			elemdisplayNameNin1 := odisplayNameNin0
			elemdisplayNameNin0 = append(elemdisplayNameNin0, elemdisplayNameNin1)
		}
		to.DisplayNameNin = &elemdisplayNameNin0
	}
	if o.DisplayNameExists != nil {
		elemdisplayNameExists0 := o.DisplayNameExists
		to.DisplayNameExists = elemdisplayNameExists0
	}
	if o.DisplayNameLike != nil {
		elemdisplayNameLike0 := o.DisplayNameLike
		to.DisplayNameLike = elemdisplayNameLike0
	}
	if o.DisplayNameNlike != nil {
		elemdisplayNameNlike0 := o.DisplayNameNlike
		to.DisplayNameNlike = elemdisplayNameNlike0
	}
	if o.EnabledEq != nil {
		elemenabledEq0 := o.EnabledEq
		to.EnabledEq = elemenabledEq0
	}
	if o.EnabledNe != nil {
		elemenabledNe0 := o.EnabledNe
		to.EnabledNe = elemenabledNe0
	}
	if o.EnabledGt != nil {
		elemenabledGt0 := o.EnabledGt
		to.EnabledGt = elemenabledGt0
	}
	if o.EnabledGte != nil {
		elemenabledGte0 := o.EnabledGte
		to.EnabledGte = elemenabledGte0
	}
	if o.EnabledLt != nil {
		elemenabledLt0 := o.EnabledLt
		to.EnabledLt = elemenabledLt0
	}
	if o.EnabledLte != nil {
		elemenabledLte0 := o.EnabledLte
		to.EnabledLte = elemenabledLte0
	}
	if o.EnabledIn != nil {
		elemenabledIn0 := make([]bool, 0)
		for _, oenabledIn0 := range *o.EnabledIn {
			elemenabledIn1 := oenabledIn0
			elemenabledIn0 = append(elemenabledIn0, elemenabledIn1)
		}
		to.EnabledIn = &elemenabledIn0
	}
	if o.EnabledNin != nil {
		elemenabledNin0 := make([]bool, 0)
		for _, oenabledNin0 := range *o.EnabledNin {
			elemenabledNin1 := oenabledNin0
			elemenabledNin0 = append(elemenabledNin0, elemenabledNin1)
		}
		to.EnabledNin = &elemenabledNin0
	}
	if o.EnabledExists != nil {
		elemenabledExists0 := o.EnabledExists
		to.EnabledExists = elemenabledExists0
	}
	if o.InstructionsEq != nil {
		eleminstructionsEq0 := o.InstructionsEq
		to.InstructionsEq = eleminstructionsEq0
	}
	if o.InstructionsNe != nil {
		eleminstructionsNe0 := o.InstructionsNe
		to.InstructionsNe = eleminstructionsNe0
	}
	if o.InstructionsGt != nil {
		eleminstructionsGt0 := o.InstructionsGt
		to.InstructionsGt = eleminstructionsGt0
	}
	if o.InstructionsGte != nil {
		eleminstructionsGte0 := o.InstructionsGte
		to.InstructionsGte = eleminstructionsGte0
	}
	if o.InstructionsLt != nil {
		eleminstructionsLt0 := o.InstructionsLt
		to.InstructionsLt = eleminstructionsLt0
	}
	if o.InstructionsLte != nil {
		eleminstructionsLte0 := o.InstructionsLte
		to.InstructionsLte = eleminstructionsLte0
	}
	if o.InstructionsIn != nil {
		eleminstructionsIn0 := make([]string, 0)
		for _, oinstructionsIn0 := range *o.InstructionsIn {
			eleminstructionsIn1 := oinstructionsIn0
			eleminstructionsIn0 = append(eleminstructionsIn0, eleminstructionsIn1)
		}
		to.InstructionsIn = &eleminstructionsIn0
	}
	if o.InstructionsNin != nil {
		eleminstructionsNin0 := make([]string, 0)
		for _, oinstructionsNin0 := range *o.InstructionsNin {
			eleminstructionsNin1 := oinstructionsNin0
			eleminstructionsNin0 = append(eleminstructionsNin0, eleminstructionsNin1)
		}
		to.InstructionsNin = &eleminstructionsNin0
	}
	if o.InstructionsExists != nil {
		eleminstructionsExists0 := o.InstructionsExists
		to.InstructionsExists = eleminstructionsExists0
	}
	if o.InstructionsLike != nil {
		eleminstructionsLike0 := o.InstructionsLike
		to.InstructionsLike = eleminstructionsLike0
	}
	if o.InstructionsNlike != nil {
		eleminstructionsNlike0 := o.InstructionsNlike
		to.InstructionsNlike = eleminstructionsNlike0
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
	if o.PaymentUrlEq != nil {
		elempaymentUrlEq0 := o.PaymentUrlEq
		to.PaymentUrlEq = elempaymentUrlEq0
	}
	if o.PaymentUrlNe != nil {
		elempaymentUrlNe0 := o.PaymentUrlNe
		to.PaymentUrlNe = elempaymentUrlNe0
	}
	if o.PaymentUrlGt != nil {
		elempaymentUrlGt0 := o.PaymentUrlGt
		to.PaymentUrlGt = elempaymentUrlGt0
	}
	if o.PaymentUrlGte != nil {
		elempaymentUrlGte0 := o.PaymentUrlGte
		to.PaymentUrlGte = elempaymentUrlGte0
	}
	if o.PaymentUrlLt != nil {
		elempaymentUrlLt0 := o.PaymentUrlLt
		to.PaymentUrlLt = elempaymentUrlLt0
	}
	if o.PaymentUrlLte != nil {
		elempaymentUrlLte0 := o.PaymentUrlLte
		to.PaymentUrlLte = elempaymentUrlLte0
	}
	if o.PaymentUrlIn != nil {
		elempaymentUrlIn0 := make([]string, 0)
		for _, opaymentUrlIn0 := range *o.PaymentUrlIn {
			elempaymentUrlIn1 := opaymentUrlIn0
			elempaymentUrlIn0 = append(elempaymentUrlIn0, elempaymentUrlIn1)
		}
		to.PaymentUrlIn = &elempaymentUrlIn0
	}
	if o.PaymentUrlNin != nil {
		elempaymentUrlNin0 := make([]string, 0)
		for _, opaymentUrlNin0 := range *o.PaymentUrlNin {
			elempaymentUrlNin1 := opaymentUrlNin0
			elempaymentUrlNin0 = append(elempaymentUrlNin0, elempaymentUrlNin1)
		}
		to.PaymentUrlNin = &elempaymentUrlNin0
	}
	if o.PaymentUrlExists != nil {
		elempaymentUrlExists0 := o.PaymentUrlExists
		to.PaymentUrlExists = elempaymentUrlExists0
	}
	if o.PaymentUrlLike != nil {
		elempaymentUrlLike0 := o.PaymentUrlLike
		to.PaymentUrlLike = elempaymentUrlLike0
	}
	if o.PaymentUrlNlike != nil {
		elempaymentUrlNlike0 := o.PaymentUrlNlike
		to.PaymentUrlNlike = elempaymentUrlNlike0
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
	if o.RecipientEmailEq != nil {
		elemrecipientEmailEq0 := o.RecipientEmailEq
		to.RecipientEmailEq = elemrecipientEmailEq0
	}
	if o.RecipientEmailNe != nil {
		elemrecipientEmailNe0 := o.RecipientEmailNe
		to.RecipientEmailNe = elemrecipientEmailNe0
	}
	if o.RecipientEmailGt != nil {
		elemrecipientEmailGt0 := o.RecipientEmailGt
		to.RecipientEmailGt = elemrecipientEmailGt0
	}
	if o.RecipientEmailGte != nil {
		elemrecipientEmailGte0 := o.RecipientEmailGte
		to.RecipientEmailGte = elemrecipientEmailGte0
	}
	if o.RecipientEmailLt != nil {
		elemrecipientEmailLt0 := o.RecipientEmailLt
		to.RecipientEmailLt = elemrecipientEmailLt0
	}
	if o.RecipientEmailLte != nil {
		elemrecipientEmailLte0 := o.RecipientEmailLte
		to.RecipientEmailLte = elemrecipientEmailLte0
	}
	if o.RecipientEmailIn != nil {
		elemrecipientEmailIn0 := make([]string, 0)
		for _, orecipientEmailIn0 := range *o.RecipientEmailIn {
			elemrecipientEmailIn1 := orecipientEmailIn0
			elemrecipientEmailIn0 = append(elemrecipientEmailIn0, elemrecipientEmailIn1)
		}
		to.RecipientEmailIn = &elemrecipientEmailIn0
	}
	if o.RecipientEmailNin != nil {
		elemrecipientEmailNin0 := make([]string, 0)
		for _, orecipientEmailNin0 := range *o.RecipientEmailNin {
			elemrecipientEmailNin1 := orecipientEmailNin0
			elemrecipientEmailNin0 = append(elemrecipientEmailNin0, elemrecipientEmailNin1)
		}
		to.RecipientEmailNin = &elemrecipientEmailNin0
	}
	if o.RecipientEmailExists != nil {
		elemrecipientEmailExists0 := o.RecipientEmailExists
		to.RecipientEmailExists = elemrecipientEmailExists0
	}
	if o.RecipientEmailLike != nil {
		elemrecipientEmailLike0 := o.RecipientEmailLike
		to.RecipientEmailLike = elemrecipientEmailLike0
	}
	if o.RecipientEmailNlike != nil {
		elemrecipientEmailNlike0 := o.RecipientEmailNlike
		to.RecipientEmailNlike = elemrecipientEmailNlike0
	}
	if o.RecipientPhoneEq != nil {
		elemrecipientPhoneEq0 := o.RecipientPhoneEq
		to.RecipientPhoneEq = elemrecipientPhoneEq0
	}
	if o.RecipientPhoneNe != nil {
		elemrecipientPhoneNe0 := o.RecipientPhoneNe
		to.RecipientPhoneNe = elemrecipientPhoneNe0
	}
	if o.RecipientPhoneGt != nil {
		elemrecipientPhoneGt0 := o.RecipientPhoneGt
		to.RecipientPhoneGt = elemrecipientPhoneGt0
	}
	if o.RecipientPhoneGte != nil {
		elemrecipientPhoneGte0 := o.RecipientPhoneGte
		to.RecipientPhoneGte = elemrecipientPhoneGte0
	}
	if o.RecipientPhoneLt != nil {
		elemrecipientPhoneLt0 := o.RecipientPhoneLt
		to.RecipientPhoneLt = elemrecipientPhoneLt0
	}
	if o.RecipientPhoneLte != nil {
		elemrecipientPhoneLte0 := o.RecipientPhoneLte
		to.RecipientPhoneLte = elemrecipientPhoneLte0
	}
	if o.RecipientPhoneIn != nil {
		elemrecipientPhoneIn0 := make([]string, 0)
		for _, orecipientPhoneIn0 := range *o.RecipientPhoneIn {
			elemrecipientPhoneIn1 := orecipientPhoneIn0
			elemrecipientPhoneIn0 = append(elemrecipientPhoneIn0, elemrecipientPhoneIn1)
		}
		to.RecipientPhoneIn = &elemrecipientPhoneIn0
	}
	if o.RecipientPhoneNin != nil {
		elemrecipientPhoneNin0 := make([]string, 0)
		for _, orecipientPhoneNin0 := range *o.RecipientPhoneNin {
			elemrecipientPhoneNin1 := orecipientPhoneNin0
			elemrecipientPhoneNin0 = append(elemrecipientPhoneNin0, elemrecipientPhoneNin1)
		}
		to.RecipientPhoneNin = &elemrecipientPhoneNin0
	}
	if o.RecipientPhoneExists != nil {
		elemrecipientPhoneExists0 := o.RecipientPhoneExists
		to.RecipientPhoneExists = elemrecipientPhoneExists0
	}
	if o.RecipientPhoneLike != nil {
		elemrecipientPhoneLike0 := o.RecipientPhoneLike
		to.RecipientPhoneLike = elemrecipientPhoneLike0
	}
	if o.RecipientPhoneNlike != nil {
		elemrecipientPhoneNlike0 := o.RecipientPhoneNlike
		to.RecipientPhoneNlike = elemrecipientPhoneNlike0
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
	if o.TypeEq != nil {
		elemtypeEq0 := o.TypeEq
		to.TypeEq = elemtypeEq0
	}
	if o.TypeNe != nil {
		elemtypeNe0 := o.TypeNe
		to.TypeNe = elemtypeNe0
	}
	if o.TypeGt != nil {
		elemtypeGt0 := o.TypeGt
		to.TypeGt = elemtypeGt0
	}
	if o.TypeGte != nil {
		elemtypeGte0 := o.TypeGte
		to.TypeGte = elemtypeGte0
	}
	if o.TypeLt != nil {
		elemtypeLt0 := o.TypeLt
		to.TypeLt = elemtypeLt0
	}
	if o.TypeLte != nil {
		elemtypeLte0 := o.TypeLte
		to.TypeLte = elemtypeLte0
	}
	if o.TypeIn != nil {
		elemtypeIn0 := make([]enum_payment_method_type.Value, 0)
		for _, otypeIn0 := range *o.TypeIn {
			elemtypeIn1 := otypeIn0
			elemtypeIn0 = append(elemtypeIn0, elemtypeIn1)
		}
		to.TypeIn = &elemtypeIn0
	}
	if o.TypeNin != nil {
		elemtypeNin0 := make([]enum_payment_method_type.Value, 0)
		for _, otypeNin0 := range *o.TypeNin {
			elemtypeNin1 := otypeNin0
			elemtypeNin0 = append(elemtypeNin0, elemtypeNin1)
		}
		to.TypeNin = &elemtypeNin0
	}
	if o.TypeExists != nil {
		elemtypeExists0 := o.TypeExists
		to.TypeExists = elemtypeExists0
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
	OwnerId    *int8 `json:"ownerId,omitempty"`
	RegistryId *int8 `json:"registryId,omitempty"`
	UpdatedAt  *int8 `json:"updatedAt,omitempty"`
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	if s.CreatedAt != nil {
		to.CreatedAt = *s.CreatedAt
	}
	if s.OwnerId != nil {
		to.OwnerId = *s.OwnerId
	}
	if s.RegistryId != nil {
		to.RegistryId = *s.RegistryId
	}
	if s.UpdatedAt != nil {
		to.UpdatedAt = *s.UpdatedAt
	}
	return to
}
