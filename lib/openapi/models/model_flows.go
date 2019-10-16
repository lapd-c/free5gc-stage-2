/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Identifies the flows
type Flows struct {
	ContVers []int32 `json:"contVers,omitempty" bson:"contVers"`
	FNums    []int32 `json:"fNums,omitempty" bson:"fNums"`
	MedCompN int32   `json:"medCompN" bson:"medCompN"`
}