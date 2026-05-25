package registry_api

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
)

type HTTPQueryResult struct {
	Data     []HTTPModel    `json:"data"`
	Total    int            `json:"total"`
	Skip     int            `json:"skip"`
	Metadata map[string]any `json:"metadata"`
}

func ToHTTPQueryResult(r QueryResult, projection Projection) (HTTPQueryResult, error) {
	data, err := ToHTTPModelList(r.Data, projection)
	return HTTPQueryResult{
		Data:     data,
		Total:    r.Total,
		Skip:     r.Skip,
		Metadata: make(map[string]any),
	}, err
}

type HTTPMutationResult struct {
	Data     registry.HTTPRecord `json:"data"`
	Metadata map[string]any      `json:"metadata"`
}

func ToHTTPMutationResult(obj registry.Model, projection registry.Projection) (HTTPMutationResult, error) {
	data, err := obj.ToHTTPRecord(projection)
	return HTTPMutationResult{
		Data:     data,
		Metadata: make(map[string]any),
	}, err
}

type HTTPDeleteResult struct {
	Id       string         `json:"id"`
	Metadata map[string]any `json:"metadata"`
}

func ToHTTPDeleteResult(id string) HTTPDeleteResult {
	return HTTPDeleteResult{
		Id:       id,
		Metadata: make(map[string]any),
	}
}

type HTTPModel struct {
	registry.HTTPRecord     `json:"registry"`
	AddressAccessSessions   *[]address_access_session.HTTPRecord   `json:"addressAccessSessions,omitempty"`
	RegistryApprovedGuests  *[]registry_approved_guest.HTTPRecord  `json:"registryApprovedGuests,omitempty"`
	RegistryItems           *[]registry_item.HTTPRecord            `json:"registryItems,omitempty"`
	Reservations            *[]reservation.HTTPRecord              `json:"reservations,omitempty"`
	ShippingAddressRequests *[]shipping_address_request.HTTPRecord `json:"shippingAddressRequests,omitempty"`
	Owner                   *owner_user.HTTPRecord                 `json:"owner,omitempty"`
}

type HTTPModelList []HTTPModel

func (r *HTTPModel) ToDomainModel() (Model, error) {
	m := Model{}
	var err error
	m.Model, err = r.ToModel()
	if r.AddressAccessSessions != nil {
		val := make([]address_access_session.Model, 0)
		var err error
		for _, rr := range *r.AddressAccessSessions {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.AddressAccessSessions = &val
	}
	if r.RegistryApprovedGuests != nil {
		val := make([]registry_approved_guest.Model, 0)
		var err error
		for _, rr := range *r.RegistryApprovedGuests {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.RegistryApprovedGuests = &val
	}
	if r.RegistryItems != nil {
		val := make([]registry_item.Model, 0)
		var err error
		for _, rr := range *r.RegistryItems {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.RegistryItems = &val
	}
	if r.Reservations != nil {
		val := make([]reservation.Model, 0)
		var err error
		for _, rr := range *r.Reservations {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.Reservations = &val
	}
	if r.ShippingAddressRequests != nil {
		val := make([]shipping_address_request.Model, 0)
		var err error
		for _, rr := range *r.ShippingAddressRequests {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.ShippingAddressRequests = &val
	}
	if r.Owner != nil {
		val, toModelErr := r.Owner.ToModel()
		if toModelErr != nil {
			err = errors.Join(err, toModelErr)
		}
		m.Owner = &val
	}
	return m, err
}

func (rs *HTTPModelList) ToDomainModel() ([]Model, error) {
	ms := make([]Model, len(*rs))
	var err error
	for i, r := range *rs {
		var iErr error
		ms[i], iErr = r.ToDomainModel()
		if iErr != nil {
			err = errors.Join(err, iErr)
		}
	}
	return ms, err
}

func ToHTTPModel(r Model, projection Projection) (HTTPModel, error) {
	m := HTTPModel{}
	var err error
	m.HTTPRecord, err = r.ToHTTPRecord(projection.Projection)
	if r.AddressAccessSessions != nil && projection.AddressAccessSessions != nil {
		refProjection := *projection.AddressAccessSessions
		val := make([]address_access_session.HTTPRecord, 0)
		for _, rr := range *r.AddressAccessSessions {
			nextVal, nextErr := rr.ToHTTPRecord(refProjection)
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.AddressAccessSessions = &val
	}
	if r.RegistryApprovedGuests != nil && projection.RegistryApprovedGuests != nil {
		refProjection := *projection.RegistryApprovedGuests
		val := make([]registry_approved_guest.HTTPRecord, 0)
		for _, rr := range *r.RegistryApprovedGuests {
			nextVal, nextErr := rr.ToHTTPRecord(refProjection)
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.RegistryApprovedGuests = &val
	}
	if r.RegistryItems != nil && projection.RegistryItems != nil {
		refProjection := *projection.RegistryItems
		val := make([]registry_item.HTTPRecord, 0)
		for _, rr := range *r.RegistryItems {
			nextVal, nextErr := rr.ToHTTPRecord(refProjection)
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.RegistryItems = &val
	}
	if r.Reservations != nil && projection.Reservations != nil {
		refProjection := *projection.Reservations
		val := make([]reservation.HTTPRecord, 0)
		for _, rr := range *r.Reservations {
			nextVal, nextErr := rr.ToHTTPRecord(refProjection)
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.Reservations = &val
	}
	if r.ShippingAddressRequests != nil && projection.ShippingAddressRequests != nil {
		refProjection := *projection.ShippingAddressRequests
		val := make([]shipping_address_request.HTTPRecord, 0)
		for _, rr := range *r.ShippingAddressRequests {
			nextVal, nextErr := rr.ToHTTPRecord(refProjection)
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.ShippingAddressRequests = &val
	}
	if r.Owner != nil && projection.Owner != nil {
		refProjection := *projection.Owner
		val, toHTTPRecordErr := r.Owner.ToHTTPRecord(refProjection)
		if toHTTPRecordErr != nil {
			err = errors.Join(err, toHTTPRecordErr)
		}
		m.Owner = &val
	}
	return m, err
}

func ToHTTPModelList(rs []Model, projection Projection) ([]HTTPModel, error) {
	ms := make([]HTTPModel, len(rs))
	var err error
	for i, r := range rs {
		var iErr error
		ms[i], iErr = ToHTTPModel(r, projection)
		if iErr != nil {
			err = errors.Join(err, iErr)
		}
	}
	return ms, err
}
