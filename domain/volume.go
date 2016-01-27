package domain

type Volume struct {
	TargetVolume string `xml:"targetvolume"`
	ActualVolume string `xml:"actualvolume"`
	MuteEnabled  string `xml:"muteenabled"`
}
