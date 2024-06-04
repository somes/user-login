## user-login

This is the backend part of a demonstration user login system that supports multiple authentication methods, including logging in with username/password alone or with a digital certificate. It also allows for two-factor authentication login.

### Configuration File

The project's configuration information is stored in the `config/application.yaml` file, which includes the project's running port and database connection information.

### How to Run

```bash
# Run the program
go run cmd/main.go
```

### Features

Below are the implemented and planned features of the project along with their completion status:

| Feature                            | Status  |
|------------------------------------|---------|
| User registration                  | &#9744; |
| Username / Password login          | &#9745; |
| Digital certificate authentication | &#9744; |
| Two-factor authentication          | &#9744; |
| User management                    | &#9744; |

**Note:** The project is still in development, and feature sets may be subject to change.
