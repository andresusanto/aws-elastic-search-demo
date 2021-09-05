# Elastic Search Index Strategy

**Considerations taken into account:**

1. Time-series based data, old-data may lose relevance overtime.
2. `user_id` field is a very high cardinality field.

**Design Decision:**

![image](https://user-images.githubusercontent.com/7076809/132135012-f6119ec7-0afa-4858-970f-6b51a323640e.png)

To achieve highest performance and maintainability, the above design is chosen. There are three important components in the design:

1. **Indices:** instead of using one large index with high number of shards, smaller but more indices are used. As the new arriving data is placed into smaller indices, [global ordinal calculation](https://www.elastic.co/guide/en/elasticsearch/reference/current/eager-global-ordinals.html#_what_are_global_ordinals) is less expensive as the number of shards involved in the calculation is less compared to using one big index.
2. **Alias:** To make data querying easy, instead of querying to the index directly (e.g `/events_2021_09_XX`), `alias` is used so clients can query `/events` instead.
3. **Index Template:** To automatically create an index with correct mapping and settings and attach it automatically to the alias, an index template is used.

### Index Template Definition

The index template definition can be found [here](./mapping.json.tmpl).
