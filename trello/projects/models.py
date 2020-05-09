from django.db import models
from users.models import User
from django.contrib.postgres.fields import ArrayField


def get_default_statuses():
    return ['Backlog', 'InProgress', 'Done']


class ProjectBoard(models.Model):
    pbid = models.AutoField(primary_key=True)
    title = models.CharField(max_length=255)
    description = models.TextField(blank=True, null=True)
    archived = models.BooleanField(default=False)
    statuses = ArrayField(
        models.CharField(max_length=255),
        default=get_default_statuses
    )
    created_by = models.ForeignKey(
        'users.User',
        on_delete=models.SET_NULL,
        to_field='username',
        null=True,
        blank=True
    )
    timestamp = models.DateTimeField(auto_now_add=True)

    class Meta:
        app_label = 'projects'


class UserProjectRelation(models.Model):
    upid = models.AutoField(primary_key=True)
    pbid = models.ForeignKey(
        'projects.ProjectBoard',
        on_delete=models.CASCADE,
        db_index=True
    )
    user = models.ForeignKey(
        'users.User',
        on_delete=models.SET_NULL,
        to_field='username',
        null=True,
        blank=True
    )
    admin = models.BooleanField(default=False)
    active = models.BooleanField(default=True)

    class Meta:
        unique_together = ['pbid', 'user']
