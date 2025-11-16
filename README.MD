# 📝 Run System BE Technical Test

## Brief Explanation
This Backend Project is built using GO, gin framework, and sqlc.
Using Gin provides fast and most populer for building REST APIs in GO.
The project is built clean enough but not complex for simple develop.


---
## Main Test

### 1. Create Entry Endpoint
**POST** `/person/add`

Request
```json
{
  "name": "M. Radhito Bil Atho",
  "ambition": "Make usefull innovation to help alongside non-tech activites user"
}

```

### 2. Get Entry Endpoint
**GET** `/person/{id}`

Response
```json
{
    "id": 1,
    "name": "M. Radhito Bil Atho",
    "ambition": "Make usefull innovation to help alongside non-tech activites user"
}

```

### Reason behind My Ambition
#### I have a strong will to be impactful, creates something meaningful with technology is the way i can get involved in civilization. Many non-tech users or everyday individuals still face problems they don’t even realize can be solved through tech. There's still a gap I from many non-tech users facing daily problem that can be solved through tech.

## 📁 Project Structure
```
be/
│
├── api/
│ ├── task.go
│ ├── middleware.go
│ ├── main_test.go
│ ├── person.go
│ ├── users.go
│ └── server.go
│
├── db/
│ ├── migration/
│ ├── query/
│ │ └── *.sql
│ ├── sqlc/
│
├── util/
│ ├── config.go
│ ├── random.go
│ ├── token.go
│ └── password.go
│
└── main.go
```

## 🗄️ Database Schema

### **Person Table**

| Column   | Type    | Description        |
|----------|---------|--------------------|
| id       | BIGINT  | Primary key        |
| name     | VARCHAR | Unique username    |
| ambition | VARCHAR |                    |

## Optional Test Features

### 🔐 Authentication
- Register a new user  
- Login user  
- JWT-based session  
- Auth middleware to protect routes  

### 📌 Task System
- Add a task  
- Get all user tasks  
- Finish a task (gain XP + task removed)

### 🧠 User Progression
- Tasks give XP  
- XP contributes to user level  
- Auto-level-up based on progression formula  

---

## 🗄️ Database Schema

### **Users Table**

| Column   | Type    | Description        |
|----------|---------|--------------------|
| id       | BIGINT  | Primary key        |
| username | VARCHAR | Unique username    |
| password | TEXT    | Hashed password    |
| exp      | INT     | Total XP           |
| level    | INT     | User level         |

### **Tasks Table**

| Column   | Type    | Description              |
|----------|---------|--------------------------|
| id       | BIGINT  | Primary key              |
| todo     | TEXT    | Task description         |
| exp      | INT     | XP from completing task  |
| user_id  | BIGINT  | Foreign key (users.id)   |

---

## 🔧 Tech Stack

- **Golang**
- **Gin Web Framework**
- **sqlc**
- **MySQL / PostgreSQL**
- **JWT Authentication**
- **BCrypt Password Hashing**

---

## 🔧 app.env
```bash
DB_DRIVER = postgres
SERVER_ADDRESS = 0.0.0.0:8080
DB_SOURCE = postgresql://postgres:secret@localhost:5433/runsystem_be_test?sslmode=disable
```

## ▶️ Run Locally

### 1️⃣ Clone Repository
```bash
git clone https://github.com/Roditu/BE_RS_TEST.git
cd be