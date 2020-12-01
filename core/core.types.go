package core

type Generic struct {
	ID   int64  `json:"-"`
	UUID string `json:"id"`
}

type Asset struct {
	Description string `json:"description"`
	Link        string `json:"link"`
}

type Resource struct {
	Videos []Asset `json:"videos,omitempty"`
	Images []Asset `json:"images,omitempty"`
}

type General struct {
	Generic
	Name        string   `json:"name,omitmepty"`
	Description string   `json:"description,omitempty"`
	Resources   Resource `json:"resources"`
}
