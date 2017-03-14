package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// VolumeDetailResponse volume detail response
// swagger:model VolumeDetailResponse
type VolumeDetailResponse struct {

	// availability zone
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// consistencygroup Id
	ConsistencygroupID string `json:"consistencygroupId,omitempty"`

	// create at
	CreateAt strfmt.DateTime `json:"createAt,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// encrypted
	Encrypted *bool `json:"encrypted,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// migration status
	MigrationStatus string `json:"migrationStatus,omitempty"`

	// multiattach
	Multiattach *bool `json:"multiattach,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// protected
	Protected *bool `json:"protected,omitempty"`

	// replication status
	ReplicationStatus string `json:"replicationStatus,omitempty"`

	// size
	Size int32 `json:"size,omitempty"`

	// snapshot Id
	SnapshotID string `json:"snapshotId,omitempty"`

	// source volid
	SourceVolid string `json:"sourceVolid,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`

	// volume type
	VolumeType string `json:"volumeType,omitempty"`
}

// Validate validates this volume detail response
func (m *VolumeDetailResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
