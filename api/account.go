package api

import (
	"database/sql"
	db "github.com/JCSong-89/go-udemy-master-class/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateAccountReq struct {
	Owner    string `json:"owner" binding:"required"`
	Balance  int64  `json:"balance" binding:"required,min=0"`
	Currency string `json:"currency" binding:"required, oneof=USD EUR"` // oneof는 주어진 값 중 하나여야 한다는 의미
}

/*
* 기본적으로 gin을 사용할 때 gin.Context를 사용하고 핸들러 내부에서 수행하는 모든 작업에는 이 컨텍스트 개체가 포함된다.
 */
func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountReq

	// gin.Context.BindJSON() 함수는 JSON 바디를 읽어서 input 구조체에 매핑하고, binding 지정된 필드의 유효섬 검사를 확인한다.
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// store.CreateAccount() 함수의 인수 타입은 db.CreateAccountParams이므로 변환한다.
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  req.Balance,
		Currency: req.Currency,
	}

	//계정생성로직
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
	return
}

type GetAccountRqe struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req GetAccountRqe

	// gin.ShouldBindUri는 URI 매개변수를 읽어서 input 구조체에 매핑한다.
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//계정조회로직
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		//sql.ErrNoRows는 데이터베이스에서 행을 찾을 수 없을 때 반환된다.
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type ListAccountReq struct {
	Limit  int32 `form:"limit" binding:"required,min=1"`
	Offset int32 `form:"offset" binding:"required,min=5, max=10"`
}

func (server *Server) listAccount(ctx *gin.Context) {
	var req ListAccountReq

	// gin.ShouldBindQuery()는 쿼리 매개변수를 읽어서 input 구조체에 매핑한다.
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit:  req.Limit,
		Offset: (req.Limit - 1) * req.Offset,
	}

	//계정목록조회로직
	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
