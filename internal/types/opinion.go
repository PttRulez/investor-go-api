package types

type OpinionType string

const (
	Flat      OpinionType = "Flat"
	General   OpinionType = "General"
	Growth    OpinionType = "Growth"
	Reduction OpinionType = "Reduction"
)

func (e OpinionType) Validate() bool {
	switch e {
	case Flat:
	case General:
	case Growth:
	case Reduction:
	default:
		return false
	}
	return true
}
