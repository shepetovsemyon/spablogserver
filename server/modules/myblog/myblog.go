package myblog

//Blog is slice of articles
type Blog struct {
	Autor    string
	Articles []Article
}

//Article of Blog
type Article struct {
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

//About blog
type About struct {
	Autor string `json:"Autor"`
	Count int    `json:"Count"`
}

//Count return count of articles in blog
func (b Blog) Count() int {
	return len(b.Articles)
}

//About blog func
func (b Blog) About() About {
	return About{
		b.Autor,
		b.Count()}
}
