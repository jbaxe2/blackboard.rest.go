package course_groups

import (
  "time"

  "github.com/google/uuid"
)

/**
 * The [Group] type...
 */
type Group struct {
  Id, ExternalId, GroupSetId, Name, Description string

  Uuid uuid.UUID

  Availability GroupAvailability

  Enrollment GroupEnrollment

  Created time.Time

  Modified time.Time
}

/**
 * The [GroupAvailability] type...
 */
type GroupAvailability struct {
  Available GroupAvailable
}

/**
 * The [GroupEnrollment] type...
 */
type GroupEnrollment struct {
  Type GroupEnrollmentType

  Limit int

  SignupSheet SignupSheet
}

/**
 * The [SignupSheet] type...
 */
type SignupSheet struct {
  Name, Description string

  ShowMembers bool
}

/**
 * The [GroupAvailable] type...
 */
type GroupAvailable string

const (
  Yes        GroupAvailable = "Yes"
  No         GroupAvailable = "No"
  SignupOnly GroupAvailable = "SignupOnly"
)

/**
 * The [GroupEnrollmentType] type...
 */
type GroupEnrollmentType string

const (
  InstructorOnly GroupEnrollmentType = "InstructorOnly"
  SelfEnroll     GroupEnrollmentType = "SelfEnroll"
)
