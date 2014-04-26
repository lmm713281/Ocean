package Meta

type Category byte

const (
	CategoryBUSINESS = Category(iota)
	CategorySYSTEM   = Category(iota)
	CategoryAPP      = Category(iota)
	CategoryUSER     = Category(iota)
)

func FormatCategory(cat Category) (result string) {
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
