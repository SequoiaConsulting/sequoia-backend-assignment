from django.contrib import admin
from django.contrib.auth import get_user_model
from .models import User
from django.contrib.auth.admin import UserAdmin as BaseUserAdmin
from django.contrib.auth.forms import UserCreationForm
from django import forms

# Register your models here.


class MyUserAdmin(BaseUserAdmin):
    list_display = ('uid', 'username', 'name',
                    'email', 'admin',)
    list_filter = ('admin',)
    fieldsets = (
        (None, {'fields': ('username', 'password')}),
        ('Personal info', {'fields': ('name', 'email',)}),
        ('Permissions', {'fields': ('admin', 'staff', 'active')}),
    )
    add_fieldsets = (
        (None, {
            'classes': ('wide',),
            'fields': ('name', 'username', 'email', 'password1', 'password2')}),
    )
    ordering = ('uid',)
    filter_horizontal = ()


admin.site.register(User, MyUserAdmin)