package models

func (_ *_Helper) ModifyLimit(limit int) int {
	if limit < MinLimit || limit > MaxLimit {
		limit = MaxLimit
	}
	return limit
}

type _Helper struct{}

var (
	Helper *_Helper
)
