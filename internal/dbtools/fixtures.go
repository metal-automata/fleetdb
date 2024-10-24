//go:build testtools
// +build testtools

package dbtools

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"

	"github.com/metal-automata/fleetdb/internal/models"
)

//nolint:revive
var (
	// Namespaces used in Attributes and VersionedAttributes
	FixtureNamespaceMetadata    = "hollow.metadata"
	FixtureNamespaceOtherdata   = "hollow.other_data"
	FixtureNamespaceVersioned   = "hollow.versioned"
	FixtureNamespaceVersionedV2 = "hollow.versioned.v2"

	// Server Component Types
	FixtureFinType *models.ServerComponentType

	NemoID                       uuid.UUID
	FixtureNemo                  *models.Server
	FixtureNemoMetadata          *models.Attribute
	FixtureNemoOtherdata         *models.Attribute
	FixtureNemoLeftFin           *models.ServerComponent
	FixtureNemoRightFin          *models.ServerComponent
	FixtureNemoLeftFinVersioned  *models.VersionedAttribute
	FixtureNemoVersionedNew      *models.VersionedAttribute
	FixtureNemoVersionedOld      *models.VersionedAttribute
	FixtureNemoVersionedV2       *models.VersionedAttribute
	FixtureNemoBMCSecret         *models.ServerCredential
	FixtureNemoRightFinOtherData *models.Attribute

	FixtureDory          *models.Server
	FixtureDoryMetadata  *models.Attribute
	FixtureDoryOtherdata *models.Attribute
	FixtureDoryLeftFin   *models.ServerComponent
	FixtureDoryRightFin  *models.ServerComponent

	FixtureMarlin          *models.Server
	FixtureMarlinMetadata  *models.Attribute
	FixtureMarlinOtherdata *models.Attribute
	FixtureMarlinLeftFin   *models.ServerComponent
	FixtureMarlinRightFin  *models.ServerComponent

	// FixtureChuckles represents the fish that was deleted
	// https://pixar.fandom.com/wiki/Chuckles_(Finding_Nemo)
	FixtureChuckles          *models.Server
	FixtureChucklesMetadata  *models.Attribute
	FixtureChucklesOtherdata *models.Attribute
	FixtureChucklesLeftFin   *models.ServerComponent

	FixtureServers        models.ServerSlice
	FixtureDeletedServers models.ServerSlice
	FixtureAllServers     models.ServerSlice

	// ComponentFirmwareVersion fixtures
	FixtureDellR640BMC      *models.ComponentFirmwareVersion
	FixtureDellR640BIOS     *models.ComponentFirmwareVersion
	FixtureDellR640CPLD     *models.ComponentFirmwareVersion
	FixtureDellR6515BMC     *models.ComponentFirmwareVersion
	FixtureDellR6515BIOS    *models.ComponentFirmwareVersion
	FixtureServerComponents models.ServerComponentSlice

	// ComponentFirmwareSet fixtures
	FixtureSuperMicroX11DPHTBMC        *models.ComponentFirmwareVersion
	FixtureFirmwareUUIDsSuperMicro     []string
	FixtureFirmwareSetX11DPHT          *models.ComponentFirmwareSet
	FixtureFirmwareSetX11DPHTAttribute *models.AttributesFirmwareSet
	FixtureFirmwareInbandNIC           *models.ComponentFirmwareVersion
	FixtureFirmwareOem                 *models.ComponentFirmwareVersion

	FixtureFirmwareUUIDsR6515        []string
	FixtureFirmwareSetR6515          *models.ComponentFirmwareSet
	FixtureFirmwareSetR6515Attribute *models.AttributesFirmwareSet

	FixtureFirmwareUUIDsR640        []string
	FixtureFirmwareSetR640          *models.ComponentFirmwareSet
	FixtureFirmwareSetR640Attribute *models.AttributesFirmwareSet

	// Inventory fixtures
	FixtureInventoryServer *models.Server

	FixtureBiosConfigSet        *models.BiosConfigSet
	FixtureBiosConfigComponents []*models.BiosConfigComponent
	FixtureBiosConfigSettings   [][]*models.BiosConfigSetting

	FixtureEventHistoryServer    *models.Server
	FixtureEventHistoryRelatedID uuid.UUID
	FixtureEventHistories        []*models.EventHistory

	FixtureFWValidationServer *models.Server
	FixtureFWValidationSet    *models.ComponentFirmwareSet

	FixtureHardwareVendors   []*models.HardwareVendor
	FixtureHardwareVendorBaz *models.HardwareVendor
	FixtureHardwareVendorBar *models.HardwareVendor
	// this fixture is deleted in tests - do not create relations
	FixtureHardwareVendorFoo     *models.HardwareVendor
	FixtureHardwareVendorNameBaz = "baz"
	FixtureHardwareVendorNameBar = "bar"
	FixtureHardwareVendorNameFoo = "foo"

	FixtureHardwareModels      []*models.HardwareModel
	FixtureHardwareModelBaz123 *models.HardwareModel
	FixtureHardwareModelBar123 *models.HardwareModel
	// this fixture is deleted in tests - do not create relations
	FixtureHardwareModelFoo789     *models.HardwareModel
	FixtureHardwareModelBaz123Name = "123"
	FixtureHardwareModelBar456Name = "456"
	FixtureHardwareModelFoo789Name = "789"

	FixtureServerBMCs []*models.ServerBMC
	FixtureServerBMC1 *models.ServerBMC
	FixtureServerBMC2 *models.ServerBMC
)

