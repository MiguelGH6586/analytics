package models

type BykeGeneralInfo struct {
	Brand     string  `bson:"brand,omitempty" validate:"required"`
	Model     string  `bson:"model,omitempty" validate:"required"`
	FullName  string  `bson:"full_name,omitempty" validate:"required"`
	YearModel string  `bson:"year_model,omitempty" validate:"required"`
	Weight    string  `bson:"weight,omitempty" validate:"required"`
	Engine    string  `bson:"engine,omitempty" validate:"required"`
	Cylinder  string  `bson:"cylinder,omitempty" validate:"required"`
	Prices    []Price `bson:"prices,omitempty" validate:"required"`
}

type Price struct {
	DateFound string `bson:"date_found,omitempty" validate:"required"`
	Km        string `bson:"km,omitempty" validate:"required"`
	Price     string `bson:"price,omitempty" validate:"required"`
}
