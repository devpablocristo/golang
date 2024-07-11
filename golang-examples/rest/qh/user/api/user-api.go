package api

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	domain "github.com/devpablocristo/golang/06-projects/qh/user/domain"
)

func LoadData(wg *sync.WaitGroup) {

	u1 := domain.User{
		UUID:     "1",
		Username: "HSimp",
		Password: "12345",
		Email:    "homero@simpson.com",
	}

	u2 := domain.User{
		UUID:     "2",
		Username: "MSimp",
		Password: "12345",
		Email:    "marge@simpson.com",
	}

	users := []domain.User{
		u1,
		u2,
	}

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}

func StartApi(wg *sync.WaitGroup, port string) {
	defer wg.Done()

	//gotenv.Load()
	// db := postgres.ConnectToDB()
	// defer db.Close()
	// mdb := mapdb.NewMapDB()
	// sdb := slicedb.NewSliceDB()
	// pse := application.NewPersonService(mdb)
	// pse := application.NewPersonService(sdb)
	// han := chihandler.NewChiHandler(pse)
	// rou := SetupChiRoutes(han)

	han := GinRouter()
	GinServer(port, han) //, rou)
}

func logErrors(err error) {
	if err != nil {
		log.Fatal("ERROR!!! ", err)
	}
}
