package registry_item

import (
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_model"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/event_subject"
)

func (m *Model) AsEventSubject() (event_subject.Model, error) {
	if m.Id == "" {
		return event_subject.Model{}, errors.New("registryItem does not have an id")
	}
	return event_subject.Model{
		SubjectId:   m.Id,
		SubjectType: enum_model.RegistryItem,
	}, nil
}
