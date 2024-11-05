package main

var (
	articles []string
	courselist = make(map[string]int)
)

func main() {
		displayCourse()
		addArticle("pomme")
		displayCourse()

}

func addArticle(article string) {
	articles = append(articles, article)
courseList[]}