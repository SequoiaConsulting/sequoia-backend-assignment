from rest_framework.permissions import BasePermission


class Permission(BasePermission):

    def has_permission(self, request, view):

        if (request.method in view.PUB):
            return True
        elif ((request.method in view.CON) and (request.auth)):
            return True
        elif ((request.method in view.SEC) and (request.user.is_staff)):
            return True
        elif ((request.auth) and (request.user.is_admin)):
            return True
        else:
            return False
