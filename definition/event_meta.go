package definition

import csndra "github.com/TerrexTech/go-cassandrautils/cassandra"

var eventMeta map[string]csndra.TableColumn

// EventMeta returns the defintion for the event_meta table.
// EventMeta table tracks the hydration for aggregate-projections
// using versions.
func EventMeta() map[string]csndra.TableColumn {
	if eventMeta == nil {
		eventMeta = map[string]csndra.TableColumn{
			"aggregateVersion": csndra.TableColumn{
				Name:            "aggregate_version",
				DataType:        "bigint",
				PrimaryKeyIndex: "2",
				PrimaryKeyOrder: "DESC",
			},
			"aggregateId": csndra.TableColumn{
				Name:            "aggregate_id",
				DataType:        "int",
				PrimaryKeyIndex: "1",
			},
			"yearBucket": csndra.TableColumn{
				Name:            "year_bucket",
				DataType:        "smallint",
				PrimaryKeyIndex: "0",
			},
		}
	}

	return eventMeta
}
