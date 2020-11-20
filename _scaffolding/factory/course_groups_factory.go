package factory

import (
  "github.com/google/uuid"
  "github.com/jbaxe2/blackboard.rest/course_groups"
)

/**
 * The [NewCourseGroups] function...
 */
func NewCourseGroups (
  rawCourseGroups []map[string]interface{},
) []course_groups.Group {
  courseGroups := make ([]course_groups.Group, len (rawCourseGroups))

  for i, rawCourseGroup := range rawCourseGroups {
    courseGroups[i] = NewCourseGroup (rawCourseGroup)
  }

  return courseGroups
}

/**
 * The [NewCourseGroup] function...
 */
func NewCourseGroup (rawCourseGroup map[string]interface{}) course_groups.Group {
  groupUuid, _ := uuid.Parse (rawCourseGroup["uuid"].(string))
  description, _ := rawCourseGroup["description"].(string)

  return course_groups.Group {
    Id: rawCourseGroup["id"].(string),
    ExternalId: rawCourseGroup["externalId"].(string),
    GroupSetId: rawCourseGroup["groupSetId"].(string),
    Name: rawCourseGroup["name"].(string),
    Description: description,
    Uuid: groupUuid,
  }
}
