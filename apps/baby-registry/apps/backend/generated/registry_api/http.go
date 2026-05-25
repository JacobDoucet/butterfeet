package registry_api

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
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
	registry.HTTPRecord `json:"registry"`
	RegistryItems       *[]registry_item.HTTPRecord `json:"registryItems,omitempty"`
	Reservations        *[]reservation.HTTPRecord   `json:"reservations,omitempty"`
	Owner               *owner_user.HTTPRecord      `json:"owner,omitempty"`
}

type HTTPModelList []HTTPModel

func (r *HTTPModel) ToDomainModel() (Model, error) {
	m := Model{}
	var err error
	m.Model, err = r.ToModel()
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
