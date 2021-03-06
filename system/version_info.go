package system

/**
 * The [VersionInfo] type...
 */
type VersionInfo struct {
  Learn Version
}

/**
 * The [Version] type...
 */
type Version struct {
  Major, Minor, Patch float64

  Build string
}