func addFixtures(t *testing.T) error {
	ctx := context.TODO()

	FixtureFinType = &models.ServerComponentType{
		Name: "Fins",
		Slug: "fins",
	}

	if err := FixtureFinType.Insert(ctx, testDB, boil.Infer()); err != nil {
		return err
	}

	if err := setupNemo(ctx, testDB, t); err != nil {
		return err
	}

	if err := setupDory(ctx, testDB); err != nil {
		return err
	}

	if err := setupMarlin(ctx, testDB); err != nil {
		return err
	}

	if err := setupChuckles(ctx, testDB); err != nil {
		return err
	}

	if err := setupFirmwareDellR640(ctx, testDB); err != nil {
		return err
	}

	if err := setupFirmwareDellR6515(ctx, testDB); err != nil {
		return err
	}

	if err := setupFirmwareSuperMicro(ctx, testDB); err != nil {
		return err
	}

	if err := setupFirmwareSetR6515(ctx, testDB); err != nil {
		return err
	}

	if err := setupFirmwareSetR640(ctx, testDB); err != nil {
		return err
	}

	if err := setupFirmwareSetR640WithLabels(ctx, testDB, "org", `"organization": "2a0834e9-4720-4193-93e4-24185ac4949c"`); err != nil {
		return err
	}

	if err := setupFirmwareSetR640WithLabels(ctx, testDB, "project", `"project": "69096c3f-1f96-434c-bd61-b9aef0b5746b"`); err != nil {
		return err
	}

	if err := setupFirmwareSetSupermicroX11DPHT(ctx, testDB); err != nil {
		return err
	}

	if err := setupFirmwareInbandNIC(ctx, testDB); err != nil {
		return err
	}

	if err := SetupComponentTypes(ctx, testDB); err != nil {
		return err
	}

	if err := setupInventoryFixture(ctx, testDB); err != nil {
		return err
	}

	if err := setupEventHistoryFixtures(ctx, testDB); err != nil {
		return err
	}

	if err := setupConfigSet(ctx, testDB); err != nil {
		return err
	}

	if err := setupFWValidationFixtures(ctx, testDB); err != nil {
		return err
	}

	if err := setupHardwareVendorFixtures(ctx, testDB); err != nil {
		return err
	}

	if err := setupHardwareModelFixtures(ctx, testDB); err != nil {
		return err
	}

	if err := setupServerBMCFixtures(ctx, testDB); err != nil {
		return err
	}

	// excluding Chuckles here since that server is deleted
	FixtureServers = models.ServerSlice{FixtureNemo, FixtureDory, FixtureMarlin}
	FixtureDeletedServers = models.ServerSlice{FixtureChuckles}

	FixtureServerComponents = models.ServerComponentSlice{FixtureDoryLeftFin, FixtureDoryRightFin}

	//nolint:gocritic
	FixtureAllServers = append(FixtureServers, FixtureDeletedServers...)

	return nil
}

