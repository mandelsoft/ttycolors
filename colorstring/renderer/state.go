package renderer

import (
	"github.com/mandelsoft/ttycolors/ansi"
)

type state struct {
	FG           ansi.Code
	FGRendered   bool
	EndFGpending bool
	BG           ansi.Code
	BGRendered   bool
	EndBGPending bool
}

func (s *state) GetFG() ansi.Code {
	if s != nil {
		return s.FG
	}
	return nil
}

func (s *state) GetBG() ansi.Code {
	if s != nil {
		return s.BG
	}
	return nil
}

func (s *state) SetFG(c ansi.Code) *state {
	if s != nil {
		return &state{
			FG:         c,
			BG:         s.BG,
			BGRendered: s.BGRendered,
		}
	}
	return &state{
		FG: c,
	}
}

func (s *state) SetBG(c ansi.Code) *state {
	if s != nil {
		return &state{
			FG:         s.FG,
			FGRendered: s.FGRendered,
			BG:         c,
		}
	}
	return &state{
		BG: c,
	}
}

func (s *state) RepeatAll() string {
	if s == nil {
		return ""
	}
	var seq ansi.Sequence

	if s.FG != nil && !s.FGRendered {
		s.FGRendered = true
		seq = seq.Append(s.FG)
	} else {
		if s.EndFGpending {
			s.EndFGpending = false
			seq = seq.Append(ansi.COLOR_FG_DEFAULT)
		}
	}

	if s.BG != nil && !s.BGRendered {
		s.BGRendered = true
		seq = seq.Append(s.BG)
	} else {
		if s.EndBGPending {
			s.EndBGPending = false
			seq = seq.Append(ansi.COLOR_BG_DEFAULT)
		}
	}

	if seq == nil {
		return ""
	}

	return ansi.Mode(seq)
}

func (s *state) RepeatNonBG() string {
	if s == nil {
		return ""
	}
	if s.FG != nil && !s.FGRendered {
		s.FGRendered = true
		return ansi.Mode(s.FG)
	}
	return s.EmitEndFG()
}

func (s *state) RepeatNonFG() string {
	if s == nil {
		return ""
	}
	if s.BG != nil && !s.BGRendered {
		s.BGRendered = true
		return ansi.Mode(s.BG)
	}
	return s.EmitEndBG()
}

func (s *state) RequestEndBG() {
	if s == nil || s.BG != nil {
		return
	}
	s.EndBGPending = true
}

func (s *state) EmitEndBG() string {
	if s == nil || !s.EndBGPending {
		return ""
	}
	s.EndBGPending = false
	return ansi.Mode(ansi.COLOR_BG_DEFAULT)
}

func (s *state) RequestEndFG() {
	if s == nil || s.FG != nil {
		return
	}
	s.EndFGpending = true
}

func (s *state) EmitEndFG() string {
	if s == nil || !s.EndFGpending {
		return ""
	}
	return ansi.Mode(ansi.COLOR_FG_DEFAULT)
}
