# Lucky Draw API Server

The Lucky Draw API Server is a Golang application built with the Gin framework.

## Introduction

This API server allows customers to participate in a lucky draw campaign. Each customer can draw multiple times during the campaign period, but they are only allowed to draw once per day.

## Running the Server

To run the server, execute the following command in your terminal:

``` sh
# Using Docker
docker-compose up
```

After running the server, you can access the Swagger documentation at:
<http://localhost:8080/api/swagger/index.html>

## Endpoints

### Draw

To enter the draw, make a request to the following endpoint:

``` sh
# Enter draw with customerId param
http GET http://localhost:8080/api/draw/:customerId

```

- The customerId parameter should contain the ID of the customer making the draw request.

The behavior of this endpoint is as follows:

- It first checks if the customerId is valid. If it is not, an error is returned.
- If the customerId is valid, it checks if the customer has already made a draw on the current day.
  - If the customer has already made a draw, they are reminded to try again tomorrow.
  - If the customer has not made a draw, they are allowed to proceed.
- If the customer does not win any prize, they are reminded to try again tomorrow.
- Otherwise, it checks if the prize has reached its daily or total quota.
  - If the quota has been reached, the customer is reminded to try again tomorrow.
  - If the quota has not been reached, the daily and total quota are reduced by one, and the customer is notified of their prize.

### Redeem

To redeem a prize, make a request to the following endpoint:

``` sh
# Request user to enter mobile phone number
http POST http://localhost:8080/api/redeem/:customerId mobile={mobileNumber}
# e.g. http POST http://localhost:8080/api/redeem/2 mobile=98765432

```

- The customerId parameter should contain the ID of the customer making the draw request.

In this endpoint:

- The user is required to provide their mobile phone number for reference and SMS.
- The endpoint checks if the provided mobile phone number is valid (8 digits).
- If the mobile phone number is valid, it is saved in the database in the mobile table.