func setupNemo(ctx context.Context, db *sqlx.DB, t *testing.T) error {
	FixtureNemo = &models.Server{
		Name:         null.StringFrom("Nemo"),
		FacilityCode: null.StringFrom("Sydney"),
	}

	if err := FixtureNemo.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	bmc, err := models.ServerCredentialTypes(models.ServerCredentialTypeWhere.Slug.EQ("bmc")).One(ctx, db)
	if err != nil {
		return err
	}

	keeper := TestSecretKeeper(t)

	value, err := Encrypt(ctx, keeper, "super-secret-bmc-password")
	if err != nil {
		return err
	}

	FixtureNemoBMCSecret = &models.ServerCredential{
		ServerCredentialTypeID: bmc.ID,
		Password:               value,
	}

	if err := FixtureNemo.AddServerCredentials(ctx, db, true, FixtureNemoBMCSecret); err != nil {
		return err
	}

	FixtureNemoMetadata = &models.Attribute{
		Namespace: FixtureNamespaceMetadata,
		Data:      types.JSON([]byte(`{"age":6,"location":"Fishbowl"}`)),
	}

	FixtureNemoOtherdata = &models.Attribute{
		Namespace: FixtureNamespaceOtherdata,
		Data:      types.JSON([]byte(`{"enabled": true, "type": "clown", "lastUpdated": 1624960800, "nested": {"tag": "finding-nemo", "number": 1}}`)),
	}

	if err := FixtureNemo.AddAttributes(ctx, db, true, FixtureNemoMetadata, FixtureNemoOtherdata); err != nil {
		return err
	}

	FixtureNemoLeftFin = &models.ServerComponent{
		ServerComponentTypeID: FixtureFinType.ID,
		Name:                  null.StringFrom("Normal Fin"),
		Model:                 null.StringFrom("Normal Fin"),
		Serial:                null.StringFrom("Left"),
	}

	FixtureNemoRightFin = &models.ServerComponent{
		ServerComponentTypeID: FixtureFinType.ID,
		Name:                  null.StringFrom("My Lucky Fin"),
		Vendor:                null.StringFrom("Barracuda"),
		Model:                 null.StringFrom("A Lucky Fin"),
		Serial:                null.StringFrom("Right"),
	}

	if err := FixtureNemo.AddServerComponents(ctx, db, true, FixtureNemoLeftFin, FixtureNemoRightFin); err != nil {
		return err
	}

	FixtureNemoRightFinOtherData = &models.Attribute{
		Namespace: FixtureNamespaceOtherdata,
		Data:      types.JSON([]byte(`{"twitchy": true}`)),
	}

	if err := FixtureNemoRightFin.AddAttributes(ctx, db, true, FixtureNemoRightFinOtherData); err != nil {
		return err
	}

	FixtureNemoLeftFinVersioned = &models.VersionedAttribute{
		Namespace: FixtureNamespaceVersioned,
		Data:      types.JSON([]byte(`{"something": "cool"}`)),
	}

	if err := FixtureNemoLeftFin.AddVersionedAttributes(ctx, db, true, FixtureNemoLeftFinVersioned); err != nil {
		return err
	}

	FixtureNemoVersionedV2 = &models.VersionedAttribute{
		Namespace: FixtureNamespaceVersionedV2,
		Data:      types.JSON([]byte(`{"something": "cool"}`)),
	}

	if err := FixtureNemo.AddVersionedAttributes(ctx, db, true, FixtureNemoVersionedV2); err != nil {
		return err
	}

	FixtureNemoVersionedOld = &models.VersionedAttribute{
		Namespace: FixtureNamespaceVersioned,
		Data:      types.JSON([]byte(`{"name": "old"}`)),
	}

	FixtureNemoVersionedNew = &models.VersionedAttribute{
		Namespace: FixtureNamespaceVersioned,
		Data:      types.JSON([]byte(`{"name": "new"}`)),
	}

	// Insert old and new in a separate transaction to ensure the new one has a later timestamp and is indeed new
	if err := FixtureNemo.AddVersionedAttributes(ctx, db, true, FixtureNemoVersionedOld); err != nil {
		return err
	}

	if err := FixtureNemo.AddVersionedAttributes(ctx, db, true, FixtureNemoVersionedNew); err != nil {
		return err
	}
	NemoID = uuid.MustParse(FixtureNemo.ID)
	return nil
}

func setupDory(ctx context.Context, db *sqlx.DB) error {
	FixtureDory = &models.Server{
		Name:         null.StringFrom("Dory"),
		FacilityCode: null.StringFrom("Ocean"),
	}

	if err := FixtureDory.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	FixtureDoryMetadata = &models.Attribute{
		Namespace: FixtureNamespaceMetadata,
		Data:      types.JSON([]byte(`{"age":12,"location":"East Australian Current"}`)),
	}

	FixtureDoryOtherdata = &models.Attribute{
		Namespace: FixtureNamespaceOtherdata,
		Data:      types.JSON([]byte(`{"enabled": true, "type": "blue-tang", "lastUpdated": 1624960400, "nested": {"tag": "finding-nemo", "number": 2}}`)),
	}

	if err := FixtureDory.AddAttributes(ctx, db, true, FixtureDoryMetadata, FixtureDoryOtherdata); err != nil {
		return err
	}

	FixtureDoryLeftFin = &models.ServerComponent{
		ServerComponentTypeID: FixtureFinType.ID,
		Name:                  null.StringFrom("Normal Fin"),
		Model:                 null.StringFrom("Normal Fin"),
		Serial:                null.StringFrom("Left"),
	}

	FixtureDoryRightFin = &models.ServerComponent{
		ServerComponentTypeID: FixtureFinType.ID,
		Name:                  null.StringFrom("Normal Fin"),
		Serial:                null.StringFrom("Right"),
	}

	return FixtureDory.AddServerComponents(ctx, db, true, FixtureDoryLeftFin, FixtureDoryRightFin)
}

