package step

//Step Step data type
type Step struct {
	ID        uint64 `json:"id"`
	ReaderNum int    `json:"reader_num" binding:"required"`
	TagNum    uint64 `json:"tag_num" binding:"required"`
	UserNum   uint64 `json:"user_num"`
	Style     int    `json:"style"`
	Start     string `json:"start"`
	End       string `json:"end"`
	IsRemove  bool   `json:"is_remove"`
}

//New new one defalut step
func New() Step {
	defalutStep := Step{
		ID:        defaultID,
		ReaderNum: defaultReaderNum,
		TagNum:    defaultTagNum,
		UserNum:   defaultUserNum,
		Style:     defaultStyle,
		Start:     defaultStart,
		End:       defaultEnd,
		IsRemove:  defaultIsRemove,
	}
	return defalutStep
}
