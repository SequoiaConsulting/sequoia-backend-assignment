from django.contrib import admin
from tasks.models import Task


class TaskAdmin(admin.ModelAdmin):
    list_display = (
        'tid', 'pbid', 'title', 'description', 'assignee', 'assigner',
        'due_date', 'creation_timestamp')


admin.site.register(Task, TaskAdmin)
