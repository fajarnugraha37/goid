package uuid

// NewV2 returns DCE Security UUID based on POSIX UID/GID.
func NewV2(domain Domain, id uint32) (UUID, error) {
	uuid, err := NewDCESecurity(domain, id)
	if err != nil {
		return Nil, err
	}
	uuid.SetVersion(V2)
	uuid.SetVariant(RFC4122)

	return uuid, nil
}
