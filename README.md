
# Chirpy

A Twitter/X API clone that provides several API features such as Creating Chirps (like Tweets but Chirps), Deleting Chirps, Getting Chirps, Logging into Chirpy, Creating API token, Refreshing Token, Revoking Token, Upgrading to Premium Account (Chirp Red), Performing a Server Healthcheck, and Getting Server Metric.

# Installation

Run the below command inside a Go module:
``` go get github.com/tobib-dev/chirpy ```

# Requirements To Run

This app uses Go, Postgres, SQLC, and Goose. Install Go with the below command:
``` curl -sS https://webi.sh/golang | sh; \ source ~/.config/envman/PATH.env ```

Verify installation with the below command:

Go Version: ``` go version ```

Install Postgres with the below command:

Install Postgtes: ``` curl -sS https://webi.sh/postgres | sh; \ source ~/.config/envman/PATH.env ```

Set database connection with connection string, see below sample to setup connection string:

Connection String: ``` postgres://username:@localhost:5432/chirpy ```

Install SQLC and Goose in your Go module with the below commands:

Install SQLC: ``` go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest ```

Install Goose: ``` go install github.com/pressly/goose/v3/cmd/goose@latest ```

**Note: Environment variables are needed to do API features besides Readiness test** Therefore, make sure to set platform, JWT server token, and polka API key inorder to use features that require this variables

# Configuration

Set database url using the sample from the requirements section, set platform to *"dev"* as delete is only allowed in "dev", generate a JWT server token and assign it to the SERVER_TOKEN variable. POLKA_KEY is neccesary to upgrade membership to Chirpy Red, so make sure to set POLKA_KEY.

# Usage

Here are some sample usage of Chirpy

## Create User

You can create a user by making a POST request to the user endpoint like below:
**POST /api/users** with Request Body:
{
  "email": "example@email.com",
  "password": "0000"
}

Responds with a 201 status code and email, password isn't in response.

## Update User

Updating user requires passing user access token in authorization header
*Sample authorization header*
- Authorization: Bearer ${sampleAccessToken}

**PUT /api/users** with Request Body:
{
  "email": "example_new@email.com",
  password: "newPassw1d"
}

Responds with a 200 status code, similar to create password isn't included in response but email is.

## User Login

Login returns access token in the response body
**POST /api/login** with Request Body:
{
  "email": "example@email.com",
  "password": "0000"
}

Responds with a 200 status code and access token variable if access token is valid else responds with 401.

## Create Chirp

Pass access token into authorization header similar to update user request.
**POST /api/chirps** with Request body:
{
  "body": "Just setting up my chpy"
}

Responds with a 200 status code and body equals to the given chirp body if access token is valid else responds with a 401

## Get Chirps

There are several variation of get chirps:
1. Get all Chirps
2. Get chirps from a specific user
3. Get chirps by chirpID

To get all chirps simply make a GET request to the /api/chirps endpoint and you can get all chirps.
**GET /api/chirps**

To get chirps from a specific user you will have to provide an optional parameter of the author's id. **Note: If you don't provide the author's id then all chirps from the database will be returned in response body**
**GET /api/chirps?author_id=author_id**

To get chirps by the chirp id, simply include the chirp id in the request.
**GET /api/chirps/{chirpID}**

All three variation return a list of chirps and a 200 status code
**Note: You can also sort chirps based on the chirp created time, simply pass asc or desc into your request similar to how you will pass author id**
**GET /api/chirps?sort=asc**

## Delete Chirp

To delete chirp use Request
**DELETE /api/chirp/{chirpID}**

Delete returns a 204 status code if successful else 403 if user's access token doesn't match database. 404 if chirp is not found, 401 if token is not valid and 400 if the access token is malformed

## Upgrade Membership

To upgrade membership use Request
**POST /api/webhooks**

Upgrade returns a 204 if successful, 404 if user is not found. Upgrade also returns a 204 if *"user.upgrade"* is not provided as event value. Upgrade also returns a 401 if POLKA_KEY is not valid or not provided in the header

# Contact

Message me on discord with the link below:
[Discord](https://discord.gg/EFHrXVEd)
