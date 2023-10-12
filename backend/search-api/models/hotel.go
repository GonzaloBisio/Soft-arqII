package models

type Hotel struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

type Hotels []Hotel 