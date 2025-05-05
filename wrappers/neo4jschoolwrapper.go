package wrappers

import (
	"time"

	"github.com/AlekSi/pointer"
	"github.com/GigaDesk/eardrum-server/graph/model"
)

//Wraps a School struct and provides additional functionality to meet the `Neo4jSchool` interface.
type Neo4jSchoolWrapper struct{
*model.School
}

// Returns the unique identifier of the school
func (n Neo4jSchoolWrapper)GetID() int64  {
	return int64(n.ID)
}   
// Returns the creation timestamp of the school       
func (n Neo4jSchoolWrapper) GetCreatedAt() time.Time{
	return n.CreatedAt.UTC()
}

// Returns the last update timestamp of the school
func (n Neo4jSchoolWrapper) GetUpdatedAt() time.Time {
	return n.UpdatedAt.UTC()
}

// Returns the name of the school
func (n Neo4jSchoolWrapper) GetName() string{
	return n.Name
}

// Returns the phone number of the school
func (n Neo4jSchoolWrapper) GetPhoneNumber() string {
	return n.PhoneNumber
}

// Returns the password associated with the school (e.g., for admin access)
func (n Neo4jSchoolWrapper) GetPassword() string {
	return n.Password
}  

// Returns school badge's image url
func (n Neo4jSchoolWrapper) GetBadge() string  {
	return pointer.GetString(n.Badge)
} 

// Returns the website URL of the school
func (n Neo4jSchoolWrapper) GetWebsite() string  {
	return pointer.GetString(n.Website)
}