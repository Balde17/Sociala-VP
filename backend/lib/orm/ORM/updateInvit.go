package orm

func (o *ORM) UpdateInvite(IdUser, IdUserInvited, IdGroup int, Accepted string) error {
	query := "UPDATE Invite SET  Accepted = ? WHERE  IdUserInvited = ? AND IdGroup = ?"
	_, err := o.Db.Exec(query, Accepted, IdUserInvited, IdGroup)
	if err != nil {
		return err
	}
	return nil
}
