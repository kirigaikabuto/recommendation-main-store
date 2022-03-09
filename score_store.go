package recommendation_system_main_store

type ScoreStore interface {
	Create(score *Score) (*Score, error)
	List() ([]Score, error)
}
