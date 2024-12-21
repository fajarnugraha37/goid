package goid

import (
	"github.com/fajarnugraha37/goid/ulid"
	"github.com/fajarnugraha37/goid/uuid"
)

func ULID() ulid.ULID {
	u := ulid.Make()
	return *u
}

func UUIDv1() uuid.UUID {
	return uuid.NewV1()
}

func UUIDv2(domain uuid.Domain, id uint32) uuid.UUID {
	return uuid.NewV2(domain, id)
}

func UUIDv3(name string) uuid.UUID {
	return uuid.NewV3(name)
}

func UUIDv4() uuid.UUID {
	return uuid.NewV4()
}

func UUIDv5(name string) uuid.UUID {
	return uuid.NewV5(name)
}

func UUIDv6() uuid.UUID {
	return uuid.NewV6()
}

func UUIDv7(name string) uuid.UUID {
	return uuid.NewV7()
}
