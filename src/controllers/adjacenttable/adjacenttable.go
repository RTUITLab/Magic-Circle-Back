package adjacenttable

import (
	"context"
	"errors"
	"net/http"

	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/ent/adjacenttable"
	"github.com/0B1t322/Magic-Circle/ent/direction"
	"github.com/0B1t322/Magic-Circle/ent/institute"
	"github.com/0B1t322/Magic-Circle/ent/profile"
	"github.com/0B1t322/Magic-Circle/ent/sector"
	"github.com/0B1t322/Magic-Circle/ent/variant"
	model "github.com/0B1t322/Magic-Circle/models/adjacenttable"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	ErrSectorNotFound     = errors.New("Sector not found")
	ErrInstituteNotFound  = errors.New("Institute not found")
	ErrDirectionNotFound  = errors.New("Direction not found")
	ErrProfileNotFound    = errors.New("Profile not found")
	ErrAdjacentTableExist = errors.New("Adjacent Table with this variant and sector exist")
)

func newLogFields(method string, err error) log.Fields {
	return log.Fields{
		"controller": "adjacenttable",
		"method":     method,
		"err":        err,
	}
}

type AdjacentTableController struct {
	Client *ent.Client
}

type AdjacentTable = model.AdjacentTable

func New(client *ent.Client) *AdjacentTableController {
	return &AdjacentTableController{
		Client: client,
	}
}

type CreateSectorReq struct {
	Coords      string  `json:"coords"`
	Description *string `json:"description,omitempty" extensions:"x-nullable"`
}

type CreateVariantReq struct {
	Institute string `json:"instituteName"`
	Direction string `json:"directionName"`
	Profile   string `json:"profileName"`
}

type CreateAdjacentTableReq struct {
	Sector 		CreateSectorReq  `json:"sector"`
	CreateVariantReq `json:",inline"`
}

type CreateAdjacentTableResp struct {
	AdjacentTable `json:",inline"`
}

