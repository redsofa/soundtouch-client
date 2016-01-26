package domain

type Preset struct {
	Id          string      `xml:"id,attr"`
	ContentItem ContentItem `xml:"ContentItem"`
}

type ContentItem struct {
	Source   string `xml:"source,attr"`
	Location string `xml:"location,attr"`
	ItemName string `xml:"itemName"`
}

type Presets struct {
	Presets []Preset `xml:"preset"`
}
