package wrappers

import (
	"time"

	"github.com/AlekSi/pointer"
	"github.com/GigaDesk/eardrum-server/graph/model"
)

type Neo4jStudentWrapper struct{
*model.Student
}

// Returns the unique identifier of the student
func (n Neo4jStudentWrapper) GetID() int64 {
return int64(n.ID)
} 
// Returns the creation timestamp of the student           
func (n Neo4jStudentWrapper) GetCreatedAt() time.Time{
return n.CreatedAt.UTC()
} 
// Returns the last update timestamp of the student
func (n Neo4jStudentWrapper) GetUpdatedAt() time.Time {
return n.UpdatedAt.UTC()
}
// Returns the student's registration number
func (n Neo4jStudentWrapper) GetRegistrationNumber() string {
return n.RegistrationNumber
}
// Returns the name of the student
func (n Neo4jStudentWrapper) GetName() string {
return n.Name
} 
// Returns the phone number of the student
func (n Neo4jStudentWrapper) GetPhoneNumber() string {
return n.PhoneNumber
}
// Returns the password associated with the student (e.g., for student access)
func (n Neo4jStudentWrapper) GetPassword() string {
return n.Password
}
// Returns the student's date of admission
func (n Neo4jStudentWrapper) GetDateOfAdmission()  time.Time {
return pointer.GetTime(n.DateOfAdmission).UTC()
} 
// Returns the student's birthday
func (n Neo4jStudentWrapper) GetDateofBirth() time.Time {
return pointer.GetTime(n.DateOfBirth).UTC()
}
//Returns the student profile picture's image url
func (n Neo4jStudentWrapper) GetProfilePicture() string {
return  pointer.GetString(n.ProfilePicture)
}