package controllers

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gorilla/mux"

	"bob-bank/utils"
	"bob-bank/models"
	"bob-bank/validations"
)

var (
	ErrInvalidCash = errors.New("Valor transferido é inválido")
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := models.GetTransactions()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, transactions)
}

func PostTransaction(w http.ResponseWriter, r *http.Request) {
	transaction, err := verifyTransaction(r)
	if err != nil{
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = models.NewTransaction(transaction)
	if err != nil{
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{"Tansação concluída com sucesso!", http.StatusCreated})
}

func verifyTransaction(r *http.Request) (models.Transaction, error) {
	//recebe a chave pública que irá receber o valor
	params := mux.Vars(r)
	targetKey := params["public_key"]
	target, err := models.GetWalletByPublicKey(targetKey)
	if err != nil {
		return models.Transaction{}, err
	}
	//recebe o json da carteira que irá enviar um valor com saldo e a chave pública
	body, _ := ioutil.ReadAll(r.Body)
	var origin models.Wallet
	err = json.Unmarshal(body, &origin)
	if err != nil {
		return models.Transaction{}, err
	}
	//verifica se a carteira existe
	originVerify, err := models.GetWalletByPublicKey(origin.PublicKey)
	if err != nil {
		return models.Transaction{}, err
	}
	if validations.IsEmpty(target.PublicKey) || validations.IsEmpty(originVerify.PublicKey) {
		return models.Transaction{}, models.ErrWalletNotFound
	}
	//verifica se o saldo que será transferido é maior que o saldo da carteira ou menor que zero
	if origin.Balance > originVerify.Balance || origin.Balance < 0 {
		return models.Transaction{}, ErrInvalidCash
	}
	var transaction models.Transaction
	transaction.Cash = origin.Balance
	transaction.Message = fmt.Sprintf("%s transferiu %.2f $, para %s", originVerify.User.Nickname, origin.Balance, target.User.Nickname)
	transaction.Origin = origin
	transaction.Target = target
	return transaction, nil
}