package lib

import "bytes"

// Stringectomy facilitates trimming away parts of a text that the caller deems superfluous.
type Stringectomy struct{ b []byte }

func NewStringectomy(b []byte) *Stringectomy { return &Stringectomy{b: b} }

func clamp(lo, n, hi int) int {
	return max(lo, min(n, hi))
}

func (this *Stringectomy) update(b []byte) (modified bool) {
	if bytes.Equal(this.b, b) {
		return false
	}
	this.b = b
	return true
}

func (this *Stringectomy) Bytes() []byte { return this.b }

func (this *Stringectomy) Excise(sub string, n int) bool {
	return this.update(bytes.Replace(this.b, []byte(sub), []byte(""), n))
}
func (this *Stringectomy) ExciseAll(sub string) bool {
	return this.update(bytes.ReplaceAll(this.b, []byte(sub), []byte("")))
}

func (this *Stringectomy) CutAfter(s string) bool {
	before, _, _ := bytes.Cut(this.b, []byte(s))
	return this.update(before)
}
func (this *Stringectomy) CutBefore(s string) bool {
	_, after, ok := bytes.Cut(this.b, []byte(s))
	if !ok {
		return false
	}
	return this.update(after)
}

func (this *Stringectomy) DiscardNLeft(n int) bool {
	return this.update(this.b[clamp(0, n, len(this.b)):])
}
func (this *Stringectomy) DiscardNRight(n int) bool {
	return this.update(this.b[:len(this.b)-clamp(0, n, len(this.b))])
}

func (this *Stringectomy) TrimLeft(set string) bool {
	return this.update(bytes.TrimLeft(this.b, set))
}
func (this *Stringectomy) TrimRight(set string) bool {
	return this.update(bytes.TrimRight(this.b, set))
}

func (this *Stringectomy) CutPrefix(prefix string) (ok bool) {
	this.b, ok = bytes.CutPrefix(this.b, []byte(prefix))
	return ok
}
func (this *Stringectomy) CutSuffix(suffix string) (ok bool) {
	this.b, ok = bytes.CutSuffix(this.b, []byte(suffix))
	return ok
}
