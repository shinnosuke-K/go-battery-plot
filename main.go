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

func main() {
	file, err := os.Open("./data.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvFile := csv.NewReader(file)

	data, err := csvFile.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	pts := make(plotter.XYs, len(data))

	for x, y := range data {
		pts[x].X = float64(x)
		pts[x].Y, err = strconv.ParseFloat(y[0], 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	p.Title.Text = "Current Capacity"
	p.Y.Label.Text = "Capacity"
	p.Y.Max = 4500

	err = plotutil.AddLinePoints(p, pts)
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
