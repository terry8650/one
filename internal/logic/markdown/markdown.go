package markdown

import (
	"context"
	"github.com/88250/lute"
	"one2.3/internal/service"
)

type sMarkdown struct {
	*lute.Lute
}

func init() {
	service.RegisterMarkdown(New())
}
func New() *sMarkdown {
	luteEngine := lute.New()
	return &sMarkdown{luteEngine}
}
func (s *sMarkdown) ToHtml(ctx context.Context, md string) (ht string) {
	ht = s.Lute.MarkdownStr("", md)
	return
}
