// The model package contains key data structures and interfaces.
//
// The subpackage designlanguage is generated and shoudl not be modified.
// The subpackage dl is provided for manually written code relating to the design language.
// TODO Maybe find a way for the generated code to be in a differently named package.
package model

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/mattmunz/sundries/collections"
	"github.com/pkg/errors"
)

const (
	// TODO Hardcoded paths
	SundriesDataDirPath = "/Users/mattmunz/Dropbox/.sundries"
	// Encrypted JSON
	SundriesDataFilePath = "/Users/mattmunz/.sundries/data.jsone"
	SundriesDBPath       = "/Users/mattmunz/.sundries/db"

	// TODO Make this a flag
	KeychainAccountName = "mattmunz"
	BaseServicePath     = "github.com/mattmunz/sundries"

	// TODO Want to use different keys for different data backends
	// Secrets
	EncryptionKeyName = "encryptedDataKey"
)

type Identifiable interface {
	ID() string
}

type IdentifiableModel struct {
	ID_ string `json:"id"`
}

func NewIdentifiableModel(id string) *IdentifiableModel {
	return &IdentifiableModel{id}
}

func (m *IdentifiableModel) ID() string {
	return m.ID_
}

func (m *IdentifiableModel) String() string {
	return m.ID()
}

func GetIdentifiable(id string, items ...Identifiable) Identifiable {
	for _, item := range items {
		if item.ID() == id {
			item2 := item
			return item2
		}
	}
	return nil
}

// TODO Eventually sequence num can be replaced with a UUID
// This would require some way to map long IDs to short numbers in the UI
func NewID(reservedIDs *collections.StringSet) string {
	newID := ""
	for newIDNum := 1; true; newIDNum = newIDNum + 1 {
		newID = fmt.Sprintf("l%d", newIDNum)
		if !reservedIDs.Contains(newID) {
			break
		}
	}
	return newID
}

func NewUUID() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

type Time struct {
	time.Time
}

func NewTime(aTime time.Time) Time {
	return Time{aTime.UTC()}
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format(time.RFC3339Nano))), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	jsonText := string(data)
	if len(jsonText) < 3 {
		return errors.Errorf("Bad time format: %s", jsonText)
	}
	newTime, err := time.Parse(time.RFC3339Nano, jsonText[1:len(jsonText)-1])
	if err != nil {
		return errors.Wrap(err, "Couldn't parse time text from JSON")
	}

	t.Time = newTime
	return nil
}

func (t *Time) AsTime() time.Time {
	return t.Time
}

// ACK, needs to return pointer
func Now() Time { return NewTime(time.Now()) }

type TimeOfDay string

const (
	Morning, Afternoon, Evening TimeOfDay = "Morning", "Afternoon", "Evening"
)

type AppointmentTime struct {
	DayOfWeek time.Weekday `json:"dayOfWeek"`
	TimeOfDay TimeOfDay    `json:"timeOfDay"`
}

func NewAppointmentTime(dayOfWeek time.Weekday, timeOfDay TimeOfDay) AppointmentTime {
	return AppointmentTime{dayOfWeek, timeOfDay}
}

func ParseAppointmentTime(text string) (AppointmentTime, error) {
	trimmedText := strings.TrimSpace(text)
	if len(trimmedText) < 1 {
		return AppointmentTime{}, errors.New("Not enough chars")
	}

	if len(trimmedText) > 3 {
		return AppointmentTime{}, errors.New("Too many chars")
	}

	dayText := ""
	todText := ""
	for i, char := range trimmedText {
		if len(todText) > 0 {
			return AppointmentTime{}, errors.Errorf("Too many TOD chars: %s", text)
		}

		if i == 0 {
			if unicode.IsLower(char) {
				return AppointmentTime{}, errors.New("First char is lower")
			}

			dayText = string(char)
			continue
		}

		if i == 1 && unicode.IsLower(char) {
			dayText = dayText + string(char)
			continue
		}

		todText = string(char)
	}

	dayOfWeek, err := newWeekday(dayText)
	if err != nil {
		return AppointmentTime{}, errors.New("Bad weekday")
	}

	tod := Morning
	if len(todText) > 0 {
		tod, err = newTimeOfDay(todText)
		if err != nil {
			return AppointmentTime{}, errors.Wrapf(err, "Couldn't parse tod from: %s", text)
		}
	}

	return NewAppointmentTime(dayOfWeek, tod), nil
}

func newWeekday(dayText string) (time.Weekday, error) {
	switch dayText {
	case "M":
		return time.Monday, nil
	case "Tu":
		return time.Tuesday, nil
	case "W":
		return time.Wednesday, nil
	case "Th":
		return time.Thursday, nil
	case "F":
		return time.Friday, nil
	case "Sa":
		return time.Saturday, nil
	case "Su":
		return time.Sunday, nil
	}

	return 0, fmt.Errorf("Unknown day: %s", dayText)
}

func newTimeOfDay(text string) (TimeOfDay, error) {
	switch text {
	case "M":
		return Morning, nil
	case "A":
		return Afternoon, nil
	case "E":
		return Evening, nil
	}

	return "", fmt.Errorf("Unknown time of day: %s", text)
}
