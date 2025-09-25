package surgeon

import "bytes"

// Textectomy facilitates trimming away parts of a text that the caller deems superfluous.
type Textectomy struct{ Text []byte }

// String implements fmt.Stringer
func (this *Textectomy) String() string {
	return string(this.Text)
}

func (this *Textectomy) clamp(n int) int {
	return max(0, min(n, len(this.Text)))
}
func (this *Textectomy) update(b []byte) (modified bool) {
	if bytes.Equal(this.Text, b) {
		return false
	}
	this.Text = b
	return true
}

// Contains checks whether []byte(s) is found in the underlying bytes.
func (this *Textectomy) Contains(s string) bool {
	return bytes.Contains(this.Text, []byte(s))
}

// HasPrefix checks whether []byte(s) is a prefix of the underlying bytes.
func (this *Textectomy) HasPrefix(s string) bool {
	return bytes.HasPrefix(this.Text, []byte(s))
}

// HasSuffix checks whether []byte(s) is a prefix of the underlying bytes.
func (this *Textectomy) HasSuffix(s string) bool {
	return bytes.HasSuffix(this.Text, []byte(s))
}

// Excise removes the first n occurrences of sub from the underlying bytes.
func (this *Textectomy) Excise(sub string, n int) bool {
	return this.update(bytes.Replace(this.Text, []byte(sub), []byte(""), n))
}

// ExciseAll removes all occurrences of sub from the underlying bytes.
func (this *Textectomy) ExciseAll(sub string) bool {
	return this.update(bytes.ReplaceAll(this.Text, []byte(sub), []byte("")))
}

// CutAfter cuts the string at the occurrence of s and discards all but what came before.
func (this *Textectomy) CutAfter(s string) bool {
	before, _, _ := bytes.Cut(this.Text, []byte(s))
	return this.update(before)
}

// CutBefore cuts the string at the occurrence of s and discards all but what came after.
func (this *Textectomy) CutBefore(s string) bool {
	_, after, ok := bytes.Cut(this.Text, []byte(s))
	if !ok {
		return false
	}
	return this.update(after)
}

// DiscardNLeft discards up to n bytes from the beginning of the underlying bytes or until none remain.
func (this *Textectomy) DiscardNLeft(n int) bool {
	return this.update(this.Text[this.clamp(n):])
}

// DiscardNRight discards up to n bytes from the
func (this *Textectomy) DiscardNRight(n int) bool {
	return this.update(this.Text[:len(this.Text)-this.clamp(n)])
}

// TrimLeft trims all characters in set from the beginning of the underlying bytes.
func (this *Textectomy) TrimLeft(set string) bool {
	return this.update(bytes.TrimLeft(this.Text, set))
}

// TrimRight trims all characters in set from the end of the underlying bytes.
func (this *Textectomy) TrimRight(set string) bool {
	return this.update(bytes.TrimRight(this.Text, set))
}

// CutPrefix removes the prefix (if present) from the underlying bytes.
func (this *Textectomy) CutPrefix(prefix string) (ok bool) {
	this.Text, ok = bytes.CutPrefix(this.Text, []byte(prefix))
	return ok
}

// CutSuffix removes the suffix (if present) from the underlying bytes.
func (this *Textectomy) CutSuffix(suffix string) (ok bool) {
	this.Text, ok = bytes.CutSuffix(this.Text, []byte(suffix))
	return ok
}
