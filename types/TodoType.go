package types

type TodoType struct {
	ID   string `json:"id"`
	Text string `json:"description"`
	Done bool   `json:"done"`
}
