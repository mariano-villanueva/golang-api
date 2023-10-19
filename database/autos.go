package database

import (
	"api/models"
	"strconv"
)

func PostAuto(nuevoAuto models.Auto) (models.Auto, error) {
	sqlQuery := "INSERT INTO autos ( marca, modelo, precio) VALUES (?,?,?)"
	res, err := DB.Exec(sqlQuery, nuevoAuto.Marca, nuevoAuto.Modelo, nuevoAuto.Precio)

	if err != nil {
		return models.Auto{}, err
	}

	nuevoCodigo, _ := res.LastInsertId()
	nuevoAuto.Codigo = strconv.FormatInt(nuevoCodigo, 10)
	return nuevoAuto, nil
}

func UpdateAuto(nuevosDatos models.Auto) (models.Auto, error) {
	sqlQuery := "UPDATE autos SET marca = ?, modelo = ?, precio = ? WHERE codigo = ?"
	res, err := DB.Exec(sqlQuery, nuevosDatos.Marca, nuevosDatos.Modelo, nuevosDatos.Precio, nuevosDatos.Codigo)

	if err != nil {
		return models.Auto{}, err
	}

	if filas, _ := res.RowsAffected(); filas == 0 {
		return models.Auto{}, nil
	}

	return nuevosDatos, nil
}

func GetAutos() ([]models.Auto, error) {
	var autos []models.Auto
	sqlQuery := "SELECT * FROM autos"
	regs, err := DB.Query(sqlQuery)

	if err != nil {
		return autos, err
	}

	for regs.Next() {
		var autoActual models.Auto
		regs.Scan(&autoActual.Codigo, &autoActual.Marca, &autoActual.Modelo, &autoActual.Precio)
		autos = append(autos, autoActual)
	}

	return autos, nil
}

func GetAutoPorCodigo(codigo string) models.Auto {
	sqlQuery := "SELECT * FROM autos WHERE codigo = ?"
	regs, _ := DB.Query(sqlQuery, codigo)
	var autoActual models.Auto
	for regs.Next() {
		regs.Scan(&autoActual.Codigo, &autoActual.Marca, &autoActual.Modelo, &autoActual.Precio)
	}

	return autoActual
}

func DeleteAuto(codigo string) {
	sqlQuery := "DELETE FROM autos WHERE codigo = ?"
	DB.Exec(sqlQuery, codigo)
}
