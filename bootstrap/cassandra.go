package bootstrap

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/TerrexTech/go-eventstore-models/definition"

	csndra "github.com/TerrexTech/go-cassandrautils/cassandra"
	"github.com/TerrexTech/go-cassandrautils/cassandra/driver"
	"github.com/TerrexTech/go-commonutils/commonutil"
	cql "github.com/gocql/gocql"
	"github.com/pkg/errors"
)

// cassandra allows ceonvniently connecting to cassandra
// and creates the required Keyspace and Table.
type cassandra struct {
	DataCenters []string
	Hosts       []string
	Keyspace    string
	Password    string
	Table       string
	Username    string
	TableDef    map[string]csndra.TableColumn
	// Currently used for closing session (in tests, for example).
	session *driver.Session
}

// evenStoreSession creates a Cassandra session, used for storing events.
func (ca *cassandra) newSession(
	clusterHosts *[]string,
	username string,
	password string,
) (*driver.Session, error) {
	cluster := cql.NewCluster(*clusterHosts...)
	cluster.ConnectTimeout = time.Millisecond * 3000
	cluster.Timeout = time.Millisecond * 3000
	cluster.ProtoVersion = 4
	cluster.RetryPolicy = &cql.ExponentialBackoffRetryPolicy{
		NumRetries: 5,
	}

	if username != "" && password != "" {
		cluster.Authenticator = cql.PasswordAuthenticator{
			Username: username,
			Password: password,
		}
	}
	return csndra.GetSession(cluster)
}

// createKeyspaceTable creates the required Keyspace and Table
// as specified in CassandraAdapter.
func (ca *cassandra) createKeyspaceTable(
	session driver.SessionI,
	keyspaceName string,
	tableName string,
	datacenters *[]string,
) (*csndra.Table, error) {
	datacenterMap := map[string]int{}
	for _, dcStr := range *datacenters {
		dc := strings.Split(dcStr, ":")
		centerID, err := strconv.Atoi(dc[1])
		if err != nil {
			return nil, errors.Wrap(
				err,
				"Cassandra Keyspace Create Error (CASSANDRA_DATA_CENTERS format mismatch)"+
					"CASSANDRA_DATA_CENTERS must be of format \"<ID>:<replication_factor>\"",
			)
		}
		datacenterMap[dc[0]] = centerID
	}

	keyspaceConfig := csndra.KeyspaceConfig{
		Name:                    keyspaceName,
		ReplicationStrategy:     "NetworkTopologyStrategy",
		ReplicationStrategyArgs: datacenterMap,
	}

	keyspace, err := csndra.NewKeyspace(session, keyspaceConfig)
	if err != nil {
		return nil, errors.Wrap(err, "Error Creating Keyspace")
	}

	tc := &csndra.TableConfig{
		Keyspace: keyspace,
		Name:     tableName,
	}

	t, err := csndra.NewTable(session, tc, &ca.TableDef)
	if err != nil {
		return nil, errors.Wrap(err, "Error Creating Table")
	}
	return t, nil
}

// ensureKeyspaceTable creates the specified Keyspace and Table
// if they don't exist, and returns a go-cassandrautils table.
// If the keyspace/table exist, this just returns the go-cassandrautils table.
func (ca *cassandra) ensureKeyspaceTable() (*csndra.Table, error) {
	if ca.Hosts == nil || len(ca.Hosts) == 0 {
		return nil, errors.New(
			"Error connecting to Cassandra: Empty or Nil Hosts provided",
		)
	}
	log.Println("Initializing CassandraIO")

	// Create Cassandra session
	session, err := ca.newSession(&ca.Hosts, ca.Username, ca.Password)
	if err != nil {
		err = errors.Wrap(err, "Cassandra Session Creation Error")
		return nil, err
	}
	ca.session = session
	log.Println("Created Cassandra Session")

	// Create Keyspace and Table
	t, err := ca.createKeyspaceTable(
		ca.session,
		ca.Keyspace,
		ca.Table,
		&ca.DataCenters,
	)
	// session.Close()
	if err != nil {
		err = errors.Wrap(err, "Keyspace Creation Error")
		return nil, err
	}

	log.Println("Cassandra Table Bootstrapped")
	return t, nil
}

// initCassandra verifies required env variables, and then
// creates the Cassandra table as per set env-vars.
func initCassandra(
	tableDef map[string]csndra.TableColumn,
	tableName string,
) (*csndra.Table, error) {
	envVars := []string{
		"CASSANDRA_HOSTS",
		"CASSANDRA_DATA_CENTERS",
		"CASSANDRA_USERNAME",
		"CASSANDRA_PASSWORD",
		"CASSANDRA_KEYSPACE",
		"CASSANDRA_EVENT_TABLE",
		"CASSANDRA_EVENT_META_TABLE",
	}

	for _, varname := range envVars {
		envVar := os.Getenv(varname)
		if envVar == "" {
			err := errors.New(
				"Error while bootstrapping Cassandra table: " +
					"Following env-var is required but was not found: " +
					varname,
			)
			return nil, err
		}
	}

	hosts := os.Getenv("CASSANDRA_HOSTS")
	dataCenters := os.Getenv("CASSANDRA_DATA_CENTERS")
	username := os.Getenv("CASSANDRA_USERNAME")
	password := os.Getenv("CASSANDRA_PASSWORD")
	keyspaceName := os.Getenv("CASSANDRA_KEYSPACE")

	c := cassandra{
		DataCenters: *commonutil.ParseHosts(dataCenters),
		Hosts:       *commonutil.ParseHosts(hosts),
		Keyspace:    keyspaceName,
		Password:    password,
		Table:       tableName,
		Username:    username,
		TableDef:    tableDef,
	}

	return c.ensureKeyspaceTable()
}

// Event creates the event table if it doesn't exist.
// This also returns the go-cassandrautils table reference to the table.
// Following env vars can be used for configuration:
//  CASSANDRA_HOSTS
//  CASSANDRA_DATA_CENTERS
//  CASSANDRA_USERNAME
//  CASSANDRA_PASSWORD
//  CASSANDRA_KEYSPACE
//  CASSANDRA_EVENT_TABLE
func Event() (*csndra.Table, error) {
	tableName := os.Getenv("CASSANDRA_EVENT_TABLE")
	return initCassandra(definition.Event(), tableName)
}

// EventMeta creates the event-meta table if it doesn't exist.
// This also returns the go-cassandrautils table reference to the table.
// Following env vars can be used for configuration:
//  CASSANDRA_HOSTS
//  CASSANDRA_DATA_CENTERS
//  CASSANDRA_USERNAME
//  CASSANDRA_PASSWORD
//  CASSANDRA_KEYSPACE
//  CASSANDRA_EVENT_META_TABLE
func EventMeta() (*csndra.Table, error) {
	tableName := os.Getenv("CASSANDRA_EVENT_META_TABLE")
	return initCassandra(definition.EventMeta(), tableName)
}
