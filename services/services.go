package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"recipe-api/errors"
	"recipe-api/models"
	"recipe-api/utils"
	"strings"
)

var recipePuppyBaseUrl = "http://www.recipepuppy.com/api/?i="
var giphyBaseUrl = "http://api.giphy.com/v1/gifs/search?q=%s&api_key=%s&limit=1"
var apiKeyGiphy = utils.EnvVar("GIPHY_API_KEY")

func GetRecipes(ingredients []string) ([]models.Recipe, error) {
	puppyResponse, err := GetRecipesFromPuppy(ingredients)
	if err != nil {
		return nil, err
	}
	var recipes []models.Recipe
	for _, p := range puppyResponse.Results {
		gifUrl, err := GetGif(p.Title)
		if err != nil {
			return []models.Recipe{}, err
		}
		recipe := models.Recipe{
			Titulo:       strings.TrimSpace(p.Title),
			Ingredientes: strings.Split(p.Ingredients, ", "),
			ReceitaURL:   p.Href,
			ReceitaGif:   gifUrl,
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil

}

func GetRecipesFromPuppy(terms []string) (models.PuppyResponse, error) {
	url := recipePuppyBaseUrl + strings.Join(terms, ",")
	res, err := http.Get(url)
	if err != nil {
		return models.PuppyResponse{}, errors.ErroPuppy{}
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.PuppyResponse{}, errors.ErroPuppy{}
	}
	puppyResponse := models.PuppyResponse{}
	err = json.Unmarshal(bodyBytes, &puppyResponse)
	if err != nil {
		return models.PuppyResponse{}, errors.ErroPuppy{}
	}
	return puppyResponse, nil
}

func GetGif(title string) (string, error) {
	title = strings.Replace(strings.TrimSpace(title), " ", "+", -1)
	url := fmt.Sprintf(giphyBaseUrl, title, apiKeyGiphy)
	res, err := http.Get(url)
	if err != nil {
		return "", errors.ErroGiphy{}
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.ErroGiphy{}
	}
	giphyResponse := models.GiphyResponse{}
	err = json.Unmarshal(bodyBytes, &giphyResponse)
	if err != nil {
		return "", errors.ErroGiphy{}
	}
	gifUrl := giphyResponse.Data[0].Images.Original.URL
	gifUrl = strings.Split(gifUrl, "?")[0]
	return gifUrl, nil
}
