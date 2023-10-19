package models

type Auto struct {
	Codigo string  `json:"codigo"`
	Marca  string  `json:"marca"`
	Modelo string  `json:"modelo"`
	Precio float64 `json:"precio"`
}
