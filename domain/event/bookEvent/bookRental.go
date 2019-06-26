package bookEvent

import "cqrs-es/domain/model/common"

type BookRental struct {
	UserId  common.UserUUID
	BookId  common.BookUUId
	DueDate common.Date
}
