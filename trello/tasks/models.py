from django.db import models
from users.models import User
from projects.models import ProjectBoard
from trello.api_exceptions import CustomApiException


class Task(models.Model):
    tid = models.AutoField(primary_key=True)
    pbid = models.ForeignKey(
        'projects.ProjectBoard',
        on_delete=models.CASCADE,
        db_index=True
    )
    title = models.CharField(max_length=255)
    description = models.TextField(blank=True, null=True)
    assignee = models.ForeignKey(
        'users.User',
        on_delete=models.SET_NULL,
        to_field='username',
        related_name='assignee',
        null=True,
        blank=True
    )
    assigner = models.ForeignKey(
        'users.User',
        on_delete=models.SET_NULL,
        to_field='username',
        related_name='assigner',
        null=True,
        blank=True
    )
    due_date = models.DateTimeField(blank=True, null=True)
    status = models.CharField(max_length=255)
    creation_timestamp = models.DateTimeField(auto_now_add=True)

    def save(self, *args, **kwargs):
        pbid = self.pbid.pbid
        project = ProjectBoard.objects.get(pbid=pbid)
        print(self.status)
        if(self.status not in project.statuses):
            error_msg = self.status + (" status doesn't belong to this "
                                       "project's statuses")
            raise CustomApiException(400, error_msg)

        super(Task, self).save(*args, **kwargs)
