package root

import (
	"context"
	"errors"
	"net/http"

	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/ent/adjacenttable"
	"github.com/0B1t322/Magic-Circle/ent/profile"
	"github.com/0B1t322/Magic-Circle/ent/sector"
	"github.com/0B1t322/Magic-Circle/models/role"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	InstituteNotFound = errors.New("Institute not Found")
	InstituteExist    = errors.New("Institute with this nmae exist")

	ProfileNotFound = errors.New("Profile not found")

	DirectionNotFound     = errors.New("Direction not found")
	CantCreateRelate      = errors.New("Can't create relate with sectors because you should to give inst/dir/prof")
	UnexpectedBody        = errors.New("Unexpected body")
	ErrSectorNotFound     = errors.New("Sector not found")
	ErrAdjacentTableExist = errors.New("Adjecent table exist")

	DirectionWithOutInst = errors.New("Can't create direction without institute")
	ProfWithOutDir       = errors.New("Can't create profile without direction")
)

func newLogFields(method string, err error) log.Fields {
	return log.Fields{
		"controller": "root",
		"method":     method,
		"err":        err,
	}
}

type RootController struct {
	Client *ent.Client
}

func New(client *ent.Client) *RootController {
	return &RootController{
		Client: client,
	}
}

/*
{
	institute: {
		name: "some_name"
	},

}
*/

type Sectors struct {
	Coords []string `json:"coords"`
}

type GetOrCreateReq struct {
	ID        *int `json:"id,omitempty" extenstions:"x-nullable"`
	CreateReq `json:",inline"`
}

type CreateReq struct {
	Name *string `json:"name,omitempty" extenstion:"x-nullable"`
}

type CreateInstDirProf struct {
	Inst    *GetOrCreateReq `json:"institute"`
	Dir     *GetOrCreateReq `json:"direction"`
	Prof    *GetOrCreateReq `json:"profile"`
	Sectors *Sectors        `json:"sectors,omitempty"`
}

func (r RootController) getOrCreateInst(ctx context.Context, inst *GetOrCreateReq) (*ent.Institute, error) {
	// Если реквест пустой то ничего не делаем
	if inst == nil {
		return nil, nil
	}

	if inst.ID != nil {
		get, err := r.Client.Institute.Get(ctx, *inst.ID)
		if ent.IsNotFound(err) {
			return nil, InstituteNotFound
		} else if err != nil {
			return nil, err
		}

		return get, nil
	} else if inst.Name != nil {
		create, err := r.Client.Institute.Create().
			SetName(*inst.Name).
			Save(ctx)

		if ent.IsConstraintError(err) {
			return nil, InstituteExist
		} else if err != nil {
			return nil, err
		}

		return create, nil
	}
	// Если пустое оба поля то ошибка ибо непонятно что хотят сделать
	return nil, nil
}

func (r RootController) getOrCreateDir(ctx context.Context, dir *GetOrCreateReq, inst *ent.Institute) (*ent.Direction, error) {
	if dir == nil {
		return nil, nil
	}

	if dir.ID != nil {
		get, err := r.Client.Direction.Get(ctx, *dir.ID)
		if ent.IsNotFound(err) {
			return nil, DirectionNotFound
		} else if err != nil {
			return nil, err
		}

		return get, nil
	} else if dir.Name != nil {
		if inst == nil {
			return nil, DirectionWithOutInst
		}
		create, err := r.Client.Direction.Create().
			SetName(*dir.Name).SetInstitute(inst).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		return create, nil
	}
	// Если пустое возможно имелось ввиду создания института поэтому ошибки нет
	return nil, nil
}

func (r RootController) getOrCreateProfile(
	ctx context.Context,
	prof *GetOrCreateReq,
	dir *ent.Direction,
) (*ent.Profile, error) {
	if prof == nil {
		return nil, nil
	}
	
	if prof.ID != nil {
		get, err := r.Client.Profile.Get(ctx, *prof.ID)
		if ent.IsNotFound(err) {
			return nil, ProfileNotFound
		} else if err != nil {
			return nil, err
		}

		return get, nil
	} else if prof.Name != nil {
		if dir == nil {
			return nil, ProfWithOutDir
		}
		create, err := r.Client.Profile.Create().SetName(*prof.Name).SetDirection(dir).Save(ctx)
		if err != nil {
			return nil, err
		}

		return create, nil
	}

	return nil, nil
}

