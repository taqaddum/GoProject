package opstatus

type Generic interface {
	Common
	String() string
}
