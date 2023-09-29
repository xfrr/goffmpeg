package media

type Metadata struct {
	Streams []Streams `json:"streams"`
	Format  Format    `json:"format"`
}
