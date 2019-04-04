package test

import (
  "github.com/jbaxe2/blackboard.rest.go/src"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  error2 "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/error"
  "testing"
)

/**
 * The [UsersTester] type...
 */
type UsersTester struct {
  t *testing.T

  Testable
}

/**
 * The [Run] method...
 */
func (tester *UsersTester) Run() {
  println ("\nUsers:")

  _testGetUsersInstance (tester.t)
  _testGetUserByPrimaryId (tester.t)
}

/**
 * The [_getUsersInstance] function...
 */
func _getUsersInstance() blackboard_rest.Users {
  authorizer := TestersAuthorizer{}
  _ = authorizer.AuthorizeForTests()

  return blackboard_rest.GetUsersInstance (
    config.Host, authorizer.accessToken,
  )
}

/**
 * The [_testGetUsersInstance] function...
 */
func _testGetUsersInstance (t *testing.T) {
  println ("Obtain a valid Users service instance.")

  usersService := _getUsersInstance()

  if nil == usersService {
    t.Error ("Obtaining a valid Users service instance failed.\n")
  }
}

/**
 * The [_testGetUserByPrimaryId] function...
 */
func _testGetUserByPrimaryId (t *testing.T) {
  println ("Get a user by his or her primary ID.")

  usersService := _getUsersInstance()
  user, err := usersService.GetUser ("_27_1")

  if (nil != err) && (error2.RestError{} != err) {
    t.Error ("Error while retrieving the user:\n" + err.Error())

    return
  }

  if "_27_1" != user.Id {
    t.Error ("The retrieved user does not match the one selected.")
  }
}
