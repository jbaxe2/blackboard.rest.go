package factory

import (
  "github.com/jbaxe2/blackboard.rest.go/src/terms"
  "time"
)

/**
 * The [NewTerms] function...
 */
func NewTerms (rawTerms []map[string]interface{}) []terms.Term {
  theTerms := make ([]terms.Term, len (rawTerms))

  for i, rawTerm := range rawTerms {
    theTerms[i] = NewTerm (rawTerm)
  }

  return theTerms
}

/**
 * The [NewTerm] function...
 */
func NewTerm (rawTerm map[string]interface{}) terms.Term {
  return terms.Term {
    Id: rawTerm["id"].(string),
    ExternalId: rawTerm["externalId"].(string),
    Name: rawTerm["name"].(string),
    Availability: _parseTermAvailability (rawTerm["availability"]),
  }
}

/**
 * The [_parseTermAvailability] function...
 */
func _parseTermAvailability (rawAvailability interface{}) terms.TermAvailability {
  mappedAvailability := rawAvailability.(map[string]interface{})

  return terms.TermAvailability {
    Available: terms.Availability (mappedAvailability["available"].(string)),
    Duration: _parseDuration (mappedAvailability["duration"]),
  }
}

/**
 * The [_parseDuration] function...
 */
func _parseDuration (rawDuration interface{}) terms.TermDuration {
  mappedDuration := rawDuration.(map[string]interface{})

  start, _ := time.Parse (time.RFC3339, mappedDuration["start"].(string))
  end, _ := time.Parse (time.RFC3339, mappedDuration["end"].(string))

  return terms.TermDuration {
    Type: terms.DurationType (mappedDuration["type"].(string)),
    Start: start,
    End: end,
    DaysOfUse: mappedDuration["daysOfUse"].(int),
  }
}
