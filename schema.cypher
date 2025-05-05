//Create a constraint requiring School nodes to have unique pk properties
CREATE CONSTRAINT school_pk
FOR (school:School) REQUIRE school.pk IS UNIQUE;

//Create a constraint requiring School nodes to have unique phonenumber properties
CREATE CONSTRAINT school_phonenumber
FOR (school:School) REQUIRE school.phonenumber IS UNIQUE;

//Create a constraint requiring Student nodes to have unique pk properties
CREATE CONSTRAINT student_pk
FOR (student:Student) REQUIRE student.pk IS UNIQUE;

//Create a constraint requiring Student nodes to have unique registration_number properties
CREATE CONSTRAINT student_registration_number
FOR (student:Student) REQUIRE student.registration_number IS UNIQUE;