func (a AdjacentTableController) create(
	ctx context.Context,
	req CreateAdjacentTableReq,
) (*ent.AdjacentTable, error) {
	s, err := a.createOrGetSector(ctx, req.Sector)
	if err != nil {
		return nil, err
	}

	i, err := a.createOrGetInstitute(ctx, req.Institute)
	if err != nil {
		return nil, err
	}

	d, err := a.createOrGetDirection(ctx, req.Direction)
	if err != nil {
		return nil, err
	}

	p, err := a.createOrGetProfile(ctx, req.Profile)
	if err != nil {
		return nil, err
	}

	v, err := a.createOrGetVariant(
		ctx,
		i,
		d,
		p,
	)
	if err != nil {
		return nil, err
	}

	created, err := a.createAdjacentTable(ctx, s, v)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (a AdjacentTableController) createALot(
	ctx 	context.Context,
	req		CreateAdjacentTablesReq,
) ([]*ent.AdjacentTable, error) {
	s, err := a.getSectors(ctx, req.Sectors)
	if err != nil {
		return nil, err
	}

	i, err := a.createOrGetInstitute(ctx, req.Institute)
	if err != nil {
		return nil, err
	}

	d, err := a.createOrGetDirection(ctx, req.Direction)
	if err != nil {
		return nil, err
	}

	p, err := a.createOrGetProfile(ctx, req.Profile)
	if err != nil {
		return nil, err
	}

	v, err := a.createOrGetVariant(
		ctx,
		i,
		d,
		p,
	)
	if err != nil {
		return nil, err
	}

	created, err := a.createAdjacentTables(ctx, s, v)
	if err != nil {
		return nil, err
	}

	return created, nil	
}

func (a AdjacentTableController) getSectors(ctx context.Context, coords []string) ([]*ent.Sector, error) {
	s, err := a.Client.Sector.Query().Where(
		sector.CoordsIn(coords...),
	).All(ctx)

	if ent.IsNotFound(err) {
		return nil, ErrSectorNotFound
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

// if exist retturn error
func (a AdjacentTableController) createAdjacentTable(
	ctx context.Context,
	s *ent.Sector,
	v *ent.Variant,
) (*ent.AdjacentTable, error) {
	_, err := a.Client.AdjacentTable.Query().Where(
		adjacenttable.HasSectorWith(
			sector.ID(s.ID),
		),
		adjacenttable.HasVariantWith(
			variant.ID(v.ID),
		),
	).Only(ctx)
	if ent.IsNotFound(err) {
		// Pass
	} else if err != nil {
		return nil, err
	} else {
		return nil, ErrAdjacentTableExist
	}

	return a.Client.AdjacentTable.Create().SetSector(s).SetVariant(v).Save(ctx)
}

// if exist retturn error
func (a AdjacentTableController) createAdjacentTables(
	ctx context.Context,
	s []*ent.Sector,
	v *ent.Variant,
) ([]*ent.AdjacentTable, error) {
	ids := func(s []*ent.Sector) (ids []int) {
		for _, sect := range s {
			ids = append(ids, sect.ID)
		}
		return ids
	}(s)

	get, err := a.Client.AdjacentTable.Query().Where(
		adjacenttable.HasSectorWith(
			sector.IDIn(ids...),
		),
		adjacenttable.HasVariantWith(
			variant.ID(v.ID),
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
			bulk = append(bulk, a.Client.AdjacentTable.Create().SetSector(sect).SetVariant(v))
		}
	}


	return a.Client.AdjacentTable.CreateBulk(bulk...).Save(ctx)
}

// Create
//
// @Summary Create
//
// @Description Create adjacent table
// 
// @Description you can create sector with this method just add description and coords to sector field
// 
// @Description also you can just add coords fields and they will find sector
// 
// @Description this endpoint also can get or create institute/profile/direction by name, because all names in this object is unique string
// 
// @Description if adjacent table with this sector and variant exist return bad request
//
// @Router /v1/adjacenttable [post]
//
// @Accept json
//
// @Produce json
//
// @Param body body adjacenttable.CreateAdjacentTableReq true "body"
//
// @Success 201 {object} adjacenttable.CreateAdjacentTableResp
//
// @Failure 400 {string} srting
//
// @Failure 500 {string} string
func (a AdjacentTableController) Create(c *gin.Context) {
	var req CreateAdjacentTableReq
	{
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpeted body")
			c.Abort()
			return
		}
	}

	created, err := a.create(c, req)
	if err == ErrAdjacentTableExist {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("Create", err)).Error("Failed to create adjecent table")
		c.String(http.StatusInternalServerError, "Failed to create adjecent table")
		c.Abort()
		return
	}

	get, err := a.Client.AdjacentTable.Query().Where(
		adjacenttable.ID(created.ID),
	).
	WithSector().
	WithVariant(
		func(vq *ent.VariantQuery) {
			vq.WithInsitute().
			WithDirection().
			WithProfile()
		},
	).Only(c)
	if err != nil {
		log.WithFields(newLogFields("Create", err)).Error("Failed to create adjecent table")
		c.String(http.StatusInternalServerError, "Failed to create adjecent table")
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, CreateAdjacentTableResp{AdjacentTable: model.AdjacentTableFromEnt(get)})
}


type CreateAdjacentTablesReq struct {
	Sectors 		[]string  `json:"sectors"`
	CreateVariantReq `json:",inline"`
}

type CreateAdjacentTablesResp struct {
	AdjacentTables []AdjacentTable `json:"adjacentTables"`
}

// Create
//
// @Summary Create
//
// @Description Create adjacent tables
// 
// @Description this method create or institute/profile/direction but require created sector in array
// 
// @Descrription also if adjacent table with this sector and variant exist return Bad request
// 
// @Description if adjacent table with this sector and variant exist return bad request
//
// @Router /v1/adjacenttables [post]
//
// @Accept json
//
// @Produce json
//
// @Param body body adjacenttable.CreateAdjacentTablesReq true "body"
//
// @Success 201 {object} adjacenttable.CreateAdjacentTablesResp
//
// @Failure 400 {string} srting
// 
// @Failure 404 {string} srting
//
// @Failure 500 {string} string
func (a AdjacentTableController) CreateALot(c *gin.Context) {
	var req CreateAdjacentTablesReq
	{
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpeted body")
			c.Abort()
			return
		}
	}

	created, err := a.createALot(c, req)
	if err == ErrAdjacentTableExist {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	} else if err == ErrSectorNotFound {
		c.String(http.StatusNotFound, err.Error())
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("CreateALot", err)).Error("Failed to create adjecent table")
		c.String(http.StatusInternalServerError, "Failed to create adjecent table")
		c.Abort()
		return
	}

	ids := func(as []*ent.AdjacentTable) (slice []int) {
		for _, a := range as {
			slice = append(slice, a.ID)
		}
		return slice
	}(created)

	get, err := a.Client.AdjacentTable.Query().Where(
		adjacenttable.IDIn(ids...),
	).
	WithSector().
	WithVariant(
		func(vq *ent.VariantQuery) {
			vq.WithInsitute().
			WithDirection().
			WithProfile()
		},
	).All(c)
	if err != nil {
		log.WithFields(newLogFields("CreateALot", err)).Error("Failed to create adjecent table")
		c.String(http.StatusInternalServerError, "Failed to create adjecent table")
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, CreateAdjacentTablesResp{AdjacentTables: model.AdjacentTablesFromEnt(get)})
}

