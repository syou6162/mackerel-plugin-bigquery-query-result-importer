package mpbigqueryqueryresultimporter

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"cloud.google.com/go/bigquery"
	mp "github.com/mackerelio/go-mackerel-plugin-helper"
	"google.golang.org/api/iterator"
)

type BigQueryQueryResultImporterPlugin struct {
	ProjectId string
	Query     string
	GraphName string
	IsStacked bool
}

// GraphDefinition of BigQueryQueryResultImporterPlugin
func (p BigQueryQueryResultImporterPlugin) GraphDefinition() map[string](mp.Graphs) {
	graphdef := map[string](mp.Graphs){
		fmt.Sprintf("bigquery_query_result_importer.%s.#", p.GraphName): mp.Graphs{
			Label: strings.Title(strings.ReplaceAll(p.GraphName, "_", " ")),
			Unit:  "float",
			Metrics: [](mp.Metrics){
				mp.Metrics{Name: "value", Label: "Value", Stacked: p.IsStacked},
			},
		},
	}
	return graphdef
}

type Metric struct {
	Label string
	Value float64
}

// FetchMetrics fetch the metrics
func (p BigQueryQueryResultImporterPlugin) FetchMetrics() (map[string]interface{}, error) {
	stat := make(map[string]interface{})
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, p.ProjectId)
	if err != nil {
		return nil, err
	}

	it, err := client.Query(p.Query).Read(ctx)
	if err != nil {
		return nil, err
	}

	for {
		var m Metric
		err := it.Next(&m)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		k := fmt.Sprintf("bigquery_query_result_importer.%s.%s.value", p.GraphName, m.Label)
		stat[k] = m.Value
	}
	return stat, nil
}

// Do the plugin
func Do() {
	var plugin BigQueryQueryResultImporterPlugin

	projectId := flag.String("project_id", "", "GCP Project ID")
	graphName := flag.String("graph_name", "sample", "Graph Name")
	query := flag.String("query", "SELECT 'sample_key' AS Label, 1 AS Value", "Query for BigQuery")
	isStacked := flag.Bool("is_stacked", true, "Graph Definition of Stacked")
	flag.Parse()

	plugin.ProjectId = *projectId
	plugin.GraphName = *graphName
	plugin.Query = *query
	plugin.IsStacked = *isStacked

	helper := mp.NewMackerelPlugin(plugin)
	helper.Run()
}
