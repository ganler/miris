package main

import (
	"github.com/favyen/miris/data"
	"github.com/favyen/miris/miris"
	"github.com/favyen/miris/planner"

	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	predName := os.Args[1]
	freq, _ := strconv.Atoi(os.Args[2])
	bound, _ := strconv.ParseFloat(os.Args[3], 64)

	var existingPlan miris.PlannerConfig
	var qSamples map[int][]float64
	if len(os.Args) >= 5 {
		miris.ReadJSON(os.Args[4], &existingPlan)
		qSamples = existingPlan.QSamples
	}

	ppCfg, modelCfg := data.Get(predName)
	if predName == "beach-runner" {
		fmt.Println(modelCfg)
		for idx, _ := range modelCfg.Filters {
			modelCfg.Filters[idx].Cfg["model_path"] = strings.ReplaceAll(modelCfg.Filters[idx].Cfg["model_path"], "beach-runner", "beach")
		}
		for idx, _ := range modelCfg.Refiners {
			modelCfg.Refiners[idx].Cfg["model_path"] = strings.ReplaceAll(modelCfg.Refiners[idx].Cfg["model_path"], "beach-runner", "beach")
		}
		predName = "beach"
		fmt.Println(modelCfg)
	}

	if qSamples == nil {
		qSamples = planner.GetQSamples(2*freq, ppCfg, modelCfg)
	}
	q := planner.PlanQ(qSamples, bound)
	log.Println("finished planning q", q)
	plan := miris.PlannerConfig{
		Freq: freq,
		Bound: bound,
		QSamples: qSamples,
		Q: q,
	}
	miris.WriteJSON(fmt.Sprintf("logs/%s/%d/%v/plan.json", predName, freq, bound), plan)
	filterPlan, refinePlan := planner.PlanFilterRefine(ppCfg, modelCfg, freq, bound, nil)
	plan.Filter = filterPlan
	plan.Refine = refinePlan
	log.Println(plan)
	miris.WriteJSON(fmt.Sprintf("logs/%s/%d/%v/plan.json", predName, freq, bound), plan)
}

