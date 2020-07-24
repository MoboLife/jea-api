package database

// ConnectionInfo database connection info
type ConnectionInfo struct {
	Host     		string		`json:"host"`
	Port     		string		`json:"port"`
	User     		string		`json:"user"`
	Database 		string		`json:"database"`
	Password 		string		`json:"password"`
	Driver			string		`json:"driver"`
	URL				string		`json:"url"`
}