func setupMarlin(ctx context.Context, db *sqlx.DB) error {
	FixtureMarlin = &models.Server{
		Name:         null.StringFrom("Marlin"),
		FacilityCode: null.StringFrom("Ocean"),
	}

	if err := FixtureMarlin.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	FixtureMarlinMetadata = &models.Attribute{
		Namespace: FixtureNamespaceMetadata,
		Data:      types.JSON([]byte(`{"age":10,"location":"East Australian Current"}`)),
	}

	FixtureMarlinOtherdata = &models.Attribute{
		Namespace: FixtureNamespaceOtherdata,
		Data:      types.JSON([]byte(`{"enabled": false, "type": "clown", "lastUpdated": 1624960000, "nested": {"tag": "finding-nemo", "number": 3}}`)),
	}

	if err := FixtureMarlin.AddAttributes(ctx, db, true, FixtureMarlinMetadata, FixtureMarlinOtherdata); err != nil {
		return err
	}

	FixtureMarlinLeftFin = &models.ServerComponent{
		ServerComponentTypeID: FixtureFinType.ID,
		Name:                  null.StringFrom("Normal Fin"),
		Model:                 null.StringFrom("Normal Fin"),
		Serial:                null.StringFrom("Left"),
	}

	FixtureMarlinRightFin = &models.ServerComponent{
		ServerComponentTypeID: FixtureFinType.ID,
		Name:                  null.StringFrom("Normal Fin"),
		Serial:                null.StringFrom("Right"),
	}

	return FixtureMarlin.AddServerComponents(ctx, db, true, FixtureMarlinLeftFin, FixtureMarlinRightFin)
}

func setupChuckles(ctx context.Context, db *sqlx.DB) error {
	FixtureChuckles = &models.Server{
		Name:         null.StringFrom("Chuckles"),
		FacilityCode: null.StringFrom("Aquarium"),
		DeletedAt:    null.TimeFrom(time.Date(2003, 5, 30, 0, 0, 0, 0, time.UTC)),
	}

	if err := FixtureChuckles.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	FixtureChucklesMetadata = &models.Attribute{
		Namespace: FixtureNamespaceMetadata,
		Data:      types.JSON([]byte(`{"age":1,"location":"Old shipwreck"}`)),
	}

	FixtureChucklesOtherdata = &models.Attribute{
		Namespace: FixtureNamespaceOtherdata,
		Data:      types.JSON([]byte(`{"enabled": false, "type": "goldfish", "lastUpdated": 1624960000, "nested": {"tag": "finding-nemo", "number": 4}}`)),
	}

	if err := FixtureChuckles.AddAttributes(ctx, db, true, FixtureChucklesMetadata, FixtureChucklesOtherdata); err != nil {
		return err
	}

	FixtureChucklesLeftFin = &models.ServerComponent{
		ServerComponentTypeID: FixtureFinType.ID,
		Model:                 null.StringFrom("Belly"),
		Serial:                null.StringFrom("Up"),
	}

	return FixtureChuckles.AddServerComponents(ctx, db, true, FixtureChucklesLeftFin)
}

func setupFirmwareDellR640(ctx context.Context, db *sqlx.DB) error {
	FixtureFirmwareUUIDsR640 = []string{}

	FixtureDellR640BMC = &models.ComponentFirmwareVersion{
		Vendor:        "Dell",
		Model:         types.StringArray{"R640"},
		Filename:      "iDRAC-with-Lifecycle-Controller_Firmware_P8HC9_WN64_5.10.00.00_A00.EXE",
		Version:       "5.10.00.00",
		Component:     "bmc",
		Checksum:      "98db2fe5bca0745151d678ddeb26679464ccb13ca3f1a3d289b77e211344402f",
		UpstreamURL:   "https://vendor.com/firmwares/iDRAC-with-Lifecycle-Controller_Firmware_P8HC9_WN64_5.10.00.00_A00.EXE",
		RepositoryURL: "https://example-firmware-bucket.s3.amazonaws.com/firmware/dell/r640/bmc/iDRAC-with-Lifecycle-Controller_Firmware_P8HC9_WN64_5.10.00.00_A00.EXE",
	}

	if err := FixtureDellR640BMC.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	FixtureFirmwareUUIDsR640 = append(FixtureFirmwareUUIDsR640, FixtureDellR640BMC.ID)

	FixtureDellR640BIOS = &models.ComponentFirmwareVersion{
		Vendor:        "Dell",
		Model:         types.StringArray{"R640"},
		Filename:      "bios.exe",
		Version:       "2.4.4",
		Component:     "bios",
		Checksum:      "78ad2fe5bca0745151d678ddeb26679464ccb13ca3f1a3d289b77e211344402f",
		UpstreamURL:   "https://vendor.com/firmwares/bios-2.4.4.EXE",
		RepositoryURL: "https://example-firmware-bucket.s3.amazonaws.com/firmware/dell/r640/bios/bios-2.4.4.EXE",
	}

	if err := FixtureDellR640BIOS.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	FixtureFirmwareUUIDsR640 = append(FixtureFirmwareUUIDsR640, FixtureDellR640BIOS.ID)

	// This fixture is not included in FixtureFirmwareUUIDsR640 slice
	// since its part of the test TestIntegrationServerComponentFirmwareSetUpdate
	// where its added into the firmware set.
	FixtureDellR640CPLD = &models.ComponentFirmwareVersion{
		Vendor:        "Dell",
		Model:         types.StringArray{"R640"},
		Filename:      "cpld.exe",
		Version:       "1.0.1",
		Component:     "cpld",
		Checksum:      "676d2fe5bca0745151d678ddeb26679464ccb13ca3f1a3d289b77e211344402f",
		UpstreamURL:   "https://vendor.com/firmwares/cpld.exe",
		RepositoryURL: "https://example-firmware-bucket.s3.amazonaws.com/firmware/dell/r640/cpld/cpld.EXE",
	}

	return FixtureDellR640CPLD.Insert(ctx, db, boil.Infer())
}

