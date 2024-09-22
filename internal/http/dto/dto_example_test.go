package dto_test

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Customer struct {
	Name       string          `json:"name"`
	Gender     string          `json:"gender"`
	Email      string          `json:"email"`
	Address    Address         `json:"address"`
	Additional AdditionalMap   `json:"additional"`
	AddArr     []AdditionalMap `json:"additional_arrs"`
}

func (c Customer) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, &c,
		// Name cannot be empty, and the length must be between 5 and 20.
		validation.Field(&c.Name, validation.Required, validation.Length(5, 20)),
		// Gender is optional, and should be either "Female" or "Male".
		validation.Field(&c.Gender, validation.In("Female", "Male")),
		// Email cannot be empty and should be in a valid email format.
		validation.Field(&c.Email, validation.Required, is.Email),
		// Validate Address using its own validation rules
		validation.Field(&c.Address),
		// Validate Additional Address using its own validation rules
		validation.Field(&c.Additional),
		// Validate Additional Array Address using its own validation rules
		validation.Field(&c.AddArr),
	)
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

func (a Address) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, &a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Street, validation.Required, validation.Length(5, 50)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&a.City, validation.Required, validation.Length(5, 50)),
		// State cannot be empty, and must be a string consisting of two letters in upper case
		validation.Field(&a.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		// State cannot be empty, and must be a string consisting of five digits
		validation.Field(&a.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	)
}

type AdditionalMap map[string]any

func (m AdditionalMap) ValidateWithContext(ctx context.Context) error {
	addressRule := validation.Map(
		// Street cannot be empty, and the length must between 5 and 50
		validation.Key("street", validation.Required, validation.Length(5, 50)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Key("city", validation.Required, validation.Length(5, 50)),
		// State cannot be empty, and must be a string consisting of two letters in upper case
		validation.Key("state", validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		// State cannot be empty, and must be a string consisting of five digits
		validation.Key("zip", validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	)

	return validation.ValidateWithContext(ctx, m,
		validation.Map(
			// Name cannot be empty, and the length must be between 5 and 20.
			validation.Key("name", validation.Required, validation.Length(5, 20)),
			// Email cannot be empty and should be in a valid email format.
			validation.Key("email", validation.Required, is.Email),
			// Validate Address using its own validation rules
			validation.Key("address", addressRule),
			// Validate Address using its own validation rules
			validation.Key("arrs", validation.Each(addressRule)),
		),
	)
}

// Try Command:
// go test ./internal/http/dto/dto_example_test.go -v
func Test_DTO_Example(t *testing.T) {
	c := Customer{
		Name:  "Qiang Xue",
		Email: "q",
		Address: Address{
			Street: "123 Main Street",
			City:   "Unknown",
			State:  "Virginia",
			Zip:    "12345",
		},
		Additional: map[string]any{
			"name":  "Qiang Xue",
			"email": "q",
			"address": map[string]any{
				"street": "123",
				"city":   "Unknown",
				"state":  "Virginia",
				"zip":    "12345",
			},
		},
		AddArr: []AdditionalMap{
			{
				"name":  "",
				"email": "q",
				"address": map[string]any{
					"street": "123",
					"city":   "Unknown",
					"state":  "Virginia",
					"zip":    "12345",
				},
				"arrs": []map[string]any{
					{
						"street": "123",
						"city":   "Unknown",
						"state":  "Virginia",
						"zip":    "12345",
					},
				},
			},
		},
	}

	err := c.ValidateWithContext(context.Background())
	errs, _ := err.(validation.Errors)
	bytes, _ := json.MarshalIndent(errs, "", "\t")
	fmt.Println(string(bytes))
}
