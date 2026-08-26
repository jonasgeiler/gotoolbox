package gotoolbox

// SelectByLibc returns [glibc] if glibc was detected, otherwise returns [other]
// on e.g. musl libc environments. If glibc was not explicitly detected it will
// default to [other].
func SelectByLibc(glibc, other string) string {
	if IsGlibcEnv() {
		return glibc
	}
	return other
}
