package types

type PlaceDetails struct {
	Name         string       `json:"displayed_what"`
	Address      string       `json:"displayed_where"`
	OpeningHours OpeningHours `json:"opening_hours"`
}

type OpeningHours struct {
	Days map[string][]WorkingBlock `json:"days"`
}

type WorkingBlock struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Type  string `json:"type"` // TODO: Use enums
}

type PlaceDetailsOut struct {
	Name         string             `json:"name"`
	Address      string             `json:"address"`
	IsOpen       bool               `json:"isOpen"`
	NextChange   string             `json:"nextChange"`
	OpeningHours []*OpeningHoursOut `json:"openingHours"`
}

type OpeningHoursOut struct {
	Days          []string
	WorkingBlocks []WorkingBlock
}
