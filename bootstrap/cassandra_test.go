package bootstrap

import (
	"log"
	"os"

	csndra "github.com/TerrexTech/go-cassandrautils/cassandra"
	"github.com/TerrexTech/go-cassandrautils/cassandra/driver"

	"github.com/TerrexTech/go-commonutils/commonutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Event", func() {
	var (
		keyspace string
		session  *driver.Session
		// We'll replace the Keyspace-name (to use for testing) in env
		// var before starting tests, and restore it after tests complete.
		oldKeyspace string
	)

	// Create a session to use in all tests
	BeforeSuite(func() {
		oldKeyspace = os.Getenv("CASSANDRA_KEYSPACE")
		testKeyspace := os.Getenv("CASSANDRA_TEST_KEYSPACE")
		os.Setenv("CASSANDRA_KEYSPACE", testKeyspace)
		log.Println("Switched \"CASSANDRA_KEYSPACE\" to: " + testKeyspace)
	})

	BeforeEach(func() {
		hosts := os.Getenv("CASSANDRA_HOSTS")
		username := os.Getenv("CASSANDRA_USERNAME")
		password := os.Getenv("CASSANDRA_PASSWORD")
		keyspace = os.Getenv("CASSANDRA_TEST_KEYSPACE")

		// Create Cassandra Session
		c := cassandra{}
		var err error
		session, err = c.newSession(commonutil.ParseHosts(hosts), username, password)
		Expect(err).ToNot(HaveOccurred())

		// Start fresh for every test
		q := session.GoCqlSession().Query("DROP KEYSPACE IF EXISTS " + keyspace)
		err = q.Exec()
		Expect(err).ToNot(HaveOccurred())
		q.Release()

		// Create Keyspace
		keyspaceConfig := csndra.KeyspaceConfig{
			Name:                keyspace,
			ReplicationStrategy: "NetworkTopologyStrategy",
			ReplicationStrategyArgs: map[string]int{
				"datacenter1": 1,
			},
		}
		_, err = csndra.NewKeyspace(session, keyspaceConfig)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		session.Close()
		log.Println("Closed Cassandra Session")
	})

	AfterSuite(func() {
		if oldKeyspace == "" {
			os.Unsetenv("CASSANDRA_KEYSPACE")
		} else {
			os.Setenv("CASSANDRA_KEYSPACE", oldKeyspace)
			log.Println("Switched \"CASSANDRA_KEYSPACE\" to: " + oldKeyspace)
		}
	})

	It("should create the event table", func() {
		_, err := Event()
		Expect(err).ToNot(HaveOccurred())

		keyspaceMeta, err := session.GoCqlSession().KeyspaceMetadata(keyspace)

		eventTable := os.Getenv("CASSANDRA_EVENT_TABLE")
		Expect(keyspaceMeta.Tables[eventTable]).ToNot(BeNil())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should create the event-meta table", func() {
		_, err := EventMeta()
		Expect(err).ToNot(HaveOccurred())

		keyspaceMeta, err := session.GoCqlSession().KeyspaceMetadata(keyspace)

		eventMetaTable := os.Getenv("CASSANDRA_EVENT_META_TABLE")
		Expect(keyspaceMeta.Tables[eventMetaTable]).ToNot(BeNil())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should return error if a critical env-var is not set", func() {
		envVars := []string{
			"CASSANDRA_HOSTS",
			"CASSANDRA_DATA_CENTERS",
			"CASSANDRA_KEYSPACE",
			"CASSANDRA_EVENT_TABLE",
			"CASSANDRA_EVENT_META_TABLE",
		}

		for _, v := range envVars {
			envVal := os.Getenv(v)
			err := os.Unsetenv(v)
			Expect(err).ToNot(HaveOccurred())

			if v != "CASSANDRA_EVENT_META_TABLE" {
				_, err = Event()
				Expect(err).To(HaveOccurred())
			}
			if v != "CASSANDRA_EVENT_TABLE" {
				_, err = EventMeta()
				Expect(err).To(HaveOccurred())
			}

			err = os.Setenv(v, envVal)
			Expect(err).ToNot(HaveOccurred())
		}
	})
})
