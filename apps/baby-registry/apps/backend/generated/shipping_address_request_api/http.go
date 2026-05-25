package shipping_address_request_api

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
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
	Data     shipping_address_request.HTTPRecord `json:"data"`
	Metadata map[string]any                      `json:"metadata"`
}

func ToHTTPMutationResult(obj shipping_address_request.Model, projection shipping_address_request.Projection) (HTTPMutationResult, error) {
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
	shipping_address_request.HTTPRecord `json:"shippingAddressRequest"`
	Owner                               *owner_user.HTTPRecord    `json:"owner,omitempty"`
	Registry                            *registry.HTTPRecord      `json:"registry,omitempty"`
	RegistryItem                        *registry_item.HTTPRecord `json:"registryItem,omitempty"`
}

type HTTPModelList []HTTPModel

func (r *HTTPModel) ToDomainModel() (Model, error) {
	m := Model{}
	var err error
	m.Model, err = r.ToModel()
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
	if r.RegistryItem != nil {
		val, toModelErr := r.RegistryItem.ToModel()
		if toModelErr != nil {
			err = errors.Join(err, toModelErr)
		}
		m.RegistryItem = &val
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
	if r.RegistryItem != nil && projection.RegistryItem != nil {
		refProjection := *projection.RegistryItem
		val, toHTTPRecordErr := r.RegistryItem.ToHTTPRecord(refProjection)
		if toHTTPRecordErr != nil {
			err = errors.Join(err, toHTTPRecordErr)
		}
		m.RegistryItem = &val
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
