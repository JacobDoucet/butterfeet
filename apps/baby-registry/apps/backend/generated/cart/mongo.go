package cart

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_cart_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_payment_method_type"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoRecord struct {
	Id                 *primitive.ObjectID             `bson:"_id,omitempty"`
	AmountCents        *int                            `bson:"amountCents,omitempty"`
	ClaimedAt          *time.Time                      `bson:"claimedAt,omitempty"`
	ContributorEmail   *string                         `bson:"contributorEmail,omitempty"`
	ContributorName    *string                         `bson:"contributorName,omitempty"`
	Created            *actor_trace.MongoRecord        `bson:"created,omitempty"`
	Currency           *string                         `bson:"currency,omitempty"`
	DecidedAt          *time.Time                      `bson:"decidedAt,omitempty"`
	DecisionReason     *string                         `bson:"decisionReason,omitempty"`
	Message            *string                         `bson:"message,omitempty"`
	MethodDisplayName  *string                         `bson:"methodDisplayName,omitempty"`
	MethodType         *enum_payment_method_type.Value `bson:"methodType,omitempty"`
	OwnerId            *primitive.ObjectID             `bson:"ownerId,omitempty"`
	PaymentMethodId    *primitive.ObjectID             `bson:"paymentMethodId,omitempty"`
	ReferenceCode      *string                         `bson:"referenceCode,omitempty"`
	RegistryId         *primitive.ObjectID             `bson:"registryId,omitempty"`
	Status             *enum_cart_status.Value         `bson:"status,omitempty"`
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
		elemownerId0 := r.OwnerId.Hex()
		m.OwnerId = elemownerId0
	}
	if r.PaymentMethodId != nil {
		elempaymentMethodId0 := r.PaymentMethodId.Hex()
		m.PaymentMethodId = elempaymentMethodId0
	}
	if r.ReferenceCode != nil {
		elemreferenceCode0 := r.ReferenceCode
		m.ReferenceCode = *elemreferenceCode0
	}
	if r.RegistryId != nil {
		elemregistryId0 := r.RegistryId.Hex()
		m.RegistryId = elemregistryId0
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

type MongoSelectByIdQuery struct {
	Id primitive.ObjectID
}
type MongoSelectByReferenceUniqueQuery struct {
	ReferenceCode string
}

type MongoWhereClause struct {
	// id (Ref<Cart>) search options
	IdEq     *primitive.ObjectID
	IdIn     *[]primitive.ObjectID
	IdNin    *[]primitive.ObjectID
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
	Created *actor_trace.MongoWhereClause
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
	OwnerIdEq     *primitive.ObjectID
	OwnerIdIn     *[]primitive.ObjectID
	OwnerIdNin    *[]primitive.ObjectID
	OwnerIdExists *bool
	// paymentMethodId (Ref<RegistryPaymentMethod>) search options
	PaymentMethodIdEq     *primitive.ObjectID
	PaymentMethodIdIn     *[]primitive.ObjectID
	PaymentMethodIdNin    *[]primitive.ObjectID
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
	RegistryIdEq     *primitive.ObjectID
	RegistryIdIn     *[]primitive.ObjectID
	RegistryIdNin    *[]primitive.ObjectID
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
	if o.AmountCentsEq != nil {
		query := bson.M{}
		query["amountCents"] = o.AmountCentsEq
		and = append(and, query)
	}
	if o.AmountCentsNe != nil {
		query := bson.M{}
		query["amountCents"] = bson.M{"$ne": o.AmountCentsNe}
		and = append(and, query)
	}
	if o.AmountCentsGt != nil {
		query := bson.M{}
		query["amountCents"] = bson.M{"$gt": o.AmountCentsGt}
		and = append(and, query)
	}
	if o.AmountCentsGte != nil {
		query := bson.M{}
		query["amountCents"] = bson.M{"$gte": o.AmountCentsGte}
		and = append(and, query)
	}
	if o.AmountCentsLt != nil {
		query := bson.M{}
		query["amountCents"] = bson.M{"$lt": o.AmountCentsLt}
		and = append(and, query)
	}
	if o.AmountCentsLte != nil {
		query := bson.M{}
		query["amountCents"] = bson.M{"$lte": o.AmountCentsLte}
		and = append(and, query)
	}
	if o.AmountCentsIn != nil {
		query := bson.M{}
		query["amountCents"] = bson.M{"$in": o.AmountCentsIn}
		and = append(and, query)
	}
	if o.AmountCentsNin != nil {
		query := bson.M{}
		query["amountCents"] = bson.M{"$nin": o.AmountCentsNin}
		and = append(and, query)
	}
	if o.AmountCentsExists != nil {
		query := bson.M{}
		query["amountCents"] = bson.M{"$exists": *o.AmountCentsExists}
		and = append(and, query)
	}
	if o.ClaimedAtEq != nil {
		query := bson.M{}
		query["claimedAt"] = o.ClaimedAtEq
		and = append(and, query)
	}
	if o.ClaimedAtNe != nil {
		query := bson.M{}
		query["claimedAt"] = bson.M{"$ne": o.ClaimedAtNe}
		and = append(and, query)
	}
	if o.ClaimedAtGt != nil {
		query := bson.M{}
		query["claimedAt"] = bson.M{"$gt": o.ClaimedAtGt}
		and = append(and, query)
	}
	if o.ClaimedAtGte != nil {
		query := bson.M{}
		query["claimedAt"] = bson.M{"$gte": o.ClaimedAtGte}
		and = append(and, query)
	}
	if o.ClaimedAtLt != nil {
		query := bson.M{}
		query["claimedAt"] = bson.M{"$lt": o.ClaimedAtLt}
		and = append(and, query)
	}
	if o.ClaimedAtLte != nil {
		query := bson.M{}
		query["claimedAt"] = bson.M{"$lte": o.ClaimedAtLte}
		and = append(and, query)
	}
	if o.ClaimedAtIn != nil {
		query := bson.M{}
		query["claimedAt"] = bson.M{"$in": o.ClaimedAtIn}
		and = append(and, query)
	}
	if o.ClaimedAtNin != nil {
		query := bson.M{}
		query["claimedAt"] = bson.M{"$nin": o.ClaimedAtNin}
		and = append(and, query)
	}
	if o.ClaimedAtExists != nil {
		query := bson.M{}
		query["claimedAt"] = bson.M{"$exists": *o.ClaimedAtExists}
		and = append(and, query)
	}
	if o.ContributorEmailEq != nil {
		query := bson.M{}
		query["contributorEmail"] = o.ContributorEmailEq
		and = append(and, query)
	}
	if o.ContributorEmailNe != nil {
		query := bson.M{}
		query["contributorEmail"] = bson.M{"$ne": o.ContributorEmailNe}
		and = append(and, query)
	}
	if o.ContributorEmailGt != nil {
		query := bson.M{}
		query["contributorEmail"] = bson.M{"$gt": o.ContributorEmailGt}
		and = append(and, query)
	}
	if o.ContributorEmailGte != nil {
		query := bson.M{}
		query["contributorEmail"] = bson.M{"$gte": o.ContributorEmailGte}
		and = append(and, query)
	}
	if o.ContributorEmailLt != nil {
		query := bson.M{}
		query["contributorEmail"] = bson.M{"$lt": o.ContributorEmailLt}
		and = append(and, query)
	}
	if o.ContributorEmailLte != nil {
		query := bson.M{}
		query["contributorEmail"] = bson.M{"$lte": o.ContributorEmailLte}
		and = append(and, query)
	}
	if o.ContributorEmailIn != nil {
		query := bson.M{}
		query["contributorEmail"] = bson.M{"$in": o.ContributorEmailIn}
		and = append(and, query)
	}
	if o.ContributorEmailNin != nil {
		query := bson.M{}
		query["contributorEmail"] = bson.M{"$nin": o.ContributorEmailNin}
		and = append(and, query)
	}
	if o.ContributorEmailExists != nil {
		query := bson.M{}
		query["contributorEmail"] = bson.M{"$exists": *o.ContributorEmailExists}
		and = append(and, query)
	}
	if o.ContributorEmailLike != nil {
		query := bson.M{}
		query["contributorEmail"] = bson.M{"$regex": o.ContributorEmailLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ContributorEmailNlike != nil {
		query := bson.M{}
		query["contributorEmail"] = bson.M{"$not": bson.M{"$regex": o.ContributorEmailNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.ContributorNameEq != nil {
		query := bson.M{}
		query["contributorName"] = o.ContributorNameEq
		and = append(and, query)
	}
	if o.ContributorNameNe != nil {
		query := bson.M{}
		query["contributorName"] = bson.M{"$ne": o.ContributorNameNe}
		and = append(and, query)
	}
	if o.ContributorNameGt != nil {
		query := bson.M{}
		query["contributorName"] = bson.M{"$gt": o.ContributorNameGt}
		and = append(and, query)
	}
	if o.ContributorNameGte != nil {
		query := bson.M{}
		query["contributorName"] = bson.M{"$gte": o.ContributorNameGte}
		and = append(and, query)
	}
	if o.ContributorNameLt != nil {
		query := bson.M{}
		query["contributorName"] = bson.M{"$lt": o.ContributorNameLt}
		and = append(and, query)
	}
	if o.ContributorNameLte != nil {
		query := bson.M{}
		query["contributorName"] = bson.M{"$lte": o.ContributorNameLte}
		and = append(and, query)
	}
	if o.ContributorNameIn != nil {
		query := bson.M{}
		query["contributorName"] = bson.M{"$in": o.ContributorNameIn}
		and = append(and, query)
	}
	if o.ContributorNameNin != nil {
		query := bson.M{}
		query["contributorName"] = bson.M{"$nin": o.ContributorNameNin}
		and = append(and, query)
	}
	if o.ContributorNameExists != nil {
		query := bson.M{}
		query["contributorName"] = bson.M{"$exists": *o.ContributorNameExists}
		and = append(and, query)
	}
	if o.ContributorNameLike != nil {
		query := bson.M{}
		query["contributorName"] = bson.M{"$regex": o.ContributorNameLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ContributorNameNlike != nil {
		query := bson.M{}
		query["contributorName"] = bson.M{"$not": bson.M{"$regex": o.ContributorNameNlike, "$options": "i"}}
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
	if o.CurrencyEq != nil {
		query := bson.M{}
		query["currency"] = o.CurrencyEq
		and = append(and, query)
	}
	if o.CurrencyNe != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$ne": o.CurrencyNe}
		and = append(and, query)
	}
	if o.CurrencyGt != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$gt": o.CurrencyGt}
		and = append(and, query)
	}
	if o.CurrencyGte != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$gte": o.CurrencyGte}
		and = append(and, query)
	}
	if o.CurrencyLt != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$lt": o.CurrencyLt}
		and = append(and, query)
	}
	if o.CurrencyLte != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$lte": o.CurrencyLte}
		and = append(and, query)
	}
	if o.CurrencyIn != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$in": o.CurrencyIn}
		and = append(and, query)
	}
	if o.CurrencyNin != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$nin": o.CurrencyNin}
		and = append(and, query)
	}
	if o.CurrencyExists != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$exists": *o.CurrencyExists}
		and = append(and, query)
	}
	if o.CurrencyLike != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$regex": o.CurrencyLike, "$options": "i"}
		and = append(and, query)
	}
	if o.CurrencyNlike != nil {
		query := bson.M{}
		query["currency"] = bson.M{"$not": bson.M{"$regex": o.CurrencyNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.DecidedAtEq != nil {
		query := bson.M{}
		query["decidedAt"] = o.DecidedAtEq
		and = append(and, query)
	}
	if o.DecidedAtNe != nil {
		query := bson.M{}
		query["decidedAt"] = bson.M{"$ne": o.DecidedAtNe}
		and = append(and, query)
	}
	if o.DecidedAtGt != nil {
		query := bson.M{}
		query["decidedAt"] = bson.M{"$gt": o.DecidedAtGt}
		and = append(and, query)
	}
	if o.DecidedAtGte != nil {
		query := bson.M{}
		query["decidedAt"] = bson.M{"$gte": o.DecidedAtGte}
		and = append(and, query)
	}
	if o.DecidedAtLt != nil {
		query := bson.M{}
		query["decidedAt"] = bson.M{"$lt": o.DecidedAtLt}
		and = append(and, query)
	}
	if o.DecidedAtLte != nil {
		query := bson.M{}
		query["decidedAt"] = bson.M{"$lte": o.DecidedAtLte}
		and = append(and, query)
	}
	if o.DecidedAtIn != nil {
		query := bson.M{}
		query["decidedAt"] = bson.M{"$in": o.DecidedAtIn}
		and = append(and, query)
	}
	if o.DecidedAtNin != nil {
		query := bson.M{}
		query["decidedAt"] = bson.M{"$nin": o.DecidedAtNin}
		and = append(and, query)
	}
	if o.DecidedAtExists != nil {
		query := bson.M{}
		query["decidedAt"] = bson.M{"$exists": *o.DecidedAtExists}
		and = append(and, query)
	}
	if o.DecisionReasonEq != nil {
		query := bson.M{}
		query["decisionReason"] = o.DecisionReasonEq
		and = append(and, query)
	}
	if o.DecisionReasonNe != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$ne": o.DecisionReasonNe}
		and = append(and, query)
	}
	if o.DecisionReasonGt != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$gt": o.DecisionReasonGt}
		and = append(and, query)
	}
	if o.DecisionReasonGte != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$gte": o.DecisionReasonGte}
		and = append(and, query)
	}
	if o.DecisionReasonLt != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$lt": o.DecisionReasonLt}
		and = append(and, query)
	}
	if o.DecisionReasonLte != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$lte": o.DecisionReasonLte}
		and = append(and, query)
	}
	if o.DecisionReasonIn != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$in": o.DecisionReasonIn}
		and = append(and, query)
	}
	if o.DecisionReasonNin != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$nin": o.DecisionReasonNin}
		and = append(and, query)
	}
	if o.DecisionReasonExists != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$exists": *o.DecisionReasonExists}
		and = append(and, query)
	}
	if o.DecisionReasonLike != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$regex": o.DecisionReasonLike, "$options": "i"}
		and = append(and, query)
	}
	if o.DecisionReasonNlike != nil {
		query := bson.M{}
		query["decisionReason"] = bson.M{"$not": bson.M{"$regex": o.DecisionReasonNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.MessageEq != nil {
		query := bson.M{}
		query["message"] = o.MessageEq
		and = append(and, query)
	}
	if o.MessageNe != nil {
		query := bson.M{}
		query["message"] = bson.M{"$ne": o.MessageNe}
		and = append(and, query)
	}
	if o.MessageGt != nil {
		query := bson.M{}
		query["message"] = bson.M{"$gt": o.MessageGt}
		and = append(and, query)
	}
	if o.MessageGte != nil {
		query := bson.M{}
		query["message"] = bson.M{"$gte": o.MessageGte}
		and = append(and, query)
	}
	if o.MessageLt != nil {
		query := bson.M{}
		query["message"] = bson.M{"$lt": o.MessageLt}
		and = append(and, query)
	}
	if o.MessageLte != nil {
		query := bson.M{}
		query["message"] = bson.M{"$lte": o.MessageLte}
		and = append(and, query)
	}
	if o.MessageIn != nil {
		query := bson.M{}
		query["message"] = bson.M{"$in": o.MessageIn}
		and = append(and, query)
	}
	if o.MessageNin != nil {
		query := bson.M{}
		query["message"] = bson.M{"$nin": o.MessageNin}
		and = append(and, query)
	}
	if o.MessageExists != nil {
		query := bson.M{}
		query["message"] = bson.M{"$exists": *o.MessageExists}
		and = append(and, query)
	}
	if o.MessageLike != nil {
		query := bson.M{}
		query["message"] = bson.M{"$regex": o.MessageLike, "$options": "i"}
		and = append(and, query)
	}
	if o.MessageNlike != nil {
		query := bson.M{}
		query["message"] = bson.M{"$not": bson.M{"$regex": o.MessageNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.MethodDisplayNameEq != nil {
		query := bson.M{}
		query["methodDisplayName"] = o.MethodDisplayNameEq
		and = append(and, query)
	}
	if o.MethodDisplayNameNe != nil {
		query := bson.M{}
		query["methodDisplayName"] = bson.M{"$ne": o.MethodDisplayNameNe}
		and = append(and, query)
	}
	if o.MethodDisplayNameGt != nil {
		query := bson.M{}
		query["methodDisplayName"] = bson.M{"$gt": o.MethodDisplayNameGt}
		and = append(and, query)
	}
	if o.MethodDisplayNameGte != nil {
		query := bson.M{}
		query["methodDisplayName"] = bson.M{"$gte": o.MethodDisplayNameGte}
		and = append(and, query)
	}
	if o.MethodDisplayNameLt != nil {
		query := bson.M{}
		query["methodDisplayName"] = bson.M{"$lt": o.MethodDisplayNameLt}
		and = append(and, query)
	}
	if o.MethodDisplayNameLte != nil {
		query := bson.M{}
		query["methodDisplayName"] = bson.M{"$lte": o.MethodDisplayNameLte}
		and = append(and, query)
	}
	if o.MethodDisplayNameIn != nil {
		query := bson.M{}
		query["methodDisplayName"] = bson.M{"$in": o.MethodDisplayNameIn}
		and = append(and, query)
	}
	if o.MethodDisplayNameNin != nil {
		query := bson.M{}
		query["methodDisplayName"] = bson.M{"$nin": o.MethodDisplayNameNin}
		and = append(and, query)
	}
	if o.MethodDisplayNameExists != nil {
		query := bson.M{}
		query["methodDisplayName"] = bson.M{"$exists": *o.MethodDisplayNameExists}
		and = append(and, query)
	}
	if o.MethodDisplayNameLike != nil {
		query := bson.M{}
		query["methodDisplayName"] = bson.M{"$regex": o.MethodDisplayNameLike, "$options": "i"}
		and = append(and, query)
	}
	if o.MethodDisplayNameNlike != nil {
		query := bson.M{}
		query["methodDisplayName"] = bson.M{"$not": bson.M{"$regex": o.MethodDisplayNameNlike, "$options": "i"}}
		and = append(and, query)
	}
	if o.MethodTypeEq != nil {
		query := bson.M{}
		query["methodType"] = o.MethodTypeEq
		and = append(and, query)
	}
	if o.MethodTypeNe != nil {
		query := bson.M{}
		query["methodType"] = bson.M{"$ne": o.MethodTypeNe}
		and = append(and, query)
	}
	if o.MethodTypeGt != nil {
		query := bson.M{}
		query["methodType"] = bson.M{"$gt": o.MethodTypeGt}
		and = append(and, query)
	}
	if o.MethodTypeGte != nil {
		query := bson.M{}
		query["methodType"] = bson.M{"$gte": o.MethodTypeGte}
		and = append(and, query)
	}
	if o.MethodTypeLt != nil {
		query := bson.M{}
		query["methodType"] = bson.M{"$lt": o.MethodTypeLt}
		and = append(and, query)
	}
	if o.MethodTypeLte != nil {
		query := bson.M{}
		query["methodType"] = bson.M{"$lte": o.MethodTypeLte}
		and = append(and, query)
	}
	if o.MethodTypeIn != nil {
		query := bson.M{}
		query["methodType"] = bson.M{"$in": o.MethodTypeIn}
		and = append(and, query)
	}
	if o.MethodTypeNin != nil {
		query := bson.M{}
		query["methodType"] = bson.M{"$nin": o.MethodTypeNin}
		and = append(and, query)
	}
	if o.MethodTypeExists != nil {
		query := bson.M{}
		query["methodType"] = bson.M{"$exists": *o.MethodTypeExists}
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
	if o.PaymentMethodIdEq != nil {
		query := bson.M{}
		query["paymentMethodId"] = o.PaymentMethodIdEq
		and = append(and, query)
	}
	if o.PaymentMethodIdIn != nil {
		query := bson.M{}
		query["paymentMethodId"] = bson.M{"$in": o.PaymentMethodIdIn}
		and = append(and, query)
	}
	if o.PaymentMethodIdNin != nil {
		query := bson.M{}
		query["paymentMethodId"] = bson.M{"$nin": o.PaymentMethodIdNin}
		and = append(and, query)
	}
	if o.PaymentMethodIdExists != nil {
		query := bson.M{}
		query["paymentMethodId"] = bson.M{"$exists": *o.PaymentMethodIdExists}
		and = append(and, query)
	}
	if o.ReferenceCodeEq != nil {
		query := bson.M{}
		query["referenceCode"] = o.ReferenceCodeEq
		and = append(and, query)
	}
	if o.ReferenceCodeNe != nil {
		query := bson.M{}
		query["referenceCode"] = bson.M{"$ne": o.ReferenceCodeNe}
		and = append(and, query)
	}
	if o.ReferenceCodeGt != nil {
		query := bson.M{}
		query["referenceCode"] = bson.M{"$gt": o.ReferenceCodeGt}
		and = append(and, query)
	}
	if o.ReferenceCodeGte != nil {
		query := bson.M{}
		query["referenceCode"] = bson.M{"$gte": o.ReferenceCodeGte}
		and = append(and, query)
	}
	if o.ReferenceCodeLt != nil {
		query := bson.M{}
		query["referenceCode"] = bson.M{"$lt": o.ReferenceCodeLt}
		and = append(and, query)
	}
	if o.ReferenceCodeLte != nil {
		query := bson.M{}
		query["referenceCode"] = bson.M{"$lte": o.ReferenceCodeLte}
		and = append(and, query)
	}
	if o.ReferenceCodeIn != nil {
		query := bson.M{}
		query["referenceCode"] = bson.M{"$in": o.ReferenceCodeIn}
		and = append(and, query)
	}
	if o.ReferenceCodeNin != nil {
		query := bson.M{}
		query["referenceCode"] = bson.M{"$nin": o.ReferenceCodeNin}
		and = append(and, query)
	}
	if o.ReferenceCodeExists != nil {
		query := bson.M{}
		query["referenceCode"] = bson.M{"$exists": *o.ReferenceCodeExists}
		and = append(and, query)
	}
	if o.ReferenceCodeLike != nil {
		query := bson.M{}
		query["referenceCode"] = bson.M{"$regex": o.ReferenceCodeLike, "$options": "i"}
		and = append(and, query)
	}
	if o.ReferenceCodeNlike != nil {
		query := bson.M{}
		query["referenceCode"] = bson.M{"$not": bson.M{"$regex": o.ReferenceCodeNlike, "$options": "i"}}
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
	if o.StatusEq != nil {
		query := bson.M{}
		query["status"] = o.StatusEq
		and = append(and, query)
	}
	if o.StatusNe != nil {
		query := bson.M{}
		query["status"] = bson.M{"$ne": o.StatusNe}
		and = append(and, query)
	}
	if o.StatusGt != nil {
		query := bson.M{}
		query["status"] = bson.M{"$gt": o.StatusGt}
		and = append(and, query)
	}
	if o.StatusGte != nil {
		query := bson.M{}
		query["status"] = bson.M{"$gte": o.StatusGte}
		and = append(and, query)
	}
	if o.StatusLt != nil {
		query := bson.M{}
		query["status"] = bson.M{"$lt": o.StatusLt}
		and = append(and, query)
	}
	if o.StatusLte != nil {
		query := bson.M{}
		query["status"] = bson.M{"$lte": o.StatusLte}
		and = append(and, query)
	}
	if o.StatusIn != nil {
		query := bson.M{}
		query["status"] = bson.M{"$in": o.StatusIn}
		and = append(and, query)
	}
	if o.StatusNin != nil {
		query := bson.M{}
		query["status"] = bson.M{"$nin": o.StatusNin}
		and = append(and, query)
	}
	if o.StatusExists != nil {
		query := bson.M{}
		query["status"] = bson.M{"$exists": *o.StatusExists}
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
	ContributorEmail int8
	CreatedAt        int8
	OwnerId          int8
	ReferenceCode    int8
	RegistryId       int8
	UpdatedAt        int8
}
