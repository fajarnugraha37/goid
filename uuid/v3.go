package uuid

// NewV3 returns UUID based on MD5 hash of namespace UUID and name.
func NewV3(ns UUID, name string) (UUID, error) {
	uuid, err := NewV1()
	if err != nil {
		return Nil, nil
	}

	u := NewMD5(uuid, []byte(name))
	u.SetVersion(V3)
	u.SetVariant(RFC4122)

	return u, nil
}
