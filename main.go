package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/btussupb/blog/models"
)

var posts map[string]*models.Post

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/index.html", "templates/footer.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	if err = tmp.ExecuteTemplate(w, "index", nil); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/footer.html", "templates/header.html", "templates/write.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if err = tmp.ExecuteTemplate(w, "write", nil); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/footer.html", "templates/header.html", "templates/write.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	id := r.FormValue("id")
	post, found := posts[id]
	if !found {
		http.NotFound(w, r)
	}

	if err = tmp.ExecuteTemplate(w, "write", post); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post *models.Post

	if id != "" {
		post = posts[id]
		post.Title = title
		post.Content = content
	} else {
		id = GenerateId()
		post := models.NewPost(id, title, content)
		posts[post.Id] = post
	}

	posts[post.Id] = post

	http.Redirect(w, r, "/", 302)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id != "" {
		http.NotFound(w, r)
	}

	delete(posts, id)
	http.Redirect(w, r, "/", 302)
}

func main() {
	fmt.Println("Listening on port localhost:3000")

	posts = make(map[string]*models.Post, 0)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/savePost", savePostHandler)

	http.ListenAndServe(":3000", nil)
}
