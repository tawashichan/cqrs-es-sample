package bookModel

import "cqrs-es/domain/event/bookEvent"

type IBookEventRepository interface {
	BookRental(rental bookEvent.BookRental) error
}
