## Golang Api with postgresSQL
    # Packages Necessary
        https://github.com/gorilla/mux
        https://github.com/jinzhu/gorm
        https://github.com/joho/godotenv
        https://github.com/lib/pg

        *You have to fill in .env with the information from the database.

     
## File Structure

```
Api
.
├── README.md
├── .env
├── main.go
└── api
    ├── controller
    │   ├── baseController.go
    │   └── userController.go
    ├── models
    │   └── user.go
    └── responses
        └── json.go
```