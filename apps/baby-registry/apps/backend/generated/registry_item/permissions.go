package registry_item;

import (
    "github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
    "github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_role"
    "github.com/butterfeetlabs/baby-registry/apps/backend/generated/coded_error"
)




func GetAbacProjection(actor permissions.Actor) Projection {
    return Projection{
    }
}

func HasWritePermissions(m Model, actor permissions.Actor) bool {
    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        case enum_role.Owner: 
            // Actor has full permissions
            return true
        case enum_role.Super: 
            // Actor has full permissions
            return true
        }
    }
    return false
}

func ProjectReadPermissions(p Projection, actor permissions.Actor) Projection {
    return p
}

func ProjectWritePermissions(p Projection, actor permissions.Actor) Projection {
    return p
}

// Permissions on query



func ApplyActorReadPermissionsToWhereClause(actor permissions.Actor, query WhereClause) (WhereClause, error) {
    

    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        case enum_role.Owner: 
            // Actor has full permissions
            return query, nil
        case enum_role.Super: 
            // Actor has full permissions
            return query, nil
        }
    }
    return query, coded_error.NewUnauthorizedError()

}

func ApplyActorWritePermissionsToWhereClause(actor permissions.Actor, query WhereClause) (WhereClause, error) {
    

    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        case enum_role.Owner: 
            // Actor has full permissions
            return query, nil
        case enum_role.Super: 
            // Actor has full permissions
            return query, nil
        }
    }
    return query, coded_error.NewUnauthorizedError()

}
