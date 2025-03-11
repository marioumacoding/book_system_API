# 📚 Book Management API

This is a simple **Book Management API** built using **Go** and the **Gin** framework. It allows users to perform CRUD (Create, Read, Update, Delete) operations on books.

---

## 🚀 Features
- Get all books
- Get a book by ID
- Add a new book
- Update an existing book
- Partially update a book
- Delete a book

---

## 🛠️ Tech Stack
- **Go** (Golang)
- **Gin** (Web Framework)
- **MySQL** (Database - Future Implementation)
- **Redis** (Caching - Future Implementation)

---

## 📦 Installation & Setup
### 1️⃣ Clone the Repository
```sh
git clone https://github.com/your-username/book-management-api.git
cd book-management-api
```

### 2️⃣ Install Dependencies
```sh
go mod tidy
```

### 3️⃣ Run the Application
```sh
go run main.go
```
The API will run on `http://localhost:8080`

---

## 📌 API Endpoints
| Method  | Endpoint         | Description            |
|---------|----------------|------------------------|
| GET     | `/books`        | Get all books         |
| GET     | `/books/:id`    | Get book by ID        |
| POST    | `/books`        | Add a new book        |
| PUT     | `/books/:id`    | Update a book         |
| PATCH   | `/books/:id`    | Partially update book |
| DELETE  | `/books/:id`    | Delete a book         |

---

## 🔥 Future Enhancements
- Connect to a **MySQL database** for persistent storage
- Implement **Redis caching**
- Add **user authentication** (JWT)

---

## 📝 License
This project is open-source and available under the **MIT License**.

💡 **Feel free to contribute or give feedback!** 🚀

