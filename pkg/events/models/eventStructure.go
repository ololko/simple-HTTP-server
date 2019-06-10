//Data structure as it is stored in firestore
package models

import "github.com/jinzhu/gorm"

type EventT struct {
	gorm.Model
	Count     int64
	Type      string
	Timestamp int64
}
