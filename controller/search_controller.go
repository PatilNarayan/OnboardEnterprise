package controller

import (
	"core/generated"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	_ generated.SearchAPIAPI = &searchApiController{}
)

type searchApiController struct{}

// NewSearchApiController returns a new instance of the controller
func NewSearchApiController() *searchApiController {
	return &searchApiController{}
}

func (a *searchApiController) PagesSearchGet(c *gin.Context) {
	// 1️⃣ Read query parameters
	orgName := c.Query("org_name")
	teamName := c.Query("team_name") // optional
	userName := c.Query("user_name") // or user_id if you prefer
	query := c.Query("query")

	if orgName == "" || userName == "" || query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "org_name, user_name, and query are required",
		})
		return
	}

	// ✅ At this point you have:
	// - orgName
	// - teamName (may be empty)
	// - userName
	// - query string

	// 2️⃣ Call your DB/search functions
	// Example:
	// results := searchDocuments(orgName, teamName, userName, query)

	// 3️⃣ Return a placeholder response for now
	c.JSON(http.StatusOK, gin.H{
		"org":     orgName,
		"team":    teamName,
		"user":    userName,
		"query":   query,
		"results": []interface{}{}, // replace with actual search results
		"message": "Search parameters received successfully",
	})
}
