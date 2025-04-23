package orc

import "github.com/scritchley/orc/proto"

func (r *Reader) PostScript() *proto.PostScript {
	return r.postScript
}
