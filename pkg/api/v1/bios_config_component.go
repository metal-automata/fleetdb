package fleetdbapi

import (
	"time"

	"github.com/metal-automata/fleetdb/internal/models"
)

// BiosConfigComponent represents a BIOS Configuration Component
type BiosConfigComponent struct {
	ID        string              `json:"id"`
	Name      string              `json:"name" binding:"required"`
	Vendor    string              `json:"vendor"`
	Model     string              `json:"model"`
	Settings  []BiosConfigSetting `json:"settings" binding:"required"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}

func (cc *BiosConfigComponent) toDBModelBiosConfigComponent() *models.BiosConfigComponent {
	dbcc := &models.BiosConfigComponent{
		Name:   cc.Name,
		Vendor: cc.Vendor,
		Model:  cc.Model,
	}

	return dbcc
}

// fromDBModelBiosConfigComponent converts a models.BiosConfigComponent (created by sqlboiler) into a BiosConfigComponent
func (cc *BiosConfigComponent) fromDBModelBiosConfigComponent(component *models.BiosConfigComponent) {
	cc.ID = component.ID
	cc.Name = component.Name
	cc.Vendor = component.Vendor
	cc.Model = component.Model
	cc.CreatedAt = component.CreatedAt.Time
	cc.UpdatedAt = component.CreatedAt.Time

	if component.R != nil {
		cc.Settings = make([]BiosConfigSetting, len(component.R.FKBiosConfigComponentBiosConfigSettings))
		for i, dbSetting := range component.R.FKBiosConfigComponentBiosConfigSettings {
			cc.Settings[i].fromDBModelBiosConfigSetting(dbSetting)
		}
	}
}
