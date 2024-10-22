package user

import "context"

func (u *useCases) findUserByCuit(ctx context.Context, cuit string) (bool, error) {
	person, err := u.personRepo.FindByCuit(ctx, cuit)
	if err != nil {
		return false, err
	}
	if person != nil {
		return true, nil
	}

	company, err := u.companyRepo.FindByCuit(ctx, cuit)
	if err != nil {
		return false, nil
	}
	if company != nil {
		return true, nil
	}

	return false, nil
}

func (u *useCases) findAdministrativeRequestByCuit(ctx context.Context, cuit string) (bool, error) {
	// consulta ms administrative requests

	//return false, nil
	return true, nil
}
