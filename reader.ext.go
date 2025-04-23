package orc

import "github.com/scritchley/orc/proto"

func (r *Reader) PostScript() *proto.PostScript {
	return r.postScript
}

func (r *Reader) Footer() *proto.Footer {
	return r.footer
}

func (r *Reader) StripeInformation() *proto.StripeInformation {
	return r.currentStripeInformation
}
