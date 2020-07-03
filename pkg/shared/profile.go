package shared

// Contact represents contact details for a relative, to be included
// on a Profile. It does *not* represent the contact details on a
// member Profile.
type Contact struct {
	Name   string // The person's name (or the nickname for the contact)
	Phone  string // Landline phone number, subject to some unspecified checks
	Mobile string // Mobile phone number, subject to some unspecified checks
	Email  string // Contact email, subject to some checks but ownership is not verified
}

// Address represents an address.
type Address struct {
	AddressType int
	Street      string
	ZipCode     string
	ZipName     string
	Country     Country
}

type Profile struct {
	Dob       string
	Ssno      int
	Note      string
	FirstName string
	LastName  string
	Gender    string
	Email     string
	Mobile    string
	Phone     string
	Address   Address
	Relative1 Contact
	Relative2 Contact
}
