package uuid

// NewV5 returns UUID based on SHA-1 hash of namespace UUID and name.
func NewV5(ns UUID, name string) (UUID, error) {
	uuid, err := NewV1()
	if err != nil {
		return Nil, nil
	}

	u := NewSHA1(uuid, []byte(name))
	u.SetVersion(V5)
	u.SetVariant(RFC4122)

	return u, nil
}
