package orm

func (o *ORM) UpdateLike(Liked, IdObject int, Type string) error {
	query := "UPDATE Like SET  Liked = ? , CreatedAt =CURRENT_TIMESTAMP  WHERE IdObject = ? AND Type = ?"
	_, err := o.Db.Exec(query, Liked, IdObject, Type)
	if err != nil {
		return err
	}
	return nil
}
