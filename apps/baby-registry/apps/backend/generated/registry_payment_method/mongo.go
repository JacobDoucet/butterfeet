package registry_payment_method

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_payment_method_type"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRecord struct {
	Id                 *primitive.ObjectID             `bson:"_id,omitempty"`
	BankAccountName    *string                         `bson:"bankAccountName,omitempty"`
	BankAccountNumber  *string                         `bson:"bankAccountNumber,omitempty"`
	BankIban           *string                         `bson:"bankIban,omitempty"`
	BankName           *string                         `bson:"bankName,omitempty"`
	BankRoutingNumber  *string                         `bson:"bankRoutingNumber,omitempty"`
	BankSwift          *string                         `bson:"bankSwift,omitempty"`
	Created            *actor_trace.MongoRecord        `bson:"created,omitempty"`
	DisplayName        *string                         `bson:"displayName,omitempty"`
	Enabled            *bool                           `bson:"enabled,omitempty"`
	Instructions       *string                         `bson:"instructions,omitempty"`
	OwnerId            *primitive.ObjectID             `bson:"ownerId,omitempty"`
	PaymentUrl         *string                         `bson:"paymentUrl,omitempty"`
	Position           *int                            `bson:"position,omitempty"`
	RecipientEmail     *string                         `bson:"recipientEmail,omitempty"`
	RecipientPhone     *string                         `bson:"recipientPhone,omitempty"`
	RegistryId         *primitive.ObjectID             `bson:"registryId,omitempty"`
	Type               *enum_payment_method_type.Value `bson:"type,omitempty"`
	Updated            *actor_trace.MongoRecord        `bson:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.MongoRecord        `bson:"updatedByOwnerUser,omitempty"`
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
		elemownerId0 := r.OwnerId.Hex()
		m.OwnerId = elemownerId0
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
		elemregistryId0 := r.RegistryId.Hex()
		m.RegistryId = elemregistryId0
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

type MongoSelectByIdQuery struct {
	Id primitive.ObjectID
}

type MongoWhereClause struct {
	// id (Ref<RegistryPaymentMethod>) search options
	IdEq     *primitive.ObjectID
	IdIn     *[]primitive.ObjectID
	IdNin    *[]primitive.ObjectID
	IdExists *bool
	// bankAccountName (string) search options
	BankAccountNameEq     *string
	BankAccountNameNe     *string
	BankAccountNameGt     *string
	BankAccountNameGte    *string
	BankAccountNameLt     *string
	BankAccountNameLte    *string
	BankAccountNameIn     *[]string
	BankAccountNameNin    *[]string
	BankAccountNameExists *bool
	BankAccountNameLike   *string
	BankAccountNameNlike  *string
	// bankAccountNumber (string) search options
	BankAccountNumberEq     *string
	BankAccountNumberNe     *string
	BankAccountNumberGt     *string
	BankAccountNumberGte    *string
	BankAccountNumberLt     *string
	BankAccountNumberLte    *string
	BankAccountNumberIn     *[]string
	BankAccountNumberNin    *[]string
	BankAccountNumberExists *bool
	BankAccountNumberLike   *string
	BankAccountNumberNlike  *string
	// bankIban (string) search options
	BankIbanEq     *string
	BankIbanNe     *string
	BankIbanGt     *string
	BankIbanGte    *string
	BankIbanLt     *string
	BankIbanLte    *string
	BankIbanIn     *[]string
	BankIbanNin    *[]string
	BankIbanExists *bool
	BankIbanLike   *string
	BankIbanNlike  *string
	// bankName (string) search options
	BankNameEq     *string
	BankNameNe     *string
	BankNameGt     *string
	BankNameGte    *string
	BankNameLt     *string
	BankNameLte    *string
	BankNameIn     *[]string
	BankNameNin    *[]string
	BankNameExists *bool
	BankNameLike   *string
	BankNameNlike  *string
	// bankRoutingNumber (string) search options
	BankRoutingNumberEq     *string
	BankRoutingNumberNe     *string
	BankRoutingNumberGt     *string
	BankRoutingNumberGte    *string
	BankRoutingNumberLt     *string
	BankRoutingNumberLte    *string
	BankRoutingNumberIn     *[]string
	BankRoutingNumberNin    *[]string
	BankRoutingNumberExists *bool
	BankRoutingNumberLike   *string
	BankRoutingNumberNlike  *string
	// bankSwift (string) search options
	BankSwiftEq     *string
	BankSwiftNe     *string
	BankSwiftGt     *string
	BankSwiftGte    *string
	BankSwiftLt     *string
	BankSwiftLte    *string
	BankSwiftIn     *[]string
	BankSwiftNin    *[]string
	BankSwiftExists *bool
	BankSwiftLike   *string
	BankSwiftNlike  *string
	// created (ActorTrace) search options
	Created *actor_trace.MongoWhereClause
	// displayName (string) search options
	DisplayNameEq     *string
	DisplayNameNe     *string
	DisplayNameGt     *string
	DisplayNameGte    *string
	DisplayNameLt     *string
	DisplayNameLte    *string
	DisplayNameIn     *[]string
	DisplayNameNin    *[]string
	DisplayNameExists *bool
	DisplayNameLike   *string
	DisplayNameNlike  *string
	// enabled (bool) search options
	EnabledEq     *bool
	EnabledNe     *bool
	EnabledGt     *bool
	EnabledGte    *bool
	EnabledLt     *bool
	EnabledLte    *bool
	EnabledIn     *[]bool
	EnabledNin    *[]bool
	EnabledExists *bool
	// instructions (string) search options
	InstructionsEq     *string
	InstructionsNe     *string
	InstructionsGt     *string
	InstructionsGte    *string
	InstructionsLt     *string
	InstructionsLte    *string
	InstructionsIn     *[]string
	InstructionsNin    *[]string
	InstructionsExists *bool
	InstructionsLike   *string
	InstructionsNlike  *string
	// ownerId (Ref<OwnerUser>) search options
	OwnerIdEq     *primitive.ObjectID
	OwnerIdIn     *[]primitive.ObjectID
	OwnerIdNin    *[]primitive.ObjectID
	OwnerIdExists *bool
	// paymentUrl (string) search options
	PaymentUrlEq     *string
	PaymentUrlNe     *string
	PaymentUrlGt     *string
	PaymentUrlGte    *string
	PaymentUrlLt     *string
	PaymentUrlLte    *string
	PaymentUrlIn     *[]string
	PaymentUrlNin    *[]string
	PaymentUrlExists *bool
	PaymentUrlLike   *string
	PaymentUrlNlike  *string
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
	// recipientEmail (string) search options
	RecipientEmailEq     *string
	RecipientEmailNe     *string
	RecipientEmailGt     *string
	RecipientEmailGte    *string
	RecipientEmailLt     *string
	RecipientEmailLte    *string
	RecipientEmailIn     *[]string
	RecipientEmailNin    *[]string
	RecipientEmailExists *bool
	RecipientEmailLike   *string
	RecipientEmailNlike  *string
	// recipientPhone (string) search options
	RecipientPhoneEq     *string
	RecipientPhoneNe     *string
	RecipientPhoneGt     *string
	RecipientPhoneGte    *string
	RecipientPhoneLt     *string
	RecipientPhoneLte    *string
	RecipientPhoneIn     *[]string
	RecipientPhoneNin    *[]string
	RecipientPhoneExists *bool
	RecipientPhoneLike   *string
	RecipientPhoneNlike  *string
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *primitive.ObjectID
	RegistryIdIn     *[]primitive.ObjectID
	RegistryIdNin    *[]primitive.ObjectID
	RegistryIdExists *bool
	// type (PaymentMethodType) search options
	TypeEq     *enum_payment_method_type.Value
	TypeNe     *enum_payment_method_type.Value
	TypeGt     *enum_payment_method_type.Value
	TypeGte    *enum_payment_method_type.Value
	TypeLt     *enum_payment_method_type.Value
	TypeLte    *enum_payment_method_type.Value
	TypeIn     *[]enum_payment_method_type.Value
	TypeNin    *[]enum_payment_method_type.Value
	TypeExists *bool
	// updated (ActorTrace) search options
	Updated *actor_trace.MongoWhereClause
	// updatedByOwnerUser (ActorTrace) search options
	UpdatedByOwnerUser *actor_trace.MongoWhereClause
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
	if o.BankAccountNameEq != nil {
		query := bson.M{}
		query["bankAccountName"] = o.BankAccountNameEq
		and = append(and, query)
	}
	if o.BankAccountNameNe != nil {
		query := bson.M{}
		query["bankAccountName"] = bson.M{"$ne": o.BankAccountNameNe}
		and = append(and, query)
	}
	if o.BankAccountNameGt != nil {
		query := bson.M{}
		query["bankAccountName"] = bson.M{"$gt": o.BankAccountNameGt}
		and = append(and, query)
	}
	if o.BankAccountNameGte != nil {
		query := bson.M{}
		query["bankAccountName"] = bson.M{"$gte": o.BankAccountNameGte}
		and = append(and, query)
	}
	if o.BankAccountNameLt != nil {
		query := bson.M{}
		query["bankAccountName"] = bson.M{"$lt": o.BankAccountNameLt}
		and = append(and, query)
	}
	if o.BankAccountNameLte != nil {
		query := bson.M{}
		query["bankAccountName"] = bson.M{"$lte": o.BankAccountNameLte}
		and = append(and, query)
	}
	if o.BankAccountNameIn != nil {
		query := bson.M{}
		query["bankAccountName"] = bson.M{"$in": o.BankAccountNameIn}
		and = append(and, query)
	}
	if o.BankAccountNameNin != nil {
		query := bson.M{}
		query["bankAccountName"] = bson.M{"$nin": o.BankAccountNameNin}
		and = append(and, query)
	}
	if o.BankAccountNameExists != nil {
		query := bson.M{}
		query["bankAccountName"] = bson.M{"$exists": *o.BankAccountNameExists}
		and = append(and, query)
	}
	if o.BankAccountNameLike != nil {
		query := bson.M{}
		query["bankAccountName"] = bson.M{"$regex": o.BankAccountNameLike, "$options": "i"}
		and = append(and, query)
	}
	if o.BankAccountNameNlike != nil {
		query := bson.M{}
		query["bankAccountName"] = bson.M{"$not": bson.M{"$regex": o.BankAccountNameNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.BankAccountNumberEq != nil {
		query := bson.M{}
		query["bankAccountNumber"] = o.BankAccountNumberEq
		and = append(and, query)
	}
	if o.BankAccountNumberNe != nil {
		query := bson.M{}
		query["bankAccountNumber"] = bson.M{"$ne": o.BankAccountNumberNe}
		and = append(and, query)
	}
	if o.BankAccountNumberGt != nil {
		query := bson.M{}
		query["bankAccountNumber"] = bson.M{"$gt": o.BankAccountNumberGt}
		and = append(and, query)
	}
	if o.BankAccountNumberGte != nil {
		query := bson.M{}
		query["bankAccountNumber"] = bson.M{"$gte": o.BankAccountNumberGte}
		and = append(and, query)
	}
	if o.BankAccountNumberLt != nil {
		query := bson.M{}
		query["bankAccountNumber"] = bson.M{"$lt": o.BankAccountNumberLt}
		and = append(and, query)
	}
	if o.BankAccountNumberLte != nil {
		query := bson.M{}
		query["bankAccountNumber"] = bson.M{"$lte": o.BankAccountNumberLte}
		and = append(and, query)
	}
	if o.BankAccountNumberIn != nil {
		query := bson.M{}
		query["bankAccountNumber"] = bson.M{"$in": o.BankAccountNumberIn}
		and = append(and, query)
	}
	if o.BankAccountNumberNin != nil {
		query := bson.M{}
		query["bankAccountNumber"] = bson.M{"$nin": o.BankAccountNumberNin}
		and = append(and, query)
	}
	if o.BankAccountNumberExists != nil {
		query := bson.M{}
		query["bankAccountNumber"] = bson.M{"$exists": *o.BankAccountNumberExists}
		and = append(and, query)
	}
	if o.BankAccountNumberLike != nil {
		query := bson.M{}
		query["bankAccountNumber"] = bson.M{"$regex": o.BankAccountNumberLike, "$options": "i"}
		and = append(and, query)
	}
	if o.BankAccountNumberNlike != nil {
		query := bson.M{}
		query["bankAccountNumber"] = bson.M{"$not": bson.M{"$regex": o.BankAccountNumberNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.BankIbanEq != nil {
		query := bson.M{}
		query["bankIban"] = o.BankIbanEq
		and = append(and, query)
	}
	if o.BankIbanNe != nil {
		query := bson.M{}
		query["bankIban"] = bson.M{"$ne": o.BankIbanNe}
		and = append(and, query)
	}
	if o.BankIbanGt != nil {
		query := bson.M{}
		query["bankIban"] = bson.M{"$gt": o.BankIbanGt}
		and = append(and, query)
	}
	if o.BankIbanGte != nil {
		query := bson.M{}
		query["bankIban"] = bson.M{"$gte": o.BankIbanGte}
		and = append(and, query)
	}
	if o.BankIbanLt != nil {
		query := bson.M{}
		query["bankIban"] = bson.M{"$lt": o.BankIbanLt}
		and = append(and, query)
	}
	if o.BankIbanLte != nil {
		query := bson.M{}
		query["bankIban"] = bson.M{"$lte": o.BankIbanLte}
		and = append(and, query)
	}
	if o.BankIbanIn != nil {
		query := bson.M{}
		query["bankIban"] = bson.M{"$in": o.BankIbanIn}
		and = append(and, query)
	}
	if o.BankIbanNin != nil {
		query := bson.M{}
		query["bankIban"] = bson.M{"$nin": o.BankIbanNin}
		and = append(and, query)
	}
	if o.BankIbanExists != nil {
		query := bson.M{}
		query["bankIban"] = bson.M{"$exists": *o.BankIbanExists}
		and = append(and, query)
	}
	if o.BankIbanLike != nil {
		query := bson.M{}
		query["bankIban"] = bson.M{"$regex": o.BankIbanLike, "$options": "i"}
		and = append(and, query)
	}
	if o.BankIbanNlike != nil {
		query := bson.M{}
		query["bankIban"] = bson.M{"$not": bson.M{"$regex": o.BankIbanNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.BankNameEq != nil {
		query := bson.M{}
		query["bankName"] = o.BankNameEq
		and = append(and, query)
	}
	if o.BankNameNe != nil {
		query := bson.M{}
		query["bankName"] = bson.M{"$ne": o.BankNameNe}
		and = append(and, query)
	}
	if o.BankNameGt != nil {
		query := bson.M{}
		query["bankName"] = bson.M{"$gt": o.BankNameGt}
		and = append(and, query)
	}
	if o.BankNameGte != nil {
		query := bson.M{}
		query["bankName"] = bson.M{"$gte": o.BankNameGte}
		and = append(and, query)
	}
	if o.BankNameLt != nil {
		query := bson.M{}
		query["bankName"] = bson.M{"$lt": o.BankNameLt}
		and = append(and, query)
	}
	if o.BankNameLte != nil {
		query := bson.M{}
		query["bankName"] = bson.M{"$lte": o.BankNameLte}
		and = append(and, query)
	}
	if o.BankNameIn != nil {
		query := bson.M{}
		query["bankName"] = bson.M{"$in": o.BankNameIn}
		and = append(and, query)
	}
	if o.BankNameNin != nil {
		query := bson.M{}
		query["bankName"] = bson.M{"$nin": o.BankNameNin}
		and = append(and, query)
	}
	if o.BankNameExists != nil {
		query := bson.M{}
		query["bankName"] = bson.M{"$exists": *o.BankNameExists}
		and = append(and, query)
	}
	if o.BankNameLike != nil {
		query := bson.M{}
		query["bankName"] = bson.M{"$regex": o.BankNameLike, "$options": "i"}
		and = append(and, query)
	}
	if o.BankNameNlike != nil {
		query := bson.M{}
		query["bankName"] = bson.M{"$not": bson.M{"$regex": o.BankNameNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.BankRoutingNumberEq != nil {
		query := bson.M{}
		query["bankRoutingNumber"] = o.BankRoutingNumberEq
		and = append(and, query)
	}
	if o.BankRoutingNumberNe != nil {
		query := bson.M{}
		query["bankRoutingNumber"] = bson.M{"$ne": o.BankRoutingNumberNe}
		and = append(and, query)
	}
	if o.BankRoutingNumberGt != nil {
		query := bson.M{}
		query["bankRoutingNumber"] = bson.M{"$gt": o.BankRoutingNumberGt}
		and = append(and, query)
	}
	if o.BankRoutingNumberGte != nil {
		query := bson.M{}
		query["bankRoutingNumber"] = bson.M{"$gte": o.BankRoutingNumberGte}
		and = append(and, query)
	}
	if o.BankRoutingNumberLt != nil {
		query := bson.M{}
		query["bankRoutingNumber"] = bson.M{"$lt": o.BankRoutingNumberLt}
		and = append(and, query)
	}
	if o.BankRoutingNumberLte != nil {
		query := bson.M{}
		query["bankRoutingNumber"] = bson.M{"$lte": o.BankRoutingNumberLte}
		and = append(and, query)
	}
	if o.BankRoutingNumberIn != nil {
		query := bson.M{}
		query["bankRoutingNumber"] = bson.M{"$in": o.BankRoutingNumberIn}
		and = append(and, query)
	}
	if o.BankRoutingNumberNin != nil {
		query := bson.M{}
		query["bankRoutingNumber"] = bson.M{"$nin": o.BankRoutingNumberNin}
		and = append(and, query)
	}
	if o.BankRoutingNumberExists != nil {
		query := bson.M{}
		query["bankRoutingNumber"] = bson.M{"$exists": *o.BankRoutingNumberExists}
		and = append(and, query)
	}
	if o.BankRoutingNumberLike != nil {
		query := bson.M{}
		query["bankRoutingNumber"] = bson.M{"$regex": o.BankRoutingNumberLike, "$options": "i"}
		and = append(and, query)
	}
	if o.BankRoutingNumberNlike != nil {
		query := bson.M{}
		query["bankRoutingNumber"] = bson.M{"$not": bson.M{"$regex": o.BankRoutingNumberNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.BankSwiftEq != nil {
		query := bson.M{}
		query["bankSwift"] = o.BankSwiftEq
		and = append(and, query)
	}
	if o.BankSwiftNe != nil {
		query := bson.M{}
		query["bankSwift"] = bson.M{"$ne": o.BankSwiftNe}
		and = append(and, query)
	}
	if o.BankSwiftGt != nil {
		query := bson.M{}
		query["bankSwift"] = bson.M{"$gt": o.BankSwiftGt}
		and = append(and, query)
	}
	if o.BankSwiftGte != nil {
		query := bson.M{}
		query["bankSwift"] = bson.M{"$gte": o.BankSwiftGte}
		and = append(and, query)
	}
	if o.BankSwiftLt != nil {
		query := bson.M{}
		query["bankSwift"] = bson.M{"$lt": o.BankSwiftLt}
		and = append(and, query)
	}
	if o.BankSwiftLte != nil {
		query := bson.M{}
		query["bankSwift"] = bson.M{"$lte": o.BankSwiftLte}
		and = append(and, query)
	}
	if o.BankSwiftIn != nil {
		query := bson.M{}
		query["bankSwift"] = bson.M{"$in": o.BankSwiftIn}
		and = append(and, query)
	}
	if o.BankSwiftNin != nil {
		query := bson.M{}
		query["bankSwift"] = bson.M{"$nin": o.BankSwiftNin}
		and = append(and, query)
	}
	if o.BankSwiftExists != nil {
		query := bson.M{}
		query["bankSwift"] = bson.M{"$exists": *o.BankSwiftExists}
		and = append(and, query)
	}
	if o.BankSwiftLike != nil {
		query := bson.M{}
		query["bankSwift"] = bson.M{"$regex": o.BankSwiftLike, "$options": "i"}
		and = append(and, query)
	}
	if o.BankSwiftNlike != nil {
		query := bson.M{}
		query["bankSwift"] = bson.M{"$not": bson.M{"$regex": o.BankSwiftNlike, "$options": "i"}}
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
	if o.DisplayNameEq != nil {
		query := bson.M{}
		query["displayName"] = o.DisplayNameEq
		and = append(and, query)
	}
	if o.DisplayNameNe != nil {
		query := bson.M{}
		query["displayName"] = bson.M{"$ne": o.DisplayNameNe}
		and = append(and, query)
	}
	if o.DisplayNameGt != nil {
		query := bson.M{}
		query["displayName"] = bson.M{"$gt": o.DisplayNameGt}
		and = append(and, query)
	}
	if o.DisplayNameGte != nil {
		query := bson.M{}
		query["displayName"] = bson.M{"$gte": o.DisplayNameGte}
		and = append(and, query)
	}
	if o.DisplayNameLt != nil {
		query := bson.M{}
		query["displayName"] = bson.M{"$lt": o.DisplayNameLt}
		and = append(and, query)
	}
	if o.DisplayNameLte != nil {
		query := bson.M{}
		query["displayName"] = bson.M{"$lte": o.DisplayNameLte}
		and = append(and, query)
	}
	if o.DisplayNameIn != nil {
		query := bson.M{}
		query["displayName"] = bson.M{"$in": o.DisplayNameIn}
		and = append(and, query)
	}
	if o.DisplayNameNin != nil {
		query := bson.M{}
		query["displayName"] = bson.M{"$nin": o.DisplayNameNin}
		and = append(and, query)
	}
	if o.DisplayNameExists != nil {
		query := bson.M{}
		query["displayName"] = bson.M{"$exists": *o.DisplayNameExists}
		and = append(and, query)
	}
	if o.DisplayNameLike != nil {
		query := bson.M{}
		query["displayName"] = bson.M{"$regex": o.DisplayNameLike, "$options": "i"}
		and = append(and, query)
	}
	if o.DisplayNameNlike != nil {
		query := bson.M{}
		query["displayName"] = bson.M{"$not": bson.M{"$regex": o.DisplayNameNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.EnabledEq != nil {
		query := bson.M{}
		query["enabled"] = o.EnabledEq
		and = append(and, query)
	}
	if o.EnabledNe != nil {
		query := bson.M{}
		query["enabled"] = bson.M{"$ne": o.EnabledNe}
		and = append(and, query)
	}
	if o.EnabledGt != nil {
		query := bson.M{}
		query["enabled"] = bson.M{"$gt": o.EnabledGt}
		and = append(and, query)
	}
	if o.EnabledGte != nil {
		query := bson.M{}
		query["enabled"] = bson.M{"$gte": o.EnabledGte}
		and = append(and, query)
	}
	if o.EnabledLt != nil {
		query := bson.M{}
		query["enabled"] = bson.M{"$lt": o.EnabledLt}
		and = append(and, query)
	}
	if o.EnabledLte != nil {
		query := bson.M{}
		query["enabled"] = bson.M{"$lte": o.EnabledLte}
		and = append(and, query)
	}
	if o.EnabledIn != nil {
		query := bson.M{}
		query["enabled"] = bson.M{"$in": o.EnabledIn}
		and = append(and, query)
	}
	if o.EnabledNin != nil {
		query := bson.M{}
		query["enabled"] = bson.M{"$nin": o.EnabledNin}
		and = append(and, query)
	}
	if o.EnabledExists != nil {
		query := bson.M{}
		query["enabled"] = bson.M{"$exists": *o.EnabledExists}
		and = append(and, query)
	}
	if o.InstructionsEq != nil {
		query := bson.M{}
		query["instructions"] = o.InstructionsEq
		and = append(and, query)
	}
	if o.InstructionsNe != nil {
		query := bson.M{}
		query["instructions"] = bson.M{"$ne": o.InstructionsNe}
		and = append(and, query)
	}
	if o.InstructionsGt != nil {
		query := bson.M{}
		query["instructions"] = bson.M{"$gt": o.InstructionsGt}
		and = append(and, query)
	}
	if o.InstructionsGte != nil {
		query := bson.M{}
		query["instructions"] = bson.M{"$gte": o.InstructionsGte}
		and = append(and, query)
	}
	if o.InstructionsLt != nil {
		query := bson.M{}
		query["instructions"] = bson.M{"$lt": o.InstructionsLt}
		and = append(and, query)
	}
	if o.InstructionsLte != nil {
		query := bson.M{}
		query["instructions"] = bson.M{"$lte": o.InstructionsLte}
		and = append(and, query)
	}
	if o.InstructionsIn != nil {
		query := bson.M{}
		query["instructions"] = bson.M{"$in": o.InstructionsIn}
		and = append(and, query)
	}
	if o.InstructionsNin != nil {
		query := bson.M{}
		query["instructions"] = bson.M{"$nin": o.InstructionsNin}
		and = append(and, query)
	}
	if o.InstructionsExists != nil {
		query := bson.M{}
		query["instructions"] = bson.M{"$exists": *o.InstructionsExists}
		and = append(and, query)
	}
	if o.InstructionsLike != nil {
		query := bson.M{}
		query["instructions"] = bson.M{"$regex": o.InstructionsLike, "$options": "i"}
		and = append(and, query)
	}
	if o.InstructionsNlike != nil {
		query := bson.M{}
		query["instructions"] = bson.M{"$not": bson.M{"$regex": o.InstructionsNlike, "$options": "i"}}
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
	if o.PaymentUrlEq != nil {
		query := bson.M{}
		query["paymentUrl"] = o.PaymentUrlEq
		and = append(and, query)
	}
	if o.PaymentUrlNe != nil {
		query := bson.M{}
		query["paymentUrl"] = bson.M{"$ne": o.PaymentUrlNe}
		and = append(and, query)
	}
	if o.PaymentUrlGt != nil {
		query := bson.M{}
		query["paymentUrl"] = bson.M{"$gt": o.PaymentUrlGt}
		and = append(and, query)
	}
	if o.PaymentUrlGte != nil {
		query := bson.M{}
		query["paymentUrl"] = bson.M{"$gte": o.PaymentUrlGte}
		and = append(and, query)
	}
	if o.PaymentUrlLt != nil {
		query := bson.M{}
		query["paymentUrl"] = bson.M{"$lt": o.PaymentUrlLt}
		and = append(and, query)
	}
	if o.PaymentUrlLte != nil {
		query := bson.M{}
		query["paymentUrl"] = bson.M{"$lte": o.PaymentUrlLte}
		and = append(and, query)
	}
	if o.PaymentUrlIn != nil {
		query := bson.M{}
		query["paymentUrl"] = bson.M{"$in": o.PaymentUrlIn}
		and = append(and, query)
	}
	if o.PaymentUrlNin != nil {
		query := bson.M{}
		query["paymentUrl"] = bson.M{"$nin": o.PaymentUrlNin}
		and = append(and, query)
	}
	if o.PaymentUrlExists != nil {
		query := bson.M{}
		query["paymentUrl"] = bson.M{"$exists": *o.PaymentUrlExists}
		and = append(and, query)
	}
	if o.PaymentUrlLike != nil {
		query := bson.M{}
		query["paymentUrl"] = bson.M{"$regex": o.PaymentUrlLike, "$options": "i"}
		and = append(and, query)
	}
	if o.PaymentUrlNlike != nil {
		query := bson.M{}
		query["paymentUrl"] = bson.M{"$not": bson.M{"$regex": o.PaymentUrlNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.PositionEq != nil {
		query := bson.M{}
		query["position"] = o.PositionEq
		and = append(and, query)
	}
	if o.PositionNe != nil {
		query := bson.M{}
		query["position"] = bson.M{"$ne": o.PositionNe}
		and = append(and, query)
	}
	if o.PositionGt != nil {
		query := bson.M{}
		query["position"] = bson.M{"$gt": o.PositionGt}
		and = append(and, query)
	}
	if o.PositionGte != nil {
		query := bson.M{}
		query["position"] = bson.M{"$gte": o.PositionGte}
		and = append(and, query)
	}
	if o.PositionLt != nil {
		query := bson.M{}
		query["position"] = bson.M{"$lt": o.PositionLt}
		and = append(and, query)
	}
	if o.PositionLte != nil {
		query := bson.M{}
		query["position"] = bson.M{"$lte": o.PositionLte}
		and = append(and, query)
	}
	if o.PositionIn != nil {
		query := bson.M{}
		query["position"] = bson.M{"$in": o.PositionIn}
		and = append(and, query)
	}
	if o.PositionNin != nil {
		query := bson.M{}
		query["position"] = bson.M{"$nin": o.PositionNin}
		and = append(and, query)
	}
	if o.PositionExists != nil {
		query := bson.M{}
		query["position"] = bson.M{"$exists": *o.PositionExists}
		and = append(and, query)
	}
	if o.RecipientEmailEq != nil {
		query := bson.M{}
		query["recipientEmail"] = o.RecipientEmailEq
		and = append(and, query)
	}
	if o.RecipientEmailNe != nil {
		query := bson.M{}
		query["recipientEmail"] = bson.M{"$ne": o.RecipientEmailNe}
		and = append(and, query)
	}
	if o.RecipientEmailGt != nil {
		query := bson.M{}
		query["recipientEmail"] = bson.M{"$gt": o.RecipientEmailGt}
		and = append(and, query)
	}
	if o.RecipientEmailGte != nil {
		query := bson.M{}
		query["recipientEmail"] = bson.M{"$gte": o.RecipientEmailGte}
		and = append(and, query)
	}
	if o.RecipientEmailLt != nil {
		query := bson.M{}
		query["recipientEmail"] = bson.M{"$lt": o.RecipientEmailLt}
		and = append(and, query)
	}
	if o.RecipientEmailLte != nil {
		query := bson.M{}
		query["recipientEmail"] = bson.M{"$lte": o.RecipientEmailLte}
		and = append(and, query)
	}
	if o.RecipientEmailIn != nil {
		query := bson.M{}
		query["recipientEmail"] = bson.M{"$in": o.RecipientEmailIn}
		and = append(and, query)
	}
	if o.RecipientEmailNin != nil {
		query := bson.M{}
		query["recipientEmail"] = bson.M{"$nin": o.RecipientEmailNin}
		and = append(and, query)
	}
	if o.RecipientEmailExists != nil {
		query := bson.M{}
		query["recipientEmail"] = bson.M{"$exists": *o.RecipientEmailExists}
		and = append(and, query)
	}
	if o.RecipientEmailLike != nil {
		query := bson.M{}
		query["recipientEmail"] = bson.M{"$regex": o.RecipientEmailLike, "$options": "i"}
		and = append(and, query)
	}
	if o.RecipientEmailNlike != nil {
		query := bson.M{}
		query["recipientEmail"] = bson.M{"$not": bson.M{"$regex": o.RecipientEmailNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.RecipientPhoneEq != nil {
		query := bson.M{}
		query["recipientPhone"] = o.RecipientPhoneEq
		and = append(and, query)
	}
	if o.RecipientPhoneNe != nil {
		query := bson.M{}
		query["recipientPhone"] = bson.M{"$ne": o.RecipientPhoneNe}
		and = append(and, query)
	}
	if o.RecipientPhoneGt != nil {
		query := bson.M{}
		query["recipientPhone"] = bson.M{"$gt": o.RecipientPhoneGt}
		and = append(and, query)
	}
	if o.RecipientPhoneGte != nil {
		query := bson.M{}
		query["recipientPhone"] = bson.M{"$gte": o.RecipientPhoneGte}
		and = append(and, query)
	}
	if o.RecipientPhoneLt != nil {
		query := bson.M{}
		query["recipientPhone"] = bson.M{"$lt": o.RecipientPhoneLt}
		and = append(and, query)
	}
	if o.RecipientPhoneLte != nil {
		query := bson.M{}
		query["recipientPhone"] = bson.M{"$lte": o.RecipientPhoneLte}
		and = append(and, query)
	}
	if o.RecipientPhoneIn != nil {
		query := bson.M{}
		query["recipientPhone"] = bson.M{"$in": o.RecipientPhoneIn}
		and = append(and, query)
	}
	if o.RecipientPhoneNin != nil {
		query := bson.M{}
		query["recipientPhone"] = bson.M{"$nin": o.RecipientPhoneNin}
		and = append(and, query)
	}
	if o.RecipientPhoneExists != nil {
		query := bson.M{}
		query["recipientPhone"] = bson.M{"$exists": *o.RecipientPhoneExists}
		and = append(and, query)
	}
	if o.RecipientPhoneLike != nil {
		query := bson.M{}
		query["recipientPhone"] = bson.M{"$regex": o.RecipientPhoneLike, "$options": "i"}
		and = append(and, query)
	}
	if o.RecipientPhoneNlike != nil {
		query := bson.M{}
		query["recipientPhone"] = bson.M{"$not": bson.M{"$regex": o.RecipientPhoneNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.RegistryIdEq != nil {
		query := bson.M{}
		query["registryId"] = o.RegistryIdEq
		and = append(and, query)
	}
	if o.RegistryIdIn != nil {
		query := bson.M{}
		query["registryId"] = bson.M{"$in": o.RegistryIdIn}
		and = append(and, query)
	}
	if o.RegistryIdNin != nil {
		query := bson.M{}
		query["registryId"] = bson.M{"$nin": o.RegistryIdNin}
		and = append(and, query)
	}
	if o.RegistryIdExists != nil {
		query := bson.M{}
		query["registryId"] = bson.M{"$exists": *o.RegistryIdExists}
		and = append(and, query)
	}
	if o.TypeEq != nil {
		query := bson.M{}
		query["type"] = o.TypeEq
		and = append(and, query)
	}
	if o.TypeNe != nil {
		query := bson.M{}
		query["type"] = bson.M{"$ne": o.TypeNe}
		and = append(and, query)
	}
	if o.TypeGt != nil {
		query := bson.M{}
		query["type"] = bson.M{"$gt": o.TypeGt}
		and = append(and, query)
	}
	if o.TypeGte != nil {
		query := bson.M{}
		query["type"] = bson.M{"$gte": o.TypeGte}
		and = append(and, query)
	}
	if o.TypeLt != nil {
		query := bson.M{}
		query["type"] = bson.M{"$lt": o.TypeLt}
		and = append(and, query)
	}
	if o.TypeLte != nil {
		query := bson.M{}
		query["type"] = bson.M{"$lte": o.TypeLte}
		and = append(and, query)
	}
	if o.TypeIn != nil {
		query := bson.M{}
		query["type"] = bson.M{"$in": o.TypeIn}
		and = append(and, query)
	}
	if o.TypeNin != nil {
		query := bson.M{}
		query["type"] = bson.M{"$nin": o.TypeNin}
		and = append(and, query)
	}
	if o.TypeExists != nil {
		query := bson.M{}
		query["type"] = bson.M{"$exists": *o.TypeExists}
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
	return and, nil
}

type MongoSortParams struct {
	CreatedAt  int8
	OwnerId    int8
	RegistryId int8
	UpdatedAt  int8
}
