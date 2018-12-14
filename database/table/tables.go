package table

import (
	"github.com/dibaggioj/sensor-api/database"
	"github.com/pkg/errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/dibaggioj/sensor-api/models"
	"fmt"
)

const SQL_TABLE_CREATION_DATA_SET = `
CREATE TABLE IF NOT EXISTS dataset (
	id SERIAL PRIMARY KEY,
	timestamp timestamp with time zone DEFAULT current_timestamp
	data INTEGER REFERENCES measurements(id)
)`

const SQL_TABLE_CREATION_MEASUREMENTS = `
CREATE TABLE IF NOT EXISTS measurements (
	id SERIAL PRIMARY KEY,
	temperature FLOAT,
	humidity FLOAT
)`

// DataTable holds the internals of the table, i.e,
// the manager of this instance's database pool (Connection).
// Here you could also add things like a `logger` with
// some predefined fields (for structured logging with
// context).
type Table struct {
	Connection *database.Connection
	Name string
}

// DataTableConfig holds the configuration passed to
// the DataTable "constructor" (`NewDataTable`).
type DataTableConfig struct {
	Connection *database.Connection
	Name string
	CreateSql string
}


// NewDataTable creates an instance of DataTable.
// It performs all of its operation against a pool of connections that is managed by `Connection`.
func NewDataTable(cfg DataTableConfig) (table Table, err error) {
	if cfg.Connection == nil {
		err = errors.New(
			"Can't create table without Connection instance")
		return
	}

	table.Connection = cfg.Connection
	table.Name = cfg.Name
	// Always try to create the table just in case we don't create them at the database startup.
	// This won't fail in case the table already exists.
	if err = table.createTable(cfg.CreateSql); err != nil {
		err = errors.Wrapf(err,
			"Couldn't create table during initialization")
		return
	}

	return
}

// createTable tries to create a table. If it already exists or not, no error is thrown.
// The operation only fails in case there's a mismatch in table definition of if there's a connection error.
func (table *Table) createTable(qry string) (err error) {
	// Exec executes a query without returning any rows.
	if _, err = table.Connection.Db.Exec(qry); err != nil {
		err = errors.Wrapf(err,
			"Data table creation query failed (%s)",
			qry)
		return
	}
	return
}

func (table *Table) InsertDataPoint(row models.DataPoint) (newRow models.DataPoint, err error) {
	// TODO: data validation

	if !row.IsValid() {
		fmt.Println("Data point is missing data, unable to add to table.");
		return
	}

	qryDataPoint := `
INSERT INTO dataset (data)
VALUES ($1)
RETURNING id`

	qryData := `
INSERT INTO measurements (temperature, humidity)
VALUES ($1, $2)
RETURNING id, temperature, humidity`

	sensorDataPtr := models.SensorData{}
	newRow.Data = &sensorDataPtr
	err = table.Connection.Db.QueryRow(qryData, row.Data.Temperature, row.Data.Humidity).Scan(&newRow.Data.ID,
		&newRow.Data.Temperature, &newRow.Data.Humidity)
	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't insert row into DB (%s)",
			spew.Sdump(row))
		return
	}

	fmt.Printf("## newRow.Data: %d, %d, %d", newRow.Data.ID, newRow.Data.Temperature, newRow.Data.Humidity)

	err = table.Connection.Db.QueryRow(qryDataPoint, newRow.Data.ID).Scan(&newRow.ID)
	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't insert row into DB (%s)",
			spew.Sdump(row))
		return
	}

	fmt.Printf("## &newRow.ID: %d", newRow.ID)

	return
}

func (table *Table) InsertSensorData(row models.SensorData) (newRow models.SensorData, err error) {
	// TODO: data validation
	qry := `
INSERT INTO measurements (temperature, humidity)
VALUES ($1, $2)
RETURNING id`

	// `QueryRow` is a single-row query that, unlike `Query()`, doesn't hold a connection. Errors from `QueryRow` are
	// forwarded to `Scan` where we can get errors from both. Here we perform such query for inserting because we want
	// to grab right from the Database the entry that was inserted (plus the fields that the database generated). If we
	// were just getting a value, we could also check if the query was successful but returned 0 rows with
	// `if err == sql.ErrNoRows`. result, err := table.Connection.Db.Exec(qry, row.Temperature, row.Humidity)
	err = table.Connection.Db.QueryRow(qry, row.Temperature, row.Humidity).Scan(&newRow.ID)
	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't insert row into DB (%s)",
			spew.Sdump(row))
		return
	}

	return
}