func setupFirmwareDellR6515(ctx context.Context, db *sqlx.DB) error {
	FixtureFirmwareUUIDsR6515 = []string{}

	FixtureDellR6515BIOS = &models.ComponentFirmwareVersion{
		Vendor:        "Dell",
		Model:         types.StringArray{"R6515"},
		Filename:      "BIOS_C4FT0_WN64_2.6.6.EXE",
		Version:       "2.6.6",
		Component:     "bios",
		Checksum:      "1ddcb3c3d0fc5925ef03a3dde768e9e245c579039dd958fc0f3a9c6368b6c5f4",
		UpstreamURL:   "https://vendor.com/firmwares/BIOS_C4FT0_WN64_2.6.6.EXE",
		RepositoryURL: "https://example-firmware-bucket.s3.amazonaws.com/firmware/dell/r6515/bios/BIOS_C4FT0_WN64_2.6.6.EXE",
	}

	if err := FixtureDellR6515BIOS.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	FixtureFirmwareUUIDsR6515 = append(FixtureFirmwareUUIDsR6515, FixtureDellR6515BIOS.ID)

	FixtureDellR6515BMC = &models.ComponentFirmwareVersion{
		Vendor:        "Dell",
		Model:         types.StringArray{"R6515"},
		Filename:      "BMC-5.20.20.20.EXE",
		Version:       "5.20.20.20",
		Component:     "bmc",
		Checksum:      "abccb3c3d0fc5925ef03a3dde768e9e245c579039dd958fc0f3a9c6368b6c5f4",
		UpstreamURL:   "https://vendor.com/firmwares/BMC-5.20.20.20.EXE",
		RepositoryURL: "https://example-firmware-bucket.s3.amazonaws.com/firmware/dell/r6515/bmc/BMC-5.20.20.20.EXE",
	}

	if err := FixtureDellR6515BMC.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	FixtureFirmwareUUIDsR6515 = append(FixtureFirmwareUUIDsR6515, FixtureDellR6515BMC.ID)

	return nil
}

func setupFirmwareSuperMicro(ctx context.Context, db *sqlx.DB) error {
	FixtureFirmwareUUIDsSuperMicro = []string{}

	FixtureSuperMicroX11DPHTBMC = &models.ComponentFirmwareVersion{
		Vendor:        "SuperMicro",
		Model:         types.StringArray{"X11DPH-T"},
		Filename:      "SMT_X11AST2500_173_11.bin",
		Version:       "1.73.11",
		Component:     "bmc",
		Checksum:      "83d220484495e79a3c20e16c21a0d751a71519ac7058350d8a38e1f55efb0211",
		UpstreamURL:   "https://vendor.com/firmwares/SMT_X11AST2500_173_11.bin",
		RepositoryURL: "https://example-firmware-bucket.s3.amazonaws.com/firmware/supermicro/X11DPH-T/bmc/SMT_X11AST2500_173_11.bin",
	}

	if err := FixtureSuperMicroX11DPHTBMC.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	FixtureFirmwareUUIDsSuperMicro = append(FixtureFirmwareUUIDsSuperMicro, FixtureSuperMicroX11DPHTBMC.ID)

	return nil
}

func setupFirmwareInbandNIC(ctx context.Context, db *sqlx.DB) error {
	FixtureFirmwareInbandNIC = &models.ComponentFirmwareVersion{
		Vendor:        "Intel",
		Model:         types.StringArray{"e810"},
		Filename:      "blob.bin",
		Version:       "0.00.7",
		Component:     "nic",
		Checksum:      "83d220484495e79a3c20e16c21a0d751a71519ac7058350d8a38e1f55efb0222",
		UpstreamURL:   "https://vendor.com/firmwares/blob.bin",
		RepositoryURL: "https://example-firmware-bucket.s3.amazonaws.com/firmware/intel/blob.bin",
		InstallInband: true,
		Oem:           true,
	}

	return FixtureFirmwareInbandNIC.Insert(ctx, db, boil.Infer())
}

