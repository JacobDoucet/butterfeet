package owner_user_api

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
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
	Data     owner_user.HTTPRecord `json:"data"`
	Metadata map[string]any        `json:"metadata"`
}

func ToHTTPMutationResult(obj owner_user.Model, projection owner_user.Projection) (HTTPMutationResult, error) {
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
	owner_user.HTTPRecord   `json:"ownerUser"`
	AddressAccessSessions   *[]address_access_session.HTTPRecord   `json:"addressAccessSessions,omitempty"`
	RegistryApprovedGuests  *[]registry_approved_guest.HTTPRecord  `json:"registryApprovedGuests,omitempty"`
	Registrys               *[]registry.HTTPRecord                 `json:"registrys,omitempty"`
	ShippingAddressRequests *[]shipping_address_request.HTTPRecord `json:"shippingAddressRequests,omitempty"`
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
	if r.Registrys != nil {
		val := make([]registry.Model, 0)
		var err error
		for _, rr := range *r.Registrys {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.Registrys = &val
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
	if r.Registrys != nil && projection.Registrys != nil {
		refProjection := *projection.Registrys
		val := make([]registry.HTTPRecord, 0)
		for _, rr := range *r.Registrys {
			nextVal, nextErr := rr.ToHTTPRecord(refProjection)
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.Registrys = &val
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