// CreateInstDirProfile
//
// @Summary create institute or direction or profile
//
// @Description to create only institute you need to put into body only name of institute according to schema
//
// @Description to create some relation you need to put to institute id and put into direction name
//
// @Param body body root.CreateInstDirProf true "body"
//
// @Accept json
//
// @Router /v1/ [post]
//
// @Security ApiKeyAuth
// 
// @Produce json
//
// @Success 201 {object} root.CreateInstDirProfResp
//
// @Failure 500
//
// @Failure 404 {string} string
//
// @Failure 400 {string} string
func (r RootController) CreateInstDirProf(c *gin.Context) {
	var req CreateInstDirProf
	{
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "Failed to read body")
			c.Abort()
			return
		}
	}

	var (
		inst *ent.Institute
		dir  *ent.Direction
		prof *ent.Profile

		resp CreateInstDirProfResp
		err  error
	)

	claims := jwt.ExtractClaims(c)
	// Check that user is admin and if its admin and they try to create intitite or pass not their institute
	// Pass error
	if claims["role"].(string) == string(role.ADMIN) && req.Inst != nil {
		if req.Inst.ID == nil || float64(*req.Inst.ID) != claims["intstituteId"].(float64){
			c.String(http.StatusForbidden, "You are not superadmin or admin of this institute")
			c.Abort()
			return
		}
	}

	// Если создание института будет достаточно указания этого в реквесте
	inst, err = r.getOrCreateInst(c, req.Inst)
	if err == InstituteNotFound {
		c.String(http.StatusNotFound, err.Error())
		c.Abort()
		return
	} else if err == InstituteExist {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	} else if err == UnexpectedBody {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("CreateInstDirProf", err)).Error()
		c.Status(http.StatusInternalServerError)
		return
	}

	// если институт был создан или получен
	if inst != nil {
		resp.Institute = &CreatedInstitute{
			ID:   inst.ID,
			Name: inst.Name,
		}
	}

	// Check that if user is admin and try pass direction this dir is in this inst
	if claims["role"].(string) == string(role.ADMIN) && req.Dir != nil {
		if req.Dir.ID != nil {
			dirId := *req.Dir.ID
			getDir, err := r.Client.Direction.Get(c, dirId)
			if ent.IsNotFound(err) {
				// Pass because error would be in create method
			} else if err != nil {
				log.WithFields(newLogFields("CreateInstDirProf", err)).Error()
				c.Status(http.StatusInternalServerError)
				c.Abort()
				return
			}

			if float64(getDir.InstituteID) != claims["intstituteId"].(float64) {
				c.String(http.StatusForbidden, "You are not superadmin or admin of this institute")
				c.Abort()
				return
			}
		}
	}

	dir, err = r.getOrCreateDir(
		c,
		req.Dir,
		inst,
	)
	if err == DirectionNotFound {
		c.String(http.StatusNotFound, err.Error())
		c.Abort()
		return
	} else if err == DirectionWithOutInst {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("CreateInstDirProf", err)).Error()
		c.Status(http.StatusInternalServerError)
		c.Abort()
		return
	}

	if dir != nil {
		resp.Direction = &CreatedDirection{
			ID:   dir.ID,
			Name: dir.Name,
		}
	}

	// Check that if user is admin and try pass profile and this profile not in this inst
	if claims["role"].(string) == string(role.ADMIN) && req.Prof != nil {
		if req.Prof.ID != nil {
			profId := *req.Prof.ID
			getProf, err := r.Client.Profile.Query().
				WithDirection().
				Where(
					profile.ID(profId),
				).
				Only(c)
			if ent.IsNotFound(err) {
				// Pass
			} else if err != nil {
				log.WithFields(newLogFields("CreateInstDirProf", err)).Error()
				c.Status(http.StatusInternalServerError)
				c.Abort()
				return
			}

			if float64(getProf.Edges.Direction.InstituteID) != claims["intstituteId"].(float64) {
				c.String(http.StatusForbidden, "You are not superadmin or admin of this institute")
				c.Abort()
				return
			}
		}
	}

	prof, err = r.getOrCreateProfile(
		c,
		req.Prof,
		dir,
	)
	if err == ProfWithOutDir {
		c.String(http.StatusNotFound, err.Error())
		c.Abort()
		return
	} else if err == ProfileNotFound {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("CreateInstDirProf", err)).Error()
		c.Status(http.StatusInternalServerError)
		c.Abort()
		return
	}

	if prof != nil {
		resp.Profile = &CreatedProfile{
			ID:   prof.ID,
			Name: prof.Name,
		}
	}

	if req.Sectors != nil {
		if prof == nil {
			c.String(http.StatusBadRequest, CantCreateRelate.Error())
			c.Abort()
			return
		}

		sectors, err := r.getSectors(c, req.Sectors.Coords)
		if err == ErrSectorNotFound {
			c.String(http.StatusNotFound, err.Error())
			c.Abort()
			return
		} else if err != nil {
			log.WithFields(newLogFields("CreateInstDirProf", err)).Error()
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		}

		_, err = r.relateWithSectors(
			c,
			prof,
			sectors,
		)

		if err == ErrAdjacentTableExist {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		} else if err != nil {
			log.WithFields(newLogFields("CreateInstDirProf", err)).Error()
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		}

	}

	c.JSON(
		http.StatusCreated,
		resp,
	)
}

