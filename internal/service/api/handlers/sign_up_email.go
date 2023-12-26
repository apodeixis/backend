package handlers

import (
	"net/http"
	"strings"

	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/helpers"
	"github.com/apodeixis/backend/internal/service/api/requests"
	"github.com/apodeixis/backend/internal/service/api/responses"
	"github.com/apodeixis/backend/internal/types"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func SignUpEmail(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)
	request, err := requests.NewEmail(r)
	if err != nil {
		log.WithError(err).Error("invalid request")
		errObjects := problems.BadRequest(err)
		if strings.Contains(err.Error(), is.ErrEmail.Error()) {
			if len(errObjects) == 1 {
				errObjects[0].Code = types.ErrInvalidEmail
			}
		}
		ape.RenderErr(w, errObjects...)
		return
	}
	user, err := ctx.UsersQ(r).New().FilterByEmail(request.Data.Attributes.Email).Get()
	if err != nil {
		log.Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if user == nil {
		log.Error("not found user")
		ape.RenderErr(w, problems.NotFound(types.ErrNotFoundUser))
		return
	}
	if user.EmailVerified {
		log.Warn("user has already verified email, skipping send email")
		ape.RenderErr(w, problems.UnprocessableEntity(types.ErrEmailAlreadyVerified))
		return
	}
	if err := helpers.SendConfirmationEmail(user, r); err != nil {
		log.Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	response := responses.ComposeEmail(user.Email)
	ape.Render(w, response)
}
