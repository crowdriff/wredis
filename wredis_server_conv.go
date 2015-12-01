package wredis

// SelectAndFlushDb is a convenience method for flushing a specified DB
func (w *Wredis) SelectAndFlushDb(db uint) error {
	if w.safe {
		return unsafeError("SelectAndFlushDb")
	}
	err := w.Select(db)
	if err != nil {
		return err
	}
	return w.FlushDb()
}
