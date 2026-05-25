package reservation

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_reservation_status"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id                 string
	ContactEmail       string
	Created            actor_trace.Model
	IsAnonymous        bool
	ItemId             string
	Message            string
	Quantity           int
	RegistryId         string
	ReserverName       string
	Status             enum_reservation_status.Value
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
	if projection.ContactEmail {
		elemcontactEmail0 := m.ContactEmail
		r.ContactEmail = &elemcontactEmail0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToMongoRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.IsAnonymous {
		elemisAnonymous0 := m.IsAnonymous
		r.IsAnonymous = &elemisAnonymous0
	}
	if projection.ItemId && m.ItemId != "" {
		elemitemId0, err := primitive.ObjectIDFromHex(m.ItemId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.ItemId"), err)
		}
		r.ItemId = &elemitemId0
	}
	if projection.Message {
		elemmessage0 := m.Message
		r.Message = &elemmessage0
	}
	if projection.Quantity {
		elemquantity0 := m.Quantity
		r.Quantity = &elemquantity0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0, err := primitive.ObjectIDFromHex(m.RegistryId)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.RegistryId"), err)
		}
		r.RegistryId = &elemregistryId0
	}
	if projection.ReserverName {
		elemreserverName0 := m.ReserverName
		r.ReserverName = &elemreserverName0
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
	if projection.ContactEmail {
		elemcontactEmail0 := m.ContactEmail
		r.ContactEmail = &elemcontactEmail0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToHTTPRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.IsAnonymous {
		elemisAnonymous0 := m.IsAnonymous
		r.IsAnonymous = &elemisAnonymous0
	}
	if projection.ItemId && m.ItemId != "" {
		elemitemId0 := m.ItemId
		r.ItemId = &elemitemId0
	}
	if projection.Message {
		elemmessage0 := m.Message
		r.Message = &elemmessage0
	}
	if projection.Quantity {
		elemquantity0 := m.Quantity
		r.Quantity = &elemquantity0
	}
	if projection.RegistryId && m.RegistryId != "" {
		elemregistryId0 := m.RegistryId
		r.RegistryId = &elemregistryId0
	}
	if projection.ReserverName {
		elemreserverName0 := m.ReserverName
		r.ReserverName = &elemreserverName0
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

type WhereClause struct {
	// id (Ref<Reservation>) search options
	IdEq     *string
	IdIn     *[]string
	IdNin    *[]string
	IdExists *bool
	// contactEmail (string) search options
	ContactEmailEq     *string
	ContactEmailNe     *string
	ContactEmailGt     *string
	ContactEmailGte    *string
	ContactEmailLt     *string
	ContactEmailLte    *string
	ContactEmailIn     *[]string
	ContactEmailNin    *[]string
	ContactEmailExists *bool
	ContactEmailLike   *string
	ContactEmailNlike  *string
	// created (ActorTrace) search options
	Created *actor_trace.WhereClause
	// isAnonymous (bool) search options
	IsAnonymousEq     *bool
	IsAnonymousNe     *bool
	IsAnonymousGt     *bool
	IsAnonymousGte    *bool
	IsAnonymousLt     *bool
	IsAnonymousLte    *bool
	IsAnonymousIn     *[]bool
	IsAnonymousNin    *[]bool
	IsAnonymousExists *bool
	// itemId (ParentRef<RegistryItem>) search options
	ItemIdEq     *string
	ItemIdIn     *[]string
	ItemIdNin    *[]string
	ItemIdExists *bool
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
	// quantity (int) search options
	QuantityEq     *int
	QuantityNe     *int
	QuantityGt     *int
	QuantityGte    *int
	QuantityLt     *int
	QuantityLte    *int
	QuantityIn     *[]int
	QuantityNin    *[]int
	QuantityExists *bool
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *string
	RegistryIdIn     *[]string
	RegistryIdNin    *[]string
	RegistryIdExists *bool
	// reserverName (string) search options
	ReserverNameEq     *string
	ReserverNameNe     *string
	ReserverNameGt     *string
	ReserverNameGte    *string
	ReserverNameLt     *string
	ReserverNameLte    *string
	ReserverNameIn     *[]string
	ReserverNameNin    *[]string
	ReserverNameExists *bool
	ReserverNameLike   *string
	ReserverNameNlike  *string
	// status (ReservationStatus) search options
	StatusEq     *enum_reservation_status.Value
	StatusNe     *enum_reservation_status.Value
	StatusGt     *enum_reservation_status.Value
	StatusGte    *enum_reservation_status.Value
	StatusLt     *enum_reservation_status.Value
	StatusLte    *enum_reservation_status.Value
	StatusIn     *[]enum_reservation_status.Value
	StatusNin    *[]enum_reservation_status.Value
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
	if o.ContactEmailEq != nil {
		elemcontactEmailEq0 := o.ContactEmailEq
		to.ContactEmailEq = elemcontactEmailEq0
	}
	if o.ContactEmailNe != nil {
		elemcontactEmailNe0 := o.ContactEmailNe
		to.ContactEmailNe = elemcontactEmailNe0
	}
	if o.ContactEmailGt != nil {
		elemcontactEmailGt0 := o.ContactEmailGt
		to.ContactEmailGt = elemcontactEmailGt0
	}
	if o.ContactEmailGte != nil {
		elemcontactEmailGte0 := o.ContactEmailGte
		to.ContactEmailGte = elemcontactEmailGte0
	}
	if o.ContactEmailLt != nil {
		elemcontactEmailLt0 := o.ContactEmailLt
		to.ContactEmailLt = elemcontactEmailLt0
	}
	if o.ContactEmailLte != nil {
		elemcontactEmailLte0 := o.ContactEmailLte
		to.ContactEmailLte = elemcontactEmailLte0
	}
	if o.ContactEmailIn != nil {
		elemcontactEmailIn0 := make([]string, 0)
		for _, ocontactEmailIn0 := range *o.ContactEmailIn {
			elemcontactEmailIn1 := ocontactEmailIn0
			elemcontactEmailIn0 = append(elemcontactEmailIn0, elemcontactEmailIn1)
		}
		to.ContactEmailIn = &elemcontactEmailIn0
	}
	if o.ContactEmailNin != nil {
		elemcontactEmailNin0 := make([]string, 0)
		for _, ocontactEmailNin0 := range *o.ContactEmailNin {
			elemcontactEmailNin1 := ocontactEmailNin0
			elemcontactEmailNin0 = append(elemcontactEmailNin0, elemcontactEmailNin1)
		}
		to.ContactEmailNin = &elemcontactEmailNin0
	}
	if o.ContactEmailExists != nil {
		elemcontactEmailExists0 := o.ContactEmailExists
		to.ContactEmailExists = elemcontactEmailExists0
	}
	if o.ContactEmailLike != nil {
		elemcontactEmailLike0 := o.ContactEmailLike
		to.ContactEmailLike = elemcontactEmailLike0
	}
	if o.ContactEmailNlike != nil {
		elemcontactEmailNlike0 := o.ContactEmailNlike
		to.ContactEmailNlike = elemcontactEmailNlike0
	}
	if o.Created != nil {
		elemcreated0, err := o.Created.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.Created = &elemcreated0
	}
	if o.IsAnonymousEq != nil {
		elemisAnonymousEq0 := o.IsAnonymousEq
		to.IsAnonymousEq = elemisAnonymousEq0
	}
	if o.IsAnonymousNe != nil {
		elemisAnonymousNe0 := o.IsAnonymousNe
		to.IsAnonymousNe = elemisAnonymousNe0
	}
	if o.IsAnonymousGt != nil {
		elemisAnonymousGt0 := o.IsAnonymousGt
		to.IsAnonymousGt = elemisAnonymousGt0
	}
	if o.IsAnonymousGte != nil {
		elemisAnonymousGte0 := o.IsAnonymousGte
		to.IsAnonymousGte = elemisAnonymousGte0
	}
	if o.IsAnonymousLt != nil {
		elemisAnonymousLt0 := o.IsAnonymousLt
		to.IsAnonymousLt = elemisAnonymousLt0
	}
	if o.IsAnonymousLte != nil {
		elemisAnonymousLte0 := o.IsAnonymousLte
		to.IsAnonymousLte = elemisAnonymousLte0
	}
	if o.IsAnonymousIn != nil {
		elemisAnonymousIn0 := make([]bool, 0)
		for _, oisAnonymousIn0 := range *o.IsAnonymousIn {
			elemisAnonymousIn1 := oisAnonymousIn0
			elemisAnonymousIn0 = append(elemisAnonymousIn0, elemisAnonymousIn1)
		}
		to.IsAnonymousIn = &elemisAnonymousIn0
	}
	if o.IsAnonymousNin != nil {
		elemisAnonymousNin0 := make([]bool, 0)
		for _, oisAnonymousNin0 := range *o.IsAnonymousNin {
			elemisAnonymousNin1 := oisAnonymousNin0
			elemisAnonymousNin0 = append(elemisAnonymousNin0, elemisAnonymousNin1)
		}
		to.IsAnonymousNin = &elemisAnonymousNin0
	}
	if o.IsAnonymousExists != nil {
		elemisAnonymousExists0 := o.IsAnonymousExists
		to.IsAnonymousExists = elemisAnonymousExists0
	}
	if o.ItemIdEq != nil {
		elemitemIdEq0, err := primitive.ObjectIDFromHex(*o.ItemIdEq)
		if err != nil {
			return to, errors.Join(errors.New("invalid o.ItemIdEq"), err)
		}
		to.ItemIdEq = &elemitemIdEq0
	}
	if o.ItemIdIn != nil {
		elemitemIdIn0 := make([]primitive.ObjectID, 0)
		for _, oitemIdIn0 := range *o.ItemIdIn {
			elemitemIdIn1, err := primitive.ObjectIDFromHex(oitemIdIn0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oitemIdIn0"), err)
			}
			elemitemIdIn0 = append(elemitemIdIn0, elemitemIdIn1)
		}
		to.ItemIdIn = &elemitemIdIn0
	}
	if o.ItemIdNin != nil {
		elemitemIdNin0 := make([]primitive.ObjectID, 0)
		for _, oitemIdNin0 := range *o.ItemIdNin {
			elemitemIdNin1, err := primitive.ObjectIDFromHex(oitemIdNin0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oitemIdNin0"), err)
			}
			elemitemIdNin0 = append(elemitemIdNin0, elemitemIdNin1)
		}
		to.ItemIdNin = &elemitemIdNin0
	}
	if o.ItemIdExists != nil {
		elemitemIdExists0 := o.ItemIdExists
		to.ItemIdExists = elemitemIdExists0
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
	if o.ReserverNameEq != nil {
		elemreserverNameEq0 := o.ReserverNameEq
		to.ReserverNameEq = elemreserverNameEq0
	}
	if o.ReserverNameNe != nil {
		elemreserverNameNe0 := o.ReserverNameNe
		to.ReserverNameNe = elemreserverNameNe0
	}
	if o.ReserverNameGt != nil {
		elemreserverNameGt0 := o.ReserverNameGt
		to.ReserverNameGt = elemreserverNameGt0
	}
	if o.ReserverNameGte != nil {
		elemreserverNameGte0 := o.ReserverNameGte
		to.ReserverNameGte = elemreserverNameGte0
	}
	if o.ReserverNameLt != nil {
		elemreserverNameLt0 := o.ReserverNameLt
		to.ReserverNameLt = elemreserverNameLt0
	}
	if o.ReserverNameLte != nil {
		elemreserverNameLte0 := o.ReserverNameLte
		to.ReserverNameLte = elemreserverNameLte0
	}
	if o.ReserverNameIn != nil {
		elemreserverNameIn0 := make([]string, 0)
		for _, oreserverNameIn0 := range *o.ReserverNameIn {
			elemreserverNameIn1 := oreserverNameIn0
			elemreserverNameIn0 = append(elemreserverNameIn0, elemreserverNameIn1)
		}
		to.ReserverNameIn = &elemreserverNameIn0
	}
	if o.ReserverNameNin != nil {
		elemreserverNameNin0 := make([]string, 0)
		for _, oreserverNameNin0 := range *o.ReserverNameNin {
			elemreserverNameNin1 := oreserverNameNin0
			elemreserverNameNin0 = append(elemreserverNameNin0, elemreserverNameNin1)
		}
		to.ReserverNameNin = &elemreserverNameNin0
	}
	if o.ReserverNameExists != nil {
		elemreserverNameExists0 := o.ReserverNameExists
		to.ReserverNameExists = elemreserverNameExists0
	}
	if o.ReserverNameLike != nil {
		elemreserverNameLike0 := o.ReserverNameLike
		to.ReserverNameLike = elemreserverNameLike0
	}
	if o.ReserverNameNlike != nil {
		elemreserverNameNlike0 := o.ReserverNameNlike
		to.ReserverNameNlike = elemreserverNameNlike0
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
		elemstatusIn0 := make([]enum_reservation_status.Value, 0)
		for _, ostatusIn0 := range *o.StatusIn {
			elemstatusIn1 := ostatusIn0
			elemstatusIn0 = append(elemstatusIn0, elemstatusIn1)
		}
		to.StatusIn = &elemstatusIn0
	}
	if o.StatusNin != nil {
		elemstatusNin0 := make([]enum_reservation_status.Value, 0)
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
	CreatedAt  int8
	ItemId     int8
	RegistryId int8
	UpdatedAt  int8
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	to.CreatedAt = s.CreatedAt
	to.ItemId = s.ItemId
	to.RegistryId = s.RegistryId
	to.UpdatedAt = s.UpdatedAt
	return to
}
