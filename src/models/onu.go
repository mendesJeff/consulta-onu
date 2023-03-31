package models

import (
	"errors"
	"strings"
)

type Onu struct {
	// ID          uint64 `json:"id,omitempty"`
	Olt_name    string `json:"olt_name,omitempty"`
	Olt_ip      string `json:"olt_ip,omitempty"`
	Slot_number uint64 `json:"slot_number,omitempty"`
	Pon_number  uint64 `json:"pon_number,omitempty"`
	Onu_number  uint64 `json:"onu_number,omitempty"`
	Onu_serial  string `json:"onu_serial,omitempty"`
}

func (onu *Onu) Prepare() error {
	if error := onu.validate(); error != nil {
		return error
	}

	onu.formate()
	return nil
}

func (onu *Onu) validate() error {
	switch {
	case onu.Onu_serial == "":
		return errors.New("onu serial invalid")

	}
	return nil
}

func (onu *Onu) formate() {
	onu.Olt_name = strings.TrimSpace(onu.Olt_name)
	onu.Olt_ip = strings.TrimSpace(onu.Olt_ip)
	onu.Onu_serial = strings.TrimSpace(onu.Onu_serial)
}
