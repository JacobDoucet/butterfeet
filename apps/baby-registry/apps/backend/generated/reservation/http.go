package reservation

import (
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/actor_trace"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_reservation_status"
)

type HTTPRecord struct {
	Id                 *string                        `json:"id,omitempty"`
	ContactEmail       *string                        `json:"contactEmail,omitempty"`
	Created            *actor_trace.HTTPRecord        `json:"created,omitempty"`
	IsAnonymous        *bool                          `json:"isAnonymous,omitempty"`
	ItemId             *string                        `json:"itemId,omitempty"`
	Message            *string                        `json:"message,omitempty"`
	Quantity           *int                           `json:"quantity,omitempty"`
	RegistryId         *string                        `json:"registryId,omitempty"`
	ReserverName       *string                        `json:"reserverName,omitempty"`
	Status             *enum_reservation_status.Value `json:"status,omitempty"`
	Updated            *actor_trace.HTTPRecord        `json:"updated,omitempty"`
	UpdatedByOwnerUser *actor_trace.HTTPRecord        `json:"updatedByOwnerUser,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id
		m.Id = *elemid0
	}
	if r.ContactEmail != nil {
		elemcontactEmail0 := r.ContactEmail
		m.ContactEmail = *elemcontactEmail0
	}
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.IsAnonymous != nil {
		elemisAnonymous0 := r.IsAnonymous
		m.IsAnonymous = *elemisAnonymous0
	}
	if r.ItemId != nil {
		elemitemId0 := r.ItemId
		m.ItemId = *elemitemId0
	}
	if r.Message != nil {
		elemmessage0 := r.Message
		m.Message = *elemmessage0
	}
	if r.Quantity != nil {
		elemquantity0 := r.Quantity
		m.Quantity = *elemquantity0
	}
	if r.RegistryId != nil {
		elemregistryId0 := r.RegistryId
		m.RegistryId = *elemregistryId0
	}
	if r.ReserverName != nil {
		elemreserverName0 := r.ReserverName
		m.ReserverName = *elemreserverName0
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
	if r.ContactEmail != nil {
		p.ContactEmail = true
	}
	if r.Created != nil {
		p.Created = true
		p.CreatedFields = actor_trace.NewProjection(true)
	}
	if r.IsAnonymous != nil {
		p.IsAnonymous = true
	}
	if r.ItemId != nil {
		p.ItemId = true
	}
	if r.Message != nil {
		p.Message = true
	}
	if r.Quantity != nil {
		p.Quantity = true
	}
	if r.RegistryId != nil {
		p.RegistryId = true
	}
	if r.ReserverName != nil {
		p.ReserverName = true
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

type HTTPWhereClause struct {
	// id (Ref<Reservation>) search options
	IdEq     *string   `json:"idEq,omitempty"`
	IdIn     *[]string `json:"idIn,omitempty"`
	IdNin    *[]string `json:"idNin,omitempty"`
	IdExists *bool     `json:"idExists,omitempty"`
	// contactEmail (string) search options
	ContactEmailEq     *string   `json:"contactEmailEq,omitempty"`
	ContactEmailNe     *string   `json:"contactEmailNe,omitempty"`
	ContactEmailGt     *string   `json:"contactEmailGt,omitempty"`
	ContactEmailGte    *string   `json:"contactEmailGte,omitempty"`
	ContactEmailLt     *string   `json:"contactEmailLt,omitempty"`
	ContactEmailLte    *string   `json:"contactEmailLte,omitempty"`
	ContactEmailIn     *[]string `json:"contactEmailIn,omitempty"`
	ContactEmailNin    *[]string `json:"contactEmailNin,omitempty"`
	ContactEmailExists *bool     `json:"contactEmailExists,omitempty"`
	ContactEmailLike   *string   `json:"contactEmailLike,omitempty"`
	ContactEmailNlike  *string   `json:"contactEmailNlike,omitempty"`
	// created (ActorTrace) search options
	Created *actor_trace.HTTPWhereClause `json:"created,omitempty"`
	// isAnonymous (bool) search options
	IsAnonymousEq     *bool   `json:"isAnonymousEq,omitempty"`
	IsAnonymousNe     *bool   `json:"isAnonymousNe,omitempty"`
	IsAnonymousGt     *bool   `json:"isAnonymousGt,omitempty"`
	IsAnonymousGte    *bool   `json:"isAnonymousGte,omitempty"`
	IsAnonymousLt     *bool   `json:"isAnonymousLt,omitempty"`
	IsAnonymousLte    *bool   `json:"isAnonymousLte,omitempty"`
	IsAnonymousIn     *[]bool `json:"isAnonymousIn,omitempty"`
	IsAnonymousNin    *[]bool `json:"isAnonymousNin,omitempty"`
	IsAnonymousExists *bool   `json:"isAnonymousExists,omitempty"`
	// itemId (ParentRef<RegistryItem>) search options
	ItemIdEq     *string   `json:"itemIdEq,omitempty"`
	ItemIdIn     *[]string `json:"itemIdIn,omitempty"`
	ItemIdNin    *[]string `json:"itemIdNin,omitempty"`
	ItemIdExists *bool     `json:"itemIdExists,omitempty"`
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
	// quantity (int) search options
	QuantityEq     *int   `json:"quantityEq,omitempty"`
	QuantityNe     *int   `json:"quantityNe,omitempty"`
	QuantityGt     *int   `json:"quantityGt,omitempty"`
	QuantityGte    *int   `json:"quantityGte,omitempty"`
	QuantityLt     *int   `json:"quantityLt,omitempty"`
	QuantityLte    *int   `json:"quantityLte,omitempty"`
	QuantityIn     *[]int `json:"quantityIn,omitempty"`
	QuantityNin    *[]int `json:"quantityNin,omitempty"`
	QuantityExists *bool  `json:"quantityExists,omitempty"`
	// registryId (Ref<Registry>) search options
	RegistryIdEq     *string   `json:"registryIdEq,omitempty"`
	RegistryIdIn     *[]string `json:"registryIdIn,omitempty"`
	RegistryIdNin    *[]string `json:"registryIdNin,omitempty"`
	RegistryIdExists *bool     `json:"registryIdExists,omitempty"`
	// reserverName (string) search options
	ReserverNameEq     *string   `json:"reserverNameEq,omitempty"`
	ReserverNameNe     *string   `json:"reserverNameNe,omitempty"`
	ReserverNameGt     *string   `json:"reserverNameGt,omitempty"`
	ReserverNameGte    *string   `json:"reserverNameGte,omitempty"`
	ReserverNameLt     *string   `json:"reserverNameLt,omitempty"`
	ReserverNameLte    *string   `json:"reserverNameLte,omitempty"`
	ReserverNameIn     *[]string `json:"reserverNameIn,omitempty"`
	ReserverNameNin    *[]string `json:"reserverNameNin,omitempty"`
	ReserverNameExists *bool     `json:"reserverNameExists,omitempty"`
	ReserverNameLike   *string   `json:"reserverNameLike,omitempty"`
	ReserverNameNlike  *string   `json:"reserverNameNlike,omitempty"`
	// status (ReservationStatus) search options
	StatusEq     *enum_reservation_status.Value   `json:"statusEq,omitempty"`
	StatusNe     *enum_reservation_status.Value   `json:"statusNe,omitempty"`
	StatusGt     *enum_reservation_status.Value   `json:"statusGt,omitempty"`
	StatusGte    *enum_reservation_status.Value   `json:"statusGte,omitempty"`
	StatusLt     *enum_reservation_status.Value   `json:"statusLt,omitempty"`
	StatusLte    *enum_reservation_status.Value   `json:"statusLte,omitempty"`
	StatusIn     *[]enum_reservation_status.Value `json:"statusIn,omitempty"`
	StatusNin    *[]enum_reservation_status.Value `json:"statusNin,omitempty"`
	StatusExists *bool                            `json:"statusExists,omitempty"`
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
		elemcreated0, err := o.Created.ToWhereClause()
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
		elemitemIdEq0 := o.ItemIdEq
		to.ItemIdEq = elemitemIdEq0
	}
	if o.ItemIdIn != nil {
		elemitemIdIn0 := make([]string, 0)
		for _, oitemIdIn0 := range *o.ItemIdIn {
			elemitemIdIn1 := oitemIdIn0
			elemitemIdIn0 = append(elemitemIdIn0, elemitemIdIn1)
		}
		to.ItemIdIn = &elemitemIdIn0
	}
	if o.ItemIdNin != nil {
		elemitemIdNin0 := make([]string, 0)
		for _, oitemIdNin0 := range *o.ItemIdNin {
			elemitemIdNin1 := oitemIdNin0
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
	ItemId     *int8 `json:"itemId,omitempty"`
	RegistryId *int8 `json:"registryId,omitempty"`
	UpdatedAt  *int8 `json:"updatedAt,omitempty"`
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	if s.CreatedAt != nil {
		to.CreatedAt = *s.CreatedAt
	}
	if s.ItemId != nil {
		to.ItemId = *s.ItemId
	}
	if s.RegistryId != nil {
		to.RegistryId = *s.RegistryId
	}
	if s.UpdatedAt != nil {
		to.UpdatedAt = *s.UpdatedAt
	}
	return to
}
