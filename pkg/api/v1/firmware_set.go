package fleetdbapi

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/types"

	"github.com/metal-automata/fleetdb/internal/models"
)

// ComponentFirmwareSet represents a group of firmwares
type ComponentFirmwareSet struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	// TODO: remove the dependency on attributes and purge attributes storage
	Attributes        []Attributes               `json:"attributes"`
	ComponentFirmware []ComponentFirmwareVersion `json:"component_firmware"`
	UUID              uuid.UUID                  `json:"uuid"`
}

func convertFromDBModelAttributesFirmwareSet(dbAttrsFirmwareSet models.AttributesFirmwareSetSlice) ([]Attributes, error) {
	data := []Attributes{}

	if dbAttrsFirmwareSet == nil {
		return data, nil
	}

	for _, a := range dbAttrsFirmwareSet {
		data = append(data, Attributes{
			Namespace: a.Namespace,
			Data:      json.RawMessage(a.Data),
			CreatedAt: a.CreatedAt.Time,
			UpdatedAt: a.UpdatedAt.Time,
		})
	}

	return data, nil
}

func (a *Attributes) toDBModelAttributesFirmwareSet() *models.AttributesFirmwareSet {
	return &models.AttributesFirmwareSet{
		Namespace: a.Namespace,
		Data:      types.JSON(a.Data),
	}
}

func (s *ComponentFirmwareSet) fromDBModel(dbFS *models.ComponentFirmwareSet, firmwares []*models.ComponentFirmwareVersion) error {
	var err error

	s.UUID, err = uuid.Parse(dbFS.ID)
	if err != nil {
		return err
	}

	s.Name = dbFS.Name

	for _, firmware := range firmwares {
		f := ComponentFirmwareVersion{}

		errConv := f.fromDBModel(firmware)
		if errConv != nil {
			return errConv
		}

		s.ComponentFirmware = append(s.ComponentFirmware, f)
	}

	// relation attributes
	if dbFS.R != nil {
		s.Attributes, err = convertFromDBModelAttributesFirmwareSet(dbFS.R.FirmwareSetAttributesFirmwareSets)
		if err != nil {
			return err
		}
	}

	s.CreatedAt = dbFS.CreatedAt.Time
	s.UpdatedAt = dbFS.UpdatedAt.Time

	return nil
}

// ComponentFirmwareSetRequest represents the payload to create a firmware set
type ComponentFirmwareSetRequest struct {
	Name                   string       `json:"name"`
	Attributes             []Attributes `json:"attributes"`
	ComponentFirmwareUUIDs []string     `json:"component_firmware_uuids"`
	ID                     uuid.UUID    `json:"uuid"`
}

func (sc *ComponentFirmwareSetRequest) toDBModelFirmwareSet() (*models.ComponentFirmwareSet, error) {
	s := &models.ComponentFirmwareSet{
		ID:   sc.ID.String(),
		Name: sc.Name,
	}

	if sc.ID == uuid.Nil {
		s.ID = uuid.NewString()
	}

	return s, nil
}

type FirmwareSetValidation struct {
	TargetServer uuid.UUID `json:"target_server" binding:"required"`
	FirmwareSet  uuid.UUID `json:"firmware_set" binding:"required"`
	PerformedOn  time.Time `json:"performed_on" binding:"required"`
}
