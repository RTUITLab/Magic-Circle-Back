package direction

import (
	"context"
	"net/http"

	"github.com/0B1t322/Magic-Circle/ent"
	. "github.com/0B1t322/Magic-Circle/models/direction"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func newLogFields(method string, err error) log.Fields {
	return log.Fields{
		"controller": "direction",
		"method": method,
		"err": err,
	}
}

type DirectionController struct {
	Client *ent.Client
}

func New(client *ent.Client) *DirectionController {
	return &DirectionController{
		Client: client,
	}
}

func (d DirectionController) getAll(ctx context.Context) ([]*ent.Direction, error) {
	return d.Client.Direction.Query().All(ctx)
}

type GetDirectionsReq struct {

}

type GetDirectionsResp struct {
	Dirs	[]Direction		`json:"directions"`
}

// GetAll
// 
// @Summary Get all directions
// 
// @Description return all directions
// 
// @Router /v1/direction [get]
// 
// @Produce json
// 
// @Success 200 {object} direction.GetDirectionsResp
// 
// @Failure 500 {string} srting
func (d DirectionController) GetAll(c *gin.Context) {
	ds, err := d.getAll(c)
	if err != nil {
		log.WithFields(newLogFields("GetAll", err)).Error("Failed to get institutes")
		c.String(http.StatusInternalServerError, "Failed to get institutes")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, GetDirectionsResp{Dirs: DirectionsFromEnt(ds)})
}