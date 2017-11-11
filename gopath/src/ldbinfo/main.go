package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	//"github.com/golang/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"encoding/hex"
	"github.com/fatih/color"
)

var cred    = color.New(color.FgRed)
var cgreen  = color.New(color.FgGreen)
var cyellow = color.New(color.FgYellow)
var cblue   = color.New(color.FgBlue)

func iterateKeys(db *leveldb.DB) {
	cyellow.Printf("##############################################################################\n") 
	cyellow.Printf("######            Keys and Values defined in the database               ######\n") 
	cyellow.Printf("##############################################################################\n") 
	// itarate over all keys
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		cblue.Printf("KEY   : "); fmt.Printf("length=%d\n",len(key))
		fmt.Print(hex.Dump(key))
		cblue.Printf("VALUE : "); fmt.Printf("length=%d\n",len(value))
		fmt.Print(hex.Dump(value))
		fmt.Println()
	}
	iter.Release()
	_ = iter.Error()
	
}


func dbProperties(db *leveldb.DB, dbdir string) {
	cyellow.Printf("##############################################################################\n") 
	cyellow.Printf("######                    levelDB database properties                   ######\n") 
	cyellow.Printf("##############################################################################\n") 
	cblue.Printf("LOCATION   : "); fmt.Printf("%s\n",dbdir)
	cblue.Printf("FILES      : ")
	files, err := ioutil.ReadDir(dbdir)
	if err != nil {
		log.Fatal(err)
	}
	i := 1;
	for _, f := range files {
		if i == 1 {
			fmt.Printf("#%03d %s\n", i, f.Name())
		} else {
			fmt.Printf("             #%03d %s\n", i, f.Name())
		}
		i++
	}
	stats, _ := db.GetProperty("leveldb.stats")
	cblue.Printf("PROPERTIES : "); cyellow.Printf("leveldb.stats        "); fmt.Printf(": \n%s", stats)
	sstables, _ := db.GetProperty("leveldb.sstables")
	cblue.Printf("             "); cyellow.Printf("leveldb.sstables     "); fmt.Printf(": \n%s", sstables)
	blockpool, _ := db.GetProperty("leveldb.blockpool")
	cblue.Printf("             "); cyellow.Printf("leveldb.blockpool    "); fmt.Printf(": %s\n", blockpool)
	cachedblock, _ := db.GetProperty("leveldb.cachedblock")
	cblue.Printf("             "); cyellow.Printf("leveldb.cachedblock  "); fmt.Printf(": %s\n", cachedblock)
	openedtables, _ := db.GetProperty("leveldb.openedtables")
	cblue.Printf("             "); cyellow.Printf("leveldb.openedtables "); fmt.Printf(": %s\n", openedtables)
	alivesnaps, _ := db.GetProperty("leveldb.alivesnaps")
	cblue.Printf("             "); cyellow.Printf("leveldb.alivesnaps   "); fmt.Printf(": %s\n", alivesnaps)
	aliveiters, _ := db.GetProperty("leveldb.aliveiters")
	cblue.Printf("             "); cyellow.Printf("leveldb.aliveiters   "); fmt.Printf(": %s\n", aliveiters)
	

}

func main() {

	dbdir := "production/ledgersData/stateLeveldb"
	//dbdir := "production/ledgersData/historyLeveldb"
	//dbdir := "production/ledgersData/chains/index"
	//dbdir := "production/ledgersData/ledgerProvider"

	cgreen.Printf("LevelDB location : "); fmt.Printf("%s\n",dbdir)

	dbopt := &opt.Options { ErrorIfMissing: true }
	db, err := leveldb.OpenFile(dbdir, dbopt)
	if err != nil {
		cred.Printf("ERROR ! "); fmt.Printf("Database %s does not exist or is corrupted\n", dbdir)
                cyellow.Printf("Details : "); log.Fatal(err)
		os.Exit(1)	
	} 
	defer db.Close()

	iterateKeys(db)
	dbProperties(db, dbdir)
}
