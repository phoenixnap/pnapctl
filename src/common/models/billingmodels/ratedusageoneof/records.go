package ratedusageoneof

type BandwidthRecord struct {
	RatedUsage
	Metadata BandwidthDetails `json:"metadata" yaml:"metadata"`
}

type OperatingSystemRecord struct {
	RatedUsage
	Metadata OperatingSystemDetails `json:"metadata" yaml:"metadata"`
}

type PublicSubnetRecord struct {
	RatedUsage
	Metadata PublicSubnetDetails `json:"metadata" yaml:"metadata"`
}

type ServerRecord struct {
	RatedUsage
	Metadata ServerDetails `json:"metadata" yaml:"metadata"`
}
