package owner_user

// GetOwnerId satisfies the permissions.Actor interface for OwnerUser.
// An OwnerUser owns themselves, so OwnerId == ActorId == Id.
func (m *Model) GetOwnerId() string {
	return m.Id
}
