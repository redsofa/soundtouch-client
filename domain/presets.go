package domain

type Preset struct {
	Id          string      `xml:"id,attr"`
	ContentItem ContentItem `xml:"ContentItem"`
}

type ContentItem struct {
	Source   string `xml:"source,attr"`
	Location string `xml:"location,attr"`
}

type Presets struct {
	Pressets []Preset `xml:"preset"`
}
