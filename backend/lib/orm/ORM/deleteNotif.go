package orm

func (o *ORM) DeleteNotif(idNotif int) error {
	query := "DELETE FROM Notifications WHERE Id = ?"
	_, err := o.Db.Exec(query, idNotif)
	if err != nil {
		return err
	}
	return nil
}
