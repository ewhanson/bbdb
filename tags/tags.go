package tags

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/resolvers"
	"github.com/pocketbase/pocketbase/tools/search"
	"log"
	"net/http"
)

func AddHooks(app *pocketbase.PocketBase) {
	addRoutes(app)
}

func addRoutes(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		_, err := e.Router.AddRoute(echo.Route{
			Method: "GET",
			Path:   "/api/bb/tags/:tagName",
			Handler: func(c echo.Context) error {
				tagName := c.PathParam("tagName")
				tag, err := app.Dao().FindFirstRecordByData("tags", "name", tagName)
				if err != nil {
					return err
				}
				tagId := tag.GetString("id")

				requestData := apis.RequestData(c)

				collection, err := app.Dao().FindCollectionByNameOrId("photos")
				if err != nil {
					return err
				}

				fieldResolver := resolvers.NewRecordFieldResolver(
					app.Dao(),
					collection,
					requestData,
					requestData.Admin != nil,
				)

				searchProvider := search.NewProvider(fieldResolver).
					Query(
						app.Dao().
							RecordQuery(collection).
							LeftJoin("tags", dbx.HashExp{"tags.id": tagId}).
							InnerJoin("json_each(iif(json_valid(tags), tags, null)) je", nil).
							AndWhere(dbx.HashExp{"je.value": tagId}),
					)

				if requestData.Admin == nil && collection.ListRule != nil {
					searchProvider.AddFilter(search.FilterData(*collection.ListRule))
				}

				records := []*models.Record{}

				result, err := searchProvider.ParseAndExec(c.QueryParams().Encode(), &records)
				if err != nil {
					return apis.NewBadRequestError("Invalid filter parameters.", err)
				}

				if err := apis.EnrichRecords(c, app.Dao(), records); err != nil && app.IsDebug() {
					log.Println(err)
				}

				return c.JSON(http.StatusOK, result)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.RequireAdminOrRecordAuth(),
				requireTagsAccess(app),
			},
			Name: "",
		})
		if err != nil {
			return err
		}

		return nil
	})
}

// requireTagsAccess checks if auth user has allowed list permission for tags collection
func requireTagsAccess(app *pocketbase.PocketBase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestData := apis.RequestData(c)
			tagName := c.PathParam("tagName")

			collection, err := app.Dao().FindCollectionByNameOrId("tags")
			if err != nil {
				return err
			}
			listRule := collection.ListRule

			ruleFunc := func(q *dbx.SelectQuery) error {
				if *listRule == "" {
					return nil // empty public rule
				}

				resolver := resolvers.NewRecordFieldResolver(app.Dao(), collection, requestData, true)
				expr, err := search.FilterData(*listRule).BuildExpr(resolver)
				if err != nil {
					return err
				}
				err = resolver.UpdateQuery(q)
				if err != nil {
					return err
				}
				q.AndWhere(expr)

				return nil
			}

			tag, err := app.Dao().FindFirstRecordByData(collection.Id, "name", tagName)
			if err != nil {
				return err
			}

			foundRecord, err := app.Dao().FindRecordById(collection.Id, tag.Id, ruleFunc)
			if err == nil && foundRecord != nil {
				return next(c)
			}

			return apis.NewForbiddenError("The authorized record model is not allowed to perform this action.", nil)
		}
	}
}
