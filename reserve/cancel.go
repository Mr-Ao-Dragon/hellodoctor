package reserve

import (
	"log"

	"github.com/Mr-Ao-Dragon/hellodoctor/database"
)

func CancelReserve(reserveID int32) (err error) {
	err = database.CancelReserve(reserveID)
	if err != nil {
		log.Printf("Cancel reserve error: %#v\n", err)
		return err
	}
	return nil
}
