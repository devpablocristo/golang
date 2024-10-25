package user

import "context"

func (u *useCases) findUserByCuit(ctx context.Context, cuil string) (bool, error) {
	person, err := u.personUseCases.FindPersonByCuil(ctx, cuil)
	if err != nil {
		return false, err
	}
	if person != nil {
		return true, nil
	}

	return false, nil
}

func (u *useCases) findAdministrativeRequestByCuit(ctx context.Context, cuil string) (bool, error) {
	// consulta ms administrative requests

	//return false, nil
	return true, nil
}
