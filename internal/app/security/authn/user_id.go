package authn

type UserID struct {
	IDInDomain string
}

func (userId *UserID) Equal(idInDomain, domainID string) bool {
	uid := *userId
	return idInDomain == uid.IDInDomain
}

func (userId *UserID) String() string {
	return userId.IDInDomain
}
