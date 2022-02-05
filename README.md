# mackerel-plugin-bigquery-query-result-importer

## Synopsis

```console
% mackerel-plugin-bigquery-query-result-importer --project_id PROJECT --graph_name GRAPH_NAME --query QUERY
```

## Example of mackerel-agent.conf
Be careful about BigQuery Billing. Details and useful query examples are written in [blog entry](https://www.yasuhisay.info/entry/mackerel-plugin-bigquery-query-result-importer).

```conf
[plugin.metrics.bigquery-query-result-importer]
command = "/path/to/mackerel-plugin-bigquery-query-result-importer --project_id my-project --graph_name bigquery_query_sample --query \"SELECT 'hoge' AS Label, 12.3 AS Value UNION ALL SELECT 'fuga' AS Label, 45.6 AS Value UNION ALL SELECT 'piyo' AS Label, 78.9 AS Value\""
```
