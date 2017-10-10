package db

import (
  "log"
  "fmt"
  "io/ioutil"
  "os"
  "context"

  "github.com/dgraph-io/dgraph/client"

  "github.com/gogo/protobuf/proto"

  "diserve.didactia.org/lib/env"

  "google.golang.org/grpc"
)

// Init initializes the database, with the schema defined at TODO
func Init(args []string) {
  d := newClient()
  defer d.Close()
  // user
  d.addSchema(`name: string @index(exact) .
               password: password .
               title: string @index(exact) .
               prerequisite: uid .
               concept: uid .
               understander: uid @count .
               text: string .
               reasoning: uid .
               comment: uid .
               old: uid .
               next: uid .
               rating: uid @count .
               expression: uid .
               response: uid .`)
}

// newClient returns a dgraph client.
func newClient() (*client.Dgraph) {
  conn, err := grpc.Dial(fmt.Sprintf("%s:%s", env.Vars.DBIP, env.Vars.DBPORT), grpc.WithInsecure())
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()

  clientDir, err := ioutil.TempDir("", "client_")
  if err != nil {
    log.Fatal(err)
  }
  defer os.RemoveAll(clientDir)
  return client.NewDgraphClient([]*grpc.ClientConn{conn}, client.DefaultOptions, clientDir)
}

func (d *client.Dgraph) addSchema(schema string) {
  req := client.Req{}
  req.SetSchema(schema)
  resp, err := d.Run(context.Background(), &req)
  if err != nil {
    log.Fatalf("Error in getting response from server, %s", err)
  }
  fmt.Printf("Response %+v\n", proto.MarshalTextString(resp))
}
