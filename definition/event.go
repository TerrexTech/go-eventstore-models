package definition

import csndra "github.com/TerrexTech/go-cassandrautils/cassandra"

var event map[string]csndra.TableColumn

// Event returns the defintion for the event table.
func Event() map[string]csndra.TableColumn {
	if event == nil {
		event = map[string]csndra.TableColumn{
			"aggregateID": csndra.TableColumn{
				Name:            "aggregate_id",
				DataType:        "smallint",
				PrimaryKeyIndex: "1",
			},
			"eventAction": csndra.TableColumn{
				Name:     "event_action",
				DataType: "text",
			},
			"correlationID": csndra.TableColumn{
				Name:     "correlation_id",
				DataType: "uuid",
			},
			"serviceAction": csndra.TableColumn{
				Name:     "service_action",
				DataType: "text",
			},
			"data": csndra.TableColumn{
				Name:     "data",
				DataType: "blob",
			},
			"nanoTime": csndra.TableColumn{
				Name:            "nano_time",
				DataType:        "timestamp",
				PrimaryKeyIndex: "3",
				PrimaryKeyOrder: "DESC",
			},
			"userUUID": csndra.TableColumn{
				Name:     "user_uuid",
				DataType: "uuid",
			},
			"UUID": csndra.TableColumn{
				Name:     "uuid",
				DataType: "uuid",
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