func (a AdjacentTableController) createOrGetVariant(
	ctx context.Context,
	i *ent.Institute,
	d *ent.Direction,
	p *ent.Profile,
) (*ent.Variant, error) {
	v, err := a.Client.Variant.Query().Where(
		variant.HasDirectionWith(
			direction.ID(d.ID),
		),
		variant.HasInsituteWith(
			institute.ID(i.ID),
		),
		variant.HasProfileWith(
			profile.ID(p.ID),
		),
	).Only(ctx)

	if ent.IsNotFound(err) {
		v, err := a.Client.Variant.Create().
			SetDirection(d).
			SetInsitute(i).
			SetProfile(p).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return v, nil
	} else if err != nil {
		return nil, err
	}

	return v, nil
}

func (a AdjacentTableController) createOrGetSector(
	ctx context.Context,
	req CreateSectorReq,
) (*ent.Sector, error) {
	s, err := a.getSector(ctx, req.Coords)
	if err == ErrSectorNotFound {

		s, err = a.createSector(ctx, req)
		if err != nil {
			return nil, err
		}

		return s, nil

	} else if err != nil {
		return nil, err
	}

	return s, nil
}

func (a AdjacentTableController) getSector(
	ctx context.Context,
	coords string,
) (*ent.Sector, error) {
	get, err := a.Client.Sector.Query().Where(
		sector.Coords(coords),
	).Only(ctx)
	if ent.IsNotFound(err) {
		return nil, ErrSectorNotFound
	} else if err != nil {
		return nil, err
	}

	return get, nil
}

func (a AdjacentTableController) createSector(
	ctx context.Context,
	req CreateSectorReq,
) (*ent.Sector, error) {
	create, err := a.Client.Sector.Create().
		SetDescription(*req.Description).
		SetCoords(req.Coords).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return create, nil
}

func (a AdjacentTableController) createOrGetInstitute(
	ctx context.Context,
	name string,
) (*ent.Institute, error) {
	i, err := a.getInstitute(ctx, name)
	if err == ErrInstituteNotFound {
		i, err = a.createInstitute(ctx, name)
		if err != nil {
			return nil, err
		}
		return i, nil
	} else if err != nil {
		return nil, err
	}

	return i, nil
}

func (a AdjacentTableController) getInstitute(
	ctx context.Context,
	name string,
) (*ent.Institute, error) {
	i, err := a.Client.Institute.Query().Where(
		institute.Name(name),
	).Only(ctx)
	if ent.IsNotFound(err) {
		return nil, ErrInstituteNotFound
	} else if err != nil {
		return nil, err
	}

	return i, nil
}

func (a AdjacentTableController) createInstitute(
	ctx context.Context,
	name string,
) (*ent.Institute, error) {
	i, err := a.Client.Institute.Create().
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (a AdjacentTableController) createOrGetProfile(
	ctx context.Context,
	name string,
) (*ent.Profile, error) {
	i, err := a.getProfile(ctx, name)
	if err == ErrProfileNotFound {
		i, err = a.createProfile(ctx, name)
		if err != nil {
			return nil, err
		}
		return i, nil
	} else if err != nil {
		return nil, err
	}

	return i, nil
}

func (a AdjacentTableController) getProfile(
	ctx context.Context,
	name string,
) (*ent.Profile, error) {
	i, err := a.Client.Profile.Query().Where(
		profile.Name(name),
	).Only(ctx)
	if ent.IsNotFound(err) {
		return nil, ErrProfileNotFound
	} else if err != nil {
		return nil, err
	}

	return i, nil
}

func (a AdjacentTableController) createProfile(
	ctx context.Context,
	name string,
) (*ent.Profile, error) {
	i, err := a.Client.Profile.Create().
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (a AdjacentTableController) createOrGetDirection(
	ctx context.Context,
	name string,
) (*ent.Direction, error) {
	i, err := a.getDirection(ctx, name)
	if err == ErrDirectionNotFound {
		i, err = a.createDirection(ctx, name)
		if err != nil {
			return nil, err
		}
		return i, nil
	} else if err != nil {
		return nil, err
	}

	return i, nil
}

func (a AdjacentTableController) getDirection(
	ctx context.Context,
	name string,
) (*ent.Direction, error) {
	i, err := a.Client.Direction.Query().Where(
		direction.Name(name),
	).Only(ctx)
	if ent.IsNotFound(err) {
		return nil, ErrDirectionNotFound
	} else if err != nil {
		return nil, err
	}

	return i, nil
}

func (a AdjacentTableController) createDirection(
	ctx context.Context,
	name string,
) (*ent.Direction, error) {
	i, err := a.Client.Direction.Create().
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return i, nil
}