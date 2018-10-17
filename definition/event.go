package definition

import csndra "github.com/TerrexTech/go-cassandrautils/cassandra"

var event map[string]csndra.TableColumn

// Event returns the defintion for the event table.
func Event() map[string]csndra.TableColumn {
	if event == nil {
		event = map[string]csndra.TableColumn{
			"action": csndra.TableColumn{
				Name:            "action",
				DataType:        "text",
				PrimaryKeyIndex: "4",
			},
			"aggregateID": csndra.TableColumn{
				Name:            "aggregate_id",
				DataType:        "smallint",
				PrimaryKeyIndex: "1",
			},
			"correlationID": csndra.TableColumn{
				Name:     "correlation_id",
				DataType: "uuid",
			},
			"data": csndra.TableColumn{
				Name:     "data",
				DataType: "blob",
			},
			"timestamp": csndra.TableColumn{
				Name:     "timestamp",
				DataType: "timestamp",
			},
			"userUUID": csndra.TableColumn{
				Name:     "user_uuid",
				DataType: "uuid",
			},
			"timeUUID": csndra.TableColumn{
				Name:            "time_uuid",
				DataType:        "timeuuid",
				PrimaryKeyIndex: "3",
				PrimaryKeyOrder: "DESC",
			},
			"version": csndra.TableColumn{
				Name:            "version",
				DataType:        "bigint",
				PrimaryKeyIndex: "2",
				PrimaryKeyOrder: "DESC",
			},
			"yearBucket": csndra.TableColumn{
				Name:            "year_bucket",
				DataType:        "smallint",
				PrimaryKeyIndex: "0",
			},
		}
	}

	return event
}
