resource "elasticsearch_index_template" "events" {
  name = "events"
  body = templatefile("${path.module}/mapping.json.tmpl", {
    index_shard_count = var.index_shard_count
  })
}
