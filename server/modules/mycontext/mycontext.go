package mycontext

import (
	"strconv"

	"../myblog"
)

//Context is the context of application
type Context struct {
	Blogs []myblog.Blog
}

//CreateDefContext context
func CreateDefContext() *Context {

	Context := Context{}
	Context.Blogs = []myblog.Blog{}

	blog1 := myblog.Blog{
		Autor:    "Autor1",
		Articles: []myblog.Article{}}

	var newArticle myblog.Article

	for i := 1; i < 100; i++ {

		newArticle = myblog.Article{
			Title:   "Title" + strconv.Itoa(i),
			Content: "Context" + strconv.Itoa(i)}

		blog1.Articles = append(blog1.Articles, newArticle)
	}

	Context.Blogs = append(Context.Blogs, blog1)

	return &Context
}
