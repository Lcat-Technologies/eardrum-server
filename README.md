# Eardrum Server    

## Overview  
Eardrum Server is a campus-focused fintech engine designed to power secure, cashless, deviceless, and contactless transactions. By integrating dynamic QR code identification with facial recognition technology, the system enables friction-free payment flows and identity verification across school ecosystems—allowing students and users to execute point-of-sale (POS) transactions without requiring physical cards or personal devices.

The platform provides core functionality including:

* **User & Merchant Management**: Account creation, lifecycle tracking, and verification workflows for system users and merchants.
* **Internal Wallet & Ledger Engine**: Tracks internal balances (denominated in cents) with support for overdraft facilities, direct account debits, and merchant settlements.
* **Biometric Vector Processing**: Stores array sets of base64-encoded facial embeddings and photo URLs for 1:1 facial verification during authorization.
* **Dynamic QR Code Identification**: Links unique dynamic UUID QR codes directly to user accounts to resolve targets during POS checkouts.
* **Transaction Logging & Scan Auditing**: Logs real-time and offline transactions (combining device IMEI and timestamps) alongside base64 audit scan logs.

## Table of Contents

* [Technologies](#technologies)
* [Frameworks and Libraries](#frameworks-and-libraries)
* [Third-party Services](#third-party-services)
* [Architecture](#architecture)
* [Deployment Guide](#deployment-guide)
* [Database Design](#database-design)
* [Testing and Quality Assurance](#testing-and-quality-assurance)
* [Security Considerations](#security-considerations)
* [Logging and Error Handling](#logging-and-error-handling)

## Technologies 

- Golang
- GraphQL
- PostgreSQL

## Frameworks and Libraries

- gqlgen
- autogql
- godotenv
- twilio-go
- gorm
- eardrum-prefix
- zerolog

## Third-party Services 

- Twilio Verify

## Architecture
![birds-eye](https://github.com/user-attachments/assets/9d52fa2c-7b6b-42bc-be92-542f12f6f6d3)
        

## Deployment Guide
* Install [gcloud CLI](https://cloud.google.com/sdk/docs/install)
* Run `gcloud init` to link your local CLI with your Google Cloud project.
* Ensure Docker Engine is installed and running on your computer. If needed, install it [here](https://docs.docker.com/engine/install/).
* Clone the `eardrum-server` repository to your local machine.
* Build the Docker image for your preferred `eardrum-server` release:
  ```bash
  docker build -t [DOCKERHUB_USERNAME]/eardrum-server:[TAG] .
  ```

### Docker Deployment
Authenticate with Docker Hub
```bash
docker login
```
Push the Tagged Image to Docker Hub
```bash
docker push [DOCKERHUB_USERNAME]/eardrum-server:[TAG]
```
### Deploy to Google Cloud Run

1. Open the Google Cloud Run Console.
2. Click Deploy Container → Service.
3. Under Container image URL, enter your Docker Hub image path:
```
docker.io/[DOCKERHUB_USERNAME]/eardrum-server:[TAG]
```
4. Configure the service with the following required environment variables:

```
POSTGRES_DBURL
JWT_SECRET_KEY
TWILIO_ACCOUNT_SID
TWILIO_AUTH_TOKEN
TWILIO_VERIFY_SERVICE_SID
DEFAULT_PORT
```

## Database Design

The relational database layer is built on PostgreSQL and managed via GORM. It handles user balances, merchant accounts, facial vectors, dynamic QR references, and transaction ledgers.

### Relational Database Schema (SQL DDL)
#### 1. users Table

Purpose: Stores verified customer profiles, internal wallet and overdraft balances (in cents), dynamic QR code UUIDs, and facial embedding vectors for POS verification.
```
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    user_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    password TEXT NOT NULL,
    account_balance_in_cents BIGINT NOT NULL DEFAULT 0,
    overdraft_balance_in_cents BIGINT NOT NULL DEFAULT 0,
    pin_code TEXT,
    qr_code UUID,
    facial_embeddings TEXT[],
    facial_images TEXT[],
    passport TEXT
);

CREATE UNIQUE INDEX idx_users_user_name ON users(user_name);
CREATE UNIQUE INDEX idx_users_phone_number ON users(phone_number);
CREATE UNIQUE INDEX idx_users_qr_code ON users(qr_code);
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
```
#### 2. unverified_users Table

Purpose: Holds pending user registrations before phone/identity verification is complete.
```
CREATE TABLE unverified_users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    user_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    password TEXT NOT NULL,
    account_balance_in_cents BIGINT NOT NULL DEFAULT 0,
    overdraft_balance_in_cents BIGINT NOT NULL DEFAULT 0,
    pin_code TEXT,
    qr_code UUID,
    facial_embeddings TEXT[],
    facial_images TEXT[],
    passport TEXT
);

CREATE UNIQUE INDEX idx_unverified_users_user_name
    ON unverified_users(user_name);

CREATE UNIQUE INDEX idx_unverified_users_phone_number
    ON unverified_users(phone_number);

CREATE UNIQUE INDEX idx_unverified_users_qr_code
    ON unverified_users(qr_code);

CREATE INDEX idx_unverified_users_deleted_at
    ON unverified_users(deleted_at);
```
#### 3. merchants Table

Purpose: Manages official campus vendor accounts and settlement balances in cents.
```
CREATE TABLE merchants (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    user_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    password TEXT NOT NULL,
    account_balance_in_cents BIGINT NOT NULL DEFAULT 0,
    pin_code TEXT
);

CREATE UNIQUE INDEX idx_merchants_user_name
    ON merchants(user_name);

CREATE UNIQUE INDEX idx_merchants_phone_number
    ON merchants(phone_number);

CREATE INDEX idx_merchants_deleted_at
    ON merchants(deleted_at);
```
#### 4. unverified_merchants Table

Purpose: Holds merchant onboarding requests prior to verification.
```
CREATE TABLE unverified_merchants (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    user_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    password TEXT NOT NULL,
    account_balance_in_cents BIGINT NOT NULL DEFAULT 0,
    pin_code TEXT
);

CREATE UNIQUE INDEX idx_unverified_merchants_user_name
    ON unverified_merchants(user_name);

CREATE UNIQUE INDEX idx_unverified_merchants_phone_number
    ON unverified_merchants(phone_number);

CREATE INDEX idx_unverified_merchants_deleted_at
    ON unverified_merchants(deleted_at);
```
#### 5. transactions Table

Purpose: Records online and offline transactions between users and merchants, including fee deductions and authorization scan audit images.
```
CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    transaction_id VARCHAR(12) NOT NULL,
    offline_transaction_id VARCHAR(36),
    total_amount_in_cents BIGINT NOT NULL,
    transaction_cost_in_cents BIGINT NOT NULL,
    scan_log TEXT,
    user_user_name TEXT NOT NULL,
    merchant_user_name TEXT NOT NULL,

    CONSTRAINT fk_transactions_user
        FOREIGN KEY (user_user_name)
        REFERENCES users(user_name),

    CONSTRAINT fk_transactions_merchant
        FOREIGN KEY (merchant_user_name)
        REFERENCES merchants(user_name)
);

CREATE UNIQUE INDEX idx_transactions_transaction_id
    ON transactions(transaction_id);

CREATE UNIQUE INDEX idx_transactions_offline_transaction_id
    ON transactions(offline_transaction_id);

CREATE INDEX idx_transactions_deleted_at
    ON transactions(deleted_at);
```
## Testing and Quality Assurance
### GraphQL Playground

GraphQL Playground is used to test queries and mutations for:

User and merchant onboarding

QR code resolution

Facial vector authorization

Internal balance debiting

Responses are validated against the target schema conditions and financial constraints.

## Security Considerations
### Credential Protection

Passwords and PIN codes are stored strictly as salted, secure hashes (e.g., bcrypt) and are never exposed in plaintext.

### JSON Web Tokens (JWT)

JWTs are issued upon authentication to enforce role-based access control, such as distinguishing between user and merchant capabilities.

### Two-Factor Contactless Verification (1:1)

Point-of-sale transactions use a two-step verification process:

A scanned dynamic QR UUID is used to resolve the user's profile.

A 1:1 facial vector match is performed to help prevent fraudulent checkouts.

### Dynamic QR Codes

Dynamic QR codes regenerate on demand. Once a new QR code is generated, previous codes are immediately invalidated to prevent replay attacks.

## Logging and Error Handling
### Logging

Structured JSON logging is implemented using zerolog and records operational events across standard log levels:

info — Tracks routine API execution, such as processing a payment request.

trace — Logs granular setup steps, such as database connection pool initialization.

error — Identifies non-fatal errors, such as JWT generation failures, that allow the system to remain online.

warn / fatal — Logs critical application flaws or structural failures before executing a safe termination.

Common log fields include:

level
id
role
path
time
message

### Error Handling
#### Minor Errors

Minor errors are returned in the GraphQL response body and logged under the error level.

#### Critical Errors

Critical errors are logged under the warn or fatal levels, allowing system administrators to diagnose transaction anomalies and integrity failures promptly.
