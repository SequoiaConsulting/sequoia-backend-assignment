from django.db import models
from django.contrib.auth.models import (
    BaseUserManager, AbstractBaseUser, PermissionsMixin)
from django.core.validators import RegexValidator
# Create your models here.

alphanumeric = RegexValidator(
    r'^[0-9a-zA-Z]*$', 'Only alphanumeric characters are allowed.')


class UserManager(BaseUserManager):
    def create_user(self, username, name, email, password=None):
        if not name:
            raise ValueError("User must have a name")
        if not email:
            raise ValueError("User must have an email id")
        if not username:
            raise ValueError("User must have a username")
        if not password:
            raise ValueError("User must have a password")

        user = self.model(
            username=username,
            name=name,
            email=email,
        )
        user.set_password(password)
        user.save(using=self._db)
        return user

    def create_staffuser(self, username, name, email, password=None):
        user = self.create_user(
            username=username,
            name=name,
            email=email,
            password=password,
        )
        user.staff = True
        user.save(using=self._db)
        return user

    def create_superuser(self, username, name, email, password=None):
        user = self.create_user(
            username=username,
            name=name,
            email=email,
            password=password,
        )
        user.staff = True
        user.admin = True
        user.save(using=self._db)
        return user


class User(AbstractBaseUser, PermissionsMixin):
    uid = models.AutoField(primary_key=True)
    email = models.EmailField(max_length=255, unique=True, db_index=True)
    username = models.CharField(
        max_length=255,
        unique=True,
        validators=[alphanumeric],
        db_index=True
    )
    name = models.CharField(max_length=255)
    active = models.BooleanField(default=True)
    staff = models.BooleanField(default=False)
    admin = models.BooleanField(default=False)

    USERNAME_FIELD = 'username'
    REQUIRED_FIELDS = ['name', 'email']

    objects = UserManager()

    class Meta:
        app_label = 'users'

    def get_name(self):
        return self.name

    def get_email(self):
        return self.email

    def get_username(self):
        return self.username

    def has_perm(self, perm, obj=None):
        return True

    def has_module_perms(self, app_label):
        return True

    @property
    def is_staff(self):
        return self.staff

    @property
    def is_admin(self):
        return self.admin

    @property
    def is_active(self):
        return self.active
