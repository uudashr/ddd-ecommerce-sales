package order

import (
	"errors"
	"fmt"
)

type PostalAddress struct {
	line1    string
	line2    string
	city     string
	postcode string
	country  string
}

func NewPostalAddress(line1, line2, city, postcode, country string) (*PostalAddress, error) {
	if line1 == "" {
		return nil, errors.New("empty address line1")
	}

	if city == "" {
		return nil, errors.New("empty city")
	}

	if postcode == "" {
		return nil, errors.New("empty postcode")
	}

	if country == "" {
		return nil, errors.New("empty country")
	}

	return &PostalAddress{line1, line2, city, postcode, country}, nil
}

func (addr PostalAddress) Line1() string {
	return addr.line1
}

func (addr PostalAddress) Line2() string {
	return addr.line2
}

func (addr PostalAddress) City() string {
	return addr.city
}

func (addr PostalAddress) Postcode() string {
	return addr.postcode
}

func (addr PostalAddress) Country() string {
	return addr.country
}

func (addr PostalAddress) Equal(other *PostalAddress) bool {
	return addr.line1 == other.line1 &&
		addr.line2 == other.line2 &&
		addr.city == other.city &&
		addr.postcode == other.postcode &&
		addr.country == other.country
}

func (addr PostalAddress) String() string {
	return fmt.Sprintf("%s, %s, %s %s, %s", addr.line1, addr.line2, addr.city, addr.postcode, addr.country)
}