func setupFirmwareSetSupermicroX11DPHT(ctx context.Context, db *sqlx.DB) error {
	// setup firmware fixtures if they haven't been
	if len(FixtureFirmwareUUIDsSuperMicro) == 0 {
		if err := setupFirmwareSuperMicro(ctx, db); err != nil {
			return err
		}
	}

	FixtureFirmwareSetX11DPHT = &models.ComponentFirmwareSet{Name: "x11dph-t"}

	if err := FixtureFirmwareSetX11DPHT.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	FixtureFirmwareSetX11DPHTAttribute = &models.AttributesFirmwareSet{
		FirmwareSetID: null.StringFrom(FixtureFirmwareSetX11DPHT.ID),
		Namespace:     "sh.hollow.firmware_set.labels",
		Data:          types.JSON([]byte(`{"vendor": "supermicro", "model": "x11dph-t"}`)),
	}

	if err := FixtureFirmwareSetX11DPHT.AddFirmwareSetAttributesFirmwareSets(ctx, db, true, FixtureFirmwareSetX11DPHTAttribute); err != nil {
		return err
	}

	for _, firmwareID := range FixtureFirmwareUUIDsSuperMicro {
		m := &models.ComponentFirmwareSetMap{
			FirmwareSetID: FixtureFirmwareSetX11DPHT.ID,
			FirmwareID:    firmwareID,
		}

		if err := m.Insert(ctx, db, boil.Infer()); err != nil {
			return err
		}
	}

	return nil
}

func setupFirmwareSetR6515(ctx context.Context, db *sqlx.DB) error {
	// setup firmware fixtures if they haven't been
	if len(FixtureFirmwareUUIDsR6515) == 0 {
		if err := setupFirmwareDellR6515(ctx, db); err != nil {
			return err
		}
	}

	FixtureFirmwareSetR6515 = &models.ComponentFirmwareSet{Name: "r6515"}

	if err := FixtureFirmwareSetR6515.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	FixtureFirmwareSetR6515Attribute = &models.AttributesFirmwareSet{
		FirmwareSetID: null.StringFrom(FixtureFirmwareSetR6515.ID),
		Namespace:     "sh.hollow.firmware_set.labels",
		Data:          types.JSON([]byte(`{"vendor": "dell", "model": "r6515"}`)),
	}

	if err := FixtureFirmwareSetR6515.AddFirmwareSetAttributesFirmwareSets(ctx, db, true, FixtureFirmwareSetR6515Attribute); err != nil {
		return err
	}

	for _, firmwareID := range FixtureFirmwareUUIDsR6515 {
		m := &models.ComponentFirmwareSetMap{
			FirmwareSetID: FixtureFirmwareSetR6515.ID,
			FirmwareID:    firmwareID,
		}

		if err := m.Insert(ctx, db, boil.Infer()); err != nil {
			return err
		}
	}

	return nil
}

func setupFirmwareSetR640(ctx context.Context, db *sqlx.DB) error {
	// setup firmware fixtures if they haven't been
	if len(FixtureFirmwareUUIDsR640) == 0 {
		if err := setupFirmwareDellR640(ctx, db); err != nil {
			return err
		}
	}

	FixtureFirmwareSetR640 = &models.ComponentFirmwareSet{Name: "r640"}

	if err := FixtureFirmwareSetR640.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	FixtureFirmwareSetR640Attribute = &models.AttributesFirmwareSet{
		FirmwareSetID: null.StringFrom(FixtureFirmwareSetR640.ID),
		Namespace:     "sh.hollow.firmware_set.labels",
		Data:          types.JSON([]byte(`{"vendor": "dell", "model": "r640", "default": "true"}`)),
	}

	if err := FixtureFirmwareSetR640.AddFirmwareSetAttributesFirmwareSets(ctx, db, true, FixtureFirmwareSetR640Attribute); err != nil {
		return err
	}

	for _, firmwareID := range FixtureFirmwareUUIDsR640 {
		m := &models.ComponentFirmwareSetMap{
			FirmwareSetID: FixtureFirmwareSetR640.ID,
			FirmwareID:    firmwareID,
		}

		if err := m.Insert(ctx, db, boil.Infer()); err != nil {
			return err
		}
	}

	return nil
}

func setupFirmwareSetR640WithLabels(ctx context.Context, db *sqlx.DB, name, labels string) error {
	fixtureFirmwareSetR640 := &models.ComponentFirmwareSet{Name: "r640-" + name}

	if err := fixtureFirmwareSetR640.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	data := `"vendor": "dell", "model": "r640"`
	if labels != "" {
		data += "," + labels
	}

	data = fmt.Sprintf("{%v}", data)
	fixtureFirmwareSetR640Attribute := &models.AttributesFirmwareSet{
		FirmwareSetID: null.StringFrom(fixtureFirmwareSetR640.ID),
		Namespace:     "sh.hollow.firmware_set.labels",
		Data:          types.JSON([]byte(data)),
	}

	if err := fixtureFirmwareSetR640.AddFirmwareSetAttributesFirmwareSets(ctx, db, true, fixtureFirmwareSetR640Attribute); err != nil {
		return err
	}

	return nil
}

