package imgprocessing

import "github.com/leductoan3082004/go-sdk/sdkcm"

type Response struct {
	sdkcm.AppError
	Data *sdkcm.Image `json:"data"`
}
