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
				Name:     "aggregateVersion",
				DataType: "bigint",
			},
			"aggregateID": csndra.TableColumn{
				Name:            "aggregateID",
				DataType:        "smallint",
				PrimaryKeyIndex: "1",
			},
			"partitionKey": csndra.TableColumn{
				Name:            "partitionKey",
				DataType:        "smallint",
				PrimaryKeyIndex: "0",
			},
		}
	}

	return eventMeta
}
