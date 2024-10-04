package orm

func (o *ORM) UpdateStatutRequet(IdFollowee, IdFollower int, Status string) error {
	query := "UPDATE Followers SET  Status = ? WHERE IdFollowee = ? AND IdFollower = ?"
	_, err := o.Db.Exec(query, Status, IdFollowee, IdFollower)
	if err != nil {
		return err
	}
	return nil
}
