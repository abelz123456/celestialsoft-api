package mapping

import (
	"github.com/abelz123456/celestial-api/api/auth/domain"
	"github.com/abelz123456/celestial-api/entity"
)

func ToPermissionPolicyUserResponse(permissionPolicyUser entity.PermissionPolicyUser) domain.PermissionPolicyUserVm {
	return domain.PermissionPolicyUserVm{
		Oid:       permissionPolicyUser.Oid,
		EmailName: permissionPolicyUser.EmailName,
		Password:  permissionPolicyUser.Password,
	}
}

func ToPermissionPolicyUserResponses(permissionPolicyUsers []entity.PermissionPolicyUser) []domain.PermissionPolicyUserVm {
	var permissionPolicyUserResponses []domain.PermissionPolicyUserVm
	for _, permissionPolicyUser := range permissionPolicyUsers {
		permissionPolicyUserResponses = append(permissionPolicyUserResponses, ToPermissionPolicyUserResponse(permissionPolicyUser))
	}
	return permissionPolicyUserResponses
}

func ToPermissionPolicyUserResponseAuth(permissionPolicyUser entity.PermissionPolicyUser,
	token string, refreshToken string, role string) domain.PermissionPolicyUserAuthVm {
	return domain.PermissionPolicyUserAuthVm{
		Oid:       permissionPolicyUser.Oid,
		EmailName: permissionPolicyUser.EmailName,
		// Password:    permissionPolicyUser.Password,
		Token:        &token,
		RefreshToken: &refreshToken,
		Role:         &role,
	}
}
