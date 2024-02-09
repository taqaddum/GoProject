package enum

type Gender uint8

const (
	Female Gender = 0
	Male   Gender = 1
)

func (gen Gender) ToString() string {
	switch gen {
	case Male:
		return "男"
	case Female:
		return "女"
	}
	return "unknown"
}
