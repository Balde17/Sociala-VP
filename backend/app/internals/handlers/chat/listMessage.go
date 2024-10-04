package chat

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"net/http"
)

func ListMessage(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	me := r.URL.Query().Get("me")
	him := r.URL.Query().Get("him")
	offset := r.URL.Query().Get("offset")

	if me != "" && him != "" && offset != "" {
		rslt, err := queryBuilder.NewSelectBuilder().
			Select(
				"Content",
				"Sender",
				"Receiver",
				"c.CreatedAt AS MessageDate").
			From("chat c").
			Join("INNER", "User u", "u.Id= c.Receiver").
			Join("INNER", "User v", "v.Id= c.Sender").
			Where("(u.Id = " + me + " AND v.Id = " + him + ") OR (u.Id = " + him + " AND v.Id = " + me + ")").
			And("Type='CHAT'").
			OrderBy("MessageDate DESC").
			Limit("10").
			Offset(offset).
			SelectQuery(models.OrmInstance.Db)

		var posts []map[string]string

		if err == nil {
			for _, row := range rslt {
				posts = append(posts, map[string]string{
					"Content":     row[0],
					"Sender":      row[1],
					"Receiver":    row[2],
					"MessageDate": row[3],
				})
			}
		}
		data.Content = posts
		utils.SendJSON(w, http.StatusOK, data)
	} else {
		errorMsg := map[string]string{
			"error": "cannot get messages",
		}
		utils.SendJSON(w, http.StatusBadRequest, errorMsg)
	}
}

func ListMessageGroup(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	him := r.URL.Query().Get("him")
	offset := r.URL.Query().Get("offset")

	if him != "" && offset != "" {
		rslt, err := queryBuilder.NewSelectBuilder().
			Select(
				"CONCAT(u.FirstName,' ',u.LastName) AS Name",
				"Content",
				"Sender",
				"Receiver",
				"c.CreatedAt AS MessageDate").
			From("chat c").
			Join("INNER", "User u", "u.Id = c.Sender").
			Where("(c.Receiver = " + him + ")").
			And("Type='Groups'").
			OrderBy("MessageDate DESC").
			Limit("10").
			Offset(offset).
			SelectQuery(models.OrmInstance.Db)

		var posts []map[string]string

		if err == nil {
			for _, row := range rslt {
				posts = append(posts, map[string]string{
					"Name":        row[0],
					"Content":     row[1],
					"Sender":      row[2],
					"Receiver":    row[3],
					"MessageDate": row[4],
				})
			}
		}
		data.Content = posts
		utils.SendJSON(w, http.StatusOK, data)
	} else {
		errorMsg := map[string]string{
			"error": "cannot get messages",
		}
		utils.SendJSON(w, http.StatusBadRequest, errorMsg)
	}
}
