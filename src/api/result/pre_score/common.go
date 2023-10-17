package pre_score

type CommonRequest struct {
	PreScoreID uint `uri:"pre_score_id"`

	ID        uint   `json:"ID"`
	Processor string `json:"Processor"`
}
