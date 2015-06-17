package Meta

type Category byte

const (
	CategoryBUSINESS = Category(iota) // Business category
	CategorySYSTEM   = Category(iota) // System category
	CategoryAPP      = Category(iota) // Application category
	CategoryUSER     = Category(iota) // User category
)

// Formats a category as string.
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

// Parse a category from a string.
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
