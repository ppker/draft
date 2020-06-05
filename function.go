package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type function struct {
	seq int16
}

func (rcv *function) nextID() string {
	rcv.seq++
	return fmt.Sprintf("fn%d", rcv.seq)
}

func (rcv *function) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.Label(comp.Provider))

	el := node.New(cl, id,
		node.Label(comp.Label),
		node.Rounded(comp.Rounded),
		node.FontColor(comp.FontColor),
		node.FillColor(comp.FillColor, "#abd9e9ff"),
		node.Shape("signature"),
	)
	el.Attr("fontsize", "6")
	el.Attr("height", "0.5")
}
