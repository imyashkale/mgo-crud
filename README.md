# mgo-crud
REST API [ CRUD ] using GO &amp; MongoDB .
### Below functionality inmplemented in this repo.

```

	router := mux.NewRouter()

	// GET :  Get All the Books
	router.HandleFunc("/api/v1/book/", handlers.GetBooks).Methods(http.MethodGet)

	// GET : Get Book By ID
	router.HandleFunc("/api/v1/book/{id}", handlers.GetBook).Methods(http.MethodGet)

	// POST : Create New Book
	router.HandleFunc("/api/v1/book/", handlers.CreateBook).Methods(http.MethodPost)

	// PUT : Update the Book By ID
	router.HandleFunc("/api/v1/book/{id}", handlers.UpdateBook).Methods(http.MethodPut)

	// DELETE : Delete the book by ID
	router.HandleFunc("/api/v1/book/{id}", handlers.DeleteBook).Methods(http.MethodDelete)
 
 ```
