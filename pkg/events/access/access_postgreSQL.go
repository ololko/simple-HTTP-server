package access

import(
	"github.com/ololko/simple-HTTP-server/pkg/events/models"

	"database/sql"

	log "github.com/sirupsen/logrus"
	_ "github.com/lib/pq"
)

type PostgreAccess struct {
	Client *sql.DB
}

func (d *PostgreAccess) ReadEvent (request models.RequestT, answer chan<- models.AnswerT, chanErr chan<- error) {

}

func (d *PostgreAccess) WriteEvent (insert models.EventT, chanErr chan<- error) {
	sqlStatement := "INSERT INTO events (type, count, timestamp) VALUES ($1, $2, $3)"
		_, err := d.Client.Exec(sqlStatement, insert.Type, insert.Count, insert.Timestamp)
		if err != nil {
			log.WithFields(log.Fields{
				"type": insert.Type,
			}).Error("Unexpected error while creating new event in database")
			chanErr <- err
			return
		}
		chanErr <- nil
		return
}