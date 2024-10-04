package orm

func (o *ORM) UpdateStatut(Profil string, Id int) error {
	query := "UPDATE User SET  Profil = ? WHERE Id = ?"
	_, err := o.Db.Exec(query, Profil, Id)
	if err != nil {
		return err
	}
	return nil
}
