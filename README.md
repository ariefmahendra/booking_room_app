# Booking Room

## Prerequisites

Before running the Booking Room application, make sure you have the following prerequisites in place:

- **Go Installation:** Ensure that Go (Golang) is installed on your system.
- **PostgreSQL Setup:** Install PostgreSQL, create the required tables from the `ddl.sql` file, and populate them with dummy data using the `dml.sql` file.
- **Internet Connection:** An active internet connection is needed to download Go dependencies.

# Running the Application

Once the prerequisites are set up, run the Booking Room application. Access it via a web browser or use an API client like Postman or cURL. Log in using an admin-created account. The application provides APIs for managing Rooms, Facilities, Employees, and Transactions.

To register a new admin account, follow these steps:

1. **Use the Registration API:** Make a POST request to the registration endpoint to create a new account with role admin. Provide the necessary details, including email and a strong password.

2. **Log in as Admin:** Once the admin account is set up, log in using the credentials to access the admin functionalities and get token.

The application provides APIs for managing Rooms, Facilities, Employees, and Transactions.

## Using the API

Follow these instructions to utilize the API based on the features provided by the Booking Room application:

# API Documentation

### Guest API

- [Guest API Documentation](/api/guest_api.md)

### Employee Management API

- [Employee Management API Documentation](/api/employee_management_api.md)

### Room Management API

- [Room Management API Documentation](/api/room_management_api.md)

### Transaction Management API

- [Transaction Management API Documentation](/api/reservation_management.md)

### Room Management API 

- [Room Management API Documentation](/api/room_management_api.md)

### Report API

- [Report API Documentation](/api/room_management_api.md)