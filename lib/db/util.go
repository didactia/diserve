package db

import (
  "log"
  "fmt"
  "io/ioutil"
  "os"
  "context"
  "errors"
  "strconv"

  "github.com/dgraph-io/dgraph/client"
  "github.com/dgraph-io/dgraph/protos"

  "github.com/gogo/protobuf/proto"

  "diserve.didactia.org/lib/env"
  "diserve.didactia.org/lib/templater"

  "google.golang.org/grpc"
)
// ErrResponseUnmarshalling occurs on response unmarshalling error,
// datasensitive version.
var ErrResponseUnmarshalling = errors.New("database error: unmarshalling")

// ErrResponseQuery occurs on response query error, datasensitive version
var ErrResponseQuery = errors.New("database error: query")

// DatabaseClient holds the dgraph client, and a function to gracefully close
type DatabaseClient struct {
  Dgraph *client.Dgraph
  Templater *templater.Templater
  Close func()
}

var dbc *DatabaseClient

// NewDatabaseClient returns a dgraph client and the function to close it.
func NewDatabaseClient(ip string, port string) (*DatabaseClient) {
  conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, port), grpc.WithInsecure())
  if err != nil {
    log.Fatal(err)
  }

  clientDir, err := ioutil.TempDir("", "client_")
  if err != nil {
    log.Fatal(err)
  }
  dgraph := client.NewDgraphClient([]*grpc.ClientConn{conn}, client.DefaultOptions, clientDir)
  closer := func() {
    dgraph.Close()
    os.RemoveAll(clientDir)
    conn.Close()
  }
  templater := templater.NewTemplater(env.Vars.DBTMPLPATH)
  dbclient := &DatabaseClient{
    Dgraph: dgraph,
    Templater: templater,
    Close: closer,
  }
  return dbclient
}

// InitializeDatabaseClient intializes the database client, this must be
// called before running other database functions.
func InitializeDatabaseClient(ip string, port string) {
  dbc = NewDatabaseClient(ip, port)
}

// Close the database client gracefully.
func Close() {
  dbc.Close()
}

// AddSchema adds the given schema through the databaseclient
func AddSchema(schema string) {
  req := client.Req{}
  req.SetSchema(schema)
  resp, err := dbc.Dgraph.Run(context.Background(), &req)
  if err != nil {
    log.Fatalf("Error in getting response from server, %s", err)
  }
  fmt.Printf("Response %+v\n", proto.MarshalTextString(resp))
}

// Query queries with the query string through the database client
func Query(name string, data interface{}) (*protos.Response, error) {
  req := client.Req{}
  query, err := dbc.Templater.RenderString(name, data)
  fmt.Println(query)
  return nil, ErrResponseQuery
  if err != nil {
    return nil, err
  }
  req.SetQuery(query)
  resp, err := dbc.Dgraph.Run(context.Background(), &req)
  if err != nil {
    return nil, err
  }
  return resp, nil
}

func uidString(i uint64) string {
  return "0x" + strconv.FormatUint(i, 16)
}
