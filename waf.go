package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type waf struct {
	seq int16
}

func (rcv *waf) nextID() string {
	rcv.seq++
	return fmt.Sprintf("waf%d", rcv.seq)
}

func (rcv *waf) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))

	el := node.New(cl, id,
		node.Label("FW", false),
		node.FontColor(comp.FontColor, "#fafafaff"),
		node.FillColor(comp.FillColor, "#f3190b"),
		node.Shape("invhouse"),
	)
	el.Attr("width", "0.3")
}
