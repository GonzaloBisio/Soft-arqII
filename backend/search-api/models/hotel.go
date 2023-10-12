package model

type Hotel struct {
	ID          string   `json:"id"` 
	Name        string   `json:"name"`
	City        string   `json:"city"`
	Description string   `json:"description"`
	Thumbnail   string   `json:"thumbnail"`
	Images      []string `json:"images"`
	Amenities   []string `json:"amenities"`
}

type Hotels []Hotel