package registry_approved_guest;

import (
    "github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
    "github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_role"
    "github.com/butterfeetlabs/baby-registry/apps/backend/generated/coded_error"
    "fmt"
)



func (m *Model) GetOwnerId() string {
    return m.OwnerId
}


func GetAbacProjection(actor permissions.Actor) Projection {
    return Projection{
        OwnerId: true,
    }
}

func HasWritePermissions(m Model, actor permissions.Actor) bool {
    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        case enum_role.Owner: 
            hasPerm := true
            hasPerm = hasPerm && m.GetOwnerId() == actorRole.OwnerId
            if hasPerm {
                return true
            }
            break
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
     
    ownerIdIn := make(map[string]struct{})

    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        case enum_role.Owner: 
            if actorRole.OwnerId == "" {
                return query, coded_error.NewUnauthorizedError("OwnerId is required for role Owner")
            }
            ownerIdIn[actorRole.OwnerId] = struct{}{}
            break
        case enum_role.Super: 
            // Actor has full permissions
            return query, nil
        }
    }
    if len(ownerIdIn) == 0 {
        // If actor does not have granted permissions, return error
        return query, coded_error.NewUnauthorizedError("actor has no permissions")
    }
     
    if err := func() error {
        // If no query, return
        if len(ownerIdIn) == 0 {
            return nil
        }
        // If Eq the query, check if the actor has permissions
        if query.OwnerIdEq != nil {
            if _, ok := ownerIdIn[*query.OwnerIdEq]; !ok {
                return coded_error.NewUnauthorizedError(fmt.Sprintf("actor %s has no permissions to OwnerId %v", actor.GetActorId(), *query.OwnerIdEq))
            }
            return nil
        }
        // If filtering on a range of values, ensure the actor has permissions for all values
        if query.OwnerIdIn != nil {
           l := *query.OwnerIdIn
           var noPerms []string
           for _, v := range l {
                if _, ok := ownerIdIn[v]; !ok {
                    noPerms = append(noPerms, v)
                }
           }
           if len(noPerms) == 0 {
                return coded_error.NewUnauthorizedError(fmt.Sprintf("actor %s has no permissions to OwnerId %v", actor.GetActorId(), noPerms))
           }
           return nil
        }
        // Only include values the actor has permissions for
        l := make([]string, 0, len(ownerIdIn))
        for k := range ownerIdIn {
            l = append(l, k)
        }
        query.OwnerIdIn = &l
        return nil
    }(); err != nil {
        return query, err
    }
    
    return query, nil


}

func ApplyActorWritePermissionsToWhereClause(actor permissions.Actor, query WhereClause) (WhereClause, error) {
     
    ownerIdIn := make(map[string]struct{})

    for _, actorRole := range actor.GetActorRoles() {
        switch actorRole.Role {
        case enum_role.Owner: 
            if actorRole.OwnerId == "" {
                return query, coded_error.NewUnauthorizedError("OwnerId is required for role Owner")
            }
            ownerIdIn[actorRole.OwnerId] = struct{}{}
            break
        case enum_role.Super: 
            // Actor has full permissions
            return query, nil
        }
    }
    if len(ownerIdIn) == 0 {
        // If actor does not have granted permissions, return error
        return query, coded_error.NewUnauthorizedError("actor has no permissions")
    }
     
    if err := func() error {
        // If no query, return
        if len(ownerIdIn) == 0 {
            return nil
        }
        // If Eq the query, check if the actor has permissions
        if query.OwnerIdEq != nil {
            if _, ok := ownerIdIn[*query.OwnerIdEq]; !ok {
                return coded_error.NewUnauthorizedError(fmt.Sprintf("actor %s has no permissions to OwnerId %v", actor.GetActorId(), *query.OwnerIdEq))
            }
            return nil
        }
        // If filtering on a range of values, ensure the actor has permissions for all values
        if query.OwnerIdIn != nil {
           l := *query.OwnerIdIn
           var noPerms []string
           for _, v := range l {
                if _, ok := ownerIdIn[v]; !ok {
                    noPerms = append(noPerms, v)
                }
           }
           if len(noPerms) == 0 {
                return coded_error.NewUnauthorizedError(fmt.Sprintf("actor %s has no permissions to OwnerId %v", actor.GetActorId(), noPerms))
           }
           return nil
        }
        // Only include values the actor has permissions for
        l := make([]string, 0, len(ownerIdIn))
        for k := range ownerIdIn {
            l = append(l, k)
        }
        query.OwnerIdIn = &l
        return nil
    }(); err != nil {
        return query, err
    }
    
    return query, nil


}
