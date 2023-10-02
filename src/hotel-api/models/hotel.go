package models

type Hotel struct {
    ID          string `json:"id,omitempty"`
    Name        string `json:"name"`
    Description string `json:"description"`
    // ... otros campos
}
