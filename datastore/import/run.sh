export DATASTORE_PROJECT_ID

curl -X POST localhost:8081/v1/projects/${DATASTORE_PROJECT_ID}:import \
    -H 'Content-Type: application/json' \
    -d '{"input_url":"/datastore/import/2020-01-21.overall_export_metadata"}'
