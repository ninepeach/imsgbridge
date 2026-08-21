package applemsg

type CursorStore interface {
	Load() (int64, error)

	Save(int64) error
}
