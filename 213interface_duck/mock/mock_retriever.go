package mock

type Retriever struct {
	Content string
}

func (r Retriever) GetURLContent(url string) string {
	return r.Content
}
