# Practice-Basic-GO-REST-API


nouvalkaff/Practice-Basic-GO-REST-API


How to run the file:
1. Type or copy-paste this command on terminal "go build", then "Enter".
2. Type or copy-paste this command on terminal "./Practice-Basic-GO-REST-API", then "Enter".
3. {Optional} Type or copy-paste this command on Git Bash terminal "go build && ./Practice-Basic-GO-REST-API", then "Enter".
4. Local server shall be running.


Route endpoints:
("/api/books", getAllBooks).Methods("GET")
("/api/book/{id}", getBook).Methods("GET")
("/api/book/crt", createBook).Methods("POST")
("/api/book/upd/{id}", updateBook).Methods("PUT")
("/api/book/del/{id}", deleteBook).Methods("DELETE")
