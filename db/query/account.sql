/*Account 단일 계정 객체 (one) return  */
-- name: CreateAccount :one
INSERT INTO  authors (
    owner,
    balance,
    currency
) VALUES (
    $1, $2, $3
) RETURNING *; -- $1, $2, $3 은 인자로 받은 값을 사용한다.
