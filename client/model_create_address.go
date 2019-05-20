/*
 * Customers API
 *
 * Customers ...
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreateAddress struct {
	Type string `json:"type"`
	// First line of the address
	Address1 string `json:"address1"`
	// Second line of the address
	Address2 string `json:"address2"`
	City     string `json:"city"`
	// two charcer code of US state
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}