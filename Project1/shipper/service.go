package shipper

import (
	"Project1/shipper/model"
	"Project1/shipper/util"
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"gitlab.com/tolexo/aqua"
	"gopkg.in/go-playground/validator.v9"
	//"reflect
	"strconv"
	"strings"
	//"time"
	//"log"
)

var db *pg.DB

func Connect() {
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "postgres",
		Password: "123456",
		Addr:     "localhost:5432",
	})
}
func Close() {
	db.Close()
}

type Shipper struct {
	aqua.RestService `prefix:"catalog" root:"/shipper" version:"1"`
	createShipper    aqua.POST `url:"/"`
	updateShipper    aqua.PUT  `url:"/"`
	listShipper      aqua.GET  `url:"/"`
	//getShipper       aqua.GET    `url:"/{shipper_id:[0-9]+}/"`
	deleteShipper aqua.DELETE `url:"/{shipperIDs}/"`
}

/*func (*Shipper) CreateShipper(req aqua.Aide) (httpCode int, data interface{}) {

	//VCreateShipper()
	//PCreateShipper()
	var (
		shipperReq model.ShipperReq
		tx         *pg.Tx
		err        error
		shipperID  int
	)
	req.LoadVars()
	if err = json.Unmarshal([]byte(req.Body), &shipperReq); err == nil {
		V := validator.New()
		if err = V.Struct(shipperReq); err == nil {
			fmt.Println("hello")
			Connect()
			defer Close()
			if tx, err = db.Begin(); err == nil {
				fmt.Println("hello")
				if shipperID, err = util.InsertQueryForshipper(tx, shipperReq); err == nil {
					err = util.InsertQueryForaddress(tx, shipperReq.SAddress, shipperID)
				}
				if err == nil {
					tx.Commit()
					data = strconv.Itoa(shipperID)
					httpCode = 200
				} else {
					tx.Rollback()
				}
			}
		}
	}
	if err != nil {
		data = err.Error()
		httpCode = 500
	}
	//fmt.Println("here", httpCode)
	//err := db.CreateTable(&model.ShipperDB{}, nil)
	return
}*/

func (*Shipper) CreateShipper(req aqua.Aide) (httpCode int, data interface{}) {

	//VCreateShipper()
	//PCreateShipper()
	var (
		shipperReq model.ShipperReq
	)
	var err error
	if shipperReq, err = util.VCreateShipper(req); err == nil {
		Connect()
		defer Close()
		if err = util.PCreateShipper(req, db, shipperReq); err == nil {
			data = strconv.Itoa(util.ShipperID)
			httpCode = 200
		} else {
			fmt.Println(err.Error())
		}
	}
	if err != nil {
		data = err.Error()
		httpCode = 500
	}
	return
}

func (*Shipper) ListShipper(req aqua.Aide) (httpCode int, data interface{}) {
	Connect()
	defer Close()
	var (
		sp []model.Temp
		//sa  []model.AddressDB
		err error
	)

	//qry := `select s.shipper_id,a.address_id from shipper s join address a on s.shipper_id=a.address_id;`
	qry := `select s.shipper_id, s.name, s.mobile, s.description,
	json_agg(
		jsonb_build_object(
			'id', a.address_id,
			'city', a.city,
			'state', a.state,
			'address', a.address
		)
	) address_info
	from address a 
	join shipper_address sa on a.address_id=sa.fkaddress_id 
	join shipper s on s.shipper_id=sa.fkshipper_id
	group by s.shipper_id;`

	//err := db.Model(&model.ShipperDB{}).Column("shipper_id").Select(&sp)

	if _, err = db.Query(&sp, qry); err == nil {
		data = sp
		httpCode = 200
	}
	/*if err = db.Model(&model.ShipperDB{}).Select(&sp); err == nil {

		if err = db.Model(&model.AddressDB{}).Select(&sa); err == nil {
			data = sp
			httpCode = 200
		}
	}*/
	if err != nil {
		data = err.Error()
		httpCode = 500
	}
	return
}

func (*Shipper) UpdateShipper(req aqua.Aide) (httpCode int, data interface{}) {

	var (
		err error
		up  model.UpdateReq1
		qry *orm.Query
	)
	req.LoadVars()
	if err = json.Unmarshal([]byte(req.Body), &up); err == nil {
		fmt.Println("hello")
		V := validator.New()
		if err = V.Struct(up); err == nil {
			fmt.Println("hey")
			Connect()
			defer Close()
			qry = db.Model(&model.ShipperDB{}).Set("name = ?", up.Name)
			if up.Email != "" {
				qry = qry.Set("email = ?", up.Email)
			}
			if up.Mobile != "" {
				qry = qry.Set("mobile = ?", up.Mobile)
			}
			_, err = qry.Where("Shipper_id=?", up.ShipperID).Update()
		}
	}

	/*if _, err := db.Model(&[]model.ShipperDB{}).Set("Name=?", "abc").WhereIn("Shipper_id IN ?", []int{23, 24, 25}).Returning("*").Update(); err == nil {
		if _, err := db.Model(&[]model.AddressDB{}).Set("City=?", "Ahmedabad").WhereIn("Address_id IN ?", []int{1, 2, 3}).Returning("*").Update(); err != nil {
			panic(err)
		}
	}*/

	//var s1 []model.ShipperDB
	//err = db.Model(&[]model.ShipperDB{}).Select(&s1)
	if err != nil {
		httpCode = 500
		data = err.Error()
	} else {
		data = "Shipper Updated"
	}
	httpCode = 200
	return
}

func (*Shipper) DeleteShipper(shipperIDReq string, req aqua.Aide) (httpCode int, data interface{}) {
	Connect()
	defer Close()

	var (
		err        error
		shipperIDs []int
		shipperID  int
		res        orm.Result
	)
	shIDReqParam := strings.Split(shipperIDReq, ",")
	for _, shIDString := range shIDReqParam {
		if shIDString != "" {
			if shipperID, err = strconv.Atoi(shIDString); err == nil {
				shipperIDs = append(shipperIDs, shipperID)
			} else {
				break
			}
		}
	}
	if err == nil {
		if len(shipperIDs) > 0 {
			//res, err := db.Model(&[]model.ShipperDB{}).WhereIn("Shipper_id IN ?", []int{4, 5, 6}).Delete()
			if res, err = db.Model(&model.ShipperDB{}).WhereIn("Shipper_id IN ?", shipperIDs).Delete(); err == nil {
				//fmt.Println(reflect.ValueOf(res))

				//fmt.Println("deleted", res.RowsAffected())

				data = "No. of deleted shippers: " + strconv.Itoa(res.RowsAffected())

				//data = strconv.Itoa(count)
			}
		} else {
			data = "No. of deleted shippers: 0"
		}
	}
	if err != nil {
		data = err.Error()
		httpCode = 500
	} else {
		httpCode = 200
		//fmt.Println("left", count)
	}
	return
}