func setupInventoryFixture(ctx context.Context, db *sqlx.DB) error {
	FixtureInventoryServer = &models.Server{
		Name:         null.StringFrom("inventory"),
		FacilityCode: null.StringFrom("tf2"),
	}

	if err := FixtureInventoryServer.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	vendorAttrs := map[string]string{
		"model":  "myModel",
		"vendor": "Awesome Computer, Inc.",
		"serial": "1234xyz",
	}
	vaData, _ := json.Marshal(vendorAttrs)

	vendorAttrRecord := &models.Attribute{
		ServerID:  null.StringFrom(FixtureInventoryServer.ID),
		Namespace: "sh.hollow.alloy.server_vendor_attributes", // alloyVendorNamespace
		Data:      types.JSON(vaData),
	}
	if err := vendorAttrRecord.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func setupEventHistoryFixtures(ctx context.Context, db *sqlx.DB) error {
	FixtureEventHistoryServer = &models.Server{
		Name:         null.StringFrom("event-history"),
		FacilityCode: null.StringFrom("tf2"),
	}

	if err := FixtureEventHistoryServer.Insert(ctx, db, boil.Infer()); err != nil {
		return errors.Wrap(err, "event history server fixture")
	}

	FixtureEventHistoryRelatedID = uuid.New()

	FixtureEventHistories = []*models.EventHistory{
		{
			EventID:      uuid.New().String(),
			EventType:    "test event",
			EventStart:   time.Now().Add(-5 * time.Hour),
			EventEnd:     time.Now().Add(-4 * time.Hour),
			TargetServer: FixtureEventHistoryServer.ID,
			Parameters:   null.JSONFrom([]byte(`{"msg": "test event"}`)),
			FinalState:   "succeeded",
			FinalStatus:  null.JSONFrom([]byte(`{"status": "some status"}`)),
		},
		{
			EventID:      uuid.New().String(),
			EventType:    "test event",
			EventStart:   time.Now().Add(-3 * time.Hour),
			EventEnd:     time.Now().Add(-2 * time.Hour),
			TargetServer: FixtureEventHistoryServer.ID,
			Parameters:   null.JSONFrom([]byte(`{"msg": "test event"}`)),
			FinalState:   "failed",
			FinalStatus:  null.JSONFrom([]byte(`{"status": "bad status"}`)),
		},
		{
			EventID:      uuid.New().String(),
			EventType:    "test event",
			EventStart:   time.Now().Add(-1 * time.Hour),
			EventEnd:     time.Now().Add(-30 * time.Minute),
			TargetServer: FixtureEventHistoryServer.ID,
			Parameters:   null.JSONFrom([]byte(`{"msg": "test event"}`)),
			FinalState:   "succeeded",
			FinalStatus:  null.JSONFrom([]byte(`{"status": "some status"}`)),
		},
		{
			EventID:      FixtureEventHistoryRelatedID.String(),
			EventType:    "test event",
			EventStart:   time.Now().Add(-1 * time.Hour),
			EventEnd:     time.Now().Add(-50 * time.Minute),
			TargetServer: FixtureEventHistoryServer.ID,
			Parameters:   null.JSONFrom([]byte(`{"msg": "test event"}`)),
			FinalState:   "succeeded",
			FinalStatus:  null.JSONFrom([]byte(`{"status": "some status"}`)),
		},
		{
			EventID:      FixtureEventHistoryRelatedID.String(),
			EventType:    "test event 2",
			EventStart:   time.Now().Add(-1 * time.Hour),
			EventEnd:     time.Now().Add(-40 * time.Minute),
			TargetServer: FixtureEventHistoryServer.ID,
			Parameters:   null.JSONFrom([]byte(`{"msg": "test event"}`)),
			FinalState:   "succeeded",
			FinalStatus:  null.JSONFrom([]byte(`{"status": "some status"}`)),
		},
		{
			EventID:      FixtureEventHistoryRelatedID.String(),
			EventType:    "test event 3",
			EventStart:   time.Now().Add(-1 * time.Hour),
			EventEnd:     time.Now().Add(-50 * time.Minute),
			TargetServer: FixtureEventHistoryServer.ID,
			Parameters:   null.JSONFrom([]byte(`{"msg": "test event"}`)),
			FinalState:   "succeeded",
			FinalStatus:  null.JSONFrom([]byte(`{"status": "some status"}`)),
		},
	}

	for idx, evt := range FixtureEventHistories {
		if err := evt.Insert(ctx, db, boil.Infer()); err != nil {
			errStr := fmt.Sprintf("event history %d", idx)
			return errors.Wrap(err, errStr)
		}
	}
	return nil
}

func setupConfigSet(ctx context.Context, db *sqlx.DB) error {
	settings := [][]*models.BiosConfigSetting{
		{
			{
				SettingsKey:   "BootOrder",
				SettingsValue: "dev2,dev3,dev4",
			},
			{
				SettingsKey:   "Mode",
				SettingsValue: "UEFI",
			},
		},
		{
			{
				SettingsKey:   "PXEEnable",
				SettingsValue: "true",
				Raw:           null.NewJSON([]byte(`{}`), true),
			},
			{
				SettingsKey:   "SRIOVEnable",
				SettingsValue: "false",
			},
			{
				SettingsKey:   "position",
				SettingsValue: "1",
				Raw:           null.NewJSON([]byte(`{ "lanes": 8 }`), true),
			},
		},
	}

	components := []*models.BiosConfigComponent{
		{
			Name:   "Fixture Test SM Motherboard",
			Vendor: "SUPERMICRO",
			Model:  "ATX",
		},
		{
			Name:   "Fixture Test Intel Network Adapter",
			Vendor: "Intel",
			Model:  "PCIE",
		},
	}

	configSet := models.BiosConfigSet{
		Name:    "Fixture Test Config Set",
		Version: "version",
	}

	err := configSet.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}

	for c := range components {
		components[c].FKBiosConfigSetID = configSet.ID

		err = configSet.AddFKBiosConfigSetBiosConfigComponents(ctx, db, true, components[c])
		if err != nil {
			return err
		}

		for s := range settings[c] {
			settings[c][s].FKBiosConfigComponentID = components[c].ID
		}

		err = components[c].AddFKBiosConfigComponentBiosConfigSettings(ctx, db, true, settings[c]...)
		if err != nil {
			return err
		}
	}

	FixtureBiosConfigSet = &configSet
	FixtureBiosConfigComponents = components
	FixtureBiosConfigSettings = settings

	return nil
}

