package GoYandexDisk

import (
	"encoding/json"
	"net/http"
)

const baseURL = "https://cloud-api.yandex.net/v1/disk/"

//YandexDisk Структура для работы с диском
type YandexDisk struct {
	accessToken string
}

func ConnectDisk(token string) YandexDisk{
	return YandexDisk{token}
}

func (yd YandexDisk)DiskStatus() (Disk, error){
	resp, err := yd.createGETRequest("")
	if err != nil {
		return Disk{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return Disk{}, err
	}

	var result Disk
	err = json.Unmarshal(body, &result)
	return result, err
}

//Get Получение ресурса списка ресурсов
func (yd YandexDisk)Get(path string, limit int, offset int) (ResourceList, error) {
	link := "resources?path=" + path + "&limit=" + string(limit) + "&offset=" + string(offset)
	resp, err := yd.createGETRequest(link)
	if err != nil {
		return ResourceList{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return ResourceList{}, err
	}

	var result ResourceList
	err = json.Unmarshal(body, &result)
	return result, err
}

//FlatList Получение плоского списка файлов
func (yd YandexDisk)FlatList(limit int, mediatype []string, offset int) (FileResourceList, error) {
	//TODO добавить в запрос mediatype
	link := "resources/files?limit=" + string(limit) + "&offset=" + string(offset)
	resp, err := yd.createGETRequest(link)
	if err != nil {
		return FileResourceList{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return FileResourceList{}, err
	}

	var result FileResourceList
	err = json.Unmarshal(body, &result)
	return result, err
}

//LastUploaded получение списка последних загруженных файлов
func (yd YandexDisk)LastUploaded(limit int, mediatype []string) (LastUploadedResourceList, error) {
	//TODO добавить в запрос mediatype
	link := "resources/files?limit=" + string(limit)
	resp, err := yd.createGETRequest(link)
	if err != nil {
		return LastUploadedResourceList{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return LastUploadedResourceList{}, err
	}

	var result LastUploadedResourceList
	err = json.Unmarshal(body, &result)
	return result, err
}

//UploadFile
//func (yd YandexDisk)UploadFile() {}

//UploadFileFromNet Загрузка файла на диск по ссылке из интернета
func (yd YandexDisk)UploadFileFromNet(url string, path string, fields []string, disableRedirects bool) (Link, error){
	link := "resources/upload?url=" + url + "?path=" + path
	resp, err := yd.createPOSTRequest(link)

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return Link{}, err
	}

	var result Link
	err = json.Unmarshal(body, &result)
	return result, err
}

//DownloadFile вернет объект Link со ссылкой для загрузки
func (yd YandexDisk)DownloadFile(path string) (Link, error){
	link := "resources/download?path=" + path
	resp, err := yd.createGETRequest(link)
	if err != nil {
		return Link{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return Link{}, err
	}

	var result Link
	err = json.Unmarshal(body, &result)
	return result, err
}

//CopyResource копирование ресурса
func (yd YandexDisk)CopyResource(from string, path string, owerwrite bool) (Link, error){
	link := "resources/copy?from=" + from + "&path=" + path
	resp, err := yd.createGETRequest(link)
	if err != nil {
		return Link{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return Link{}, err
	}

	var result Link
	err = json.Unmarshal(body, &result)
	return result, err
}

//ReplaceResource перемещение ресурса
func (yd YandexDisk)ReplaceResource(from string, path string, owerwrite bool) (Link, error){
	link := "resources/move?from=" + from + "&path=" + path
	resp, err := yd.createGETRequest(link)
	if err != nil {
		return Link{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return Link{}, err
	}

	var result Link
	err = json.Unmarshal(body, &result)
	return result, err
}

//DeleteResource удаление ресурса
//func (yd YandexDisk)DeleteResource(path string, permanently bool) (Link, error){
//	//TODO дописать
//	link := "resources?path=" + path
//	resp, err := yd.createGETRequest(link)
//	if err != nil {
//		return Link{}, err
//	}
//
//	var body []byte
//	_, err = resp.Body.Read(body)
//	if err != nil {
//		return Link{}, err
//	}
//
//	var result Link
//	err = json.Unmarshal(body, &result)
//	return result, err
//}

//func (yd YandexDisk)CreateFolder(path string) (Link, error){
//	link := "resources/move?path=" + path
//	resp, err := yd.createGETRequest(link)
//	if err != nil {
//		return Link{}, err
//	}
//
//	var body []byte
//	_, err = resp.Body.Read(body)
//	if err != nil {
//		return Link{}, err
//	}
//
//	var result Link
//	err = json.Unmarshal(body, &result)
//	return result, err
//}

//PublishResource публикация ресурса
func (yd YandexDisk)PublishResource(path string) (Link, error){
	link := "resources/publish?path=" + path
	resp, err := yd.createGETRequest(link)
	if err != nil {
		return Link{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return Link{}, err
	}

	var result Link
	err = json.Unmarshal(body, &result)
	return result, err
}

//UnpublishResource отказаться от публикации ресурса
func (yd YandexDisk)UnpublishResource(publickey string) (Resource, error){
	//TODO доделать функцию
	link := "public/resources" + publickey
	resp, err := yd.createGETRequest(link)
	if err != nil {
		return Resource{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return Resource{}, err
	}

	var result Resource
	err = json.Unmarshal(body, &result)
	return result, err
}

//PublicResourceMeta получение мета информации о публичном ресурсе
func (yd YandexDisk)PublicResourceMeta(publikKey string) (Resource, error){
	link := "public/resources?public_key=" + publikKey
	resp, err := yd.createGETRequest(link)
	if err != nil {
		return Resource{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return Resource{}, err
	}

	var result Resource
	err = json.Unmarshal(body, &result)
	return result, err
}

//DownloadPublicResource получение ссылки на зугрузку публичного ресурса
func (yd YandexDisk)DownloadPublicResource(publikKey string) (Link, error){
	link := "public/resources/download?public_key=" + publikKey
	resp, err := yd.createGETRequest(link)
	if err != nil {
		return Link{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return Link{}, err
	}

	var result Link
	err = json.Unmarshal(body, &result)
	return result, err
}

//SavePublicResource сохранение ресурса в папку загрузки
func (yd YandexDisk)SavePublicResource(publikKey string) (Link, error){
	link := "resources/download?public_key=" + publikKey
	resp, err := yd.createPOSTRequest(link)
	if err != nil {
		return Link{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return Link{}, err
	}

	var result Link
	err = json.Unmarshal(body, &result)
	return result, err
}

//PublicResourcesMeta получение списка публичных файлов
func (yd YandexDisk)PublishResourcesList(limit int, offset int) (PublicResourcesList, error){
	link := "resources/public?limit=" + string(limit) + "&offset=" + string(offset)
	resp, err := yd.createGETRequest(link)
	if err != nil {
		return PublicResourcesList{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return PublicResourcesList{}, err
	}

	var result PublicResourcesList
	err = json.Unmarshal(body, &result)
	return result, err
}

//CleanTrash Очистка файла из корзины, если path = "" происходит полная очистка
func (yd YandexDisk)CleanTrash(path string) (Link, error){
	link := "trash/resources?path=" + path
	resp, err := yd.createGETRequest(link)
	if err != nil {
		return Link{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return Link{}, err
	}

	var result Link
	err = json.Unmarshal(body, &result)
	return result, err
}

//RestorePublicResource восстановление ранее удаленного ресурса
//func (yd YandexDisk)RestorePublicResource() {}

//Status возвращает текущий статус выполнения операции
func (yd YandexDisk)Status(id string) (Link, error){
	link := "operations/" + id
	resp, err := yd.createGETRequest(link)
	if err != nil {
		return Link{}, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		return Link{}, err
	}

	var result Link
	err = json.Unmarshal(body, &result)
	return result, err
}

func (ya YandexDisk)createGETRequest(link string) (*http.Response, error){
	r := http.Request{}
	r.Header.Add("Authorization", ya.accessToken)
	return http.Get(baseURL + link)
}

func (ya YandexDisk)createPOSTRequest(link string) (*http.Response, error){
	r := http.Request{}
	r.Header.Add("Authorization", ya.accessToken)
	return http.Post(baseURL + link, "", nil)
}

func encodeLink(link string) string{
	//TODO Описать тело функции
	return link
}