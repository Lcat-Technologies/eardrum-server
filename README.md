# Eardrum Server    

## Overview  
This project is a school management system that allows multiple schools to host their data, allow restricted access to the data from other parties such as teachers, students and other stakeholders. The sytem also provides various functionalities such as:

* **Creation of a school account**: This project allows schools to register in order to use the system to host their data.
* **Adding of student data**: Allows for schools to load their students' data into their network in the system.
* **School and Student login**: Allows for schools to log into their accounts and for students to log into their accounts in the respective school network they were created in.
* **Retrieving Student and School data**: Allows for retrieval of schools and students data that is already loaded into the system.

## Table of Contents

* [Technologies](#technologies)
* [Frameworks and Libraries](#frameworks-and-libraries)
* [Third-party Services](#third-party-services)
* [Architecture](#architecture)
* [Deployment Guide](#deployment-guide)
* [Database Design](#database-design)
* [Data Synchronization](#data-synchronization)
* [Testing and Quality Assurance](#testing-and-quality-assurance)
* [Security Considerations](#security-considerations)
* [Logging and Error Handling](#logging-and-error-handling)

## Technologies 

- Golang
- GraphQL
- Postgres
- Neo4j

## Frameworks and Libraries

- gqlgen
- autogql
- godotenv
- twilio-go
- gorm
- eardrum-graph
- eardrum-prefix
- zerolog

## Third-party Services 

- Twilio verify

## Architecture
![birds-eye](https://github.com/user-attachments/assets/b1beea2d-5cd8-42e3-b11a-3ac1b3d10c3b)        
## Deployment Guide
* Install [gcloud CLI](https://cloud.google.com/sdk/docs/install)
* `gcloud init` to link up your gcloud CLI with your google cloud project
* Ensure you have installed Docker engine on your local computer. If you haven't install [here](https://docs.docker.com/engine/install/)
* Clone the eardrum-server repo into your local machine
* Build a docker-image off your preffered eardrum-server version or release
* Setup an Artifact Registry Docker repository in Google Cloud. Follow this [guide](https://www.youtube.com/watch?v=b-rg71O3ZwU&t=9s)
* `gcloud auth configure-docker` to configure authentication to Artifact Registry for Docker. For more information [here](https://cloud.google.com/artifact-registry/docs/docker/authentication)
* Tag the local image that you build using this commands
```
docker tag [IMAGE_NAME] \
gcr.io/[PROJECT_ID]/[REPO_NAME]/[IMAGE_NAME]:[TAG]
```
* Next, push the tagged image to Artifact Registry by running:
```
docker push gcr.io/[PROJECT_ID]/[REPO_NAME]/[IMAGE_NAME]:[TAG]
```
If it fails for some reason run `gcloud auth configure-docker` and then push it again
* Now go to Google Cloud Run console > Navigate to "DEPLOY CONTAINER" > pick "Service" option > select where it says "Container image URL" > Select Artifact Registry > Select your repository and image and deploy with the following environment variables:
1. POSTGRES_DBURL(From Render).
2. NEO4J_DBURI(From Aura).
3. NEO4J_DBUSER
4. NEO4J_DBPASSWORD
5. JWT_SECRET_KEY
6. TWILIO_ACCOUNT_SID(From Twilio verify)
7. TWILIO_AUTH_TOKEN
8. TWILIO_VERIFY_SERVICE_SID
9. DEFAULT_PORT
## Database Design     
The database is layer of the application is designed to store data on entities and the details of their relationship.     
### Database Management Systems
The following database management systems are used:   
- Relational Database Management System (RDBMS).
- Graph Database Management System (GDBMS).
#### Relational Database Management System (RDBMS)    
PostgreSQL is used to manage relational data in table format.
##### Table Schema      
###### schools     
***Purpose***: Stores data for officially verified school accounts.     
***Schema***:
- id (BIGINT, PRIMARY KEY, AUTO_INCREMENT, NOT NULL)
- created_at (DATETIME, NOT NULL)
- updated_at (DATETIME, NOT NULL)
- deleted_at (DATETIME)
- name (TEXT, NOT NULL)
- phone_number (TEXT, NOT NULL, UNIQUE)
- password (TEXT, NOT NULL)
- badge (TEXT)
- website (TEXT)
###### unverified_schools    
***Purpose***: Stores data for registered but yet to be verified schools       
***Schema***:
- id (BIGINT, PRIMARY KEY, AUTO_INCREMENT, NOT NULL)
- created_at (DATETIME, NOT NULL)
- updated_at (DATETIME, NOT NULL)
- deleted_at (DATETIME)
- name (TEXT, NOT NULL)
- phone_number (TEXT, NOT NULL, UNIQUE)
- password (TEXT, NOT NULL)
- badge (TEXT)
- website (TEXT)
###### students    
***Purpose***: Stores data for student accounts that exist with school networks.     
***Schema***:
- id (BIGINT, PRIMARY KEY, AUTO_INCREMENT, NOT NULL)
- created_at (DATETIME, NOT NULL)
- updated_at (DATETIME, NOT NULL)
- deleted_at (DATETIME)
- name (TEXT, NOT NULL)
- registration_number (TEXT, NOT NULL, UNIQUE)
- phone_number (TEXT, NOT NULL)
- date_of_admission (DATETIME)
- date_of_birth (DATETIME)
- profile_picture (TEXT)
##### Data Flow     
![DatabaseDataFlow](https://github.com/user-attachments/assets/2be496b2-632f-4250-a892-b7ca4c828c8b)        
#### Graph Database Management System (GDBMS)    
Neo4j is used to manage relationship between connected data in graph format. Data representing entities is represented in form of **nodes** and the data detailing the relationships between enities is represented in form of **edges** or in other words **relationships**.
##### Node Schema    
###### School   
***Purpose***: Stores data for official school accounts.     
***Schema***:
- pk: bigint (UNIQUE)
- createdat: DateTime
- updatedat: DateTime
- name: string
- phonenumber: string (UNIQUE)
- badge: string
- website: string
###### Student    
***Purpose***: Stores data for student accounts that exist with school networks.     
***Schema***:
- pk: bigint (UNIQUE)
- createdat: DateTime
- updatedat: DateTime
- date_of_admission: DateTime
- date_of_birth: DateTime
- name: string
- phonenumber: string 
- profile_picture: string
- registration_number: string (UNIQUE)
###### STUDENT_AT  
`STUDENT_AT` has no properties. It serves the primary role of signifying that a particular `Student` goes to a particular `School`.   
## Data Synchronization     
### Purpose  
Data synchronization between PostgreSQL and Neo4j is important because it is key to leveraging the strengths of both relational and graph databases such as: 
- **Enhanced Data Flexibity**:  Neo4j, as a graph database, excels at modeling complex relationships between entities. By synchronizing data from PostgreSQL, you can leverage Neo4j's powerful graph querying capabilities to extract data from inter-dependent entities. e.g Students and Schools.
- **Improved Application Performance**: Neo4j's graph data model is optimized for traversing relationships, making it ideal for complex queries that would be inefficient in a relational database.
- **Data Integrity and Consistency**: PostgreSQL's strong data integrity features, such as robust data constraints and referential integrity, help ensure data consistency.
### Scope      
#### Data Entities   
- School data
- Student data
#### Data Systems
- PostgreSQL
- Neo4j
### Frequency   
Data is synchronized in real-time.   
### Source System
- PostgreSQL
### Target System  
- Neo4j
### Synchronization Process  
#### Trigger Mechanism    
The synchronization is dependent on event-driven triggers. The process is initiated once data is successfully inserted into the source system e.g successful insertion of school data in PostgreSQL into the school table. 
#### Data Extraction    
Once input data is successfully inserted into PostgreSQL, additional data fields such as id are generated and stored in runtime memory.  
#### Data Transformation  
Some fields are left out. Such fields include `password` and `deleted_at` fields.
#### Data Loading   
Data is loaded into the target system using incremental updates.    
#### Error Handling    
If data successfully loaded into PostgreSQL fails to be successfully loaded into Neo4j, all resolvers are blocked from receiving any incoming requests, a `Warn` message is logged and a countdown of 10 minutes is initiated after which the system will shutdown.
### Data Synchronization Diagram
![Data-Synchronization](https://github.com/user-attachments/assets/a0d947cd-2496-44d9-ba2a-574d8f814011)
## Testing and Quality Assurance   
- **GraphQL Playground**: An in-browser GraphQL IDE that allows you to explore and test your GraphQL API. It includes features like syntax highlighting, auto-complete, 
   query history, and more. GraphQL Playground was used to test query and mutation types. Responses were evaluated if they met the expected results.
## Security Considerations     
* **School Password Hashing** : plaitext passwords provided by schools is encrypted before being stored in the database
* **JSON Web Tokens**: The application leverages JSON Web Tokens to identify users that are logged into the system. Particularly, their id and role. For example: id: "1", 
   role: "school".
## Logging and Error Handling    
### Logging
Logging is done using the zerolog framework and logs are in json formats that contain the following fields:
- **level** - This field shows the severity of the log, e.g {"level":"fatal"} signifies a very severe error that an application cannot recover from. Log levels include:
  - **info** -  messages describing the normal operation of an application typically triggered by a request. e.g if a school makes a request to retrieve its students,
                the system will log a message "getting schools students" which is an informs whoever is monitoring the system of the operation currently performed.
  - **trace** - for tracing the code execution path, typically for providing information on the steps a process has reached. e.g on starting the server, a trace message
                "initializing postgresql" and then later "completed postgresql initialization" is logged. These messages informs progress on the steps of connecting to a
                postgresql database.
  - **error** -  messages describing system errors that can be easily recovered from and can be patched later in the software development cycle. e.g if a system
                is failing to generate json web tokens for some reason.
  - **Warn** - messages describing a critical error that has to be checked as soon as possible, e.g a failure of synchronizing records between postgres and neo4j that
              threaten data corruption. Such messages are usually followed by a notice of system shutdown.
  - **fatal** - messages that signifies that a certain critical error has brought an issue that cannot be recovered from and so the application has shutdown.
               Once such messages have been logged the application shuts down immediately.
- **id** - The primary key of the entity who was making the request that lead to the particular log message.
- **role** - The role of the entity who was making the request that lead to the particular log message.
- **path** - The resolver path that the request was made to.
- **time** - Time of the logging event.
- **message** - The actual message being logged.

**Note** - Other contextual information like registration_number, and phone_number might appear in log mesaages.
### Error Handling     
The following details how different scenarions of errors are handled:     
- **Minor errors** - In case of a minor error, the error message is logged with level "error" and a response detailing the error is returned to the client.
- **Critical errors** - In case of a critical error, all resolvers are blocked from receiving any incoming requests, the error message is logged with level "warn", a
                        response detailing the error is returned to the client, a shutdown countdown of ten minutes is initiated after which an the error message
                        logged again, this time with level "fatal".  