# Simple Stateless Go REST API

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

This is a Go project that includes a stateless REST API structure used the Gin web framework.

## Prequisites

- Go 1.21 or later


### Mahasiswa API Routes

- `GET /mahasiswa/` - Get all the data from mahasiswa.
- `POST /mahasiswa` - Create a new data for mahasiswa.
- `GET /mahasiswa/:nim` - Get each mahasiswa's data by NIM.
- `PUT /mahasiswa/:nim` - Edit each mahasiswa'data by NIM.
- `DELETE /mahasiswa/:nim` - Delete a mahasiswa by NIM.

## Getting Started
1. Clone the repository:

```sh
git clone https://github.com/ahmadammarm/simple-go-rest.git
```

2. Navigate to the project directory:

```sh
cd simple-go-rest
```

3. Install the project dependencies:

```sh
go mod download
```

4. Run the project:

```sh
go run main.go
```



The project will be available at `http://localhost:8080`.
