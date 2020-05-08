from users.models import User
from users.serializers import (
    UserSerializer, UserSerializerCreate, TokenSerializer)
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.parsers import MultiPartParser, FormParser
from rest_framework import status
from django.shortcuts import get_object_or_404
from django.contrib.auth import authenticate
from rest_framework_jwt.settings import api_settings
from trello.permissions import Permission
from rest_framework.parsers import JSONParser
from trello.api_exceptions import CustomApiException

jwt_payload_handler = api_settings.JWT_PAYLOAD_HANDLER
jwt_encode_handler = api_settings.JWT_ENCODE_HANDLER


class UserView(APIView):
    parser_classes = (MultiPartParser, FormParser)

    PUB = ['POST']
    CON = []
    SEC = []

    permission_classes = (Permission,)

    def post(self, request, *args, **kwargs):
        data = JSONParser().parse(request)
      
        serializer = UserSerializerCreate(data=data,
                                          context={'request': request})
        if serializer.is_valid():
            user = serializer.save()
            user.set_password(serializer.validated_data['password'])
            user.save()

            token_serializer = TokenSerializer(data={
                "token": jwt_encode_handler(
                    jwt_payload_handler(user)
                )})
            token_serializer.is_valid()
            return Response(token_serializer.data, status=status.HTTP_200_OK)

        raise CustomApiException(400, serializer.errors)


class UserDetail(APIView):
    parser_classes = (MultiPartParser, FormParser)

    PUB = []
    CON = ['GET']
    SEC = []

    permission_classes = (Permission,)

    def get(self, request, username):
        if(request.user.username == username or request.user.staff):
            user = get_object_or_404(User, username=username)
            return Response(
                UserSerializer(user).data, status=status.HTTP_200_OK)

        raise CustomApiException(403, "Permission Denied")


class UserExist(APIView):
    parser_classes = (MultiPartParser, FormParser)

    PUB = ['GET']
    CON = []
    SEC = []

    permission_classes = (Permission,)

    def get(self, request, username):
        user = get_object_or_404(User, username=username)
        data = {
            'username': user.username
        }
        return Response(data, status=status.HTTP_200_OK)


class LoginView(APIView):
    parser_classes = (MultiPartParser, FormParser)

    PUB = ['POST']
    CON = []
    SEC = []

    permission_classes = (Permission,)

    def post(self, request, *args, **kwargs):

        data = JSONParser().parse(request)
        try:
            user = User.objects.get(username=data["username"])
        except:
            try:
                user = User.objects.get(email=data["username"])
            except:
                raise CustomApiException(404, "Invalid Username or Password")
        try:
            user = authenticate(username=user.username,
                                password=data['password'])
        except:
            raise CustomApiException(404, "Invalid Username or Password")

        if user is not None:
            serializer = TokenSerializer(data={
                "token": jwt_encode_handler(
                    jwt_payload_handler(user)
                )})
            serializer.is_valid()
            return Response(serializer.data, status=status.HTTP_200_OK)
        raise CustomApiException(401, "Credentials Mismatch")
