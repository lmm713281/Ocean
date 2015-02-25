package Meta

type Category byte

const (
	CategoryBUSINESS = Category(iota)
	CategorySYSTEM   = Category(iota)
	CategoryAPP      = Category(iota)
	CategoryUSER     = Category(iota)
)

func (cat Category) Format() (result string) {
	switch cat {
	case CategoryBUSINESS:
		result = `C:BUSINESS`
	case CategoryAPP:
		result = `C:APP`
	case CategorySYSTEM:
		result = `C:SYSTEM`
	case CategoryUSER:
		result = `C:USER`
	default:
		result = `C:N/A`
	}

	return
}

func ParseCategory(cat string) (value Category) {
	switch cat {
	case `C:BUSINESS`:
		value = CategoryBUSINESS
	case `C:APP`:
		value = CategoryAPP
	case `C:SYSTEM`:
		value = CategorySYSTEM
	case `C:USER`:
		value = CategoryUSER
	default:
		value = CategoryAPP
	}

	return
}
