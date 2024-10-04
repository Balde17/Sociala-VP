package orm

func (o *ORM) DeleteRequest(IdFollowee, IdFollower int) error {
	query := "DELETE FROM Followers WHERE IdFollowee = ? AND IdFollower = ?"
	_, err := o.Db.Exec(query, IdFollowee, IdFollower)
	if err != nil {
		return err
	}
	return nil
}
