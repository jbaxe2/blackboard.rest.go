package test

import (
  "github.com/jbaxe2/blackboard.rest.go/src"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  error2 "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/error"
  "github.com/jbaxe2/blackboard.rest.go/src/course_memberships"
  "testing"
)

/**
 * The [CourseMembershipsTester] type...
 */
type CourseMembershipsTester struct {
  t *testing.T

  Testable
}

/**
 * The [Run] method...
 */
func (tester *CourseMembershipsTester) Run() {
  println ("\nCourse Memberships:")

  _testGetCourseMembershipsInstance (tester.t)
  _testGetCourseMembershipsByCoursePrimaryId (tester.t)
  _testGetCourseMembershipsByUserPrimaryId (tester.t)
  _testGetMembershipByCourseAndUserPrimaryIds (tester.t)
}

/**
 * The [_getCourseMembershipsInstance] function...
 */
func _getCourseMembershipsInstance() blackboard_rest.CourseMemberships {
  authorizer := TestersAuthorizer{}
  _ = authorizer.AuthorizeForTests()

  return blackboard_rest.GetCourseMembershipsInstance (
    config.Host, authorizer.accessToken,
  )
}

/**
 * The [_testGetCourseMembershipsInstance] function...
 */
func _testGetCourseMembershipsInstance (t *testing.T) {
  println ("Obtain a valid CourseMemberships service instance.")

  courseMembershipsService := _getCourseMembershipsInstance()

  if nil == courseMembershipsService {
    t.Error ("Obtaining a valid CourseMemberships instance failed.\n")
  }
}

/**
 * The [_testGetCourseMembershipsByCoursePrimaryId] function...
 */
func _testGetCourseMembershipsByCoursePrimaryId (t *testing.T) {
  println ("Get a list of course memberships by the course primary ID.")

  membershipsService := _getCourseMembershipsInstance()

  memberships, err := membershipsService.GetMembershipsForCourse ("_121_1")

  if (nil == memberships) || (error2.RestError{} != err) {
    t.Error ("Failed to obtain the list of course memberships (course).\n")

    return
  }

  if 0 == len (memberships) {
    t.Error ("Retrieved an empty list of enrollments that should not be empty.")

    return
  }

  for _, membership := range memberships {
    if "_121_1" != membership.CourseId {
      t.Error ("Membership retrieved does not match what was specified.")

      return
    }
  }
}

/**
 * The [_testGetCourseMembershipsByUserPrimaryId] function...
 */
func _testGetCourseMembershipsByUserPrimaryId (t *testing.T) {
  println ("Get a list of course memberships by the user primary ID.")

  membershipsService := _getCourseMembershipsInstance()

  memberships, err := membershipsService.GetMembershipsForUser ("_27_1")

  if (nil == memberships) || (error2.RestError{} != err) {
    t.Error ("Failed to obtain the list of course memberships (user).\n")

    return
  }

  if 0 == len (memberships) {
    t.Error ("Retrieved an empty list of enrollments that should not be empty.")

    return
  }

  for _, membership := range memberships {
    if "_27_1" != membership.UserId {
      t.Error ("Membership retrieved does not match what was specified.")

      return
    }
  }
}

/**
 * The [_testGetMembershipByCourseAndUserPrimaryIds] function...
 */
func _testGetMembershipByCourseAndUserPrimaryIds (t *testing.T) {
  println ("Get a membership by the course and user primary ID's.")

  membershipsService := _getCourseMembershipsInstance()

  membership, err := membershipsService.GetMembership ("_121_1", "_27_1")

  if (course_memberships.Membership{} == membership) ||
     (nil != err) ||
     (error2.RestError{} != err) {
    t.Error ("Failed to obtain the membership for the course and user.")

    return
  }

  if ("_121_1" != membership.CourseId) && ("_27_1" != membership.UserId) {
    t.Error ("Membership retrieved does not match what was specified.")
  }
}
