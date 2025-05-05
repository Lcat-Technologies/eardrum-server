/*add a unique constraint to unverified schools' phone_number*/

ALTER TABLE unverified_schools
ADD CONSTRAINT unique_phone_number UNIQUE (phone_number);

/*add a unique constraint to schools' phone_number*/

ALTER TABLE schools
ADD CONSTRAINT unique_school_phone_number UNIQUE (phone_number);

/*add a unique constraint to  students' registration_number*/

ALTER TABLE students
ADD CONSTRAINT unique_registration_number UNIQUE (registration_number);