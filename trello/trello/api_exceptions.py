from rest_framework.exceptions import APIException
from rest_framework.views import exception_handler
from rest_framework.exceptions import Throttled


class CustomApiException(APIException):

    # public fields
    detail = None
    status_code = None

    # create constructor
    def __init__(self, status_code, message):
        # override public fields
        CustomApiException.status_code = status_code
        CustomApiException.detail = message


def custom_exception_handler(exc, context):
    response = exception_handler(exc, context)

    if isinstance(exc, Throttled):
        custom_response_data = {
            'message': 'Request limit exceeded'
        }
        response.data = custom_response_data

    return response