func setupFWValidationFixtures(ctx context.Context, db *sqlx.DB) error {
	FixtureFWValidationServer = &models.Server{
		Name:         null.StringFrom("firmware-validation"),
		FacilityCode: null.StringFrom("tf2"),
	}

	if err := FixtureFWValidationServer.Insert(ctx, db, boil.Infer()); err != nil {
		return errors.Wrap(err, "firmware validation server fixture")
	}

	FixtureFWValidationSet = &models.ComponentFirmwareSet{Name: "firmware-validation"}
	if err := FixtureFWValidationSet.Insert(ctx, db, boil.Infer()); err != nil {
		return errors.Wrap(err, "firmware validation set fixture")
	}

	return nil
}

func setupHardwareVendorFixtures(ctx context.Context, db *sqlx.DB) error {
	FixtureHardwareVendorBar = &models.HardwareVendor{Name: FixtureHardwareVendorNameBar}
	FixtureHardwareVendorBaz = &models.HardwareVendor{Name: FixtureHardwareVendorNameBaz}

	// this fixture is deleted in tests
	FixtureHardwareVendorFoo = &models.HardwareVendor{Name: FixtureHardwareVendorNameFoo}

	FixtureHardwareVendors = []*models.HardwareVendor{
		FixtureHardwareVendorBar,
		FixtureHardwareVendorBaz,
		FixtureHardwareVendorFoo,
	}

	for _, hv := range FixtureHardwareVendors {
		if err := hv.Insert(ctx, db, boil.Infer()); err != nil {
			return errors.Wrap(err, "hardware vendor insert fixture")
		}
	}

	return nil
}

func setupHardwareModelFixtures(ctx context.Context, db *sqlx.DB) error {
	FixtureHardwareModelBar123 = &models.HardwareModel{
		Name:             FixtureHardwareModelBaz123Name,
		HardwareVendorID: FixtureHardwareVendorBaz.ID,
	}

	FixtureHardwareModelBaz123 = &models.HardwareModel{
		Name:             FixtureHardwareModelBar456Name,
		HardwareVendorID: FixtureHardwareVendorBar.ID,
	}

	// this fixture is deleted in tests
	FixtureHardwareModelFoo789 = &models.HardwareModel{
		Name:             FixtureHardwareModelFoo789Name,
		HardwareVendorID: FixtureHardwareVendorFoo.ID,
	}

	FixtureHardwareModels = []*models.HardwareModel{
		FixtureHardwareModelBar123,
		FixtureHardwareModelBaz123,
		FixtureHardwareModelFoo789,
	}

	for _, hm := range FixtureHardwareModels {
		if err := hm.Insert(ctx, db, boil.Infer()); err != nil {
			return errors.Wrap(err, "hardware model insert fixture")
		}
	}

	return nil
}

func setupServerBMCFixtures(ctx context.Context, db *sqlx.DB) error {
	FixtureServerBMCs = []*models.ServerBMC{
		{
			ServerID:         FixtureNemo.ID,
			HardwareVendorID: FixtureHardwareVendorBar.ID,
			HardwareModelID:  FixtureHardwareModelBar123.ID,
			Username:         "user",
			IPAddress:        "127.0.0.1",
			MacAddress:       null.StringFrom("de:ad:be:ef:ca:fe"),
		},
		{
			ServerID:         FixtureDory.ID,
			HardwareVendorID: FixtureHardwareVendorBaz.ID,
			HardwareModelID:  FixtureHardwareModelBaz123.ID,
			Username:         "user",
			IPAddress:        "127.0.0.2",
			MacAddress:       null.StringFrom("de:ad:be:ef:ef:fe"),
		},
	}

	for _, serverBMC := range FixtureServerBMCs {
		if err := serverBMC.Insert(ctx, db, boil.Infer()); err != nil {
			return errors.Wrap(err, "server BMC insert fixture")
		}
	}

	return nil
}
