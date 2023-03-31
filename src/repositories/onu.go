package repositories

import (
	"database/sql"
	"database/src/models"
)

type onus struct {
	db *sql.DB
}

func NewOnusRepository(db *sql.DB) *onus {
	return &onus{db}
}

func (repository onus) SearchBySerial(onuSerial string) ([]models.Onu, error) {

	lines, error := repository.db.Query(
		"SELECT OLT.cobjectname, OLT.cipaddress, cslotno, cponno, cauthno, contmacquery FROM t_ontdevice ONU JOIN t_nedevice OLT ON ONU.cneid = OLT.cobjectid WHERE ONU.contmacquery = ?", onuSerial)

	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var onus []models.Onu
	for lines.Next() {
		var onu models.Onu

		if error = lines.Scan(
			&onu.Olt_name,
			&onu.Olt_ip,
			&onu.Slot_number,
			&onu.Pon_number,
			&onu.Onu_number,
			&onu.Onu_serial,
		); error != nil {
			return nil, error
		}

		onus = append(onus, onu)
	}

	return onus, nil
}

func (repository onus) SearchByOlt(oltIpOrOltName string) ([]models.Onu, error) {

	lines, error := repository.db.Query(
		"SELECT OLT.cobjectname, OLT.cipaddress, cslotno, cponno, cauthno, contmacquery FROM t_ontdevice ONU JOIN t_nedevice OLT ON ONU.cneid = OLT.cobjectid WHERE LOWER(OLT.cobjectname) = ? OR OLT.cipaddress = ?", oltIpOrOltName, oltIpOrOltName)

	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var onus []models.Onu
	for lines.Next() {
		var onu models.Onu

		if error = lines.Scan(
			&onu.Olt_name,
			&onu.Olt_ip,
			&onu.Slot_number,
			&onu.Pon_number,
			&onu.Onu_number,
			&onu.Onu_serial,
		); error != nil {
			return nil, error
		}

		onus = append(onus, onu)
	}

	return onus, nil
}
