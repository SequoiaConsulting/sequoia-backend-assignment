from datetime import datetime
from rest_framework_jwt.settings import api_settings
import time


def jwt_payload_handler(user):
    time_now = int(time.time())
    return {
        'uid': user.pk,
        'username': user.username,
        'name': user.name,
        'email': user.email,
        'staff': user.staff,
        'exp': datetime.utcnow() + api_settings.JWT_EXPIRATION_DELTA,
        'orig_iat': time_now
    }
