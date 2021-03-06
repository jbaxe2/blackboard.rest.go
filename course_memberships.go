package blackboard_rest

import (
  "net/url"
  "strings"

  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/factory"
  "github.com/jbaxe2/blackboard.rest/course_memberships"
  "github.com/jbaxe2/blackboard.rest/oauth2"
)

/**
 * The [CourseMemberships] interface...
 */
type CourseMemberships interface {
  GetMembershipsForCourse (
    courseId string,
  ) ([]course_memberships.Membership, error)

  GetMembershipsForUser (userId string) ([]course_memberships.Membership, error)

  GetMembership (
    courseId string, userId string,
  ) (course_memberships.Membership, error)

  UpdateMembership (
    courseId string, userId string, membership course_memberships.Membership,
  ) error

  CreateMembership (
    courseId string, userId string, membership course_memberships.Membership,
  ) error
}

/**
 * The [_BbRestCourseMemberships] type...
 */
type _BbRestCourseMemberships struct {
  _BlackboardRest

  CourseMemberships
}

/**
 * The [GetCourseMembershipsInstance] function...
 */
func GetCourseMembershipsInstance (
  host string, accessToken oauth2.AccessToken,
) CourseMemberships {
  hostUri, _ := url.Parse (host)

  membershipsService := new (_BbRestCourseMemberships)

  membershipsService.host = *hostUri
  membershipsService.accessToken = accessToken

  membershipsService.service.SetHost (host)
  membershipsService.service.SetAccessToken (accessToken)

  return membershipsService
}

/**
 * The [GetMembershipsForCourse] method...
 */
func (restMemberships *_BbRestCourseMemberships) GetMembershipsForCourse (
  courseId string,
) ([]course_memberships.Membership, error) {
  endpoint := config.CourseMembershipsEndpoints["course_memberships"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)

  return restMemberships._getMemberships (endpoint, make (map[string]interface{}))
}

/**
 * The [GetMembershipsForUser] method...
 */
func (restMemberships *_BbRestCourseMemberships) GetMembershipsForUser (
  userId string,
) ([]course_memberships.Membership, error) {
  endpoint := config.CourseMembershipsEndpoints["user_memberships"]
  endpoint = strings.Replace (endpoint, "{userId}", userId, -1)

  return restMemberships._getMemberships (endpoint, make (map[string]interface{}))
}

/**
 * The [GetMembership] method...
 */
func (restMemberships *_BbRestCourseMemberships) GetMembership (
  courseId string, userId string,
) (course_memberships.Membership, error) {
  var courseMembership course_memberships.Membership

  endpoint := config.CourseMembershipsEndpoints["membership"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)
  endpoint = strings.Replace (endpoint, "{userId}", userId, -1)

  result, err := restMemberships.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 1,
  )

  if nil != err {
    return courseMembership, err
  }

  courseMembership = factory.NewMembership (result.(map[string]interface{}))

  return courseMembership, err
}

/**
 * The [_getMemberships] method...
 */
func (restMemberships *_BbRestCourseMemberships) _getMemberships (
  endpoint string, data map[string]interface{},
) ([]course_memberships.Membership, error) {
  result, err := restMemberships.service.Connector.SendBbRequest (
    endpoint, "GET", data, 1,
  )

  if nil != err {
    return []course_memberships.Membership{}, err
  }

  rawMemberships := result.(map[string]interface{})["results"]

  courseMemberships := factory.NewMemberships (
    _normalizeRawMemberships(rawMemberships.([]interface{})),
  )

  return courseMemberships, err
}

/**
 * The [_normalizeRawMemberships] function...
 */
func _normalizeRawMemberships (
  rawMemberships []interface{},
) []map[string]interface{} {
  mappedRawMemberships := make ([]map[string]interface{}, len (rawMemberships))

  for i, rawMembership := range rawMemberships {
    mappedRawMemberships[i] = rawMembership.(map[string]interface{})
  }

  return mappedRawMemberships
}
