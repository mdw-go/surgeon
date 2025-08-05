package surgeon

import "bytes"

// Stringectomy facilitates trimming away parts of a text that the caller deems superfluous.
type Stringectomy struct{ Text []byte }

func clamp(lo, n, hi int) int {
	return max(lo, min(n, hi))
}

func (this *Stringectomy) update(b []byte) (modified bool) {
	if bytes.Equal(this.Text, b) {
		return false
	}
	this.Text = b
	return true
}

// Contains checks whether []byte(s) is found in the underlying bytes.
func (this *Stringectomy) Contains(s string) bool {
	return bytes.Contains(this.Text, []byte(s))
}

// Excise removes the first n occurrences of sub from the underlying bytes.
func (this *Stringectomy) Excise(sub string, n int) bool {
	return this.update(bytes.Replace(this.Text, []byte(sub), []byte(""), n))
}

// ExciseAll removes all occurrences of sub from the underlying bytes.
func (this *Stringectomy) ExciseAll(sub string) bool {
	return this.update(bytes.ReplaceAll(this.Text, []byte(sub), []byte("")))
}

// CutAfter cuts the string at the occurrence of s and discards all but what came before.
func (this *Stringectomy) CutAfter(s string) bool {
	before, _, _ := bytes.Cut(this.Text, []byte(s))
	return this.update(before)
}

// CutBefore cuts the string at the occurrence of s and discards all but what came after.
func (this *Stringectomy) CutBefore(s string) bool {
	_, after, ok := bytes.Cut(this.Text, []byte(s))
	if !ok {
		return false
	}
	return this.update(after)
}

// DiscardNLeft discards up to n bytes from the beginning of the underlying bytes or until none remain.
func (this *Stringectomy) DiscardNLeft(n int) bool {
	return this.update(this.Text[clamp(0, n, len(this.Text)):])
}

// DiscardNRight discards up to n bytes from the
func (this *Stringectomy) DiscardNRight(n int) bool {
	return this.update(this.Text[:len(this.Text)-clamp(0, n, len(this.Text))])
}

// TrimLeft trims all characters in set from the beginning of the underlying bytes.
func (this *Stringectomy) TrimLeft(set string) bool {
	return this.update(bytes.TrimLeft(this.Text, set))
}

// TrimRight trims all characters in set from the end of the underlying bytes.
func (this *Stringectomy) TrimRight(set string) bool {
	return this.update(bytes.TrimRight(this.Text, set))
}

// CutPrefix removes the prefix (if present) from the underlying bytes.
func (this *Stringectomy) CutPrefix(prefix string) (ok bool) {
	this.Text, ok = bytes.CutPrefix(this.Text, []byte(prefix))
	return ok
}

// CutSuffix removes the suffix (if present) from the underlying bytes.
func (this *Stringectomy) CutSuffix(suffix string) (ok bool) {
	this.Text, ok = bytes.CutSuffix(this.Text, []byte(suffix))
	return ok
}
