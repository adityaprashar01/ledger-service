# 💼 Ledger Service API

A backend service to manage customer accounts and transactions, built with **Golang**, **MongoDB Atlas**, and **Gin** framework.

🔗 **Live Demo**: [https://ledger-service.onrender.com](https://ledger-service.onrender.com)  
📘 **API Docs (Swagger UI)**: [https://ledger-service.onrender.com/swagger/index.html](https://ledger-service.onrender.com/swagger/index.html)

---

## ✨ Features

- Create customer accounts
- Add transactions (credit/debit)
- Get customer balance
- View transaction history (With pagination)
- RESTful API with OpenAPI (Swagger) documentation
- Hosted on Render with MongoDB Atlas

---

## 🚀 Tech Stack

- **Go** (Golang)
- **Gin** Web Framework
- **MongoDB Atlas**
- **Render** (Deployment)
- **Swagger/OpenAPI** (API documentation)

---

## ⚙️ Local Setup Instructions

### 1️⃣ Install Go

Download from: https://go.dev/dl/  
Ensure Go is installed:
```bash
go version
```

### 2️⃣ Clone the Repository

```bash
git clone https://github.com/your-username/ledger-service.git
cd ledger-service
```

### 3️⃣ Setup Environment Variables

Create a `.env` file in the root directory:

```env
MONGO_URI=mongodb+srv://<your_username>:<your_password>@<cluster>.mongodb.net/?retryWrites=true&w=majority
```

(Use your **MongoDB Atlas connection string** here.)

### 4️⃣ Install Dependencies

```bash
go mod tidy
```

### 5️⃣ Run the Server

```bash
go run main.go
```

It will start on `http://localhost:8080`

---

## 🧪 Testing the API Locally

Use tools like **Postman** or **cURL**.


## 📄 API Documentation

All endpoints accept and return **JSON**.

### 1. Create Customer

```http
POST /customers
```

**Request Body:**

```json
{
  "name": "John Doe",
  "initial_balance": 5000
}
```

**Response:**

```json
{
  "_id": "64f1234...",
  "name": "John Doe",
  "balance": 5000
}
```

---

### 2. Get Customer Balance

```http
GET /customers/:customer_id/balance
```

**Response:**

```json
{
  "customer_id": "64f1234...",
  "balance": 5000
}
```

---

### 3. Create Transaction

```http
POST /transactions
```

**Request Body:**

```json
{
  "customer_id": "64f1234...",
  "amount": 1000,
  "type": "credit" // or "debit"
}
```

**Response:**

```json
{
  "transaction_id": "...",
  "customer_id": "...",
  "amount": 1000,
  "type": "credit",
  "timestamp": "..."
}
```

---

### 4. Get Transaction History 

Retrieve the transaction history for a specific customer.

```http
GET /customers/:customer_id/transactions
```

**Path Parameter:**

- `customer_id` (string) – The ID of the customer

**Optional Query Parameters:**

- `page` (integer) – Page number (default: 1)
- `limit` (integer) – Number of transactions per page (default: 10)

---

### 🔄 Example Request (Paginated)

```http
GET /customers/6613a1234abcde23456789ff/transactions?page=1&limit=2
```

### Response:

```json
{
  "page": 1,
  "limit": 2,
  "total": 4,
  "transactions": [
    {
      "transaction_id": "6613aabcde123456789fff01",
      "amount": 1000,
      "type": "credit",
      "timestamp": "2025-04-01T12:34:56Z"
    },
    {
      "transaction_id": "6613aabcde123456789fff02",
      "amount": 500,
      "type": "debit",
      "timestamp": "2025-04-02T09:30:21Z"
    }
  ]
}
```

### Status Codes:

- `200 OK` – Success
- `404 Not Found` – Customer not found
- `500 Internal Server Error` – Server-side error


## 📘 Swagger API Documentation

Once the server is running, visit:

```
http://localhost:8080/swagger/index.html
```

OR for deployed:

```
https://ledger-service.onrender.com/swagger/index.html
```

---

## ☁️ Deploying on Render (Optional)

1. Push your code to GitHub.
2. Go to [https://dashboard.render.com](https://dashboard.render.com).
3. Create a new Web Service → Connect your repo.
4. Add environment variable:

```
Key: MONGO_URI
Value: <your MongoDB Atlas URI>
```

5. Render will build and deploy your service.

---

## 📂 Project Structure

```
.
├── config/          # Configuration files
├── controllers/     # API handler functions
├── database/        # MongoDB connection
├── docs/            # Swagger auto-generated files
├── models/          # Data models
├── routes/          # Router setup
├── tests/           # Unit tests
├── main.go          # Entry point
├── go.mod           # Go module file
├── .env             # Environment variables
```

---

## 📦 Generate Swagger Docs (If needed)

Install swag CLI:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate docs:

```bash
swag init --generalInfo main.go
```

---

## 🛠️ Future Improvements

- Multi-currency support
- Rate limiting
- Role-based access control (RBAC)
- Docker support

---

## 🧑‍💻 Author

**Aditya Prashar**  
[LinkedIn](https://www.linkedin.com/in/aditya-prashar03) 

---
