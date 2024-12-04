package fixtures

import (
	"github.com/bmc-toolbox/common"
	"github.com/jinzhu/copier"
)

// CopyDevice returns a pointer to a copy of the given ironlib device object
func CopyDevice(src *common.Device) *common.Device {
	dst := &common.Device{}

	copyOptions := copier.Option{IgnoreEmpty: true, DeepCopy: true}

	err := copier.CopyWithOption(&dst, &src, copyOptions)
	if err != nil {
		panic(err)
	}

	return dst
}

// nolint:dupl,misspell,revive,stylecheck
// r6515 inventory taken from ironlib

var (
	DellR6515 = &common.Device{
		Common: common.Common{
			Oem:         false,
			Description: "",
			Vendor:      "dell",
			Model:       "r6515",
			Serial:      "11WLK93",
			ProductName: "",
			Firmware:    nil,
			Status:      nil,
			Metadata:    map[string]string(nil), // p0
		},
		HardwareType: "",
		Chassis:      "",
		BIOS: &common.BIOS{
			Common: common.Common{
				Oem:         false,
				Description: "BIOS",
				Vendor:      "Dell Inc.",
				Model:       "r6515",
				Serial:      "",
				ProductName: "",
				Firmware: &common.Firmware{
					Installed:  "1.7.4",
					Available:  "",
					SoftwareID: "",
					Previous:   nil,
				},
				Status: nil,
				Capabilities: []*common.Capability{
					{
						Name:        "acpi",
						Description: "ACPI",
						Enabled:     true,
					},
					{
						Name:        "biosbootspecification",
						Description: "BIOS boot specification",
						Enabled:     true,
					},
					{
						Name:        "bootselect",
						Description: "Selectable boot path",
						Enabled:     true,
					},
					{
						Name:        "cdboot",
						Description: "Booting from CD-ROM/DVD",
						Enabled:     true,
					},
					{
						Name:        "edd",
						Description: "Enhanced Disk Drive extensions",
						Enabled:     true,
					},
					{
						Name:        "int10video",
						Description: "INT10 CGA/Mono video",
						Enabled:     true,
					},
					{
						Name:        "int13floppy1200",
						Description: "5.25\" 1.2MB floppy",
						Enabled:     true,
					},
					{
						Name:        "int13floppy360",
						Description: "5.25\" 360KB floppy",
						Enabled:     true,
					},
					{
						Name:        "int13floppy720",
						Description: "3.5\" 720KB floppy",
						Enabled:     true,
					},
					{
						Name:        "int13floppytoshiba",
						Description: "Toshiba floppy",
						Enabled:     true,
					},
					{
						Name:        "int14serial",
						Description: "INT14 serial line control",
						Enabled:     true,
					},
					{
						Name:        "int9keyboard",
						Description: "i8042 keyboard controller",
						Enabled:     true,
					},
					{
						Name:        "isa",
						Description: "ISA bus",
						Enabled:     true,
					},
					{
						Name:        "netboot",
						Description: "Function-key initiated network service boot",
						Enabled:     true,
					},
					{
						Name:        "pci",
						Description: "PCI bus",
						Enabled:     true,
					},
					{
						Name:        "pnp",
						Description: "Plug-and-Play",
						Enabled:     true,
					},
					{
						Name:        "shadowing",
						Description: "BIOS shadowing",
						Enabled:     true,
					},
					{
						Name:        "uefi",
						Description: "UEFI specification is supported",
						Enabled:     true,
					},
					{
						Name:        "upgrade",
						Description: "BIOS EEPROM can be upgraded",
						Enabled:     true,
					},
					{
						Name:        "usb",
						Description: "USB legacy emulation",
						Enabled:     true,
					},
				},
			},
			SizeBytes:     65536,
			CapacityBytes: 33554432,
		},
		BMC: &common.BMC{
			Common: common.Common{
				Oem:         false,
				Description: "",
				Vendor:      "",
				Model:       "r6515",
				Serial:      "",
				ProductName: "",
				Firmware:    nil,
				Status:      nil,
			},
			ID: "",
			NIC: &common.NIC{
				Common: common.Common{
					Oem:         false,
					Description: "",
					Vendor:      "",
					Model:       "",
					Serial:      "",
					ProductName: "",
					Firmware:    nil,
					Status:      nil,
				},
				ID:       "",
				NICPorts: nil,
			},
		},
		Mainboard: &common.Mainboard{
			Common: common.Common{
				Oem:         false,
				Description: "Motherboard",
				Vendor:      "Dell Inc.",
				Model:       "0R4CNN",
				Serial:      ".11WLK93.CNCMS0009G0078.",
				ProductName: "0R4CNN",
				Firmware:    nil,
				Status:      nil,
			},
			PhysicalID: "0",
		},
		CPLDs: []*common.CPLD{},
		TPMs:  []*common.TPM{},
		GPUs:  []*common.GPU{},
		CPUs: []*common.CPU{
			{
				Common: common.Common{
					Oem:         false,
					Description: "CPU",
					Vendor:      "Advanced Micro Devices [AMD]",
					Model:       "AMD EPYC 7502P 32-Core Processor",
					Serial:      "",
					ProductName: "AMD EPYC 7502P 32-Core Processor",
					Firmware:    &common.Firmware{Installed: "137367629", Metadata: map[string]string{}},
					Status:      nil,
					Capabilities: []*common.Capability{
						{
							Name:        "3dnowprefetch",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "abm",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "adx",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "aes",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "aperfmperf",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "apic",
							Description: "on-chip advanced programmable interrupt controller (APIC)",
							Enabled:     true,
						},
						{
							Name:        "arat",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "avic",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "avx",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "avx2",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "bmi1",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "bmi2",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "bpext",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "cat_l3",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "cdp_l3",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "clflush",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "clflushopt",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "clwb",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "clzero",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "cmov",
							Description: "conditional move instruction",
							Enabled:     true,
						},
						{
							Name:        "cmp_legacy",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "constant_tsc",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "cpb",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "cpuid",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "cqm",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "cqm_llc",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "cqm_mbm_local",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "cqm_mbm_total",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "cqm_occup_llc",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "cr8_legacy",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "cx16",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "cx8",
							Description: "compare and exchange 8-byte",
							Enabled:     true,
						},
						{
							Name:        "de",
							Description: "debugging extensions",
							Enabled:     true,
						},
						{
							Name:        "decodeassists",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "extapic",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "extd_apicid",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "f16c",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "flushbyasid",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "fma",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "fpu",
							Description: "mathematical co-processor",
							Enabled:     true,
						},
						{
							Name:        "fpu_exception",
							Description: "FPU exceptions reporting",
							Enabled:     true,
						},
						{
							Name:        "fsgsbase",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "fxsr",
							Description: "fast floating point save/restore",
							Enabled:     true,
						},
						{
							Name:        "fxsr_opt",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "ht",
							Description: "HyperThreading",
							Enabled:     true,
						},
						{
							Name:        "hw_pstate",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "ibpb",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "ibrs",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "ibs",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "irperf",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "lahf_lm",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "lbrv",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "lm",
							Description: "64bits extensions (x86-64)",
							Enabled:     true,
						},
						{
							Name:        "mba",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "mca",
							Description: "machine check architecture",
							Enabled:     true,
						},
						{
							Name:        "mce",
							Description: "machine check exceptions",
							Enabled:     true,
						},
						{
							Name:        "misalignsse",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "mmx",
							Description: "multimedia extensions (MMX)",
							Enabled:     true,
						},
						{
							Name:        "mmxext",
							Description: "multimedia extensions (MMXExt)",
							Enabled:     true,
						},
						{
							Name:        "monitor",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "movbe",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "msr",
							Description: "model-specific registers",
							Enabled:     true,
						},
						{
							Name:        "mtrr",
							Description: "memory type range registers",
							Enabled:     true,
						},
						{
							Name:        "mwaitx",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "nonstop_tsc",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "nopl",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "npt",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "nrip_save",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "nx",
							Description: "no-execute bit (NX)",
							Enabled:     true,
						},
						{
							Name:        "osvw",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "overflow_recov",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "pae",
							Description: "4GB+ memory addressing (Physical Address Extension)",
							Enabled:     true,
						},
						{
							Name:        "pat",
							Description: "page attribute table",
							Enabled:     true,
						},
						{
							Name:        "pausefilter",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "pclmulqdq",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "pdpe1gb",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "perfctr_core",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "perfctr_llc",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "perfctr_nb",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "pfthreshold",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "pge",
							Description: "page global enable",
							Enabled:     true,
						},
						{
							Name:        "pni",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "popcnt",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "pse",
							Description: "page size extensions",
							Enabled:     true,
						},
						{
							Name:        "pse36",
							Description: "36-bit page size extensions",
							Enabled:     true,
						},
						{
							Name:        "rdpid",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "rdrand",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "rdseed",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "rdt_a",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "rdtscp",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "rep_good",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "sep",
							Description: "fast system calls",
							Enabled:     true,
						},
						{
							Name:        "sev",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "sha_ni",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "skinit",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "smap",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "smca",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "sme",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "smep",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "ssbd",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "sse",
							Description: "streaming SIMD extensions (SSE)",
							Enabled:     true,
						},
						{
							Name:        "sse2",
							Description: "streaming SIMD extensions (SSE2)",
							Enabled:     true,
						},
						{
							Name:        "sse4_1",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "sse4_2",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "sse4a",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "ssse3",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "stibp",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "succor",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "svm",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "svm_lock",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "syscall",
							Description: "fast system calls",
							Enabled:     true,
						},
						{
							Name:        "tce",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "topoext",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "tsc",
							Description: "time stamp counter",
							Enabled:     true,
						},
						{
							Name:        "tsc_scale",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "umip",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "v_vmsave_vmload",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "vgif",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "vmcb_clean",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "vme",
							Description: "virtual mode extensions",
							Enabled:     true,
						},
						{
							Name:        "vmmcall",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "wbnoinvd",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "wdt",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "wp",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "x2apic",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "x86-64",
							Description: "64bits extensions (x86-64)",
							Enabled:     true,
						},
						{
							Name:        "xgetbv1",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "xsave",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "xsavec",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "xsaveerptr",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "xsaveopt",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "xsaves",
							Description: "",
							Enabled:     true,
						},
					},
				},
				ID:           "",
				Slot:         "CPU1",
				Architecture: "",
				ClockSpeedHz: 2000000000,
				Cores:        32,
				Threads:      64,
			},
		},
		Memory: []*common.Memory{
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 3200 MHz (0.3 ns)",
					Vendor:      "Hynix Semiconductor (Hyundai Electronics)",
					Model:       "HMA84GR7DJR4N-XN",
					Serial:      "3510CB7F",
					ProductName: "HMA84GR7DJR4N-XN",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A1",
				Type:         "",
				SizeBytes:    34359738368,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 3200000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 3200 MHz (0.3 ns)",
					Vendor:      "Hynix Semiconductor (Hyundai Electronics)",
					Model:       "HMA84GR7DJR4N-XN",
					Serial:      "3510CBDD",
					ProductName: "HMA84GR7DJR4N-XN",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A2",
				Type:         "",
				SizeBytes:    34359738368,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 3200000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 3200 MHz (0.3 ns)",
					Vendor:      "Hynix Semiconductor (Hyundai Electronics)",
					Model:       "HMA84GR7DJR4N-XN",
					Serial:      "3510CC97",
					ProductName: "HMA84GR7DJR4N-XN",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A3",
				Type:         "",
				SizeBytes:    34359738368,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 3200000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 3200 MHz (0.3 ns)",
					Vendor:      "Hynix Semiconductor (Hyundai Electronics)",
					Model:       "HMA84GR7DJR4N-XN",
					Serial:      "3510CC8C",
					ProductName: "HMA84GR7DJR4N-XN",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A4",
				Type:         "",
				SizeBytes:    34359738368,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 3200000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 3200 MHz (0.3 ns)",
					Vendor:      "Hynix Semiconductor (Hyundai Electronics)",
					Model:       "HMA84GR7DJR4N-XN",
					Serial:      "3510CBDE",
					ProductName: "HMA84GR7DJR4N-XN",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A5",
				Type:         "",
				SizeBytes:    34359738368,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 3200000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 3200 MHz (0.3 ns)",
					Vendor:      "Hynix Semiconductor (Hyundai Electronics)",
					Model:       "HMA84GR7DJR4N-XN",
					Serial:      "3510CB77",
					ProductName: "HMA84GR7DJR4N-XN",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A6",
				Type:         "",
				SizeBytes:    34359738368,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 3200000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 3200 MHz (0.3 ns)",
					Vendor:      "Hynix Semiconductor (Hyundai Electronics)",
					Model:       "HMA84GR7DJR4N-XN",
					Serial:      "3510CC78",
					ProductName: "HMA84GR7DJR4N-XN",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A7",
				Type:         "",
				SizeBytes:    34359738368,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 3200000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "DIMM DDR4 Synchronous Registered (Buffered) 3200 MHz (0.3 ns)",
					Vendor:      "Hynix Semiconductor (Hyundai Electronics)",
					Model:       "HMA84GR7DJR4N-XN",
					Serial:      "3510CC93",
					ProductName: "HMA84GR7DJR4N-XN",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A8",
				Type:         "",
				SizeBytes:    34359738368,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 3200000000,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "[empty]",
					Vendor:      "",
					Model:       "",
					Serial:      "",
					ProductName: "",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A9",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "[empty]",
					Vendor:      "",
					Model:       "",
					Serial:      "",
					ProductName: "",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A10",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "[empty]",
					Vendor:      "",
					Model:       "",
					Serial:      "",
					ProductName: "",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A11",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "[empty]",
					Vendor:      "",
					Model:       "",
					Serial:      "",
					ProductName: "",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A12",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "[empty]",
					Vendor:      "",
					Model:       "",
					Serial:      "",
					ProductName: "",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A13",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "[empty]",
					Vendor:      "",
					Model:       "",
					Serial:      "",
					ProductName: "",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A14",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "[empty]",
					Vendor:      "",
					Model:       "",
					Serial:      "",
					ProductName: "",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A15",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 0,
			},
			{
				Common: common.Common{
					Oem:         false,
					Description: "[empty]",
					Vendor:      "",
					Model:       "",
					Serial:      "",
					ProductName: "",
					Firmware:    nil,
					Status:      nil,
				},
				ID:           "",
				Slot:         "A16",
				Type:         "",
				SizeBytes:    0,
				FormFactor:   "",
				PartNumber:   "",
				ClockSpeedHz: 0,
			},
		},
		NICs: []*common.NIC{
			{
				Common: common.Common{
					Oem:         false,
					Description: "Ethernet interface",
					Vendor:      "Intel Corporation",
					Model:       "Ethernet Controller XXV710 for 25GbE SFP28",
					Serial:      "40:a6:b7:4e:8a:a0",
					ProductName: "Ethernet Controller XXV710 for 25GbE SFP28",
					Firmware: &common.Firmware{
						Installed:  "20.0.17",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
					Metadata: map[string]string{
						"driver":   "i40e",
						"duplex":   "full",
						"firmware": "8.15 0x800096ca 20.0.17",
						"link":     "yes",
					},
					PCIVendorID:  "dead",
					PCIProductID: "beef",
					Capabilities: []*common.Capability{
						{
							Name:        "autonegotiation",
							Description: "Auto-negotiation",
							Enabled:     true,
						},
						{
							Name:        "bus_master",
							Description: "bus mastering",
							Enabled:     true,
						},
						{
							Name:        "cap_list",
							Description: "PCI capabilities listing",
							Enabled:     true,
						},
						{
							Name:        "ethernet",
							Description: "",
							Enabled:     true,
						},
						{
							Name:        "fibre",
							Description: "optical fibre",
							Enabled:     true,
						},
						{
							Name:        "msi",
							Description: "Message Signalled Interrupts",
							Enabled:     true,
						},
						{
							Name:        "msix",
							Description: "MSI-X",
							Enabled:     true,
						},
						{
							Name:        "pciexpress",
							Description: "PCI Express",
							Enabled:     true,
						},
						{
							Name:        "physical",
							Description: "Physical interface",
							Enabled:     true,
						},
						{
							Name:        "pm",
							Description: "Power Management",
							Enabled:     true,
						},
						{
							Name:        "rom",
							Description: "extension ROM",
							Enabled:     true,
						},
						{
							Name:        "vpd",
							Description: "Vital Product Data",
							Enabled:     true,
						},
					},
				},
				ID: "",
				NICPorts: []*common.NICPort{
					{
						SpeedBits:  0,
						PhysicalID: "0",
						BusInfo:    "pci@0000:41:00.0",
						MacAddress: "",
					},
				},
			},
		},
		Drives: []*common.Drive{
			{
				Common: common.Common{
					Oem:          false,
					Description:  "NVMe device",
					Vendor:       "micron",
					Model:        "Micron_9300_MTFDHAL3T8TDP",
					Serial:       "202728F691F5",
					ProductName:  "Micron_9300_MTFDHAL3T8TDP",
					LogicalName:  "/dev/nvme0",
					Capabilities: nvmeDriveCapabilities,
					Firmware: &common.Firmware{
						Installed:  "11300DN0",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                       "",
				OemID:                    "",
				Type:                     "NVMe-PCIe-SSD",
				StorageController:        "",
				StorageControllerDriveID: -1,
				BusInfo:                  "pci@0000:45:00.0",
				WWN:                      "",
				Protocol:                 "nvme",
				SmartStatus:              "ok",
				SmartErrors:              nil,
				CapacityBytes:            0,
				BlockSizeBytes:           0,
				CapableSpeedGbps:         0,
				NegotiatedSpeedGbps:      0,
			},
			{
				Common: common.Common{
					Oem:          false,
					Description:  "NVMe device",
					Vendor:       "micron",
					Model:        "Micron_9300_MTFDHAL3T8TDP",
					Serial:       "202728F691C6",
					ProductName:  "Micron_9300_MTFDHAL3T8TDP",
					LogicalName:  "/dev/nvme1",
					Capabilities: nvmeDriveCapabilities,
					Firmware: &common.Firmware{
						Installed:  "11300DN0",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                       "",
				OemID:                    "",
				Type:                     "NVMe-PCIe-SSD",
				StorageController:        "",
				StorageControllerDriveID: -1,
				BusInfo:                  "pci@0000:46:00.0",
				WWN:                      "",
				Protocol:                 "nvme",
				SmartStatus:              "ok",
				SmartErrors:              nil,
				CapacityBytes:            0,
				BlockSizeBytes:           0,
				CapableSpeedGbps:         0,
				NegotiatedSpeedGbps:      0,
			},
			{
				Common: common.Common{
					Oem:          false,
					Description:  "ATA Disk",
					Vendor:       "micron",
					Model:        "MTFDDAV240TDU",
					Serial:       "203329F89392",
					ProductName:  "MTFDDAV240TDU",
					LogicalName:  "/dev/sda",
					Capabilities: hddCapabilities,
					Firmware: &common.Firmware{
						Installed:  "D3DJ004",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                       "",
				OemID:                    "DELL(tm)",
				Type:                     "Sata-SSD",
				StorageController:        "",
				StorageControllerDriveID: -1,
				BusInfo:                  "scsi@10:0.0.0",
				WWN:                      "",
				Protocol:                 "sata",
				SmartStatus:              "ok",
				SmartErrors:              nil,
				CapacityBytes:            240057409536,
				BlockSizeBytes:           0,
				CapableSpeedGbps:         0,
				NegotiatedSpeedGbps:      0,
			},
			{
				Common: common.Common{
					Oem:          false,
					Description:  "ATA Disk",
					Vendor:       "micron",
					Model:        "MTFDDAV240TDU",
					Serial:       "203329F89796",
					ProductName:  "MTFDDAV240TDU",
					LogicalName:  "/dev/sdb",
					Capabilities: hddCapabilities,
					Firmware: &common.Firmware{
						Installed:  "D3DJ004",
						Available:  "",
						SoftwareID: "",
						Previous:   nil,
					},
					Status: nil,
				},
				ID:                       "",
				OemID:                    "DELL(tm)",
				Type:                     "Sata-SSD",
				StorageController:        "",
				StorageControllerDriveID: -1,
				BusInfo:                  "scsi@11:0.0.0",
				WWN:                      "",
				Protocol:                 "sata",
				SmartStatus:              "ok",
				SmartErrors:              nil,
				CapacityBytes:            240057409536,
				BlockSizeBytes:           0,
				CapableSpeedGbps:         0,
				NegotiatedSpeedGbps:      0,
			},
		},
		StorageControllers: []*common.StorageController{
			{
				Common: common.Common{
					Oem:          false,
					Description:  "Serial Attached SCSI controller",
					Vendor:       "broadcom",
					Model:        "SAS3008 PCI-Express Fusion-MPT SAS-3",
					Serial:       "dead:beef",
					ProductName:  "SAS3008 PCI-Express Fusion-MPT SAS-3",
					Firmware:     nil,
					Status:       nil,
					PCIVendorID:  "dead",
					PCIProductID: "beef",
				},
				ID:                           "",
				SupportedControllerProtocols: "",
				SupportedDeviceProtocols:     "SAS",
				SupportedRAIDTypes:           "",
				PhysicalID:                   "0",
				BusInfo:                      "pci@0000:01:00.0",
				SpeedGbps:                    0,
			},
			{
				Common: common.Common{
					Oem:          false,
					Description:  "SATA controller",
					Vendor:       "Advanced Micro Devices, Inc. [AMD]",
					Model:        "FCH SATA Controller [AHCI mode]",
					Serial:       "dead:beef",
					ProductName:  "FCH SATA Controller [AHCI mode]",
					Firmware:     nil,
					Status:       nil,
					PCIVendorID:  "dead",
					PCIProductID: "beef",
				},
				ID:                           "",
				SupportedControllerProtocols: "",
				SupportedDeviceProtocols:     "SATA",
				SupportedRAIDTypes:           "",
				PhysicalID:                   "0",
				BusInfo:                      "pci@0000:85:00.0",
				SpeedGbps:                    0,
			},
			{
				Common: common.Common{
					Oem:          false,
					Description:  "SATA controller",
					Vendor:       "Advanced Micro Devices, Inc. [AMD]",
					Model:        "FCH SATA Controller [AHCI mode]",
					Serial:       "dead:beef",
					ProductName:  "FCH SATA Controller [AHCI mode]",
					Firmware:     nil,
					Status:       nil,
					PCIVendorID:  "dead",
					PCIProductID: "beef",
				},
				ID:                           "",
				SupportedControllerProtocols: "",
				SupportedDeviceProtocols:     "SATA",
				SupportedRAIDTypes:           "",
				PhysicalID:                   "0",
				BusInfo:                      "pci@0000:86:00.0",
				SpeedGbps:                    0,
			},
			{
				Common: common.Common{
					Oem:          false,
					Description:  "SATA controller",
					Vendor:       "marvell",
					Model:        "88SE9230 PCIe SATA 6Gb/s Controller",
					Serial:       common.VendorMarvellPciID + ":beef",
					ProductName:  "88SE9230 PCIe SATA 6Gb/s Controller",
					Firmware:     nil,
					Status:       nil,
					PCIVendorID:  common.VendorMarvellPciID,
					PCIProductID: "beef",
				},
				ID:                           "",
				SupportedControllerProtocols: "",
				SupportedDeviceProtocols:     "SATA",
				SupportedRAIDTypes:           "",
				PhysicalID:                   "0",
				BusInfo:                      "pci@0000:c3:00.0",
				SpeedGbps:                    0,
			},
		},
		PSUs: []*common.PSU{
			{
				Common: common.Common{
					Description: "PWR SPLY,550W,RDNT,LTON",
					Vendor:      "DELL",
					Model:       "0NCNFFA02",
					Serial:      "CNLOD000B229D6",
					ProductName: "0NCNFFA02",
				},
				ID:                 "1",
				PowerCapacityWatts: 550,
			},
			{
				Common: common.Common{
					Description: "PWR SPLY,550W,RDNT,LTON",
					Vendor:      "DELL",
					Model:       "0NCNFFA02",
					Serial:      "CNLOD000B232CD",
					ProductName: "0NCNFFA02",
				},
				ID:                 "2",
				PowerCapacityWatts: 550,
			},
		},
		Enclosures: []*common.Enclosure{},
	}
)