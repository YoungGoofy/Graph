package graph

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type Graph struct {
	nodes []node
	edges []edge
}

func (g *Graph) Print() {
	fmt.Println("NODES")
	g.PrintNodes()
	fmt.Println("\nEDGES")
	g.PrintEdges()
}

func (g *Graph) DrawGraph() *image.RGBA {
	// Find node with maximum edges
	maxNode := findMaxEdgesNode(g.nodes)

	// Set canvas size
	size := int(math.Sqrt(float64(len(g.nodes)))) * 1500
	canvas := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Calculate node positions
	theta := (2 * math.Pi) / float64(len(g.nodes))
	for i := range g.nodes {
		if i == maxNode.ID {
			g.nodes[i].X = size / 2
			g.nodes[i].Y = size / 2
		} else {
			g.nodes[i].X = int(float64(size)/2 + math.Cos(float64(i)*theta)*float64(size)*0.4)
			g.nodes[i].Y = int(float64(size)/2 + math.Sin(float64(i)*theta)*float64(size)*0.4)
		}
	}

	// Draw edges
	for _, edge := range g.edges {
		start := g.nodes[edge.StartID-1]
		end := g.nodes[edge.EndID-1]
		drawLine(canvas, start.X, start.Y, end.X, end.Y)
	}

	// Draw nodes
	for _, node := range g.nodes {
		drawCircle(canvas, node.X, node.Y, 40)
		drawCircle(canvas, node.X, node.Y, 36)
		drawText(canvas, &node)
	}

	return canvas
}

func findMaxEdgesNode(nodes []node) node {
	var maxNode node
	maxEdges := 0

	for _, node := range nodes {
		if len(node.Edges) > maxEdges {
			maxEdges = len(node.Edges)
			maxNode = node
		}
	}

	return maxNode
}

func drawLine(img *image.RGBA, x0, y0, x1, y1 int) {
	clr := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	dx := abs(x1 - x0)
	dy := -abs(y1 - y0)
	sx := 1
	sy := 1
	if x0 > x1 {
		sx = -1
	}
	if y0 > y1 {
		sy = -1
	}
	err := dx + dy
	for {
		img.Set(x0, y0, clr)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x0 += sx
		}
		if e2 <= dx {
			err += dx
			y0 += sy
		}
	}
}

func drawCircle(img *image.RGBA, x, y, r int) {
	for i := x - r; i <= x+r; i++ {
		for j := y - r; j <= y+r; j++ {
			if (i-x)*(i-x)+(j-y)*(j-y) < r*r {
				img.Set(i, j, color.Black)
			}
		}
	}
}

func drawText(img *image.RGBA, node *node) {
	col := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	point := fixed.Point26_6{
		X: fixed.Int26_6(node.X * 64),
		Y: fixed.Int26_6(node.Y * 64),
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(node.Name)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
