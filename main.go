package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Data struct {
	d [][]string
}

func (d Data) createPoint(n int) *plotter.XYs {
	var err error
	pts := make(plotter.XYs, len(d.d))
	for x, y := range d.d {
		pts[x].X = float64(x)
		pts[x].Y, err = strconv.ParseFloat(y[n], 64)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &pts
}

func main() {
	file, err := os.Open("./data.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvFile := csv.NewReader(file)

	var csvData Data

	csvData.d, err = csvFile.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Current Capacity"
	p.Y.Label.Text = "Capacity"
	p.Y.Max = 4500

	err = plotutil.AddLinePoints(p, "Current", csvData.createPoint(0))
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
