package main

func main() {
	server := NewServer(":8000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/home", server.AddMiddleware(HandleHome, CheckAuth(), Login()))
	server.Handle("POST", "/create", server.AddMiddleware(PostRequest, Login()))
	server.Handle("POST", "/user", server.AddMiddleware(UserPostRequest, Login()))
	server.Listen()
}
