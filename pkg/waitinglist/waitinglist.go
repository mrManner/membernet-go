package waitinglist // import "mrmanner.eu/go/membernet/pkg/waitinglist"

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
	"mrmanner.eu/go/membernet/pkg/shared"
)

// Contact describes contact details for a single person
type Contact struct {
	Name   string
	Phone  string
	Mobile string
	Email  string
}

// Address describes a geographic address for a member
type Address struct {
	AddressType int
	Street      string
	ZipCode     string
	ZipName     string
	Country     shared.Country
}

// Profile describes a full profile for the waiting list
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

// Register registers the new member described by profile in group
func Register(p Profile, leader bool, group string, apikey string, host string) error {
	u := &url.URL{
		Scheme: "https",
		Host:   host,
		User:   url.UserPassword(group, apikey),
		Path:   "api/organisation/register/member",
	}

	payload := url.Values{}
	payload.Set("profile[first_name]", p.FirstName)
	payload.Set("profile[last_name]", p.LastName)
	payload.Set("profile[email]", p.Email)
	payload.Set("profile[date_of_birth]", p.Dob)
	payload.Set("profile[ssno]", fmt.Sprintf("%04d", p.Ssno))
	payload.Set("profile[note]", p.Note)
	payload.Set("profile[sex]", p.Gender)
	payload.Set("address_list[addresses][address_1][address_type]", strconv.Itoa(p.Address.AddressType))
	payload.Set("address_list[addresses][address_1][address_line1]", p.Address.Street)
	payload.Set("address_list[addresses][address_1][country_code]", strconv.Itoa(int(p.Address.Country)))
	payload.Set("address_list[addresses][address_1][zip_code]", p.Address.ZipCode)
	payload.Set("address_list[addresses][address_1][zip_name]", p.Address.ZipName)

	payload.Set("contact_list[contacts][contact_1][contact_type_id]", "14")
	payload.Set("contact_list[contacts][contact_1][details]", p.Relative1.Name)
	payload.Set("contact_list[contacts][contact_2][contact_type_id]", "33")
	payload.Set("contact_list[contacts][contact_2][details]", p.Relative1.Email)
	payload.Set("contact_list[contacts][contact_3][contact_type_id]", "38")
	payload.Set("contact_list[contacts][contact_3][details]", p.Relative1.Mobile)
	payload.Set("contact_list[contacts][contact_4][contact_type_id]", "43")
	payload.Set("contact_list[contacts][contact_4][details]", p.Relative1.Phone)
	payload.Set("contact_list[contacts][contact_5][contact_type_id]", "16")
	payload.Set("contact_list[contacts][contact_5][details]", p.Relative2.Name)
	payload.Set("contact_list[contacts][contact_6][contact_type_id]", "34")
	payload.Set("contact_list[contacts][contact_6][details]", p.Relative2.Email)
	payload.Set("contact_list[contacts][contact_7][contact_type_id]", "39")
	payload.Set("contact_list[contacts][contact_7][details]", p.Relative2.Mobile)
	payload.Set("contact_list[contacts][contact_8][contact_type_id]", "44")
	payload.Set("contact_list[contacts][contact_8][details]", p.Relative2.Phone)
	payload.Set("contact_list[contacts][contact_9][contact_type_id]", "60")

	if leader {
		payload.Set("contact_list[contacts][contact_9][details]", "1")
	} else {
		payload.Set("contact_list[contacts][contact_9][details]", "0")
	}

	payload.Set("contact_list[contacts][contact_10][contact_type_id]", "1")
	payload.Set("contact_list[contacts][contact_10][details]", p.Mobile)

	resp, err := http.PostForm(u.String(), payload)
	if err != nil {
		return errors.Wrap(err, "Got error posting profile to Membernet")
	}
	if resp.StatusCode > 201 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrapf(err, "Got status %s from Membernet. When trying to read error message from API an error occured.", resp.Status)
		}
		return errors.New(fmt.Sprintf("Got status %s from Membernet: %s", resp.Status, body))
	}
	return nil
}
