package http

type RequestDelete struct {
	TargetID int `json:"target_id"`
}
type UpdateRequest struct {
	NewID   int    `json:"new id,omitempty"`
	NewName string `json:"new name,omitempty"`
}

type ListInput struct {
	ID     []int    `json:"id" in:"query=id[],id"`
	Name   []string `json:"name" in:"query=name[],name"`
	SortBy string   `json:"sort_by" in:"query=sort_by"`
}
