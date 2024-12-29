package models

type DatabaseRecord struct {
	Id   int
	Name string
	Data struct {
		Complex1 string
		Complex2 string
		Complex3 string
	}
}
