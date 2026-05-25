package address_access_session_api

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
	Data     address_access_session.HTTPRecord `json:"data"`
	Metadata map[string]any                    `json:"metadata"`
}

func ToHTTPMutationResult(obj address_access_session.Model, projection address_access_session.Projection) (HTTPMutationResult, error) {
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
	address_access_session.HTTPRecord `json:"addressAccessSession"`
	ApprovedGuest                     *registry_approved_guest.HTTPRecord `json:"approvedGuest,omitempty"`
	Owner                             *owner_user.HTTPRecord              `json:"owner,omitempty"`
	Registry                          *registry.HTTPRecord                `json:"registry,omitempty"`
}

type HTTPModelList []HTTPModel

func (r *HTTPModel) ToDomainModel() (Model, error) {
	m := Model{}
	var err error
	m.Model, err = r.ToModel()
	if r.ApprovedGuest != nil {
		val, toModelErr := r.ApprovedGuest.ToModel()
		if toModelErr != nil {
			err = errors.Join(err, toModelErr)
		}
		m.ApprovedGuest = &val
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
	if r.ApprovedGuest != nil && projection.ApprovedGuest != nil {
		refProjection := *projection.ApprovedGuest
		val, toHTTPRecordErr := r.ApprovedGuest.ToHTTPRecord(refProjection)
		if toHTTPRecordErr != nil {
			err = errors.Join(err, toHTTPRecordErr)
		}
		m.ApprovedGuest = &val
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
