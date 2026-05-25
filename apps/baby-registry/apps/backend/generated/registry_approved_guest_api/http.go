package registry_approved_guest_api

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
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
	Data     registry_approved_guest.HTTPRecord `json:"data"`
	Metadata map[string]any                     `json:"metadata"`
}

func ToHTTPMutationResult(obj registry_approved_guest.Model, projection registry_approved_guest.Projection) (HTTPMutationResult, error) {
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
	registry_approved_guest.HTTPRecord `json:"registryApprovedGuest"`
	AddressAccessSessions              *[]address_access_session.HTTPRecord `json:"addressAccessSessions,omitempty"`
	Owner                              *owner_user.HTTPRecord               `json:"owner,omitempty"`
	Registry                           *registry.HTTPRecord                 `json:"registry,omitempty"`
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
	if r.Owner != nil {
		val, toModelErr := r.Owner.ToModel()
		if toModelErr != nil {
			err = errors.Join(err, toModelErr)
		}
		m.Owner = &val
	}
	if r.Registry != nil {
		val, toModelErr := r.Registry.ToModel()
		if toModelErr != nil {
			err = errors.Join(err, toModelErr)
		}
		m.Registry = &val
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
	if r.Owner != nil && projection.Owner != nil {
		refProjection := *projection.Owner
		val, toHTTPRecordErr := r.Owner.ToHTTPRecord(refProjection)
		if toHTTPRecordErr != nil {
			err = errors.Join(err, toHTTPRecordErr)
		}
		m.Owner = &val
	}
	if r.Registry != nil && projection.Registry != nil {
		refProjection := *projection.Registry
		val, toHTTPRecordErr := r.Registry.ToHTTPRecord(refProjection)
		if toHTTPRecordErr != nil {
			err = errors.Join(err, toHTTPRecordErr)
		}
		m.Registry = &val
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
