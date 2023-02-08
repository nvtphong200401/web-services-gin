package models

//    landmark, err := UnmarshalLandmark(bytes)
//    bytes, err = landmark.Marshal()

import "encoding/json"

func UnmarshalLandmark(data []byte) (Landmark, error) {
	var r Landmark
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Landmark) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Landmark struct {
	Name        string      `json:"name"`
	Category    string      `json:"category"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	ID          int64       `json:"id"`
	IsFeatured  bool        `json:"isFeatured"`
	IsFavorite  bool        `json:"isFavorite"`
	Park        string      `json:"park"`
	Coordinates Coordinates `json:"coordinates"`
	Description string      `json:"description"`
	ImageName   string      `json:"imageName"`
}

type Coordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
