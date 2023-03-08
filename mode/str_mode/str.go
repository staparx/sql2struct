package str_mode

type strMode struct {
}

func NewStrMode() *strMode {
	return &strMode{}
}

func (m *strMode) DbTable(param string) []string {
	return []string{param}
}
