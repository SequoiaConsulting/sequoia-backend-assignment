from users.models import User
from rest_framework import serializers
from django.contrib.auth.password_validation import validate_password
from rest_framework.authtoken.models import Token


class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = (
            'uid',
            'name',
            'username',
            'email',
            'active',
            'staff',
            'admin',
        )


class UserSerializerCreate(serializers.ModelSerializer):

    class Meta:
        model = User
        fields = ('username', 'name', 'email', 'password')

    @staticmethod
    def validate_password(password):
        """
        Validate password
        """

        validate_password(password)
        return password


class UserSerializerInfo(UserSerializer):

    class Meta:
        model = User
        fields = (
            'uid',
            'name',
            'username',
            'email',
        )


class TokenSerializer(serializers.Serializer):
    """
    This serializer serializes the token data
    """
    token = serializers.CharField(max_length=255)

