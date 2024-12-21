package uuid

// NewV3 returns UUID based on MD5 hash of namespace UUID and name.
func NewV3(name string) UUID {
	return Must(newV3(name))
}
func newV3(name string) (UUID, error) {
	uuid, err := newV1()
	if err != nil {
		return Nil, nil
	}

	u := NewMD5(uuid, []byte(name))
	u.SetVersion(V3)
	u.SetVariant(RFC4122)

	return u, nil
}
