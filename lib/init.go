package lyndbdump

import (
	"fmt"
	"strings"
)

// dump begins dumping process
func dump(conf Config) {
	confHosts := conf.Hosts
	confLocalDb := conf.LocalDb
	for i := 0; i < len(confHosts); i++ {
		confHost := confHosts[i]
		if confHost.Enabled {
			mysql := &MySQL{
				Host:     confHost.Db.Host,
				Port:     confHost.Db.Port,
				DB:       confHost.Db.Name,
				User:     confHost.Db.User,
				Password: confHost.Db.Pass,
			}

			mysqlDump := mysql.Dump()

			if confHost.WriteToFile {
				writeToFile(confHost.Name+".sql", convertDumpToLocal(mysqlDump, confHost))
				fmt.Println("Storing " + confHost.Db.Name + " database to './tmp/" + confHost.Name + ".sql'")
			} else {
				//Import Directly to Database
				mysql.Import(string(convertDumpToLocal(mysqlDump, confHost)), confLocalDb)
			}

		}
	}
	fmt.Println("Process Successfully Compelete.")

}

// convertDumpToLocal replace remote host name with local hostname
func convertDumpToLocal(mysqlDump []byte, host Hosts) []byte {
	return []byte(strings.Replace(string(mysqlDump), host.Protocol+host.Name, "http://"+host.LocalName, -1))
}

// Init initialize the application
func Init() {
	conf := ConfParse()
	dump(conf)
}