func (r RootController) getSectors(ctx context.Context, coords []string) ([]*ent.Sector, error) {
	s, err := r.Client.Sector.Query().Where(
		sector.CoordsIn(coords...),
	).All(ctx)

	if ent.IsNotFound(err) || len(s) != len(coords) {
		return nil, ErrSectorNotFound
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

func (r RootController) relateWithSectors(
	ctx context.Context,
	prof *ent.Profile,
	s []*ent.Sector,
) ([]*ent.AdjacentTable, error) {
	ids := func(s []*ent.Sector) (ids []int) {
		for _, sect := range s {
			ids = append(ids, sect.ID)
		}
		return ids
	}(s)

	get, err := r.Client.AdjacentTable.Query().Where(
		adjacenttable.HasSectorWith(
			sector.IDIn(ids...),
		),
		adjacenttable.HasProfileWith(
			profile.ID(prof.ID),
		),
	).All(ctx)
	if err != nil {
		return nil, err
	}

	if len(get) > 0 {
		return nil, ErrAdjacentTableExist
	}

	var bulk []*ent.AdjacentTableCreate
	{
		for _, sect := range s {
			bulk = append(bulk, r.Client.AdjacentTable.Create().SetSector(sect).SetProfile(prof))
		}
	}

	return r.Client.AdjacentTable.CreateBulk(bulk...).Save(ctx)
}

type CreatedInstitute struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

type CreatedDirection struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

type CreatedProfile struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

type RelateSectors struct {
	Coords []string `json:"coords"`
}

type CreateInstDirProfResp struct {
	Institute     *CreatedInstitute `json:"institute,omitempty"`
	Direction     *CreatedDirection `json:"direction,omitempty"`
	Profile       *CreatedProfile   `json:"profile,omitempty"`
	// CreatedRelate *RelateSectors    `json:"relateSectors,omitempty"`
}

type DeleteRelateReq struct {
	ProfileID int     `json:"profile_id"`
	Sectors   Sectors `json:"sectors"`
}


// DeleteRelate
//
// @Summary delete relate between profile and sectors
//
// @Param body body root.DeleteRelateReq true "body"
//
// @Accept json
//
// @Router /v1/ [delete]
//
// @Security ApiKeyAuth
// 
// @Produce json
//
// @Success 200
//
// @Failure 500
//
// @Failure 404 {string} string
//
// @Failure 400 {string} string
func (r RootController) DeleteRelate(c *gin.Context) {
	var req DeleteRelateReq
	{
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpected body")
			c.Abort()
			return
		}
	}

	claims := jwt.ExtractClaims(c)

	if claims["role"].(string) == string(role.ADMIN) {
		getProf, err := r.Client.Profile.Query().
				WithDirection().
				Where(
					profile.ID(req.ProfileID),
				).
				Only(c)
		if ent.IsNotFound(err) {
			// Pass
		} else if err != nil {
			log.WithFields(newLogFields("DeleteRelate", err)).Error()
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		}

		if float64(getProf.Edges.Direction.InstituteID) != claims["intstituteId"].(float64) {
			c.String(http.StatusForbidden, "You are not superadmin or admin of this institute")
			c.Abort()
			return
		}
	}

	if deleteCount, err := r.Client.AdjacentTable.Delete().Where(
		adjacenttable.And(
			adjacenttable.HasProfileWith(
				profile.ID(req.ProfileID),
			),
			adjacenttable.HasSectorWith(
				sector.CoordsIn(req.Sectors.Coords...),
			),
		),
	).Exec(c); ent.IsNotFound(err) || deleteCount == 0 {
		c.String(http.StatusNotFound, "Don't find sectors that relate with this profile")
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("DeleteRelate", err)).Error()
		c.Status(http.StatusInternalServerError)
		c.Abort()
		return
	}
}
