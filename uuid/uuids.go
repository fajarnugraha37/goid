package uuid

// UUIDs is a slice of UUID types.
type UUIDs []UUID

// Strings returns a string slice containing the string form of each UUID in uuids.
func (uuids UUIDs) Strings() []string {
	var uuidStrs = make([]string, len(uuids))
	for i, uuid := range uuids {
		uuidStrs[i] = uuid.String()
	}
	return uuidStrs
}
