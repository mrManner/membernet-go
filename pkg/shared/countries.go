// Package constants contains definitions used by the Membernet APIs, such as
// numeric codes for countries, membership statuses, etc.
package shared // import "mrmanner.eu/go/membernet/pkg/shared"

// Country represents a country in an address or in a membership profile as
// an ISO 3166-1 numeric country code: https://www.iso.org/obp/ui/#search.
// Constant names are available both as a CamelCase version of the country's
// official name according to the standard, and as a two letter code
// (ISO 3166-1 Alpha-1).
//
// For copyright reasons only the most necessary subset is included here. If
// your application requires more country codes, please include them somehow.
type Country int

const (
	Belgium                                       Country = 056
	BE                                            Country = 056
	Denmark                                       Country = 208
	DK                                            Country = 208
	Finland                                       Country = 246
	FI                                            Country = 246
	Norway                                        Country = 578
	NO                                            Country = 578
	Sweden                                        Country = 752
	SE                                            Country = 752
	UnitedKingdomOfGreatBritainAndNorthernIreland Country = 826 // Official name according to standard
	UnitedKingdom                                 Country = 826 // Practical name...
	UK                                            Country = 826
	ÅlandIslands                                  Country = 248
	AX                                            Country = 248
)
