package services

var (
	ItemsService itemsServiceInterface=&itemsService{}
)
type itemsService struct{}

type itemsServiceInterface interface{
	GetItem()
	SaveItem()
}

func (*itemsService) GetItem() {

}

func (*itemsService) SaveItem(){

}