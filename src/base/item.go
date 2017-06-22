package base

type Item map[string]interface{}

// 检查条目是否合法
func (item Item) Valid() bool {
	return item != nil
}
