from django.contrib import admin
from projects.models import ProjectBoard, UserProjectRelation


class ProjectBoardAdmin(admin.ModelAdmin):
    list_display = (
        'pbid', 'title', 'archived', 'statuses', 'created_by', 'timestamp')


class UserProjectRelationAdmin(admin.ModelAdmin):
    list_display = (
        'upid', 'pbid', 'user', 'admin', 'active')


admin.site.register(ProjectBoard, ProjectBoardAdmin)
admin.site.register(UserProjectRelation, UserProjectRelationAdmin)
