/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NetworkSliceCond struct {
	SnssaiList *[]Snssai `json:"snssaiList" bson:"snssaiList"`

	NsiList []string `json:"nsiList,omitempty" bson:"nsiList"`
}
